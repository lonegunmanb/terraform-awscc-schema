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
              "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_arn": {
              "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
              "description_kind": "plain",
              "required": true,
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
      "output_arn": {
        "computed": true,
        "description": "The ARN of the output.",
        "description_kind": "plain",
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
