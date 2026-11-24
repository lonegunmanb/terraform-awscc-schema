package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcPeeringConnection = `{
  "block": {
    "attributes": {
      "assume_role_region": {
        "computed": true,
        "description": "The Region code to use when calling Security Token Service (STS) to assume the PeerRoleArn, if provided. You can update this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "multipath_resilience_status": {
        "computed": true,
        "description": "The multipath resilience status of the VPC peering connection: enabled when both the requester and accepter sides allow multipath resilience, pending when only one side allows it, and disabled when neither side allows it.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_owner_id": {
        "computed": true,
        "description": "The AWS account ID of the owner of the accepter VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_region": {
        "computed": true,
        "description": "The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in another AWS account. You can update this property. To manage PeerVpcPeeringConnectionOptions on a peering connection with another AWS account, the role must allow ec2:DescribeVpcPeeringConnections and ec2:ModifyVpcPeeringConnectionOptions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_vpc_id": {
        "description": "The ID of the VPC with which you are creating the VPC peering connection. You must specify this parameter in the request.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_vpc_peering_connection_options": {
        "computed": true,
        "description": "The VPC peering connection options corresponding to the accepter side of the connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_dns_resolution_from_remote_vpc": {
              "computed": true,
              "description": "If true, enables a local VPC to resolve public DNS hostnames to private IP addresses when queried from instances in the peer VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_multipath_resilience": {
              "computed": true,
              "description": "If true, allows multipath resilience on this side of the peering connection. Multipath resilience is enabled only when both the requester and accepter sides allow it; MultipathResilienceStatus reports the result.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "multipath_resilience_mode": {
              "computed": true,
              "description": "The mode of multipath resilience on the peering connection. Applies when AllowMultipathResilience is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "description": "The ID of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_peering_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_peering_connection_options": {
        "computed": true,
        "description": "The VPC peering connection options corresponding to the requester side of the connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_dns_resolution_from_remote_vpc": {
              "computed": true,
              "description": "If true, enables a local VPC to resolve public DNS hostnames to private IP addresses when queried from instances in the peer VPC.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_multipath_resilience": {
              "computed": true,
              "description": "If true, allows multipath resilience on this side of the peering connection. Multipath resilience is enabled only when both the requester and accepter sides allow it; MultipathResilienceStatus reports the result.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "multipath_resilience_mode": {
              "computed": true,
              "description": "The mode of multipath resilience on the peering connection. Applies when AllowMultipathResilience is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCPeeringConnection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcPeeringConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcPeeringConnection), &result)
	return &result
}
