package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2IpamScope = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_authority_configuration": {
        "computed": true,
        "description": "External service configuration to connect your AWS IPAM scope.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "external_resource_identifier": {
              "computed": true,
              "description": "Resource identifier of the scope in the external service connecting to your AWS IPAM scope.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipam_scope_external_authority_type": {
              "computed": true,
              "description": "An external service connecting to your AWS IPAM scope.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IPAM this scope is a part of.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_id": {
        "description": "The Id of the IPAM this scope is a part of.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ipam_scope_id": {
        "computed": true,
        "description": "Id of the IPAM scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipam_scope_type": {
        "computed": true,
        "description": "Determines whether this scope contains publicly routable space or space for a private network",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "Is this one of the default scopes created with the IPAM.",
        "description_kind": "plain",
        "type": "bool"
      },
      "pool_count": {
        "computed": true,
        "description": "The number of pools that currently exist in this scope.",
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
      }
    },
    "description": "Resource Schema of AWS::EC2::IPAMScope Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2IpamScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2IpamScope), &result)
	return &result
}
