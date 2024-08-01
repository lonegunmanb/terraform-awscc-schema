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
        "description": "Note: Internal load balancers must use the ` + "`" + `` + "`" + `ipv4` + "`" + `` + "`" + ` IP address type.\n [Application Load Balancers] The IP address type. The possible values are ` + "`" + `` + "`" + `ipv4` + "`" + `` + "`" + ` (for only IPv4 addresses), ` + "`" + `` + "`" + `dualstack` + "`" + `` + "`" + ` (for IPv4 and IPv6 addresses), and ` + "`" + `` + "`" + `dualstack-without-public-ipv4` + "`" + `` + "`" + ` (for IPv6 only public addresses, with private IPv4 and IPv6 addresses).\n [Network Load Balancers] The IP address type. The possible values are ` + "`" + `` + "`" + `ipv4` + "`" + `` + "`" + ` (for only IPv4 addresses) and ` + "`" + `` + "`" + `dualstack` + "`" + `` + "`" + ` (for IPv4 and IPv6 addresses). You can?t specify ` + "`" + `` + "`" + `dualstack` + "`" + `` + "`" + ` for a load balancer with a UDP or TCP_UDP listener.\n [Gateway Load Balancers] The IP address type. The possible values are ` + "`" + `` + "`" + `ipv4` + "`" + `` + "`" + ` (for only IPv4 addresses) and ` + "`" + `` + "`" + `dualstack` + "`" + `` + "`" + ` (for IPv4 and IPv6 addresses).",
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
              "description": "The name of the attribute.\n The following attributes are supported by all load balancers:\n  +   ` + "`" + `` + "`" + `deletion_protection.enabled` + "`" + `` + "`" + ` - Indicates whether deletion protection is enabled. The value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `load_balancing.cross_zone.enabled` + "`" + `` + "`" + ` - Indicates whether cross-zone load balancing is enabled. The possible values are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default for Network Load Balancers and Gateway Load Balancers is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default for Application Load Balancers is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, and cannot be changed.\n  \n The following attributes are supported by both Application Load Balancers and Network Load Balancers:\n  +   ` + "`" + `` + "`" + `access_logs.s3.enabled` + "`" + `` + "`" + ` - Indicates whether access logs are enabled. The value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `access_logs.s3.bucket` + "`" + `` + "`" + ` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.\n  +   ` + "`" + `` + "`" + `access_logs.s3.prefix` + "`" + `` + "`" + ` - The prefix for the location in the S3 bucket for the access logs.\n  +   ` + "`" + `` + "`" + `ipv6.deny_all_igw_traffic` + "`" + `` + "`" + ` - Blocks internet gateway (IGW) access to the load balancer. It is set to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` for internet-facing load balancers and ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for internal load balancers, preventing unintended access to your internal load balancer through an internet gateway.\n  \n The following attributes are supported by only Application Load Balancers:\n  +   ` + "`" + `` + "`" + `idle_timeout.timeout_seconds` + "`" + `` + "`" + ` - The idle timeout value, in seconds. The valid range is 1-4000 seconds. The default is 60 seconds.\n  +   ` + "`" + `` + "`" + `client_keep_alive.seconds` + "`" + `` + "`" + ` - The client keep alive value, in seconds. The valid range is 60-604800 seconds. The default is 3600 seconds.\n  +   ` + "`" + `` + "`" + `connection_logs.s3.enabled` + "`" + `` + "`" + ` - Indicates whether connection logs are enabled. The value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `connection_logs.s3.bucket` + "`" + `` + "`" + ` - The name of the S3 bucket for the connection logs. This attribute is required if connection logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.\n  +   ` + "`" + `` + "`" + `connection_logs.s3.prefix` + "`" + `` + "`" + ` - The prefix for the location in the S3 bucket for the connection logs.\n  +   ` + "`" + `` + "`" + `routing.http.desync_mitigation_mode` + "`" + `` + "`" + ` - Determines how the load balancer handles requests that might pose a security risk to your application. The possible values are ` + "`" + `` + "`" + `monitor` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `defensive` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `strictest` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `defensive` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `routing.http.drop_invalid_header_fields.enabled` + "`" + `` + "`" + ` - Indicates whether HTTP headers with invalid header fields are removed by the load balancer (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `) or routed to targets (` + "`" + `` + "`" + `false` + "`" + `` + "`" + `). The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `routing.http.preserve_host_header.enabled` + "`" + `` + "`" + ` - Indicates whether the Application Load Balancer should preserve the ` + "`" + `` + "`" + `Host` + "`" + `` + "`" + ` header in the HTTP request and send it to the target without any change. The possible values are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `routing.http.x_amzn_tls_version_and_cipher_suite.enabled` + "`" + `` + "`" + ` - Indicates whether the two headers (` + "`" + `` + "`" + `x-amzn-tls-version` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `x-amzn-tls-cipher-suite` + "`" + `` + "`" + `), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. The ` + "`" + `` + "`" + `x-amzn-tls-version` + "`" + `` + "`" + ` header has information about the TLS protocol version negotiated with the client, and the ` + "`" + `` + "`" + `x-amzn-tls-cipher-suite` + "`" + `` + "`" + ` header has information about the cipher suite negotiated with the client. Both headers are in OpenSSL format. The possible values for the attribute are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `routing.http.xff_client_port.enabled` + "`" + `` + "`" + ` - Indicates whether the ` + "`" + `` + "`" + `X-Forwarded-For` + "`" + `` + "`" + ` header should preserve the source port that the client used to connect to the load balancer. The possible values are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `routing.http.xff_header_processing.mode` + "`" + `` + "`" + ` - Enables you to modify, preserve, or remove the ` + "`" + `` + "`" + `X-Forwarded-For` + "`" + `` + "`" + ` header in the HTTP request before the Application Load Balancer sends the request to the target. The possible values are ` + "`" + `` + "`" + `append` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `preserve` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `remove` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `append` + "`" + `` + "`" + `.\n  +  If the value is ` + "`" + `` + "`" + `append` + "`" + `` + "`" + `, the Application Load Balancer adds the client IP address (of the last hop) to the ` + "`" + `` + "`" + `X-Forwarded-For` + "`" + `` + "`" + ` header in the HTTP request before it sends it to targets.\n  +  If the value is ` + "`" + `` + "`" + `preserve` + "`" + `` + "`" + ` the Application Load Balancer preserves the ` + "`" + `` + "`" + `X-Forwarded-For` + "`" + `` + "`" + ` header in the HTTP request, and sends it to targets without any change.\n  +  If the value is ` + "`" + `` + "`" + `remove` + "`" + `` + "`" + `, the Application Load Balancer removes the ` + "`" + `` + "`" + `X-Forwarded-For` + "`" + `` + "`" + ` header in the HTTP request before it sends it to targets.\n  \n  +   ` + "`" + `` + "`" + `routing.http2.enabled` + "`" + `` + "`" + ` - Indicates whether HTTP/2 is enabled. The possible values are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens.\n  +   ` + "`" + `` + "`" + `waf.fail_open.enabled` + "`" + `` + "`" + ` - Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. The possible values are ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  \n The following attributes are supported by only Network Load Balancers:\n  +   ` + "`" + `` + "`" + `dns_record.client_routing_policy` + "`" + `` + "`" + ` - Indicates how traffic is distributed among the load balancer Availability Zones. The possible values are ` + "`" + `` + "`" + `availability_zone_affinity` + "`" + `` + "`" + ` with 100 percent zonal affinity, ` + "`" + `` + "`" + `partial_availability_zone_affinity` + "`" + `` + "`" + ` with 85 percent zonal affinity, and ` + "`" + `` + "`" + `any_availability_zone` + "`" + `` + "`" + ` with 0 percent zonal affinity.",
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
        "description": "The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP addresses for your subnets.",
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
        "description": "The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.",
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
