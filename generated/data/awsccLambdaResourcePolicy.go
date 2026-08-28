package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The policy document you want to add to your LAM resource. This is formatted as a JSON string.\n For more information, see [Working with resource-based policies in](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the LAM resource you want to add the policy to. For a function, you can use a qualified or an unqualified ARN. The value must be a complete ARN, and the operation does not accept wildcard characters.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lambda::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaResourcePolicy), &result)
	return &result
}
