package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksNodegroup = `{
  "block": {
    "attributes": {
      "ami_type": {
        "computed": true,
        "description": "The AMI type for your node group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_type": {
        "computed": true,
        "description": "The capacity type of your managed node group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_name": {
        "description": "Name of the cluster to create the node group in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "disk_size": {
        "computed": true,
        "description": "The root device disk size (in GiB) for your node group instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "force_update_enabled": {
        "computed": true,
        "description": "Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_types": {
        "computed": true,
        "description": "Specify the instance types for a node group.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "labels": {
        "computed": true,
        "description": "The Kubernetes labels to be applied to the nodes in the node group when they are created.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "launch_template": {
        "computed": true,
        "description": "An object representing a node group's launch template specification.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "node_role": {
        "description": "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nodegroup_name": {
        "computed": true,
        "description": "The unique name to give your node group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_version": {
        "computed": true,
        "description": "The AMI version of the Amazon EKS-optimized AMI to use with your node group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_access": {
        "computed": true,
        "description": "The remote access (SSH) configuration to use with your node group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ec_2_ssh_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_security_groups": {
              "computed": true,
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
        "optional": true
      },
      "scaling_config": {
        "computed": true,
        "description": "The scaling configuration details for the Auto Scaling group that is created for your node group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "subnets": {
        "description": "The subnets to use for the Auto Scaling group that is created for your node group.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "taints": {
        "computed": true,
        "description": "The Kubernetes taints to be applied to the nodes in the node group when they are created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "effect": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
      },
      "update_config": {
        "computed": true,
        "description": "The node group update configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_unavailable": {
              "computed": true,
              "description": "The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. ",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_unavailable_percentage": {
              "computed": true,
              "description": "The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "The Kubernetes version to use for your managed nodes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::EKS::Nodegroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEksNodegroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksNodegroup), &result)
	return &result
}
