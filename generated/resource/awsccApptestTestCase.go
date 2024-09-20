package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApptestTestCase = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "steps": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "compare_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "input": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "file": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "file_metadata": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "data_sets": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "ccsid": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "format": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "length": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "name": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "database_cdc": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "source_metadata": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "capture_tool": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "type": {
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
                                                "target_metadata": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "capture_tool": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "type": {
                                                        "computed": true,
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  },
                                                  "optional": true
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "source_location": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "target_location": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "output": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "file": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "file_location": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "mainframe_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "batch": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "batch_job_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "batch_job_parameters": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "export_data_set_names": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "tn_3270": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "export_data_set_names": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "script": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "script_location": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "type": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "dms_task_arn": {
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
                        "resource": {
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
                  "resource_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cloudformation_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "action_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "resource": {
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
                        "m2_managed_application_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "action_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "properties": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "force_stop": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "import_data_set_location": {
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
                              "resource": {
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
                        "m2_non_managed_application_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "action_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "resource": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "test_case_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "test_case_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "test_case_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Represents a Test Case that can be captured and executed",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApptestTestCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApptestTestCase), &result)
	return &result
}
