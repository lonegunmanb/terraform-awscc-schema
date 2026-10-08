package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworksecuritymanagerTemplate = `{
  "block": {
    "attributes": {
      "associated_rule_list": {
        "computed": true,
        "description": "List of rules associated with this template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rule_arn": {
              "computed": true,
              "description": "ARN of the associated rule.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "firewall_type": {
        "computed": true,
        "description": "The type of firewall.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the template.",
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
        "description": "The ARN of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_description": {
        "computed": true,
        "description": "A description of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_id": {
        "computed": true,
        "description": "The unique identifier of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_name": {
        "computed": true,
        "description": "The name of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "An ISO 8601 timestamp indicating when the template was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkSecurityManager::Template",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworksecuritymanagerTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworksecuritymanagerTemplate), &result)
	return &result
}
