package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamStackFleetAssociation = `{
  "block": {
    "attributes": {
      "fleet_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "stack_fleet_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppStream::StackFleetAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamStackFleetAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamStackFleetAssociation), &result)
	return &result
}
