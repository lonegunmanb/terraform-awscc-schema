package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisanalyticsv2Snapshot = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "The name of the Managed Service for Apache Flink application to snapshot. Used only at create time to select the source application; immutable after create. Specify either ApplicationName or ImportSourceConfiguration, not both.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "import_source_configuration": {
        "computed": true,
        "description": "The Amazon S3 location and role to import the snapshot from, with ImportSnapshot. Used only at create time; immutable after create. Specify either ApplicationName or ImportSourceConfiguration, not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "import_role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that the service assumes to read the snapshot files from Amazon S3 (and to decrypt them, if they are encrypted with a customer managed key).",
              "description_kind": "plain",
              "type": "string"
            },
            "source_encryption_configuration": {
              "computed": true,
              "description": "The encryption of the snapshot files in Amazon S3.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key_id": {
                    "computed": true,
                    "description": "The ARN of the customer managed KMS key. Required when KeyType is CUSTOMER_MANAGED_KEY; must be omitted for AWS_OWNED_KEY.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description": "The type of KMS key: AWS_OWNED_KEY or CUSTOMER_MANAGED_KEY.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_s3_prefix_location": {
              "computed": true,
              "description": "The Amazon S3 bucket and key prefix that hold the snapshot files, for example the prefix an ExportSnapshot operation wrote them to.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_prefix": {
                    "computed": true,
                    "description": "The S3 key prefix.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_type": {
              "computed": true,
              "description": "The type of storage the snapshot is imported from.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "runtime_environment": {
        "computed": true,
        "description": "The runtime environment of the source application when the snapshot was taken with ApplicationName (CreateSnapshot), for example FLINK-1_20. Populated by DescribeSnapshot. Not returned for snapshots created with ImportSourceConfiguration (ImportSnapshot), which carry no runtime version.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the snapshot. Account-scoped: arn:\u003cpartition\u003e:kinesisanalytics:\u003cregion\u003e:\u003caccount\u003e:snapshot/\u003cSnapshotName\u003e. Assigned by the service at create time and used as the resource's primary identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_creation_timestamp": {
        "computed": true,
        "description": "The time at which the service created the snapshot, in ISO-8601 format. Populated by DescribeSnapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_description": {
        "computed": true,
        "description": "Optional customer-supplied description of the snapshot. Immutable after create.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_encryption_configuration_description": {
        "computed": true,
        "description": "The encryption of the snapshot. Can be set only with ImportSourceConfiguration (sent as the ImportSnapshot DestinationEncryptionConfiguration; the AWS owned key when omitted). A snapshot taken from ApplicationName uses the application's encryption: DescribeSnapshot returns it for such snapshots, but it must not be set with ApplicationName. Immutable after create.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_id": {
              "computed": true,
              "description": "The ARN of the customer managed KMS key. Required when KeyType is CUSTOMER_MANAGED_KEY; must be omitted for AWS_OWNED_KEY.",
              "description_kind": "plain",
              "type": "string"
            },
            "key_type": {
              "computed": true,
              "description": "The type of KMS key: AWS_OWNED_KEY or CUSTOMER_MANAGED_KEY.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snapshot_format": {
        "computed": true,
        "description": "The format of the snapshot (CANONICAL). Sent to ImportSnapshot with ImportSourceConfiguration; with ApplicationName it may only be CANONICAL, the format CreateSnapshot produces. Returned by DescribeSnapshot. Immutable after create.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_initiator": {
        "computed": true,
        "description": "Indicates whether the snapshot was initiated by the customer or by the service. Populated by DescribeSnapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_initiator_operation": {
        "computed": true,
        "description": "The operation that initiated the snapshot. Populated by DescribeSnapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_name": {
        "computed": true,
        "description": "The customer-chosen name of the snapshot, unique within the account. Immutable after create.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_size_in_bytes": {
        "computed": true,
        "description": "The size of the snapshot in bytes. Populated by DescribeSnapshot once the snapshot reaches READY.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_source_description": {
        "computed": true,
        "description": "The source of the snapshot: the application it was taken from, or the Amazon S3 location it was imported from. Populated by DescribeSnapshot.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_prefix_location": {
              "computed": true,
              "description": "An Amazon S3 prefix location.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_prefix": {
                    "computed": true,
                    "description": "The S3 key prefix.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_application_description": {
              "computed": true,
              "description": "Details about the application the snapshot was taken from.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "application_creation_timestamp": {
                    "computed": true,
                    "description": "The time at which the source application was created.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "application_name": {
                    "computed": true,
                    "description": "The name of the source application.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "application_version_id": {
                    "computed": true,
                    "description": "The version of the source application captured in the snapshot.",
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
      "snapshot_status": {
        "computed": true,
        "description": "The lifecycle status of the snapshot as reported by the service (CREATING, READY, DELETING, FAILED). Populated by DescribeSnapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the snapshot. Applied at create time (with CreateSnapshot, or with TagResource right after ImportSnapshot) and updated in place with TagResource and UntagResource. Tag keys must not start with aws:.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::KinesisAnalyticsV2::Snapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKinesisanalyticsv2SnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisanalyticsv2Snapshot), &result)
	return &result
}
