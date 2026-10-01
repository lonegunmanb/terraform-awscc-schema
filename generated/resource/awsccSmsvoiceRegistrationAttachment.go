package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceRegistrationAttachment = `{
  "block": {
    "attributes": {
      "attachment_body": {
        "computed": true,
        "description": "The registration file to upload. The maximum file size is 1500KB and valid file extensions are PDF, JPEG and PNG.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "attachment_status": {
        "computed": true,
        "description": "The status of the registration attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "attachment_url": {
        "computed": true,
        "description": "A URL to the required registration file. For example, the URL to an MMS/shortcode form.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The time when the registration attachment was created, in UNIX epoch time format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "registration_attachment_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the registration attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "registration_attachment_id": {
        "computed": true,
        "description": "The unique identifier for the registration attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of tags (key and value pairs) to associate with the registration attachment.",
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
      },
      "uploaded_attachment_url": {
        "computed": true,
        "description": "The URL to the document that was uploaded as the registration attachment, as returned by the service.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SMSVOICE::RegistrationAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSmsvoiceRegistrationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceRegistrationAttachment), &result)
	return &result
}
