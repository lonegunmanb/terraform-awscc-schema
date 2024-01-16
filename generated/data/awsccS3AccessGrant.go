package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3AccessGrant = `{
  "block": {
    "attributes": {
      "access_grant_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified access grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_grant_id": {
        "computed": true,
        "description": "The ID assigned to this access grant.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_grants_location_configuration": {
        "computed": true,
        "description": "The configuration options of the grant location, which is the S3 path to the data to which you are granting access.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_sub_prefix": {
              "computed": true,
              "description": "The S3 sub prefix of a registered location in your S3 Access Grants instance",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "access_grants_location_id": {
        "computed": true,
        "description": "The custom S3 location to be accessed by the grantee",
        "description_kind": "plain",
        "type": "string"
      },
      "application_arn": {
        "computed": true,
        "description": "The ARN of the application grantees will use to access the location",
        "description_kind": "plain",
        "type": "string"
      },
      "grant_scope": {
        "computed": true,
        "description": "The S3 path of the data to which you are granting access. It is a combination of the S3 path of the registered location and the subprefix.",
        "description_kind": "plain",
        "type": "string"
      },
      "grantee": {
        "computed": true,
        "description": "The principal who will be granted permission to access S3.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "grantee_identifier": {
              "computed": true,
              "description": "The unique identifier of the Grantee",
              "description_kind": "plain",
              "type": "string"
            },
            "grantee_type": {
              "computed": true,
              "description": "Configures the transfer acceleration state for an Amazon S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission": {
        "computed": true,
        "description": "The level of access to be afforded to the grantee",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_prefix_type": {
        "computed": true,
        "description": "The type of S3SubPrefix.",
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
      }
    },
    "description": "Data Source schema for AWS::S3::AccessGrant",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3AccessGrantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3AccessGrant), &result)
	return &result
}
