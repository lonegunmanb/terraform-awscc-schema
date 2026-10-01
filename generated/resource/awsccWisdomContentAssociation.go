package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomContentAssociation = `{
  "block": {
    "attributes": {
      "association": {
        "description": "The identifier of the associated resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_connect_guide_association": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "flow_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "association_type": {
        "description": "The type of association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_association_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the content association.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_association_id": {
        "computed": true,
        "description": "The identifier of the content association.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_id": {
        "description": "The identifier of the content.",
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
      "knowledge_base_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_id": {
        "description": "The identifier of the knowledge base.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::Wisdom::ContentAssociation Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWisdomContentAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomContentAssociation), &result)
	return &result
}
