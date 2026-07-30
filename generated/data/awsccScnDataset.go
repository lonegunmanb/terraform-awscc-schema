package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccScnDataset = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The creation time of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The Amazon Web Services Supply Chain instance identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The last modified time of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description": "The namespace of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "partition_spec": {
        "computed": true,
        "description": "The partition specification of the dataset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fields": {
              "computed": true,
              "description": "The partition fields.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the partition field.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transform": {
                    "computed": true,
                    "description": "The transformation of the partition field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "The type of partitioning transformation.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "schema": {
        "computed": true,
        "description": "The schema of the dataset.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fields": {
              "computed": true,
              "description": "The list of field details of the dataset schema.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "is_required": {
                    "computed": true,
                    "description": "Indicate if the field is required or not.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "name": {
                    "computed": true,
                    "description": "The dataset field name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The dataset field type.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The name of the dataset schema.",
              "description_kind": "plain",
              "type": "string"
            },
            "primary_keys": {
              "computed": true,
              "description": "The list of primary key fields for the dataset.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the primary key field.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags for the dataset.",
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
    "description": "Data Source schema for AWS::SCN::Dataset",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccScnDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccScnDataset), &result)
	return &result
}
