package data

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
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "target_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "provisioning_role_arn": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "template_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_body": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "template_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::ProvisioningTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotProvisioningTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotProvisioningTemplate), &result)
	return &result
}
