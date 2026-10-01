package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncApiCache = `{
  "block": {
    "attributes": {
      "api_cache_id": {
        "computed": true,
        "description": "The CloudFormation primary identifier for an ApiCache. Mirrors ApiId once the resource exists; populated by the handler. Customers should reference the ApiCache by ApiId.",
        "description_kind": "plain",
        "type": "string"
      },
      "api_caching_behavior": {
        "description": "Caching behavior, such as FULL_REQUEST_CACHING, PER_RESOLVER_CACHING, or OPERATION_LEVEL_CACHING",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "api_id": {
        "description": "Unique AWS AppSync GraphQL API identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "at_rest_encryption_enabled": {
        "computed": true,
        "description": "At-rest encryption flag for cache. You cannot update this setting after creation.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_metrics_config": {
        "computed": true,
        "description": "Controls how cache health metrics will be emitted to CloudWatch. Cache health metrics include NetworkBandwidthOutAllowanceExceeded and EngineCPUUtilization",
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
      "transit_encryption_enabled": {
        "computed": true,
        "description": "Transit encryption flag when connecting to cache. You cannot update this setting after creation.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ttl": {
        "description": "TTL in seconds for cache entries. Valid values are 1-3600.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "type": {
        "description": "The cache instance type, such as SMALL, LARGE and LARGE_4X",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppSync::ApiCache",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncApiCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncApiCache), &result)
	return &result
}
