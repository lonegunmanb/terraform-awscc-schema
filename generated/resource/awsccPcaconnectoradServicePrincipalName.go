package resource

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
        "optional": true,
        "type": "string"
      },
      "directory_registration_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::PCAConnectorAD::ServicePrincipalName Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectoradServicePrincipalNameSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradServicePrincipalName), &result)
	return &result
}
