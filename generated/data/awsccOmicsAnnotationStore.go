package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOmicsAnnotationStore = `{
  "block": {
    "attributes": {
      "annotation_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reference": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reference_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sse_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "store_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "store_format": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "store_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "tsv_store_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "annotation_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "format_to_header": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "schema": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      [
                        "list",
                        "string"
                      ]
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "store_size_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Omics::AnnotationStore",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOmicsAnnotationStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsAnnotationStore), &result)
	return &result
}
