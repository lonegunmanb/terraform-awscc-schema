package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAiopsInvestigationGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Investigation Group's ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "chatbot_notification_channels": {
        "computed": true,
        "description": "An array of key-value pairs of notification channels to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "chat_configuration_arns": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "sns_topic_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp value.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "User friendly name for resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "cross_account_configurations": {
        "computed": true,
        "description": "An array of cross account configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_role_arn": {
              "computed": true,
              "description": "The Investigation Role's ARN.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "encryption_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_configuration_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "investigation_group_policy": {
        "computed": true,
        "description": "Investigation Group policy",
        "description_kind": "plain",
        "type": "string"
      },
      "is_cloud_trail_event_history_enabled": {
        "computed": true,
        "description": "Flag to enable cloud trail history",
        "description_kind": "plain",
        "type": "bool"
      },
      "last_modified_at": {
        "computed": true,
        "description": "User friendly name for resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_by": {
        "computed": true,
        "description": "User friendly name for resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "User friendly name for resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "retention_in_days": {
        "computed": true,
        "description": "The number of days to retain the investigation group",
        "description_kind": "plain",
        "type": "number"
      },
      "role_arn": {
        "computed": true,
        "description": "The Investigation Role's ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "tag_key_boundaries": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::AIOps::InvestigationGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAiopsInvestigationGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAiopsInvestigationGroup), &result)
	return &result
}
