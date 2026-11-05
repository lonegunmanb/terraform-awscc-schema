package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveInput = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Unique ARN of the input (generated, immutable).",
        "description_kind": "plain",
        "type": "string"
      },
      "attached_channels": {
        "computed": true,
        "description": "A list of channel IDs that that input is attached to (currently an input can only be attached to one channel).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "destination_urls": {
        "computed": true,
        "description": "A list of destination URLs of the input (PUSH-type).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "destinations": {
        "computed": true,
        "description": "A list of the destinations of the input (PUSH-type).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip": {
              "computed": true,
              "description": "The system-generated static IP address of endpoint.\nIt remains fixed for the lifetime of the input.",
              "description_kind": "plain",
              "type": "string"
            },
            "network": {
              "computed": true,
              "description": "The ID of the attached network.",
              "description_kind": "plain",
              "type": "string"
            },
            "network_routes": {
              "computed": true,
              "description": "If the push input has an input location of ON-PREM it's a requirement to specify what the route of the input\nis going to be on the customer local network.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description": "The CIDR of the route.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "gateway": {
                    "computed": true,
                    "description": "An optional gateway for the route.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "port": {
              "computed": true,
              "description": "The port number for the input.",
              "description_kind": "plain",
              "type": "string"
            },
            "static_ip_address": {
              "computed": true,
              "description": "The static IP address for the input destination.",
              "description_kind": "plain",
              "type": "string"
            },
            "stream_name": {
              "computed": true,
              "description": "The stream name for the input destination.",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "This represents the endpoint that the customer stream will be\npushed to.",
              "description_kind": "plain",
              "type": "string"
            },
            "vpc": {
              "computed": true,
              "description": "The properties for a VPC type input destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone": {
                    "computed": true,
                    "description": "The availability zone of the Input destination.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_interface_id": {
                    "computed": true,
                    "description": "The network interface ID of the Input destination in the VPC.",
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_class": {
        "computed": true,
        "description": "STANDARD - MediaLive expects two sources to be connected to this input. If the channel is also STANDARD, both sources will be ingested. If the channel is SINGLE_PIPELINE, only the first source will be ingested; the second source will always be ignored, even if the first source fails.\nSINGLE_PIPELINE - You can connect only one source to this input. If the ChannelClass is also SINGLE_PIPELINE, this value is valid. If the ChannelClass is STANDARD, this value is not valid because the channel requires two sources in the input.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_devices": {
        "computed": true,
        "description": "Settings for the input devices.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The unique ID for the device.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "input_id": {
        "computed": true,
        "description": "The generated ID of the input (unique for user account, immutable).",
        "description_kind": "plain",
        "type": "string"
      },
      "input_network_location": {
        "computed": true,
        "description": "The location of this input. AWS, for an input existing in the AWS Cloud, On-Prem for\nan input in a customer network.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_partner_ids": {
        "computed": true,
        "description": "A list of IDs for all Inputs which are partners of this one.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "input_security_groups": {
        "computed": true,
        "description": "A list of security groups referenced by IDs to attach to the input.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "input_source_type": {
        "computed": true,
        "description": "Certain pull input sources can be dynamic, meaning that they can have their URL's dynamically changes\nduring input switch actions. Presently, this functionality only works with MP4_FILE and TS_FILE inputs.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_connect_flows": {
        "computed": true,
        "description": "A list of MediaConnect Flows for this input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "flow_arn": {
              "computed": true,
              "description": "The unique ARN of the MediaConnect Flow being used as a source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "multicast_settings": {
        "computed": true,
        "description": "Multicast Input settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sources": {
              "computed": true,
              "description": "Placeholder documentation for __listOfMulticastSource",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_ip": {
                    "computed": true,
                    "description": "This represents the ip address of the device sending the multicast stream.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "url": {
                    "computed": true,
                    "description": "This represents the customer's source URL where multicast stream is pulled from.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The user-assigned name (This is a mutable value).",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the role this input assumes during and after creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "router_settings": {
        "computed": true,
        "description": "Information about the MediaConnect Router association for this input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description": "MediaConnect Router destinations associated with this input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_name": {
                    "computed": true,
                    "description": "The Availability Zone (AZ) name this destination is created in.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "router_output_arn": {
                    "computed": true,
                    "description": "The ARN of the MediaConnect Router output currently connected to this input.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "encryption_type": {
              "computed": true,
              "description": "Encryption configuration for the MediaConnect Router association. AUTOMATIC uses a service-managed secret; SECRETS_MANAGER requires SecretArn.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "The ARN of the Secrets Manager secret used to encrypt this input. Required when EncryptionType is SECRETS_MANAGER.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sdi_sources": {
        "computed": true,
        "description": "SDI Sources for this Input.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "security_groups": {
        "computed": true,
        "description": "A list of IDs for all the Input Security Groups attached to the input.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "smpte_2110_receiver_group_settings": {
        "computed": true,
        "description": "Include this parameter if the input is a SMPTE 2110 input, to identify the stream sources for this input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "smpte_2110_receiver_groups": {
              "computed": true,
              "description": "Placeholder documentation for __listOfSmpte2110ReceiverGroup",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sdp_settings": {
                    "computed": true,
                    "description": "The single Smpte2110ReceiverGroupSdpSettings that identify the video, audio, and ancillary streams for this receiver group.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ancillary_sdps": {
                          "computed": true,
                          "description": "A list of InputSdpLocations. Each item in the list specifies the SDP file and index for one ancillary SMPTE 2110 stream.\nEach stream encapsulates one captions stream (out of any number you can include) or the single SCTE 35 stream that you can include.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "media_index": {
                                "computed": true,
                                "description": "The index of the media stream in the SDP file for one SMPTE 2110 stream.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "sdp_url": {
                                "computed": true,
                                "description": "The URL of the SDP file for one SMPTE 2110 stream.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "audio_sdps": {
                          "computed": true,
                          "description": "A list of InputSdpLocations. Each item in the list specifies the SDP file and index for one audio SMPTE 2110 stream.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "media_index": {
                                "computed": true,
                                "description": "The index of the media stream in the SDP file for one SMPTE 2110 stream.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "sdp_url": {
                                "computed": true,
                                "description": "The URL of the SDP file for one SMPTE 2110 stream.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "video_sdp": {
                          "computed": true,
                          "description": "The InputSdpLocation that specifies the SDP file and index for the single video SMPTE 2110 stream for this 2110 input.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "media_index": {
                                "computed": true,
                                "description": "The index of the media stream in the SDP file for one SMPTE 2110 stream.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "sdp_url": {
                                "computed": true,
                                "description": "The URL of the SDP file for one SMPTE 2110 stream.",
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
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_urls": {
        "computed": true,
        "description": "A list of source URLs of the input (PULL-type).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sources": {
        "computed": true,
        "description": "A list of the sources of the input (PULL-type).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "password_param": {
              "computed": true,
              "description": "The key used to extract the password from EC2 Parameter store.",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "This represents the customer's source URL where stream is\npulled from.",
              "description_kind": "plain",
              "type": "string"
            },
            "username": {
              "computed": true,
              "description": "The username for the input source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "srt_settings": {
        "computed": true,
        "description": "The settings associated with an SRT input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "srt_caller_sources": {
              "computed": true,
              "description": "Placeholder documentation for __listOfSrtCallerSource",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "decryption": {
                    "computed": true,
                    "description": "The decryption settings for the SRT caller source. Present only if the source has decryption enabled.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "algorithm": {
                          "computed": true,
                          "description": "The algorithm used to encrypt content.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "passphrase_secret_arn": {
                          "computed": true,
                          "description": "The ARN for the secret in Secrets Manager. Someone in your organization must create a secret and provide you with its ARN. The secret holds the passphrase that MediaLive uses to decrypt the source content.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "minimum_latency": {
                    "computed": true,
                    "description": "The preferred latency (in milliseconds) for implementing packet loss and recovery. Packet recovery is a key feature of SRT.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "srt_listener_address": {
                    "computed": true,
                    "description": "The IP address at the upstream system (the listener) that MediaLive (the caller) connects to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "srt_listener_port": {
                    "computed": true,
                    "description": "The port at the upstream system (the listener) that MediaLive (the caller) connects to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stream_id": {
                    "computed": true,
                    "description": "The stream ID, if the upstream system uses this identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "srt_listener_settings": {
              "computed": true,
              "description": "Settings for an SRT listener input. Configured only when the input uses the SRT_LISTENER input type.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "decryption": {
                    "computed": true,
                    "description": "The decryption settings for the SRT listener source. Present only if the source has decryption enabled.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "algorithm": {
                          "computed": true,
                          "description": "The algorithm used to encrypt content.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "passphrase_secret_arn": {
                          "computed": true,
                          "description": "The ARN for the secret in Secrets Manager. Someone in your organization must create a secret and provide you with its ARN. The secret holds the passphrase that MediaLive uses to decrypt the source content.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "minimum_latency": {
                    "computed": true,
                    "description": "The preferred latency (in milliseconds) for implementing packet loss and recovery. Packet recovery is a key feature of SRT.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "stream_id": {
                    "computed": true,
                    "description": "The stream ID, if the upstream system uses this identifier.",
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
      "state": {
        "computed": true,
        "description": "Placeholder documentation for InputState",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of key-value pairs.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "The different types of inputs that AWS Elemental MediaLive supports.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc": {
        "computed": true,
        "description": "Settings for a private VPC Input.\nWhen this property is specified, the input destination addresses will be created in a VPC rather than with public Internet addresses.\nThis property requires setting the roleArn property on Input creation.\nNot compatible with the inputSecurityGroups property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "A list of up to 5 EC2 VPC security group IDs to attach to the Input VPC network interfaces.\nRequires subnetIds. If none are specified then the VPC default security group will be used.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of 2 VPC subnet IDs from the same VPC.\nSubnet IDs must be mapped to two unique availability zones (AZ).",
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
    "description": "Data Source schema for AWS::MediaLive::Input",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMedialiveInputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveInput), &result)
	return &result
}
