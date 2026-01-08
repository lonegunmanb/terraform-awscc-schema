package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnConcentrator = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "description": "The key of the tag.\n Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.\n Constraints: Tag values are case-sensitive and accept a maximum of 256 Unicode characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "description": "The ID of the transit gateway associated with the VPN concentrator.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The type of VPN concentrator.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_concentrator_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Describes a VPN concentrator.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpnConcentratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnConcentrator), &result)
	return &result
}
