package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationsignalsInstrumentationConfig = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the instrumentation configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "attribute_filters": {
        "computed": true,
        "description": "Client-side filters that target specific instances. Each object is AND-matched on keys, multiple objects are OR-matched.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "map",
            "string"
          ]
        ]
      },
      "capture_configuration": {
        "computed": true,
        "description": "Specifies what to capture when the instrumentation point is hit.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code_capture": {
              "computed": true,
              "description": "Defines what data to capture for code-level instrumentation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capture_arguments": {
                    "computed": true,
                    "description": "The function arguments to capture.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "capture_limits": {
                    "computed": true,
                    "description": "Safety limits that bound what is captured.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_collection_depth": {
                          "computed": true,
                          "description": "Maximum nesting depth to traverse inside collections.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_collection_width": {
                          "computed": true,
                          "description": "Maximum number of items to capture from any collection.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_fields_per_object": {
                          "computed": true,
                          "description": "Maximum number of fields to capture for any object.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_hits": {
                          "computed": true,
                          "description": "Maximum number of times the instrumentation point can be hit before disabled.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_object_depth": {
                          "computed": true,
                          "description": "Maximum depth for nested object traversal.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_stack_frames": {
                          "computed": true,
                          "description": "Maximum number of stack frames to capture.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_stack_trace_size": {
                          "computed": true,
                          "description": "Maximum total size in bytes of a captured stack trace.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_string_length": {
                          "computed": true,
                          "description": "Maximum length of captured string values in characters.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "capture_locals": {
                    "computed": true,
                    "description": "The local variables to capture by name.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "capture_return": {
                    "computed": true,
                    "description": "Whether to capture the return value. Defaults to false.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "capture_stack_trace": {
                    "computed": true,
                    "description": "Whether to capture a stack trace. Defaults to true.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The server-generated creation timestamp for this instrumentation configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional short description that explains the purpose of this instrumentation.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description": "The environment that the service is running in.",
        "description_kind": "plain",
        "type": "string"
      },
      "expires_at": {
        "computed": true,
        "description": "The timestamp after which this configuration is no longer served. For BREAKPOINT only; defaults to 24 hours.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instrumentation_type": {
        "computed": true,
        "description": "Type of instrumentation: BREAKPOINT (temporary, expires after 24 hours) or PROBE (permanent, persists until deleted).",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "The location where instrumentation should be applied.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code_location": {
              "computed": true,
              "description": "Identifies a code location to instrument.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "class_name": {
                    "computed": true,
                    "description": "The class or type name that contains the method.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "code_unit": {
                    "computed": true,
                    "description": "The package, module, or namespace that contains the target code.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "file_path": {
                    "computed": true,
                    "description": "The source file path relative to the project or source root.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "language": {
                    "computed": true,
                    "description": "The programming language for this instrumentation point.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "line_number": {
                    "computed": true,
                    "description": "The line number to instrument.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "method_name": {
                    "computed": true,
                    "description": "The method or function name to instrument.",
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
      "location_hash": {
        "computed": true,
        "description": "A stable hash computed from the location that uniquely identifies this instrumentation point.",
        "description_kind": "plain",
        "type": "string"
      },
      "service": {
        "computed": true,
        "description": "The name of the service to instrument. This should match the service.name resource attribute reported by the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "signal_type": {
        "computed": true,
        "description": "The telemetry signal type to emit for this instrumentation. The supported value is SNAPSHOT.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An optional list of key-value pairs to associate with the instrumentation configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ApplicationSignals::InstrumentationConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApplicationsignalsInstrumentationConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationsignalsInstrumentationConfig), &result)
	return &result
}
