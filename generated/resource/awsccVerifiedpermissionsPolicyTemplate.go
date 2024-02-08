package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsPolicyTemplate = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "statement": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::VerifiedPermissions::PolicyTemplate Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVerifiedpermissionsPolicyTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsPolicyTemplate), &result)
	return &result
}
