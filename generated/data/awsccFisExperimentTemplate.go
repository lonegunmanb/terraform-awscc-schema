package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFisExperimentTemplate = `{
  "block": {
    "attributes": {
      "actions": {
        "computed": true,
        "description": "The actions for the experiment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action_id": {
              "computed": true,
              "description": "The ID of the action.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "A description for the action.",
              "description_kind": "plain",
              "type": "string"
            },
            "parameters": {
              "computed": true,
              "description": "The parameters for the action, if applicable.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "start_after": {
              "computed": true,
              "description": "The names of the actions that must be completed before the current action starts.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "targets": {
              "computed": true,
              "description": "One or more targets for the action.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "map"
        }
      },
      "description": {
        "computed": true,
        "description": "A description for the experiment template.",
        "description_kind": "plain",
        "type": "string"
      },
      "experiment_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_targeting": {
              "computed": true,
              "description": "The account targeting setting for the experiment template.",
              "description_kind": "plain",
              "type": "string"
            },
            "empty_target_resolution_mode": {
              "computed": true,
              "description": "The target resolution failure mode for the experiment template.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "log_schema_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "s3_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.",
        "description_kind": "plain",
        "type": "string"
      },
      "stop_conditions": {
        "computed": true,
        "description": "One or more stop conditions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "targets": {
        "computed": true,
        "description": "The targets for the experiment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "path": {
                    "computed": true,
                    "description": "The attribute path for the filter.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
                    "description": "The attribute values for the filter.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "resource_arns": {
              "computed": true,
              "description": "The Amazon Resource Names (ARNs) of the target resources.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "resource_tags": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "resource_type": {
              "computed": true,
              "description": "The AWS resource type. The resource type must be supported for the specified action.",
              "description_kind": "plain",
              "type": "string"
            },
            "selection_mode": {
              "computed": true,
              "description": "Scopes the identified resources to a specific number of the resources at random, or a percentage of the resources.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      }
    },
    "description": "Data Source schema for AWS::FIS::ExperimentTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFisExperimentTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFisExperimentTemplate), &result)
	return &result
}
