package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneConnection = `{
  "block": {
    "attributes": {
      "aws_location": {
        "computed": true,
        "description": "AWS Location of project",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_role": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aws_account_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aws_region": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "iam_connection_id": {
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
      "connection_id": {
        "computed": true,
        "description": "The ID of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the domain in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "description": "The identifier of the domain in which the connection is created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_unit_id": {
        "computed": true,
        "description": "The ID of the domain unit in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_trusted_identity_propagation": {
        "computed": true,
        "description": "Specifies whether the trusted identity propagation is enabled",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of the environment in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_identifier": {
        "computed": true,
        "description": "The identifier of the environment in which the connection is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_user_role": {
        "computed": true,
        "description": "The role of the user in the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the project in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_identifier": {
        "computed": true,
        "description": "The identifier of the project in which the connection should be created. If ",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "props": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_q_properties": {
              "computed": true,
              "description": "Amazon Q properties of the connection.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_mode": {
                    "computed": true,
                    "description": "The authentication mode of the connection's AmazonQ properties",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "is_enabled": {
                    "computed": true,
                    "description": "Specifies whether Amazon Q is enabled for the connection",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "profile_arn": {
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
            "athena_properties": {
              "computed": true,
              "description": "Athena Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "workgroup_name": {
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
            "glue_properties": {
              "computed": true,
              "description": "Glue Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "glue_connection_input": {
                    "computed": true,
                    "description": "Glue Connection Input",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "athena_properties": {
                          "computed": true,
                          "description": "Property Map",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "authentication_configuration": {
                          "computed": true,
                          "description": "Authentication Configuration Input",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "authentication_type": {
                                "computed": true,
                                "description": "Authentication Type",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "basic_authentication_credentials": {
                                "computed": true,
                                "description": "Basic Authentication Credentials",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "password": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "user_name": {
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
                              "custom_authentication_credentials": {
                                "computed": true,
                                "description": "Credential Map",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "kms_key_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "o_auth_2_properties": {
                                "computed": true,
                                "description": "OAuth2 Properties",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "authorization_code_properties": {
                                      "computed": true,
                                      "description": "Authorization Code Properties",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "authorization_code": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "redirect_uri": {
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
                                    "o_auth_2_client_application": {
                                      "computed": true,
                                      "description": "OAuth2 Client Application",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "aws_managed_client_application_reference": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "user_managed_client_application_client_id": {
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
                                    "o_auth_2_credentials": {
                                      "computed": true,
                                      "description": "Glue OAuth2 Credentials",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "access_token": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "jwt_token": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "refresh_token": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "user_managed_client_application_client_secret": {
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
                                    "o_auth_2_grant_type": {
                                      "computed": true,
                                      "description": "OAuth2 Grant Type",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "token_url": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "token_url_parameters_map": {
                                      "computed": true,
                                      "description": "The token URL parameters.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "secret_arn": {
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
                        "connection_properties": {
                          "computed": true,
                          "description": "Connection Properties",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "connection_type": {
                          "computed": true,
                          "description": "Glue Connection Type",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "description": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "match_criteria": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "physical_connection_requirements": {
                          "computed": true,
                          "description": "Physical Connection Requirements",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "availability_zone": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "security_group_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "subnet_id_list": {
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
                        "python_properties": {
                          "computed": true,
                          "description": "Property Map",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "spark_properties": {
                          "computed": true,
                          "description": "Property Map",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "validate_credentials": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "validate_for_compute_environments": {
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "hyper_pod_properties": {
              "computed": true,
              "description": "HyperPod Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_name": {
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
            "iam_properties": {
              "computed": true,
              "description": "IAM Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "glue_lineage_sync_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "mlflow_properties": {
              "computed": true,
              "description": "MLflow Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "tracking_server_arn": {
                    "computed": true,
                    "description": "The ARN of the MLflow tracking server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "redshift_properties": {
              "computed": true,
              "description": "Redshift Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credentials": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "username_password": {
                          "computed": true,
                          "description": "The username and password to be used for authentication.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "password": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "username": {
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
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "host": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "lineage_sync": {
                    "computed": true,
                    "description": "Redshift Lineage Sync Configuration Input",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "schedule": {
                          "computed": true,
                          "description": "Lineage Sync Schedule",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "schedule": {
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
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "storage": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cluster_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "workgroup_name": {
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
            "s3_properties": {
              "computed": true,
              "description": "S3 Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_access_grant_location_id": {
                    "computed": true,
                    "description": "The Amazon S3 Access Grant location ID that's part of the Amazon S3 properties of a connection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_uri": {
                    "computed": true,
                    "description": "The Amazon S3 URI that's part of the Amazon S3 properties of a connection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "spark_emr_properties": {
              "computed": true,
              "description": "Spark EMR Properties Input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "compute_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_profile_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "java_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "managed_endpoint_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "python_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "runtime_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "trusted_certificates_s3_uri": {
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
            "spark_glue_properties": {
              "computed": true,
              "description": "Spark Glue Properties Input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "additional_args": {
                    "computed": true,
                    "description": "Spark Glue Args.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "connection": {
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
                  "glue_connection_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "glue_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "idle_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "java_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "number_of_workers": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "python_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "worker_type": {
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
      "scope": {
        "computed": true,
        "description": "The scope of the connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Connection Type",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Connections enables users to connect their DataZone resources (domains, projects, and environments) to external resources/services (data, compute, etc)",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneConnection), &result)
	return &result
}
