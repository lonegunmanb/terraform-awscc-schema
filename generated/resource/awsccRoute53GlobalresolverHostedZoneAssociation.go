package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53GlobalresolverHostedZoneAssociation = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hosted_zone_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Route53GlobalResolver::HostedZoneAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53GlobalresolverHostedZoneAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53GlobalresolverHostedZoneAssociation), &result)
	return &result
}
