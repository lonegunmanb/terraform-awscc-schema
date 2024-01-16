package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "broker_node_group_info": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "broker_az_distribution": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "client_subnets": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "connectivity_info": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "public_access": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
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
                  "vpc_connectivity": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_authentication": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "sasl": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "iam": {
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
                                    "scram": {
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
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "tls": {
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "instance_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_groups": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "storage_info": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ebs_storage_info": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "provisioned_throughput": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "enabled": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "volume_throughput": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "volume_size": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "client_authentication": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sasl": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iam": {
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
                  "scram": {
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "tls": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_authority_arn_list": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enabled": {
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
            "unauthenticated": {
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cluster_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configuration_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "revision": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "current_version": {
        "computed": true,
        "description": "The current version of the MSK cluster",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_at_rest": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_volume_kms_key_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "encryption_in_transit": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_broker": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "in_cluster": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "enhanced_monitoring": {
        "computed": true,
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
      "kafka_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "logging_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "broker_logs": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_logs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "log_group": {
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
                  "firehose": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream": {
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
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "enabled": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "prefix": {
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
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "number_of_broker_nodes": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "open_monitoring": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "prometheus": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "jmx_exporter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled_in_broker": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "node_exporter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled_in_broker": {
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
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "storage_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::MSK::Cluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskCluster), &result)
	return &result
}
