package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamPool = `{
  "block": {
    "attributes": {
      "address_family": {
        "description": "The address family of the address space in this pool. Either IPv4 or IPv6.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "allocation_default_netmask_length": {
        "computed": true,
        "description": "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_max_netmask_length": {
        "computed": true,
        "description": "The maximum allowed netmask length for allocations made from this pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_min_netmask_length": {
        "computed": true,
        "description": "The minimum allowed netmask length for allocations made from this pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "allocation_resource_tags": {
        "computed": true,
        "description": "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_import": {
        "computed": true,
        "description": "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "aws_service": {
        "computed": true,
        "description": "Limits which service in Amazon Web Services that the pool can be used in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
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
      "ipam_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_pool_id": {
        "computed": true,
        "description": "Id of the IPAM Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_scope_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_scope_id": {
        "description": "The Id of the scope this pool is a part of.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_scope_type": {
        "computed": true,
        "description": "Determines whether this scope contains publicly routable space or space for a private network",
        "description_kind": "plain",
        "type": "string"
      },
      "locale": {
        "computed": true,
        "description": "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pool_depth": {
        "computed": true,
        "description": "The depth of this pool in the source pool hierarchy.",
        "description_kind": "plain",
        "type": "number"
      },
      "provisioned_cidrs": {
        "computed": true,
        "description": "A list of cidrs representing the address space available for allocation in this pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr": {
              "description": "Represents a single IPv4 or IPv6 CIDR",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "public_ip_source": {
        "computed": true,
        "description": "The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is ` + "`" + `byoip` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_advertisable": {
        "computed": true,
        "description": "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source_ipam_pool_id": {
        "computed": true,
        "description": "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_resource": {
        "computed": true,
        "description": "The resource associated with this pool's space. Depending on the ResourceType, setting a SourceResource changes which space can be provisioned in this pool and which types of resources can receive allocations",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_owner": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_region": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "state": {
        "computed": true,
        "description": "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "An explanation of how the pool arrived at it current state.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Schema of AWS::EC2::IPAMPool Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamPool), &result)
	return &result
}
