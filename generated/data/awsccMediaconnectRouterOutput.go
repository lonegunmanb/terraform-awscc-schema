package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectRouterOutput = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone where you want to create the router output. This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "The configuration settings for a router output.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "media_connect_flow": {
              "computed": true,
              "description": "Configuration settings for connecting a router output to a MediaConnect flow source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_transit_encryption": {
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
                  },
                  "flow_arn": {
                    "computed": true,
                    "description": "The ARN of the flow to connect to this router output.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "flow_source_arn": {
                    "computed": true,
                    "description": "The ARN of the flow source to connect to this router output.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "media_live_input": {
              "computed": true,
              "description": "Configuration settings for connecting a router output to a MediaLive input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_transit_encryption": {
                    "computed": true,
                    "description": "The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive. This configuration determines whether encryption keys are automatically managed by the service or manually managed through AWS Secrets Manager.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_key_configuration": {
                          "computed": true,
                          "description": "Configuration settings for the MediaLive transit encryption key.",
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
                  "media_live_input_arn": {
                    "computed": true,
                    "description": "The ARN of the MediaLive input to connect to this router output.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "media_live_pipeline_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "standard": {
              "computed": true,
              "description": "The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "network_interface_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the network interface associated with the standard router output.",
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
                    "description": "The protocol configuration settings for a router output.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rist": {
                          "computed": true,
                          "description": "The configuration settings for a router output using the RIST (Reliable Internet Stream Transport) protocol, including the destination address and port.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination_address": {
                                "computed": true,
                                "description": "The destination IP address for the RIST protocol in the router output configuration.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "destination_port": {
                                "computed": true,
                                "description": "The destination port number for the RIST protocol in the router output configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "rtp": {
                          "computed": true,
                          "description": "The configuration settings for a router output using the RTP (Real-Time Transport Protocol) protocol, including the destination address and port, and forward error correction state.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination_address": {
                                "computed": true,
                                "description": "The destination IP address for the RTP protocol in the router output configuration.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "destination_port": {
                                "computed": true,
                                "description": "The destination port number for the RTP protocol in the router output configuration.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "forward_error_correction": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "srt_caller": {
                          "computed": true,
                          "description": "The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in caller mode, including the destination address and port, minimum latency, stream ID, and encryption key configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination_address": {
                                "computed": true,
                                "description": "The destination IP address for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "destination_port": {
                                "computed": true,
                                "description": "The destination port number for the SRT protocol in caller mode.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "encryption_configuration": {
                                "computed": true,
                                "description": "Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.",
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
                          "description": "The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "encryption_configuration": {
                                "computed": true,
                                "description": "Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.",
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
        "description": "The timestamp when the router output was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "The IP address of the router output.",
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
        "description": "The maximum bitrate for the router output.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the router output.",
        "description_kind": "plain",
        "type": "string"
      },
      "output_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region_name": {
        "computed": true,
        "description": "The AWS Region for the router output. Defaults to the current region if not specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "routed_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "router_output_id": {
        "computed": true,
        "description": "The unique identifier of the router output.",
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
        "description": "Key-value pairs that can be used to tag this router output.",
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the router output was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaConnect::RouterOutput",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediaconnectRouterOutputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectRouterOutput), &result)
	return &result
}
