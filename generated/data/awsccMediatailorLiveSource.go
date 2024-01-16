package data

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
        "computed": true,
        "description": "\u003cp\u003eA list of HTTP package configuration parameters for this live source.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "path": {
              "computed": true,
              "description": "\u003cp\u003eThe relative path to the URL for this VOD source. This is combined with \u003ccode\u003eSourceLocation::HttpConfiguration::BaseUrl\u003c/code\u003e to form a valid URL.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "source_group": {
              "computed": true,
              "description": "\u003cp\u003eThe name of the source group. This has to match one of the \u003ccode\u003eChannel::Outputs::SourceGroup\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
      "live_source_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "source_location_name": {
        "computed": true,
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::MediaTailor::LiveSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediatailorLiveSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorLiveSource), &result)
	return &result
}
