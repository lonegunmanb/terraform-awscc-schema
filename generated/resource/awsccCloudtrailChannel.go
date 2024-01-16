package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudtrailChannel = `{
  "block": {
    "attributes": {
      "channel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of a channel.",
        "description_kind": "plain",
        "type": "string"
      },
      "destinations": {
        "computed": true,
        "description": "One or more resources to which events arriving through a channel are logged and stored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "location": {
              "description": "The ARN of a resource that receives events from a channel.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "The type of destination for events arriving from a channel.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the channel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "The ARN of an on-premises storage solution or application, or a partner event source.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "A channel receives events from a specific source (such as an on-premises storage solution or application, or a partner event data source), and delivers the events to one or more event data stores. You use channels to ingest events into CloudTrail from sources outside AWS.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudtrailChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudtrailChannel), &result)
	return &result
}
