package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2LoadBalancer = `{
  "block": {
    "attributes": {
      "canonical_hosted_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enforce_security_group_inbound_rules_on_private_link_traffic": {
        "computed": true,
        "description": "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.",
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
      "ip_address_type": {
        "computed": true,
        "description": "The IP address type. The possible values are ` + "`" + `` + "`" + `ipv4` + "`" + `` + "`" + ` (for IPv4 addresses) and ` + "`" + `` + "`" + `dualstack` + "`" + `` + "`" + ` (for IPv4 and IPv6 addresses). You can?t specify ` + "`" + `` + "`" + `dualstack` + "`" + `` + "`" + ` for a load balancer with a UDP or TCP_UDP listener.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_attributes": {
        "computed": true,
        "description": "The load balancer attributes.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name of the attribute.\n The following attributes are supported by all load balancers:\n  +   ` + "`" + `` + "`" + `deletion_protection.enabled` + "`" + `` + "`" + ` - Indicates whether deletion protection is enabled. The value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `load_balancing.cross_zone.enabled` + "`" + `` + "`" + ` - Indicates whether cross-zone load balancing is enabled. The possible values are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default for Network Load Balancers and Gateway Load Balancers is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default for Application Load Balancers is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, and cannot be changed.\n  \n The following attributes are supported by both Application Load Balancers and Network Load Balancers:\n  +   ` + "`" + `` + "`" + `access_logs.s3.enabled` + "`" + `` + "`" + ` - Indicates whether access logs are enabled. The value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `access_logs.s3.bucket` + "`" + `` + "`" + ` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and h",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the attribute.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "load_balancer_full_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with \"internal-\".\n If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheme": {
        "computed": true,
        "description": "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.\n The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.\n The default is an Internet-facing load balancer.\n You cannot specify a scheme for a Gateway Load Balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "[Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_mappings": {
        "computed": true,
        "description": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_id": {
              "computed": true,
              "description": "[Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "i_pv_6_address": {
              "computed": true,
              "description": "[Network Load Balancers] The IPv6 address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_i_pv_4_address": {
              "computed": true,
              "description": "[Network Load Balancers] The private IPv4 address for an internal load balancer.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description": "The ID of the subnet.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "subnets": {
        "computed": true,
        "description": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the load balancer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key of the tag.",
              "description_kind": "plain",
              "required": true,
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
      "type": {
        "computed": true,
        "description": "The type of load balancer. The default is ` + "`" + `` + "`" + `application` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticloadbalancingv2LoadBalancerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2LoadBalancer), &result)
	return &result
}
