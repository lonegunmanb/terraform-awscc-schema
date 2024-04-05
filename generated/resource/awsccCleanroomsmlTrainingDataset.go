package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsmlTrainingDataset = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "training_data": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_source": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "glue_data_source": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "catalog_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "database_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "table_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "schema": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "column_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "column_types": {
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
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "training_dataset_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::CleanRoomsML::TrainingDataset Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCleanroomsmlTrainingDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsmlTrainingDataset), &result)
	return &result
}
