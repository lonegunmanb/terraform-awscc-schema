package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsAdConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Ad configuration ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "media_tailor_playback_configurations": {
        "computed": true,
        "description": "List of integration configurations with MediaTailor resources. The first item in the list is the default playback configuration used for the ad configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "playback_configuration_arn": {
              "computed": true,
              "description": "ARN of the customer-created EMT PlaybackConfiguration resource in the same region and account.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description": "Ad configuration name. The value does not need to be unique.",
        "description_kind": "plain",
        "type": "string"
      },
      "post_roll_configuration": {
        "computed": true,
        "description": "Configuration for the post-roll ad break to use for this ad configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "duration_seconds": {
              "computed": true,
              "description": "Duration of the post-roll ad break, in seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "enabled": {
              "computed": true,
              "description": "Whether the post-roll ad configuration is enabled.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags attached to the resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::IVS::AdConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIvsAdConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsAdConfiguration), &result)
	return &result
}
