package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveNode = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Node. It is automatically assigned when the Node is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_placement_groups": {
        "computed": true,
        "description": "An array of IDs. Each ID is one ChannelPlacementGroup that is associated with this Node.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_id": {
        "computed": true,
        "description": "The ID of the Cluster that the Node belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_state": {
        "computed": true,
        "description": "The current connection state of the Node.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The ARN of the EC2 instance hosting the Node.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The user-specified name of the Node.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_id": {
        "computed": true,
        "description": "The unique ID of the Node. Unique in the Cluster. The ID is the resource-id portion of the ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_interface_mappings": {
        "computed": true,
        "description": "An array of interface mappings for the Node.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "logical_interface_name": {
              "computed": true,
              "description": "The logical name for this interface.",
              "description_kind": "plain",
              "type": "string"
            },
            "network_interface_mode": {
              "computed": true,
              "description": "The network interface mode.",
              "description_kind": "plain",
              "type": "string"
            },
            "physical_interface_name": {
              "computed": true,
              "description": "The physical interface name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "role": {
        "computed": true,
        "description": "The role of the Node in the Cluster. ACTIVE means the Node is available for encoding. BACKUP means the Node is a redundant Node and might get used if an ACTIVE Node fails.",
        "description_kind": "plain",
        "type": "string"
      },
      "sdi_source_mappings": {
        "computed": true,
        "description": "An array of SDI source mappings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "card_number": {
              "computed": true,
              "description": "The card number.",
              "description_kind": "plain",
              "type": "number"
            },
            "channel_number": {
              "computed": true,
              "description": "The channel number.",
              "description_kind": "plain",
              "type": "number"
            },
            "sdi_source": {
              "computed": true,
              "description": "The SDI source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "state": {
        "computed": true,
        "description": "The current state of the Node.",
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
    "description": "Data Source schema for AWS::MediaLive::Node",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMedialiveNodeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveNode), &result)
	return &result
}
