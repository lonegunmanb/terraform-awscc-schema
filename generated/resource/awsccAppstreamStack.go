package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamStack = `{
  "block": {
    "attributes": {
      "access_endpoints": {
        "computed": true,
        "description": "The list of virtual private cloud (VPC) interface endpoint objects. Users of the stack can connect to AppStream 2.0 only through the specified endpoints.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_type": {
              "computed": true,
              "description": "The type of interface endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpce_id": {
              "computed": true,
              "description": "The identifier (ID) of the VPC in which the interface endpoint is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "application_settings": {
        "computed": true,
        "description": "The persistent application settings for users of the stack. When these settings are enabled, changes that users make to applications and Windows settings are automatically saved after each session and applied to the next session.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Enables or disables persistent application settings for users during their streaming sessions.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "settings_group": {
              "computed": true,
              "description": "The path prefix for the S3 bucket where users? persistent application settings are stored. You can allow the same persistent application settings to be used across multiple stacks by specifying the same settings group for each stack.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "attributes_to_delete": {
        "computed": true,
        "description": "The stack attributes to delete.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "delete_storage_connectors": {
        "computed": true,
        "description": "This parameter has been deprecated. Deletes the storage connectors currently enabled for the stack.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "The description to display.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The stack name to display.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "embed_host_domains": {
        "computed": true,
        "description": "The domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "feedback_url": {
        "computed": true,
        "description": "The URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the stack.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redirect_url": {
        "computed": true,
        "description": "The URL that users are redirected to after their streaming session ends.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_connectors": {
        "computed": true,
        "description": "The storage connectors to enable.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connector_type": {
              "computed": true,
              "description": "The type of storage connector.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "domains": {
              "computed": true,
              "description": "The names of the domains for the account.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_identifier": {
              "computed": true,
              "description": "The ARN of the storage connector.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "streaming_experience_settings": {
        "computed": true,
        "description": "The streaming protocol that you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "preferred_protocol": {
              "computed": true,
              "description": "The preferred protocol that you want to use while streaming your application.",
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
        "description": "An array of key-value pairs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "user_settings": {
        "computed": true,
        "description": "The actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description": "The action that is enabled or disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maximum_length": {
              "computed": true,
              "description": "Specifies the number of characters that can be copied by end users from the local device to the remote session, and to the local device from the remote session. This can be specified only for the CLIPBOARD_COPY_FROM_LOCAL_DEVICE and CLIPBOARD_COPY_TO_LOCAL_DEVICE actions. This defaults to 20,971,520 (20 MB) when unspecified and the permission is ENABLED. This can't be specified when the permission is DISABLED. The value can be between 1 and 20,971,520 (20 MB).",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "permission": {
              "computed": true,
              "description": "Indicates whether the action is enabled or disabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::AppStream::Stack",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamStackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamStack), &result)
	return &result
}
