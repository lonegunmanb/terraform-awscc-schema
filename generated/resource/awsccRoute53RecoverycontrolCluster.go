package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoverycontrolCluster = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_endpoints": {
        "computed": true,
        "description": "Endpoints for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "region": {
              "computed": true,
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
        "description": "Name of a Cluster. You can use any non-white space character in the name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
        "description_kind": "plain",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "AWS Route53 Recovery Control Cluster resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoverycontrolClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoverycontrolCluster), &result)
	return &result
}
