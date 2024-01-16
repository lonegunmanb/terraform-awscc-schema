package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGreengrassv2ComponentVersion = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "component_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "component_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inline_recipe": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lambda_function": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_dependencies": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dependency_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version_requirement": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "map"
              }
            },
            "component_lambda_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "environment_variables": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "event_sources": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "topic": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "exec_args": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "input_payload_encoding_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "linux_process_params": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_params": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "devices": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "add_group_owner": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "path": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "permission": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "memory_size_in_kb": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "mount_ro_sysfs": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "volumes": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "add_group_owner": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "destination_path": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "permission": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "source_path": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "isolation_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "max_idle_time_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_instances_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_queue_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "pinned": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "status_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "component_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "component_platforms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "component_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "lambda_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
    "description": "Data Source schema for AWS::GreengrassV2::ComponentVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGreengrassv2ComponentVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGreengrassv2ComponentVersion), &result)
	return &result
}
