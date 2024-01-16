package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectView = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "The actions of the view in an array.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "The description of the view.",
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
      "instance_arn": {
        "description": "The Amazon Resource Name (ARN) of the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the view.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "template": {
        "description": "The template of the view as JSON.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "view_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the view.",
        "description_kind": "plain",
        "type": "string"
      },
      "view_content_sha_256": {
        "computed": true,
        "description": "The view content hash.",
        "description_kind": "plain",
        "type": "string"
      },
      "view_id": {
        "computed": true,
        "description": "The view id of the view.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::View",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectViewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectView), &result)
	return &result
}
