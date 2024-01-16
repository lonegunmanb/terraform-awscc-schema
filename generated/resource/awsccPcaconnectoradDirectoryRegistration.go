package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradDirectoryRegistration = `{
  "block": {
    "attributes": {
      "directory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "directory_registration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Definition of AWS::PCAConnectorAD::DirectoryRegistration Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectoradDirectoryRegistrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradDirectoryRegistration), &result)
	return &result
}
