package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsguruResourceCollection = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_collection_filter": {
        "computed": true,
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
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  },
                  "tag_values": {
                    "computed": true,
                    "description": "Tag values of DevOps Guru app boundary.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_collection_type": {
        "computed": true,
        "description": "The type of ResourceCollection",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DevOpsGuru::ResourceCollection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDevopsguruResourceCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsguruResourceCollection), &result)
	return &result
}
