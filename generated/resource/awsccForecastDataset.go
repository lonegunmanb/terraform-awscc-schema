package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccForecastDataset = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_frequency": {
        "computed": true,
        "description": "Frequency of data collection. This parameter is required for RELATED_TIME_SERIES",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dataset_name": {
        "description": "A name for the dataset",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dataset_type": {
        "description": "The dataset type",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain": {
        "description": "The domain associated with the dataset",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "KMS key used to encrypt the Dataset data",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "computed": true,
                    "description": "Name of the dataset field",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "attribute_type": {
                    "computed": true,
                    "description": "Data type of the field",
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
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
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
    "description": "Resource Type Definition for AWS::Forecast::Dataset",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccForecastDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccForecastDataset), &result)
	return &result
}
