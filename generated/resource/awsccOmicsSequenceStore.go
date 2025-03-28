package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOmicsSequenceStore = `{
  "block": {
    "attributes": {
      "access_log_location": {
        "computed": true,
        "description": "Location of the access logs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The store's ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "When the store was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the store.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "e_tag_algorithm_family": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fallback_location": {
        "computed": true,
        "description": "An S3 location that is used to store files that have failed a direct upload.",
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
      "name": {
        "description": "A name for the store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "propagated_set_level_tags": {
        "computed": true,
        "description": "The tags keys to propagate to the S3 objects associated with read sets in the sequence store.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "s3_access_point_arn": {
        "computed": true,
        "description": "This is ARN of the access point associated with the S3 bucket storing read sets.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_access_policy": {
        "computed": true,
        "description": "The resource policy that controls S3 access on the store",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_uri": {
        "computed": true,
        "description": "The S3 URI of the sequence store.",
        "description_kind": "plain",
        "type": "string"
      },
      "sequence_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sse_config": {
        "computed": true,
        "description": "Server-side encryption (SSE) settings for a store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_arn": {
              "computed": true,
              "description": "An encryption key ARN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
        "description": "The status message of the sequence store.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The last-updated time of the sequence store.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Omics::SequenceStore",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOmicsSequenceStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOmicsSequenceStore), &result)
	return &result
}
