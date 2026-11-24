package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayNetworkFunctionAttachment = `{
  "block": {
    "attributes": {
      "availability_zone_ids": {
        "computed": true,
        "description": "The Availability Zone IDs for the attachment.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "availability_zones": {
        "computed": true,
        "description": "The Availability Zones for the attachment.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "creation_time": {
        "computed": true,
        "description": "The creation time of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The resource ID of the network function.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description": "The name of the network function service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
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
        "description": "The ID of the network function attachment.",
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
    "description": "The AWS::EC2::TransitGatewayNetworkFunctionAttachment type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayNetworkFunctionAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayNetworkFunctionAttachment), &result)
	return &result
}
