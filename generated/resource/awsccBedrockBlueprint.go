package resource

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
        "description": "Name of the Blueprint",
        "description_kind": "plain",
        "required": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_encryption_context": {
        "computed": true,
        "description": "KMS encryption context",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "kms_key_id": {
        "computed": true,
        "description": "KMS key identifier",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Last modified timestamp",
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "description": "Schema of the blueprint",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value for the tag",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "description": "Modality Type",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::Blueprint Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockBlueprintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockBlueprint), &result)
	return &result
}
