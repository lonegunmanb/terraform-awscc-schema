package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppflowFlow = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Description of the flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_flow_config_list": {
        "computed": true,
        "description": "List of Destination connectors of the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_version": {
              "computed": true,
              "description": "The API version that the destination connector uses.",
              "description_kind": "plain",
              "type": "string"
            },
            "connector_profile_name": {
              "computed": true,
              "description": "Name of destination connector profile",
              "description_kind": "plain",
              "type": "string"
            },
            "connector_type": {
              "computed": true,
              "description": "Destination connector type",
              "description_kind": "plain",
              "type": "string"
            },
            "destination_connector_properties": {
              "computed": true,
              "description": "Destination connector details",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_connector": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "custom_properties": {
                          "computed": true,
                          "description": "A map for properties for custom connector.",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "entity_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "id_field_names": {
                          "computed": true,
                          "description": "List of fields used as ID when performing a write operation.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "write_operation_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "event_bridge": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lookout_metrics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "marketo": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "redshift": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "intermediate_bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_output_format_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aggregation_config": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "aggregation_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "target_file_size": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "file_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prefix_config": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "path_prefix_hierarchy": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "prefix_format": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "prefix_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "preserve_source_data_typing": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "salesforce": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_transfer_api": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "id_field_names": {
                          "computed": true,
                          "description": "List of fields used as ID when performing a write operation.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "write_operation_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sapo_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "id_field_names": {
                          "computed": true,
                          "description": "List of fields used as ID when performing a write operation.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "object_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "success_response_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "write_operation_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "snowflake": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "intermediate_bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "upsolver": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_output_format_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aggregation_config": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "aggregation_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "target_file_size": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "file_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prefix_config": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "path_prefix_hierarchy": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "prefix_format": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "prefix_type": {
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "zendesk": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "error_handling_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "fail_on_first_error": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "id_field_names": {
                          "computed": true,
                          "description": "List of fields used as ID when performing a write operation.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "write_operation_type": {
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
            }
          },
          "nesting_mode": "list"
        }
      },
      "flow_arn": {
        "computed": true,
        "description": "ARN identifier of the flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_name": {
        "computed": true,
        "description": "Name of the flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_status": {
        "computed": true,
        "description": "Flow activation status for Scheduled- and Event-triggered flows",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_arn": {
        "computed": true,
        "description": "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_catalog_config": {
        "computed": true,
        "description": "Configurations of metadata catalog of the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "glue_data_catalog": {
              "computed": true,
              "description": "Configurations of glue data catalog of the flow.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description": "A string containing the value for the tag",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "A string containing the value for the tag",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_prefix": {
                    "computed": true,
                    "description": "A string containing the value for the tag",
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
      "source_flow_config": {
        "computed": true,
        "description": "Configurations of Source connector of the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_version": {
              "computed": true,
              "description": "The API version that the destination connector uses.",
              "description_kind": "plain",
              "type": "string"
            },
            "connector_profile_name": {
              "computed": true,
              "description": "Name of source connector profile",
              "description_kind": "plain",
              "type": "string"
            },
            "connector_type": {
              "computed": true,
              "description": "Type of source connector",
              "description_kind": "plain",
              "type": "string"
            },
            "incremental_pull_config": {
              "computed": true,
              "description": "Configuration for scheduled incremental data pull",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "datetime_type_field_name": {
                    "computed": true,
                    "description": "Name of the datetime/timestamp data type field to be used for importing incremental records from the source",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source_connector_properties": {
              "computed": true,
              "description": "Source connector details required to query a connector",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "amplitude": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "custom_connector": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "custom_properties": {
                          "computed": true,
                          "description": "A map for properties for custom connector.",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "data_transfer_api": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
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
                            "nesting_mode": "single"
                          }
                        },
                        "entity_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "datadog": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "dynatrace": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "google_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "infor_nexus": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "marketo": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "pardot": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_input_format_config": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_input_file_type": {
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
                  "salesforce": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_transfer_api": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "enable_dynamic_field_update": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "include_deleted_records": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "sapo_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "pagination_config": {
                          "computed": true,
                          "description": "SAP Source connector page size",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_page_size": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "parallelism_config": {
                          "computed": true,
                          "description": "SAP Source connector parallelism factor",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_parallelism": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "service_now": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "singular": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "slack": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "trendmicro": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "veeva": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "document_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "include_all_versions": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "include_renditions": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "include_source_files": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "object": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "zendesk": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "object": {
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "List of Tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tasks": {
        "computed": true,
        "description": "List of tasks for the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connector_operator": {
              "computed": true,
              "description": "Operation to be performed on provided source fields",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "amplitude": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "custom_connector": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "datadog": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dynatrace": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "google_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "infor_nexus": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "marketo": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "pardot": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "salesforce": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sapo_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "service_now": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "singular": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "slack": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "trendmicro": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "veeva": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "zendesk": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "destination_field": {
              "computed": true,
              "description": "A field value on which source field should be validated",
              "description_kind": "plain",
              "type": "string"
            },
            "source_fields": {
              "computed": true,
              "description": "Source fields on which particular task will be applied",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "task_properties": {
              "computed": true,
              "description": "A Map used to store task related info",
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
            "task_type": {
              "computed": true,
              "description": "Type of task",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "trigger_config": {
        "computed": true,
        "description": "Trigger settings of the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "trigger_properties": {
              "computed": true,
              "description": "Details required based on the type of trigger",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_pull_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "first_execution_from": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "flow_error_deactivation_threshold": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "schedule_end_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "schedule_expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "schedule_offset": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "schedule_start_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "time_zone": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "trigger_type": {
              "computed": true,
              "description": "Trigger type of the flow",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AppFlow::Flow",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppflowFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppflowFlow), &result)
	return &result
}
