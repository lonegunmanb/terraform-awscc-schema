package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStepfunctionsStateMachineAlias = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_preference": {
        "computed": true,
        "description": "The settings to enable gradual state machine deployments.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarms": {
              "computed": true,
              "description": "A list of CloudWatch alarm names that will be monitored during the deployment. The deployment will fail and rollback if any alarms go into ALARM state.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "interval": {
              "computed": true,
              "description": "The time in minutes between each traffic shifting increment.",
              "description_kind": "plain",
              "type": "number"
            },
            "percentage": {
              "computed": true,
              "description": "The percentage of traffic to shift to the new version in each increment.",
              "description_kind": "plain",
              "type": "number"
            },
            "state_machine_version_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of deployment to perform.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "An optional description of the alias.",
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
        "description": "The alias name.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_configuration": {
        "computed": true,
        "description": "The routing configuration of the alias. One or two versions can be mapped to an alias to split StartExecution requests of the same state machine.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "state_machine_version_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) that identifies one or two state machine versions defined in the routing configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description": "The percentage of traffic you want to route to the state machine version. The sum of the weights in the routing configuration must be equal to 100.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::StepFunctions::StateMachineAlias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccStepfunctionsStateMachineAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStepfunctionsStateMachineAlias), &result)
	return &result
}
