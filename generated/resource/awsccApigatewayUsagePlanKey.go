package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayUsagePlanKey = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_id": {
        "description": "The Id of the UsagePlanKey resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "usage_plan_id": {
        "description": "The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "usage_plan_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::UsagePlanKey` + "`" + `` + "`" + ` resource associates an API key with a usage plan. This association determines which users the usage plan is applied to.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayUsagePlanKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayUsagePlanKey), &result)
	return &result
}
