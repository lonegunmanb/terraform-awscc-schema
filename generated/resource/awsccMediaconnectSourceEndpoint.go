package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectSourceEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associated_resource_details": {
        "computed": true,
        "description": "The details of the resource that is associated with a source endpoint. This is a union type, so exactly one of its members is set, depending on the type of the source endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "router_output": {
              "computed": true,
              "description": "The details of the router output that is associated with a source endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum_bitrate": {
                    "computed": true,
                    "description": "The maximum bitrate of the router output, in bits per second.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the router output.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "state": {
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
              "description": "The associated resource details for a standard source endpoint. This type does not require any additional settings.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone in which the source endpoint is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The ID of the Availability Zone in which the source endpoint is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "description": "The configuration settings for a source endpoint. This is a union type, so exactly one of its members is set, depending on the type of the source endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "router_output": {
              "computed": true,
              "description": "The configuration settings for a source endpoint that represents a router output.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "router_output_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the router output that the source endpoint represents.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "standard": {
              "computed": true,
              "description": "The configuration settings for a standard source endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the source endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol_configuration": {
                    "computed": true,
                    "description": "The protocol configuration for a standard source endpoint. This is a union type, so exactly one of its members is set, depending on the protocol that the source endpoint uses.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rist": {
                          "computed": true,
                          "description": "The configuration settings for a source endpoint that uses the RIST protocol.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination_ingest_port": {
                                "computed": true,
                                "description": "The port on the destination that receives the content. The valid range is 3000 to 30000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "recovery_latency_milliseconds": {
                                "computed": true,
                                "description": "The amount of time, in milliseconds, that the source endpoint allows for recovering lost packets. The valid range is 10 to 10000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "rtp": {
                          "computed": true,
                          "description": "The configuration settings for a source endpoint that uses the RTP protocol.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination_ingest_port": {
                                "computed": true,
                                "description": "The port on the destination that receives the content. The valid range is 3000 to 30000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "forward_error_correction_state": {
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
                        "srt_caller": {
                          "computed": true,
                          "description": "The configuration settings for a source endpoint that uses SRT in caller mode.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination_ingest_port": {
                                "computed": true,
                                "description": "The port on which the destination listens for the connection. The valid range is 3000 to 30000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "encryption_passphrase": {
                                "computed": true,
                                "description": "The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "secret_arn": {
                                      "computed": true,
                                      "description": "The ARN of the Secrets Manager secret that holds the key or passphrase. The secret must be encrypted with a customer managed Amazon Web Services KMS key that is enabled and symmetric; it cannot use the Amazon Web Services managed key aws/secretsmanager. The secret must be in the same Amazon Web Services account and Region as the endpoint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "minimum_latency_milliseconds": {
                                "computed": true,
                                "description": "The minimum latency, in milliseconds, for the SRT connection. The valid range is 10 to 10000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "srt_listener": {
                          "computed": true,
                          "description": "The configuration settings for a source endpoint that uses SRT in listener mode.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "encryption_passphrase": {
                                "computed": true,
                                "description": "The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "secret_arn": {
                                      "computed": true,
                                      "description": "The ARN of the Secrets Manager secret that holds the key or passphrase. The secret must be encrypted with a customer managed Amazon Web Services KMS key that is enabled and symmetric; it cannot use the Amazon Web Services managed key aws/secretsmanager. The secret must be in the same Amazon Web Services account and Region as the endpoint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "minimum_latency_milliseconds": {
                                "computed": true,
                                "description": "The minimum latency, in milliseconds, for the SRT connection. The valid range is 10 to 10000.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "source_address": {
                                "computed": true,
                                "description": "The IP address that the source endpoint listens on for incoming connections.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "source_listener_port": {
                                "computed": true,
                                "description": "The port that the source endpoint listens on for incoming connections. The valid range is 1024 to 65535.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "stream_id": {
                                "computed": true,
                                "description": "The stream ID that identifies the content for the SRT connection.",
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
                  "source_networking_configuration": {
                    "computed": true,
                    "description": "The networking configuration for a standard source endpoint. This is a union type, so exactly one of its members is set: public for a publicly accessible endpoint, or vpc for an endpoint in a virtual private cloud (VPC).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "public": {
                          "computed": true,
                          "description": "The networking configuration for a publicly accessible source endpoint, including the CIDR blocks that are allowed to connect.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source_allow_rules": {
                                "computed": true,
                                "description": "The rules that specify which CIDR blocks can connect to the source endpoint. You can specify up to 10 rules.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cidr": {
                                      "computed": true,
                                      "description": "The CIDR block that is allowed to connect to the source endpoint, in CIDR notation.",
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
                        "vpc": {
                          "computed": true,
                          "description": "The networking configuration for a source endpoint in a virtual private cloud (VPC).",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "security_group_ids": {
                                "computed": true,
                                "description": "The IDs of the security groups to associate with the source endpoint. You can specify from 1 to 5 security groups.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id": {
                                "computed": true,
                                "description": "The ID of the subnet to use for the source endpoint.",
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
                  "source_networking_type": {
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
      },
      "connection_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time when the source endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the source endpoint.",
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
      "lifecycle_management": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the source endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "region_name": {
        "computed": true,
        "description": "The name of the Amazon Web Services Region in which the source endpoint is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "registration_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_endpoint_id": {
        "computed": true,
        "description": "The unique identifier of the source endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the source endpoint. Each tag is a key-value pair.",
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
      "updated_at": {
        "computed": true,
        "description": "The date and time when the source endpoint was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Represents a source endpoint in AWS Elemental MediaConnect that can be connected to an input",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectSourceEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectSourceEndpoint), &result)
	return &result
}
