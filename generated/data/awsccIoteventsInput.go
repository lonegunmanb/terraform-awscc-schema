package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIoteventsInput = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_definition": {
        "computed": true,
        "description": "The definition of the input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the ITE system using ` + "`" + `` + "`" + `BatchPutMessage` + "`" + `` + "`" + `. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` expressions used by detectors that monitor this input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "json_path": {
                    "computed": true,
                    "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to ITE (` + "`" + `` + "`" + `BatchPutMessage` + "`" + `` + "`" + `). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` expressions used by detectors. \n Syntax: ` + "`" + `` + "`" + `\u003cfield-name\u003e.\u003cfield-name\u003e...` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "input_description": {
        "computed": true,
        "description": "A brief description of the input.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_name": {
        "computed": true,
        "description": "The name of the input.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
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
    "description": "Data Source schema for AWS::IoTEvents::Input",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIoteventsInputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIoteventsInput), &result)
	return &result
}
