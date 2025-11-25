package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksCluster = `{
  "block": {
    "attributes": {
      "access_config": {
        "computed": true,
        "description": "An object representing the Access Config to use for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_mode": {
              "computed": true,
              "description": "Specify the authentication mode that should be used to create your cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "bootstrap_cluster_creator_admin_permissions": {
              "computed": true,
              "description": "Set this value to false to avoid creating a default cluster admin Access Entry using the IAM principal used to create the cluster.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod.",
        "description_kind": "plain",
        "type": "string"
      },
      "bootstrap_self_managed_addons": {
        "computed": true,
        "description": "Set this value to false to avoid creating the default networking add-ons when the cluster is created.",
        "description_kind": "plain",
        "type": "bool"
      },
      "certificate_authority_data": {
        "computed": true,
        "description": "The certificate-authority-data for your cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_id": {
        "computed": true,
        "description": "The unique ID given to your cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_security_group_id": {
        "computed": true,
        "description": "The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_config": {
        "computed": true,
        "description": "Todo: add description",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Todo: add description",
              "description_kind": "plain",
              "type": "bool"
            },
            "node_pools": {
              "computed": true,
              "description": "Todo: add description",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "node_role_arn": {
              "computed": true,
              "description": "Todo: add description",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "control_plane_scaling_config": {
        "computed": true,
        "description": "Configuration for provisioned control plane scaling.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "tier": {
              "computed": true,
              "description": "The scaling tier for the provisioned control plane.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "deletion_protection": {
        "computed": true,
        "description": "Set this value to true to enable deletion protection for the cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "encryption_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "provider": {
              "computed": true,
              "description": "The encryption provider for the cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key_arn": {
                    "computed": true,
                    "description": "Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "resources": {
              "computed": true,
              "description": "Specifies the resources to be encrypted. The only supported value is \"secrets\".",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "encryption_config_key_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) or alias of the customer master key (CMK).",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com.",
        "description_kind": "plain",
        "type": "string"
      },
      "force": {
        "computed": true,
        "description": "Force cluster version update",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kubernetes_network_config": {
        "computed": true,
        "description": "The Kubernetes network configuration for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "elastic_load_balancing": {
              "computed": true,
              "description": "Todo: add description",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Todo: add description",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ip_family": {
              "computed": true,
              "description": "Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on",
              "description_kind": "plain",
              "type": "string"
            },
            "service_ipv_4_cidr": {
              "computed": true,
              "description": "The CIDR block to assign Kubernetes service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. ",
              "description_kind": "plain",
              "type": "string"
            },
            "service_ipv_6_cidr": {
              "computed": true,
              "description": "The CIDR block to assign Kubernetes service IP addresses from.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "logging": {
        "computed": true,
        "description": "Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cluster_logging": {
              "computed": true,
              "description": "The cluster control plane logging configuration for your cluster. ",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled_types": {
                    "computed": true,
                    "description": "Enable control plane logs for your cluster, all log types will be disabled if the array is empty",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "name of the log type",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
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
        "description": "The unique name to give to your cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "open_id_connect_issuer_url": {
        "computed": true,
        "description": "The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template.",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_config": {
        "computed": true,
        "description": "An object representing the Outpost configuration to use for AWS EKS outpost cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "control_plane_instance_type": {
              "computed": true,
              "description": "Specify the Instance type of the machines that should be used to create your cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "control_plane_placement": {
              "computed": true,
              "description": "Specify the placement group of the control plane machines for your cluster.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_name": {
                    "computed": true,
                    "description": "Specify the placement group name of the control place machines for your cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "outpost_arns": {
              "computed": true,
              "description": "Specify one or more Arn(s) of Outpost(s) on which you would like to create your cluster.",
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
      "remote_network_config": {
        "computed": true,
        "description": "Configuration fields for specifying on-premises node and pod CIDRs that are external to the VPC passed during cluster creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "remote_node_networks": {
              "computed": true,
              "description": "Network configuration of nodes run on-premises with EKS Hybrid Nodes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidrs": {
                    "computed": true,
                    "description": "Specifies the list of remote node CIDRs.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "remote_pod_networks": {
              "computed": true,
              "description": "Network configuration of pods run on-premises with EKS Hybrid Nodes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidrs": {
                    "computed": true,
                    "description": "Specifies the list of remote pod CIDRs.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resources_vpc_config": {
        "computed": true,
        "description": "An object representing the VPC configuration to use for an Amazon EKS cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_private_access": {
              "computed": true,
              "description": "Set this value to true to enable private access for your cluster's Kubernetes API server endpoint. If you enable private access, Kubernetes API requests from within your cluster's VPC use the private VPC endpoint. The default value for this parameter is false, which disables private access for your Kubernetes API server. If you disable private access and you have nodes or AWS Fargate pods in the cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks for communication with the nodes or Fargate pods.",
              "description_kind": "plain",
              "type": "bool"
            },
            "endpoint_public_access": {
              "computed": true,
              "description": "Set this value to false to disable public access to your cluster's Kubernetes API server endpoint. If you disable public access, your cluster's Kubernetes API server can only receive requests from within the cluster VPC. The default value for this parameter is true, which enables public access for your Kubernetes API server.",
              "description_kind": "plain",
              "type": "bool"
            },
            "public_access_cidrs": {
              "computed": true,
              "description": "The CIDR blocks that are allowed access to your cluster's public Kubernetes API server endpoint. Communication to the endpoint from addresses outside of the CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If you've disabled private endpoint access and you have nodes or AWS Fargate pods in the cluster, then ensure that you specify the necessary CIDR blocks.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "security_group_ids": {
              "computed": true,
              "description": "Specify one or more security groups for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane. If you don't specify a security group, the default security group for your VPC is used.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your nodes and the Kubernetes control plane.",
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
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_config": {
        "computed": true,
        "description": "Todo: add description",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_storage": {
              "computed": true,
              "description": "Todo: add description",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Todo: add description",
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
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
      "upgrade_policy": {
        "computed": true,
        "description": "An object representing the Upgrade Policy to use for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "support_type": {
              "computed": true,
              "description": "Specify the support type for your cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version": {
        "computed": true,
        "description": "The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "zonal_shift_config": {
        "computed": true,
        "description": "The current zonal shift configuration to use for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Set this value to true to enable zonal shift for the cluster.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::EKS::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEksClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksCluster), &result)
	return &result
}
