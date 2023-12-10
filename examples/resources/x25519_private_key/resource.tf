resource "x25519_private_key" "example" {}

output "x25519_keys" {
  value = {
    private_key = x25519_private_key.example.private_key
    public_key  = x25519_private_key.example.public_key
  }
}
