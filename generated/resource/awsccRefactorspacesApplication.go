package resource

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
              "optional": true,
              "type": "string"
            },
            "stage_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
      "name": {
        "description_kind": "plain",
        "required": true,
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
        "description_kind": "plain",
        "required": true,
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
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_link_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::RefactorSpaces::Application Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRefactorspacesApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRefactorspacesApplication), &result)
	return &result
}
