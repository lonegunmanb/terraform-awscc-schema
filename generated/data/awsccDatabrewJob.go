package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatabrewJob = `{
  "block": {
    "attributes": {
      "data_catalog_outputs": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "database_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temp_directory": {
                    "computed": true,
                    "description": "S3 Output location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
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
            "overwrite": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "s3_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "location": {
                    "computed": true,
                    "description": "S3 Output location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
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
            "table_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "database_outputs": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "database_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "table_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "temp_directory": {
                    "computed": true,
                    "description": "S3 Output location",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
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
            "database_output_mode": {
              "computed": true,
              "description": "Database table name",
              "description_kind": "plain",
              "type": "string"
            },
            "glue_connection_name": {
              "computed": true,
              "description": "Glue connection name",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "dataset_name": {
        "computed": true,
        "description": "Dataset name",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "Encryption Key Arn",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_mode": {
        "computed": true,
        "description": "Encryption mode",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_sample": {
        "computed": true,
        "description": "Job Sample",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Sample configuration mode for profile jobs.",
              "description_kind": "plain",
              "type": "string"
            },
            "size": {
              "computed": true,
              "description": "Sample configuration size for profile jobs.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "log_subscription": {
        "computed": true,
        "description": "Log subscription",
        "description_kind": "plain",
        "type": "string"
      },
      "max_capacity": {
        "computed": true,
        "description": "Max capacity",
        "description_kind": "plain",
        "type": "number"
      },
      "max_retries": {
        "computed": true,
        "description": "Max retries",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Job name",
        "description_kind": "plain",
        "type": "string"
      },
      "output_location": {
        "computed": true,
        "description": "Output location",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "bucket_owner": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "outputs": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compression_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "format_options": {
              "computed": true,
              "description": "Format options for job Output",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "csv": {
                    "computed": true,
                    "description": "Output Csv options",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delimiter": {
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
            "location": {
              "computed": true,
              "description": "S3 Output location",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "max_output_files": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "overwrite": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "partition_columns": {
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
      "profile_configuration": {
        "computed": true,
        "description": "Profile Job configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column_statistics_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "selectors": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "regex": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "statistics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "included_statistics": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "overrides": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "parameters": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "statistic": {
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
                  }
                },
                "nesting_mode": "list"
              }
            },
            "dataset_statistics_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "included_statistics": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "statistic": {
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
            "entity_detector_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_statistics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "statistics": {
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
                  "entity_types": {
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
            "profile_columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "regex": {
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
      "project_name": {
        "computed": true,
        "description": "Project name",
        "description_kind": "plain",
        "type": "string"
      },
      "recipe": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "Recipe name",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "Recipe version",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description": "Role arn",
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
      },
      "timeout": {
        "computed": true,
        "description": "Timeout",
        "description_kind": "plain",
        "type": "number"
      },
      "type": {
        "computed": true,
        "description": "Job type",
        "description_kind": "plain",
        "type": "string"
      },
      "validation_configurations": {
        "computed": true,
        "description": "Data quality rules configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ruleset_arn": {
              "computed": true,
              "description": "Arn of the Ruleset",
              "description_kind": "plain",
              "type": "string"
            },
            "validation_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::DataBrew::Job",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatabrewJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewJob), &result)
	return &result
}
