package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatabrewProject = `{
  "block": {
    "attributes": {
      "dataset_name": {
        "computed": true,
        "description": "Dataset name",
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
        "description": "Project name",
        "description_kind": "plain",
        "type": "string"
      },
      "recipe_name": {
        "computed": true,
        "description": "Recipe name",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Role arn",
        "description_kind": "plain",
        "type": "string"
      },
      "sample": {
        "computed": true,
        "description": "Sample",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "size": {
              "computed": true,
              "description": "Sample size",
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description": "Sample type",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::DataBrew::Project",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatabrewProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewProject), &result)
	return &result
}
