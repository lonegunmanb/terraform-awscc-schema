package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupLegalHold = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the legal hold.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date": {
        "computed": true,
        "description": "The time when the legal hold was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The description of the legal hold.",
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
      "legal_hold_id": {
        "computed": true,
        "description": "The ID of the legal hold.",
        "description_kind": "plain",
        "type": "string"
      },
      "recovery_point_selection": {
        "description": "The criteria to assign a set of resources, such as resource types or backup vaults.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "date_range": {
              "computed": true,
              "description": "A date range for filtering recovery points.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "from_date": {
                    "computed": true,
                    "description": "The beginning date, inclusive. ISO 8601 date-time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "to_date": {
                    "computed": true,
                    "description": "The end date, inclusive. ISO 8601 date-time.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "resource_identifiers": {
              "computed": true,
              "description": "The resources included in the resource selection.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vault_names": {
              "computed": true,
              "description": "The names of the vaults in which the selected recovery points are contained.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "The status of the legal hold.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Optional tags to include.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "title": {
        "description": "The title of the legal hold.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Creates a legal hold on recovery points (backups). A legal hold prevents backups from being deleted while under hold.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupLegalHoldSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupLegalHold), &result)
	return &result
}
