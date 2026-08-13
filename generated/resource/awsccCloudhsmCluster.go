package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudhsmCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_policy": {
        "computed": true,
        "description": "The cluster's backup policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_retention_policy": {
        "computed": true,
        "description": "A policy that defines how the service retains backups.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "The type of backup retention policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Use a value between 7 - 379.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cluster_id": {
        "computed": true,
        "description": "The cluster's identifier (ID).",
        "description_kind": "plain",
        "type": "string"
      },
      "hsm_type": {
        "description": "The type of HSM to use in the cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "mode": {
        "computed": true,
        "description": "The mode to use in the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The NetworkType to create a cluster with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group": {
        "computed": true,
        "description": "The identifier (ID) of the cluster's security group.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The cluster's state.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "The identifiers (IDs) of the subnets where the cluster is created. You must specify at least one subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_mapping": {
        "computed": true,
        "description": "A map from availability zone to the cluster's subnet in that availability zone.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the CloudHSM cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description": "The identifier (ID) of the virtual private cloud (VPC) that contains the cluster.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates and manages an AWS CloudHSM cluster.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudhsmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudhsmCluster), &result)
	return &result
}
