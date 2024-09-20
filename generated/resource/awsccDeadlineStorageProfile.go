package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineStorageProfile = `{
  "block": {
    "attributes": {
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "farm_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "file_system_locations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "os_family": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Deadline::StorageProfile Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineStorageProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineStorageProfile), &result)
	return &result
}
