package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccHealthimagingDatastore = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the data store was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_arn": {
        "computed": true,
        "description": "The Datastore's ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_name": {
        "computed": true,
        "description": "User friendly name for Datastore.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_status": {
        "computed": true,
        "description": "A string to denote the Datastore's state.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "ARN referencing a KMS key or KMS key alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A Map of key value pairs for Tags.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the data store was created.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::HealthImaging::Datastore",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccHealthimagingDatastoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccHealthimagingDatastore), &result)
	return &result
}
