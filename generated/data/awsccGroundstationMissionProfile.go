package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGroundstationMissionProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "contact_post_pass_duration_seconds": {
        "computed": true,
        "description": "Post-pass time needed after the contact.",
        "description_kind": "plain",
        "type": "number"
      },
      "contact_pre_pass_duration_seconds": {
        "computed": true,
        "description": "Pre-pass time needed before the contact.",
        "description_kind": "plain",
        "type": "number"
      },
      "dataflow_edges": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "minimum_viable_contact_duration_seconds": {
        "computed": true,
        "description": "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
        "description_kind": "plain",
        "type": "number"
      },
      "mission_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A name used to identify a mission profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "streams_kms_key": {
        "computed": true,
        "description": "The ARN of a KMS Key used for encrypting data during transmission from the source to destination locations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_alias_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_alias_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "streams_kms_role": {
        "computed": true,
        "description": "The ARN of the KMS Key or Alias Key role used to define permissions on KMS Key usage.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tracking_config_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::GroundStation::MissionProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGroundstationMissionProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGroundstationMissionProfile), &result)
	return &result
}
