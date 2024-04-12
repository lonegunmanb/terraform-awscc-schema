package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodedeployApplication = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "A name for the application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_platform": {
        "computed": true,
        "description": "The compute platform that CodeDeploy deploys the application to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to CodeDeploy applications to help you organize and categorize them. Each tag consists of a key and an optional value, both of which you define. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The AWS::CodeDeploy::Application resource creates an AWS CodeDeploy application",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodedeployApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodedeployApplication), &result)
	return &result
}
