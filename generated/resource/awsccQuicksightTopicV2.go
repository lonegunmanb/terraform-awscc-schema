package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightTopicV2 = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_instructions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_instructions_string": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "data_set_relations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "left": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_set_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "right": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_set_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "data_sets": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_set_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_set_name": {
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
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "folder_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "topic_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of the AWS::QuickSight::TopicV2 Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightTopicV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightTopicV2), &result)
	return &result
}
