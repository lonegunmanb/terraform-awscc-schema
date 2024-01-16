package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradTemplateGroupAccessControlEntry = `{
  "block": {
    "attributes": {
      "access_rights": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_enroll": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enroll": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "group_display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_security_identifier": {
        "computed": true,
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
      "template_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::PCAConnectorAD::TemplateGroupAccessControlEntry Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectoradTemplateGroupAccessControlEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradTemplateGroupAccessControlEntry), &result)
	return &result
}
