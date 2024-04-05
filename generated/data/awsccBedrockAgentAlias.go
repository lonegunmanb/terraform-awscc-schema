package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockAgentAlias = `{
  "block": {
    "attributes": {
      "agent_alias_arn": {
        "computed": true,
        "description": "Arn representation of the Agent Alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_alias_history_events": {
        "computed": true,
        "description": "The list of history events for an alias for an Agent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_date": {
              "computed": true,
              "description": "Time Stamp.",
              "description_kind": "plain",
              "type": "string"
            },
            "routing_configuration": {
              "computed": true,
              "description": "Routing configuration for an Agent alias.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_version": {
                    "computed": true,
                    "description": "Agent Version.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "start_date": {
              "computed": true,
              "description": "Time Stamp.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "agent_alias_id": {
        "computed": true,
        "description": "Id for an Agent Alias generated at the server side.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_alias_name": {
        "computed": true,
        "description": "Name for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_alias_status": {
        "computed": true,
        "description": "The statuses an Agent Alias can be in.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_id": {
        "computed": true,
        "description": "Identifier for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_configuration": {
        "computed": true,
        "description": "Routing configuration for an Agent alias.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "agent_version": {
              "computed": true,
              "description": "Agent Version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::AgentAlias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockAgentAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockAgentAlias), &result)
	return &result
}
