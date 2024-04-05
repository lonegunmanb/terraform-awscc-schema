package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverConfig = `{
  "block": {
    "attributes": {
      "autodefined_reverse": {
        "computed": true,
        "description": "ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.",
        "description_kind": "plain",
        "type": "string"
      },
      "autodefined_reverse_flag": {
        "computed": true,
        "description": "Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "AccountId",
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_config_id": {
        "computed": true,
        "description": "Id",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "ResourceId",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53Resolver::ResolverConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ResolverResolverConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverConfig), &result)
	return &result
}
