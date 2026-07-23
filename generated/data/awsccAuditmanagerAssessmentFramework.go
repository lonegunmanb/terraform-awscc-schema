package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAuditmanagerAssessmentFramework = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "compliance_type": {
        "computed": true,
        "description": "The compliance type that the framework supports, such as CIS or HIPAA.",
        "description_kind": "plain",
        "type": "string"
      },
      "control_sets": {
        "computed": true,
        "description": "The control sets that are associated with the framework.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "controls": {
              "computed": true,
              "description": "The list of controls within the control set.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "The unique identifier of the control.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The name of the control set.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The time when the framework was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The user or role that created the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "framework_id": {
        "computed": true,
        "description": "The unique identifier for the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The time when the framework was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_by": {
        "computed": true,
        "description": "The user or role that most recently updated the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the framework.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
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
        "description": "The framework type, such as a standard framework or a custom framework.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AuditManager::AssessmentFramework",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAuditmanagerAssessmentFrameworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAuditmanagerAssessmentFramework), &result)
	return &result
}
