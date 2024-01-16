package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3OutpostsBucketPolicy = `{
  "block": {
    "attributes": {
      "bucket": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified bucket.",
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
        "description": "A policy document containing permissions to add to the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Outposts::BucketPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3OutpostsBucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3OutpostsBucketPolicy), &result)
	return &result
}
