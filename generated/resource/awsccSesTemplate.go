package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesTemplate = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template": {
        "computed": true,
        "description": "The content of the email, composed of a subject line, an HTML part, and a text-only part",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "html_part": {
              "computed": true,
              "description": "The HTML body of the email.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subject_part": {
              "description": "The subject line of the email.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "template_name": {
              "computed": true,
              "description": "The name of the template.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "text_part": {
              "computed": true,
              "description": "The email body that is visible to recipients whose email clients do not display HTML content.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SES::Template",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesTemplate), &result)
	return &result
}
