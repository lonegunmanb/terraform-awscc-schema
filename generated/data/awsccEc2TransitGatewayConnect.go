package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayConnect = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "options": {
        "computed": true,
        "description": "The Connect attachment options.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "protocol": {
              "computed": true,
              "description": "The tunnel protocol.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "state": {
        "computed": true,
        "description": "The state of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the Connect attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of the transit gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "transport_transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the attachment from which the Connect attachment was created.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayConnect",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayConnectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayConnect), &result)
	return &result
}
