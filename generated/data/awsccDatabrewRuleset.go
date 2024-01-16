package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the Ruleset",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description": "List of the data quality rules in the ruleset",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "check_expression": {
              "computed": true,
              "description": "Expression with rule conditions",
              "description_kind": "plain",
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
                    "type": "string"
                  },
                  "regex": {
                    "computed": true,
                    "description": "A regular expression for selecting a column from a dataset",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "disabled": {
              "computed": true,
              "description": "Boolean value to disable/enable a rule",
              "description_kind": "plain",
              "type": "bool"
            },
            "name": {
              "computed": true,
              "description": "Name of the rule",
              "description_kind": "plain",
              "type": "string"
            },
            "substitution_map": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value": {
                    "computed": true,
                    "description": "Value or column name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value_reference": {
                    "computed": true,
                    "description": "Variable name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
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
                    "type": "string"
                  },
                  "unit": {
                    "computed": true,
                    "description": "Threshold unit for a rule",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "Threshold value for a rule",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
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
      },
      "target_arn": {
        "computed": true,
        "description": "Arn of the target resource (dataset) to apply the ruleset to",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataBrew::Ruleset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatabrewRulesetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewRuleset), &result)
	return &result
}
