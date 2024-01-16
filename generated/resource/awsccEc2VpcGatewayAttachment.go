package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcGatewayAttachment = `{
  "block": {
    "attributes": {
      "attachment_type": {
        "computed": true,
        "description": "Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment ",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "internet_gateway_id": {
        "computed": true,
        "description": "The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_gateway_id": {
        "computed": true,
        "description": "The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCGatewayAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcGatewayAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcGatewayAttachment), &result)
	return &result
}
