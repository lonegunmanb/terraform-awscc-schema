package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecuritylakeDataLake = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) created by you to provide to the subscriber.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "Provides encryption details of Amazon Security Lake object.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The id of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_configuration": {
        "computed": true,
        "description": "Provides lifecycle details of Amazon Security Lake object.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expiration": {
              "computed": true,
              "description": "Provides data expiration details of Amazon Security Lake object.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "days": {
                    "computed": true,
                    "description": "Number of days before data expires in the Amazon Security Lake object.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "transitions": {
              "computed": true,
              "description": "Provides data storage transition details of Amazon Security Lake object.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "days": {
                    "computed": true,
                    "description": "Number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage_class": {
                    "computed": true,
                    "description": "The range of storage classes that you can choose from based on the data access, resiliency, and cost requirements of your workloads.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "meta_store_manager_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) used to index AWS Glue table partitions that are generated by the ingestion and normalization of AWS log sources and custom sources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_configuration": {
        "computed": true,
        "description": "Provides replication details of Amazon Security Lake object.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "regions": {
              "computed": true,
              "description": "Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. Amazon S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different AWS Regions or within the same Region as the source bucket.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "role_arn": {
              "computed": true,
              "description": "Replication settings for the Amazon S3 buckets. This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake, to ensure the replication setting is correct.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "s3_bucket_arn": {
        "computed": true,
        "description": "The ARN for the Amazon Security Lake Amazon S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, ` + "`" + `_` + "`" + `, ` + "`" + `.` + "`" + `, ` + "`" + `/` + "`" + `, ` + "`" + `=` + "`" + `, ` + "`" + `+` + "`" + `, and ` + "`" + `-` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SecurityLake::DataLake",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecuritylakeDataLakeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecuritylakeDataLake), &result)
	return &result
}
