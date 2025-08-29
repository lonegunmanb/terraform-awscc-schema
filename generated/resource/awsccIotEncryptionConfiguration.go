package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotEncryptionConfiguration = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "error_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "error_message": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "encryption_type": {
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
      "kms_access_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IoT::EncryptionConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotEncryptionConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotEncryptionConfiguration), &result)
	return &result
}
