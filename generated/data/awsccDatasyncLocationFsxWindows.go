package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationFsxWindows = `{
  "block": {
    "attributes": {
      "domain": {
        "computed": true,
        "description": "The name of the Windows domain that the FSx for Windows server belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "fsx_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx for Windows file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the FSx for Windows location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "The password of the user who has the permissions to access files and folders in the FSx for Windows file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_arns": {
        "computed": true,
        "description": "The ARNs of the security groups that are to use to configure the FSx for Windows file system.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subdirectory": {
        "computed": true,
        "description": "A subdirectory in the location's path.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "user": {
        "computed": true,
        "description": "The user who has the permissions to access files and folders in the FSx for Windows file system.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataSync::LocationFSxWindows",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationFsxWindowsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationFsxWindows), &result)
	return &result
}
