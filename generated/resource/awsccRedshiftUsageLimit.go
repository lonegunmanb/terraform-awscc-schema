package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftUsageLimit = `{
  "block": {
    "attributes": {
      "amount": {
        "description": "The limit amount. If time-based, this amount is in minutes. If data-based, this amount is in terabytes (TB). The value must be a positive number.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the usage limit.",
        "description_kind": "plain",
        "type": "string"
      },
      "breach_action": {
        "computed": true,
        "description": "The action that Amazon Redshift takes when the limit is reached. The default is log.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_identifier": {
        "description": "The identifier of the cluster that you want to limit usage.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "feature_type": {
        "description": "The Amazon Redshift feature that you want to limit.",
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
      "limit_type": {
        "description": "The type of limit. Depending on the feature type, this can be based on a time duration or data size. If FeatureType is spectrum, then LimitType must be data-scanned. If FeatureType is concurrency-scaling, then LimitType must be time. If FeatureType is cross-region-datasharing, then LimitType must be data-scanned.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "period": {
        "computed": true,
        "description": "The time period that the amount applies to. A weekly period begins on Sunday. The default is monthly.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tag instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key, or name, for the resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "usage_limit_id": {
        "computed": true,
        "description": "The identifier of the usage limit.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Redshift::UsageLimit. Creates a usage limit for a specified Amazon Redshift feature on a cluster. The usage limit is identified by the returned usage limit identifier.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftUsageLimitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftUsageLimit), &result)
	return &result
}
