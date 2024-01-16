package resource

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
        "optional": true,
        "type": "number"
      },
      "contact_pre_pass_duration_seconds": {
        "computed": true,
        "description": "Pre-pass time needed before the contact.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dataflow_edges": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "minimum_viable_contact_duration_seconds": {
        "description": "Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "name": {
        "description": "A name used to identify a mission profile.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "streams_kms_role": {
        "computed": true,
        "description": "The ARN of the KMS Key or Alias Key role used to define permissions on KMS Key usage.",
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
      },
      "tracking_config_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "AWS Ground Station Mission Profile resource type for CloudFormation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGroundstationMissionProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGroundstationMissionProfile), &result)
	return &result
}
