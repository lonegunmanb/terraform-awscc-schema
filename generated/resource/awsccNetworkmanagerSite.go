package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerSite = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The date and time that the device was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the site.",
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
      "location": {
        "computed": true,
        "description": "The location of the site.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The physical address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "latitude": {
              "computed": true,
              "description": "The latitude.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "longitude": {
              "computed": true,
              "description": "The longitude.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "site_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the site.",
        "description_kind": "plain",
        "type": "string"
      },
      "site_id": {
        "computed": true,
        "description": "The ID of the site.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the site.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the site.",
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
    "description": "The AWS::NetworkManager::Site type describes a site.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerSiteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerSite), &result)
	return &result
}
