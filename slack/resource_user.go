package slack

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/slack-go/slack"
)

func resourceSlackUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceSlackUserRead,
		// CreateContext: resourceSlackUserGroupCreate,
		// UpdateContext: resourceSlackUserGroupUpdate,
		// DeleteContext: resourceSlackUserGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSlackUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*slack.Client)
	id := d.Id()
	var diags diag.Diagnostics

	user, err := client.GetUserByEmailContext(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("user %s not found", id),
		})
		d.SetId("")

		return diags
	}

	d.Set("email", user.Profile.Email)
	d.Set("name", user.Name)
	d.SetId(user.Profile.Email)

	return nil
}
