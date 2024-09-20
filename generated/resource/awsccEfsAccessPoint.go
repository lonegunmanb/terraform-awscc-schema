package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEfsAccessPoint = `{
  "block": {
    "attributes": {
      "access_point_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_point_tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key (String). The key can't start with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "The opaque string specified in the request to ensure idempotent creation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_system_id": {
        "description": "The ID of the EFS file system that the access point applies to. Accepts only the ID format for input when specifying a file system, for example ` + "`" + `` + "`" + `fs-0123456789abcedf2` + "`" + `` + "`" + `.",
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
      "posix_user": {
        "computed": true,
        "description": "The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "gid": {
              "computed": true,
              "description": "The POSIX group ID used for all file system operations using this access point.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secondary_gids": {
              "computed": true,
              "description": "Secondary POSIX group IDs used for all file system operations using this access point.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "uid": {
              "computed": true,
              "description": "The POSIX user ID used for all file system operations using this access point.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "root_directory": {
        "computed": true,
        "description": "The directory on the EFS file system that the access point exposes as the root directory to NFS clients using the access point.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "creation_info": {
              "computed": true,
              "description": "(Optional) Specifies the POSIX IDs and permissions to apply to the access point's ` + "`" + `` + "`" + `RootDirectory` + "`" + `` + "`" + `. If the ` + "`" + `` + "`" + `RootDirectory` + "`" + `` + "`" + ` \u003e ` + "`" + `` + "`" + `Path` + "`" + `` + "`" + ` specified does not exist, EFS creates the root directory using the ` + "`" + `` + "`" + `CreationInfo` + "`" + `` + "`" + ` settings when a client connects to an access point. When specifying the ` + "`" + `` + "`" + `CreationInfo` + "`" + `` + "`" + `, you must provide values for all properties. \n  If you do not provide ` + "`" + `` + "`" + `CreationInfo` + "`" + `` + "`" + ` and the specified ` + "`" + `` + "`" + `RootDirectory` + "`" + `` + "`" + ` \u003e ` + "`" + `` + "`" + `Path` + "`" + `` + "`" + ` does not exist, attempts to mount the file system using the access point will fail.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "owner_gid": {
                    "computed": true,
                    "description": "Specifies the POSIX group ID to apply to the ` + "`" + `` + "`" + `RootDirectory` + "`" + `` + "`" + `. Accepts values from 0 to 2^32 (4294967295).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "owner_uid": {
                    "computed": true,
                    "description": "Specifies the POSIX user ID to apply to the ` + "`" + `` + "`" + `RootDirectory` + "`" + `` + "`" + `. Accepts values from 0 to 2^32 (4294967295).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "permissions": {
                    "computed": true,
                    "description": "Specifies the POSIX permissions to apply to the ` + "`" + `` + "`" + `RootDirectory` + "`" + `` + "`" + `, in the format of an octal number representing the file's mode bits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "path": {
              "computed": true,
              "description": "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ` + "`" + `` + "`" + `CreationInfo` + "`" + `` + "`" + `.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::EFS::AccessPoint` + "`" + `` + "`" + ` resource creates an EFS access point. An access point is an application-specific view into an EFS file system that applies an operating system user and group, and a file system path, to any file system request made through the access point. The operating system user and group override any identity information provided by the NFS client. The file system path is exposed as the access point's root directory. Applications using the access point can only access data in its own directory and below. To learn more, see [Mounting a file system using EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html).\n This operation requires permissions for the ` + "`" + `` + "`" + `elasticfilesystem:CreateAccessPoint` + "`" + `` + "`" + ` action.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEfsAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEfsAccessPoint), &result)
	return &result
}
