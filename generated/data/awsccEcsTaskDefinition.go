package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsTaskDefinition = `{
  "block": {
    "attributes": {
      "container_definitions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "command": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "depends_on": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "container_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "disable_networking": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "dns_search_domains": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "dns_servers": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "docker_labels": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "docker_security_options": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "entry_point": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "environment": {
              "computed": true,
              "description": "The environment variables to pass to a container",
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
                "nesting_mode": "set"
              }
            },
            "environment_files": {
              "computed": true,
              "description": "The list of one or more files that contain the environment variables to pass to a container",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
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
            "essential": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "extra_hosts": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hostname": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ip_address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "firelens_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "options": {
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
                "nesting_mode": "single"
              }
            },
            "health_check": {
              "computed": true,
              "description": "The health check command and associated configuration parameters for the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "A string array representing the command that the container runs to determine if it is healthy.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "interval": {
                    "computed": true,
                    "description": "The time period in seconds between each health check execution. You may specify between 5 and 300 seconds. The default value is 30 seconds.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "retries": {
                    "computed": true,
                    "description": "The number of times to retry a failed health check before the container is considered unhealthy. You may specify between 1 and 10 retries. The default value is three retries.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "start_period": {
                    "computed": true,
                    "description": "The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries. You may specify between 0 and 300 seconds. The startPeriod is disabled by default.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "timeout": {
                    "computed": true,
                    "description": "The time period in seconds to wait for a health check to succeed before it is considered a failure. You may specify between 2 and 60 seconds. The default value is 5 seconds.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "hostname": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "image": {
              "computed": true,
              "description": "The image used to start a container. This string is passed directly to the Docker daemon.",
              "description_kind": "plain",
              "type": "string"
            },
            "interactive": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "links": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "linux_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capabilities": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "add": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "drop": {
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
                  },
                  "devices": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "host_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "permissions": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "init_process_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "max_swap": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "shared_memory_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "swappiness": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "tmpfs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "mount_options": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "size": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
            "memory": {
              "computed": true,
              "description": "The amount (in MiB) of memory to present to the container. If your container attempts to exceed the memory specified here, the container is killed.",
              "description_kind": "plain",
              "type": "number"
            },
            "memory_reservation": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "mount_points": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "read_only": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "source_volume": {
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
              "description": "The name of a container. Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed",
              "description_kind": "plain",
              "type": "string"
            },
            "port_mappings": {
              "computed": true,
              "description": "Port mappings allow containers to access ports on the host container instance to send or receive traffic.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "container_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "container_port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "privileged": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "pseudo_terminal": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "readonly_root_filesystem": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "repository_credentials": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credentials_parameter": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "resource_requirements": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
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
            "secrets": {
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
            },
            "start_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "stop_timeout": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "system_controls": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "namespace": {
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
            "ulimits": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hard_limit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "soft_limit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "user": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "volumes_from": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_only": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "source_container": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "working_directory": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ephemeral_storage": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "size_in_gi_b": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "execution_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "family": {
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
      "inference_accelerators": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "device_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "ipc_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pid_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
          "nesting_mode": "set"
        }
      },
      "proxy_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "proxy_configuration_properties": {
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
                "nesting_mode": "set"
              }
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "requires_compatibilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "runtime_platform": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu_architecture": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "operating_system_family": {
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
      "task_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon ECS task definition",
        "description_kind": "plain",
        "type": "string"
      },
      "task_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volumes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configured_at_launch": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "docker_volume_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "autoprovision": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "driver": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "driver_opts": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "labels": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "scope": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "efs_volume_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_point_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "iam": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "filesystem_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "root_directory": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transit_encryption": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transit_encryption_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "host": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_path": {
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::ECS::TaskDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsTaskDefinition), &result)
	return &result
}
