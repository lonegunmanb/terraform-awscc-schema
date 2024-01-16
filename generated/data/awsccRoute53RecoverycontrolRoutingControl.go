package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoverycontrolRoutingControl = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "Arn associated with Control Panel",
        "description_kind": "plain",
        "type": "string"
      },
      "control_panel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the control panel.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the routing control. You can use any non-white space character in the name.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_control_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the routing control.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53RecoveryControl::RoutingControl",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53RecoverycontrolRoutingControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoverycontrolRoutingControl), &result)
	return &result
}
