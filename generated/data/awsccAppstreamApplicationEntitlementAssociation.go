package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamApplicationEntitlementAssociation = `{
  "block": {
    "attributes": {
      "application_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "entitlement_name": {
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
      "stack_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppStream::ApplicationEntitlementAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamApplicationEntitlementAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamApplicationEntitlementAssociation), &result)
	return &result
}
