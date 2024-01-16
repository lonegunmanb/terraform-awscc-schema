package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerCustomerGatewayAssociation = `{
  "block": {
    "attributes": {
      "customer_gateway_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the customer gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_id": {
        "computed": true,
        "description": "The ID of the device",
        "description_kind": "plain",
        "type": "string"
      },
      "global_network_id": {
        "computed": true,
        "description": "The ID of the global network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "link_id": {
        "computed": true,
        "description": "The ID of the link",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::CustomerGatewayAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerCustomerGatewayAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerCustomerGatewayAssociation), &result)
	return &result
}
