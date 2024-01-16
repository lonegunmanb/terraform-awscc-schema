package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectPhoneNumber = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description": "The phone number e164 address.",
        "description_kind": "plain",
        "type": "string"
      },
      "country_code": {
        "computed": true,
        "description": "The phone number country code.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the phone number.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "phone_number_arn": {
        "computed": true,
        "description": "The phone number ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "prefix": {
        "computed": true,
        "description": "The phone number prefix.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_phone_number_arn": {
        "computed": true,
        "description": "The source phone number arn.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "target_arn": {
        "computed": true,
        "description": "The ARN of the target the phone number is claimed to.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The phone number type",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::PhoneNumber",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectPhoneNumberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectPhoneNumber), &result)
	return &result
}
