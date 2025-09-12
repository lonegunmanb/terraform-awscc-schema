package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceProtectConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the protect configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "country_rule_set": {
        "computed": true,
        "description": "An array of CountryRule containing the rules for the NumberCapability.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "country_code": {
                    "computed": true,
                    "description": "The two-letter ISO country code",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protect_status": {
                    "computed": true,
                    "description": "The types of protection that can be used.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "sms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "country_code": {
                    "computed": true,
                    "description": "The two-letter ISO country code",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protect_status": {
                    "computed": true,
                    "description": "The types of protection that can be used.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "voice": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "country_code": {
                    "computed": true,
                    "description": "The two-letter ISO country code",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protect_status": {
                    "computed": true,
                    "description": "The types of protection that can be used.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "When set to true deletion protection is enabled and protect configuration cannot be deleted. By default this is set to false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protect_configuration_id": {
        "computed": true,
        "description": "The unique identifier for the protect configuration.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::SMSVOICE::ProtectConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSmsvoiceProtectConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceProtectConfiguration), &result)
	return &result
}
