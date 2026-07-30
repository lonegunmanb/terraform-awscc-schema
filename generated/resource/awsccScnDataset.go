package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The Amazon Web Services Supply Chain instance identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The last modified time of the dataset.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the dataset.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description": "The namespace of the dataset.",
        "description_kind": "plain",
        "required": true,
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
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
                    "optional": true,
                    "type": "bool"
                  },
                  "name": {
                    "computed": true,
                    "description": "The dataset field name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The dataset field type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "name": {
              "computed": true,
              "description": "The name of the dataset schema.",
              "description_kind": "plain",
              "optional": true,
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
        "optional": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "Represents an AWS Supply Chain data lake dataset.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccScnDatasetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccScnDataset), &result)
	return &result
}
