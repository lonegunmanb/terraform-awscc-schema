package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendrarankingExecutionPlan = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_units": {
        "computed": true,
        "description": "Capacity units",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rescore_capacity_units": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "A description for the execution plan",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Unique ID of rescore execution plan",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of kendra ranking rescore execution plan",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for labeling the execution plan",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "A KendraRanking Rescore execution plan",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKendrarankingExecutionPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendrarankingExecutionPlan), &result)
	return &result
}
