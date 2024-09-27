package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "options": {
        "computed": true,
        "description": "The options for the transit gateway vpc attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "appliance_mode_support": {
              "computed": true,
              "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dns_support": {
              "computed": true,
              "description": "Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv_6_support": {
              "computed": true,
              "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_referencing_support": {
              "computed": true,
              "description": "Indicates whether to enable Security Group referencing support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "subnet_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::TransitGatewayAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayAttachment), &result)
	return &result
}
