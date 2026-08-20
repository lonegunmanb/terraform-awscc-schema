package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsStorageTierPolicy = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID that owns this storage tier policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "Timestamp (milliseconds after Jan 1, 1970 00:00:00 UTC) when the storage tier policy was last updated.",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_tier": {
        "computed": true,
        "description": "The storage tier to apply. Only INTELLIGENT_TIERING is accepted for creation.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::StorageTierPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsStorageTierPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsStorageTierPolicy), &result)
	return &result
}
