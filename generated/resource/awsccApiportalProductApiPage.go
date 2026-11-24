package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalProductApiPage = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The route (Path + Method) on the source-service resource, resolved from the Managed API's signature. One ResourceArn covers API Gateway, ALB, and CloudFront.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier_parts": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "method": {
                    "computed": true,
                    "description": "The HTTP method of the route.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "The route path within the source-service resource.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "ARN of the source-service resource (API Gateway RestApi+Stage, ALB, or CloudFront distribution).",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "display_content": {
        "computed": true,
        "description": "Three-mode content: write with either Overrides or None; the service populates ApiUrl, OperationName and Body on Read.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_url": {
              "computed": true,
              "description": "Service-rendered API URL. Read-only.",
              "description_kind": "plain",
              "type": "string"
            },
            "body": {
              "computed": true,
              "description": "Service-rendered documentation body. Read-only.",
              "description_kind": "plain",
              "type": "string"
            },
            "none": {
              "computed": true,
              "description": "Marker selecting the default service-rendered content with no overrides.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "operation_name": {
              "computed": true,
              "description": "Service-rendered operation name. Read-only.",
              "description_kind": "plain",
              "type": "string"
            },
            "overrides": {
              "computed": true,
              "description": "Customer-supplied overrides for the service-rendered API documentation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_url": {
                    "computed": true,
                    "description": "Override for the API URL.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "body": {
                    "computed": true,
                    "description": "Override for the documentation body.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operation_name": {
                    "computed": true,
                    "description": "Override for the operation name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_api_id": {
        "computed": true,
        "description": "Identifier of the Managed API this page documents.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "product_api_page_arn": {
        "computed": true,
        "description": "The ARN of the product API page.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_api_page_id": {
        "computed": true,
        "description": "The product API page identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "description": "The identifier of the parent Product.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Provisioning status of the API documentation.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_exception": {
        "computed": true,
        "description": "Reported when Status is FAILED.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exception": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the product API page.",
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
      "try_it_state": {
        "computed": true,
        "description": "Whether the interactive TryIt UI is enabled for this API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the product API page was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::ProductApiPage. Per-API documentation page attached to an AWS::APIPortal::Product.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalProductApiPageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalProductApiPage), &result)
	return &result
}
