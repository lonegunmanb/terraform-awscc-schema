package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectRouterInput = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone where you want to create the router input. This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "The configuration settings for a router input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failover": {
              "computed": true,
              "description": "Configuration settings for a failover router input that allows switching between two input sources.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "network_interface_arn": {
                    "computed": true,
                    "description": "The ARN of the network interface to use for this failover router input.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "primary_source_index": {
                    "computed": true,
                    "description": "The index (0 or 1) that specifies which source in the protocol configurations list is currently active. Used to control which of the two failover sources is currently selected. This field is ignored when sourcePriorityMode is set to NO_PRIORITY",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "protocol_configurations": {
                    "computed": true,
                    "description": "A list of exactly two protocol configurations for the failover input sources. Both must use the same protocol type.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rist": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "port": {
                                "computed": true,
                                "description": "The port number used for the RIST protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "recovery_latency_milliseconds": {
                                "computed": true,
                                "description": "The recovery latency in milliseconds for the RIST protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "rtp": {
                          "computed": true,
                          "description": "The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward_error_correction": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "The port number used for the RTP protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "srt_caller": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "decryption_configuration": {
                                "computed": true,
                                "description": "Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "encryption_key": {
                                      "computed": true,
                                      "description": "The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "role_arn": {
                                            "computed": true,
                                            "description": "The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "secret_arn": {
                                            "computed": true,
                                            "description": "The ARN of the AWS Secrets Manager secret used for transit encryption.",
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
                              "minimum_latency_milliseconds": {
                                "computed": true,
                                "description": "The minimum latency in milliseconds for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "source_address": {
                                "computed": true,
                                "description": "The source IP address for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "source_port": {
                                "computed": true,
                                "description": "The source port number for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "stream_id": {
                                "computed": true,
                                "description": "The stream ID for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "srt_listener": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "decryption_configuration": {
                                "computed": true,
                                "description": "Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "encryption_key": {
                                      "computed": true,
                                      "description": "The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "role_arn": {
                                            "computed": true,
                                            "description": "The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "secret_arn": {
                                            "computed": true,
                                            "description": "The ARN of the AWS Secrets Manager secret used for transit encryption.",
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
                              "minimum_latency_milliseconds": {
                                "computed": true,
                                "description": "The minimum latency in milliseconds for the SRT protocol in listener mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "port": {
                                "computed": true,
                                "description": "The port number for the SRT protocol in listener mode.",
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
                  "source_priority_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "media_connect_flow": {
              "computed": true,
              "description": "Configuration settings for connecting a router input to a flow output.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "flow_arn": {
                    "computed": true,
                    "description": "The ARN of the flow to connect to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "flow_output_arn": {
                    "computed": true,
                    "description": "The ARN of the flow output to connect to this router input.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_transit_decryption": {
                    "computed": true,
                    "description": "The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_key_configuration": {
                          "computed": true,
                          "description": "Configuration settings for flow transit encryption keys.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "automatic": {
                                "computed": true,
                                "description": "Configuration settings for automatic encryption key management, where MediaConnect handles key creation and rotation.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "secrets_manager": {
                                "computed": true,
                                "description": "The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "role_arn": {
                                      "computed": true,
                                      "description": "The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "secret_arn": {
                                      "computed": true,
                                      "description": "The ARN of the AWS Secrets Manager secret used for transit encryption.",
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
                        "encryption_key_type": {
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
            "merge": {
              "computed": true,
              "description": "Configuration settings for a merge router input that combines two input sources.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "merge_recovery_window_milliseconds": {
                    "computed": true,
                    "description": "The time window in milliseconds for merging the two input sources.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "network_interface_arn": {
                    "computed": true,
                    "description": "The ARN of the network interface to use for this merge router input.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol_configurations": {
                    "computed": true,
                    "description": "A list of exactly two protocol configurations for the merge input sources. Both must use the same protocol type.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rist": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "port": {
                                "computed": true,
                                "description": "The port number used for the RIST protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "recovery_latency_milliseconds": {
                                "computed": true,
                                "description": "The recovery latency in milliseconds for the RIST protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "rtp": {
                          "computed": true,
                          "description": "The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward_error_correction": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "The port number used for the RTP protocol in the router input configuration.",
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
                  }
                },
                "nesting_mode": "single"
              }
            },
            "standard": {
              "computed": true,
              "description": "The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "network_interface_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the network interface associated with the standard router input.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol_configuration": {
                    "computed": true,
                    "description": "The protocol configuration settings for a router input.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rist": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "port": {
                                "computed": true,
                                "description": "The port number used for the RIST protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "recovery_latency_milliseconds": {
                                "computed": true,
                                "description": "The recovery latency in milliseconds for the RIST protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "rtp": {
                          "computed": true,
                          "description": "The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward_error_correction": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "The port number used for the RTP protocol in the router input configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "srt_caller": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "decryption_configuration": {
                                "computed": true,
                                "description": "Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "encryption_key": {
                                      "computed": true,
                                      "description": "The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "role_arn": {
                                            "computed": true,
                                            "description": "The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "secret_arn": {
                                            "computed": true,
                                            "description": "The ARN of the AWS Secrets Manager secret used for transit encryption.",
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
                              "minimum_latency_milliseconds": {
                                "computed": true,
                                "description": "The minimum latency in milliseconds for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "source_address": {
                                "computed": true,
                                "description": "The source IP address for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "source_port": {
                                "computed": true,
                                "description": "The source port number for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "stream_id": {
                                "computed": true,
                                "description": "The stream ID for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "srt_listener": {
                          "computed": true,
                          "description": "The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "decryption_configuration": {
                                "computed": true,
                                "description": "Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "encryption_key": {
                                      "computed": true,
                                      "description": "The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "role_arn": {
                                            "computed": true,
                                            "description": "The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "secret_arn": {
                                            "computed": true,
                                            "description": "The ARN of the AWS Secrets Manager secret used for transit encryption.",
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
                              "minimum_latency_milliseconds": {
                                "computed": true,
                                "description": "The minimum latency in milliseconds for the SRT protocol in listener mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "port": {
                                "computed": true,
                                "description": "The port number for the SRT protocol in listener mode.",
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
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the router input was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "The IP address of the router input.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_configuration": {
        "computed": true,
        "description": "The configuration settings for maintenance operations, including preferred maintenance windows and schedules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default": {
              "computed": true,
              "description": "Configuration settings for default maintenance scheduling.",
              "description_kind": "plain",
              "type": "string"
            },
            "preferred_day_time": {
              "computed": true,
              "description": "Configuration for preferred day and time maintenance settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "time": {
                    "computed": true,
                    "description": "The preferred time for maintenance operations.",
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
      "maintenance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maximum_bitrate": {
        "computed": true,
        "description": "The maximum bitrate for the router input.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the router input.",
        "description_kind": "plain",
        "type": "string"
      },
      "region_name": {
        "computed": true,
        "description": "The AWS Region for the router input. Defaults to the current region if not specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "routed_outputs": {
        "computed": true,
        "description": "The number of router outputs associated with the router input.",
        "description_kind": "plain",
        "type": "number"
      },
      "router_input_id": {
        "computed": true,
        "description": "The unique identifier of the router input.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_scope": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to tag and organize this router input.",
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
      "tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_encryption": {
        "computed": true,
        "description": "The transit encryption settings for a router input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_key_configuration": {
              "computed": true,
              "description": "Defines the configuration settings for transit encryption keys.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "automatic": {
                    "computed": true,
                    "description": "Configuration settings for automatic encryption key management, where MediaConnect handles key creation and rotation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secrets_manager": {
                    "computed": true,
                    "description": "The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "role_arn": {
                          "computed": true,
                          "description": "The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secret_arn": {
                          "computed": true,
                          "description": "The ARN of the AWS Secrets Manager secret used for transit encryption.",
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
            "encryption_key_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the router input was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaConnect::RouterInput",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediaconnectRouterInputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectRouterInput), &result)
	return &result
}
