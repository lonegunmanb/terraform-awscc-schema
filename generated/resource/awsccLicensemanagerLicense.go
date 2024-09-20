package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerLicense = `{
  "block": {
    "attributes": {
      "beneficiary": {
        "computed": true,
        "description": "Beneficiary of the license.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "consumption_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "borrow_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allow_early_check_in": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "max_time_to_live_in_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "provisional_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_time_to_live_in_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "renew_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "entitlements": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_check_in": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "max_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "overage": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "unit": {
              "description_kind": "plain",
              "required": true,
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
        "required": true
      },
      "home_region": {
        "description": "Home region for the created license.",
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
      "issuer": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sign_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "license_arn": {
        "computed": true,
        "description": "Amazon Resource Name is a unique name for each resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_metadata": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
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
      "license_name": {
        "description": "Name for the created license.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_name": {
        "description": "Product name for the created license.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "product_sku": {
        "computed": true,
        "description": "ProductSKU of the license.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "validity": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "begin": {
              "description": "Validity begin date for the license.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "end": {
              "description": "Validity begin date for the license.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "version": {
        "computed": true,
        "description": "The version of the license.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::LicenseManager::License",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLicensemanagerLicenseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerLicense), &result)
	return &result
}
