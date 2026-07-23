package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceRegistration = `{
  "block": {
    "attributes": {
      "created_timestamp": {
        "computed": true,
        "description": "The time when the registration was created, in UNIX epoch time format.",
        "description_kind": "plain",
        "type": "string"
      },
      "current_version_number": {
        "computed": true,
        "description": "The current version number of the registration.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "registration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the registration.",
        "description_kind": "plain",
        "type": "string"
      },
      "registration_id": {
        "computed": true,
        "description": "The unique identifier for the registration.",
        "description_kind": "plain",
        "type": "string"
      },
      "registration_status": {
        "computed": true,
        "description": "The status of the registration.",
        "description_kind": "plain",
        "type": "string"
      },
      "registration_type": {
        "description": "The type of registration form to create.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of tags (key and value pairs) to associate with the registration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The string value associated with the key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "A registration that has been created.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSmsvoiceRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceRegistration), &result)
	return &result
}
