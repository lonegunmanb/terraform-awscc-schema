package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerCoreNetwork = `{
  "block": {
    "attributes": {
      "core_network_arn": {
        "computed": true,
        "description": "The ARN (Amazon resource name) of core network",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_id": {
        "computed": true,
        "description": "The Id of core network",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The creation time of core network",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of core network",
        "description_kind": "plain",
        "type": "string"
      },
      "edges": {
        "computed": true,
        "description": "The edges within a core network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "asn": {
              "computed": true,
              "description": "The ASN of a core network edge.",
              "description_kind": "plain",
              "type": "number"
            },
            "edge_location": {
              "computed": true,
              "description": "The Region where a core network edge is located.",
              "description_kind": "plain",
              "type": "string"
            },
            "inside_cidr_blocks": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "global_network_id": {
        "computed": true,
        "description": "The ID of the global network that your core network is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account": {
        "computed": true,
        "description": "Owner of the core network",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "Live policy document for the core network, you must provide PolicyDocument in Json Format",
        "description_kind": "plain",
        "type": "string"
      },
      "segments": {
        "computed": true,
        "description": "The segments within a core network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "edge_locations": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "Name of segment",
              "description_kind": "plain",
              "type": "string"
            },
            "shared_segments": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "state": {
        "computed": true,
        "description": "The state of core network",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the global network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::CoreNetwork",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerCoreNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerCoreNetwork), &result)
	return &result
}
