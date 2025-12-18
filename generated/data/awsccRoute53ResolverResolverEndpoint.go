package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resolver endpoint, such as arn:aws:route53resolver:us-east-1:123456789012:resolver-endpoint/resolver-endpoint-a1bzhi.",
        "description_kind": "plain",
        "type": "string"
      },
      "direction": {
        "computed": true,
        "description": "Indicates whether the Resolver endpoint allows inbound or outbound DNS queries:\n- INBOUND: allows DNS queries to your VPC from your network \n- OUTBOUND: allows DNS queries from your VPC to your network \n- INBOUND_DELEGATION: allows DNS queries to your VPC from your network with authoritative answers from private hosted zones",
        "description_kind": "plain",
        "type": "string"
      },
      "host_vpc_id": {
        "computed": true,
        "description": "The ID of the VPC that you want to create the resolver endpoint in.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address_count": {
        "computed": true,
        "description": "The number of IP addresses that the resolver endpoint can use for DNS queries.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_addresses": {
        "computed": true,
        "description": "The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints). The subnet ID uniquely identifies a VPC.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip": {
              "computed": true,
              "description": "The IPv4 address that you want to use for DNS queries.",
              "description_kind": "plain",
              "type": "string"
            },
            "ipv_6": {
              "computed": true,
              "description": "The IPv6 address that you want to use for DNS queries.",
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_id": {
              "computed": true,
              "description": "The ID of the subnet that contains the IP address.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description": "A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description": "The ARN (Amazon Resource Name) for the Outpost.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_instance_type": {
        "computed": true,
        "description": "The Amazon EC2 instance type.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocols": {
        "computed": true,
        "description": "Protocols used for the endpoint. DoH-FIPS is applicable for inbound endpoints only.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resolver_endpoint_id": {
        "computed": true,
        "description": "The ID of the resolver endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_endpoint_type": {
        "computed": true,
        "description": "The Resolver endpoint IP address type.",
        "description_kind": "plain",
        "type": "string"
      },
      "rni_enhanced_metrics_enabled": {
        "computed": true,
        "description": "Specifies whether RNI enhanced metrics are enabled for the Resolver Endpoints. When set to true, one-minute granular metrics are published in CloudWatch for each RNI associated with this endpoint. When set to false, metrics are not published. Default is false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The ID of one or more security groups that control access to this VPC. The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name for the tag. For example, if you want to associate Resolver resources with the account IDs of your customers for billing purposes, the value of Key might be account-id.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. For example, if Key is account-id, then Value might be the ID of the customer account that you're creating the resource for.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_name_server_metrics_enabled": {
        "computed": true,
        "description": "Specifies whether target name server metrics are enabled for the Outbound Resolver Endpoint. When set to true, one-minute granular metrics are published in CloudWatch for each target name server associated with this endpoint. When set to false, metrics are not published. Default is false.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::Route53Resolver::ResolverEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ResolverResolverEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverEndpoint), &result)
	return &result
}
