package hello

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"hello_ping":        pingOrder(),
			"hello_append_item": appendItem(),
		},
	}
}
