package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectTrafficDistributionGroup = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description for the traffic distribution group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance that has been replicated.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "If this is the default traffic distribution group.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The name for the traffic distribution group.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the traffic distribution group.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "traffic_distribution_group_arn": {
        "computed": true,
        "description": "The identifier of the traffic distribution group.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::TrafficDistributionGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectTrafficDistributionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectTrafficDistributionGroup), &result)
	return &result
}
