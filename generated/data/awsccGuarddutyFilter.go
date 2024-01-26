package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyFilter = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "detector_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "finding_criteria": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criterion": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eq": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "equals": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "greater_than": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "greater_than_or_equal": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "gt": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "gte": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "less_than": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "less_than_or_equal": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "lt": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "lte": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "neq": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "not_equals": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "map"
              }
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rank": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::GuardDuty::Filter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyFilter), &result)
	return &result
}
