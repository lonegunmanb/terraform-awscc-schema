package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamPrefixListResolver = `{
  "block": {
    "attributes": {
      "address_family": {
        "computed": true,
        "description": "The address family of the address space in this Prefix List Resolver. Either IPv4 or IPv6.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM this Prefix List Resolver is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_id": {
        "computed": true,
        "description": "The Id of the IPAM this Prefix List Resolver is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_prefix_list_resolver_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM Prefix List Resolver",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_prefix_list_resolver_id": {
        "computed": true,
        "description": "Id of the IPAM Prefix List Resolver.",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description": "Rules define the business logic for selecting CIDRs from IPAM.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "conditions": {
              "computed": true,
              "description": "Two of the rule types allow you to add conditions to the rules. (1) For IPAM Pool CIDR rules, you can specify an ipamPoolId; if not specified, the rule will apply to all IPAM Pool CIDRs in the scope.  (2) For IPAM Resource CIDR rules, you can specify resourceId, resourceOwner, resourceRegion, cidr, or resourceTag.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description": "Condition for the IPAM Resource CIDR rule type.  CIDR (like 10.24.34.0/23).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ipam_pool_id": {
                    "computed": true,
                    "description": "Condition for the IPAM Pool CIDR rule type.  If not chosen, the resolver applies to all IPAM Pool CIDRs in the scope.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "operation": {
                    "computed": true,
                    "description": "Equals, Not equals, or Subnet Of.  The subnet-of operation only applies to cidr conditions.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_id": {
                    "computed": true,
                    "description": "Condition for the IPAM Resource CIDR rule type.  The unique ID of a resource (like vpc-1234567890abcdef0).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_owner": {
                    "computed": true,
                    "description": "Condition for the IPAM Resource CIDR rule type.  Resource owner (like 111122223333).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_region": {
                    "computed": true,
                    "description": "Condition for the IPAM Resource CIDR rule type.  Resource region (like us-east-1).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_tag": {
                    "computed": true,
                    "description": "Condition for the IPAM Resource CIDR rule type.  Resource Tag (like dev-vpc-1).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "ipam_scope_id": {
              "computed": true,
              "description": "This rule will only match resources that are in this IPAM Scope.",
              "description_kind": "plain",
              "type": "string"
            },
            "resource_type": {
              "computed": true,
              "description": "The resourceType property only applies to ipam-resource-cidr rules; this property specifies what type of resources this rule will apply to, such as VPCs or Subnets.",
              "description_kind": "plain",
              "type": "string"
            },
            "rule_type": {
              "computed": true,
              "description": "There are three rule types: (1) Static CIDR: A fixed list of CIDRs that don't change (like a manual list replicated across Regions). (2) IPAM pool CIDR: CIDRs from specific IPAM pools (like all CIDRs from your IPAM production pool).  (3) IPAM resource CIDR: CIDRs for AWS resources like VPCs, subnets, and EIPs within a specific IPAM scope.",
              "description_kind": "plain",
              "type": "string"
            },
            "static_cidr": {
              "computed": true,
              "description": "A fixed CIDR that doesn't change",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::IPAMPrefixListResolver",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2IpamPrefixListResolverSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamPrefixListResolver), &result)
	return &result
}
