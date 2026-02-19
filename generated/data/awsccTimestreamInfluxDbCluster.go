package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTimestreamInfluxDbCluster = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The allocated storage for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is associated with the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "computed": true,
        "description": "The bucket for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_type": {
        "computed": true,
        "description": "The compute instance of the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_parameter_group_identifier": {
        "computed": true,
        "description": "The name of an existing InfluxDB parameter group.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_storage_type": {
        "computed": true,
        "description": "The storage type of the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description": "Deployment type of the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The connection endpoint for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_type": {
        "computed": true,
        "description": "The engine type for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "failover_mode": {
        "computed": true,
        "description": "Failover mode of the InfluxDB cluster.",
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
        "description": "The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "influx_db_cluster_id": {
        "computed": true,
        "description": "The service generated unique identifier for InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_delivery_configuration": {
        "computed": true,
        "description": "Configuration for sending logs to customer account from the InfluxDB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_configuration": {
              "computed": true,
              "description": "S3 configuration for sending logs to customer account from the InfluxDB cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description": "The bucket name for logs to be sent from the InfluxDB cluster",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description": "Specifies whether logging to customer specified bucket is enabled.",
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
      "name": {
        "computed": true,
        "description": "The unique name that is associated with the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "Network type of the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "organization": {
        "computed": true,
        "description": "The organization for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "The password for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number on which InfluxDB accepts connections.",
        "description_kind": "plain",
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Attach a public IP to the customer ENI.",
        "description_kind": "plain",
        "type": "bool"
      },
      "reader_endpoint": {
        "computed": true,
        "description": "The reader endpoint for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this DB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "username": {
        "computed": true,
        "description": "The username for the InfluxDB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "A list of Amazon EC2 VPC security groups to associate with this InfluxDB cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpc_subnet_ids": {
        "computed": true,
        "description": "A list of EC2 subnet IDs for this InfluxDB cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Timestream::InfluxDBCluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTimestreamInfluxDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTimestreamInfluxDbCluster), &result)
	return &result
}
