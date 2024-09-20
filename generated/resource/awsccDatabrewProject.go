package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatabrewProject = `{
  "block": {
    "attributes": {
      "dataset_name": {
        "description": "Dataset name",
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
      "name": {
        "description": "Project name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recipe_name": {
        "description": "Recipe name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "Role arn",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "number"
            },
            "type": {
              "computed": true,
              "description": "Sample type",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
      }
    },
    "description": "Resource schema for AWS::DataBrew::Project.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatabrewProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewProject), &result)
	return &result
}
