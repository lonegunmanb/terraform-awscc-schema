package data

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
        "type": "string"
      },
      "consumption_configuration": {
        "computed": true,
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
                    "type": "bool"
                  },
                  "max_time_to_live_in_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "provisional_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_time_to_live_in_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "renew_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "entitlements": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_check_in": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "max_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "overage": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "unit": {
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
          "nesting_mode": "list"
        }
      },
      "home_region": {
        "computed": true,
        "description": "Home region for the created license.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "issuer": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sign_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "license_name": {
        "computed": true,
        "description": "Name for the created license.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_name": {
        "computed": true,
        "description": "Product name for the created license.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_sku": {
        "computed": true,
        "description": "ProductSKU of the license.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "validity": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "begin": {
              "computed": true,
              "description": "Validity begin date for the license.",
              "description_kind": "plain",
              "type": "string"
            },
            "end": {
              "computed": true,
              "description": "Validity begin date for the license.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version": {
        "computed": true,
        "description": "The version of the license.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::LicenseManager::License",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLicensemanagerLicenseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerLicense), &result)
	return &result
}
