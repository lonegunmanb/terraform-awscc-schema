package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoAssignment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The sso instance that the permission set is owned.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission_set_arn": {
        "computed": true,
        "description": "The permission set that the assignemt will be assigned",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_id": {
        "computed": true,
        "description": "The assignee's identifier, user id/group id",
        "description_kind": "plain",
        "type": "string"
      },
      "principal_type": {
        "computed": true,
        "description": "The assignee's type, user/group",
        "description_kind": "plain",
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description": "The account id to be provisioned.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description": "The type of resource to be provsioned to, only aws account now",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSO::Assignment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsoAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoAssignment), &result)
	return &result
}
