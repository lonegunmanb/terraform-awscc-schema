package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverResolverQueryLoggingConfig = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "association_count": {
        "computed": true,
        "description": "Count",
        "description_kind": "plain",
        "type": "number"
      },
      "creation_time": {
        "computed": true,
        "description": "Rfc3339TimeString",
        "description_kind": "plain",
        "type": "string"
      },
      "creator_request_id": {
        "computed": true,
        "description": "The id of the creator request.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_arn": {
        "computed": true,
        "description": "destination arn",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "ResolverQueryLogConfigName",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "AccountId",
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_query_logging_config_id": {
        "computed": true,
        "description": "ResourceId",
        "description_kind": "plain",
        "type": "string"
      },
      "share_status": {
        "computed": true,
        "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverResolverQueryLoggingConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverResolverQueryLoggingConfig), &result)
	return &result
}
