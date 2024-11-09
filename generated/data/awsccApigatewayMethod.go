package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayMethod = `{
  "block": {
    "attributes": {
      "api_key_required": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "authorization_scopes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "authorization_type": {
        "computed": true,
        "description": "The method's authorization type. This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference*.\n  If you specify the ` + "`" + `` + "`" + `AuthorizerId` + "`" + `` + "`" + ` property, specify ` + "`" + `` + "`" + `CUSTOM` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` for this property.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_method": {
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
      "integration": {
        "computed": true,
        "description": "` + "`" + `` + "`" + `Integration` + "`" + `` + "`" + ` is a property of the [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html) resource that specifies information about the target backend that a method calls.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cache_key_parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "cache_namespace": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "connection_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "connection_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "content_handling": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "credentials": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "integration_http_method": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "integration_responses": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_handling": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "response_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "response_templates": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "selection_pattern": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "passthrough_behavior": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "request_parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "request_templates": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "timeout_in_millis": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "method_responses": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "response_models": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "response_parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "bool"
              ]
            },
            "status_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "operation_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "request_models": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "request_parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "bool"
        ]
      },
      "request_validator_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::Method",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayMethodSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayMethod), &result)
	return &result
}
