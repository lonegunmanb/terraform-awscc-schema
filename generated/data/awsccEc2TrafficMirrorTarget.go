package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TrafficMirrorTarget = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the Traffic Mirror target.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_load_balancer_endpoint_id": {
        "computed": true,
        "description": "The ID of the Gateway Load Balancer endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The network interface ID that is associated with the target.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_load_balancer_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": " The tags to assign to the Traffic Mirror target.",
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
      },
      "traffic_mirror_target_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TrafficMirrorTarget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TrafficMirrorTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TrafficMirrorTarget), &result)
	return &result
}
