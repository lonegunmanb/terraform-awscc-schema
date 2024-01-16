package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncAgent = `{
  "block": {
    "attributes": {
      "activation_key": {
        "computed": true,
        "description": "Activation key of the Agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_arn": {
        "computed": true,
        "description": "The DataSync Agent ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_name": {
        "computed": true,
        "description": "The name configured for the agent. Text reference used to identify the agent in the console.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_type": {
        "computed": true,
        "description": "The service endpoints that the agent will connect to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_arns": {
        "computed": true,
        "description": "The ARNs of the security group used to protect your data transfer task subnets.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_arns": {
        "computed": true,
        "description": "The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_endpoint_id": {
        "computed": true,
        "description": "The ID of the VPC endpoint that the agent has access to.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataSync::Agent",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncAgent), &result)
	return &result
}
