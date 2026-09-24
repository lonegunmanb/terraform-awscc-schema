package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallContainerAssociation = `{
  "block": {
    "attributes": {
      "container_association_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the container association.",
        "description_kind": "plain",
        "type": "string"
      },
      "container_association_name": {
        "computed": true,
        "description": "The descriptive name of the container association. You can't change the name of a container association after you create it.",
        "description_kind": "plain",
        "type": "string"
      },
      "container_monitoring_configurations": {
        "computed": true,
        "description": "The monitoring configurations for the container association. Each configuration specifies an Amazon ECS or Amazon EKS cluster to monitor and optional attribute filters to narrow which containers are tracked.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_filters": {
              "computed": true,
              "description": "Key-value pairs that filter which containers are tracked. For Amazon EKS, you can filter by namespace and Kubernetes labels. For Amazon ECS, you can filter by container instance attributes (EC2 launch type only).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The attribute key to filter on.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The attribute value to match.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "cluster_arn": {
              "computed": true,
              "description": "The ARN of the Amazon ECS or Amazon EKS cluster to monitor. The cluster must be in the same Region and account as the container association.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the container association.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resolved_cidr_count": {
        "computed": true,
        "description": "The number of CIDR blocks resolved from the monitored containers.",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The current status of the container association.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The part of the key:value pair that defines a tag. Tag keys are case-sensitive.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The part of the key:value pair that defines a tag. Tag values are case-sensitive.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of containers to monitor. You can't change the container type after creation.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkFirewall::ContainerAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkfirewallContainerAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallContainerAssociation), &result)
	return &result
}
