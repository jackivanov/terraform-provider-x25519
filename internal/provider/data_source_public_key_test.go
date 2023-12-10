package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccX25519DataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
				data "x25519_public_key" "test" {
					private_key = "iM4KhF7Zu6vYcTdamVOQsiNftCdlu0ceBZonXb02KmU="
				}
			`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.x25519_public_key.test",
						"public_key",
						"rXSwT/bVUlMB0URSwXrx1uPbGuo9GjYfKYDyYbV22TA=",
					),
					resource.TestCheckResourceAttrSet(
						"data.x25519_public_key.test",
						"id",
					),
				),
			},
		},
	})
}
