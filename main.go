package main

import (
	"flag"

	"github.com/BlueMedoraPublic/bpcli/cmd"
)

func main() {
	cmd.Execute()
}

func init() {
	// init the glog config
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Parse()
}
