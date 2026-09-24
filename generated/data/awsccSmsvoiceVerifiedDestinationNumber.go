package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceVerifiedDestinationNumber = `{
  "block": {
    "attributes": {
      "created_timestamp": {
        "computed": true,
        "description": "The time when the verified destination phone number was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_phone_number": {
        "computed": true,
        "description": "The verified destination phone number, in E.164 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the verified destination phone number. PENDING means the phone number has not been verified yet; VERIFIED means it is verified and can receive messages.",
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
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "verified_destination_number_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the verified destination phone number.",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_destination_number_id": {
        "computed": true,
        "description": "The unique identifier for the verified destination phone number.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SMSVOICE::VerifiedDestinationNumber",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSmsvoiceVerifiedDestinationNumberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceVerifiedDestinationNumber), &result)
	return &result
}
