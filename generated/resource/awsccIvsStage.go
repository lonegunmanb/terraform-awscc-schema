package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsStage = `{
  "block": {
    "attributes": {
      "active_session_id": {
        "computed": true,
        "description": "ID of the active session within the stage.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Stage ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_participant_recording_configuration": {
        "computed": true,
        "description": "Configuration object for individual participant recording, to attach to the new stage.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "media_types": {
              "computed": true,
              "description": "Types of media to be recorded. Default: AUDIO_VIDEO.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "storage_configuration_arn": {
              "description": "ARN of the StorageConfiguration resource to use for individual participant recording.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Stage name",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Definition for type AWS::IVS::Stage.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvsStageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsStage), &result)
	return &result
}
