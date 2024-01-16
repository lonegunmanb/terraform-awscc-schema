package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Account` + "`" + `` + "`" + ` resource specifies the IAM role that Amazon API Gateway uses to write API logs to Amazon CloudWatch Logs. To avoid overwriting other roles, you should only have one ` + "`" + `` + "`" + `AWS::ApiGateway::Account` + "`" + `` + "`" + ` resource per region per account.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayAccount), &result)
	return &result
}
