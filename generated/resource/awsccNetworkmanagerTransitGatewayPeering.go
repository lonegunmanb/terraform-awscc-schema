package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerTransitGatewayPeering = `{
  "block": {
    "attributes": {
      "core_network_arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) of the core network that you want to peer a transit gateway to.",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_id": {
        "description": "The Id of the core network that you want to peer a transit gateway to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The creation time of the transit gateway peering",
        "description_kind": "plain",
        "type": "string"
      },
      "edge_location": {
        "computed": true,
        "description": "The location of the transit gateway peering",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "Peering owner account Id",
        "description_kind": "plain",
        "type": "string"
      },
      "peering_id": {
        "computed": true,
        "description": "The Id of the transit gateway peering",
        "description_kind": "plain",
        "type": "string"
      },
      "peering_type": {
        "computed": true,
        "description": "Peering type (TransitGatewayPeering)",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) of the resource that you will peer to a core network",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the transit gateway peering",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "transit_gateway_arn": {
        "description": "The ARN (Amazon Resource Name) of the transit gateway that you will peer to a core network",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_peering_attachment_id": {
        "computed": true,
        "description": "The ID of the TransitGatewayPeeringAttachment",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "AWS::NetworkManager::TransitGatewayPeering Resoruce Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerTransitGatewayPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerTransitGatewayPeering), &result)
	return &result
}
