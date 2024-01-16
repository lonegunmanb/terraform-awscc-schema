package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncDomainNameApiAssociation = `{
  "block": {
    "attributes": {
      "api_association_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_name": {
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
    "description": "Resource Type definition for AWS::AppSync::DomainNameApiAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncDomainNameApiAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncDomainNameApiAssociation), &result)
	return &result
}
