package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackageOriginEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) assigned to the OriginEndpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorization": {
        "computed": true,
        "description": "CDN Authorization credentials",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cdn_identifier_secret": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the secret in Secrets Manager that your Content Distribution Network (CDN) uses for authorization to access your endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "channel_id": {
        "computed": true,
        "description": "The ID of the Channel the OriginEndpoint is associated with.",
        "description_kind": "plain",
        "type": "string"
      },
      "cmaf_package": {
        "computed": true,
        "description": "A Common Media Application Format (CMAF) packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption": {
              "computed": true,
              "description": "A Common Media Application Format (CMAF) encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constant_initialization_vector": {
                    "computed": true,
                    "description": "An optional 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting blocks. If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "encryption_method": {
                    "computed": true,
                    "description": "The encryption method used",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_rotation_interval_seconds": {
                    "computed": true,
                    "description": "Time (in seconds) between each encryption key rotation.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "certificate_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of a Certificate Manager certificate that MediaPackage will use for enforcing secure end-to-end data transfer with the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
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
                        "resource_id": {
                          "computed": true,
                          "description": "The resource ID to include in key requests.",
                          "description_kind": "plain",
                          "type": "string"
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
              "description": "A list of HLS manifest configurations",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_markers": {
                    "computed": true,
                    "description": "This setting controls how ad markers are included in the packaged OriginEndpoint. \"NONE\" will omit all SCTE-35 ad markers from the output. \"PASSTHROUGH\" causes the manifest to contain a copy of the SCTE-35 ad markers (comments) taken directly from the input HTTP Live Streaming (HLS) manifest. \"SCTE35_ENHANCED\" generates ad markers and blackout tags based on SCTE-35 messages in the input source. \"DATERANGE\" inserts EXT-X-DATERANGE tags to signal ad and program transition events in HLS and CMAF manifests. For this option, you must set a programDateTimeIntervalSeconds value that is greater than 0.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ad_triggers": {
                    "computed": true,
                    "description": "A list of SCTE-35 message types that are treated as ad markers in the output.  If empty, no ad markers are output.  Specify multiple items to create ad markers for all of the included message types.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ads_on_delivery_restrictions": {
                    "computed": true,
                    "description": "This setting allows the delivery restriction flags on SCTE-35 segmentation descriptors to determine whether a message signals an ad.  Choosing \"NONE\" means no SCTE-35 messages become ads.  Choosing \"RESTRICTED\" means SCTE-35 messages of the types specified in AdTriggers that contain delivery restrictions will be treated as ads.  Choosing \"UNRESTRICTED\" means SCTE-35 messages of the types specified in AdTriggers that do not contain delivery restrictions will be treated as ads.  Choosing \"BOTH\" means all SCTE-35 messages of the types specified in AdTriggers will be treated as ads.  Note that Splice Insert messages do not have these flags and are always treated as ads if specified in AdTriggers.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description": "The ID of the manifest. The ID must be unique within the OriginEndpoint and it cannot be changed after it is created.",
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
                    "description": "An optional short string appended to the end of the OriginEndpoint URL. If not specified, defaults to the manifestName for the OriginEndpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "playlist_type": {
                    "computed": true,
                    "description": "The HTTP Live Streaming (HLS) playlist type. When either \"EVENT\" or \"VOD\" is specified, a corresponding EXT-X-PLAYLIST-TYPE entry will be included in the media playlist.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "playlist_window_seconds": {
                    "computed": true,
                    "description": "Time window (in seconds) contained in each parent manifest.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "program_date_time_interval_seconds": {
                    "computed": true,
                    "description": "The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into manifests. Additionally, when an interval is specified ID3Timed Metadata messages will be generated every 5 seconds using the ingest time of the content. If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME tags will be inserted into manifests and no ID3Timed Metadata messages will be generated. Note that irrespective of this parameter, if any ID3 Timed Metadata is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS output.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "url": {
                    "computed": true,
                    "description": "The URL of the packaged OriginEndpoint for consumption.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each segment. Actual segments will be rounded to the nearest multiple of the source segment duration.",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_prefix": {
              "computed": true,
              "description": "An optional custom string that is prepended to the name of each segment. If not specified, it defaults to the ChannelId.",
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
          "nesting_mode": "single"
        }
      },
      "dash_package": {
        "computed": true,
        "description": "A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ad_triggers": {
              "computed": true,
              "description": "A list of SCTE-35 message types that are treated as ad markers in the output.  If empty, no ad markers are output.  Specify multiple items to create ad markers for all of the included message types.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "ads_on_delivery_restrictions": {
              "computed": true,
              "description": "This setting allows the delivery restriction flags on SCTE-35 segmentation descriptors to determine whether a message signals an ad.  Choosing \"NONE\" means no SCTE-35 messages become ads.  Choosing \"RESTRICTED\" means SCTE-35 messages of the types specified in AdTriggers that contain delivery restrictions will be treated as ads.  Choosing \"UNRESTRICTED\" means SCTE-35 messages of the types specified in AdTriggers that do not contain delivery restrictions will be treated as ads.  Choosing \"BOTH\" means all SCTE-35 messages of the types specified in AdTriggers will be treated as ads.  Note that Splice Insert messages do not have these flags and are always treated as ads if specified in AdTriggers.",
              "description_kind": "plain",
              "type": "string"
            },
            "encryption": {
              "computed": true,
              "description": "A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key_rotation_interval_seconds": {
                    "computed": true,
                    "description": "Time (in seconds) between each encryption key rotation.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "certificate_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of a Certificate Manager certificate that MediaPackage will use for enforcing secure end-to-end data transfer with the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
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
                        "resource_id": {
                          "computed": true,
                          "description": "The resource ID to include in key requests.",
                          "description_kind": "plain",
                          "type": "string"
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
            "include_iframe_only_stream": {
              "computed": true,
              "description": "When enabled, an I-Frame only stream will be included in the output.",
              "description_kind": "plain",
              "type": "bool"
            },
            "manifest_layout": {
              "computed": true,
              "description": "Determines the position of some tags in the Media Presentation Description (MPD).  When set to FULL, elements like SegmentTemplate and ContentProtection are included in each Representation.  When set to COMPACT, duplicate elements are combined and presented at the AdaptationSet level.",
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_window_seconds": {
              "computed": true,
              "description": "Time window (in seconds) contained in each manifest.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_buffer_time_seconds": {
              "computed": true,
              "description": "Minimum duration (in seconds) that a player will buffer media before starting the presentation.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_update_period_seconds": {
              "computed": true,
              "description": "Minimum duration (in seconds) between potential changes to the Dynamic Adaptive Streaming over HTTP (DASH) Media Presentation Description (MPD).",
              "description_kind": "plain",
              "type": "number"
            },
            "period_triggers": {
              "computed": true,
              "description": "A list of triggers that controls when the outgoing Dynamic Adaptive Streaming over HTTP (DASH) Media Presentation Description (MPD) will be partitioned into multiple periods. If empty, the content will not be partitioned into more than one period. If the list contains \"ADS\", new periods will be created where the Channel source contains SCTE-35 ad markers.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "profile": {
              "computed": true,
              "description": "The Dynamic Adaptive Streaming over HTTP (DASH) profile type.  When set to \"HBBTV_1_5\", HbbTV 1.5 compliant output is enabled.",
              "description_kind": "plain",
              "type": "string"
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each segment. Actual segments will be rounded to the nearest multiple of the source segment duration.",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_template_format": {
              "computed": true,
              "description": "Determines the type of SegmentTemplate included in the Media Presentation Description (MPD).  When set to NUMBER_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Number$ media URLs.  When set to TIME_WITH_TIMELINE, a full timeline is presented in each SegmentTemplate, with $Time$ media URLs. When set to NUMBER_WITH_DURATION, only a duration is included in each SegmentTemplate, with $Number$ media URLs.",
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
            },
            "suggested_presentation_delay_seconds": {
              "computed": true,
              "description": "Duration (in seconds) to delay live content before presentation.",
              "description_kind": "plain",
              "type": "number"
            },
            "utc_timing": {
              "computed": true,
              "description": "Determines the type of UTCTiming included in the Media Presentation Description (MPD)",
              "description_kind": "plain",
              "type": "string"
            },
            "utc_timing_uri": {
              "computed": true,
              "description": "Specifies the value attribute of the UTCTiming field when utcTiming is set to HTTP-ISO, HTTP-HEAD or HTTP-XSDATE",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A short text description of the OriginEndpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "hls_package": {
        "computed": true,
        "description": "An HTTP Live Streaming (HLS) packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ad_markers": {
              "computed": true,
              "description": "This setting controls how ad markers are included in the packaged OriginEndpoint. \"NONE\" will omit all SCTE-35 ad markers from the output. \"PASSTHROUGH\" causes the manifest to contain a copy of the SCTE-35 ad markers (comments) taken directly from the input HTTP Live Streaming (HLS) manifest. \"SCTE35_ENHANCED\" generates ad markers and blackout tags based on SCTE-35 messages in the input source. \"DATERANGE\" inserts EXT-X-DATERANGE tags to signal ad and program transition events in HLS and CMAF manifests. For this option, you must set a programDateTimeIntervalSeconds value that is greater than 0.",
              "description_kind": "plain",
              "type": "string"
            },
            "ad_triggers": {
              "computed": true,
              "description": "A list of SCTE-35 message types that are treated as ad markers in the output.  If empty, no ad markers are output.  Specify multiple items to create ad markers for all of the included message types.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "ads_on_delivery_restrictions": {
              "computed": true,
              "description": "This setting allows the delivery restriction flags on SCTE-35 segmentation descriptors to determine whether a message signals an ad.  Choosing \"NONE\" means no SCTE-35 messages become ads.  Choosing \"RESTRICTED\" means SCTE-35 messages of the types specified in AdTriggers that contain delivery restrictions will be treated as ads.  Choosing \"UNRESTRICTED\" means SCTE-35 messages of the types specified in AdTriggers that do not contain delivery restrictions will be treated as ads.  Choosing \"BOTH\" means all SCTE-35 messages of the types specified in AdTriggers will be treated as ads.  Note that Splice Insert messages do not have these flags and are always treated as ads if specified in AdTriggers.",
              "description_kind": "plain",
              "type": "string"
            },
            "encryption": {
              "computed": true,
              "description": "An HTTP Live Streaming (HLS) encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constant_initialization_vector": {
                    "computed": true,
                    "description": "A constant initialization vector for encryption (optional). When not specified the initialization vector will be periodically rotated.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "encryption_method": {
                    "computed": true,
                    "description": "The encryption method to use.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key_rotation_interval_seconds": {
                    "computed": true,
                    "description": "Interval (in seconds) between each encryption key rotation.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "repeat_ext_x_key": {
                    "computed": true,
                    "description": "When enabled, the EXT-X-KEY tag will be repeated in output manifests.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "certificate_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of a Certificate Manager certificate that MediaPackage will use for enforcing secure end-to-end data transfer with the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
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
                        "resource_id": {
                          "computed": true,
                          "description": "The resource ID to include in key requests.",
                          "description_kind": "plain",
                          "type": "string"
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
            "include_dvb_subtitles": {
              "computed": true,
              "description": "When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_iframe_only_stream": {
              "computed": true,
              "description": "When enabled, an I-Frame only stream will be included in the output.",
              "description_kind": "plain",
              "type": "bool"
            },
            "playlist_type": {
              "computed": true,
              "description": "The HTTP Live Streaming (HLS) playlist type. When either \"EVENT\" or \"VOD\" is specified, a corresponding EXT-X-PLAYLIST-TYPE entry will be included in the media playlist.",
              "description_kind": "plain",
              "type": "string"
            },
            "playlist_window_seconds": {
              "computed": true,
              "description": "Time window (in seconds) contained in each parent manifest.",
              "description_kind": "plain",
              "type": "number"
            },
            "program_date_time_interval_seconds": {
              "computed": true,
              "description": "The interval (in seconds) between each EXT-X-PROGRAM-DATE-TIME tag inserted into manifests. Additionally, when an interval is specified ID3Timed Metadata messages will be generated every 5 seconds using the ingest time of the content. If the interval is not specified, or set to 0, then no EXT-X-PROGRAM-DATE-TIME tags will be inserted into manifests and no ID3Timed Metadata messages will be generated. Note that irrespective of this parameter, if any ID3 Timed Metadata is found in HTTP Live Streaming (HLS) input, it will be passed through to HLS output.",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "Duration (in seconds) of each fragment. Actual fragments will be rounded to the nearest multiple of the source fragment duration.",
              "description_kind": "plain",
              "type": "number"
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
      "manifest_name": {
        "computed": true,
        "description": "A short string appended to the end of the OriginEndpoint URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "mss_package": {
        "computed": true,
        "description": "A Microsoft Smooth Streaming (MSS) packaging configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption": {
              "computed": true,
              "description": "A Microsoft Smooth Streaming (MSS) encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "speke_key_provider": {
                    "computed": true,
                    "description": "A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "certificate_arn": {
                          "computed": true,
                          "description": "An Amazon Resource Name (ARN) of a Certificate Manager certificate that MediaPackage will use for enforcing secure end-to-end data transfer with the key provider service.",
                          "description_kind": "plain",
                          "type": "string"
                        },
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
                        "resource_id": {
                          "computed": true,
                          "description": "The resource ID to include in key requests.",
                          "description_kind": "plain",
                          "type": "string"
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
            "manifest_window_seconds": {
              "computed": true,
              "description": "The time window (in seconds) contained in each manifest.",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "The duration (in seconds) of each segment.",
              "description_kind": "plain",
              "type": "number"
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
          "nesting_mode": "single"
        }
      },
      "origination": {
        "computed": true,
        "description": "Control whether origination of video is allowed for this OriginEndpoint. If set to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of access control. If set to DENY, the OriginEndpoint may not be requested. This can be helpful for Live to VOD harvesting, or for temporarily disabling origination",
        "description_kind": "plain",
        "type": "string"
      },
      "startover_window_seconds": {
        "computed": true,
        "description": "Maximum duration (seconds) of content to retain for startover playback. If not specified, startover playback will be disabled for the OriginEndpoint.",
        "description_kind": "plain",
        "type": "number"
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
      },
      "time_delay_seconds": {
        "computed": true,
        "description": "Amount of delay (seconds) to enforce on the playback of live content. If not specified, there will be no time delay in effect for the OriginEndpoint.",
        "description_kind": "plain",
        "type": "number"
      },
      "url": {
        "computed": true,
        "description": "The URL of the packaged OriginEndpoint for consumption.",
        "description_kind": "plain",
        "type": "string"
      },
      "whitelist": {
        "computed": true,
        "description": "A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::MediaPackage::OriginEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackageOriginEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackageOriginEndpoint), &result)
	return &result
}
