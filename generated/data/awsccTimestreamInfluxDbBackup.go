package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTimestreamInfluxDbBackup = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The allocated storage of the resource at the time of backup, in GiB.",
        "description_kind": "plain",
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_configuration": {
        "computed": true,
        "description": "The cluster configuration of the resource at the time of backup.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dedicated_compactor": {
              "computed": true,
              "description": "Indicates if the compactor instance is a standalone instance or not.",
              "description_kind": "plain",
              "type": "bool"
            },
            "ingest_query_instances": {
              "computed": true,
              "description": "The number of ingest and query instances in the cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "query_only_instances": {
              "computed": true,
              "description": "The number of query-only instances in the cluster.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The time when the backup was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_type": {
        "computed": true,
        "description": "The DB instance type of the resource at the time of backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_parameter_group_id": {
        "computed": true,
        "description": "The identifier of the DB parameter group associated with the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_resource_id": {
        "computed": true,
        "description": "The id of the DB instance or DB cluster to back up.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_storage_type": {
        "computed": true,
        "description": "The storage type of the resource at the time of backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description": "The deployment type of the resource that the backup was created from.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_type": {
        "computed": true,
        "description": "The engine type of the resource that the backup was created from.",
        "description_kind": "plain",
        "type": "string"
      },
      "expires_after": {
        "computed": true,
        "description": "The date after which the backup will be automatically deleted, as a calendar date in YYYY-MM-DD format.",
        "description_kind": "plain",
        "type": "string"
      },
      "failover_mode": {
        "computed": true,
        "description": "The failover mode of the resource at the time of backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "influx_auth_parameters_secret_arn": {
        "computed": true,
        "description": "The ARN of the Secrets Manager secret containing the InfluxDB auth parameters.",
        "description_kind": "plain",
        "type": "string"
      },
      "influx_db_backup_id": {
        "computed": true,
        "description": "Service-generated unique identifier of the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Web Services KMS key ARN used for encryption of the resource at the time of backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_delivery_configuration": {
        "computed": true,
        "description": "The log delivery configuration of the resource at the time of backup.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_configuration": {
              "computed": true,
              "description": "Configuration for S3 bucket log delivery.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description": "The name of the S3 bucket to deliver logs to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "Indicates whether log delivery to the S3 bucket is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "maintenance_schedule": {
        "computed": true,
        "description": "The maintenance schedule of the resource at the time of backup.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "preferred_maintenance_window": {
              "computed": true,
              "description": "The preferred maintenance window in the format ddd:HH:MM-ddd:HH:MM.",
              "description_kind": "plain",
              "type": "string"
            },
            "timezone": {
              "computed": true,
              "description": "The IANA timezone identifier for the maintenance window.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the backup. Must be unique within the account and region.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The network type of the resource at the time of backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number of the resource at the time of backup.",
        "description_kind": "plain",
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Indicates whether the resource was publicly accessible at the time of backup.",
        "description_kind": "plain",
        "type": "bool"
      },
      "retention_days": {
        "computed": true,
        "description": "The number of days to retain the backup. Valid values are 1 to 3650.",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The current status of the backup.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to associate with the backup.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of backup. Backups created through this resource type are always ON_DEMAND.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "The VPC security group IDs associated with the resource at the time of backup.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_subnet_ids": {
        "computed": true,
        "description": "The VPC subnet IDs associated with the resource at the time of backup.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Timestream::InfluxDBBackup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTimestreamInfluxDbBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTimestreamInfluxDbBackup), &result)
	return &result
}
