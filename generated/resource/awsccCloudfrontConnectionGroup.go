package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontConnectionGroup = `{
  "block": {
    "attributes": {
      "anycast_ip_list_id": {
        "computed": true,
        "description": "The ID of the Anycast static IP list.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "e_tag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "Whether the connection group is enabled.",
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
      "ipv_6_enabled": {
        "computed": true,
        "description": "IPv6 is enabled for the connection group.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_default": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the connection group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A complex type that contains zero or more ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` elements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string that contains ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` key.\n The string length should be between 1 and 128 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string that contains an optional ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` value.\n The string length should be between 0 and 256 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
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
    "description": "The connection group for your distribution tenants. When you first create a distribution tenant and you don't specify a connection group, CloudFront will automatically create a default connection group for you. When you create a new distribution tenant and don't specify a connection group, the default one will be associated with your distribution tenant.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontConnectionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontConnectionGroup), &result)
	return &result
}
