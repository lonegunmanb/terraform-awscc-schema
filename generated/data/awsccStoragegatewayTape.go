package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccStoragegatewayTape = `{
  "block": {
    "attributes": {
      "gateway_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Tape Gateway that hosts the virtual tape.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_encrypted": {
        "computed": true,
        "description": "Set to true to use Amazon S3 server-side encryption with your own KMS key, or false to use a key managed by Amazon S3. Optional.",
        "description_kind": "plain",
        "type": "bool"
      },
      "kms_key": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption. This value must be set if KMSEncrypted is true.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description": "The ID of the pool that you want to add your tape to for archiving. Tapes in this pool are archived in the S3 storage class that is associated with the pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of up to 50 tags to assign to the virtual tape.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Cannot be prefixed with aws:.",
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
          "nesting_mode": "set"
        }
      },
      "tape_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the virtual tape.",
        "description_kind": "plain",
        "type": "string"
      },
      "tape_barcode": {
        "computed": true,
        "description": "The barcode that you want to assign to the virtual tape. Barcodes cannot be reused, even after a tape is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "tape_created_date": {
        "computed": true,
        "description": "The date and time that the virtual tape was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "tape_size_in_bytes": {
        "computed": true,
        "description": "The size, in bytes, of the virtual tape that you want to create.",
        "description_kind": "plain",
        "type": "number"
      },
      "tape_status": {
        "computed": true,
        "description": "The current status of the virtual tape.",
        "description_kind": "plain",
        "type": "string"
      },
      "tape_used_in_bytes": {
        "computed": true,
        "description": "The size, in bytes, of data stored on the virtual tape.",
        "description_kind": "plain",
        "type": "number"
      },
      "worm": {
        "computed": true,
        "description": "Set to true to create a write-once-read-many (WORM) virtual tape.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::StorageGateway::Tape",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccStoragegatewayTapeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccStoragegatewayTape), &result)
	return &result
}
