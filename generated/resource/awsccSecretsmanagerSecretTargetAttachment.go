package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecretsmanagerSecretTargetAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_target_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "target_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SecretsManager::SecretTargetAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecretsmanagerSecretTargetAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerSecretTargetAttachment), &result)
	return &result
}
