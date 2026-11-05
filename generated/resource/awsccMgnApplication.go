package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMgnApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The ID of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date_time": {
        "computed": true,
        "description": "The date and time the application was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the application.",
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
      "is_archived": {
        "computed": true,
        "description": "Whether the application is archived.",
        "description_kind": "plain",
        "type": "bool"
      },
      "last_modified_date_time": {
        "computed": true,
        "description": "The date and time the application was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags to associate with the application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "An MGN Application is a logical grouping of source servers being migrated together using AWS Application Migration Service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMgnApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnApplication), &result)
	return &result
}
