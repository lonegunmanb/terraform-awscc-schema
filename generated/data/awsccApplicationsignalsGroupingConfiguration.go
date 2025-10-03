package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationsignalsGroupingConfiguration = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The identifier for the specified AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "grouping_attribute_definitions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_grouping_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "grouping_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "grouping_source_keys": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApplicationSignals::GroupingConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApplicationsignalsGroupingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationsignalsGroupingConfiguration), &result)
	return &result
}
