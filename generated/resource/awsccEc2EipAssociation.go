package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2EipAssociation = `{
  "block": {
    "attributes": {
      "allocation_id": {
        "computed": true,
        "description": "The allocation ID. This is required.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eip": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "eip_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance. The instance must have exactly one attached network interface. You can specify either the instance ID or the network interface ID, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_interface_id": {
        "computed": true,
        "description": "The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.\n You can specify either the instance ID or the network interface ID, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Associates an Elastic IP address with an instance or a network interface. Before you can use an Elastic IP address, you must allocate it to your account. For more information about working with Elastic IP addresses, see [Elastic IP address concepts and rules](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#vpc-eip-overview).\n You must specify ` + "`" + `` + "`" + `AllocationId` + "`" + `` + "`" + ` and either ` + "`" + `` + "`" + `InstanceId` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `NetworkInterfaceId` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `PrivateIpAddress` + "`" + `` + "`" + `.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2EipAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2EipAssociation), &result)
	return &result
}
