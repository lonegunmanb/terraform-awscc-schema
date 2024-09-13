package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_ids": {
        "computed": true,
        "description": "The MediaLive Channels that are currently running on Nodes in this Cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_id": {
        "computed": true,
        "description": "The unique ID of the Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_type": {
        "computed": true,
        "description": "The hardware type for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_role_arn": {
        "computed": true,
        "description": "The IAM role your nodes will use.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The user-specified name of the Cluster to be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_settings": {
        "computed": true,
        "description": "On premises settings which will have the interface network mappings and default Output logical interface",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_route": {
              "computed": true,
              "description": "Default value if the customer does not define it in channel Output API",
              "description_kind": "plain",
              "type": "string"
            },
            "interface_mappings": {
              "computed": true,
              "description": "Network mappings for the cluster",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "logical_interface_name": {
                    "computed": true,
                    "description": "logical interface name, unique in the list",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_id": {
                    "computed": true,
                    "description": "Network Id to be associated with the logical interface name, can be duplicated in list",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "state": {
        "computed": true,
        "description": "The current state of the Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of key-value pairs.",
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
    "description": "Data Source schema for AWS::MediaLive::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMedialiveClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveCluster), &result)
	return &result
}
