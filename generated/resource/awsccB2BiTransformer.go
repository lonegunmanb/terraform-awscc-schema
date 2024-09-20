package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccB2BiTransformer = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "edi_type": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "x12_details": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "transaction_set": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
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
        "required": true
      },
      "file_format": {
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
      "mapping_template": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sample_document": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "transformer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transformer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::B2BI::Transformer Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccB2BiTransformerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccB2BiTransformer), &result)
	return &result
}
