package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerTransitGatewayRegistration = `{
  "block": {
    "attributes": {
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
      "transit_gateway_arn": {
        "description": "The Amazon Resource Name (ARN) of the transit gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::NetworkManager::TransitGatewayRegistration type registers a transit gateway in your global network. The transit gateway can be in any AWS Region, but it must be owned by the same AWS account that owns the global network. You cannot register a transit gateway in more than one global network.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerTransitGatewayRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerTransitGatewayRegistration), &result)
	return &result
}
