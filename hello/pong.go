package hello

import (
	"context"

	hello "github.com/Omar-Khawaja/hello-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pongOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: pongOrderCreate,
		ReadContext:   pongOrderRead,
		UpdateContext: pongOrderUpdate,
		DeleteContext: pongOrderDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name of pong command",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func pongOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = d.Get("name").(string)
	config := &hello.Config{}
	client, err := hello.NewClient(config)
	if err != nil {
		return diag.FromErr(err)
	}
	err = client.PongOrder()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("name")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func pongOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func pongOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func pongOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
