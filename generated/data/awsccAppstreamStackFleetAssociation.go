package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamStackFleetAssociation = `{
  "block": {
    "attributes": {
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
      },
      "stack_fleet_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppStream::StackFleetAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamStackFleetAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamStackFleetAssociation), &result)
	return &result
}
