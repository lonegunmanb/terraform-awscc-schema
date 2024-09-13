package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveNetwork = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Network.",
        "description_kind": "plain",
        "type": "string"
      },
      "associated_cluster_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_pools": {
        "description": "The list of IP address cidr pools for the network",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr": {
              "computed": true,
              "description": "IP address cidr pool",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "name": {
        "description": "The user-specified name of the Network to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_id": {
        "computed": true,
        "description": "The unique ID of the Network.",
        "description_kind": "plain",
        "type": "string"
      },
      "routes": {
        "computed": true,
        "description": "The routes for the network",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr": {
              "computed": true,
              "description": "Ip address cidr",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gateway": {
              "computed": true,
              "description": "IP address for the route packet paths",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "state": {
        "computed": true,
        "description": "The current state of the Network.",
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
    "description": "Resource schema for AWS::MediaLive::Network.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveNetwork), &result)
	return &result
}
