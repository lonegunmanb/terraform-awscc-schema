package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsAnalysisTemplate = `{
  "block": {
    "attributes": {
      "analysis_parameters": {
        "computed": true,
        "description": "The member who can query can provide this placeholder for a literal data value in an analysis template",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "analysis_template_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "format": {
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
      "membership_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "membership_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "referenced_tables": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "source": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "text": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template.",
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
    "description": "Represents a stored analysis within a collaboration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCleanroomsAnalysisTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsAnalysisTemplate), &result)
	return &result
}
