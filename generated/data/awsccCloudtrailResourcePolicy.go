package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudtrailResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The ARN of the AWS CloudTrail resource to which the policy applies.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_policy": {
        "computed": true,
        "description": "A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudTrail::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudtrailResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudtrailResourcePolicy), &result)
	return &result
}
