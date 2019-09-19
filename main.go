package main

import (
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
)

// Config ...
type Config struct {
	GMCloudSaaSInstanceUUID string `env:"gmcloud_saas_instance_uuid,required"`
}

// failf prints an error and terminates the step.
func failf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	os.Exit(1)
}

func main() {
	var c Config
	if err := stepconf.Parse(&c); err != nil {
		failf("Issue with input: %s", err)
	}
	stepconf.Print(c)

	log.Infof("Stop Android devices on Genymotion Cloud SaaS")
	cmd := exec.Command("gmsaas", "instances", "stop", c.GMCloudSaaSInstanceUUID)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		failf("Failed to stop a device, error: %#v | output: %s", err, stdout)
	} else {
		log.Donef("Device stopped %s", c.GMCloudSaaSInstanceUUID)
	}

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
