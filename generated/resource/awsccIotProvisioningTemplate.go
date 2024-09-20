package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotProvisioningTemplate = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "pre_provisioning_hook": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "payload_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_arn": {
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
      "provisioning_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "template_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_body": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Creates a fleet provisioning template.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotProvisioningTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotProvisioningTemplate), &result)
	return &result
}
