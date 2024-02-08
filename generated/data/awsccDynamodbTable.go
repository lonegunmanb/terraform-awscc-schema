package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDynamodbTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attribute_definitions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "attribute_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "billing_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "contributor_insights_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "global_secondary_indexes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "contributor_insights_specification": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "index_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_schema": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "projection": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "non_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "projection_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "provisioned_throughput": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_capacity_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "write_capacity_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "import_source_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_compression_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "input_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "input_format_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "csv": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delimiter": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "header_list": {
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
            "s3_bucket_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_key_prefix": {
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
      "key_schema": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kinesis_stream_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "approximate_creation_date_time_precision": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stream_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "local_secondary_indexes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "index_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_schema": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "projection": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "non_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "projection_type": {
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
      },
      "point_in_time_recovery_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "point_in_time_recovery_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "provisioned_throughput": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "read_capacity_units": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "write_capacity_units": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sse_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_master_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sse_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "sse_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "stream_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stream_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "stream_view_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "table_class": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
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
      },
      "time_to_live_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DynamoDB::Table",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDynamodbTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDynamodbTable), &result)
	return &result
}
