package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoAssignment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The sso instance that the permission set is owned.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission_set_arn": {
        "description": "The permission set that the assignemt will be assigned",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_id": {
        "description": "The assignee's identifier, user id/group id",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_type": {
        "description": "The assignee's type, user/group",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_id": {
        "description": "The account id to be provisioned.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_type": {
        "description": "The type of resource to be provsioned to, only aws account now",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for SSO assignmet",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsoAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoAssignment), &result)
	return &result
}
