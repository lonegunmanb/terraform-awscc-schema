package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLaunchwizardDeployment = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN of the LaunchWizard deployment",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp of LaunchWizard deployment creation",
        "description_kind": "plain",
        "type": "string"
      },
      "deleted_at": {
        "computed": true,
        "description": "Timestamp of LaunchWizard deployment deletion",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_id": {
        "computed": true,
        "description": "Deployment ID of the LaunchWizard deployment",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_pattern_name": {
        "computed": true,
        "description": "Workload deployment pattern name",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of LaunchWizard deployment",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group": {
        "computed": true,
        "description": "Resource Group Name created for LaunchWizard deployment",
        "description_kind": "plain",
        "type": "string"
      },
      "specifications": {
        "computed": true,
        "description": "LaunchWizard deployment specifications",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "Status of LaunchWizard deployment",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for LaunchWizard deployment",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "workload_name": {
        "computed": true,
        "description": "Workload Name for LaunchWizard deployment",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::LaunchWizard::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLaunchwizardDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLaunchwizardDeployment), &result)
	return &result
}
