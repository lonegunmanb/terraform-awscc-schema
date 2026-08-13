package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAccessanalyzerArchiveRule = `{
  "block": {
    "attributes": {
      "analyzer_name": {
        "computed": true,
        "description": "The name of the analyzer for the archive rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the archive rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time at which the archive rule was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "filter": {
        "computed": true,
        "description": "The criteria for the archive rule. A map of filter criteria property names to their criterion values.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "contains": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "eq": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "exists": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "neq": {
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
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_name": {
        "computed": true,
        "description": "The name of the archive rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The time at which the archive rule was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AccessAnalyzer::ArchiveRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAccessanalyzerArchiveRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAccessanalyzerArchiveRule), &result)
	return &result
}
