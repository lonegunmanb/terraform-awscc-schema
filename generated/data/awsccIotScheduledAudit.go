package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotScheduledAudit = `{
  "block": {
    "attributes": {
      "day_of_month": {
        "computed": true,
        "description": "The day of the month on which the scheduled audit takes place. Can be 1 through 31 or LAST. This field is required if the frequency parameter is set to MONTHLY.",
        "description_kind": "plain",
        "type": "string"
      },
      "day_of_week": {
        "computed": true,
        "description": "The day of the week on which the scheduled audit takes place. Can be one of SUN, MON, TUE,WED, THU, FRI, or SAT. This field is required if the frequency parameter is set to WEEKLY or BIWEEKLY.",
        "description_kind": "plain",
        "type": "string"
      },
      "frequency": {
        "computed": true,
        "description": "How often the scheduled audit takes place. Can be one of DAILY, WEEKLY, BIWEEKLY, or MONTHLY.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scheduled_audit_arn": {
        "computed": true,
        "description": "The ARN (Amazon resource name) of the scheduled audit.",
        "description_kind": "plain",
        "type": "string"
      },
      "scheduled_audit_name": {
        "computed": true,
        "description": "The name you want to give to the scheduled audit.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_check_names": {
        "computed": true,
        "description": "Which checks are performed during the scheduled audit. Checks must be enabled for your account.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::IoT::ScheduledAudit",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotScheduledAuditSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotScheduledAudit), &result)
	return &result
}
