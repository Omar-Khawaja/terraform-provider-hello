package hello

import (
	"context"

	hello "github.com/Omar-Khawaja/hello-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pingOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: pingOrderCreate,
		ReadContext:   pingOrderRead,
		UpdateContext: pingOrderUpdate,
		DeleteContext: pingOrderDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name of ping command",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func pingOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = d.Get("name").(string)
	config := &hello.Config{}
	client, err := hello.NewClient(config)
	if err != nil {
		return diag.FromErr(err)
	}
	err = client.PingOrder()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("name")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func pingOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func pingOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func pingOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
