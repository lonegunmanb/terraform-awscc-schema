package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskServerlessCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_authentication": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sasl": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iam": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "cluster_name": {
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
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_configs": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_groups": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::MSK::ServerlessCluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskServerlessClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskServerlessCluster), &result)
	return &result
}
