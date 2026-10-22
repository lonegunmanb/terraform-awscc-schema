package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsGatewayClusterAssociations = `{
  "block": {
    "attributes": {
      "cluster": {
        "computed": true,
        "description": "The name or full Amazon Resource Name (ARN) of the cluster to associate the gateway with.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_arn": {
        "computed": true,
        "description": "The full Amazon Resource Name (ARN) of the associated cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway": {
        "computed": true,
        "description": "The name or full Amazon Resource Name (ARN) of the gateway to associate with the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_arn": {
        "computed": true,
        "description": "The full Amazon Resource Name (ARN) of the associated gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECS::GatewayClusterAssociations",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsGatewayClusterAssociationsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsGatewayClusterAssociations), &result)
	return &result
}
