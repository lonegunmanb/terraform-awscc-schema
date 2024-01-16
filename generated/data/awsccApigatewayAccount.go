package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayAccount = `{
  "block": {
    "attributes": {
      "cloudwatch_role_arn": {
        "computed": true,
        "description": "The ARN of an Amazon CloudWatch role for the current Account.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::Account",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayAccount), &result)
	return &result
}
