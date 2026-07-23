package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerContext = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the context.",
        "description_kind": "plain",
        "type": "string"
      },
      "context_name": {
        "description": "The name of the context. Must be unique to your account in an AWS Region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "context_type": {
        "description": "The context type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "When the context was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the context.",
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
      "last_modified_time": {
        "computed": true,
        "description": "When the context was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "A list of properties to add to the context.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "source": {
        "description": "The source type, ID, and URI.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_id": {
              "computed": true,
              "description": "The ID of the source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description": "The type of the source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_uri": {
              "description": "The URI of the source.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the context.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
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
    "description": "Resource type definition for AWS::SageMaker::Context. A context is a lineage tracking entity that represents a logical grouping of other tracking or experiment entities.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerContextSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerContext), &result)
	return &result
}
