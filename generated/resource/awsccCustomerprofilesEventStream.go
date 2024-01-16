package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesEventStream = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the export was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_details": {
        "computed": true,
        "description": "Details regarding the Kinesis stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "The status of enabling the Kinesis stream as a destination for export.",
              "description_kind": "plain",
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "The StreamARN of the destination to deliver profile events to. For example, arn:aws:kinesis:region:account-id:stream/stream-name",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_name": {
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_stream_arn": {
        "computed": true,
        "description": "A unique identifier for the event stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_stream_name": {
        "description": "The name of the event stream.",
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
      "state": {
        "computed": true,
        "description": "The operational state of destination stream for export.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "uri": {
        "description": "The StreamARN of the destination to deliver profile events to. For example, arn:aws:kinesis:region:account-id:stream/stream-name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "An Event Stream resource of Amazon Connect Customer Profiles",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCustomerprofilesEventStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesEventStream), &result)
	return &result
}
