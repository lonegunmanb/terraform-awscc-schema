package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTableBucketPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_policy": {
        "description": "A policy document containing permissions to add to the specified table bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_bucket_arn": {
        "description": "The Amazon Resource Name (ARN) of the table bucket to which the policy applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Applies an IAM resource policy to a table bucket.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3TablesTableBucketPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTableBucketPolicy), &result)
	return &result
}
