package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayPeeringAttachment = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time the transit gateway peering attachment was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_account_id": {
        "description": "The ID of the peer account",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_region": {
        "description": "Peer Region",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_transit_gateway_id": {
        "description": "The ID of the peer transit gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the transit gateway peering attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The status code.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The status message, if applicable.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags for the transit gateway peering attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
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
        "description": "The ID of the transit gateway peering attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "description": "The ID of the transit gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::EC2::TransitGatewayPeeringAttachment type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayPeeringAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayPeeringAttachment), &result)
	return &result
}
