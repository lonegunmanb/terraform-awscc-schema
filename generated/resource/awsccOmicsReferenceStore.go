package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOmicsReferenceStore = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The store's ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "When the store was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the store.",
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
      "name": {
        "description": "A name for the store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reference_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sse_config": {
        "computed": true,
        "description": "Server-side encryption (SSE) settings for a store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_arn": {
              "computed": true,
              "description": "An encryption key ARN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Definition of AWS::Omics::ReferenceStore Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOmicsReferenceStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsReferenceStore), &result)
	return &result
}
