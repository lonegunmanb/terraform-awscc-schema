package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEvidentlyExperiment = `{
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_goals": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_change": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
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
        "required": true
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "online_ab_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "control_treatment_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "treatment_weights": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "split_weight": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "treatment": {
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
          "nesting_mode": "single"
        },
        "required": true
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
      "remove_segment": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "running_status": {
        "computed": true,
        "description": "Start Experiment. Default is False",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "analysis_complete_time": {
              "computed": true,
              "description": "Provide the analysis Completion time for an experiment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "desired_state": {
              "computed": true,
              "description": "Provide CANCELLED or COMPLETED desired state when stopping an experiment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reason": {
              "computed": true,
              "description": "Reason is a required input for stopping the experiment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "Provide START or STOP action to apply on an experiment",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sampling_rate": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "segment": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      },
      "treatments": {
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
            "treatment_name": {
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
      }
    },
    "description": "Resource Type definition for AWS::Evidently::Experiment.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEvidentlyExperimentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEvidentlyExperiment), &result)
	return &result
}
