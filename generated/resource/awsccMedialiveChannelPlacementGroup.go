package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveChannelPlacementGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the channel placement group.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_placement_group_id": {
        "computed": true,
        "description": "Unique internal identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "channels": {
        "computed": true,
        "description": "List of channel IDs added to the channel placement group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_id": {
        "computed": true,
        "description": "The ID of the cluster the node is on.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the channel placement group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nodes": {
        "computed": true,
        "description": "List of nodes added to the channel placement group",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description": "The current state of the ChannelPlacementGroupState",
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
    "description": "Definition of AWS::MediaLive::ChannelPlacementGroup Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveChannelPlacementGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveChannelPlacementGroup), &result)
	return &result
}
