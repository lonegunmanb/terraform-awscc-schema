package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maximum_duration": {
        "computed": true,
        "description": "The maximum running time of the simulation.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the simulation.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Role ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema_s3_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description": "The Schema S3 bucket name.",
              "description_kind": "plain",
              "type": "string"
            },
            "object_key": {
              "computed": true,
              "description": "This is the schema S3 object key, which includes the full path of \"folders\" from the bucket root to the schema.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snapshot_s3_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description": "The Schema S3 bucket name.",
              "description_kind": "plain",
              "type": "string"
            },
            "object_key": {
              "computed": true,
              "description": "This is the schema S3 object key, which includes the full path of \"folders\" from the bucket root to the schema.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SimSpaceWeaver::Simulation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSimspaceweaverSimulationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSimspaceweaverSimulation), &result)
	return &result
}
