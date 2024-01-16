package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheServerlessCache = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_usage_limits": {
        "computed": true,
        "description": "The cache capacity limit of the Serverless Cache.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_storage": {
              "computed": true,
              "description": "The cached data capacity of the Serverless Cache.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum": {
                    "computed": true,
                    "description": "The maximum cached data capacity of the Serverless Cache.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "unit": {
                    "computed": true,
                    "description": "The unix of cached data capacity of the Serverless Cache.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ecpu_per_second": {
              "computed": true,
              "description": "The ECPU per second of the Serverless Cache.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum": {
                    "computed": true,
                    "description": "The maximum ECPU per second of the Serverless Cache.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "create_time": {
        "computed": true,
        "description": "The creation time of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "daily_snapshot_time": {
        "computed": true,
        "description": "The daily time range (in UTC) during which the service takes automatic snapshot of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The address and the port.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "Endpoint address.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "Endpoint port.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "engine": {
        "computed": true,
        "description": "The engine name of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "final_snapshot_name": {
        "computed": true,
        "description": "The final snapshot name which is taken before Serverless Cache is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "full_engine_version": {
        "computed": true,
        "description": "The full engine version of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ID of the KMS key used to encrypt the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "major_engine_version": {
        "computed": true,
        "description": "The major engine version of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "reader_endpoint": {
        "computed": true,
        "description": "The address and the port.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "Endpoint address.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "Endpoint port.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "security_group_ids": {
        "computed": true,
        "description": "One or more Amazon VPC security groups associated with this Serverless Cache.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "serverless_cache_name": {
        "computed": true,
        "description": "The name of the Serverless Cache. This value must be unique.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_arns_to_restore": {
        "computed": true,
        "description": "The ARN's of snapshot to restore Serverless Cache.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description": "The snapshot retention limit of the Serverless Cache.",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The status of the Serverless Cache.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The subnet id's of the Serverless Cache.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this Serverless Cache.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "user_group_id": {
        "computed": true,
        "description": "The ID of the user group.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ElastiCache::ServerlessCache",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticacheServerlessCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheServerlessCache), &result)
	return &result
}
