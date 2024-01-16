package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerCustomerGatewayAssociation = `{
  "block": {
    "attributes": {
      "customer_gateway_arn": {
        "description": "The Amazon Resource Name (ARN) of the customer gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "device_id": {
        "description": "The ID of the device",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "global_network_id": {
        "description": "The ID of the global network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "link_id": {
        "computed": true,
        "description": "The ID of the link",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerCustomerGatewayAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerCustomerGatewayAssociation), &result)
	return &result
}
