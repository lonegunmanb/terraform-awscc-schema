package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region, used as part of the primary identifier for the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::ARCZonalShift::AutoshiftObserverNotificationStatus Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccArczonalshiftAutoshiftObserverNotificationStatusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccArczonalshiftAutoshiftObserverNotificationStatus), &result)
	return &result
}
