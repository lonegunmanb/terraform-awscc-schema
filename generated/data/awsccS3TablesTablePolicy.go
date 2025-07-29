package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTablePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "The namespace that the table belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_policy": {
        "computed": true,
        "description": "A policy document containing permissions to add to the specified table. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description": "The name for the table.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Tables::TablePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3TablesTablePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTablePolicy), &result)
	return &result
}
