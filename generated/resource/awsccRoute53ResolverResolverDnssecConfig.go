package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverDnssecConfig = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Id",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "AccountId",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "ResourceId",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "validation_status": {
        "computed": true,
        "description": "ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverResolverDnssecConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverDnssecConfig), &result)
	return &result
}
