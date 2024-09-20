package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationFsxLustre = `{
  "block": {
    "attributes": {
      "fsx_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the FSx for Lustre file system.",
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
        "description": "The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the FSx for Lustre location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_arns": {
        "description": "The ARNs of the security groups that are to use to configure the FSx for Lustre file system.",
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
    "description": "Resource schema for AWS::DataSync::LocationFSxLustre.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationFsxLustreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationFsxLustre), &result)
	return &result
}
