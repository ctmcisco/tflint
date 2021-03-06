// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCodepipelineInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_codepipeline" "foo" {
	name = "test/pipeline"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCodepipelineInvalidNameRule(),
					Message: `"test/pipeline" does not match valid pattern ^[A-Za-z0-9.@\-_]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_codepipeline" "foo" {
	name = "tf-test-pipeline"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCodepipelineInvalidNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
