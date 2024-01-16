package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrcontainersVirtualCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_provider": {
        "description": "Container provider of the virtual cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "description": "The ID of the container cluster",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "info": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "eks_info": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "namespace": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "type": {
              "description": "The type of the container provider",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Id of the virtual cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the virtual cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this virtual cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Schema of AWS::EMRContainers::VirtualCluster Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEmrcontainersVirtualClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrcontainersVirtualCluster), &result)
	return &result
}
