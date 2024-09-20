package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRolesanywhereProfile = `{
  "block": {
    "attributes": {
      "accept_role_session_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "attribute_mappings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_field": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mapping_rules": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "specifier": {
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "duration_seconds": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "managed_policy_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "require_instance_properties": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "role_arns": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "session_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::RolesAnywhere::Profile Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRolesanywhereProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRolesanywhereProfile), &result)
	return &result
}
