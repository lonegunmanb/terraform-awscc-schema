package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3FilesAccessPoint = `{
  "block": {
    "attributes": {
      "access_point_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "access_point_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "(optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_system_id": {
        "description": "The ID of the S3 Files file system that the access point provides access to.",
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
      "owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "posix_user": {
        "computed": true,
        "description": "The operating system user and group applied to all compute drive requests made using the access point.",
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
        "description": "Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory\u003ePath specified does not exist, EFS creates it and applies the CreationPermissions settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationPermissions is optional.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "creation_permissions": {
              "computed": true,
              "description": "(Optional) Specifies the POSIX IDs and permissions to apply to the access point's RootDirectory. If the RootDirectory\u003ePath specified does not exist, EFS creates the root directory using the CreationPermissions settings when a client connects to an access point. When specifying the CreationPermissions, you must provide values for all properties.   If you do not provide CreationPermissions and the specified RootDirectory\u003ePath does not exist, attempts to mount the file system using the access point will fail. ",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "owner_gid": {
                    "computed": true,
                    "description": "Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "owner_uid": {
                    "computed": true,
                    "description": "Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "permissions": {
                    "computed": true,
                    "description": "Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.",
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
              "description": "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationPermissions.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "Resource Type definition for AWS::S3Files::AccessPoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3FilesAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3FilesAccessPoint), &result)
	return &result
}
