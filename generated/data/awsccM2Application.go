package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccM2Application = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_location": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_type": {
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
      "kms_key_id": {
        "computed": true,
        "description": "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting application-related resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::M2::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccM2ApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccM2Application), &result)
	return &result
}
