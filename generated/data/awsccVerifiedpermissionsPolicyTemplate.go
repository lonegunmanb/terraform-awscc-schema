package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_template_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "statement": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::VerifiedPermissions::PolicyTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVerifiedpermissionsPolicyTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsPolicyTemplate), &result)
	return &result
}
