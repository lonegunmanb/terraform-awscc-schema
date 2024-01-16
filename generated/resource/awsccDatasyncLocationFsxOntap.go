package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationFsxOntap = `{
  "block": {
    "attributes": {
      "fsx_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx ONAP file system.",
        "description_kind": "plain",
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
        "description": "The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the FSx ONTAP file system that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocol": {
        "computed": true,
        "description": "Configuration settings for NFS or SMB protocol.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "nfs": {
              "computed": true,
              "description": "NFS protocol configuration for FSx ONTAP file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mount_options": {
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
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "smb": {
              "computed": true,
              "description": "SMB protocol configuration for FSx ONTAP file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "computed": true,
                    "description": "The name of the Windows domain that the SMB server belongs to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mount_options": {
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
                    "required": true
                  },
                  "password": {
                    "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user": {
                    "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
                    "description_kind": "plain",
                    "required": true,
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
      "security_group_arns": {
        "description": "The ARNs of the security groups that are to use to configure the FSx ONTAP file system.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "storage_virtual_machine_arn": {
        "description": "The Amazon Resource Name (ARN) for the FSx ONTAP SVM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
    "description": "Resource schema for AWS::DataSync::LocationFSxONTAP.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationFsxOntapSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationFsxOntap), &result)
	return &result
}
