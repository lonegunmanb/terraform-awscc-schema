package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resolver rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps",
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
      "name": {
        "computed": true,
        "description": "The name for the Resolver rule",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resolver_endpoint_id": {
        "computed": true,
        "description": "The ID of the endpoint that the rule is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resolver_rule_id": {
        "computed": true,
        "description": "The ID of the endpoint that the rule is associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_type": {
        "description": "When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
      "target_ips": {
        "computed": true,
        "description": "An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip": {
              "computed": true,
              "description": "One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv_6": {
              "computed": true,
              "description": "One IPv6 address that you want to forward DNS queries to. You can specify only IPv6 addresses. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port at Ip that you want to forward DNS queries to. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "The protocol that you want to use to forward DNS queries. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_name_indication": {
              "computed": true,
              "description": "The SNI of the target name servers for DoH/DoH-FIPS outbound endpoints",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Route53Resolver::ResolverRule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverResolverRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverRule), &result)
	return &result
}
