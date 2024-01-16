package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradServicePrincipalName = `{
  "block": {
    "attributes": {
      "connector_arn": {
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
      }
    },
    "description": "Data Source schema for AWS::PCAConnectorAD::ServicePrincipalName",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectoradServicePrincipalNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradServicePrincipalName), &result)
	return &result
}
