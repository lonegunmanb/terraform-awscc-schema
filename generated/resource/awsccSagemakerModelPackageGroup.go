package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerModelPackageGroup = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time at which the model package group was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_package_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the model package group.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_package_group_description": {
        "computed": true,
        "description": "The description of the model package group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_package_group_name": {
        "description": "The name of the model package group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_package_group_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_package_group_status": {
        "computed": true,
        "description": "The status of a modelpackage group job.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
    "description": "Resource Type definition for AWS::SageMaker::ModelPackageGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerModelPackageGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerModelPackageGroup), &result)
	return &result
}
