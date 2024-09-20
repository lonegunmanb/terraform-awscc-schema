package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccProtonServiceTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the service template.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "\u003cp\u003eA description of the service template.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "\u003cp\u003eThe name of the service template as displayed in the developer interface.\u003c/p\u003e",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
        "description": "\u003cp\u003eA customer provided encryption key that's used to encrypt data.\u003c/p\u003e",
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pipeline_provisioning": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "\u003cp\u003eAn optional list of metadata items that you can associate with the Proton service template. A tag is a key-value pair.\u003c/p\u003e\n         \u003cp\u003eFor more information, see \u003ca href=\"https://docs.aws.amazon.com/proton/latest/userguide/resources.html\"\u003eProton resources and tagging\u003c/a\u003e in the\n        \u003ci\u003eProton User Guide\u003c/i\u003e.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "\u003cp\u003eThe key of the resource tag.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "\u003cp\u003eThe value of the resource tag.\u003c/p\u003e",
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
    "description": "Definition of AWS::Proton::ServiceTemplate Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccProtonServiceTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccProtonServiceTemplate), &result)
	return &result
}
