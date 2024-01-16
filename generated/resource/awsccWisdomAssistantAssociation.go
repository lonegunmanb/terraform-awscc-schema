package resource

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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "association": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "knowledge_base_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "association_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::Wisdom::AssistantAssociation Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWisdomAssistantAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomAssistantAssociation), &result)
	return &result
}
