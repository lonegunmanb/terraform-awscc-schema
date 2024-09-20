package resource

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
              "optional": true,
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
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
            "overwrite": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
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
            "table_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "bucket_owner": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key": {
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
            "database_output_mode": {
              "computed": true,
              "description": "Database table name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "glue_connection_name": {
              "computed": true,
              "description": "Glue connection name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "dataset_name": {
        "computed": true,
        "description": "Dataset name",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key_arn": {
        "computed": true,
        "description": "Encryption Key Arn",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_mode": {
        "computed": true,
        "description": "Encryption mode",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "size": {
              "computed": true,
              "description": "Sample configuration size for profile jobs.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "log_subscription": {
        "computed": true,
        "description": "Log subscription",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_capacity": {
        "computed": true,
        "description": "Max capacity",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_retries": {
        "computed": true,
        "description": "Max retries",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Job name",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "bucket_owner": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
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
      "outputs": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compression_format": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "format": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
            "location": {
              "computed": true,
              "description": "S3 Output location",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
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
            "max_output_files": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "overwrite": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "partition_columns": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "regex": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "statistics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "included_statistics": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "statistic": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "dataset_statistics_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "included_statistics": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "statistic": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                  "entity_types": {
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
            "profile_columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "regex": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "project_name": {
        "computed": true,
        "description": "Project name",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "Recipe version",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "role_arn": {
        "description": "Role arn",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "timeout": {
        "computed": true,
        "description": "Timeout",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "description": "Job type",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "validation_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::DataBrew::Job.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatabrewJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatabrewJob), &result)
	return &result
}
