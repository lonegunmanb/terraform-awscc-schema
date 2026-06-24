package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaMicrovmImage = `{
  "block": {
    "attributes": {
      "additional_os_capabilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "base_image_arn": {
        "computed": true,
        "description": "ARN of the base MicroVM image.",
        "description_kind": "plain",
        "type": "string"
      },
      "base_image_version": {
        "computed": true,
        "description": "Specific version of the base MicroVM image to use.",
        "description_kind": "plain",
        "type": "string"
      },
      "build_role_arn": {
        "computed": true,
        "description": "ARN of the IAM build role.",
        "description_kind": "plain",
        "type": "string"
      },
      "code_artifact": {
        "computed": true,
        "description": "Code artifact for the active MicroVM image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "uri": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cpu_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the MicroVM image was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Human-readable description of the MicroVM image and its purpose.",
        "description_kind": "plain",
        "type": "string"
      },
      "egress_network_connectors": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "environment_variables": {
        "computed": true,
        "description": "Environment variables to set in the container during the snapshot build.",
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
          "nesting_mode": "set"
        }
      },
      "hooks": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "microvm_hooks": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resume": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resume_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "run": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "run_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "suspend": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "suspend_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "terminate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "terminate_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "microvm_image_hooks": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ready": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ready_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "validate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "validate_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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
      "image_arn": {
        "computed": true,
        "description": "ARN of the MicroVM image.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_active_image_version": {
        "computed": true,
        "description": "The latest active version of the MicroVM image.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_failed_image_version": {
        "computed": true,
        "description": "The latest failed version of the MicroVM image.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging": {
        "computed": true,
        "description": "Configuration for MicroVM image logging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "log_stream": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "disabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "Unique name for the MicroVM image within the account.",
        "description_kind": "plain",
        "type": "string"
      },
      "resources": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "minimum_memory_in_mi_b": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "set"
        }
      },
      "state": {
        "computed": true,
        "description": "Current state of the MicroVM image.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs to associate with the MicroVM image for organization and management.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "Timestamp when the MicroVM image was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lambda::MicrovmImage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaMicrovmImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaMicrovmImage), &result)
	return &result
}
