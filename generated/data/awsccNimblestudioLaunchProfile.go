package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNimblestudioLaunchProfile = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ec_2_subnet_ids": {
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
      "launch_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "launch_profile_protocol_versions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "automatic_termination_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "clipboard_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ec_2_instance_types": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "max_session_length_in_minutes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_stopped_session_length_in_minutes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "session_backup": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_backups_to_retain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "session_persistence_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "session_storage": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "root": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "linux": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "windows": {
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
            "streaming_image_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "volume_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "throughput": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "studio_component_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "studio_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::NimbleStudio::LaunchProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNimblestudioLaunchProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNimblestudioLaunchProfile), &result)
	return &result
}
