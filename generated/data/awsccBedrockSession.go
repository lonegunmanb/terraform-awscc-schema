package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockSession = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp for when the session was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the KMS key to use to encrypt the session data.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp for when the session was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the session.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_id": {
        "computed": true,
        "description": "The unique identifier of the session in UUID format.",
        "description_kind": "plain",
        "type": "string"
      },
      "session_metadata": {
        "computed": true,
        "description": "A map of key-value pairs containing attributes to be persisted across the session.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "session_status": {
        "computed": true,
        "description": "The current status of the session.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags associated with the session.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Bedrock::Session",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockSessionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockSession), &result)
	return &result
}
