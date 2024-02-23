package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3BucketPolicy = `{
  "block": {
    "attributes": {
      "bucket": {
        "computed": true,
        "description": "The name of the Amazon S3 bucket to which the policy applies.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3::BucketPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3BucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3BucketPolicy), &result)
	return &result
}
