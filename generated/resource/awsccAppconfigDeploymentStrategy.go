package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigDeploymentStrategy = `{
  "block": {
    "attributes": {
      "deployment_duration_in_minutes": {
        "description": "Total amount of time for a deployment to last.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "deployment_strategy_id": {
        "computed": true,
        "description": "The deployment strategy ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the deployment strategy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "final_bake_time_in_minutes": {
        "computed": true,
        "description": "Specifies the amount of time AWS AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete. If an alarm is triggered during this time, AWS AppConfig rolls back the deployment. You must configure permissions for AWS AppConfig to roll back based on CloudWatch alarms. For more information, see Configuring permissions for rollback based on Amazon CloudWatch alarms in the AWS AppConfig User Guide.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "growth_factor": {
        "description": "The percentage of targets to receive a deployed configuration during each interval.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "growth_type": {
        "computed": true,
        "description": "The algorithm used to define how percentage grows over time. AWS AppConfig supports the following growth types:\n\nLinear: For this type, AWS AppConfig processes the deployment by dividing the total number of targets by the value specified for Step percentage. For example, a linear deployment that uses a Step percentage of 10 deploys the configuration to 10 percent of the hosts. After those deployments are complete, the system deploys the configuration to the next 10 percent. This continues until 100% of the targets have successfully received the configuration.\n\nExponential: For this type, AWS AppConfig processes the deployment exponentially using the following formula: G*(2^N). In this formula, G is the growth factor specified by the user and N is the number of steps until the configuration is deployed to all targets. For example, if you specify a growth factor of 2, then the system rolls out the configuration as follows:\n\n2*(2^0)\n\n2*(2^1)\n\n2*(2^2)\n\nExpressed numerically, the deployment rolls out as follows: 2% of the targets, 4% of the targets, 8% of the targets, and continues until the configuration has been deployed to all targets.",
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
        "description": "A name for the deployment strategy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replicate_to": {
        "description": "Save the deployment strategy to a Systems Manager (SSM) document.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Assigns metadata to an AWS AppConfig resource. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. You can specify a maximum of 50 tags for a resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::AppConfig::DeploymentStrategy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppconfigDeploymentStrategySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigDeploymentStrategy), &result)
	return &result
}
