package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamApplicationFleetAssociation = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_name": {
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
    "description": "Data Source schema for AWS::AppStream::ApplicationFleetAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamApplicationFleetAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamApplicationFleetAssociation), &result)
	return &result
}
