package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradDirectoryRegistration = `{
  "block": {
    "attributes": {
      "directory_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_registration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::PCAConnectorAD::DirectoryRegistration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectoradDirectoryRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradDirectoryRegistration), &result)
	return &result
}
