package resource

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
        "description": "Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
        "description": "ResourceId",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Route53Resolver::ResolverConfig.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverResolverConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverConfig), &result)
	return &result
}
