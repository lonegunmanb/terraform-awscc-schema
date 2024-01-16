package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFrauddetectorEventType = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the event type.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The time when the event type was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the event type.",
        "description_kind": "plain",
        "type": "string"
      },
      "entity_types": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "type": "string"
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "event_variables": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "data_source": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "data_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "type": "string"
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
                "nesting_mode": "list"
              }
            },
            "variable_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "type": "string"
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "last_updated_time": {
        "computed": true,
        "description": "The time when the event type was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the event type",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with this event type.",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::FraudDetector::EventType",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFrauddetectorEventTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFrauddetectorEventType), &result)
	return &result
}
