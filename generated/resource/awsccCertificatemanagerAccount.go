package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCertificatemanagerAccount = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiry_events_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "days_before_expiry": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::CertificateManager::Account.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCertificatemanagerAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCertificatemanagerAccount), &result)
	return &result
}
