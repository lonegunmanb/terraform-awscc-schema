package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamServiceLinkedRole = `{
  "block": {
    "attributes": {
      "aws_service_name": {
        "computed": true,
        "description": "The service principal for the AWS service to which this role is attached.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_suffix": {
        "computed": true,
        "description": "A string that you provide, which is combined with the service-provided prefix to form the complete role name.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the role.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_name": {
        "computed": true,
        "description": "The name of the role.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IAM::ServiceLinkedRole",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamServiceLinkedRoleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamServiceLinkedRole), &result)
	return &result
}
