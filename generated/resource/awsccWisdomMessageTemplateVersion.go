package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomMessageTemplateVersion = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "message_template_arn": {
        "description": "The unqualified Amazon Resource Name (ARN) of the message template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "message_template_content_sha_256": {
        "computed": true,
        "description": "The content SHA256 of the message template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "message_template_version_arn": {
        "computed": true,
        "description": "The unqualified Amazon Resource Name (ARN) of the message template version.",
        "description_kind": "plain",
        "type": "string"
      },
      "message_template_version_number": {
        "computed": true,
        "description": "Current version number of the message template.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "A version for the specified customer-managed message template within the specified knowledge base.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWisdomMessageTemplateVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomMessageTemplateVersion), &result)
	return &result
}
