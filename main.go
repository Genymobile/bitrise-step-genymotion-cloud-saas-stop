package main

import (
	"os"
	"strings"
	"sync"

	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/log"
)

// Config ...
type Config struct {
	GMCloudSaaSInstanceUUID string `env:"instance_uuid,required"`
}

// Define variable
var isError bool = false

// printError prints an error.
func printError(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// abortf prints an error and terminates step
func abortf(format string, args ...interface{}) {
	printError(format, args...)
	os.Exit(1)
}

// setOperationFailed marked step as failed
func setOperationFailed(format string, args ...interface{}) {
	printError(format, args...)
	isError = true
}

func stopInstance(wg *sync.WaitGroup, uuid string) {
	defer wg.Done()

	cmd := command.New("gmsaas", "instances", "stop", uuid)
	out, err := cmd.RunAndReturnTrimmedCombinedOutput()
	if err != nil {
		setOperationFailed("Failed to stop instance %s, error: error: %s | output: %s", uuid, cmd.PrintableCommandArgs(), err, out)
		return

	}
	log.Donef("Instance stopped %s", uuid)
}

func main() {
	var c Config
	if err := stepconf.Parse(&c); err != nil {
		abortf("Issue with input: %s", err)
	}
	stepconf.Print(c)

	log.Infof("Stopping Genymotion Cloud SaaS instances")
	instancesList := strings.Split(c.GMCloudSaaSInstanceUUID, ",")

	var wg sync.WaitGroup
	for cptInstance := 0; cptInstance < len(instancesList); cptInstance++ {
		wg.Add(1)
		go stopInstance(&wg, instancesList[cptInstance])

	}
	wg.Wait()

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	if isError {
		// If at least one error happens, step will fail
		os.Exit(1)
	}
	os.Exit(0)
}
