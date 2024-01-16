package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverQueryLoggingConfigAssociation = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "error": {
        "computed": true,
        "description": "ResolverQueryLogConfigAssociationError",
        "description_kind": "plain",
        "type": "string"
      },
      "error_message": {
        "computed": true,
        "description": "ResolverQueryLogConfigAssociationErrorMessage",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resolver_query_log_config_id": {
        "computed": true,
        "description": "ResolverQueryLogConfigId",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "ResourceId",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "ResolverQueryLogConfigAssociationStatus",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53ResolverResolverQueryLoggingConfigAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverQueryLoggingConfigAssociation), &result)
	return &result
}
