package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEntityresolutionSchemaMapping = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time of this SchemaMapping got created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the SchemaMapping",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "has_workflows": {
        "computed": true,
        "description": "The boolean value that indicates whether or not a SchemaMapping has MatchingWorkflows that are associated with",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "mapped_input_fields": {
        "description": "The SchemaMapping attributes input",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "group_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hashed": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "match_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sub_type": {
              "computed": true,
              "description": "The subtype of the Attribute. Would be required only when type is PROVIDER_ID",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "schema_arn": {
        "computed": true,
        "description": "The SchemaMapping arn associated with the Schema",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_name": {
        "description": "The name of the SchemaMapping",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time of this SchemaMapping got last updated at",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "SchemaMapping defined in AWS Entity Resolution service",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEntityresolutionSchemaMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEntityresolutionSchemaMapping), &result)
	return &result
}
