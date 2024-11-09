package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncChannelNamespace = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "AppSync Api Id that this Channel Namespace belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "channel_namespace_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Channel Namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "code_handlers": {
        "computed": true,
        "description": "String of APPSYNC_JS code to be used by the handlers.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_s3_location": {
        "computed": true,
        "description": "The Amazon S3 endpoint where the code is located.",
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
        "description": "Namespace indentifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "publish_auth_modes": {
        "computed": true,
        "description": "List of AuthModes supported for Publish operations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_type": {
              "computed": true,
              "description": "Security configuration for your AppSync API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "subscribe_auth_modes": {
        "computed": true,
        "description": "List of AuthModes supported for Subscribe operations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_type": {
              "computed": true,
              "description": "Security configuration for your AppSync API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this AppSync API.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag. You can specify a maximum of 128 characters for a tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
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
    "description": "Resource schema for AppSync ChannelNamespace",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncChannelNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncChannelNamespace), &result)
	return &result
}
