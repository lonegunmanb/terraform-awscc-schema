package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRefactorspacesService = `{
  "block": {
    "attributes": {
      "application_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment_identifier": {
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
      "lambda_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "url_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "health_url": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::RefactorSpaces::Service Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRefactorspacesServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRefactorspacesService), &result)
	return &result
}
