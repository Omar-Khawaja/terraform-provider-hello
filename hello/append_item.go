package hello

import (
	"context"

	hello "github.com/Omar-Khawaja/hello-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func appendItem() *schema.Resource {
	return &schema.Resource{
		CreateContext: appendItemCreate,
		ReadContext:   appendItemRead,
		UpdateContext: appendItemUpdate,
		DeleteContext: appendItemDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "name of append action",
				Required:    true,
				Type:        schema.TypeString,
			},
			"item": {
				Description: "actual item to append",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func appendItemCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_ = d.Get("name").(string)
	item := d.Get("item").(string)
	config := &hello.Config{}
	client, err := hello.NewClient(config)
	if err != nil {
		return diag.FromErr(err)
	}
	err = client.AppendItem(item)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("name")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func appendItemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func appendItemUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func appendItemDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
