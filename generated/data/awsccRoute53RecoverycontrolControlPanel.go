package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoverycontrolControlPanel = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "Cluster to associate with the Control Panel",
        "description_kind": "plain",
        "type": "string"
      },
      "control_panel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_control_panel": {
        "computed": true,
        "description": "A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the control panel. You can use any non-white space character in the name.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_control_count": {
        "computed": true,
        "description": "Count of associated routing controls",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
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
    "description": "Data Source schema for AWS::Route53RecoveryControl::ControlPanel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53RecoverycontrolControlPanelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoverycontrolControlPanel), &result)
	return &result
}
