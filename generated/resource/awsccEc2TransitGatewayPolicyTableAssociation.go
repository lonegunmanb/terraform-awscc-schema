package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayPolicyTableAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the transit gateway policy table association.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "description": "The ID of transit gateway attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_policy_table_id": {
        "description": "The ID of transit gateway policy table.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "AWS::EC2::TransitGatewayPolicyTableAssociation Resource Definition",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayPolicyTableAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayPolicyTableAssociation), &result)
	return &result
}
