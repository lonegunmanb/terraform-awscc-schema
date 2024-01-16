package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEvidentlyLaunch = `{
  "block": {
    "attributes": {
      "arn": {
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
      "execution_status": {
        "computed": true,
        "description": "Start or Stop Launch Launch. Default is not started.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_state": {
              "computed": true,
              "description": "Provide CANCELLED or COMPLETED as the launch desired state. Defaults to Completed if not provided.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reason": {
              "computed": true,
              "description": "Provide a reason for stopping the launch. Defaults to empty if not provided.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "description": "Provide START or STOP action to apply on a launch",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "groups": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "feature": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "group_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "variation": {
              "description_kind": "plain",
              "required": true,
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
      "metric_monitors": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "entity_id_key": {
              "description": "The JSON path to reference the entity id in the event.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "event_pattern": {
              "computed": true,
              "description": "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "unit_label": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value_key": {
              "description": "The JSON path to reference the numerical metric value in the event.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "randomization_salt": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduled_splits_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group_weights": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "split_weight": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              },
              "required": true
            },
            "segment_overrides": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "evaluation_order": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "segment": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "weights": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "group_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "split_weight": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "start_time": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Type definition for AWS::Evidently::Launch.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEvidentlyLaunchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEvidentlyLaunch), &result)
	return &result
}
