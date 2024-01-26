package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2LoadBalancer = `{
  "block": {
    "attributes": {
      "canonical_hosted_zone_id": {
        "computed": true,
        "description": "The ID of the Amazon Route 53 hosted zone associated with the load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description": "The public DNS name of the load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "enforce_security_group_inbound_rules_on_private_link_traffic": {
        "computed": true,
        "description": "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through PrivateLink",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description": "The type of IP addresses used by the subnets for your load balancer. The possible values are ipv4 (for IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses).",
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the load balancer.",
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
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "load_balancer_full_name": {
        "computed": true,
        "description": "The full name of the load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancer_name": {
        "computed": true,
        "description": "The name of the load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "scheme": {
        "computed": true,
        "description": "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet. The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer. The default is an Internet-facing load balancer.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "The IDs of the security groups for the load balancer.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_mappings": {
        "computed": true,
        "description": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "i_pv_6_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_i_pv_4_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "subnets": {
        "computed": true,
        "description": "The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.",
        "description_kind": "plain",
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
      "type": {
        "computed": true,
        "description": "The type of load balancer. The default is application.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ElasticLoadBalancingV2::LoadBalancer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticloadbalancingv2LoadBalancerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2LoadBalancer), &result)
	return &result
}
