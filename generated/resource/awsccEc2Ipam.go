package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Ipam = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_resource_discovery_association_id": {
        "computed": true,
        "description": "The Id of the default association to the default resource discovery, created with this IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_resource_discovery_id": {
        "computed": true,
        "description": "The Id of the default resource discovery, created with this IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_resource_discovery_organizational_unit_exclusions": {
        "computed": true,
        "description": "A set of organizational unit (OU) exclusions for the default resource discovery, created with this IPAM.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "organizations_entity_path": {
              "computed": true,
              "description": "An AWS Organizations entity path. Build the path for the OU(s) using AWS Organizations IDs separated by a '/'. Include all child OUs by ending the path with '/*'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_private_gua": {
        "computed": true,
        "description": "Enable provisioning of GUA space in private pools.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_id": {
        "computed": true,
        "description": "Id of the IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "operating_regions": {
        "computed": true,
        "description": "The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "region_name": {
              "computed": true,
              "description": "The name of the region.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "private_default_scope_id": {
        "computed": true,
        "description": "The Id of the default scope for publicly routable IP space, created with this IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_default_scope_id": {
        "computed": true,
        "description": "The Id of the default scope for publicly routable IP space, created with this IPAM.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_discovery_association_count": {
        "computed": true,
        "description": "The count of resource discoveries associated with this IPAM.",
        "description_kind": "plain",
        "type": "number"
      },
      "scope_count": {
        "computed": true,
        "description": "The number of scopes that currently exist in this IPAM.",
        "description_kind": "plain",
        "type": "number"
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "tier": {
        "computed": true,
        "description": "The tier of the IPAM.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Schema of AWS::EC2::IPAM Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Ipam), &result)
	return &result
}
