package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigExperimentRun = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The resolved application ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "application_identifier": {
        "description": "The application name or ID used to create the experiment run.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the experiment run.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "experiment_definition_id": {
        "computed": true,
        "description": "The resolved experiment definition ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "experiment_definition_identifier": {
        "description": "The experiment definition name or ID used to create the experiment run.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "exposure_percentage": {
        "description": "Percentage of traffic exposed to the experiment (0-100).",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "run": {
        "computed": true,
        "description": "The run number (auto-assigned by the service).",
        "description_kind": "plain",
        "type": "string"
      },
      "started_at": {
        "computed": true,
        "description": "ISO-8601 timestamp when the run started.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the run.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the experiment run.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "treatment_overrides": {
        "computed": true,
        "description": "Treatment overrides for specific entities.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inline": {
              "computed": true,
              "description": "Map of entity ID to treatment key (t1, t2, ..., or c for control).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "ISO-8601 timestamp when the run was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppConfig::ExperimentRun",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppconfigExperimentRunSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigExperimentRun), &result)
	return &result
}
