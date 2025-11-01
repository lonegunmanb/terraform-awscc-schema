package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3VectorsVectorBucketPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description": "A policy document containing permissions to add to the specified vector bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vector_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the vector bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vector_bucket_name": {
        "computed": true,
        "description": "The name of the vector bucket",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3Vectors::VectorBucketPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3VectorsVectorBucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3VectorsVectorBucketPolicy), &result)
	return &result
}
