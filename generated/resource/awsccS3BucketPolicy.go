package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3BucketPolicy = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the Amazon S3 bucket to which the policy applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Applies an Amazon S3 bucket policy to an Amazon S3 bucket. If you are using an identity other than the root user of the AWS-account that owns the bucket, the calling identity must have the ` + "`" + `` + "`" + `PutBucketPolicy` + "`" + `` + "`" + ` permissions on the specified bucket and belong to the bucket owner's account in order to use this operation.\n If you don't have ` + "`" + `` + "`" + `PutBucketPolicy` + "`" + `` + "`" + ` permissions, Amazon S3 returns a ` + "`" + `` + "`" + `403 Access Denied` + "`" + `` + "`" + ` error. If you have the correct permissions, but you're not using an identity that belongs to the bucket owner's account, Amazon S3 returns a ` + "`" + `` + "`" + `405 Method Not Allowed` + "`" + `` + "`" + ` error.\n   As a security precaution, the root user of the AWS-account that owns a bucket can always use this operation, even if the policy explicitly denies the root user the ability to perform this action. \n  For more information, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html).\n The following operations are related to ` + "`" + `` + "`" + `PutBucketPolicy` + "`" + `` + "`" + `:\n  +   [Create",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3BucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3BucketPolicy), &result)
	return &result
}
