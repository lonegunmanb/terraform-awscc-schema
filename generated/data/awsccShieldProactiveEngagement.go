package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccShieldProactiveEngagement = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "emergency_contact_list": {
        "computed": true,
        "description": "A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.\nTo enable proactive engagement, the contact list must include at least one phone number.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "contact_notes": {
              "computed": true,
              "description": "Additional notes regarding the contact.",
              "description_kind": "plain",
              "type": "string"
            },
            "email_address": {
              "computed": true,
              "description": "The email address for the contact.",
              "description_kind": "plain",
              "type": "string"
            },
            "phone_number": {
              "computed": true,
              "description": "The phone number for the contact",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "proactive_engagement_status": {
        "computed": true,
        "description": "If ` + "`" + `ENABLED` + "`" + `, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.\nIf ` + "`" + `DISABLED` + "`" + `, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Shield::ProactiveEngagement",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccShieldProactiveEngagementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccShieldProactiveEngagement), &result)
	return &result
}
