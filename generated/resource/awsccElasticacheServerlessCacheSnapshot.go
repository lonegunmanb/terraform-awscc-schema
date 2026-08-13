package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheServerlessCacheSnapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the serverless cache snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "bytes_used_for_cache": {
        "computed": true,
        "description": "The total size of the serverless cache snapshot, in bytes.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time that the source serverless cache's metadata and cache data set was obtained for the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS KMS key used to encrypt the snapshot. Provide the key ARN: the resource returns the key ARN on read, so supplying a bare key ID or alias for this createOnly property may be reported as drift by CloudFormation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serverless_cache_configuration": {
        "computed": true,
        "description": "The configuration of the serverless cache, at the time the snapshot was taken.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "engine": {
              "computed": true,
              "description": "The engine that the serverless cache is configured with.",
              "description_kind": "plain",
              "type": "string"
            },
            "major_engine_version": {
              "computed": true,
              "description": "The engine version number that the serverless cache is configured with.",
              "description_kind": "plain",
              "type": "string"
            },
            "serverless_cache_name": {
              "computed": true,
              "description": "The identifier of the serverless cache.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "serverless_cache_name": {
        "description": "The name of an existing serverless cache. The snapshot is created from this cache.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "serverless_cache_snapshot_name": {
        "description": "The name of the serverless cache snapshot. Must be unique for the customer account. This value is stored as a lowercase string.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "snapshot_type": {
        "computed": true,
        "description": "The type of snapshot of the serverless cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the serverless cache snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to be added to the serverless cache snapshot resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for the tag. May not be null.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value. May be null.",
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
    "description": "Resource Type definition for AWS::ElastiCache::ServerlessCacheSnapshot. A serverless cache snapshot is a point-in-time backup of an ElastiCache serverless cache. Available for Valkey, Redis OSS and Serverless Memcached only.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticacheServerlessCacheSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheServerlessCacheSnapshot), &result)
	return &result
}
