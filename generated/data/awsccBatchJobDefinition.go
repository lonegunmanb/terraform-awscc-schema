package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBatchJobDefinition = `{
  "block": {
    "attributes": {
      "consumable_resource_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "consumable_resource_list": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "consumable_resource": {
                    "computed": true,
                    "description": "The ARN of the consumable resource the job definition should consume.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "quantity": {
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
      "container_properties": {
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
            "enable_execute_command": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "environment": {
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
            "fargate_platform_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "platform_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "job_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "linux_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
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
                            "list",
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
            "network_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "assign_public_ip": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "privileged": {
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
            "vcpus": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "volumes": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
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
                        "file_system_id": {
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
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "ecs_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "task_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "containers": {
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
                        "environment": {
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
                        "essential": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
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
                        "image": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "linux_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
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
                                        "list",
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
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "privileged": {
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
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "enable_execute_command": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
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
                  "ipc_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "assign_public_ip": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "pid_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "platform_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
                              "file_system_id": {
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
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "eks_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pod_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "containers": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "args": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "command": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "env": {
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
                        "image": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_pull_policy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resources": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "limits": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "requests": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "security_context": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allow_privilege_escalation": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "privileged": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "read_only_root_filesystem": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "run_as_group": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "run_as_non_root": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "run_as_user": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "volume_mounts": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "mount_path": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "read_only": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "sub_path": {
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
                  "dns_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_network": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "image_pull_secrets": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "init_containers": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "args": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "command": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "env": {
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
                        "image": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_pull_policy": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resources": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "limits": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "requests": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "security_context": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allow_privilege_escalation": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "privileged": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "read_only_root_filesystem": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "run_as_group": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "run_as_non_root": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "run_as_user": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "volume_mounts": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "mount_path": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "read_only": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "sub_path": {
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
                  "metadata": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "annotations": {
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
                        "namespace": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "service_account_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "share_process_namespace": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "volumes": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "empty_dir": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "medium": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "size_limit": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "host_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "path": {
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
                        },
                        "persistent_volume_claim": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "claim_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "read_only": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "secret": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "optional": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "secret_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
      "job_definition_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "job_definition_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "node_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "main_node": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "node_range_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "consumable_resource_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "consumable_resource_list": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "consumable_resource": {
                                "computed": true,
                                "description": "The ARN of the consumable resource the job definition should consume.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "quantity": {
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
                  "container": {
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
                        "enable_execute_command": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "environment": {
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
                        "image": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "instance_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "job_role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "linux_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
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
                                        "list",
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
                        "privileged": {
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
                        "vcpus": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "volumes": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
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
                                    "file_system_id": {
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
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "ecs_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "task_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "containers": {
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
                                    "environment": {
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
                                    "essential": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
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
                                    "image": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "linux_parameters": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
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
                                                    "list",
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
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "privileged": {
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
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "enable_execute_command": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "execution_role_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "ipc_mode": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "pid_mode": {
                                "computed": true,
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
                                          "file_system_id": {
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
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "eks_properties": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "pod_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "containers": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "args": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "command": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "env": {
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
                                    "image": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "image_pull_policy": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "resources": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "limits": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          },
                                          "requests": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "security_context": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_privilege_escalation": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "privileged": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "read_only_root_filesystem": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "run_as_group": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "run_as_non_root": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "run_as_user": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "volume_mounts": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mount_path": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "read_only": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "sub_path": {
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
                              "dns_policy": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "host_network": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "image_pull_secrets": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "init_containers": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "args": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "command": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "env": {
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
                                    "image": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "image_pull_policy": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "resources": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "limits": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          },
                                          "requests": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": [
                                              "map",
                                              "string"
                                            ]
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "security_context": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "allow_privilege_escalation": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "privileged": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "read_only_root_filesystem": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "run_as_group": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "run_as_non_root": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "run_as_user": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "volume_mounts": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mount_path": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "read_only": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "sub_path": {
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
                              "metadata": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "annotations": {
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
                                    "namespace": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "service_account_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "share_process_namespace": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "volumes": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "empty_dir": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "medium": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "size_limit": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "host_path": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "path": {
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
                                    },
                                    "persistent_volume_claim": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "claim_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "read_only": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "secret": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "optional": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "bool"
                                          },
                                          "secret_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "instance_types": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "target_nodes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "num_nodes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "platform_capabilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "propagate_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "retry_strategy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attempts": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "evaluate_on_exit": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "on_exit_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "on_reason": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "on_status_reason": {
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
      "scheduling_priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "timeout": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attempt_duration_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Batch::JobDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBatchJobDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBatchJobDefinition), &result)
	return &result
}
