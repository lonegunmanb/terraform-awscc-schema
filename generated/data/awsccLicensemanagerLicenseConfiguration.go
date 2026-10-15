package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerLicenseConfiguration = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the license configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "disassociate_when_not_found": {
        "computed": true,
        "description": "When true, disassociates a resource when software is uninstalled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "license_configuration_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the license configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_configuration_id": {
        "computed": true,
        "description": "Unique ID of the license configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_count": {
        "computed": true,
        "description": "Number of licenses managed by the license configuration.",
        "description_kind": "plain",
        "type": "number"
      },
      "license_count_hard_limit": {
        "computed": true,
        "description": "Indicates whether hard or soft license enforcement is used. Exceeding a hard limit blocks the launch of new instances.",
        "description_kind": "plain",
        "type": "bool"
      },
      "license_counting_type": {
        "computed": true,
        "description": "Dimension used to track the license inventory.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_expiry": {
        "computed": true,
        "description": "License configuration expiry as an epoch timestamp (seconds since 1970-01-01T00:00:00Z).",
        "description_kind": "plain",
        "type": "number"
      },
      "license_rules": {
        "computed": true,
        "description": "License rules. The syntax is #name=value (for example, #allowedTenancy=EC2-DedicatedHost).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Name of the license configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "Account ID of the owner of the license configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_information_list": {
        "computed": true,
        "description": "Product information.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "product_information_filter_list": {
              "computed": true,
              "description": "A product information filter list.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "product_information_filter_comparator": {
                    "computed": true,
                    "description": "Logical operator.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "product_information_filter_name": {
                    "computed": true,
                    "description": "Filter name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "product_information_filter_value": {
                    "computed": true,
                    "description": "Filter value.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "resource_type": {
              "computed": true,
              "description": "Resource type. The possible values are SSM_MANAGED | RDS.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description": "License configuration status.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to add to the license configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::LicenseManager::LicenseConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLicensemanagerLicenseConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerLicenseConfiguration), &result)
	return &result
}
