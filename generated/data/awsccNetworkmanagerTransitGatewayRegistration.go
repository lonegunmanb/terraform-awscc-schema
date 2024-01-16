package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerTransitGatewayRegistration = `{
  "block": {
    "attributes": {
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
      "transit_gateway_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the transit gateway.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::TransitGatewayRegistration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerTransitGatewayRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerTransitGatewayRegistration), &result)
	return &result
}
