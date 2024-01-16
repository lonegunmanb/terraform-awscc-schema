package resource

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
        "optional": true,
        "type": "string"
      },
      "entity_types": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
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
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "event_variables": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
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
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "variable_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "created_time": {
              "computed": true,
              "description": "The time when the event type was created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inline": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "last_updated_time": {
              "computed": true,
              "description": "The time when the event type was last updated.",
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
            "tags": {
              "computed": true,
              "description": "Tags associated with this event type.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "last_updated_time": {
        "computed": true,
        "description": "The time when the event type was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name for the event type",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with this event type.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "A resource schema for an EventType in Amazon Fraud Detector.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFrauddetectorEventTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFrauddetectorEventType), &result)
	return &result
}
