package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectView = `{
  "block": {
    "attributes": {
      "actions": {
        "computed": true,
        "description": "The actions of the view in an array.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "The description of the view.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the view.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "template": {
        "computed": true,
        "description": "The template of the view as JSON.",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::Connect::View",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectViewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectView), &result)
	return &result
}
