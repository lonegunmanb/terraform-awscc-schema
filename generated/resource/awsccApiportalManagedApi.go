package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalManagedApi = `{
  "block": {
    "attributes": {
      "api_signature": {
        "description": "Identifies a discovered API by its source resource, method, hostname, port, and path.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "hostname": {
              "description": "The front-door hostname of the API.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "method": {
              "description": "The HTTP method of the API.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description": "The route path of the API.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description": "The port the API is served on.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "resource_arn": {
              "description": "ARN of the source-service resource (API Gateway stage, load balancer, or CloudFront distribution).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the managed API was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the managed API was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::ManagedApi. Promotes an API discovered in the API inventory into the managed API list.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalManagedApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalManagedApi), &result)
	return &result
}
