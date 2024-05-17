package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectFlow = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "egress_ip": {
        "computed": true,
        "description": "The IP address from which video will be sent to output destinations.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "flow_availability_zone": {
        "computed": true,
        "description": "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly)",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance": {
        "computed": true,
        "description": "The maintenance settings you want to use for the flow. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maintenance_day": {
              "description": "A day of a week when the maintenance will happen. Use Monday/Tuesday/Wednesday/Thursday/Friday/Saturday/Sunday.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "maintenance_start_hour": {
              "description": "UTC time when the maintenance will happen. Use 24-hour HH:MM format. Minutes must be 00. Example: 13:00. The default value is 02:00.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "media_streams": {
        "computed": true,
        "description": "The media streams associated with the flow. You can associate any of these media streams with sources and outputs on the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "Attributes that are related to the media stream.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "fmtp": {
                    "computed": true,
                    "description": "A set of parameters that define the media stream.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "channel_order": {
                          "computed": true,
                          "description": "The format of the audio channel.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "colorimetry": {
                          "computed": true,
                          "description": "The format used for the representation of color.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "exact_framerate": {
                          "computed": true,
                          "description": "The frame rate for the video stream, in frames/second. For example: 60000/1001.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "par": {
                          "computed": true,
                          "description": "The pixel aspect ratio (PAR) of the video.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range": {
                          "computed": true,
                          "description": "The encoding range of the video.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scan_mode": {
                          "computed": true,
                          "description": "The type of compression that was used to smooth the video's appearance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tcs": {
                          "computed": true,
                          "description": "The transfer characteristic system (TCS) that is used in the video.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lang": {
                    "computed": true,
                    "description": "The audio language, in a format that is recognized by the receiver.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "clock_rate": {
              "computed": true,
              "description": "The sample rate for the stream. This value in measured in kHz.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "description": {
              "computed": true,
              "description": "A description that can help you quickly identify what your media stream is used for.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "fmt": {
              "computed": true,
              "description": "The format type number (sometimes referred to as RTP payload type) of the media stream. MediaConnect assigns this value to the media stream. For ST 2110 JPEG XS outputs, you need to provide this value to the receiver.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "media_stream_id": {
              "description": "A unique identifier for the media stream.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "media_stream_name": {
              "description": "A name that helps you distinguish one media stream from another.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "media_stream_type": {
              "description": "The type of media stream.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "video_format": {
              "computed": true,
              "description": "The resolution of the video.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "name": {
        "description": "The name of the flow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "description": "The source of the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "decryption": {
              "computed": true,
              "description": "The type of decryption that is used on the content ingested from this source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "algorithm": {
                    "computed": true,
                    "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "constant_initialization_vector": {
                    "computed": true,
                    "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "device_id": {
                    "computed": true,
                    "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_type": {
                    "computed": true,
                    "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_id": {
                    "computed": true,
                    "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "url": {
                    "computed": true,
                    "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "description": {
              "computed": true,
              "description": "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "entitlement_arn": {
              "computed": true,
              "description": "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "gateway_bridge_source": {
              "computed": true,
              "description": "The source configuration for cloud flows receiving a stream from a bridge.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bridge_arn": {
                    "description": "The ARN of the bridge feeding this flow.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "vpc_interface_attachment": {
                    "computed": true,
                    "description": "The name of the VPC interface attachment to use for this bridge source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "vpc_interface_name": {
                          "computed": true,
                          "description": "The name of the VPC interface to use for this resource.",
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
            "ingest_ip": {
              "computed": true,
              "description": "The IP address that the flow will be listening on for incoming content.",
              "description_kind": "plain",
              "type": "string"
            },
            "ingest_port": {
              "computed": true,
              "description": "The port that the flow will be listening on for incoming content.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_bitrate": {
              "computed": true,
              "description": "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_latency": {
              "computed": true,
              "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_sync_buffer": {
              "computed": true,
              "description": "The size of the buffer (in milliseconds) to use to sync incoming source data.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "media_stream_source_configurations": {
              "computed": true,
              "description": "The media stream that is associated with the source, and the parameters for that association.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encoding_name": {
                    "description": "The format that was used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "input_configurations": {
                    "computed": true,
                    "description": "The media streams that you want to associate with the source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "input_port": {
                          "description": "The port that the flow listens on for an incoming media stream.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "interface": {
                          "description": "The VPC interface where the media stream comes in from.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "description": "The name of the VPC interface that you want to use for the media stream associated with the output.",
                                "description_kind": "plain",
                                "required": true,
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
                  "media_stream_name": {
                    "description": "A name that helps you distinguish one media stream from another.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "min_latency": {
              "computed": true,
              "description": "The minimum latency in milliseconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "computed": true,
              "description": "The name of the source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "The protocol that is used by the source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sender_control_port": {
              "computed": true,
              "description": "The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sender_ip_address": {
              "computed": true,
              "description": "The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_arn": {
              "computed": true,
              "description": "The ARN of the source.",
              "description_kind": "plain",
              "type": "string"
            },
            "source_ingest_port": {
              "computed": true,
              "description": "The port that the flow will be listening on for incoming content.(ReadOnly)",
              "description_kind": "plain",
              "type": "string"
            },
            "source_listener_address": {
              "computed": true,
              "description": "Source IP or domain name for SRT-caller protocol.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_listener_port": {
              "computed": true,
              "description": "Source port for SRT-caller protocol.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stream_id": {
              "computed": true,
              "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_interface_name": {
              "computed": true,
              "description": "The name of the VPC Interface this Source is configured with.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "whitelist_cidr": {
              "computed": true,
              "description": "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "source_failover_config": {
        "computed": true,
        "description": "The source failover config of the flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failover_mode": {
              "computed": true,
              "description": "The type of failover you choose for this flow. MERGE combines the source streams into a single stream, allowing graceful recovery from any single-source loss. FAILOVER allows switching between different streams.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recovery_window": {
              "computed": true,
              "description": "Search window time to look for dash-7 packets",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "source_priority": {
              "computed": true,
              "description": "The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "primary_source": {
                    "description": "The name of the source you choose as the primary source for this flow.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "state": {
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
      "vpc_interfaces": {
        "computed": true,
        "description": "The VPC interfaces that you added to this flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_interface_ids": {
              "computed": true,
              "description": "IDs of the network interfaces created in customer's account by MediaConnect.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "network_interface_type": {
              "computed": true,
              "description": "The type of network adapter that you want MediaConnect to use on this interface. If you don't set this value, it defaults to ENA.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "description": "Role Arn MediaConnect can assume to create ENIs in customer's account.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "security_group_ids": {
              "description": "Security Group IDs to be used on ENI.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_id": {
              "description": "Subnet must be in the AZ of the Flow",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::MediaConnect::Flow",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectFlow), &result)
	return &result
}
