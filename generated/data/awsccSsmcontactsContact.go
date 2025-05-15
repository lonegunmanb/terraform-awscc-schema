package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmcontactsContact = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the contact.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plan": {
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
            "rotation_ids": {
              "computed": true,
              "description": "List of Rotation Ids to associate with Contact",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
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
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "Contact type, which specify type of contact. Currently supported values: ?PERSONAL?, ?SHARED?, ?OTHER?.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSMContacts::Contact",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmcontactsContactSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmcontactsContact), &result)
	return &result
}
