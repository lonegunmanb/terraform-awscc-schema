package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRekognitionDataset = `{
  "block": {
    "attributes": {
      "dataset_arn": {
        "computed": true,
        "description": "The ARN of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_type": {
        "computed": true,
        "description": "The type of the dataset. Specify TRAIN to create a training dataset. Specify TEST to create a test dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_arn": {
        "computed": true,
        "description": "The ARN of the project to which the dataset belongs.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Rekognition::Dataset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRekognitionDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRekognitionDataset), &result)
	return &result
}
