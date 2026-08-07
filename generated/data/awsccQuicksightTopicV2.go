package data

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
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_set_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "right": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "data_set_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "data_sets": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_set_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "data_set_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "folder_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
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
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "topic_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::TopicV2",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightTopicV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightTopicV2), &result)
	return &result
}
