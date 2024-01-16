package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ObjectlambdaAccessPointPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_lambda_access_point": {
        "description": "The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "description": "A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. ",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3ObjectlambdaAccessPointPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ObjectlambdaAccessPointPolicy), &result)
	return &result
}
