package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferAgreement = `{
  "block": {
    "attributes": {
      "access_role": {
        "description": "Specifies the access role for the agreement.",
        "description_kind": "plain",
        "required": true,
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
        "description": "Specifies the base directory for the agreement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A textual description for the agreement.",
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
      "local_profile_id": {
        "description": "A unique identifier for the local profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partner_profile_id": {
        "description": "A unique identifier for the partner profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_id": {
        "description": "A unique identifier for the server.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Specifies the status of the agreement.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The name assigned to the tag that you create.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Transfer::Agreement",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTransferAgreementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferAgreement), &result)
	return &result
}
