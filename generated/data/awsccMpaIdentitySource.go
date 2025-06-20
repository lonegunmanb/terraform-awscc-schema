package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMpaIdentitySource = `{
  "block": {
    "attributes": {
      "creation_time": {
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
      "identity_source_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_source_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iam_identity_center": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "approval_portal_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "identity_source_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      }
    },
    "description": "Data Source schema for AWS::MPA::IdentitySource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMpaIdentitySourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMpaIdentitySource), &result)
	return &result
}
