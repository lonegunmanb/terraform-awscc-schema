package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2ClientVpnTargetNetworkAssociation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "The ID of the association between the Client VPN endpoint and the target network.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone to associate with the Client VPN endpoint. Used for Transit Gateway-based endpoints in place of SubnetId.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The ID of the Availability Zone to associate with the Client VPN endpoint. Used for Transit Gateway-based endpoints in place of SubnetId.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_vpn_endpoint_id": {
        "description": "The ID of the Client VPN endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet to associate with the Client VPN endpoint. Required for VPC-based endpoints. For Transit Gateway-based endpoints, use AvailabilityZone or AvailabilityZoneId instead.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2ClientVpnTargetNetworkAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2ClientVpnTargetNetworkAssociation), &result)
	return &result
}
