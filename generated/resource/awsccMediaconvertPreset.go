package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconvertPreset = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the output preset, such as arn:aws:mediaconvert:us-west-2:123456789012",
        "description_kind": "plain",
        "type": "string"
      },
      "category": {
        "computed": true,
        "description": "The new category for the preset, if you are changing it.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The new description for the preset, if you are changing it.",
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
        "computed": true,
        "description": "The name of the preset that you are modifying.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "settings_json": {
        "description": "Specify, in JSON format, the transcoding job settings for this output preset. This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::MediaConvert::Preset",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconvertPresetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconvertPreset), &result)
	return &result
}
