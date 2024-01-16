package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudtrailResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "description": "The ARN of the AWS CloudTrail resource to which the policy applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_policy": {
        "description": "A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudTrail::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudtrailResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudtrailResourcePolicy), &result)
	return &result
}
