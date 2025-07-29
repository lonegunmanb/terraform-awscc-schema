package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigExtension = `{
  "block": {
    "attributes": {
      "actions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of the extension Action.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the extension action.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the role for invoking the extension action.",
              "description_kind": "plain",
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "The URI of the extension action.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "extension_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "latest_version_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Name of the extension.",
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of the extension Parameter.",
              "description_kind": "plain",
              "type": "string"
            },
            "dynamic": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "required": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "map"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value tags to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "version_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::AppConfig::Extension",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppconfigExtensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigExtension), &result)
	return &result
}
