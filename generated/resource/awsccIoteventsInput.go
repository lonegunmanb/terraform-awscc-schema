package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIoteventsInput = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "input_definition": {
        "description": "The definition of the input.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "description": "The attributes from the JSON payload that are made available by the input. Inputs are derived from messages sent to the AWS IoT Events system using ` + "`" + `BatchPutMessage` + "`" + `. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the ` + "`" + `condition` + "`" + ` expressions used by detectors that monitor this input.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "json_path": {
                    "description": "An expression that specifies an attribute-value pair in a JSON structure. Use this to specify an attribute from the JSON payload that is made available by the input. Inputs are derived from messages sent to AWS IoT Events (` + "`" + `BatchPutMessage` + "`" + `). Each such message contains a JSON payload. The attribute (and its paired value) specified here are available for use in the ` + "`" + `condition` + "`" + ` expressions used by detectors.\n\n_Syntax_: ` + "`" + `\u003cfield-name\u003e.\u003cfield-name\u003e...` + "`" + `",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "input_description": {
        "computed": true,
        "description": "A brief description of the input.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "input_name": {
        "computed": true,
        "description": "The name of the input.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Key of the Tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value of the Tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into AWS IoT Events. This is done by sending messages as *inputs* to AWS IoT Events. For more information, see [How to Use AWS IoT Events](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *AWS IoT Events Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIoteventsInputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIoteventsInput), &result)
	return &result
}
