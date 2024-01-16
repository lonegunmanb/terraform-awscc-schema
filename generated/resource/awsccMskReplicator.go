package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskReplicator = `{
  "block": {
    "attributes": {
      "current_version": {
        "computed": true,
        "description": "The current version of the MSK replicator.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A summary description of the replicator.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kafka_clusters": {
        "description": "Specifies a list of Kafka clusters which are targets of the replicator.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_msk_cluster": {
              "description": "Details of an Amazon MSK cluster. Exactly one of AmazonMskCluster is required.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "msk_cluster_arn": {
                    "description": "The ARN of an Amazon MSK cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "vpc_config": {
              "description": "Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The AWS security groups to associate with the elastic network interfaces in order to specify what the replicator has access to. If a security group is not specified, the default security group associated with the VPC is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "subnet_ids": {
                    "description": "The list of subnets to connect to in the virtual private cloud (VPC). AWS creates elastic network interfaces inside these subnets.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "replication_info_list": {
        "description": "A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumer_group_replication": {
              "description": "Configuration relating to consumer group replication.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "consumer_groups_to_exclude": {
                    "computed": true,
                    "description": "List of regular expression patterns indicating the consumer groups that should not be replicated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "consumer_groups_to_replicate": {
                    "description": "List of regular expression patterns indicating the consumer groups to copy.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "detect_and_copy_new_consumer_groups": {
                    "computed": true,
                    "description": "Whether to periodically check for new consumer groups.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "synchronise_consumer_group_offsets": {
                    "computed": true,
                    "description": "Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "source_kafka_cluster_arn": {
              "description": "Amazon Resource Name of the source Kafka cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_compression_type": {
              "description": "The type of compression to use writing records to target Kafka cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_kafka_cluster_arn": {
              "description": "Amazon Resource Name of the target Kafka cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "topic_replication": {
              "description": "Configuration relating to topic replication.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "copy_access_control_lists_for_topics": {
                    "computed": true,
                    "description": "Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "copy_topic_configurations": {
                    "computed": true,
                    "description": "Whether to periodically configure remote topics to match their corresponding upstream topics.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "detect_and_copy_new_topics": {
                    "computed": true,
                    "description": "Whether to periodically check for new topics and partitions.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "topics_to_exclude": {
                    "computed": true,
                    "description": "List of regular expression patterns indicating the topics that should not be replicated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "topics_to_replicate": {
                    "description": "List of regular expression patterns indicating the topics to copy.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "replicator_arn": {
        "computed": true,
        "description": "Amazon Resource Name for the created replicator.",
        "description_kind": "plain",
        "type": "string"
      },
      "replicator_name": {
        "description": "The name of the replicator.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_execution_role_arn": {
        "description": "The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::MSK::Replicator",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskReplicatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskReplicator), &result)
	return &result
}
