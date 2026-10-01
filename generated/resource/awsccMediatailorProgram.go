package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorProgram = `{
  "block": {
    "attributes": {
      "ad_breaks": {
        "computed": true,
        "description": "The ad break configuration settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ad_break_metadata": {
              "computed": true,
              "description": "Defines a list of key/value pairs that MediaTailor generates within the EXT-X-ASSET tag for SCTE35_ENHANCED output.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "message_type": {
              "computed": true,
              "description": "The SCTE-35 ad insertion type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "offset_millis": {
              "computed": true,
              "description": "How long (in milliseconds) after the beginning of the program that an ad starts.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "slate": {
              "computed": true,
              "description": "Slate VOD source configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_location_name": {
                    "computed": true,
                    "description": "The name of the source location where the slate VOD source is stored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vod_source_name": {
                    "computed": true,
                    "description": "The slate VOD source name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "splice_insert_message": {
              "computed": true,
              "description": "Splice insert message configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "avail_num": {
                    "computed": true,
                    "description": "This is written to splice_insert.avail_num.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "avails_expected": {
                    "computed": true,
                    "description": "This is written to splice_insert.avails_expected.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "splice_event_id": {
                    "computed": true,
                    "description": "This is written to splice_insert.splice_event_id.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "unique_program_id": {
                    "computed": true,
                    "description": "This is written to splice_insert.unique_program_id.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "time_signal_message": {
              "computed": true,
              "description": "The SCTE-35 time_signal message configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "segmentation_descriptors": {
                    "computed": true,
                    "description": "The configurations for the SCTE-35 segmentation_descriptor message(s).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "segment_num": {
                          "computed": true,
                          "description": "The segment number to assign.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "segmentation_event_id": {
                          "computed": true,
                          "description": "The Event Identifier to assign.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "segmentation_type_id": {
                          "computed": true,
                          "description": "The Type Identifier to assign.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "segmentation_upid": {
                          "computed": true,
                          "description": "The Upid to assign.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "segmentation_upid_type": {
                          "computed": true,
                          "description": "The Upid Type to assign.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "segments_expected": {
                          "computed": true,
                          "description": "The number of segments expected.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "sub_segment_num": {
                          "computed": true,
                          "description": "The sub-segment number to assign.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "sub_segments_expected": {
                          "computed": true,
                          "description": "The number of sub-segments expected.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the program.",
        "description_kind": "plain",
        "type": "string"
      },
      "audience_media": {
        "computed": true,
        "description": "The list of AudienceMedia defined in program.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alternate_media": {
              "computed": true,
              "description": "The list of AlternateMedia defined in AudienceMedia.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ad_breaks": {
                    "computed": true,
                    "description": "Ad break configuration parameters defined in AlternateMedia.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ad_break_metadata": {
                          "computed": true,
                          "description": "Defines a list of key/value pairs that MediaTailor generates within the EXT-X-ASSET tag for SCTE35_ENHANCED output.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description": "The key.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "message_type": {
                          "computed": true,
                          "description": "The SCTE-35 ad insertion type.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "offset_millis": {
                          "computed": true,
                          "description": "How long (in milliseconds) after the beginning of the program that an ad starts.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "slate": {
                          "computed": true,
                          "description": "Slate VOD source configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source_location_name": {
                                "computed": true,
                                "description": "The name of the source location where the slate VOD source is stored.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "vod_source_name": {
                                "computed": true,
                                "description": "The slate VOD source name.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "splice_insert_message": {
                          "computed": true,
                          "description": "Splice insert message configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "avail_num": {
                                "computed": true,
                                "description": "This is written to splice_insert.avail_num.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "avails_expected": {
                                "computed": true,
                                "description": "This is written to splice_insert.avails_expected.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "splice_event_id": {
                                "computed": true,
                                "description": "This is written to splice_insert.splice_event_id.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "unique_program_id": {
                                "computed": true,
                                "description": "This is written to splice_insert.unique_program_id.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "time_signal_message": {
                          "computed": true,
                          "description": "The SCTE-35 time_signal message configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "segmentation_descriptors": {
                                "computed": true,
                                "description": "The configurations for the SCTE-35 segmentation_descriptor message(s).",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "segment_num": {
                                      "computed": true,
                                      "description": "The segment number to assign.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "segmentation_event_id": {
                                      "computed": true,
                                      "description": "The Event Identifier to assign.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "segmentation_type_id": {
                                      "computed": true,
                                      "description": "The Type Identifier to assign.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "segmentation_upid": {
                                      "computed": true,
                                      "description": "The Upid to assign.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "segmentation_upid_type": {
                                      "computed": true,
                                      "description": "The Upid Type to assign.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "segments_expected": {
                                      "computed": true,
                                      "description": "The number of segments expected.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "sub_segment_num": {
                                      "computed": true,
                                      "description": "The sub-segment number to assign.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "sub_segments_expected": {
                                      "computed": true,
                                      "description": "The number of sub-segments expected.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
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
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "clip_range": {
                    "computed": true,
                    "description": "Clip range configuration for the VOD source associated with the program.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_offset_millis": {
                          "computed": true,
                          "description": "The end offset of the clip range, in milliseconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "start_offset_millis": {
                          "computed": true,
                          "description": "The start offset of the clip range, in milliseconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "duration_millis": {
                    "computed": true,
                    "description": "The duration of the alternateMedia in milliseconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "live_source_name": {
                    "computed": true,
                    "description": "The name of the live source for alternateMedia.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scheduled_start_time_millis": {
                    "computed": true,
                    "description": "The date and time that the alternateMedia is scheduled to start, in epoch milliseconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "source_location_name": {
                    "computed": true,
                    "description": "The name of the source location for alternateMedia.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vod_source_name": {
                    "computed": true,
                    "description": "The name of the VOD source for alternateMedia.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "audience": {
              "computed": true,
              "description": "The Audience defined in AudienceMedia.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "channel_name": {
        "description": "The name of the channel for this Program.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "clip_range": {
        "computed": true,
        "description": "The clip range configuration settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_offset_millis": {
              "computed": true,
              "description": "The end offset of the clip range, in milliseconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "start_offset_millis": {
              "computed": true,
              "description": "The start offset of the clip range, in milliseconds.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "creation_time": {
        "computed": true,
        "description": "The timestamp of when the program was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "duration_millis": {
        "computed": true,
        "description": "The duration of the live program in milliseconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "live_source_name": {
        "computed": true,
        "description": "The name of the LiveSource for this Program.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "program_name": {
        "description": "The name of the Program.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_configuration": {
        "computed": true,
        "description": "The schedule configuration settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "clip_range": {
              "computed": true,
              "description": "Clip range configuration for the VOD source associated with the program.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_offset_millis": {
                    "computed": true,
                    "description": "The end offset of the clip range, in milliseconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_offset_millis": {
                    "computed": true,
                    "description": "The start offset of the clip range, in milliseconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "transition": {
              "computed": true,
              "description": "Program transition configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "duration_millis": {
                    "computed": true,
                    "description": "The duration of the live program in seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "relative_position": {
                    "computed": true,
                    "description": "The position where this program will be inserted relative to the RelativePosition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relative_program": {
                    "computed": true,
                    "description": "The name of the program that this program will be inserted next to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scheduled_start_time_millis": {
                    "computed": true,
                    "description": "The date and time that the program is scheduled to start, in epoch milliseconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "Defines when the program plays in the schedule. You can set the value to ABSOLUTE or RELATIVE.",
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
      "scheduled_start_time": {
        "computed": true,
        "description": "The date and time that the program is scheduled to start.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_location_name": {
        "description": "The name of the source location.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vod_source_name": {
        "computed": true,
        "description": "The name that's used to refer to a VOD source.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::MediaTailor::Program",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediatailorProgramSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorProgram), &result)
	return &result
}
