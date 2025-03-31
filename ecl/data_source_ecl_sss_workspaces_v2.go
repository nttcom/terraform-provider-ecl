package ecl

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/workspaces"
)

func dataSourceSSSWorkspacesV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSSSWorkspacesV2Read,

		Schema: map[string]*schema.Schema{
			"workspace_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"contract_id": &schema.Schema{
				Type:     schema.TypeString,
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

// dataSourceSSSWorkspacesV2Read performs the project lookup.
func dataSourceSSSWorkspacesV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sssV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	listOpts := workspaces.ListOpts{}

	log.Printf("[DEBUG] List Options: %#v", listOpts)

	var workspace workspaces.Workspace
	allPages, err := workspaces.List(client, listOpts).AllPages()
	if err != nil {
		return fmt.Errorf("Unable to query workspaces: %s", err)
	}

	allWorkspaces, err := workspaces.ExtractWorkspaces(allPages)
	if err != nil {
		return fmt.Errorf("Unable to retrieve workspaces: %s", err)
	}

	var refinedWorkspaces []workspaces.Workspace
	if len(allWorkspaces) > 0 {
		for _, t := range allWorkspaces {
			if t.WorkspaceName == d.Get("workspace_name").(string) {
				refinedWorkspaces = append(refinedWorkspaces, t)
			}
		}
	}

	if len(refinedWorkspaces) > 1 {
		log.Printf("[DEBUG] Multiple results found: %#v", allWorkspaces)
		return fmt.Errorf("Your query returned more than one result")
	}

	if len(refinedWorkspaces) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	workspace = refinedWorkspaces[0]

	log.Printf("[DEBUG] Single workspace found: %s", workspace.WorkspaceID)

	d.SetId(workspace.WorkspaceID)
	d.Set("workspace_name", workspace.WorkspaceName)
	d.Set("description", workspace.Description)
	d.Set("contract_id", workspace.ContractID)
	d.Set("workspace_id", workspace.WorkspaceID)
	d.Set("start_time", workspace.StartTime.String())

	return nil
}
