package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMgnWave = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the wave.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_date_time": {
        "computed": true,
        "description": "The timestamp when the wave was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the wave.",
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
        "description": "Whether the wave is archived.",
        "description_kind": "plain",
        "type": "bool"
      },
      "last_modified_date_time": {
        "computed": true,
        "description": "The timestamp when the wave was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the wave.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the wave.",
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
      },
      "wave_aggregated_status": {
        "computed": true,
        "description": "The aggregated migration status of the wave.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "health_status": {
              "computed": true,
              "description": "The health status of the wave.",
              "description_kind": "plain",
              "type": "string"
            },
            "progress_status": {
              "computed": true,
              "description": "The progress status of the wave.",
              "description_kind": "plain",
              "type": "string"
            },
            "total_applications": {
              "computed": true,
              "description": "The total number of applications in the wave.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "wave_id": {
        "computed": true,
        "description": "The unique identifier of the wave.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::MGN::Wave. A Wave is a logical grouping used in Application Migration Service to organize and manage the migration of applications.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMgnWaveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMgnWave), &result)
	return &result
}
