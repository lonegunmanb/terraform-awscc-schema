package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesDomainObjectType = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the domain object type was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the domain object type.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
        "description": "The default encryption key",
        "description_kind": "plain",
        "type": "string"
      },
      "fields": {
        "computed": true,
        "description": "A map of the name and ObjectType field.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content_type": {
              "computed": true,
              "description": "The content type of the field.",
              "description_kind": "plain",
              "type": "string"
            },
            "feature_type": {
              "computed": true,
              "description": "The feature type of the field.",
              "description_kind": "plain",
              "type": "string"
            },
            "source": {
              "computed": true,
              "description": "The source field name.",
              "description_kind": "plain",
              "type": "string"
            },
            "target": {
              "computed": true,
              "description": "The target field name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp of when the domain object type was most recently edited.",
        "description_kind": "plain",
        "type": "string"
      },
      "object_type_name": {
        "computed": true,
        "description": "The name of the domain object type.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::DomainObjectType",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesDomainObjectTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesDomainObjectType), &result)
	return &result
}
