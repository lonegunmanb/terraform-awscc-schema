package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoApplicationAssignment = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description": "The ARN of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_id": {
        "description": "An identifier for an object in IAM Identity Center, such as a user or group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "description": "The entity type for which the assignment will be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for SSO application access grant to a user or group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsoApplicationAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoApplicationAssignment), &result)
	return &result
}
