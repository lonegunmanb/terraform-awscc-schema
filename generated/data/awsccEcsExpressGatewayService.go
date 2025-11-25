package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsExpressGatewayService = `{
  "block": {
    "attributes": {
      "active_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "created_at": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "execution_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "health_check_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ingress_paths": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_type": {
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
                "nesting_mode": "list"
              }
            },
            "memory": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "network_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
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
            },
            "primary_container": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_logs_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_group": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "log_stream_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "command": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
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
                  "image": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
                  }
                },
                "nesting_mode": "single"
              }
            },
            "scaling_target": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auto_scaling_metric": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "auto_scaling_target_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_task_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_task_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_revision_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "task_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
      "cpu": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_path": {
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
      "infrastructure_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "memory": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
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
      },
      "primary_container": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_logs_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "log_stream_prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "command": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "container_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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
            "image": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "scaling_target": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_scaling_metric": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "auto_scaling_target_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_task_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_task_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "service_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status_code": {
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
      "task_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECS::ExpressGatewayService",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsExpressGatewayServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsExpressGatewayService), &result)
	return &result
}
