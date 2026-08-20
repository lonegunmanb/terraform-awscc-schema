package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "Timestamp (milliseconds after Jan 1, 1970 00:00:00 UTC) when the storage tier policy was last updated.",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_tier": {
        "description": "The storage tier to apply. Only INTELLIGENT_TIERING is accepted for creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Logs::StorageTierPolicy. Manages the storage tier policy for a CloudWatch Logs account. When created, enables Intelligent-Tiering which automatically moves infrequently accessed log data to lower-cost storage tiers. Deleting this resource reverts the account to standard (single-tier) storage.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsStorageTierPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsStorageTierPolicy), &result)
	return &result
}
