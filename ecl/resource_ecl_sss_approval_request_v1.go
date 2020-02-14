package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	"github.com/nttcom/eclcloud/ecl/sss/v1/approval_requests"
)

func resourceSSSApprovalRequestV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSSApprovalRequestV1Create,
		Read:   resourceSSSApprovalRequestV1Read,
		Delete: resourceSSSApprovalRequestV1Delete,
		Schema: map[string]*schema.Schema{
			"request_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_request_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"approver_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"approver_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"actions": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"api_path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"method": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"body": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"descriptions": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lang": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"text": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"request_user": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"approver": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"approval_deadline": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"approval_expire": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"registered_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"approved", "denied", "cancelled",
				}, false),
			},
		},
	}
}

func resourceSSSApprovalRequestV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %w", err)
	}

	opts := approval_requests.UpdateOpts{Status: d.Get("status").(string)}
	log.Printf("[DEBUG] Update Options: %#v", opts)

	approval, err := approval_requests.Update(client, d.Get("request_id").(string), opts).Extract()
	if err != nil {
		return fmt.Errorf("error updating ECL approval request: %w", err)
	}

	d.SetId(d.Get("request_id").(string))

	log.Printf("[DEBUG] Updated ECL approval request %s: %#v", approval.RequestID, approval)
	return resourceSSSApprovalRequestV1Read(d, meta)
}

func resourceSSSApprovalRequestV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("error creating ECL sss client: %w", err)
	}

	approval, err := approval_requests.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "approval request")
	}
	log.Printf("[DEBUG] Retrieved ECL approval request: %#v", approval)

	d.Set("request_id", approval.RequestID)
	d.Set("external_request_id", approval.ExternalRequestID)
	d.Set("approver_type", approval.ApproverType)
	d.Set("approver_id", approval.ApproverID)
	d.Set("request_user_id", approval.RequestUserID)
	d.Set("service", approval.Service)
	d.Set("actions", approval.Actions)
	d.Set("descriptions", approval.Descriptions)
	d.Set("request_user", approval.RequestUser)
	d.Set("approver", approval.Approver)
	d.Set("approval_deadline", approval.ApprovalDeadLine)
	d.Set("approval_expire", approval.ApprovalExpire)
	d.Set("registered_time", approval.RegisteredTime)
	d.Set("updated_time", approval.UpdatedTime)
	d.Set("status", approval.Status)

	return nil
}

func resourceSSSApprovalRequestV1Delete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
