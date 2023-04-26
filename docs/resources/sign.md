---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cosign_sign Resource - terraform-provider-cosign"
subcategory: ""
description: |-
  This signs the provided image digest with cosign.
---

# cosign_sign (Resource)

This signs the provided image digest with cosign.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `image` (String) The digest of the container image to sign.

### Read-Only

- `id` (String) The immutable digest this resource signs.
- `signed_ref` (String) This always matches the input digest, but is a convenience for composition.

