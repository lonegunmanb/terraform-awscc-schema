package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamApplicationFleetAssociation = `{
  "block": {
    "attributes": {
      "application_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      }
    },
    "description": "Resource Type definition for AWS::AppStream::ApplicationFleetAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamApplicationFleetAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamApplicationFleetAssociation), &result)
	return &result
}
