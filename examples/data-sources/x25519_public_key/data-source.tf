resource "x25519_private_key" "example" {}

data "x25519_public_key" "example" {
  private_key = x25519_private_key.example.private_key
}

output "x25519_public_key" {
  value = data.x25519_public_key.example
}
