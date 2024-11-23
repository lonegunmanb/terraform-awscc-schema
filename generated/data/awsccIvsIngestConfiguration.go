package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsIngestConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "IngestConfiguration ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ingest_protocol": {
        "computed": true,
        "description": "Ingest Protocol.",
        "description_kind": "plain",
        "type": "string"
      },
      "insecure_ingest": {
        "computed": true,
        "description": "Whether ingest configuration allows insecure ingest.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "IngestConfiguration",
        "description_kind": "plain",
        "type": "string"
      },
      "participant_id": {
        "computed": true,
        "description": "Participant Id is automatically generated on creation and assigned.",
        "description_kind": "plain",
        "type": "string"
      },
      "stage_arn": {
        "computed": true,
        "description": "Stage ARN. A value other than an empty string indicates that stage is linked to IngestConfiguration. Default: \"\" (recording is disabled).",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of IngestConfiguration which determines whether IngestConfiguration is in use or not.",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_key": {
        "computed": true,
        "description": "Stream-key value.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset model.",
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
          "nesting_mode": "set"
        }
      },
      "user_id": {
        "computed": true,
        "description": "User defined indentifier for participant associated with IngestConfiguration.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IVS::IngestConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIvsIngestConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsIngestConfiguration), &result)
	return &result
}
