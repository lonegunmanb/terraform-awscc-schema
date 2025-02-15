package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferAgreement = `{
  "block": {
    "attributes": {
      "access_role": {
        "computed": true,
        "description": "Specifies the access role for the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "agreement_id": {
        "computed": true,
        "description": "A unique identifier for the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "base_directory": {
        "computed": true,
        "description": "Specifies the base directory for the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A textual description for the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "enforce_message_signing": {
        "computed": true,
        "description": "Specifies whether to enforce an AS2 message is signed for this agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "local_profile_id": {
        "computed": true,
        "description": "A unique identifier for the local profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "partner_profile_id": {
        "computed": true,
        "description": "A unique identifier for the partner profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "preserve_filename": {
        "computed": true,
        "description": "Specifies whether to preserve the filename received for this agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_id": {
        "computed": true,
        "description": "A unique identifier for the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Specifies the status of the agreement.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name assigned to the tag that you create.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Transfer::Agreement",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTransferAgreementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferAgreement), &result)
	return &result
}
