package main

import (
	"fmt"
	"strings"

	"github.com/tnnmuhandiram/terraform/modules/terraform"
	test_structure "github.com/tnnmuhandiram/terraform/modules/test-structure"

	"github.com/tnnmuhandiram/terraform/modules/random"
)

func main() {
	TestTerraformGcpExample()
}

func TestTerraformGcpExample() {
	// t.Parallel()

	exampleDir := test_structure.CopyTerraformFolderToTemp("./", "examples/terraform-gcp-example")

	// Get the Project Id to use
	projectId := "postgress-cluster"

	// Create all resources in the following zone
	zone := "us-east1-b"

	// Give the example bucket a unique name so we can distinguish it from any other bucket in your GCP account
	expectedBucketName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))

	// Also give the example instance a unique name
	expectedInstanceName := fmt.Sprintf("terratest-gcp-example-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: exampleDir,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"gcp_project_id": projectId,
			"zone":           zone,
			"instance_name":  expectedInstanceName,
			"bucket_name":    expectedBucketName,
		},
	}
	// t := *testing.T
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(terraformOptions)

	// // This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	// out := terraform.InitAndApply(terraformOptions)
	// fmt.Print(terraformOptions)
	// fmt.Print(out)
	// // Run `terraform output` to get the value of some of the output variables
	// bucketURL := terraform.Output(t, terraformOptions, "bucket_url")
	// instanceName := terraform.Output(t, terraformOptions, "instance_id")

	// // Verify that the new bucket url matches the expected url
	// expectedURL := fmt.Sprintf("gs://%s", expectedBucketName)
	// assert.Equal(t, expectedURL, bucketURL)

	// // Verify that the Storage Bucket exists
	// gcp.AssertStorageBucketExists(t, expectedBucketName)

	// // Add a tag to the Compute Instance
	// instance := gcp.FetchInstance(t, projectId, instanceName)
	// instance.SetLabels(t, map[string]string{"testing": "testing-tag-value2"})

	// // Check for the labels within a retry loop as it can sometimes take a while for the
	// // changes to propagate.
	// maxRetries := 12
	// timeBetweenRetries := 5 * time.Second
	// expectedText := "testing-tag-value2"

	// retry.DoWithRetry(t, fmt.Sprintf("Checking Instance %s for labels", instanceName), maxRetries, timeBetweenRetries, func() (string, error) {
	// 	// Look up the tags for the given Instance ID
	// 	instance := gcp.FetchInstance(t, projectId, instanceName)
	// 	instanceLabels := instance.GetLabels(t)

	// 	testingTag, containsTestingTag := instanceLabels["testing"]
	// 	actualText := strings.TrimSpace(testingTag)
	// 	if !containsTestingTag {
	// 		return "", fmt.Errorf("Expected the tag 'testing' to exist")
	// 	}

	// 	if actualText != expectedText {
	// 		return "", fmt.Errorf("Expected GetLabelsForComputeInstanceE to return '%s' but got '%s'", expectedText, actualText)
	// 	}

	// 	return "", nil
	// })
}
