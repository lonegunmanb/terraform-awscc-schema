package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectInstanceStorageConfig = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "An associationID is automatically generated when a storage config is associated with an instance",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "Connect Instance ID with which the storage config will be associated",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kinesis_firehose_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "firehose_arn": {
              "computed": true,
              "description": "An ARN is a unique AWS resource identifier.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "kinesis_stream_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "stream_arn": {
              "computed": true,
              "description": "An ARN is a unique AWS resource identifier.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "kinesis_video_stream_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encryption_type": {
                    "computed": true,
                    "description": "Specifies default encryption using AWS KMS-Managed Keys",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_id": {
                    "computed": true,
                    "description": "Specifies the encryption key id",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "prefix": {
              "computed": true,
              "description": "Prefixes are used to infer logical hierarchy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "retention_period_hours": {
              "computed": true,
              "description": "Number of hours",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "resource_type": {
        "description": "Specifies the type of storage resource available for the instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description": "A name for the S3 Bucket",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_prefix": {
              "computed": true,
              "description": "Prefixes are used to infer logical hierarchy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encryption_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encryption_type": {
                    "computed": true,
                    "description": "Specifies default encryption using AWS KMS-Managed Keys",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_id": {
                    "computed": true,
                    "description": "Specifies the encryption key id",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "storage_type": {
        "description": "Specifies the storage type to be associated with the instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::InstanceStorageConfig",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectInstanceStorageConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectInstanceStorageConfig), &result)
	return &result
}
