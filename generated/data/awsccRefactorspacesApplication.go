package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRefactorspacesApplication = `{
  "block": {
    "attributes": {
      "api_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_gateway_proxy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stage_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "application_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "nlb_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "nlb_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "proxy_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "proxy_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage_name": {
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
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_link_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RefactorSpaces::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRefactorspacesApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRefactorspacesApplication), &result)
	return &result
}
