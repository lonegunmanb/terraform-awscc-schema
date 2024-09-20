package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorLiveSource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe ARN of the live source.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "http_package_configurations": {
        "description": "\u003cp\u003eA list of HTTP package configuration parameters for this live source.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "path": {
              "description": "\u003cp\u003eThe relative path to the URL for this VOD source. This is combined with \u003ccode\u003eSourceLocation::HttpConfiguration::BaseUrl\u003c/code\u003e to form a valid URL.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "source_group": {
              "description": "\u003cp\u003eThe name of the source group. This has to match one of the \u003ccode\u003eChannel::Outputs::SourceGroup\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "live_source_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_location_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the live source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::MediaTailor::LiveSource Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediatailorLiveSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorLiveSource), &result)
	return &result
}
