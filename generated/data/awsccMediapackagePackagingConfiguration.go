package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagePackagingConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the PackagingConfiguration.",
        "description_kind": "plain",
        "type": "string"
      },
      "cmaf_package": {
        "computed": true,
        "description": "A CMAF packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption": {
              "computed": true,
              "description": "A CMAF encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_contract_configuration": {
                          "computed": true,
                          "description": "The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "preset_speke_20_audio": {
                                "computed": true,
                                "description": "A collection of audio encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "preset_speke_20_video": {
                                "computed": true,
                                "description": "A collection of video encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage will assume when accessing the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "system_ids": {
                          "computed": true,
                          "description": "The system IDs to include in key requests.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "url": {
                          "computed": true,
                          "description": "The URL of the external key provider service.",
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
            "hls_manifests": {
              "computed": true,
              "description": "A list of HLS manifest configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_markers": {
                    "computed": true,
                    "description": "This setting controls how ad markers are included in the packaged OriginEndpoint. \"NONE\" will omit all SCTE-35 ad markers from the output. \"PASSTHROUGH\" causes the manifest to contain a copy of the SCTE-35 ad markers (comments) taken directly from the input HTTP Live Streaming (HLS) manifest. \"SCTE35_ENHANCED\" generates ad markers and blackout tags based on SCTE-35 messages in the input source.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "include_iframe_only_stream": {
                    "computed": true,
                    "description": "When enabled, an I-Frame only stream will be included in the output.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "manifest_name": {
                    "computed": true,
                    "description": "An optional string to include in the name of the manifest.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "program_date_time_interval_seconds": {
                    "computed": true,
                    "description": "The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into manifests. Additionally, when an interval is specified ID3Timed Metadata messages will be generated every 5 seconds using the ingest time of the content. If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME tags will be inserted into manifests and no ID3Timed Metadata messages will be generated. Note that irrespective of this parameter, if any ID3 Timed Metadata is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS output.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "repeat_ext_x_key": {
                    "computed": true,
                    "description": "When enabled, the EXT-X-KEY tag will be repeated in output manifests.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "stream_selection": {
                    "computed": true,
                    "description": "A StreamSelection configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_video_bits_per_second": {
                          "computed": true,
                          "description": "The maximum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "min_video_bits_per_second": {
                          "computed": true,
                          "description": "The minimum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stream_order": {
                          "computed": true,
                          "description": "A directive that determines the order of streams in the output.",
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
            "include_encoder_configuration_in_segments": {
              "computed": true,
              "description": "When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment. This lets you use different SPS/PPS/VPS settings for your assets during content playback.",
              "description_kind": "plain",
              "type": "bool"
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each fragment. Actual fragments will be rounded to the nearest multiple of the source fragment duration.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "dash_package": {
        "computed": true,
        "description": "A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dash_manifests": {
              "computed": true,
              "description": "A list of DASH manifest configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "manifest_layout": {
                    "computed": true,
                    "description": "Determines the position of some tags in the Media Presentation Description (MPD). When set to FULL, elements like SegmentTemplate and ContentProtection are included in each Representation. When set to COMPACT, duplicate elements are combined and presented at the AdaptationSet level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "manifest_name": {
                    "computed": true,
                    "description": "An optional string to include in the name of the manifest.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "min_buffer_time_seconds": {
                    "computed": true,
                    "description": "Minimum duration (in seconds) that a player will buffer media before starting the presentation.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "profile": {
                    "computed": true,
                    "description": "The Dynamic Adaptive Streaming over HTTP (DASH) profile type. When set to \"HBBTV_1_5\", HbbTV 1.5 compliant output is enabled.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "scte_markers_source": {
                    "computed": true,
                    "description": "The source of scte markers used. When set to SEGMENTS, the scte markers are sourced from the segments of the ingested content. When set to MANIFEST, the scte markers are sourced from the manifest of the ingested content.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stream_selection": {
                    "computed": true,
                    "description": "A StreamSelection configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_video_bits_per_second": {
                          "computed": true,
                          "description": "The maximum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "min_video_bits_per_second": {
                          "computed": true,
                          "description": "The minimum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stream_order": {
                          "computed": true,
                          "description": "A directive that determines the order of streams in the output.",
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
            "encryption": {
              "computed": true,
              "description": "A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_contract_configuration": {
                          "computed": true,
                          "description": "The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "preset_speke_20_audio": {
                                "computed": true,
                                "description": "A collection of audio encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "preset_speke_20_video": {
                                "computed": true,
                                "description": "A collection of video encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage will assume when accessing the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "system_ids": {
                          "computed": true,
                          "description": "The system IDs to include in key requests.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "url": {
                          "computed": true,
                          "description": "The URL of the external key provider service.",
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
            "include_encoder_configuration_in_segments": {
              "computed": true,
              "description": "When includeEncoderConfigurationInSegments is set to true, MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment. This lets you use different SPS/PPS/VPS settings for your assets during content playback.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_iframe_only_stream": {
              "computed": true,
              "description": "When enabled, an I-Frame only stream will be included in the output.",
              "description_kind": "plain",
              "type": "bool"
            },
            "period_triggers": {
              "computed": true,
              "description": "A list of triggers that controls when the outgoing Dynamic Adaptive Streaming over HTTP (DASH) Media Presentation Description (MPD) will be partitioned into multiple periods. If empty, the content will not be partitioned into more than one period. If the list contains \"ADS\", new periods will be created where the Asset contains SCTE-35 ad markers.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each fragment. Actual fragments will be rounded to the nearest multiple of the source fragment duration.",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_template_format": {
              "computed": true,
              "description": "Determines the type of SegmentTemplate included in the Media Presentation Description (MPD). When set to NUMBER_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Number$ media URLs. When set to TIME_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Time$ media URLs. When set to NUMBER_WITH_DURATION, only a duration is included in each SegmentTemplate, with $Number$ media URLs.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "hls_package": {
        "computed": true,
        "description": "An HTTP Live Streaming (HLS) packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption": {
              "computed": true,
              "description": "An HTTP Live Streaming (HLS) encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constant_initialization_vector": {
                    "computed": true,
                    "description": "An HTTP Live Streaming (HLS) encryption configuration.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "encryption_method": {
                    "computed": true,
                    "description": "The encryption method to use.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_contract_configuration": {
                          "computed": true,
                          "description": "The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "preset_speke_20_audio": {
                                "computed": true,
                                "description": "A collection of audio encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "preset_speke_20_video": {
                                "computed": true,
                                "description": "A collection of video encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage will assume when accessing the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "system_ids": {
                          "computed": true,
                          "description": "The system IDs to include in key requests.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "url": {
                          "computed": true,
                          "description": "The URL of the external key provider service.",
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
            "hls_manifests": {
              "computed": true,
              "description": "A list of HLS manifest configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_markers": {
                    "computed": true,
                    "description": "This setting controls how ad markers are included in the packaged OriginEndpoint. \"NONE\" will omit all SCTE-35 ad markers from the output. \"PASSTHROUGH\" causes the manifest to contain a copy of the SCTE-35 ad markers (comments) taken directly from the input HTTP Live Streaming (HLS) manifest. \"SCTE35_ENHANCED\" generates ad markers and blackout tags based on SCTE-35 messages in the input source.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "include_iframe_only_stream": {
                    "computed": true,
                    "description": "When enabled, an I-Frame only stream will be included in the output.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "manifest_name": {
                    "computed": true,
                    "description": "An optional string to include in the name of the manifest.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "program_date_time_interval_seconds": {
                    "computed": true,
                    "description": "The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into manifests. Additionally, when an interval is specified ID3Timed Metadata messages will be generated every 5 seconds using the ingest time of the content. If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME tags will be inserted into manifests and no ID3Timed Metadata messages will be generated. Note that irrespective of this parameter, if any ID3 Timed Metadata is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS output.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "repeat_ext_x_key": {
                    "computed": true,
                    "description": "When enabled, the EXT-X-KEY tag will be repeated in output manifests.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "stream_selection": {
                    "computed": true,
                    "description": "A StreamSelection configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_video_bits_per_second": {
                          "computed": true,
                          "description": "The maximum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "min_video_bits_per_second": {
                          "computed": true,
                          "description": "The minimum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stream_order": {
                          "computed": true,
                          "description": "A directive that determines the order of streams in the output.",
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
            "include_dvb_subtitles": {
              "computed": true,
              "description": "When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.",
              "description_kind": "plain",
              "type": "bool"
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each fragment. Actual fragments will be rounded to the nearest multiple of the source fragment duration.",
              "description_kind": "plain",
              "type": "number"
            },
            "use_audio_rendition_group": {
              "computed": true,
              "description": "When enabled, audio streams will be placed in rendition groups in the output.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mss_package": {
        "computed": true,
        "description": "A Microsoft Smooth Streaming (MSS) PackagingConfiguration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption": {
              "computed": true,
              "description": "A CMAF encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_contract_configuration": {
                          "computed": true,
                          "description": "The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "preset_speke_20_audio": {
                                "computed": true,
                                "description": "A collection of audio encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "preset_speke_20_video": {
                                "computed": true,
                                "description": "A collection of video encryption presets.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage will assume when accessing the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "system_ids": {
                          "computed": true,
                          "description": "The system IDs to include in key requests.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "url": {
                          "computed": true,
                          "description": "The URL of the external key provider service.",
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
            "mss_manifests": {
              "computed": true,
              "description": "A list of MSS manifest configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "manifest_name": {
                    "computed": true,
                    "description": "An optional string to include in the name of the manifest.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stream_selection": {
                    "computed": true,
                    "description": "A StreamSelection configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_video_bits_per_second": {
                          "computed": true,
                          "description": "The maximum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "min_video_bits_per_second": {
                          "computed": true,
                          "description": "The minimum video bitrate (bps) to include in output.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stream_order": {
                          "computed": true,
                          "description": "A directive that determines the order of streams in the output.",
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
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each fragment. Actual fragments will be rounded to the nearest multiple of the source fragment duration.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "packaging_group_id": {
        "computed": true,
        "description": "The ID of a PackagingGroup.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
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
      }
    },
    "description": "Data Source schema for AWS::MediaPackage::PackagingConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackagePackagingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagePackagingConfiguration), &result)
	return &result
}
