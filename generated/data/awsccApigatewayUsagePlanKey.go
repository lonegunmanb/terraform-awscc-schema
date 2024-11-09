package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayUsagePlanKey = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_id": {
        "computed": true,
        "description": "The Id of the UsagePlanKey resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "usage_plan_id": {
        "computed": true,
        "description": "The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer.",
        "description_kind": "plain",
        "type": "string"
      },
      "usage_plan_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::UsagePlanKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayUsagePlanKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayUsagePlanKey), &result)
	return &result
}
