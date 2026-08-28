package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotStream = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date when the stream was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "files": {
        "computed": true,
        "description": "The files to stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_id": {
              "computed": true,
              "description": "The file ID.",
              "description_kind": "plain",
              "type": "number"
            },
            "s3_location": {
              "computed": true,
              "description": "The location of the file in S3.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description": "The S3 bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description": "The S3 key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "The S3 bucket version.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The date when the stream was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "An IAM role that allows the IoT service principal to access your S3 files.",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_id": {
        "computed": true,
        "description": "The stream ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_version": {
        "computed": true,
        "description": "The stream version.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "Metadata which can be used to manage streams.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IoT::Stream",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotStream), &result)
	return &result
}
