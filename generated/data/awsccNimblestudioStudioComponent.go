package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNimblestudioStudioComponent = `{
  "block": {
    "attributes": {
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_directory_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "computer_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
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
                  "directory_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "organizational_unit_distinguished_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "compute_farm_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "active_directory_user": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "license_service_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "shared_file_system_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "endpoint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "file_system_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "linux_mount_point": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "share_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "windows_mount_drive": {
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
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ec_2_security_group_ids": {
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
      "initialization_scripts": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_profile_protocol_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "platform": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "run_context": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "script": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "script_parameters": {
        "computed": true,
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
      "studio_component_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "studio_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subtype": {
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
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NimbleStudio::StudioComponent",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNimblestudioStudioComponentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNimblestudioStudioComponent), &result)
	return &result
}
