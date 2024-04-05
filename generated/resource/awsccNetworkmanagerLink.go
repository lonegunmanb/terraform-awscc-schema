package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerLink = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "description": "The Bandwidth for the link.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "download_speed": {
              "computed": true,
              "description": "Download speed in Mbps.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "upload_speed": {
              "computed": true,
              "description": "Upload speed in Mbps.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The date and time that the device was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the link.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_network_id": {
        "description": "The ID of the global network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "link_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the link.",
        "description_kind": "plain",
        "type": "string"
      },
      "link_id": {
        "computed": true,
        "description": "The ID of the link.",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "The provider of the link.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "site_id": {
        "description": "The ID of the site",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the link.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the link.",
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
      "type": {
        "computed": true,
        "description": "The type of the link.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::NetworkManager::Link type describes a link.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerLinkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerLink), &result)
	return &result
}
