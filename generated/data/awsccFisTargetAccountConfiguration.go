package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFisTargetAccountConfiguration = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The AWS account ID of the target account.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the target account.",
        "description_kind": "plain",
        "type": "string"
      },
      "experiment_template_id": {
        "computed": true,
        "description": "The ID of the experiment template.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role for the target account.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::FIS::TargetAccountConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFisTargetAccountConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFisTargetAccountConfiguration), &result)
	return &result
}
