package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightAssignment = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM policy assignment. This is a synthesized identifier composed of the namespace and assignment name.",
        "description_kind": "plain",
        "type": "string"
      },
      "assignment_id": {
        "computed": true,
        "description": "The ID for the IAM policy assignment.",
        "description_kind": "plain",
        "type": "string"
      },
      "assignment_name": {
        "computed": true,
        "description": "The name of the assignment, also called a rule. The name must be unique within the AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "assignment_status": {
        "computed": true,
        "description": "The status of the assignment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identities": {
        "computed": true,
        "description": "The QuickSight users, groups, or both that you want to assign the policy to. This maps identity type to a list of identity names.",
        "description_kind": "plain",
        "type": [
          "map",
          [
            "list",
            "string"
          ]
        ]
      },
      "namespace": {
        "computed": true,
        "description": "The namespace that contains the assignment.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_arn": {
        "computed": true,
        "description": "The ARN for the IAM policy to apply to the QuickSight users and groups specified in this assignment.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::Assignment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightAssignmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightAssignment), &result)
	return &result
}
