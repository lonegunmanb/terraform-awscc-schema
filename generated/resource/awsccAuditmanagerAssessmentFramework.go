package resource

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
        "optional": true,
        "type": "string"
      },
      "control_sets": {
        "description": "The control sets that are associated with the framework.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "controls": {
              "description": "The list of controls within the control set.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "description": "The unique identifier of the control.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "name": {
              "description": "The name of the control set.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
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
        "optional": true,
        "type": "string"
      },
      "framework_id": {
        "computed": true,
        "description": "The unique identifier for the framework.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
        "description": "The name of the framework.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "computed": true,
        "description": "The framework type, such as a standard framework or a custom framework.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a custom framework in AWS Audit Manager.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAuditmanagerAssessmentFrameworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAuditmanagerAssessmentFramework), &result)
	return &result
}
