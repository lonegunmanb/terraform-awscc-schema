package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheCacheCluster = `{
  "block": {
    "attributes": {
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "If you are running Redis engine version 6.0 or later, set this parameter to yes if you want to opt-in to the next minor version upgrade campaign.",
        "description_kind": "plain",
        "type": "bool"
      },
      "az_mode": {
        "computed": true,
        "description": "Specifies whether the nodes in this Memcached cluster are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_node_type": {
        "computed": true,
        "description": "The compute and memory capacity of the nodes in the node group (shard).",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_parameter_group_name": {
        "computed": true,
        "description": "The name of the parameter group to associate with this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_security_group_names": {
        "computed": true,
        "description": "A list of security group names to associate with this cluster.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "cache_subnet_group_name": {
        "computed": true,
        "description": "The name of the subnet group to be used for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "A name for the cache cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_endpoint": {
        "computed": true,
        "description": "Specifies the ConfigurationEndpoint address and port",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "engine": {
        "computed": true,
        "description": "The name of the cache engine to be used for this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the cache engine to be used for this cluster",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_discovery": {
        "computed": true,
        "description": "The Ip Discovery parameter for cachecluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_delivery_configurations": {
        "computed": true,
        "description": "Specifies the destination, format and type of the logs",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_details": {
              "computed": true,
              "description": "Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_logs_details": {
                    "computed": true,
                    "description": "The configuration details of the CloudWatch Logs destination",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_group": {
                          "computed": true,
                          "description": "The name of the CloudWatch Logs log group.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "kinesis_firehose_details": {
                    "computed": true,
                    "description": "The configuration details of the Kinesis Data Firehose destination.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream": {
                          "computed": true,
                          "description": "The name of the Kinesis Data Firehose delivery stream",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "destination_type": {
              "computed": true,
              "description": "Specify either CloudWatch Logs or Kinesis Data Firehose as the destination type. ",
              "description_kind": "plain",
              "type": "string"
            },
            "log_format": {
              "computed": true,
              "description": "Valid values are either json or text",
              "description_kind": "plain",
              "type": "string"
            },
            "log_type": {
              "computed": true,
              "description": "Valid value is either slow-log, which refers to slow-log or engine-log",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "network_type": {
        "computed": true,
        "description": "The network type parameter for cachecluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
        "description_kind": "plain",
        "type": "string"
      },
      "num_cache_nodes": {
        "computed": true,
        "description": "The number of cache nodes that the cache cluster should have.",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number on which each of the cache nodes accepts connections.",
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_availability_zone": {
        "computed": true,
        "description": "The EC2 Availability Zone in which the cluster is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_availability_zones": {
        "computed": true,
        "description": "A list of the Availability Zones in which cache nodes are created. The order of the zones in the list is not important.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "Specifies the weekly time range during which maintenance on the cluster is performed.",
        "description_kind": "plain",
        "type": "string"
      },
      "redis_endpoint": {
        "computed": true,
        "description": "Specifies the RedisEndPoint address and port",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snapshot_arns": {
        "computed": true,
        "description": "A single-element string list containing an Amazon Resource Name (ARN) that uniquely identifies a Redis RDB snapshot file stored in Amazon S3.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "snapshot_name": {
        "computed": true,
        "description": "The name of a Redis snapshot from which to restore data into the new node group (shard).",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description": "The number of days for which ElastiCache retains automatic snapshots before deleting them.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_window": {
        "computed": true,
        "description": "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to be added to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "transit_encryption_enabled": {
        "computed": true,
        "description": "A flag that enables in-transit encryption when set to true. You cannot modify the value of TransitEncryptionEnabled after the cluster is created",
        "description_kind": "plain",
        "type": "bool"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "One or more VPC security groups associated with the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::ElastiCache::CacheCluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticacheCacheClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheCacheCluster), &result)
	return &result
}
