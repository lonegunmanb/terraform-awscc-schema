package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectconnectDirectConnectGatewayAssociation = `{
  "block": {
    "attributes": {
      "accept_direct_connect_gateway_association_proposal_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the role to accept the Direct Connect Gateway association proposal. Needs directconnect:AcceptDirectConnectGatewayAssociationProposal permissions.",
        "description_kind": "plain",
        "type": "string"
      },
      "allowed_prefixes_to_direct_connect_gateway": {
        "computed": true,
        "description": "The Amazon VPC prefixes to advertise to the Direct Connect gateway. This parameter is required when you create an association to a transit gateway.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "associated_gateway_id": {
        "computed": true,
        "description": "The ID or ARN of the virtual private gateway or transit gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description": "The ID of the Direct Connect gateway association.",
        "description_kind": "plain",
        "type": "string"
      },
      "direct_connect_gateway_id": {
        "computed": true,
        "description": "The ID or ARN of the Direct Connect gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DirectConnect::DirectConnectGatewayAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDirectconnectDirectConnectGatewayAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectconnectDirectConnectGatewayAssociation), &result)
	return &result
}
