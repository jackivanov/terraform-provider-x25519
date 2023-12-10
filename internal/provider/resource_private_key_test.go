package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccOrderResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
				resource "x25519_private_key" "test" {}
				data "x25519_public_key" "test" {
					private_key = x25519_private_key.test.private_key
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.x25519_public_key.test", "public_key",
						"x25519_private_key.test", "public_key",
					),
					resource.TestCheckResourceAttrSet("x25519_private_key.test", "id"),
				),
			},
		},
	})
}
