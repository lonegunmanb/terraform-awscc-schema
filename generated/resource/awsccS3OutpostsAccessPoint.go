package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3OutpostsAccessPoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified AccessPoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket": {
        "description": "The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.",
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
      "name": {
        "description": "A name for the AccessPoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "The access point policy associated with this access point.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_configuration": {
        "description": "Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_id": {
              "computed": true,
              "description": "Virtual Private Cloud (VPC) Id from which AccessPoint will allow requests.",
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
    "description": "Resource Type Definition for AWS::S3Outposts::AccessPoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3OutpostsAccessPointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3OutpostsAccessPoint), &result)
	return &result
}
