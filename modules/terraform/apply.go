package terraform

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// InitAndApply runs terraform init and apply with the given options and return stdout/stderr from the apply command. Note that this
// method does NOT call destroy and assumes the caller is responsible for cleaning up any resources created by running
// apply.
func InitAndApply(options *Options) string {
	out, err := InitAndApplyE(options)
	// require.NoError(err)
	return out
}

// InitAndApplyE runs terraform init and apply with the given options and return stdout/stderr from the apply command. Note that this
// method does NOT call destroy and assumes the caller is responsible for cleaning up any resources created by running
// apply.
func InitAndApplyE(options *Options) (string, error) {
	if _, err := InitE(options); err != nil {
		return "", err
	}

	// if _, err := GetE(options); err != nil {
	// 	return "", err
	// }

	return ApplyE(options)
}

// Apply runs terraform apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func Apply(t *testing.T, options *Options) string {
	out, err := ApplyE(options)
	require.NoError(t, err)
	return out
}

// TgApplyAll runs terragrunt apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func TgApplyAll(t *testing.T, options *Options) string {
	out, err := TgApplyAllE(t, options)
	require.NoError(t, err)
	return out
}

// ApplyE runs terraform apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func ApplyE(options *Options) (string, error) {
	return RunTerraformCommandE(options, FormatArgs(options, "apply", "-input=false", "-lock=false", "-auto-approve")...)
}

// TgApplyAllE runs terragrunt apply-all with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func TgApplyAllE(t *testing.T, options *Options) (string, error) {
	if options.TerraformBinary != "terragrunt" {
		return "", TgInvalidBinary(options.TerraformBinary)
	}

	return RunTerraformCommandE(options, FormatArgs(options, "apply-all", "-input=false", "-lock=false", "-auto-approve")...)
}
