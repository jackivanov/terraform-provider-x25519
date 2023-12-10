---
page_title: "Provider: x25519"
description: |-
  The x25519 provider offers utilities for working with Curve25519 keys.
---

# X25519 Provider

The x25519 provider provides utilities for working with *Curve25519*
keys. It offers resources that enable the creation of
private keys and the corresponding public keys as part of a Terraform deployment.

This provider, on its own, may not have extensive standalone utility,
but it serves as a valuable tool for generating credentials.
These can be subsequently utilized with other providers when configuring resources
that expose x25519 services or when provisioning x25519 keys for specific use cases.

Use the navigation on the left to explore the available resources.

## Example Usage

```terraform
terraform {
  required_providers {
    x25519 = {
      source = "jackivanov/x25519"
    }
  }
}

resource "x25519_private_key" "example" {}

output "x25519_keys" {
  value = {
    private_key = x25519_private_key.example.private_key
    public_key  = x25519_private_key.example.public_key
  }
}
```

### Secrets and Terraform state

Some resources, such as `x25519_private_key`, are considered "secrets" and are marked as *sensitive*
by this provider to prevent unintentional leakage in logs or other outputs. However, it's crucial to
note that the values constituting the "state" of these resources will be stored in the
[Terraform state](https://www.terraform.io/language/state) file, including the "secrets" in an *unencrypted* form.
