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
      "dataset_config": {
        "computed": true,
        "description": "The configuration for the dataset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "session": {
              "computed": true,
              "description": "The session configuration for a SESSION dataset.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "session_end_time": {
                    "computed": true,
                    "description": "The end time of the session as an ISO 8601 UTC instant, for example 2024-12-31T23:59:59Z.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_start_time": {
                    "computed": true,
                    "description": "The start time of the session as an ISO 8601 UTC instant, for example 2024-01-01T00:00:00Z.",
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
      "dataset_description": {
        "computed": true,
        "description": "A description about the dataset, and its functionality.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dataset_id": {
        "computed": true,
        "description": "The ID of the dataset. For workspace-scoped datasets this is the workspace name and dataset ID joined by a slash, for example my-workspace/123e4567-e89b-42d3-a456-426614174000.",
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
        "computed": true,
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
              "computed": true,
              "description": "The format of the dataset source associated with the dataset.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description": "The type of data source for the dataset.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "dataset_type": {
        "computed": true,
        "description": "The type of the dataset.",
        "description_kind": "plain",
        "optional": true,
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
      },
      "workspace_name": {
        "computed": true,
        "description": "The name of the workspace associated with the dataset.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
