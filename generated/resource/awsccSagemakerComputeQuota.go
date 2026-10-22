package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerComputeQuota = `{
  "block": {
    "attributes": {
      "activation_state": {
        "computed": true,
        "description": "Whether the compute allocation rule is active. Defaults service-side to 'Enabled' when omitted on Create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_arn": {
        "description": "The ARN of the EKS-orchestrated SageMaker HyperPod cluster the compute allocation definition is bound to. Slurm-orchestrated clusters are rejected by the service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "compute_quota_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the compute allocation definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_quota_config": {
        "description": "Compute allocation rules: per-resource quotas, idle-borrowing strategy, and team-task preemption policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_quota_resources": {
              "computed": true,
              "description": "Per-instance-type compute quotas for the team. When VCpu / MemoryInGiB / Accelerators are partially specified, SageMaker auto-fills the remaining fields ratio-proportionally to the underlying instance ? Read may return values the customer did not supply (drift may flag this).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "accelerator_partition": {
                    "computed": true,
                    "description": "Multi-instance GPU (MIG) partition profile and count, for fractional accelerator allocation. Service rejects MIG profiles that the chosen InstanceType does not support.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "count": {
                          "computed": true,
                          "description": "The number of MIG partitions of the specified Type allocated to the team.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "type": {
                          "computed": true,
                          "description": "MIG profile type. Service rejects MIG profiles that the chosen InstanceType does not support.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "accelerators": {
                    "computed": true,
                    "description": "The number of accelerators (e.g., GPUs) of this type allocated to the team. SageMaker auto-fills VCpu / MemoryInGiB ratio-proportionally if only this is set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "count": {
                    "computed": true,
                    "description": "The number of instances of this type allocated to the team.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "instance_type": {
                    "computed": true,
                    "description": "The cluster instance type the quota applies to (e.g., 'ml.m5.xlarge', 'ml.p5.48xlarge'). Validated by the service against the supported HyperPod instance types.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "memory_in_gi_b": {
                    "computed": true,
                    "description": "The memory in GiB allocated to the team. Floating-point. SageMaker auto-fills Accelerators / VCpu ratio-proportionally if only this is set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "v_cpu": {
                    "computed": true,
                    "description": "The number of vCPUs allocated to the team. Floating-point. SageMaker auto-fills Accelerators / MemoryInGiB ratio-proportionally if only this is set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "preempt_team_tasks": {
              "computed": true,
              "description": "Team-task preemption policy. Defaults service-side to 'LowerPriority' when omitted on Create.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_sharing_config": {
              "computed": true,
              "description": "Idle-resource-sharing strategy for this team's quota. When omitted, the service applies defaults (Strategy=LendAndBorrow, BorrowLimit=50).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "absolute_borrow_limits": {
                    "computed": true,
                    "description": "Per-instance-type absolute caps on resources this team can borrow.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "accelerator_partition": {
                          "computed": true,
                          "description": "Multi-instance GPU (MIG) partition profile and count, for fractional accelerator allocation. Service rejects MIG profiles that the chosen InstanceType does not support.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "count": {
                                "computed": true,
                                "description": "The number of MIG partitions of the specified Type allocated to the team.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "type": {
                                "computed": true,
                                "description": "MIG profile type. Service rejects MIG profiles that the chosen InstanceType does not support.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "accelerators": {
                          "computed": true,
                          "description": "The number of accelerators (e.g., GPUs) of this type allocated to the team. SageMaker auto-fills VCpu / MemoryInGiB ratio-proportionally if only this is set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "count": {
                          "computed": true,
                          "description": "The number of instances of this type allocated to the team.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "instance_type": {
                          "computed": true,
                          "description": "The cluster instance type the quota applies to (e.g., 'ml.m5.xlarge', 'ml.p5.48xlarge'). Validated by the service against the supported HyperPod instance types.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "memory_in_gi_b": {
                          "computed": true,
                          "description": "The memory in GiB allocated to the team. Floating-point. SageMaker auto-fills Accelerators / VCpu ratio-proportionally if only this is set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "v_cpu": {
                          "computed": true,
                          "description": "The number of vCPUs allocated to the team. Floating-point. SageMaker auto-fills Accelerators / MemoryInGiB ratio-proportionally if only this is set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "borrow_limit": {
                    "computed": true,
                    "description": "Maximum percentage of the team's idle quota that can be borrowed by other teams. Smithy admits 0..10000; in practice the service-documented operational range is 1..500. Defaults service-side to 50 when omitted on Create.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "strategy": {
                    "computed": true,
                    "description": "How this team participates in idle-resource sharing.",
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
        "required": true
      },
      "compute_quota_id": {
        "computed": true,
        "description": "The unique 12-character identifier of the compute allocation definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_quota_target": {
        "description": "The team this compute allocation rule applies to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fair_share_weight": {
              "computed": true,
              "description": "Relative fair-share weight when fair-share scheduling is enabled on the parent ClusterSchedulerConfig. Defaults service-side to 0 when omitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "team_name": {
              "description": "The name of the team. Materialized as Kubernetes namespace 'hyperpod-ns-\u003cTeamName\u003e' on the EKS cluster, so the value must satisfy RFC 1123 namespace constraints. Reassignment is not supported (createOnly).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "compute_quota_version": {
        "computed": true,
        "description": "The version of the compute allocation definition. Incremented on each successful Update; used for optimistic concurrency control.",
        "description_kind": "plain",
        "type": "number"
      },
      "creation_time": {
        "computed": true,
        "description": "The ISO-8601 timestamp at which the compute allocation definition was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A free-form description of the compute allocation definition.",
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
      "last_modified_time": {
        "computed": true,
        "description": "The ISO-8601 timestamp at which the compute allocation definition was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the compute allocation definition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current lifecycle status of the compute allocation definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the compute allocation definition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "Resource Type definition for AWS::SageMaker::ComputeQuota. Defines a compute allocation rule for an EKS-orchestrated SageMaker HyperPod cluster's task governance. Available in 15 regions (us-east-1, us-east-2, us-west-1, us-west-2, ap-south-1, ap-southeast-1, ap-southeast-2, ap-southeast-3, ap-northeast-1, eu-central-1, eu-west-1, eu-west-2, eu-north-1, eu-south-2, sa-east-1).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerComputeQuotaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerComputeQuota), &result)
	return &result
}
