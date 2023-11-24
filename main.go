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

//go:generate git switch -c "release_v0.0.2+$((RANDOM % 1000))"
//go:generate git commit --allow-empty -m 'Release'
//go:generate git tag -f v0.0.2
//go:generate git push -f origin
//go:generate exit 0

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
