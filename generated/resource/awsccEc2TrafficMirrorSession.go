package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TrafficMirrorSession = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the Traffic Mirror session.",
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
      "network_interface_id": {
        "description": "The ID of the source network interface.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the account that owns the Traffic Mirror session.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "packet_length": {
        "computed": true,
        "description": "The number of bytes in each packet to mirror.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "session_number": {
        "description": "The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the Traffic Mirror session.",
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
      },
      "traffic_mirror_filter_id": {
        "description": "The ID of a Traffic Mirror filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "traffic_mirror_session_id": {
        "computed": true,
        "description": "The ID of a Traffic Mirror session.",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_mirror_target_id": {
        "description": "The ID of a Traffic Mirror target.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_network_id": {
        "computed": true,
        "description": "The VXLAN ID for the Traffic Mirror session.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource schema for AWS::EC2::TrafficMirrorSession",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TrafficMirrorSessionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TrafficMirrorSession), &result)
	return &result
}
