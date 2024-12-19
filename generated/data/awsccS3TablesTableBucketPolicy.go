package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTableBucketPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_policy": {
        "computed": true,
        "description": "A policy document containing permissions to add to the specified table bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the table bucket to which the policy applies.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Tables::TableBucketPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3TablesTableBucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTableBucketPolicy), &result)
	return &result
}
