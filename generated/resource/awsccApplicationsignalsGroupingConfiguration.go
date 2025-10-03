package resource

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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_grouping_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "grouping_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "grouping_source_keys": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ApplicationSignals::GroupingConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApplicationsignalsGroupingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationsignalsGroupingConfiguration), &result)
	return &result
}
