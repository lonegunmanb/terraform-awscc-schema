package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreBrowserCustom = `{
  "block": {
    "attributes": {
      "browser_arn": {
        "computed": true,
        "description": "The ARN of a Browser resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "browser_id": {
        "computed": true,
        "description": "The id of the browser.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the browser was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the browser.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role that the browser uses to access resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description": "The reason for failure if the browser creation or operation failed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "Timestamp when the browser was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the browser.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_configuration": {
        "description": "Network configuration for browser.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_mode": {
              "computed": true,
              "description": "Network modes supported by browser",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_config": {
              "computed": true,
              "description": "Network mode configuration for VPC",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_groups": {
                    "computed": true,
                    "description": "Security groups for VPC",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "Subnets for VPC",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "recording_config": {
        "computed": true,
        "description": "Recording configuration for browser.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "s3_location": {
              "computed": true,
              "description": "S3 Location Configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
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
      "status": {
        "computed": true,
        "description": "Status of browser.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource definition for AWS::BedrockAgentCore::BrowserCustom",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreBrowserCustomSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreBrowserCustom), &result)
	return &result
}
