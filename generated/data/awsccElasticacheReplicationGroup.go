package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticacheReplicationGroup = `{
  "block": {
    "attributes": {
      "at_rest_encryption_enabled": {
        "computed": true,
        "description": "A flag that enables encryption at rest when set to true.AtRestEncryptionEnabled after the replication group is created. To enable encryption at rest on a replication group you must set AtRestEncryptionEnabled to true when you create the replication group.",
        "description_kind": "plain",
        "type": "bool"
      },
      "auth_token": {
        "computed": true,
        "description": "Reserved parameter. The password used to access a password protected server.AuthToken can be specified only on replication groups where TransitEncryptionEnabled is true. For more information.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "This parameter is currently disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "automatic_failover_enabled": {
        "computed": true,
        "description": "Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails. AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled) replication groups.",
        "description_kind": "plain",
        "type": "bool"
      },
      "cache_node_type": {
        "computed": true,
        "description": "The compute and memory capacity of the nodes in the node group (shard).",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_parameter_group_name": {
        "computed": true,
        "description": "The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "cache_security_group_names": {
        "computed": true,
        "description": "A list of cache security group names to associate with this replication group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cache_subnet_group_name": {
        "computed": true,
        "description": "The name of the cache subnet group to be used for the replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_mode": {
        "computed": true,
        "description": "Enabled or Disabled. To modify cluster mode from Disabled to Enabled, you must first set the cluster mode to Compatible. Compatible mode allows your Redis OSS clients to connect using both cluster mode enabled and cluster mode disabled. After you migrate all Redis OSS clients to use cluster mode enabled, you can then complete cluster mode configuration and set the cluster mode to Enabled. For more information, see Modify cluster mode.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_end_point": {
        "computed": true,
        "description": "The configuration details of the replication group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The DNS hostname of the cache node.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port number that the cache engine is listening on.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_tiering_enabled": {
        "computed": true,
        "description": "Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes.",
        "description_kind": "plain",
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description": "The name of the cache engine to be used for the clusters in this replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the cache engine to be used for the clusters in this replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_replication_group_id": {
        "computed": true,
        "description": "The name of the Global datastore",
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
        "description": "The network type you choose when creating a replication group, either ipv4 | ipv6. IPv6 is supported for workloads using Redis OSS engine version 6.2 onward or Memcached engine version 1.6.6 on all instances built on the Nitro system.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ID of the KMS key used to encrypt the disk on the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_delivery_configurations": {
        "computed": true,
        "description": "Specifies the destination, format and type of the logs.",
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
                    "description": "The configuration details of the CloudWatch Logs destination. Note that this field is marked as required but only if CloudWatch Logs was chosen as the destination.",
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
                    "description": "The configuration details of the Kinesis Data Firehose destination. Note that this field is marked as required but only if Kinesis Data Firehose was chosen as the destination.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream": {
                          "computed": true,
                          "description": "The name of the Kinesis Data Firehose delivery stream.",
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
              "description": "Specify either CloudWatch Logs or Kinesis Data Firehose as the destination type. Valid values are either cloudwatch-logs or kinesis-firehose.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_format": {
              "computed": true,
              "description": "Valid values are either json or text.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_type": {
              "computed": true,
              "description": "Valid value is either slow-log, which refers to slow-log or engine-log.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "multi_az_enabled": {
        "computed": true,
        "description": "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see Minimizing Downtime: Multi-AZ.",
        "description_kind": "plain",
        "type": "bool"
      },
      "network_type": {
        "computed": true,
        "description": "Must be either ipv4 | ipv6 | dual_stack. IPv6 is supported for workloads using Redis OSS engine version 6.2 onward or Memcached engine version 1.6.6 on all instances built on the Nitro system",
        "description_kind": "plain",
        "type": "string"
      },
      "node_group_configuration": {
        "computed": true,
        "description": "NodeGroupConfiguration is a property of the AWS::ElastiCache::ReplicationGroup resource that configures an Amazon ElastiCache (ElastiCache) Redis cluster node group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "node_group_id": {
              "computed": true,
              "description": "Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.",
              "description_kind": "plain",
              "type": "string"
            },
            "primary_availability_zone": {
              "computed": true,
              "description": "The Availability Zone where the primary node of this node group (shard) is launched.",
              "description_kind": "plain",
              "type": "string"
            },
            "replica_availability_zones": {
              "computed": true,
              "description": "A list of Availability Zones to be used for the read replicas. The number of Availability Zones in this list must match the value of ReplicaCount or ReplicasPerNodeGroup if not specified.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "replica_count": {
              "computed": true,
              "description": "The number of read replica nodes in this node group (shard).",
              "description_kind": "plain",
              "type": "number"
            },
            "slots": {
              "computed": true,
              "description": "A string of comma-separated values where the first set of values are the slot numbers (zero based), and the second set of values are the keyspaces for each slot. The following example specifies three slots (numbered 0, 1, and 2): 0,1,2,0-4999,5000-9999,10000-16,383.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "notification_topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.",
        "description_kind": "plain",
        "type": "string"
      },
      "num_cache_clusters": {
        "computed": true,
        "description": "The number of clusters this replication group initially has.This parameter is not used if there is more than one node group (shard). You should use ReplicasPerNodeGroup instead.",
        "description_kind": "plain",
        "type": "number"
      },
      "num_node_groups": {
        "computed": true,
        "description": "An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group. For Redis (cluster mode disabled) either omit this parameter or set it to 1.",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number on which each member of the replication group accepts connections.",
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_cache_cluster_a_zs": {
        "computed": true,
        "description": "A list of EC2 Availability Zones in which the replication group's clusters are created. The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list. This parameter is not used if there is more than one node group (shard). You should use NodeGroupConfiguration instead.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_cluster_id": {
        "computed": true,
        "description": "The identifier of the cluster that serves as the primary for this replication group. This cluster must already exist and have a status of available.",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_end_point": {
        "computed": true,
        "description": "The primary endpoint configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The DNS hostname of the cache node.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port number that the cache engine is listening on.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "read_end_point": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "addresses": {
              "computed": true,
              "description": "A string containing a comma-separated list of endpoints for the primary and read-only replicas, formatted as [address1, address2, ...]. The order of the addresses maps to the order of the ports from the ReadEndPoint.Ports attribute.",
              "description_kind": "plain",
              "type": "string"
            },
            "addresses_list": {
              "computed": true,
              "description": "A list of endpoints for the read-only replicas. The order of the addresses maps to the order of the ports from the ReadEndPoint.Ports attribute.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "ports": {
              "computed": true,
              "description": "A string containing a comma-separated list of ports for the read-only replicas, formatted as [port1, port2, ...]. The order of the ports maps to the order of the addresses from the ReadEndPoint.Addresses attribute.",
              "description_kind": "plain",
              "type": "string"
            },
            "ports_list": {
              "computed": true,
              "description": "A list of ports for the read-only replicas. The order of the ports maps to the order of the addresses from the ReadEndPoint.Addresses attribute.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "reader_end_point": {
        "computed": true,
        "description": "The endpoint of the reader node in the replication group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The DNS hostname of the cache node.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port number that the cache engine is listening on.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "replicas_per_node_group": {
        "computed": true,
        "description": "An optional parameter that specifies the number of replica nodes in each node group (shard). Valid values are 0 to 5.",
        "description_kind": "plain",
        "type": "number"
      },
      "replication_group_description": {
        "computed": true,
        "description": "A user-created description for the replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_group_id": {
        "computed": true,
        "description": "The replication group identifier. This parameter is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "One or more Amazon VPC security groups associated with this replication group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_arns": {
        "computed": true,
        "description": "A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_name": {
        "computed": true,
        "description": "The name of a snapshot from which to restore data into the new replication group. The snapshot status changes to restoring while the new replication group is being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_retention_limit": {
        "computed": true,
        "description": "The number of days for which ElastiCache retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.",
        "description_kind": "plain",
        "type": "number"
      },
      "snapshot_window": {
        "computed": true,
        "description": "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshotting_cluster_id": {
        "computed": true,
        "description": "The cluster ID that is used as the daily snapshot source for the replication group. This parameter cannot be set for Redis (cluster mode enabled) replication groups.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of cost allocation tags to be added to this resource. Tags are comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "transit_encryption_enabled": {
        "computed": true,
        "description": "A flag that enables in-transit encryption when set to true.",
        "description_kind": "plain",
        "type": "bool"
      },
      "transit_encryption_mode": {
        "computed": true,
        "description": "A setting that allows you to migrate your clients to use in-transit encryption, with no downtime. When setting TransitEncryptionEnabled to true, you can set your TransitEncryptionMode to preferred in the same request, to allow both encrypted and unencrypted connections at the same time. Once you migrate all your Redis OSS clients to use encrypted connections you can modify the value to required to allow encrypted connections only. Setting TransitEncryptionMode to required is a two-step process that requires you to first set the TransitEncryptionMode to preferred, after that you can set TransitEncryptionMode to required. This process will not trigger the replacement of the replication group.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_group_ids": {
        "computed": true,
        "description": "The ID of user group to associate with the replication group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::ElastiCache::ReplicationGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticacheReplicationGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticacheReplicationGroup), &result)
	return &result
}
