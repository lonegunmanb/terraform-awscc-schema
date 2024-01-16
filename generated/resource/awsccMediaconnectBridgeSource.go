package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectBridgeSource = `{
  "block": {
    "attributes": {
      "bridge_arn": {
        "description": "The Amazon Resource Number (ARN) of the bridge.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "flow_source": {
        "computed": true,
        "description": "The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "flow_arn": {
              "description": "The ARN of the cloud flow used as a source of this bridge.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "flow_vpc_interface_attachment": {
              "computed": true,
              "description": "The name of the VPC interface attachment to use for this source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "vpc_interface_name": {
                    "computed": true,
                    "description": "The name of the VPC interface to use for this resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the source.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_source": {
        "computed": true,
        "description": "The source of the bridge. A network source originates at your premises.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "multicast_ip": {
              "description": "The network source multicast IP.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_name": {
              "description": "The network source's gateway network name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description": "The network source port.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "protocol": {
              "description": "The network source protocol.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::MediaConnect::BridgeSource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectBridgeSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectBridgeSource), &result)
	return &result
}
