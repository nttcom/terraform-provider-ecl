package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/workspace_roles"
)

func resourceSSSWorkspaceRolesV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSSWorkspaceRolesV2Create,
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

			"workspace_id": &schema.Schema{
			    Type:     schema.TypeString,
			    Required: true,
			},

			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSSSWorkspaceRolesV2Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	createOpts := workspaces.CreateOpts{
		UserID:        d.Get("user_id").(string),
		WorkspaceID:   d.Get("workspace_id").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	workspace_role, err := workspace_roles.Create(client, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating ECL workspace role: %s", err)
	}

	log.Printf("[DEBUG] Workspace role has successfully created.")
	d.SetId(workspace_role.WorkspaceID)

	return nil
}

func resourceSSSWorkspaceRolesV2Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	err = workspace_roles.Delete(client, d.Id()).ExtractErr()
	if err != nil {
		return fmt.Errorf("Error deleting ECL workspace role: %s", err)
	}

	log.Printf("[DEBUG] Workspace Role has successfully deleted.")
	return nil
}
