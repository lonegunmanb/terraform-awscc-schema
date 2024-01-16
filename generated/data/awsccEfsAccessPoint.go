package data

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
          "nesting_mode": "set"
        }
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "(optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_id": {
        "computed": true,
        "description": "The ID of the EFS file system that the access point provides access to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "posix_user": {
        "computed": true,
        "description": "The operating system user and group applied to all file system requests made using the access point.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "gid": {
              "computed": true,
              "description": "The POSIX group ID used for all file system operations using this access point.",
              "description_kind": "plain",
              "type": "string"
            },
            "secondary_gids": {
              "computed": true,
              "description": "Secondary POSIX group IDs used for all file system operations using this access point.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "uid": {
              "computed": true,
              "description": "The POSIX user ID used for all file system operations using this access point.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "root_directory": {
        "computed": true,
        "description": "Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory\u003ePath specified does not exist, EFS creates it and applies the CreationInfo settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationInfo is optional.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "creation_info": {
              "computed": true,
              "description": "(Optional) Specifies the POSIX IDs and permissions to apply to the access point's RootDirectory. If the RootDirectory\u003ePath specified does not exist, EFS creates the root directory using the CreationInfo settings when a client connects to an access point. When specifying the CreationInfo, you must provide values for all properties.   If you do not provide CreationInfo and the specified RootDirectory\u003ePath does not exist, attempts to mount the file system using the access point will fail. ",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "owner_gid": {
                    "computed": true,
                    "description": "Specifies the POSIX group ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "owner_uid": {
                    "computed": true,
                    "description": "Specifies the POSIX user ID to apply to the RootDirectory. Accepts values from 0 to 2^32 (4294967295).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "permissions": {
                    "computed": true,
                    "description": "Specifies the POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "path": {
              "computed": true,
              "description": "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the CreationInfo.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::EFS::AccessPoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEfsAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEfsAccessPoint), &result)
	return &result
}
