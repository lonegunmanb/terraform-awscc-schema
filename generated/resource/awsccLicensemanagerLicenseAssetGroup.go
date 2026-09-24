package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerLicenseAssetGroup = `{
  "block": {
    "attributes": {
      "associated_license_asset_ruleset_ar_ns": {
        "description": "ARNs of associated license asset rulesets.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "License asset group description.",
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
      "license_asset_group_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the license asset group.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_asset_group_configurations": {
        "description": "License asset group configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "usage_dimension": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "name": {
        "description": "License asset group name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "License asset group properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
        "description": "Tags to add to the license asset group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "Resource Type definition for AWS::LicenseManager::LicenseAssetGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLicensemanagerLicenseAssetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerLicenseAssetGroup), &result)
	return &result
}
