package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamPrefixListResolverTarget = `{
  "block": {
    "attributes": {
      "desired_version": {
        "computed": true,
        "description": "The desired version of the Prefix List Resolver that this Target should synchronize with.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_prefix_list_resolver_id": {
        "description": "The Id of the IPAM Prefix List Resolver associated with this Target.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_prefix_list_resolver_target_arn": {
        "computed": true,
        "description": "Id of the IPAM Prefix List Resolver Target.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_prefix_list_resolver_target_id": {
        "computed": true,
        "description": "Id of the IPAM Prefix List Resolver Target.",
        "description_kind": "plain",
        "type": "string"
      },
      "prefix_list_id": {
        "description": "The Id of the Managed Prefix List.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_list_region": {
        "description": "The region that the Managed Prefix List is located in.",
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "track_latest_version": {
        "description": "Indicates whether this Target automatically tracks the latest version of the Prefix List Resolver.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      }
    },
    "description": "Resource Type definition for AWS::EC2::IPAMPrefixListResolverTarget",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamPrefixListResolverTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamPrefixListResolverTarget), &result)
	return &result
}
