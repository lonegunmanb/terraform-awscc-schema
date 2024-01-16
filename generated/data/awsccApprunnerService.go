package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApprunnerService = `{
  "block": {
    "attributes": {
      "auto_scaling_configuration_arn": {
        "computed": true,
        "description": "Autoscaling configuration ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "Encryption configuration (KMS key)",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key": {
              "computed": true,
              "description": "The KMS Key",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "health_check_configuration": {
        "computed": true,
        "description": "Health check configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "healthy_threshold": {
              "computed": true,
              "description": "Health check Healthy Threshold",
              "description_kind": "plain",
              "type": "number"
            },
            "interval": {
              "computed": true,
              "description": "Health check Interval",
              "description_kind": "plain",
              "type": "number"
            },
            "path": {
              "computed": true,
              "description": "Health check Path",
              "description_kind": "plain",
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "Health Check Protocol",
              "description_kind": "plain",
              "type": "string"
            },
            "timeout": {
              "computed": true,
              "description": "Health check Timeout",
              "description_kind": "plain",
              "type": "number"
            },
            "unhealthy_threshold": {
              "computed": true,
              "description": "Health check Unhealthy Threshold",
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
      "instance_configuration": {
        "computed": true,
        "description": "Instance Configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu": {
              "computed": true,
              "description": "CPU",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_role_arn": {
              "computed": true,
              "description": "Instance Role Arn",
              "description_kind": "plain",
              "type": "string"
            },
            "memory": {
              "computed": true,
              "description": "Memory",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "network_configuration": {
        "computed": true,
        "description": "Network configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "egress_configuration": {
              "computed": true,
              "description": "Network egress configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "egress_type": {
                    "computed": true,
                    "description": "Network egress type.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc_connector_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the App Runner VpcConnector.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ingress_configuration": {
              "computed": true,
              "description": "Network ingress configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "is_publicly_accessible": {
                    "computed": true,
                    "description": "It's set to true if the Apprunner service is publicly accessible. It's set to false otherwise.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ip_address_type": {
              "computed": true,
              "description": "App Runner service endpoint IP address type",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "observability_configuration": {
        "computed": true,
        "description": "Service observability configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "observability_configuration_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the App Runner ObservabilityConfiguration.",
              "description_kind": "plain",
              "type": "string"
            },
            "observability_enabled": {
              "computed": true,
              "description": "Observability enabled",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "service_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AppRunner Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description": "The AppRunner Service Id",
        "description_kind": "plain",
        "type": "string"
      },
      "service_name": {
        "computed": true,
        "description": "The AppRunner Service Name.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_url": {
        "computed": true,
        "description": "The Service Url of the AppRunner Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_configuration": {
        "computed": true,
        "description": "Source Code configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_configuration": {
              "computed": true,
              "description": "Authentication Configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_role_arn": {
                    "computed": true,
                    "description": "Access Role Arn",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "connection_arn": {
                    "computed": true,
                    "description": "Connection Arn",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "auto_deployments_enabled": {
              "computed": true,
              "description": "Auto Deployment enabled",
              "description_kind": "plain",
              "type": "bool"
            },
            "code_repository": {
              "computed": true,
              "description": "Source Code Repository",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code_configuration": {
                    "computed": true,
                    "description": "Code Configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "code_configuration_values": {
                          "computed": true,
                          "description": "Code Configuration Values",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "build_command": {
                                "computed": true,
                                "description": "Build Command",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "Port",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "runtime": {
                                "computed": true,
                                "description": "Runtime",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "runtime_environment_secrets": {
                                "computed": true,
                                "description": "The secrets and parameters that get referenced by your service as environment variables",
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
                              "runtime_environment_variables": {
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
                              "start_command": {
                                "computed": true,
                                "description": "Start Command",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "configuration_source": {
                          "computed": true,
                          "description": "Configuration Source",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "repository_url": {
                    "computed": true,
                    "description": "Repository Url",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_code_version": {
                    "computed": true,
                    "description": "Source Code Version",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "Source Code Version Type",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "Source Code Version Value",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "source_directory": {
                    "computed": true,
                    "description": "Source Directory",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image_repository": {
              "computed": true,
              "description": "Image Repository",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "image_configuration": {
                    "computed": true,
                    "description": "Image Configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "port": {
                          "computed": true,
                          "description": "Port",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "runtime_environment_secrets": {
                          "computed": true,
                          "description": "The secrets and parameters that get referenced by your service as environment variables",
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
                        "runtime_environment_variables": {
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
                        "start_command": {
                          "computed": true,
                          "description": "Start Command",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "image_identifier": {
                    "computed": true,
                    "description": "Image Identifier",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "image_repository_type": {
                    "computed": true,
                    "description": "Image Repository Type",
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
      "status": {
        "computed": true,
        "description": "AppRunner Service status.",
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
    "description": "Data Source schema for AWS::AppRunner::Service",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApprunnerServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApprunnerService), &result)
	return &result
}
