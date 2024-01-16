package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerConnectPeer = `{
  "block": {
    "attributes": {
      "bgp_options": {
        "computed": true,
        "description": "Bgp options for connect peer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "peer_asn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "configuration": {
        "computed": true,
        "description": "Configuration of the connect peer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bgp_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "core_network_address": {
                    "computed": true,
                    "description": "The address of a core network.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "core_network_asn": {
                    "computed": true,
                    "description": "The ASN of the Coret Network.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "peer_address": {
                    "computed": true,
                    "description": "The address of a core network Connect peer.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "peer_asn": {
                    "computed": true,
                    "description": "The ASN of the Connect peer.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "core_network_address": {
              "computed": true,
              "description": "The IP address of a core network.",
              "description_kind": "plain",
              "type": "string"
            },
            "inside_cidr_blocks": {
              "computed": true,
              "description": "The inside IP addresses used for a Connect peer configuration.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "peer_address": {
              "computed": true,
              "description": "The IP address of the Connect peer.",
              "description_kind": "plain",
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "The protocol used for a Connect peer configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "connect_attachment_id": {
        "description": "The ID of the attachment to connect.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connect_peer_id": {
        "computed": true,
        "description": "The ID of the Connect peer.",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_address": {
        "computed": true,
        "description": "The IP address of a core network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "core_network_id": {
        "computed": true,
        "description": "The ID of the core network.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Connect peer creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "edge_location": {
        "computed": true,
        "description": "The Connect peer Regions where edges are located.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "inside_cidr_blocks": {
        "computed": true,
        "description": "The inside IP addresses used for a Connect peer configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "peer_address": {
        "description": "The IP address of the Connect peer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the connect peer.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_arn": {
        "computed": true,
        "description": "The subnet ARN for the connect peer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "AWS::NetworkManager::ConnectPeer Resource Type Definition.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerConnectPeerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerConnectPeer), &result)
	return &result
}
