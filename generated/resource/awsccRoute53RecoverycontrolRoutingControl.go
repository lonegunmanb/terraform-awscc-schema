package resource

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
        "optional": true,
        "type": "string"
      },
      "control_panel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the control panel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the routing control. You can use any non-white space character in the name.",
        "description_kind": "plain",
        "required": true,
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
    "description": "AWS Route53 Recovery Control Routing Control resource schema .",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoverycontrolRoutingControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoverycontrolRoutingControl), &result)
	return &result
}
