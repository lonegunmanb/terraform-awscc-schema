package resource

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
        "optional": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_groups": {
        "description": "The instance groups of the SageMaker HyperPod cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "current_count": {
              "computed": true,
              "description": "The number of instances that are currently in the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "execution_role": {
              "description": "The execution role for the instance group to assume.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_count": {
              "description": "The number of instances you specified to add to the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "instance_group_name": {
              "description": "The name of the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "required": true,
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
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "instance_type": {
              "description": "The instance type of the instance group of a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "life_cycle_config": {
              "description": "The lifecycle configuration for a SageMaker HyperPod cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_create": {
                    "description": "The file name of the entrypoint script of lifecycle scripts under SourceS3Uri. This entrypoint script runs during cluster creation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "source_s3_uri": {
                    "description": "An Amazon S3 bucket path where your lifecycle scripts are stored.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "on_start_deep_health_checks": {
              "computed": true,
              "description": "Nodes will undergo advanced stress test to detect and replace faulty instances, based on the type of deep health check(s) passed in.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "threads_per_core": {
              "computed": true,
              "description": "The number you specified to TreadsPerCore in CreateCluster for enabling or disabling multithreading. For instance types that support multithreading, you can specify 1 for disabling multithreading and 2 for enabling multithreading.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "node_recovery": {
        "computed": true,
        "description": "If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected. If set to false, nodes will be labelled when a fault is detected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "orchestrator": {
        "computed": true,
        "description": "Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "eks": {
              "description": "Specifies parameter(s) related to EKS as orchestrator, e.g. the EKS cluster nodes will attach to,",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_arn": {
                    "description": "The ARN of the EKS cluster, such as arn:aws:eks:us-west-2:123456789012:cluster/my-eks-cluster",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "Custom tags for managing the SageMaker HyperPod cluster as an AWS resource. You can add tags to your cluster in the same way you add them in other AWS services that support tagging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vpc_config": {
        "computed": true,
        "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "description": "The ID of the subnets in the VPC to which you want to connect your training job or model.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::Cluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerCluster), &result)
	return &result
}
