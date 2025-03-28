package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseDataset = `{
  "block": {
    "attributes": {
      "dataset_arn": {
        "computed": true,
        "description": "The ARN of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_description": {
        "computed": true,
        "description": "A description about the dataset, and its functionality.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dataset_id": {
        "computed": true,
        "description": "The ID of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_name": {
        "description": "The name of the dataset.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dataset_source": {
        "description": "The data source for the dataset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_detail": {
              "computed": true,
              "description": "The details of the dataset source associated with the dataset.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kendra": {
                    "computed": true,
                    "description": "Contains details about the Kendra dataset source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "knowledge_base_arn": {
                          "computed": true,
                          "description": "The knowledgeBaseArn details for the Kendra dataset source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "The roleARN details for the Kendra dataset source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "source_format": {
              "description": "The format of the dataset source associated with the dataset.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_type": {
              "description": "The type of data source for the dataset.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
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
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::IoTSiteWise::Dataset.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseDataset), &result)
	return &result
}
