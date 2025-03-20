package data

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
              "type": "string"
            },
            "aws_account_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "aws_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "iam_connection_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the domain in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The identifier of the domain in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_unit_id": {
        "computed": true,
        "description": "The ID of the domain unit in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
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
        "type": "string"
      },
      "environment_user_role": {
        "computed": true,
        "description": "The role of the user in the environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the project in which the connection is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "props": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "athena_properties": {
              "computed": true,
              "description": "Athena Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "workgroup_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                                      "type": "string"
                                    },
                                    "user_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "custom_authentication_credentials": {
                                "computed": true,
                                "description": "Credential Map",
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "kms_key_arn": {
                                "computed": true,
                                "description_kind": "plain",
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
                                            "type": "string"
                                          },
                                          "redirect_uri": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
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
                                            "type": "string"
                                          },
                                          "user_managed_client_application_client_id": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
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
                                            "type": "string"
                                          },
                                          "jwt_token": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "refresh_token": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "user_managed_client_application_client_secret": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "o_auth_2_grant_type": {
                                      "computed": true,
                                      "description": "OAuth2 Grant Type",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "token_url": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "token_url_parameters_map": {
                                      "computed": true,
                                      "description": "The token URL parameters.",
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
                              "secret_arn": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "connection_properties": {
                          "computed": true,
                          "description": "Connection Properties",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "connection_type": {
                          "computed": true,
                          "description": "Glue Connection Type",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "description": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "match_criteria": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
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
                                "type": "string"
                              },
                              "security_group_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "subnet_id_list": {
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
                        "python_properties": {
                          "computed": true,
                          "description": "Property Map",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "spark_properties": {
                          "computed": true,
                          "description": "Property Map",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "validate_credentials": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "validate_for_compute_environments": {
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
            "hyper_pod_properties": {
              "computed": true,
              "description": "HyperPod Properties Input",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
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
                                "type": "string"
                              },
                              "username": {
                                "computed": true,
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
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host": {
                    "computed": true,
                    "description_kind": "plain",
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
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
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
                          "type": "string"
                        },
                        "workgroup_name": {
                          "computed": true,
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
            "spark_emr_properties": {
              "computed": true,
              "description": "Spark EMR Properties Input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "compute_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_profile_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "java_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "log_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "python_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "runtime_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "trusted_certificates_s3_uri": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "glue_connection_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "glue_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "idle_timeout": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "java_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "number_of_workers": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "python_virtual_env": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "worker_type": {
                    "computed": true,
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
      "type": {
        "computed": true,
        "description": "Connection Type",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::Connection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneConnection), &result)
	return &result
}
