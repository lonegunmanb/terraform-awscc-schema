package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ObjectlambdaAccessPointPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "object_lambda_access_point": {
        "computed": true,
        "description": "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3ObjectLambda::AccessPointPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3ObjectlambdaAccessPointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ObjectlambdaAccessPointPolicy), &result)
	return &result
}
