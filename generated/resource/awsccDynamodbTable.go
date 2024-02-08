package resource

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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "attribute_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "billing_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "contributor_insights_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "index_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_schema": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "projection": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "non_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "projection_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "provisioned_throughput": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_capacity_units": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "write_capacity_units": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "input_format": {
              "description_kind": "plain",
              "required": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "header_list": {
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
            "s3_bucket_source": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_bucket_owner": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3_key_prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "key_schema": {
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "stream_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "local_secondary_indexes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "index_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_schema": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "projection": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "non_key_attributes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "projection_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "point_in_time_recovery_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "point_in_time_recovery_enabled": {
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
      "provisioned_throughput": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "read_capacity_units": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "write_capacity_units": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sse_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_master_key_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sse_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "sse_type": {
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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "table_class": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "time_to_live_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Version: None. Resource Type definition for AWS::DynamoDB::Table",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDynamodbTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDynamodbTable), &result)
	return &result
}
