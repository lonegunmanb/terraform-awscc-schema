package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayConnectPeer = `{
  "block": {
    "attributes": {
      "connect_peer_configuration": {
        "computed": true,
        "description": "The Connect peer details.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bgp_configurations": {
              "computed": true,
              "description": "The BGP configuration details.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bgp_status": {
                    "computed": true,
                    "description": "The BGP status.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "peer_address": {
                    "computed": true,
                    "description": "The interior BGP peer IP address for the appliance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "peer_asn": {
                    "computed": true,
                    "description": "The peer Autonomous System Number (ASN).",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "transit_gateway_address": {
                    "computed": true,
                    "description": "The interior BGP peer IP address for the transit gateway.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transit_gateway_asn": {
                    "computed": true,
                    "description": "The transit gateway Autonomous System Number (ASN).",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "inside_cidr_blocks": {
              "computed": true,
              "description": "The range of interior BGP peer IP addresses.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "peer_address": {
              "computed": true,
              "description": "The peer IP address (GRE outer IP address) on the appliance side of the Connect peer.",
              "description_kind": "plain",
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "The tunnel protocol.",
              "description_kind": "plain",
              "type": "string"
            },
            "transit_gateway_address": {
              "computed": true,
              "description": "The Connect peer IP address on the transit gateway side of the tunnel.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "creation_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the Connect peer.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the Connect Peer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws: .",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 256 Unicode characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the Connect attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_connect_peer_id": {
        "computed": true,
        "description": "The ID of the Connect peer.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayConnectPeer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayConnectPeerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayConnectPeer), &result)
	return &result
}
