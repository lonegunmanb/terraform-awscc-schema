package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecretsmanagerSecretTargetAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secret_target_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecretsManager::SecretTargetAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecretsmanagerSecretTargetAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerSecretTargetAttachment), &result)
	return &result
}
