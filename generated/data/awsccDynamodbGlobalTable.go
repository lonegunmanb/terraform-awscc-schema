package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDynamodbGlobalTable = `{
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
          "nesting_mode": "set"
        }
      },
      "billing_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_secondary_indexes": {
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
                      "set",
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
            "write_provisioned_throughput_settings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "write_capacity_auto_scaling_settings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "min_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "seed_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "target_tracking_scaling_policy_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "disable_scale_in": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "scale_in_cooldown": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "scale_out_cooldown": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "target_value": {
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
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
                      "set",
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
          "nesting_mode": "set"
        }
      },
      "replicas": {
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
                  "read_provisioned_throughput_settings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "read_capacity_auto_scaling_settings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_capacity": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "min_capacity": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "seed_capacity": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "target_tracking_scaling_policy_configuration": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "disable_scale_in": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "scale_in_cooldown": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "scale_out_cooldown": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "target_value": {
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
                        "read_capacity_units": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "set"
              }
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
            "read_provisioned_throughput_settings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_capacity_auto_scaling_settings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "min_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "seed_capacity": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "target_tracking_scaling_policy_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "disable_scale_in": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "scale_in_cooldown": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "scale_out_cooldown": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "target_value": {
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
                  "read_capacity_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
            }
          },
          "nesting_mode": "set"
        }
      },
      "sse_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
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
      "table_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      },
      "write_provisioned_throughput_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "write_capacity_auto_scaling_settings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "seed_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "target_tracking_scaling_policy_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "disable_scale_in": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "scale_in_cooldown": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "scale_out_cooldown": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "target_value": {
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
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DynamoDB::GlobalTable",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDynamodbGlobalTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDynamodbGlobalTable), &result)
	return &result
}
