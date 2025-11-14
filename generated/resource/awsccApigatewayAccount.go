package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayAccount = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudwatch_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Account` + "`" + `` + "`" + ` resource specifies the IAM role that Amazon API Gateway uses to write API logs to Amazon CloudWatch Logs. To avoid overwriting other roles, you should only have one ` + "`" + `` + "`" + `AWS::ApiGateway::Account` + "`" + `` + "`" + ` resource per region per account. \n When you delete a stack containing this resource, API Gateway can still assume the provided IAM role to write API logs to CloudWatch Logs. To deny API Gateway access to write API logs to CloudWatch logs, update the permissions policies or change the IAM role to deny access.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayAccount), &result)
	return &result
}
