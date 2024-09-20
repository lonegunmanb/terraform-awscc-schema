package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The application Id",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the application.",
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
        "description": "A name for the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata to assign to the application. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key-value string map. The valid character set is [a-zA-Z1-9 +-=._:/-]. The tag key can be up to 128 characters and must not start with aws:.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value can be up to 256 characters.",
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
    "description": "Resource Type definition for AWS::AppConfig::Application",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppconfigApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigApplication), &result)
	return &result
}
