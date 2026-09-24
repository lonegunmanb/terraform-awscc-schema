package resource

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
        "description": "The verified destination phone number, in E.164 format.",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
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
    "description": "A destination phone number that has been registered for verification with AWS End User Messaging SMS. A newly created number is in PENDING status; it becomes VERIFIED only after the recipient supplies the one-time code out of band.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSmsvoiceVerifiedDestinationNumberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceVerifiedDestinationNumber), &result)
	return &result
}
