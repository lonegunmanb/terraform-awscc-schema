package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamResourceDiscovery = `{
  "block": {
    "attributes": {
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
      "ipam_resource_discovery_arn": {
        "computed": true,
        "description": "Amazon Resource Name (Arn) for the Resource Discovery.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_resource_discovery_id": {
        "computed": true,
        "description": "Id of the IPAM Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_resource_discovery_region": {
        "computed": true,
        "description": "The region the resource discovery is setup in. ",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
        "description_kind": "plain",
        "type": "bool"
      },
      "operating_regions": {
        "computed": true,
        "description": "The regions Resource Discovery is enabled for. Allows resource discoveries to be created in these regions, as well as enabling monitoring",
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
      "owner_id": {
        "computed": true,
        "description": "Owner Account ID of the Resource Discovery",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of this Resource Discovery.",
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
      }
    },
    "description": "Resource Schema of AWS::EC2::IPAMResourceDiscovery Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamResourceDiscoverySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamResourceDiscovery), &result)
	return &result
}
