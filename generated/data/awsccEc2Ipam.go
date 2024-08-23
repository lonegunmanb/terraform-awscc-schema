package data

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
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_private_gua": {
        "computed": true,
        "description": "Enable provisioning of GUA space in private pools.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
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
      },
      "tier": {
        "computed": true,
        "description": "The tier of the IPAM.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::IPAM",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2IpamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Ipam), &result)
	return &result
}
