package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ExpressDirectoryBucket = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Returns the Amazon Resource Name (ARN) of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_name": {
        "computed": true,
        "description": "Returns the code for the Availability Zone where the directory bucket was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_encryption": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_name": {
        "computed": true,
        "description": "Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_redundancy": {
        "computed": true,
        "description": "Specifies the number of Availability Zone that's used for redundancy for the bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location_name": {
        "computed": true,
        "description": "Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Express::DirectoryBucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3ExpressDirectoryBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ExpressDirectoryBucket), &result)
	return &result
}
