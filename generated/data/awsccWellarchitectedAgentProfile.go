package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedAgentProfile = `{
  "block": {
    "attributes": {
      "aggregation_configuration": {
        "computed": true,
        "description": "The aggregation configuration entries (account, regions, access role) associated with this profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role used to access resources in this account.",
              "description_kind": "plain",
              "type": "string"
            },
            "account_id": {
              "computed": true,
              "description": "The target AWS account ID.",
              "description_kind": "plain",
              "type": "string"
            },
            "regions": {
              "computed": true,
              "description": "The target regions in the account.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "set"
        }
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Agent Profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "business_overview": {
        "computed": true,
        "description": "A business overview for the profile used to improve recommendation quality.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The identifier of the system or user that created this profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether deletion protection is enabled for the profile. When enabled, the profile cannot be deleted until deletion protection is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description of the profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The human-readable display name of the profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role assumed to execute recommendation actions.",
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
        "description": "The timestamp when the profile was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_by": {
        "computed": true,
        "description": "The identifier of the system or user that last modified this profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the profile. Unique within the account and used as the last component of the ARN.",
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
        "description": "Key-value pairs to associate with the Agent Profile.",
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
    "description": "Data Source schema for AWS::WellArchitected::AgentProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWellarchitectedAgentProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedAgentProfile), &result)
	return &result
}
