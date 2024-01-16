package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationNfs = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the NFS location.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the NFS location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "mount_options": {
        "computed": true,
        "description": "The NFS mount options that DataSync can use to mount your NFS share.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "version": {
              "computed": true,
              "description": "The specific NFS version that you want DataSync to use to mount your NFS share.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "on_prem_config": {
        "description": "Contains a list of Amazon Resource Names (ARNs) of agents that are used to connect an NFS server.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "agent_arns": {
              "description": "ARN(s) of the agent(s) to use for an NFS location.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "server_hostname": {
        "computed": true,
        "description": "The name of the NFS server. This value is the IP address or DNS name of the NFS server.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.",
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
      }
    },
    "description": "Resource schema for AWS::DataSync::LocationNFS",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationNfsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationNfs), &result)
	return &result
}
