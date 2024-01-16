package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatabrewRuleset = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the Ruleset",
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
        "description": "Name of the Ruleset",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
        "description": "List of the data quality rules in the ruleset",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "check_expression": {
              "description": "Expression with rule conditions",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "column_selectors": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of a column from a dataset",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "regex": {
                    "computed": true,
                    "description": "A regular expression for selecting a column from a dataset",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "disabled": {
              "computed": true,
              "description": "Boolean value to disable/enable a rule",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description": "Name of the rule",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "substitution_map": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value": {
                    "description": "Value or column name",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value_reference": {
                    "description": "Variable name",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "threshold": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "Threshold type for a rule",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "Threshold unit for a rule",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Threshold value for a rule",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "target_arn": {
        "description": "Arn of the target resource (dataset) to apply the ruleset to",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::DataBrew::Ruleset.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatabrewRulesetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewRuleset), &result)
	return &result
}
