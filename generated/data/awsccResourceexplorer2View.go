package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResourceexplorer2View = `{
  "block": {
    "attributes": {
      "filters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filter_string": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "included_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "scope": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "view_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "view_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ResourceExplorer2::View",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResourceexplorer2ViewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResourceexplorer2View), &result)
	return &result
}
