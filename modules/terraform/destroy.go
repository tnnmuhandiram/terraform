package terraform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Destroy runs terraform destroy with the given options and return stdout/stderr.
func Destroy(options *Options) string {
	out, err := DestroyE(options)
	// require.NoError(t, err)
	print(err)
	return out
}

// TgDestroyAll runs terragrunt destroy with the given options and return stdout.
func TgDestroyAll(t *testing.T, options *Options) string {
	out, err := TgDestroyAllE(t, options)
	require.NoError(t, err)
	return out
}

// DestroyE runs terraform destroy with the given options and return stdout/stderr.
func DestroyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "destroy", "-auto-approve", "-input=false", "-lock=false")...)
}

// TgDestroyAllE runs terragrunt destroy with the given options and return stdout.
func TgDestroyAllE(t *testing.T, options *Options) (string, error) {
	if options.TerraformBinary != "terragrunt" {
		return "", TgInvalidBinary(options.TerraformBinary)
	}

	return RunTerraformCommandE(options, FormatArgs(options, "destroy-all", "-force", "-input=false", "-lock=false")...)
}
