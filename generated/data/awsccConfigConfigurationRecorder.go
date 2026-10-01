package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigConfigurationRecorder = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the configuration recorder. By default, AWS Config assigns the name \"default\" when creating the configuration recorder. To change the configuration recorder name, you must use the DeleteConfigurationRecorder action to delete your current configuration recorder, and then you must use the PutConfigurationRecorder command to create a configuration recorder that has the desired name.",
        "description_kind": "plain",
        "type": "string"
      },
      "recording_group": {
        "computed": true,
        "description": "Specifies which resource types are in scope for the configuration recorder to record.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "all_supported": {
              "computed": true,
              "description": "Specifies whether AWS Config records configuration changes for all supported resource types, excluding the global IAM resource types.",
              "description_kind": "plain",
              "type": "bool"
            },
            "exclusion_by_resource_types": {
              "computed": true,
              "description": "An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_types": {
                    "computed": true,
                    "description": "A comma-separated list of resource types to exclude from recording by the configuration recorder.",
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
            "include_global_resource_types": {
              "computed": true,
              "description": "This option is a bundle which only applies to the global IAM resource types: IAM users, groups, roles, and customer managed policies. ",
              "description_kind": "plain",
              "type": "bool"
            },
            "recording_strategy": {
              "computed": true,
              "description": "An object that specifies the recording strategy for the configuration recorder.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "use_only": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "resource_types": {
              "computed": true,
              "description": "A comma-separated list that specifies which resource types AWS Config records.",
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
      "recording_mode": {
        "computed": true,
        "description": "Specifies the default recording frequency for the configuration recorder.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "recording_frequency": {
              "computed": true,
              "description": "The default recording frequency that AWS Config uses to record configuration changes.",
              "description_kind": "plain",
              "type": "string"
            },
            "recording_mode_overrides": {
              "computed": true,
              "description": "An array of 'RecordingModeOverride' objects for you to specify your overrides for the recording mode. ",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "A description that you provide for the override.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "recording_frequency": {
                    "computed": true,
                    "description": "The recording frequency that will be applied to all the resource types specified in the override.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_types": {
                    "computed": true,
                    "description": "A comma-separated list that specifies which resource types AWS Config includes in the override.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_arn": {
        "computed": true,
        "description": "Returns the Amazon Resource Name (ARN) for the Configuration Recorder.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the role that allows AWS Config to read S3 objects.",
        "description_kind": "plain",
        "type": "string"
      },
      "started_on_create": {
        "computed": true,
        "description": "Defaults to 'true'. Controls whether the recorder starts recording after the create operation. Set this to 'false' for development and testing purposes.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::Config::ConfigurationRecorder",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigConfigurationRecorderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigConfigurationRecorder), &result)
	return &result
}
