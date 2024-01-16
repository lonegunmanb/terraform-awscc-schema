package data

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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
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
    "description": "Data Source schema for AWS::AppSync::DomainNameApiAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppsyncDomainNameApiAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncDomainNameApiAssociation), &result)
	return &result
}
