package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsService = `{
  "block": {
    "attributes": {
      "capacity_provider_strategy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "capacity_provider": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "cluster": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "rollback": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "deployment_circuit_breaker": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "rollback": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maximum_percent": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "minimum_healthy_percent": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "deployment_controller": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "desired_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "enable_ecs_managed_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_execute_command": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "health_check_grace_period_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "launch_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "load_balancers": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "container_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "load_balancer_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target_group_arn": {
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
      "network_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "awsvpc_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "assign_public_ip": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "security_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "placement_constraints": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expression": {
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
      "placement_strategies": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field": {
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
      "platform_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "propagate_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scheduling_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_connect_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "log_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_driver": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "secret_options": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value_from": {
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
            "namespace": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "services": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_aliases": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dns_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "discovery_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ingress_port_override": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "port_name": {
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
      "service_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_registries": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "container_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "registry_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
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
      "task_definition": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "managed_ebs_volume": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encrypted": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "filesystem_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "iops": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "size_in_gi_b": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "snapshot_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tag_specifications": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "propagate_tags": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resource_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tags": {
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
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "throughput": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "volume_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ECS::Service",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsService), &result)
	return &result
}
