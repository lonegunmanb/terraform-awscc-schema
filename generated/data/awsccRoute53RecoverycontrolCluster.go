package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of a Cluster. You can use any non-white space character in the name",
        "description_kind": "plain",
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
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Route53RecoveryControl::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53RecoverycontrolClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoverycontrolCluster), &result)
	return &result
}
