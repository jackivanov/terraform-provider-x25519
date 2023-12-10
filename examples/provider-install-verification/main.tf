terraform {
  required_providers {
    x25519 = {
      source = "hashicorp.com/edu/x25519"
    }
  }
}

provider "x25519" {}

# resoutce test
resource "x25519_private_key" "example" {}

data "x25519_public_key" "example" {
  private_key = x25519_private_key.example.private_key
}

output "x25519_private_key" {
  value = x25519_private_key.example
}

output "x25519_public_key" {
  value = data.x25519_public_key.example
}
