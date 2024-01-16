package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradTemplateGroupAccessControlEntry = `{
  "block": {
    "attributes": {
      "access_rights": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_enroll": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enroll": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "group_display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "group_security_identifier": {
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
      "template_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::PCAConnectorAD::TemplateGroupAccessControlEntry",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectoradTemplateGroupAccessControlEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradTemplateGroupAccessControlEntry), &result)
	return &result
}
