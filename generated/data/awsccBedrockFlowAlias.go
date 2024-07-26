package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockFlowAlias = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Arn of the Flow Alias",
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
      "flow_alias_id": {
        "computed": true,
        "description": "Id for a Flow Alias generated at the server side.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_arn": {
        "computed": true,
        "description": "Arn representation of the Flow",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_id": {
        "computed": true,
        "description": "Identifier for a flow resource.",
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
        "description": "Name for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_configuration": {
        "computed": true,
        "description": "Routing configuration for a Flow alias.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "flow_version": {
              "computed": true,
              "description": "Version.",
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
    "description": "Data Source schema for AWS::Bedrock::FlowAlias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockFlowAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockFlowAlias), &result)
	return &result
}
