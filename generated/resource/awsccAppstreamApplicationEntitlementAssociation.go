package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamApplicationEntitlementAssociation = `{
  "block": {
    "attributes": {
      "application_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entitlement_name": {
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
      "stack_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamApplicationEntitlementAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamApplicationEntitlementAssociation), &result)
	return &result
}
