package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomAssistantAssociation = `{
  "block": {
    "attributes": {
      "assistant_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_association_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "association": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "knowledge_base_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "association_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Wisdom::AssistantAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWisdomAssistantAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomAssistantAssociation), &result)
	return &result
}
