package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2FpgaImage = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the FPGA image.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time the AFI was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_retention_support": {
        "computed": true,
        "description": "Indicates whether data retention support is enabled for the AFI.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description for the AFI.",
        "description_kind": "plain",
        "type": "string"
      },
      "fpga_image_global_id": {
        "computed": true,
        "description": "The global FPGA image identifier (AGFI ID).",
        "description_kind": "plain",
        "type": "string"
      },
      "fpga_image_id": {
        "computed": true,
        "description": "The FPGA image identifier (AFI ID).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_storage_location": {
        "computed": true,
        "description": "The location of the encrypted design checkpoint in Amazon S3. The input must be a tarball.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": "The name of the S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "key": {
              "computed": true,
              "description": "The key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "logs_storage_location": {
        "computed": true,
        "description": "The location in Amazon S3 for the output logs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": "The name of the S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "key": {
              "computed": true,
              "description": "The key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "A name for the AFI.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The ID of the AWS account that owns the AFI.",
        "description_kind": "plain",
        "type": "string"
      },
      "public": {
        "computed": true,
        "description": "Indicates whether the AFI is public.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The state of the AFI (pending | available | failed | unavailable).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the FPGA image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "update_time": {
        "computed": true,
        "description": "The time of the most recent update to the AFI.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::FpgaImage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2FpgaImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2FpgaImage), &result)
	return &result
}
