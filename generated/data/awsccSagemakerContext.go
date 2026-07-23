package data

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
        "computed": true,
        "description": "The name of the context. Must be unique to your account in an AWS Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "context_type": {
        "computed": true,
        "description": "The context type.",
        "description_kind": "plain",
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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "type": [
          "map",
          "string"
        ]
      },
      "source": {
        "computed": true,
        "description": "The source type, ID, and URI.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_id": {
              "computed": true,
              "description": "The ID of the source.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_type": {
              "computed": true,
              "description": "The type of the source.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_uri": {
              "computed": true,
              "description": "The URI of the source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Context",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerContextSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerContext), &result)
	return &result
}
