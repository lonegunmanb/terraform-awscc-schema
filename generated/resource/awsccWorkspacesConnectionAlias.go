package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesConnectionAlias = `{
  "block": {
    "attributes": {
      "alias_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "associated_account_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "association_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "connection_identifier": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "resource_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "connection_alias_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_string": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::WorkSpaces::ConnectionAlias",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspacesConnectionAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesConnectionAlias), &result)
	return &result
}
