package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ExpressBucketPolicy = `{
  "block": {
    "attributes": {
      "bucket": {
        "description": "The name of the S3 directory bucket to which the policy applies.",
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
        "description": "A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3Express::BucketPolicy.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3ExpressBucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ExpressBucketPolicy), &result)
	return &result
}
