package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalManagedApi = `{
  "block": {
    "attributes": {
      "api_signature": {
        "computed": true,
        "description": "Identifies a discovered API by its source resource, method, hostname, port, and path.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "hostname": {
              "computed": true,
              "description": "The front-door hostname of the API.",
              "description_kind": "plain",
              "type": "string"
            },
            "method": {
              "computed": true,
              "description": "The HTTP method of the API.",
              "description_kind": "plain",
              "type": "string"
            },
            "path": {
              "computed": true,
              "description": "The route path of the API.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port the API is served on.",
              "description_kind": "plain",
              "type": "number"
            },
            "resource_arn": {
              "computed": true,
              "description": "ARN of the source-service resource (API Gateway stage, load balancer, or CloudFront distribution).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the managed API was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_api_arn": {
        "computed": true,
        "description": "The ARN of the managed API.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_api_id": {
        "computed": true,
        "description": "The managed API identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the managed API.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the managed API was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APIPortal::ManagedApi",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApiportalManagedApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalManagedApi), &result)
	return &result
}
