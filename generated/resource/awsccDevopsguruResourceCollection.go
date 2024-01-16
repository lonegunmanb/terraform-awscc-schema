package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsguruResourceCollection = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_collection_filter": {
        "description": "Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudformation": {
              "computed": true,
              "description": "CloudFormation resource for DevOps Guru to monitor",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "stack_names": {
                    "computed": true,
                    "description": "An array of CloudFormation stack names.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "tags": {
              "computed": true,
              "description": "Tagged resources for DevOps Guru to monitor",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_boundary_key": {
                    "computed": true,
                    "description": "A Tag key for DevOps Guru app boundary.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_values": {
                    "computed": true,
                    "description": "Tag values of DevOps Guru app boundary.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "resource_collection_type": {
        "computed": true,
        "description": "The type of ResourceCollection",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "This resource schema represents the ResourceCollection resource in the Amazon DevOps Guru.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsguruResourceCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsguruResourceCollection), &result)
	return &result
}
