package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccArczonalshiftAutoshiftObserverNotificationStatus = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "User account id, used as part of the primary identifier for the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region, used as part of the primary identifier for the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ARCZonalShift::AutoshiftObserverNotificationStatus",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccArczonalshiftAutoshiftObserverNotificationStatusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccArczonalshiftAutoshiftObserverNotificationStatus), &result)
	return &result
}
