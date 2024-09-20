package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccForecastDatasetGroup = `{
  "block": {
    "attributes": {
      "dataset_arns": {
        "computed": true,
        "description": "An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "dataset_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the dataset group to delete.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_group_name": {
        "description": "A name for the dataset group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "description": "The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags of Application Insights application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Represents a dataset group that holds a collection of related datasets",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccForecastDatasetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccForecastDatasetGroup), &result)
	return &result
}
