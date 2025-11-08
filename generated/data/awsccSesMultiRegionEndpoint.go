package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMultiRegionEndpoint = `{
  "block": {
    "attributes": {
      "details": {
        "computed": true,
        "description": "Contains details of a multi-region endpoint (global-endpoint) being created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "route_details": {
              "computed": true,
              "description": "A list of route configuration details. Must contain exactly one route configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "region": {
                    "computed": true,
                    "description": "The name of an AWS-Region to be a secondary region for the multi-region endpoint (global-endpoint)",
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
      "endpoint_name": {
        "computed": true,
        "description": "The name of the multi-region endpoint (global-endpoint).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The optional part of a key-value pair that defines a tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::SES::MultiRegionEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesMultiRegionEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMultiRegionEndpoint), &result)
	return &result
}
