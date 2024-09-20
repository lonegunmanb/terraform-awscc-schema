package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectFlowSource = `{
  "block": {
    "attributes": {
      "decryption": {
        "computed": true,
        "description": "The type of encryption that is used on the content ingested from this source.",
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
        "description": "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entitlement_arn": {
        "computed": true,
        "description": "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flow_arn": {
        "computed": true,
        "description": "The ARN of the flow.",
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
              "computed": true,
              "description": "The ARN of the bridge feeding this flow.",
              "description_kind": "plain",
              "optional": true,
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
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
      "min_latency": {
        "computed": true,
        "description": "The minimum latency in milliseconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "The name of the source.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Resource schema for AWS::MediaConnect::FlowSource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectFlowSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectFlowSource), &result)
	return &result
}
