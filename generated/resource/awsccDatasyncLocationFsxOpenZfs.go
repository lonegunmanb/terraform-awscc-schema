package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationFsxOpenZfs = `{
  "block": {
    "attributes": {
      "fsx_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx OpenZFS file system.",
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
        "description": "The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the FSx OpenZFS that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "description": "Configuration settings for an NFS or SMB protocol, currently only support NFS",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "nfs": {
              "computed": true,
              "description": "FSx OpenZFS file system NFS protocol information",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
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
      "security_group_arns": {
        "description": "The ARNs of the security groups that are to use to configure the FSx OpenZFS file system.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subdirectory": {
        "computed": true,
        "description": "A subdirectory in the location's path.",
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
              "computed": true,
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::DataSync::LocationFSxOpenZFS.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationFsxOpenZfsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationFsxOpenZfs), &result)
	return &result
}
