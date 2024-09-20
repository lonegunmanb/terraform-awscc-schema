package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerImage = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_description": {
        "computed": true,
        "description": "A description of the image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_display_name": {
        "computed": true,
        "description": "The display name of the image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_name": {
        "description": "The name of the image.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_role_arn": {
        "description": "The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::Image",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerImage), &result)
	return &result
}
