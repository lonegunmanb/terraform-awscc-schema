package resource

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
        "optional": true,
        "type": "string"
      },
      "archive_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::SES::MailManagerArchive Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesMailManagerArchiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerArchive), &result)
	return &result
}
