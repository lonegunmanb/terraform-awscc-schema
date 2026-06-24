package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaMicrovmImage = `{
  "block": {
    "attributes": {
      "additional_os_capabilities": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "base_image_arn": {
        "description": "ARN of the base MicroVM image.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "base_image_version": {
        "description": "Specific version of the base MicroVM image to use.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "build_role_arn": {
        "description": "ARN of the IAM build role.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "code_artifact": {
        "description": "Code artifact for the active MicroVM image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "cpu_configurations": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architecture": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the MicroVM image was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Human-readable description of the MicroVM image and its purpose.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "egress_network_connectors": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "environment_variables": {
        "description": "Environment variables to set in the container during the snapshot build.",
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
          "nesting_mode": "set"
        },
        "required": true
      },
      "hooks": {
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
                    "optional": true,
                    "type": "string"
                  },
                  "resume_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "run": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "run_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "suspend": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "suspend_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "terminate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "terminate_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "microvm_image_hooks": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ready": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ready_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "validate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "validate_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
                    "optional": true,
                    "type": "string"
                  },
                  "log_stream": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "disabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "name": {
        "description": "Unique name for the MicroVM image within the account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resources": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "minimum_memory_in_mi_b": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "Timestamp when the MicroVM image was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::MicrovmImage",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaMicrovmImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaMicrovmImage), &result)
	return &result
}
