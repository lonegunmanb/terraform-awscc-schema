package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStoragegatewayTapePool = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom tape pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description": "The unique identifier of the custom tape pool, extracted from the ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_name": {
        "description": "The name of the custom tape pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "retention_lock_time_in_days": {
        "computed": true,
        "description": "Tape retention lock time in days (up to 36,500 days / 100 years).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "retention_lock_type": {
        "computed": true,
        "description": "Tape retention lock type. Governance mode allows authorized removal; compliance mode prevents all removal.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_class": {
        "description": "The storage class associated with the custom pool (S3 Glacier or S3 Glacier Deep Archive).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of up to 50 tags for the tape pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
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
    "description": "Creates a custom tape pool for archiving virtual tapes with optional retention lock.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccStoragegatewayTapePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStoragegatewayTapePool), &result)
	return &result
}
