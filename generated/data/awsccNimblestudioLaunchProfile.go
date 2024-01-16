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
        "description": "\u003cp\u003eThe description.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "ec_2_subnet_ids": {
        "computed": true,
        "description": "\u003cp\u003eSpecifies the IDs of the EC2 subnets where streaming sessions will be accessible from.\n            These subnets must support the specified instance types. \u003c/p\u003e",
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
        "description": "\u003cp\u003eThe version number of the protocol that is used by the launch profile. The only valid\n            version is \"2021-03-31\".\u003c/p\u003e",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "\u003cp\u003eThe name for the launch profile.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_configuration": {
        "computed": true,
        "description": "\u003cp\u003eA configuration for a streaming session.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe EC2 instance types that users can select from when launching a streaming session\n            with this launch profile.\u003c/p\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "max_session_length_in_minutes": {
              "computed": true,
              "description": "\u003cp\u003eThe length of time, in minutes, that a streaming session can be active before it is\n            stopped or terminated. After this point, Nimble Studio automatically terminates or\n            stops the session. The default length of time is 690 minutes, and the maximum length of\n            time is 30 days.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "max_stopped_session_length_in_minutes": {
              "computed": true,
              "description": "\u003cp\u003eInteger that determines if you can start and stop your sessions and how long a session\n            can stay in the \u003ccode\u003eSTOPPED\u003c/code\u003e state. The default value is 0. The maximum value is\n            5760.\u003c/p\u003e\n         \u003cp\u003eThis field is allowed only when \u003ccode\u003esessionPersistenceMode\u003c/code\u003e is\n                \u003ccode\u003eACTIVATED\u003c/code\u003e and \u003ccode\u003eautomaticTerminationMode\u003c/code\u003e is\n                \u003ccode\u003eACTIVATED\u003c/code\u003e.\u003c/p\u003e\n         \u003cp\u003eIf the value is set to 0, your sessions can?t be \u003ccode\u003eSTOPPED\u003c/code\u003e. If you then\n            call \u003ccode\u003eStopStreamingSession\u003c/code\u003e, the session fails. If the time that a session\n            stays in the \u003ccode\u003eREADY\u003c/code\u003e state exceeds the \u003ccode\u003emaxSessionLengthInMinutes\u003c/code\u003e\n            value, the session will automatically be terminated (instead of\n            \u003ccode\u003eSTOPPED\u003c/code\u003e).\u003c/p\u003e\n         \u003cp\u003eIf the value is set to a positive number, the session can be stopped. You can call\n                \u003ccode\u003eStopStreamingSession\u003c/code\u003e to stop sessions in the \u003ccode\u003eREADY\u003c/code\u003e state.\n            If the time that a session stays in the \u003ccode\u003eREADY\u003c/code\u003e state exceeds the\n                \u003ccode\u003emaxSessionLengthInMinutes\u003c/code\u003e value, the session will automatically be\n            stopped (instead of terminated).\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "session_backup": {
              "computed": true,
              "description": "\u003cp\u003eConfigures how streaming sessions are backed up when launched from this launch\n            profile.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_backups_to_retain": {
                    "computed": true,
                    "description": "\u003cp\u003eThe maximum number of backups that each streaming session created from this launch\n            profile can have.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe configuration for a streaming session?s upload storage.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mode": {
                    "computed": true,
                    "description": "\u003cp\u003eAllows artists to upload files to their workstations. The only valid option is\n                \u003ccode\u003eUPLOAD\u003c/code\u003e.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "root": {
                    "computed": true,
                    "description": "\u003cp\u003eThe upload storage root location (folder) on streaming workstations where files are\n            uploaded.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "linux": {
                          "computed": true,
                          "description": "\u003cp\u003eThe folder path in Linux workstations where files are uploaded.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "windows": {
                          "computed": true,
                          "description": "\u003cp\u003eThe folder path in Windows workstations where files are uploaded.\u003c/p\u003e",
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
              "description": "\u003cp\u003eThe streaming images that users can select from when launching a streaming session\n            with this launch profile.\u003c/p\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "volume_configuration": {
              "computed": true,
              "description": "\u003cp\u003eCustom volume configuration for the root volumes that are attached to streaming\n            sessions.\u003c/p\u003e\n         \u003cp\u003eThis parameter is only allowed when \u003ccode\u003esessionPersistenceMode\u003c/code\u003e is\n                \u003ccode\u003eACTIVATED\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iops": {
                    "computed": true,
                    "description": "\u003cp\u003eThe number of I/O operations per second for the root volume that is attached to\n            streaming session.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "size": {
                    "computed": true,
                    "description": "\u003cp\u003eThe size of the root volume that is attached to the streaming session. The root volume\n            size is measured in GiBs.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "throughput": {
                    "computed": true,
                    "description": "\u003cp\u003eThe throughput to provision for the root volume that is attached to the streaming\n            session. The throughput is measured in MiB/s.\u003c/p\u003e",
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
        "description": "\u003cp\u003eUnique identifiers for a collection of studio components that can be used with this\n            launch profile.\u003c/p\u003e",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "studio_id": {
        "computed": true,
        "description": "\u003cp\u003eThe studio ID. \u003c/p\u003e",
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
