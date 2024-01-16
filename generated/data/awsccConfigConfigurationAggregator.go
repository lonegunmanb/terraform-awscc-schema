package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigConfigurationAggregator = `{
  "block": {
    "attributes": {
      "account_aggregation_sources": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "all_aws_regions": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "aws_regions": {
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
      "configuration_aggregator_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the aggregator.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_aggregator_name": {
        "computed": true,
        "description": "The name of the aggregator.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization_aggregation_source": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "all_aws_regions": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "aws_regions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags for the configuration aggregator.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Config::ConfigurationAggregator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigConfigurationAggregatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigConfigurationAggregator), &result)
	return &result
}
