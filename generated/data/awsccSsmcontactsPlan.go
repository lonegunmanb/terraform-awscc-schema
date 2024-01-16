package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmcontactsPlan = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the contact.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_id": {
        "computed": true,
        "description": "Contact ID for the AWS SSM Incident Manager Contact to associate the plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rotation_ids": {
        "computed": true,
        "description": "Rotation Ids to associate with Oncall Contact for engagement.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "stages": {
        "computed": true,
        "description": "The stages that an escalation plan or engagement plan engages contacts and contact methods in.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "duration_in_minutes": {
              "computed": true,
              "description": "The time to wait until beginning the next stage.",
              "description_kind": "plain",
              "type": "number"
            },
            "targets": {
              "computed": true,
              "description": "The contacts or contact methods that the escalation plan or engagement plan is engaging.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "channel_target_info": {
                    "computed": true,
                    "description": "Information about the contact channel that SSM Incident Manager uses to engage the contact.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "channel_id": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the contact channel.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "retry_interval_in_minutes": {
                          "computed": true,
                          "description": "The number of minutes to wait to retry sending engagement in the case the engagement initially fails.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "contact_target_info": {
                    "computed": true,
                    "description": "The contact that SSM Incident Manager is engaging during an incident.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "contact_id": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the contact.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "is_essential": {
                          "computed": true,
                          "description": "A Boolean value determining if the contact's acknowledgement stops the progress of stages in the plan.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SSMContacts::Plan",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmcontactsPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmcontactsPlan), &result)
	return &result
}
