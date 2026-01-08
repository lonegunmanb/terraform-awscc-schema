package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcsCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The unique Amazon Resource Name (ARN) of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_id": {
        "computed": true,
        "description": "The generated unique ID of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoints": {
        "computed": true,
        "description": "The list of endpoints available for interaction with the scheduler.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_address": {
              "computed": true,
              "description": "The endpoint's IPv6 address.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The endpoint's connection port number.",
              "description_kind": "plain",
              "type": "string"
            },
            "private_ip_address": {
              "computed": true,
              "description": "The endpoint's private IP address.",
              "description_kind": "plain",
              "type": "string"
            },
            "public_ip_address": {
              "computed": true,
              "description": "The endpoint's public IP address.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "Indicates the type of endpoint running at the specific IP address.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "error_info": {
        "computed": true,
        "description": "The list of errors that occurred during cluster provisioning.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "The short-form error code.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The detailed error information.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name that identifies the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "networking": {
        "description": "The networking configuration for the cluster's control plane.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_type": {
              "computed": true,
              "description": "The IP of the cluster (IPV4 or IPV6)",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_group_ids": {
              "computed": true,
              "description": "The list of security group IDs associated with the Elastic Network Interface (ENI) created in subnets.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "The list of subnet IDs where AWS PCS creates an Elastic Network Interface (ENI) to enable communication between managed controllers and AWS PCS resources. The subnet must have an available IP address, cannot reside in AWS Outposts, AWS Wavelength, or an AWS Local Zone. AWS PCS currently supports only 1 subnet in this list.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "scheduler": {
        "description": "The cluster management and job scheduling software associated with the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "description": "The software AWS PCS uses to manage cluster scaling and job scheduling.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "description": "The version of the specified scheduling software that AWS PCS uses to manage cluster scaling and job scheduling.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "size": {
        "description": "The size of the cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "slurm_configuration": {
        "computed": true,
        "description": "Additional options related to the Slurm scheduler.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "accounting": {
              "computed": true,
              "description": "The accounting configuration includes configurable settings for Slurm accounting.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_purge_time_in_days": {
                    "computed": true,
                    "description": "The default value for all purge settings for ` + "`" + `slurmdbd.conf` + "`" + `. For more information, see the [slurmdbd.conf documentation at SchedMD](https://slurm.schedmd.com/slurmdbd.conf.html). The default value is ` + "`" + `-1` + "`" + `. A value of ` + "`" + `-1` + "`" + ` means there is no purge time and records persist as long as the cluster exists.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "mode": {
                    "computed": true,
                    "description": "The default value is ` + "`" + `NONE` + "`" + `. A value of ` + "`" + `STANDARD` + "`" + ` means that Slurm accounting is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "auth_key": {
              "computed": true,
              "description": "The shared Slurm key for authentication, also known as the cluster secret.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "secret_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the the shared Slurm key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secret_version": {
                    "computed": true,
                    "description": "The version of the shared Slurm key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "jwt_auth": {
              "computed": true,
              "description": "JWT authentication configuration for Slurm.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "jwt_key": {
                    "computed": true,
                    "description": "JWT key configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the JWT key secret.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_version": {
                          "computed": true,
                          "description": "The version of the JWT key secret.",
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
              "optional": true
            },
            "scale_down_idle_time_in_seconds": {
              "computed": true,
              "description": "The time before an idle node is scaled down.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "slurm_custom_settings": {
              "computed": true,
              "description": "Additional Slurm-specific configuration that directly maps to Slurm settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_name": {
                    "computed": true,
                    "description": "AWS PCS supports configuration of the following Slurm parameters for clusters: Prolog, Epilog, and SelectTypeParameters.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameter_value": {
                    "computed": true,
                    "description": "The value for the configured Slurm setting.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "slurm_rest": {
              "computed": true,
              "description": "The SlurmRest configuration includes configurable settings for Slurm Rest.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mode": {
                    "computed": true,
                    "description": "The default value is ` + "`" + `NONE` + "`" + `. A value of ` + "`" + `STANDARD` + "`" + ` means that Slurm Rest is enabled.",
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
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The provisioning status of the cluster. The provisioning status doesn't indicate the overall health of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "AWS::PCS::Cluster resource creates an AWS PCS cluster.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcsClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcsCluster), &result)
	return &result
}
