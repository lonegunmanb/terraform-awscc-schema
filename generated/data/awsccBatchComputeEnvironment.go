package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchComputeEnvironment = `{
  "block": {
    "attributes": {
      "compute_environment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_environment_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_resources": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bid_percentage": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "desiredv_cpus": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "ec_2_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "image_id_override": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "image_kubernetes_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "image_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "ec_2_key_pair": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "image_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_role": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_types": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "launch_template": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "launch_template_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "launch_template_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "launch_template_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "target_instance_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "userdata_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "userdata_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maxv_cpus": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "minv_cpus": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "placement_group": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "spot_iam_fleet_role": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subnets": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "tags": {
              "computed": true,
              "description": "A key-value pair to associate with a resource.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "update_to_latest_image_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "context": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "eks_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "eks_cluster_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kubernetes_namespace": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replace_compute_environment": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "service_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "unmanagedv_cpus": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "update_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "job_execution_timeout_minutes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "terminate_jobs_on_update": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Batch::ComputeEnvironment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBatchComputeEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchComputeEnvironment), &result)
	return &result
}
