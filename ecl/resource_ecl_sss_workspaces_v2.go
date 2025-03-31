package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/workspaces"
)

func resourceSSSWorkspacesV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSSWorkspacesV2Create,
		Read:   resourceSSSWorkspacesV2Read,
		Update: resourceSSSWorkspacesV2Update,
		Delete: resourceSSSWorkspacesV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"workspace_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"contract_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"workspace_id": &schema.Schema{
			    Type:     schema.TypeString,
			    Required: true,
			},
		},
	}
}

func resourceSSSWorkspacesV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	createOpts := workspaces.CreateOpts{
		WorkspaceName: d.Get("workspace_name").(string),
		Description:   d.Get("description").(string),
		ContractID:    d.Get("contract_id").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	workspace, err := workspaces.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL workspace: %s", err)
	}

	log.Printf("[DEBUG] Workspace has successfully created.")
	d.SetId(workspace.WorkspaceID)

	return resourceSSSWorkspacesV2Read(d, meta)
}

func resourceSSSWorkspacesV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	workspace, err := workspaces.Get(client, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "workspace")
	}
	log.Printf("[DEBUG] Retrieved ECL workspace: %#v", workspace)

	d.Set("workspace_name", workspace.WorkspaceName)
	d.Set("description", workspace.Description)
	d.Set("contract_id", workspace.ContractID)
	d.Set("start_time", workspace.StartTime.String())
    d.Set("workspace_id", workspace.WorkspaceID)

	log.Printf("[DEBUG] resourceSSSWorkspacesV2Read Succeeded")

	return nil
}

func resourceSSSWorkspacesV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	var hasChange bool
	var updateOpts workspaces.UpdateOpts

	if d.HasChange("description") {
		hasChange = true
		description := d.Get("description").(string)
		updateOpts.Description = &description
	}

	if hasChange {
		r := workspaces.Update(client, d.Id(), updateOpts)
		if r.Err != nil {
			return fmt.Errorf("Error updating ECL workspace: %s", r.Err)
		}
		log.Printf("[DEBUG] Workspace has successfully updated.")
	}

	return resourceSSSWorkspacesV2Read(d, meta)
}

func resourceSSSWorkspacesV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	err = workspaces.Delete(client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL workspace: %s", err)
	}

	log.Printf("[DEBUG] Workspace has successfully deleted.")
	return nil
}
