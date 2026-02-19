package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesCustomVerificationEmailTemplate = `{
  "block": {
    "attributes": {
      "failure_redirection_url": {
        "description": "The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "from_email_address": {
        "description": "The email address that the custom verification email is sent from.",
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
      "success_redirection_url": {
        "description": "The URL that the recipient of the verification email is sent to if his or her address is successfully verified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the tenant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the key-value tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the key-value tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "template_content": {
        "description": "The content of the custom verification email. The total size of the email must be less than 10 MB. The message body may contain HTML, with some limitations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_name": {
        "description": "The name of the custom verification email template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_subject": {
        "description": "The subject line of the custom verification email.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SES::CustomVerificationEmailTemplate.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesCustomVerificationEmailTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesCustomVerificationEmailTemplate), &result)
	return &result
}
