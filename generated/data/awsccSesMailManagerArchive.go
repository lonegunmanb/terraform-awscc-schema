package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerArchive = `{
  "block": {
    "attributes": {
      "archive_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "archive_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "archive_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "archive_state": {
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
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "retention": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "retention_period": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SES::MailManagerArchive",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesMailManagerArchiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerArchive), &result)
	return &result
}
