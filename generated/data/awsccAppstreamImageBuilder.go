package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamImageBuilder = `{
  "block": {
    "attributes": {
      "access_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpce_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "agent_access_config": {
        "computed": true,
        "description": "The configuration for agent access on the image builder. If specified, agent access is enabled for the image builder.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket where agent screenshots are stored. Required when ScreenshotsUploadEnabled is true.",
              "description_kind": "plain",
              "type": "string"
            },
            "screen_image_format": {
              "computed": true,
              "description": "The image format for agent screen captures.",
              "description_kind": "plain",
              "type": "string"
            },
            "screen_resolution": {
              "computed": true,
              "description": "The screen resolution for the agent streaming environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "screenshots_upload_enabled": {
              "computed": true,
              "description": "Indicates whether screenshot uploads to Amazon S3 are enabled for agent sessions.",
              "description_kind": "plain",
              "type": "bool"
            },
            "settings": {
              "computed": true,
              "description": "The list of agent access settings that define permissions for each agent action. You must specify at least one setting. Agent actions that aren't listed are disabled, so list only the actions that you enable.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_action": {
                    "computed": true,
                    "description": "The agent action to configure. Valid values are COMPUTER_VISION, COMPUTER_INPUT, and FORWARD_MCP_TOOLS. COMPUTER_VISION allows agents to take screenshots of the desktop. COMPUTER_INPUT allows agents to click, type, and scroll on the desktop and requires COMPUTER_VISION to also be enabled. FORWARD_MCP_TOOLS allows agents to interact with applications and the desktop operating system through direct MCP calls rather than using computer use tools.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "permission": {
                    "computed": true,
                    "description": "Whether the agent action is enabled or disabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "user_control_mode": {
              "computed": true,
              "description": "The user control mode for agent sessions. This setting determines how users can interact with agent sessions. Valid values are VIEW_ONLY, VIEW_STOP, and DISABLED.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "appstream_agent_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_imdsv1": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_join_info": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "organizational_unit_distinguished_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "enable_default_internet_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "iam_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "image_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_volume_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "volume_size_in_gb": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "softwares_to_install": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "softwares_to_uninstall": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "streaming_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      },
      "vpc_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::AppStream::ImageBuilder",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamImageBuilderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamImageBuilder), &result)
	return &result
}
