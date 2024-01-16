package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSimspaceweaverSimulation = `{
  "block": {
    "attributes": {
      "describe_payload": {
        "computed": true,
        "description": "Json object with all simulation details",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "maximum_duration": {
        "computed": true,
        "description": "The maximum running time of the simulation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the simulation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "Role ARN.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema_s3_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "description": "The Schema S3 bucket name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "object_key": {
              "description": "This is the schema S3 object key, which includes the full path of \"folders\" from the bucket root to the schema.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "snapshot_s3_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "description": "The Schema S3 bucket name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "object_key": {
              "description": "This is the schema S3 object key, which includes the full path of \"folders\" from the bucket root to the schema.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "AWS::SimSpaceWeaver::Simulation resource creates an AWS Simulation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSimspaceweaverSimulationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSimspaceweaverSimulation), &result)
	return &result
}
