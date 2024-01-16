package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "dns_support": {
              "computed": true,
              "description": "Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "type": "string"
            },
            "ipv_6_support": {
              "computed": true,
              "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "type": "string"
            },
            "security_group_referencing_support": {
              "computed": true,
              "description": "Indicates whether to enable Security Group referencing support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
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
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::TransitGatewayAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2TransitGatewayAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayAttachment), &result)
	return &result
}
