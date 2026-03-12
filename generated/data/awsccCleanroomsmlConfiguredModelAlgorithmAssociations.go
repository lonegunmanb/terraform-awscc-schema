package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsmlConfiguredModelAlgorithmAssociations = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description": "Set of Resource Identifiers.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Plural Data Source schema for AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCleanroomsmlConfiguredModelAlgorithmAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsmlConfiguredModelAlgorithmAssociations), &result)
	return &result
}
