package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3FilesFileSystemPolicy = `{
  "block": {
    "attributes": {
      "file_system_id": {
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
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3Files::FileSystemPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3FilesFileSystemPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3FilesFileSystemPolicy), &result)
	return &result
}
