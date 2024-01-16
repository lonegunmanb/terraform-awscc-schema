package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediapackageChannel = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) assigned to the Channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A short text description of the Channel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "egress_access_logs": {
        "computed": true,
        "description": "The configuration parameters for egress access logging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_name": {
              "computed": true,
              "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "hls_ingest": {
        "computed": true,
        "description": "An HTTP Live Streaming (HLS) ingest resource configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ingest_endpoints": {
              "computed": true,
              "description": "A list of endpoints to which the source stream should be sent.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "description": "The system generated unique identifier for the IngestEndpoint",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "password": {
                    "description": "The system generated password for ingest authentication.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "url": {
                    "description": "The ingest URL to which the source stream should be sent.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "The system generated username for ingest authentication.",
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
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "description": "The ID of the Channel.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ingress_access_logs": {
        "computed": true,
        "description": "The configuration parameters for egress access logging.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_name": {
              "computed": true,
              "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
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
    "description": "Resource schema for AWS::MediaPackage::Channel",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediapackageChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediapackageChannel), &result)
	return &result
}
