package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationSmb = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "description": "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "domain": {
        "computed": true,
        "description": "The name of the Windows domain that the SMB server belongs to.",
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
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the SMB location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the SMB location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "mount_options": {
        "computed": true,
        "description": "The mount options used by DataSync to access the SMB server.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "version": {
              "computed": true,
              "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "password": {
        "computed": true,
        "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_hostname": {
        "computed": true,
        "description": "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "user": {
        "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::DataSync::LocationSMB.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationSmbSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationSmb), &result)
	return &result
}
