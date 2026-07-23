package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueBlueprint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the blueprint.",
        "description_kind": "plain",
        "type": "string"
      },
      "blueprint_location": {
        "computed": true,
        "description": "Specifies a path in Amazon S3 where the blueprint is published.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_on": {
        "computed": true,
        "description": "The date and time the blueprint was registered.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the blueprint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_on": {
        "computed": true,
        "description": "The date and time the blueprint was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the blueprint.",
        "description_kind": "plain",
        "type": "string"
      },
      "parameter_spec": {
        "computed": true,
        "description": "A JSON string that indicates the list of parameter specifications for the blueprint.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the blueprint registration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to be applied to this blueprint.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Glue::Blueprint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueBlueprintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueBlueprint), &result)
	return &result
}
