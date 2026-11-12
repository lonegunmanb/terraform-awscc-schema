package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedAgentOrganizationalProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Agent Organizational Profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "business_overview": {
        "computed": true,
        "description": "A business overview for the organizational profile used to improve recommendation quality.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the organizational profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The identifier of the system or user that created this organizational profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether deletion protection is enabled for the organizational profile. When enabled, the profile cannot be deleted until deletion protection is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description of the organizational profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The human-readable display name of the organizational profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_at": {
        "computed": true,
        "description": "The timestamp when the organizational profile was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_by": {
        "computed": true,
        "description": "The identifier of the system or user that last modified this organizational profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the organizational profile. Unique within the account and used as the last component of the ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "pillars": {
        "computed": true,
        "description": "The list of Well-Architected pillars to focus on.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs to associate with the Agent Organizational Profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::WellArchitected::AgentOrganizationalProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedAgentOrganizationalProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedAgentOrganizationalProfile), &result)
	return &result
}
