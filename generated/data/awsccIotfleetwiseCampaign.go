package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotfleetwiseCampaign = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collection_scheme": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "condition_based_collection_scheme": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition_language_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "minimum_trigger_interval_ms": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "trigger_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "time_based_collection_scheme": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "period_ms": {
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
      "compression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_destination_configs": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mqtt_topic_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "execution_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mqtt_topic_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "s3_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_format": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "storage_compression_format": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "timestream_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "execution_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timestream_table_arn": {
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
      "data_extra_dimensions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "diagnostics_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiry_time": {
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
      "last_modification_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "post_trigger_collection_duration": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "priority": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "signal_catalog_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signals_to_collect": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_sample_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "minimum_sampling_interval_ms": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "signals_to_fetch": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "condition_language_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "fully_qualified_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "signal_fetch_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition_based": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_expression": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "trigger_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "time_based": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "execution_frequency_ms": {
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
            }
          },
          "nesting_mode": "list"
        }
      },
      "spooling_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
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
          "nesting_mode": "set"
        }
      },
      "target_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoTFleetWise::Campaign",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotfleetwiseCampaignSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotfleetwiseCampaign), &result)
	return &result
}
