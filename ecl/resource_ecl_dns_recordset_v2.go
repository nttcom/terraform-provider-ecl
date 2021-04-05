package ecl

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/dns/v2/recordsets"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceDNSRecordSetV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSRecordSetV2Create,
		Read:   resourceDNSRecordSetV2Read,
		Update: resourceDNSRecordSetV2Update,
		Delete: resourceDNSRecordSetV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"zone_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"A", "AAAA", "MX", "CNAME", "SRV",
					"SPF", "TXT", "PTR", "NS",
				}, true),
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"record": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDNSRecordSetV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	records := []string{d.Get("record").(string)}
	createOpts := RecordSetCreateOpts{
		recordsets.CreateOpts{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Records:     records,
			TTL:         d.Get("ttl").(int),
			Type:        d.Get("type").(string),
		},
	}

	zoneID := d.Get("zone_id").(string)

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	n, err := recordsets.Create(dnsClient, zoneID, createOpts).ExtractCreatedRecordSet()

	log.Printf("[DEBUG] DNS RecordSet Creation response: %#v", n)
	log.Printf("[DEBUG] DNS RecordSet Creation response(err): %#v", err)

	if err != nil {
		return fmt.Errorf("Error creating ECL DNS record set: %s", err)
	}

	id := fmt.Sprintf("%s/%s", zoneID, n.ID)
	d.SetId(id)

	// ECL2.0 DNS API returns 404(Not Found) just after creation.
	interval := 3 * time.Second
	maxLoop := 10
	for i := 0; i < maxLoop; i++ {
		log.Printf("[DEBUG] Trying to retrieve DNS record set as %d th try.", i)
		log.Printf("[DEBUG] Zone ID: %s , RecordSet ID: %s", zoneID, n.ID)

		_, err := recordsets.Get(dnsClient, zoneID, n.ID).Extract()
		if err != nil {
			_, ok := err.(eclcloud.ErrDefault404)
			if ok {
				log.Printf("[DEBUG] Waiting for DNS record set (%s) is created.", n.ID)
				time.Sleep(interval)
				continue
			} else {
				return fmt.Errorf("Some error occurred in getting ECL DNS record set: %s", err)
			}
		}
		break
	}

	log.Printf("[DEBUG] Created ECL DNS record set %s: %#v", n.ID, n)
	return resourceDNSRecordSetV2Read(d, meta)
}

func resourceDNSRecordSetV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	// Obtain relevant info from parsing the ID
	zoneID, recordsetID, err := parseDNSV2RecordSetID(d.Id())
	if err != nil {
		return err
	}

	n, err := recordsets.Get(dnsClient, zoneID, recordsetID).Extract()
	if err != nil {
		return CheckDeleted(d, err, "record_set")
	}

	records := n.Records.([]interface{})
	record := records[0].(string)

	d.Set("name", n.Name)
	d.Set("description", n.Description)
	d.Set("ttl", n.TTL)
	d.Set("type", n.Type)
	d.Set("record", record)
	d.Set("zone_id", zoneID)

	return nil
}

func resourceDNSRecordSetV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	var updateOpts recordsets.UpdateOpts

	// Note:
	// In ECL2.0 DNS RecordSet API, returns error like following,
	// if only send "description" or only send blank json as PUT request body.
	// Error -> E2044 At least domain name or TTL or recordset value is required.
	// So in RecordSet case, all parameter you can pudate will be sent as request body.
	// This is why this resource does not use "d.HasChange(<param name>)".
	name := d.Get("name").(string)
	updateOpts.Name = &name

	ttl := d.Get("ttl").(int)
	updateOpts.TTL = &ttl

	records := []string{d.Get("record").(string)}
	updateOpts.Records = &records

	description := d.Get("description").(string)
	updateOpts.Description = &description

	// Obtain relevant info from parsing the ID
	zoneID, recordsetID, err := parseDNSV2RecordSetID(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating  record set %s with options: %#v", recordsetID, updateOpts)

	_, err = recordsets.Update(dnsClient, zoneID, recordsetID, updateOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error updating ECL DNS record set: %s", err)
	}

	// ECL2.0 DNS API returns 404(Not Found) just after creation.
	interval := 3 * time.Second
	maxLoop := 10
	for i := 0; i < maxLoop; i++ {
		log.Printf("[DEBUG] Trying to retrieve DNS record set as %d th try.", i)
		log.Printf("[DEBUG] Zone ID: %s , RecordSet ID: %s", zoneID, recordsetID)

		rs, err := recordsets.Get(dnsClient, zoneID, recordsetID).Extract()
		if err != nil {
			return fmt.Errorf("Some error occurred in getting ECL DNS record set: %s", err)
		}
		if !isRecordSetUpdated(d, rs) {
			time.Sleep(interval)
			continue
		}
		break
	}
	return resourceDNSRecordSetV2Read(d, meta)
}

func isRecordSetUpdated(d *schema.ResourceData, rs *recordsets.RecordSet) bool {
	if d.HasChange("ttl") && d.Get("ttl").(int) != rs.TTL {
		log.Printf("[DEBUG] TTL still does not match.")
		return false
	}

	if d.HasChange("description") && d.Get("description").(string) != rs.Description {
		log.Printf("[DEBUG] Description still does not match.")
		return false
	}

	if d.HasChange("record") {
		records := []string{d.Get("record").(string)}
		realRecord := rs.Records.([]interface{})[0].(string)
		if records[0] != realRecord {
			log.Printf("[DEBUG] Record still does not match.")
			return false
		}
	}
	return true
}

func resourceDNSRecordSetV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	dnsClient, err := config.dnsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	// Obtain relevant info from parsing the ID
	zoneID, recordsetID, err := parseDNSV2RecordSetID(d.Id())
	if err != nil {
		return err
	}

	err = recordsets.Delete(dnsClient, zoneID, recordsetID).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL DNS record set: %s", err)
	}

	d.SetId("")
	return nil
}

func parseDNSV2RecordSetID(id string) (string, string, error) {
	idParts := strings.Split(id, "/")
	if len(idParts) != 2 {
		return "", "", fmt.Errorf("Unable to determine DNS record set ID from raw ID: %s", id)
	}

	zoneID := idParts[0]
	recordsetID := idParts[1]

	return zoneID, recordsetID, nil
}

func formatDNSV2Records(recordsraw []interface{}) []string {
	records := make([]string, len(recordsraw))

	// Strip out any [ ] characters in the address.
	// This is to format IPv6 records in a way that DNSaaS / Designate wants.
	re := regexp.MustCompile("[][]")
	for i, recordraw := range recordsraw {
		record := recordraw.(string)
		record = re.ReplaceAllString(record, "")
		records[i] = record
	}

	return records
}

// suppressRecordsDiffs will suppress diffs when the format of a record
// is different yet still a valid DNS record. For example, if a user
// specifies an IPv6 address using [bracket] notation, but the record is
// returned without brackets, it is still a valid record.
func suppressRecordsDiffs(k, old, new string, d *schema.ResourceData) bool {
	re := regexp.MustCompile("[][]")
	new = re.ReplaceAllString(new, "")

	if old == new {
		return true
	}

	return false
}
