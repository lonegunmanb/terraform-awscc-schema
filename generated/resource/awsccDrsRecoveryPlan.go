package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDrsRecoveryPlan = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Recovery Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the Recovery Plan was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the Recovery Plan.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the Recovery Plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_plan_id": {
        "computed": true,
        "description": "The ID of the Recovery Plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the Recovery Plan. ACTIVE means executable. INVALID means the plan has no SERVER type steps and cannot be executed.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the Recovery Plan.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the Recovery Plan was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A Recovery Plan orchestrates multi-server disaster recovery in AWS Elastic Disaster Recovery.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDrsRecoveryPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDrsRecoveryPlan), &result)
	return &result
}
