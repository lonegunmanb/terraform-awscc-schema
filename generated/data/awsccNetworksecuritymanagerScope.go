package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworksecuritymanagerScope = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_configuration": {
        "computed": true,
        "description": "The scope configuration as a JSON string.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_description": {
        "computed": true,
        "description": "A description of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_id": {
        "computed": true,
        "description": "The unique identifier of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope_name": {
        "computed": true,
        "description": "The name of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the scope.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the scope.",
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
      "updated_at": {
        "computed": true,
        "description": "An ISO 8601 timestamp indicating when the scope was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version number of the scope.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkSecurityManager::Scope",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworksecuritymanagerScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworksecuritymanagerScope), &result)
	return &result
}
