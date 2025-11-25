package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnConcentrator = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the VPN concentrator.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the transit gateway attachment",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of the transit gateway",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of VPN concentrator",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_concentrator_id": {
        "computed": true,
        "description": "The provider-assigned unique ID for this managed resource",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPNConcentrator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpnConcentratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnConcentrator), &result)
	return &result
}
