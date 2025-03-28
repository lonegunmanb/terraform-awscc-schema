package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectFlowOutput = `{
  "block": {
    "attributes": {
      "cidr_allow_list": {
        "computed": true,
        "description": "The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "A description of the output.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination": {
        "computed": true,
        "description": "The address where you want to send the output.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption": {
        "computed": true,
        "description": "The type of key used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
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
            "key_type": {
              "computed": true,
              "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "flow_arn": {
        "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_latency": {
        "computed": true,
        "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "media_stream_output_configurations": {
        "computed": true,
        "description": "The definition for each media stream that is associated with the output.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_configurations": {
              "computed": true,
              "description": "The media streams that you want to associate with the output.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_ip": {
                    "computed": true,
                    "description": "The IP address where contents of the media stream will be sent.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "destination_port": {
                    "computed": true,
                    "description": "The port to use when the content of the media stream is distributed to the output.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "interface": {
                    "computed": true,
                    "description": "The VPC interface that is used for the media stream associated with the output.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the VPC interface that you want to use for the media stream associated with the output.",
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "encoding_name": {
              "computed": true,
              "description": "The format that will be used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video streams on sources or outputs that use the CDI protocol, set the encoding name to raw. For video streams on sources or outputs that use the ST 2110 JPEG XS protocol, set the encoding name to jxsv.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encoding_parameters": {
              "computed": true,
              "description": "A collection of parameters that determine how MediaConnect will convert the content. These fields only apply to outputs on flows that have a CDI source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "compression_factor": {
                    "computed": true,
                    "description": "A value that is used to calculate compression for an output. The bitrate of the output is calculated as follows: Output bitrate = (1 / compressionFactor) * (source bitrate) This property only applies to outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol. Valid values are in the range of 3.0 to 10.0, inclusive.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "encoder_profile": {
                    "computed": true,
                    "description": "A setting on the encoder that drives compression settings. This property only applies to video media streams associated with outputs that use the ST 2110 JPEG XS protocol, with a flow source that uses the CDI protocol.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "media_stream_name": {
              "computed": true,
              "description": "A name that helps you distinguish one media stream from another.",
              "description_kind": "plain",
              "optional": true,
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
        "description": "The name of the output. This value must be unique within the current flow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ndi_program_name": {
        "computed": true,
        "description": "A suffix for the names of the NDI sources that the flow creates. If a custom name isn't specified, MediaConnect uses the output name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ndi_speed_hq_quality": {
        "computed": true,
        "description": "A quality setting for the NDI Speed HQ encoder.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "output_arn": {
        "computed": true,
        "description": "The ARN of the output.",
        "description_kind": "plain",
        "type": "string"
      },
      "output_status": {
        "computed": true,
        "description": "An indication of whether the output should transmit data or not.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port to use when content is distributed to this output.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "description": "The protocol that is used by the source or output.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remote_id": {
        "computed": true,
        "description": "The remote ID for the Zixi-pull stream.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "smoothing_latency": {
        "computed": true,
        "description": "The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.",
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
      "vpc_interface_attachment": {
        "computed": true,
        "description": "The name of the VPC interface attachment to use for this output.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_interface_name": {
              "computed": true,
              "description": "The name of the VPC interface to use for this output.",
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
    "description": "Resource schema for AWS::MediaConnect::FlowOutput",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectFlowOutputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectFlowOutput), &result)
	return &result
}
