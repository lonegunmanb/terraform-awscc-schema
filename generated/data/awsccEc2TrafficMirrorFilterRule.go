package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TrafficMirrorFilterRule = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the Traffic Mirror Filter rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_cidr_block": {
        "computed": true,
        "description": "The destination CIDR block to assign to the Traffic Mirror rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_port_range": {
        "computed": true,
        "description": "The destination port range.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "from_port": {
              "computed": true,
              "description": "The first port in the Traffic Mirror port range.",
              "description_kind": "plain",
              "type": "number"
            },
            "to_port": {
              "computed": true,
              "description": "The last port in the Traffic Mirror port range.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule.",
        "description_kind": "plain",
        "type": "number"
      },
      "rule_action": {
        "computed": true,
        "description": "The action to take on the filtered traffic (accept/reject).",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_number": {
        "computed": true,
        "description": "The number of the Traffic Mirror rule.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_cidr_block": {
        "computed": true,
        "description": "The source CIDR block to assign to the Traffic Mirror Filter rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_port_range": {
        "computed": true,
        "description": "The source port range.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "from_port": {
              "computed": true,
              "description": "The first port in the Traffic Mirror port range.",
              "description_kind": "plain",
              "type": "number"
            },
            "to_port": {
              "computed": true,
              "description": "The last port in the Traffic Mirror port range.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the Traffic Mirror Filter rule.",
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
      "traffic_direction": {
        "computed": true,
        "description": "The direction of traffic (ingress/egress).",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_mirror_filter_id": {
        "computed": true,
        "description": "The ID of the filter that this rule is associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "traffic_mirror_filter_rule_id": {
        "computed": true,
        "description": "The ID of the Traffic Mirror Filter rule.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TrafficMirrorFilterRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TrafficMirrorFilterRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TrafficMirrorFilterRule), &result)
	return &result
}
