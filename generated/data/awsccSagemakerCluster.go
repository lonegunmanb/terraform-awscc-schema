package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerCluster = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the HyperPod Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The name of the HyperPod Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_status": {
        "computed": true,
        "description": "The status of the HyperPod Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The time at which the HyperPod cluster was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_message": {
        "computed": true,
        "description": "The failure message of the HyperPod Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_groups": {
        "computed": true,
        "description": "The instance groups of the SageMaker HyperPod cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "current_count": {
              "computed": true,
              "description": "The number of instances that are currently in the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "execution_role": {
              "computed": true,
              "description": "The execution role for the instance group to assume.",
              "description_kind": "plain",
              "type": "string"
            },
            "image_id": {
              "computed": true,
              "description": "AMI Id to be used for launching EC2 instances - HyperPodPublicAmiId or CustomAmiId",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_count": {
              "computed": true,
              "description": "The number of instances you specified to add to the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "instance_group_name": {
              "computed": true,
              "description": "The name of the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_storage_configs": {
              "computed": true,
              "description": "The instance storage configuration for the instance group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ebs_volume_config": {
                    "computed": true,
                    "description": "Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "volume_size_in_gb": {
                          "computed": true,
                          "description": "The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "instance_type": {
              "computed": true,
              "description": "The instance type of the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "life_cycle_config": {
              "computed": true,
              "description": "The lifecycle configuration for a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_create": {
                    "computed": true,
                    "description": "The file name of the entrypoint script of lifecycle scripts under SourceS3Uri. This entrypoint script runs during cluster creation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_s3_uri": {
                    "computed": true,
                    "description": "An Amazon S3 bucket path where your lifecycle scripts are stored.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "on_start_deep_health_checks": {
              "computed": true,
              "description": "Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "override_vpc_config": {
              "computed": true,
              "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "The ID of the subnets in the VPC to which you want to connect your training job or model.",
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
            "scheduled_update_config": {
              "computed": true,
              "description": "The configuration object of the schedule that SageMaker follows when updating the AMI.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "deployment_config": {
                    "computed": true,
                    "description": "The configuration to use when updating the AMI versions.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_rollback_configuration": {
                          "computed": true,
                          "description": "An array that contains the alarms that SageMaker monitors to know whether to roll back the AMI update.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "alarm_name": {
                                "computed": true,
                                "description": "The name of the alarm.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "rolling_update_policy": {
                          "computed": true,
                          "description": "The policy that SageMaker uses when updating the AMI versions of the cluster.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "maximum_batch_size": {
                                "computed": true,
                                "description": "The configuration of the size measurements of the AMI update. Using this configuration, you can specify whether SageMaker should update your instance group by an amount or percentage of instances.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "type": {
                                      "computed": true,
                                      "description": "Specifies whether SageMaker should process the update by amount or percentage of instances.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "Specifies the amount or percentage of instances SageMaker updates at a time.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "rollback_maximum_batch_size": {
                                "computed": true,
                                "description": "The configuration of the size measurements of the AMI update. Using this configuration, you can specify whether SageMaker should update your instance group by an amount or percentage of instances.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "type": {
                                      "computed": true,
                                      "description": "Specifies whether SageMaker should process the update by amount or percentage of instances.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "Specifies the amount or percentage of instances SageMaker updates at a time.",
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
                        "wait_interval_in_seconds": {
                          "computed": true,
                          "description": "The duration in seconds that SageMaker waits before updating more instances in the cluster.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "schedule_expression": {
                    "computed": true,
                    "description": "A cron expression that specifies the schedule that SageMaker follows when updating the AMI.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "threads_per_core": {
              "computed": true,
              "description": "The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading. For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.",
              "description_kind": "plain",
              "type": "number"
            },
            "training_plan_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the training plan to use for this cluster instance group. For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "node_provisioning_mode": {
        "computed": true,
        "description": "Determines the scaling strategy for the SageMaker HyperPod cluster. When set to 'Continuous', enables continuous scaling which dynamically manages node provisioning. If the parameter is omitted, uses the standard scaling approach in previous release.",
        "description_kind": "plain",
        "type": "string"
      },
      "node_recovery": {
        "computed": true,
        "description": "If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected. If set to false, nodes will be labelled when a fault is detected.",
        "description_kind": "plain",
        "type": "string"
      },
      "orchestrator": {
        "computed": true,
        "description": "Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "eks": {
              "computed": true,
              "description": "Specifies parameter(s) related to EKS as orchestrator, e.g. the EKS cluster nodes will attach to,",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_arn": {
                    "computed": true,
                    "description": "The ARN of the EKS cluster, such as arn:aws:eks:us-west-2:123456789012:cluster/my-eks-cluster",
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
      "restricted_instance_groups": {
        "computed": true,
        "description": "The restricted instance groups of the SageMaker HyperPod cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "current_count": {
              "computed": true,
              "description": "The number of instances that are currently in the restricted instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "environment_config": {
              "computed": true,
              "description": "The configuration for the restricted instance groups (RIG) environment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "fsx_lustre_config": {
                    "computed": true,
                    "description": "Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "per_unit_storage_throughput": {
                          "computed": true,
                          "description": "The throughput capacity of the FSx for Lustre file system, measured in MB/s per TiB of storage.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "size_in_gi_b": {
                          "computed": true,
                          "description": "The storage capacity of the FSx for Lustre file system, specified in gibibytes (GiB).",
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
            "execution_role": {
              "computed": true,
              "description": "The execution role for the instance group to assume.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_count": {
              "computed": true,
              "description": "The number of instances you specified to add to the restricted instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "instance_group_name": {
              "computed": true,
              "description": "The name of the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_storage_configs": {
              "computed": true,
              "description": "The instance storage configuration for the instance group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ebs_volume_config": {
                    "computed": true,
                    "description": "Defines the configuration for attaching additional Amazon Elastic Block Store (EBS) volumes to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "volume_size_in_gb": {
                          "computed": true,
                          "description": "The size in gigabytes (GB) of the additional EBS volume to be attached to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each instance within the SageMaker HyperPod cluster instance group and mounted to /opt/sagemaker.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "instance_type": {
              "computed": true,
              "description": "The instance type of the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "on_start_deep_health_checks": {
              "computed": true,
              "description": "Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "override_vpc_config": {
              "computed": true,
              "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "The ID of the subnets in the VPC to which you want to connect your training job or model.",
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
            "threads_per_core": {
              "computed": true,
              "description": "The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading. For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.",
              "description_kind": "plain",
              "type": "number"
            },
            "training_plan_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the training plan to use for this cluster restricted instance group. For more information about how to reserve GPU capacity for your SageMaker HyperPod clusters using Amazon SageMaker Training Plan, see CreateTrainingPlan.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "Custom tags for managing the SageMaker HyperPod cluster as an AWS resource. You can add tags to your cluster in the same way you add them in other AWS services that support tagging.",
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
          "nesting_mode": "set"
        }
      },
      "vpc_config": {
        "computed": true,
        "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "computed": true,
              "description": "The ID of the subnets in the VPC to which you want to connect your training job or model.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerCluster), &result)
	return &result
}
