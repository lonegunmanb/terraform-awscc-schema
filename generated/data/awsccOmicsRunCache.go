package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOmicsRunCache = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The run cache ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_behavior": {
        "computed": true,
        "description": "The default cache behavior for runs using this cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_bucket_owner_id": {
        "computed": true,
        "description": "The AWS account ID of the expected owner of the S3 bucket for the run cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_s3_location": {
        "computed": true,
        "description": "The S3 location for storing the cached task outputs.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation time of the run cache (an ISO 8601 formatted string).",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the run cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A name for the run cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "run_cache_id": {
        "computed": true,
        "description": "The run cache ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The run cache status.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the run cache.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Omics::RunCache",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOmicsRunCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsRunCache), &result)
	return &result
}
