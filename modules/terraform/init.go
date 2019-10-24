package terraform

import (
	"fmt"
)

// Init calls terraform init and return stdout/stderr.
func Init(options *Options) string {
	out, err := InitE(options)
	if err != nil {
		// t.Fatal(err)
	}
	return out
}

// InitE calls terraform init and return stdout/stderr.
func InitE(options *Options) (string, error) {
	args := []string{"init", fmt.Sprintf("-upgrade=%t", options.Upgrade)}
	args = append(args, FormatTerraformBackendConfigAsArgs(options.BackendConfig)...)
	return RunTerraformCommandE(options, args...)
}
