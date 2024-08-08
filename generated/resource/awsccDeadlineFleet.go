package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineFleet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capabilities": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amounts": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
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
            "attributes": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "values": {
                    "computed": true,
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
      "configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customer_managed": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mode": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "storage_profile_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "worker_capabilities": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "accelerator_count": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "accelerator_total_memory_mi_b": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "accelerator_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "cpu_architecture_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "custom_amounts": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "custom_attributes": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "values": {
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "memory_mi_b": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "os_family": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "v_cpu_count": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_managed_ec_2": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_capabilities": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "allowed_instance_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "cpu_architecture_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "custom_amounts": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "custom_attributes": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "values": {
                                "description_kind": "plain",
                                "required": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "excluded_instance_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "memory_mi_b": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "os_family": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "root_ebs_volume": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "iops": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "size_gi_b": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "throughput_mi_b": {
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
                        "v_cpu_count": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "min": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "instance_market_options": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
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
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "farm_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fleet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_worker_count": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_worker_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "worker_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Definition of AWS::Deadline::Fleet Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineFleet), &result)
	return &result
}
