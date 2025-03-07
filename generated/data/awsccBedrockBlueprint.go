package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockBlueprint = `{
  "block": {
    "attributes": {
      "blueprint_arn": {
        "computed": true,
        "description": "ARN of a Blueprint",
        "description_kind": "plain",
        "type": "string"
      },
      "blueprint_name": {
        "computed": true,
        "description": "Name of the Blueprint",
        "description_kind": "plain",
        "type": "string"
      },
      "blueprint_stage": {
        "computed": true,
        "description": "Stage of the Blueprint",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_encryption_context": {
        "computed": true,
        "description": "KMS encryption context",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "kms_key_id": {
        "computed": true,
        "description": "KMS key identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Last modified timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description": "Schema of the blueprint",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Key for the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "Modality Type",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::Blueprint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockBlueprintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockBlueprint), &result)
	return &result
}
