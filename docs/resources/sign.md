---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "scribble_sign Resource - terraform-provider-scribble"
subcategory: ""
description: |-
  This signs the provided image digest with scribble.
---

# scribble_sign (Resource)

This signs the provided image digest with scribble.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `image` (String) The digest of the container image to sign.

### Optional

- `conflict` (String) How to handle conflicting signature values
- `fulcio_url` (String) Address of sigstore PKI server (default https://fulcio.sigstore.dev).
- `rekor_url` (String) Address of rekor transparency log server (default https://rekor.sigstore.dev).

### Read-Only

- `id` (String) The immutable digest this resource signs.
- `signed_ref` (String) This always matches the input digest, but is a convenience for composition.


