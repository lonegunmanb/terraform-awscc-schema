package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEncryptionControls = `{
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
    "description": "Plural Data Source schema for AWS::EC2::VPCEncryptionControl",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcEncryptionControlsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEncryptionControls), &result)
	return &result
}
