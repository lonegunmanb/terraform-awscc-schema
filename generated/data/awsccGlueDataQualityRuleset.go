package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueDataQualityRuleset = `{
  "block": {
    "attributes": {
      "client_token": {
        "computed": true,
        "description": "A unique token for idempotency.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the data quality ruleset.",
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
        "description": "A unique name for the data quality ruleset.",
        "description_kind": "plain",
        "type": "string"
      },
      "ruleset": {
        "computed": true,
        "description": "A Data Quality Definition Language (DQDL) ruleset.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "target_table": {
        "computed": true,
        "description": "An object representing an AWS Glue table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "database_name": {
              "computed": true,
              "description": "The name of the database where the AWS Glue table exists.",
              "description_kind": "plain",
              "type": "string"
            },
            "table_name": {
              "computed": true,
              "description": "The name of the AWS Glue table.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::DataQualityRuleset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueDataQualityRulesetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueDataQualityRuleset), &result)
	return &result
}
