package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoApplicationAssignment = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description": "The ARN of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_id": {
        "computed": true,
        "description": "An identifier for an object in IAM Identity Center, such as a user or group",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description": "The entity type for which the assignment will be created.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSO::ApplicationAssignment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsoApplicationAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoApplicationAssignment), &result)
	return &result
}
