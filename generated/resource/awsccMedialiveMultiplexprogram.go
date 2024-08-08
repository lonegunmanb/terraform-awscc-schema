package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMedialiveMultiplexprogram = `{
  "block": {
    "attributes": {
      "channel_id": {
        "computed": true,
        "description": "The MediaLive channel associated with the program.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "multiplex_id": {
        "computed": true,
        "description": "The ID of the multiplex that the program belongs to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multiplex_program_settings": {
        "computed": true,
        "description": "The settings for this multiplex program.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "preferred_channel_pipeline": {
              "computed": true,
              "description": "Indicates which pipeline is preferred by the multiplex for program ingest.\nIf set to \\\"PIPELINE_0\\\" or \\\"PIPELINE_1\\\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,\nit will switch back once that ingest is healthy again. If set to \\\"CURRENTLY_ACTIVE\\\",\nit will not switch back to the other pipeline based on it recovering to a healthy state,\nit will only switch if the active pipeline becomes unhealthy.\n",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "program_number": {
              "description": "Unique program number.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "service_descriptor": {
              "computed": true,
              "description": "Transport stream service descriptor configuration for the Multiplex program.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "provider_name": {
                    "description": "Name of the provider.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "service_name": {
                    "description": "Name of the service.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "video_settings": {
              "computed": true,
              "description": "Program video settings configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constant_bitrate": {
                    "computed": true,
                    "description": "The constant bitrate configuration for the video encode.\nWhen this field is defined, StatmuxSettings must be undefined.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "statmux_settings": {
                    "computed": true,
                    "description": "Statmux rate control settings.\nWhen this field is defined, ConstantBitrate must be undefined.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "maximum_bitrate": {
                          "computed": true,
                          "description": "Maximum statmux bitrate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minimum_bitrate": {
                          "computed": true,
                          "description": "Minimum statmux bitrate.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "priority": {
                          "computed": true,
                          "description": "The purpose of the priority is to use a combination of the\\nmultiplex rate control algorithm and the QVBR capability of the\\nencoder to prioritize the video quality of some channels in a\\nmultiplex over others.  Channels that have a higher priority will\\nget higher video quality at the expense of the video quality of\\nother channels in the multiplex with lower priority.",
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
        "optional": true
      },
      "packet_identifiers_map": {
        "computed": true,
        "description": "The packet identifier map for this multiplex program.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audio_pids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "dvb_sub_pids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "dvb_teletext_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "etv_platform_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "etv_signal_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "klv_data_pids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "pcr_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pmt_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "private_metadata_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scte_27_pids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "scte_35_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timed_metadata_pid": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "video_pid": {
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
      "pipeline_details": {
        "computed": true,
        "description": "Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_channel_pipeline": {
              "computed": true,
              "description": "Identifies the channel pipeline that is currently active for the pipeline (identified by PipelineId) in the multiplex.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pipeline_id": {
              "computed": true,
              "description": "Identifies a specific pipeline in the multiplex.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "preferred_channel_pipeline": {
        "computed": true,
        "description": "The settings for this multiplex program.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "program_name": {
        "computed": true,
        "description": "The name of the multiplex program.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::MediaLive::Multiplexprogram",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMedialiveMultiplexprogramSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMedialiveMultiplexprogram), &result)
	return &result
}
