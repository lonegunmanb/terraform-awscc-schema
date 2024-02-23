package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Eip = `{
  "block": {
    "attributes": {
      "allocation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The network (` + "`" + `` + "`" + `vpc` + "`" + `` + "`" + `).\n If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance.\n  Updates to the ` + "`" + `` + "`" + `InstanceId` + "`" + `` + "`" + ` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_border_group": {
        "computed": true,
        "description": "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.\n Use [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_ipv_4_pool": {
        "computed": true,
        "description": "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.\n  Updates to the ` + "`" + `` + "`" + `PublicIpv4Pool` + "`" + `` + "`" + ` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the Elastic IP address.\n  Updates to the ` + "`" + `` + "`" + `Tags` + "`" + `` + "`" + ` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "transfer_address": {
        "computed": true,
        "description": "The Elastic IP address you are accepting for transfer. You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#transfer-EIPs-intro) in the *Amazon Virtual Private Cloud User Guide*.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::EIP",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2EipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Eip), &result)
	return &result
}
