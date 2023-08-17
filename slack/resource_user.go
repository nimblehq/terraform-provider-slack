package slack

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/slack-go/slack"
)

func resourceSlackUser() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceSlackUserRead,
		CreateContext: resourceSlackUserCreate,
		UpdateContext: resourceSlackUserUpdate,
		DeleteContext: resourceSlackUserDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSlackUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*slack.Client)
	var diags diag.Diagnostics

	id := d.Id()

	user, err := client.GetUserByEmailContext(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  fmt.Sprintf("error getting user %s: %s", id, err.Error()),
		})
		d.SetId("")

		return diags
	}

	d.Set("email", user.Profile.Email)
	d.Set("name", user.Name)
	d.SetId(user.Profile.Email)

	return nil
}

func resourceSlackUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := m.(*slack.Client)
	// var diags diag.Diagnostics

	// email := d.Get("email").(string)
	// name := d.Get("name").(string)

	// if d.

	// user, err := client.GetUserByEmailContext(ctx, email)
	// if err != nil {
	// 	tflog.Info(ctx, fmt.Sprintf("Error getting user by email %s: %s", email, err.Error()))
	// }

	// if user != nil {
	// 	diags = append(diags, diag.Diagnostic{
	// 		Severity: diag.Warning,
	// 		Summary:  fmt.Sprintf("user with email %s already exists", email),
	// 	})
	// 	d.SetId(email)

	// 	return diags
	// }

	// _, err = client.InviteToTeam()
}

func resourceSlackUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceSlackUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
