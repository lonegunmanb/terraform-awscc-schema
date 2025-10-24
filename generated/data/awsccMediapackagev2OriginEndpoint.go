package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackagev2OriginEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) associated with the resource.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_group_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "channel_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "container_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the origin endpoint was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "dash_manifest_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "dash_manifests": {
        "computed": true,
        "description": "\u003cp\u003eA DASH manifest configuration.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_urls": {
              "computed": true,
              "description": "\u003cp\u003eThe base URL to use for retrieving segments.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dvb_priority": {
                    "computed": true,
                    "description": "\u003cp\u003eFor use with DVB-DASH profiles only. The priority of this location for servings segments. The lower the number, the higher the priority.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "dvb_weight": {
                    "computed": true,
                    "description": "\u003cp\u003eFor use with DVB-DASH profiles only. The weighting for source locations that have the same priority. \u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "service_location": {
                    "computed": true,
                    "description": "\u003cp\u003eThe name of the source location.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "url": {
                    "computed": true,
                    "description": "\u003cp\u003eA source location for segments.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "compactness": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "drm_signaling": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dvb_settings": {
              "computed": true,
              "description": "\u003cp\u003eFor endpoints that use the DVB-DASH profile only. The font download and error reporting information that you want MediaPackage to pass through to the manifest.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "error_metrics": {
                    "computed": true,
                    "description": "\u003cp\u003ePlayback device error reporting settings.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "probability": {
                          "computed": true,
                          "description": "\u003cp\u003eThe number of playback devices per 1000 that will send error reports to the reporting URL. This represents the probability that a playback device will be a reporting player for this session.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "reporting_url": {
                          "computed": true,
                          "description": "\u003cp\u003eThe URL where playback devices send error reports.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "font_download": {
                    "computed": true,
                    "description": "\u003cp\u003eFor use with DVB-DASH profiles only. The settings for font downloads that you want Elemental MediaPackage to pass through to the manifest.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "font_family": {
                          "computed": true,
                          "description": "\u003cp\u003eThe \u003ccode\u003efontFamily\u003c/code\u003e name for subtitles, as described in \u003ca href=\"https://tech.ebu.ch/publications/tech3380\"\u003eEBU-TT-D Subtitling Distribution Format\u003c/a\u003e. \u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "mime_type": {
                          "computed": true,
                          "description": "\u003cp\u003eThe \u003ccode\u003emimeType\u003c/code\u003e of the resource that's at the font download URL.\u003c/p\u003e \u003cp\u003eFor information about font MIME types, see the \u003ca href=\"https://dvb.org/wp-content/uploads/2021/06/A168r4_MPEG-DASH-Profile-for-Transport-of-ISO-BMFF-Based-DVB-Services_Draft-ts_103-285-v140_November_2021.pdf\"\u003eMPEG-DASH Profile for Transport of ISO BMFF Based DVB Services over IP Based Networks\u003c/a\u003e document. \u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "url": {
                          "computed": true,
                          "description": "\u003cp\u003eThe URL for downloading fonts for subtitles.\u003c/p\u003e",
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
            "filter_configuration": {
              "computed": true,
              "description": "\u003cp\u003eFilter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest. \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "clip_start_time": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the clip start time for all of your manifest egress requests. When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "end": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "manifest_filter": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "time_delay_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "manifest_name": {
              "computed": true,
              "description": "\u003cp\u003eA short string that's appended to the endpoint URL. The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, index. \u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_window_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe total duration (in seconds) of the manifest's content.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "min_buffer_time_seconds": {
              "computed": true,
              "description": "\u003cp\u003eMinimum amount of content (in seconds) that a player must keep available in the buffer.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "min_update_period_seconds": {
              "computed": true,
              "description": "\u003cp\u003eMinimum amount of time (in seconds) that the player should wait before requesting updates to the manifest.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "period_triggers": {
              "computed": true,
              "description": "\u003cp\u003eA list of triggers that controls when AWS Elemental MediaPackage separates the MPEG-DASH manifest into multiple periods. Leave this value empty to indicate that the manifest is contained all in one period. For more information about periods in the DASH manifest, see \u003ca href=\"https://docs.aws.amazon.com/mediapackage/latest/userguide/multi-period.html\"\u003eMulti-period DASH in AWS Elemental MediaPackage\u003c/a\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "profiles": {
              "computed": true,
              "description": "\u003cp\u003eThe profile that the output is compliant with.\u003c/p\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "program_information": {
              "computed": true,
              "description": "\u003cp\u003eDetails about the content that you want MediaPackage to pass through in the manifest to the playback device.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "copyright": {
                    "computed": true,
                    "description": "\u003cp\u003eA copyright statement about the content.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "language_code": {
                    "computed": true,
                    "description": "\u003cp\u003eThe language code for this manifest.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "more_information_url": {
                    "computed": true,
                    "description": "\u003cp\u003eAn absolute URL that contains more information about this content.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description": "\u003cp\u003eInformation about the content provider.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "title": {
                    "computed": true,
                    "description": "\u003cp\u003eThe title for the manifest.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "scte_dash": {
              "computed": true,
              "description": "\u003cp\u003eThe SCTE configuration.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_marker_dash": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "segment_template_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subtitle_configuration": {
              "computed": true,
              "description": "\u003cp\u003eThe configuration for DASH subtitles.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ttml_configuration": {
                    "computed": true,
                    "description": "\u003cp\u003eThe settings for TTML subtitles.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ttml_profile": {
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
            "suggested_presentation_delay_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe amount of time (in seconds) that the player should be from the end of the manifest.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "utc_timing": {
              "computed": true,
              "description": "\u003cp\u003eDetermines the type of UTC timing included in the DASH Media Presentation Description (MPD).\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "timing_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timing_source": {
                    "computed": true,
                    "description": "\u003cp\u003eThe the method that the player uses to synchronize to coordinated universal time (UTC) wall clock time.\u003c/p\u003e",
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
      "description": {
        "computed": true,
        "description": "\u003cp\u003eEnter any descriptive text that helps you to identify the origin endpoint.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "force_endpoint_error_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe failover settings for the endpoint.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_error_conditions": {
              "computed": true,
              "description": "\u003cp\u003eThe failover conditions for the endpoint. The options are:\u003c/p\u003e \u003cul\u003e \u003cli\u003e \u003cp\u003e \u003ccode\u003eSTALE_MANIFEST\u003c/code\u003e - The manifest stalled and there are no new segments or parts.\u003c/p\u003e \u003c/li\u003e \u003cli\u003e \u003cp\u003e \u003ccode\u003eINCOMPLETE_MANIFEST\u003c/code\u003e - There is a gap in the manifest.\u003c/p\u003e \u003c/li\u003e \u003cli\u003e \u003cp\u003e \u003ccode\u003eMISSING_DRM_KEY\u003c/code\u003e - Key rotation is enabled but we're unable to fetch the key for the current key period.\u003c/p\u003e \u003c/li\u003e \u003cli\u003e \u003cp\u003e \u003ccode\u003eSLATE_INPUT\u003c/code\u003e - The segments which contain slate content are considered to be missing content.\u003c/p\u003e \u003c/li\u003e \u003c/ul\u003e",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "hls_manifest_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "hls_manifests": {
        "computed": true,
        "description": "\u003cp\u003eAn HTTP live streaming (HLS) manifest configuration.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "child_manifest_name": {
              "computed": true,
              "description": "\u003cp\u003eA short string that's appended to the endpoint URL. The child manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default child manifest name, index_1. The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "filter_configuration": {
              "computed": true,
              "description": "\u003cp\u003eFilter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest. \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "clip_start_time": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the clip start time for all of your manifest egress requests. When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "end": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "manifest_filter": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "time_delay_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "manifest_name": {
              "computed": true,
              "description": "\u003cp\u003eA short short string that's appended to the endpoint URL. The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, index. MediaPackage automatically inserts the format extension, such as .m3u8. You can't use the same manifest name if you use HLS manifest and low-latency HLS manifest. The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_window_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe total duration (in seconds) of the manifest's content.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "program_date_time_interval_seconds": {
              "computed": true,
              "description": "\u003cp\u003eInserts EXT-X-PROGRAM-DATE-TIME tags in the output manifest at the interval that you specify. If you don't enter an interval, EXT-X-PROGRAM-DATE-TIME tags aren't included in the manifest. The tags sync the stream to the wall clock so that viewers can seek to a specific time in the playback timeline on the player.\u003c/p\u003e \u003cp\u003eIrrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "scte_hls": {
              "computed": true,
              "description": "\u003cp\u003eThe SCTE configuration.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_marker_hls": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "start_tag": {
              "computed": true,
              "description": "\u003cp\u003eTo insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset. When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "precise": {
                    "computed": true,
                    "description": "\u003cp\u003eSpecify the value for PRECISE within your EXT-X-START tag. Leave blank, or choose false, to use the default value NO. Choose yes to use the value YES.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "time_offset": {
                    "computed": true,
                    "description": "\u003cp\u003eSpecify the value for TIME-OFFSET within your EXT-X-START tag. Enter a signed floating point value which, if positive, must be less than the configured manifest duration minus three times the configured segment target duration. If negative, the absolute value must be larger than three times the configured segment target duration, and the absolute value must be smaller than the configured manifest duration.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "url": {
              "computed": true,
              "description": "\u003cp\u003eThe egress domain URL for stream delivery from MediaPackage.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "url_encode_child_manifest": {
              "computed": true,
              "description": "\u003cp\u003eWhen enabled, MediaPackage URL-encodes the query string for API requests for HLS child manifests to comply with Amazon Web Services Signature Version 4 (SigV4) signature signing protocol. For more information, see \u003ca href=\"https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html\"\u003eAmazon Web Services Signature Version 4 for API requests\u003c/a\u003e in \u003ci\u003eIdentity and Access Management User Guide\u003c/i\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "bool"
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
      "low_latency_hls_manifest_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "low_latency_hls_manifests": {
        "computed": true,
        "description": "\u003cp\u003eA low-latency HLS manifest configuration.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "child_manifest_name": {
              "computed": true,
              "description": "\u003cp\u003eA short string that's appended to the endpoint URL. The child manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default child manifest name, index_1. The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "filter_configuration": {
              "computed": true,
              "description": "\u003cp\u003eFilter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest. \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "clip_start_time": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the clip start time for all of your manifest egress requests. When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "end": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "manifest_filter": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "time_delay_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "manifest_name": {
              "computed": true,
              "description": "\u003cp\u003eA short short string that's appended to the endpoint URL. The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, index. MediaPackage automatically inserts the format extension, such as .m3u8. You can't use the same manifest name if you use HLS manifest and low-latency HLS manifest. The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_window_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe total duration (in seconds) of the manifest's content.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "program_date_time_interval_seconds": {
              "computed": true,
              "description": "\u003cp\u003eInserts EXT-X-PROGRAM-DATE-TIME tags in the output manifest at the interval that you specify. If you don't enter an interval, EXT-X-PROGRAM-DATE-TIME tags aren't included in the manifest. The tags sync the stream to the wall clock so that viewers can seek to a specific time in the playback timeline on the player.\u003c/p\u003e \u003cp\u003eIrrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "scte_hls": {
              "computed": true,
              "description": "\u003cp\u003eThe SCTE configuration.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_marker_hls": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "start_tag": {
              "computed": true,
              "description": "\u003cp\u003eTo insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset. When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "precise": {
                    "computed": true,
                    "description": "\u003cp\u003eSpecify the value for PRECISE within your EXT-X-START tag. Leave blank, or choose false, to use the default value NO. Choose yes to use the value YES.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "time_offset": {
                    "computed": true,
                    "description": "\u003cp\u003eSpecify the value for TIME-OFFSET within your EXT-X-START tag. Enter a signed floating point value which, if positive, must be less than the configured manifest duration minus three times the configured segment target duration. If negative, the absolute value must be larger than three times the configured segment target duration, and the absolute value must be smaller than the configured manifest duration.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "url": {
              "computed": true,
              "description": "\u003cp\u003eThe egress domain URL for stream delivery from MediaPackage.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "url_encode_child_manifest": {
              "computed": true,
              "description": "\u003cp\u003eWhen enabled, MediaPackage URL-encodes the query string for API requests for LL-HLS child manifests to comply with Amazon Web Services Signature Version 4 (SigV4) signature signing protocol. For more information, see \u003ca href=\"https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html\"\u003eAmazon Web Services Signature Version 4 for API requests\u003c/a\u003e in \u003ci\u003eIdentity and Access Management User Guide\u003c/i\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "list"
        }
      },
      "modified_at": {
        "computed": true,
        "description": "\u003cp\u003eThe date and time the origin endpoint was modified.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "mss_manifest_urls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "mss_manifests": {
        "computed": true,
        "description": "\u003cp\u003eThe Microsoft Smooth Streaming (MSS) manifest configurations associated with this origin endpoint.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filter_configuration": {
              "computed": true,
              "description": "\u003cp\u003eFilter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest. \u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "clip_start_time": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the clip start time for all of your manifest egress requests. When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "end": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "manifest_filter": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "time_delay_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eOptionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "manifest_layout": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_name": {
              "computed": true,
              "description": "\u003cp\u003eThe name of the MSS manifest. This name is appended to the origin endpoint URL to create the unique path for accessing this specific MSS manifest.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "manifest_window_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe duration (in seconds) of the manifest window. This represents the total amount of content available in the manifest at any given time.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "origin_endpoint_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "segment": {
        "computed": true,
        "description": "\u003cp\u003eThe segment configuration, including the segment name, duration, and other configuration values.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for encrypting content.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cmaf_exclude_segment_drm_metadata": {
                    "computed": true,
                    "description": "\u003cp\u003eExcludes SEIG and SGPD boxes from segment metadata in CMAF containers.\u003c/p\u003e \u003cp\u003eWhen set to \u003ccode\u003etrue\u003c/code\u003e, MediaPackage omits these DRM metadata boxes from CMAF segments, which can improve compatibility with certain devices and players that don't support these boxes.\u003c/p\u003e \u003cp\u003eImportant considerations:\u003c/p\u003e \u003cul\u003e \u003cli\u003e \u003cp\u003eThis setting only affects CMAF container formats\u003c/p\u003e \u003c/li\u003e \u003cli\u003e \u003cp\u003eKey rotation can still be handled through media playlist signaling\u003c/p\u003e \u003c/li\u003e \u003cli\u003e \u003cp\u003ePSSH and TENC boxes remain unaffected\u003c/p\u003e \u003c/li\u003e \u003cli\u003e \u003cp\u003eDefault behavior is preserved when this setting is disabled\u003c/p\u003e \u003c/li\u003e \u003c/ul\u003e \u003cp\u003eValid values: \u003ccode\u003etrue\u003c/code\u003e | \u003ccode\u003efalse\u003c/code\u003e \u003c/p\u003e \u003cp\u003eDefault: \u003ccode\u003efalse\u003c/code\u003e \u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "constant_initialization_vector": {
                    "computed": true,
                    "description": "\u003cp\u003eA 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting content. If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "encryption_method": {
                    "computed": true,
                    "description": "\u003cp\u003eThe encryption type.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cmaf_encryption_method": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "ism_encryption_method": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "ts_encryption_method": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "key_rotation_interval_seconds": {
                    "computed": true,
                    "description": "\u003cp\u003eThe frequency (in seconds) of key changes for live workflows, in which content is streamed real time. The service retrieves content keys before the live content begins streaming, and then retrieves them as needed over the lifetime of the workflow. By default, key rotation is set to 300 seconds (5 minutes), the minimum rotation interval, which is equivalent to setting it to 300. If you don't enter an interval, content keys aren't rotated.\u003c/p\u003e \u003cp\u003eThe following example setting causes the service to rotate keys every thirty minutes: \u003ccode\u003e1800\u003c/code\u003e \u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "speke_key_provider": {
                    "computed": true,
                    "description": "\u003cp\u003eThe parameters for the SPEKE key provider.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "drm_systems": {
                          "computed": true,
                          "description": "\u003cp\u003eThe DRM solution provider you're using to protect your content during distribution.\u003c/p\u003e",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "encryption_contract_configuration": {
                          "computed": true,
                          "description": "\u003cp\u003eConfigure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines which content keys are used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "preset_speke_20_audio": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "preset_speke_20_video": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "resource_id": {
                          "computed": true,
                          "description": "\u003cp\u003eThe unique identifier for the content. The service sends this to the key server to identify the current endpoint. How unique you make this depends on how fine-grained you want access controls to be. The service does not permit you to use the same ID for two simultaneous encryption processes. The resource ID is also known as the content ID.\u003c/p\u003e \u003cp\u003eThe following example shows a resource ID: \u003ccode\u003eMovieNight20171126093045\u003c/code\u003e \u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "\u003cp\u003eThe ARN for the IAM role granted by the key provider that provides access to the key provider API. This role must have a trust policy that allows MediaPackage to assume the role, and it must have a sufficient permissions policy to allow access to the specific key retrieval URL. Get this from your DRM solution provider.\u003c/p\u003e \u003cp\u003eValid format: \u003ccode\u003earn:aws:iam::{accountID}:role/{name}\u003c/code\u003e. The following example shows a role ARN: \u003ccode\u003earn:aws:iam::444455556666:role/SpekeAccess\u003c/code\u003e \u003c/p\u003e",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "url": {
                          "computed": true,
                          "description": "\u003cp\u003eThe URL of the API Gateway proxy that you set up to talk to your key server. The API Gateway proxy must reside in the same AWS Region as MediaPackage and must start with https://.\u003c/p\u003e \u003cp\u003eThe following example shows a URL: \u003ccode\u003ehttps://1wm2dx1f33.execute-api.us-west-2.amazonaws.com/SpekeSample/copyProtection\u003c/code\u003e \u003c/p\u003e",
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
            "include_iframe_only_streams": {
              "computed": true,
              "description": "\u003cp\u003eWhen selected, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included. MediaPackage generates an I-frame only stream from the first rendition in the manifest. The service inserts EXT-I-FRAMES-ONLY tags in the output manifest, and then generates and includes an I-frames only playlist in the stream. This playlist permits player functionality like fast forward and rewind.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "bool"
            },
            "scte": {
              "computed": true,
              "description": "\u003cp\u003eThe SCTE configuration.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "scte_filter": {
                    "computed": true,
                    "description": "\u003cp\u003eThe SCTE-35 message types that you want to be treated as ad markers in the output.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "segment_duration_seconds": {
              "computed": true,
              "description": "\u003cp\u003eThe duration (in seconds) of each segment. Enter a value equal to, or a multiple of, the input segment duration. If the value that you enter is different from the input segment duration, MediaPackage rounds segments to the nearest multiple of the input segment duration.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_name": {
              "computed": true,
              "description": "\u003cp\u003eThe name that describes the segment. The name is the base name of the segment used in all content manifests inside of the endpoint. You can't use spaces in the name.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "ts_include_dvb_subtitles": {
              "computed": true,
              "description": "\u003cp\u003eBy default, MediaPackage excludes all digital video broadcasting (DVB) subtitles from the output. When selected, MediaPackage passes through DVB subtitles into the output.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "bool"
            },
            "ts_use_audio_rendition_group": {
              "computed": true,
              "description": "\u003cp\u003eWhen selected, MediaPackage bundles all audio tracks in a rendition group. All other tracks in the stream can be used with any audio rendition from the group.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "startover_window_seconds": {
        "computed": true,
        "description": "\u003cp\u003eThe size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).\u003c/p\u003e",
        "description_kind": "plain",
        "type": "number"
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::MediaPackageV2::OriginEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediapackagev2OriginEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackagev2OriginEndpoint), &result)
	return &result
}
