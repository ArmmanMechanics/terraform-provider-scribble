terraform {
  required_providers {
    scribble  = {
        source = "ArmmanMechanics/scribble"
    }
  }
}

resource "scribble_sign" "image" {
  image = "us-docker.pkg.dev/wlynch-chainguard/public/ko-gcloud@sha256:d882f1b1ba89f712f00d955c7268d66f89774f79e922258cd6194ae18e8ac7ce"
}
