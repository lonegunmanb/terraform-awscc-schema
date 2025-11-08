package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMultiRegionEndpoint = `{
  "block": {
    "attributes": {
      "details": {
        "description": "Contains details of a multi-region endpoint (global-endpoint) being created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "route_details": {
              "description": "A list of route configuration details. Must contain exactly one route configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "region": {
                    "description": "The name of an AWS-Region to be a secondary region for the multi-region endpoint (global-endpoint)",
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
      "endpoint_name": {
        "description": "The name of the multi-region endpoint (global-endpoint).",
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
      "tags": {
        "computed": true,
        "description": "An Array of objects that define the tags (keys and values) to associate with the multi-region endpoint (global-endpoint).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "One part of a key-value pair that defines a tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The optional part of a key-value pair that defines a tag.",
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
    "description": "Resource Type definition for AWS::SES::MultiRegionEndpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesMultiRegionEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMultiRegionEndpoint), &result)
	return &result
}
