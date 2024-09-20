package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieFindingsFilter = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "Findings filter action.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Findings filter ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Findings filter description",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "finding_criteria": {
        "description": "Findings filter criteria.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criterion": {
              "computed": true,
              "description": "Map of filter criteria.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "gt": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lt": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "neq": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "map"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "findings_filter_id": {
        "computed": true,
        "description": "Findings filter ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Findings filter name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "position": {
        "computed": true,
        "description": "Findings filter position.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Macie FindingsFilter resource schema.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMacieFindingsFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieFindingsFilter), &result)
	return &result
}
