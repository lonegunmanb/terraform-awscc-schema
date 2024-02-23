package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayMethod = `{
  "block": {
    "attributes": {
      "api_key_required": {
        "computed": true,
        "description": "A boolean flag specifying whether a valid ApiKey is required to invoke this method.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "authorization_scopes": {
        "computed": true,
        "description": "A list of authorization scopes configured on the method. The scopes are used with a ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "authorization_type": {
        "computed": true,
        "description": "The method's authorization type. This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference*.\n  If you specify the ` + "`" + `` + "`" + `AuthorizerId` + "`" + `` + "`" + ` property, specify ` + "`" + `` + "`" + `CUSTOM` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` for this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorizer_id": {
        "computed": true,
        "description": "The identifier of an authorizer to use on this method. The method's authorization type must be ` + "`" + `` + "`" + `CUSTOM` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "http_method": {
        "description": "The method's HTTP verb.",
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
      "integration": {
        "computed": true,
        "description": "Represents an ` + "`" + `` + "`" + `HTTP` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `HTTP_PROXY` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AWS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AWS_PROXY` + "`" + `` + "`" + `, or Mock integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cache_key_parameters": {
              "computed": true,
              "description": "A list of request parameters whose values API Gateway caches. To be valid values for ` + "`" + `` + "`" + `cacheKeyParameters` + "`" + `` + "`" + `, these parameters must also be specified for Method ` + "`" + `` + "`" + `requestParameters` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cache_namespace": {
              "computed": true,
              "description": "Specifies a group of related cached parameters. By default, API Gateway uses the resource ID as the ` + "`" + `` + "`" + `cacheNamespace` + "`" + `` + "`" + `. You can specify the same ` + "`" + `` + "`" + `cacheNamespace` + "`" + `` + "`" + ` across resources to return the same cached data for requests to different resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_id": {
              "computed": true,
              "description": "The ID of the VpcLink used for the integration when ` + "`" + `` + "`" + `connectionType=VPC_LINK` + "`" + `` + "`" + ` and undefined, otherwise.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_type": {
              "computed": true,
              "description": "The type of the network connection to the integration endpoint. The valid value is ` + "`" + `` + "`" + `INTERNET` + "`" + `` + "`" + ` for connections through the public routable internet or ` + "`" + `` + "`" + `VPC_LINK` + "`" + `` + "`" + ` for private connections between API Gateway and a network load balancer in a VPC. The default value is ` + "`" + `` + "`" + `INTERNET` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_handling": {
              "computed": true,
              "description": "Specifies how to handle request payload content type conversions. Supported values are ` + "`" + `` + "`" + `CONVERT_TO_BINARY` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `CONVERT_TO_TEXT` + "`" + `` + "`" + `, with the following behaviors:\n If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the ` + "`" + `` + "`" + `passthroughBehavior` + "`" + `` + "`" + ` is configured to support payload pass-through.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "credentials": {
              "computed": true,
              "description": "Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string ` + "`" + `` + "`" + `arn:aws:iam::\\*:user/\\*` + "`" + `` + "`" + `. To use resource-based permissions on supported AWS services, specify null.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "integration_http_method": {
              "computed": true,
              "description": "Specifies the integration's HTTP method type. For the Type property, if you specify ` + "`" + `` + "`" + `MOCK` + "`" + `` + "`" + `, this property is optional. For Lambda integrations, you must set the integration method to ` + "`" + `` + "`" + `POST` + "`" + `` + "`" + `. For all other types, you must specify this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "integration_responses": {
              "computed": true,
              "description": "Specifies the integration's responses.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_handling": {
                    "computed": true,
                    "description": "Specifies how to handle response payload content type conversions. Supported values are ` + "`" + `` + "`" + `CONVERT_TO_BINARY` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `CONVERT_TO_TEXT` + "`" + `` + "`" + `, with the following behaviors:\n If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "response_parameters": {
                    "computed": true,
                    "description": "A key-value map specifying response parameters that are passed to the method response from the back end. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ` + "`" + `` + "`" + `method.response.header.{name}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` is a valid and unique header name. The mapped non-static value must match the pattern of ` + "`" + `` + "`" + `integration.response.header.{name}` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `integration.response.body.{JSON-expression}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` is a valid and unique response header name and ` + "`" + `` + "`" + `JSON-expression` + "`" + `` + "`" + ` is a valid JSON expression without the ` + "`" + `` + "`" + `$` + "`" + `` + "`" + ` prefix.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "response_templates": {
                    "computed": true,
                    "description": "Specifies the templates used to transform the integration response body. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "selection_pattern": {
                    "computed": true,
                    "description": "Specifies the regular expression (regex) pattern used to choose an integration response based on the response from the back end. For example, if the success response returns nothing and the error response returns some string, you could use the ` + "`" + `` + "`" + `.+` + "`" + `` + "`" + ` regex to match error response. However, make sure that the error response does not contain any newline (` + "`" + `` + "`" + `\\n` + "`" + `` + "`" + `) character in such cases. If the back end is an LAMlong function, the LAMlong function error header is matched. For all other HTTP and AWS back ends, the HTTP status code is matched.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
                    "description": "Specifies the status code that is used to map the integration response to an existing MethodResponse.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "passthrough_behavior": {
              "computed": true,
              "description": "Specifies how the method request body of an unmapped content type will be passed through the integration request to the back end without transformation. A content type is unmapped if no mapping template is defined in the integration or the content type does not match any of the mapped content types, as specified in ` + "`" + `` + "`" + `requestTemplates` + "`" + `` + "`" + `. The valid value is one of the following: ` + "`" + `` + "`" + `WHEN_NO_MATCH` + "`" + `` + "`" + `: passes the method request body through the integration request to the back end without transformation when the method request content type does not match any content type associated with the mapping templates defined in the integration request. ` + "`" + `` + "`" + `WHEN_NO_TEMPLATES` + "`" + `` + "`" + `: passes the method request body through the integration request to the back end without transformation when no mapping template is defined in the integration request. If a template is defined when this option is selected, the method request of an unmapped content-type will be rejected with an HTTP 415 Unsupported Media Type response",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_parameters": {
              "computed": true,
              "description": "A key-value map specifying request parameters that are passed from the method request to the back end. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the back end. The method request parameter value must match the pattern of ` + "`" + `` + "`" + `method.request.{location}.{name}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `location` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `querystring` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `path` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `header` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` must be a valid and unique method request parameter name.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "request_templates": {
              "computed": true,
              "description": "Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "timeout_in_millis": {
              "computed": true,
              "description": "Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "type": {
              "description": "Specifies an API method integration type. The valid value is one of the following:\n For the HTTP and HTTP proxy integrations, each integration can specify a protocol (` + "`" + `` + "`" + `http/https` + "`" + `` + "`" + `), port and path. Standard 80 and 443 ports are supported as well as custom ports above 1024. An HTTP or HTTP proxy integration with a ` + "`" + `` + "`" + `connectionType` + "`" + `` + "`" + ` of ` + "`" + `` + "`" + `VPC_LINK` + "`" + `` + "`" + ` is referred to as a private integration and uses a VpcLink to connect API Gateway to a network load balancer of a VPC.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "Specifies Uniform Resource Identifier (URI) of the integration endpoint.\n For ` + "`" + `` + "`" + `HTTP` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `HTTP_PROXY` + "`" + `` + "`" + ` integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification for standard integrations. If ` + "`" + `` + "`" + `connectionType` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `VPC_LINK` + "`" + `` + "`" + ` specify the Network Load Balancer DNS name. For ` + "`" + `` + "`" + `AWS` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `AWS_PROXY` + "`" + `` + "`" + ` integrations, the URI is of the form ` + "`" + `` + "`" + `arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}` + "`" + `` + "`" + `. Here, {Region} is the API Gateway region (e.g., us-east-1); {service} is the name of the integrated AWS service (e.g., s3); and {subdomain} is a designated subdomain supported by certain AWS service for fast host-name lookup. action can be used for an AWS service action-based API, using an Action={name}\u0026{p1}={v1}\u0026p2={v2}... query string. The ensuing {service_api} refers to a supported action {name} plus any required input parameters. Alternatively, path can be used for an AWS service path-based API. The ensuing service_ap",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "method_responses": {
        "computed": true,
        "description": "Gets a method response associated with a given HTTP status code.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "response_models": {
              "computed": true,
              "description": "Specifies the Model resources used for the response's content-type. Response models are represented as a key/value map, with a content-type as the key and a Model name as the value.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "response_parameters": {
              "computed": true,
              "description": "A key-value map specifying required or optional response parameters that API Gateway can send back to the caller. A key defines a method response header and the value specifies whether the associated method response header is required or not. The expression of the key must match the pattern ` + "`" + `` + "`" + `method.response.header.{name}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` is a valid and unique header name. API Gateway passes certain integration response data to the method response headers specified here according to the mapping you prescribe in the API's IntegrationResponse. The integration response data that can be mapped include an integration response header expressed in ` + "`" + `` + "`" + `integration.response.header.{name}` + "`" + `` + "`" + `, a static value enclosed within a pair of single quotes (e.g., ` + "`" + `` + "`" + `'application/json'` + "`" + `` + "`" + `), or a JSON expression from the back-end response payload in the form of ` + "`" + `` + "`" + `integration.response.body.{JSON-expression}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `JSON-expression` + "`" + `` + "`" + ` is a valid JSON expression without the ` + "`" + `` + "`" + `$` + "`" + `` + "`" + ` prefix.)",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "bool"
              ]
            },
            "status_code": {
              "description": "The method response's status code.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "operation_name": {
        "computed": true,
        "description": "A human-friendly operation identifier for the method. For example, you can assign the ` + "`" + `` + "`" + `operationName` + "`" + `` + "`" + ` of ` + "`" + `` + "`" + `ListPets` + "`" + `` + "`" + ` for the ` + "`" + `` + "`" + `GET /pets` + "`" + `` + "`" + ` method in the ` + "`" + `` + "`" + `PetStore` + "`" + `` + "`" + ` example.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_models": {
        "computed": true,
        "description": "A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "request_parameters": {
        "computed": true,
        "description": "A key-value map defining required or optional method request parameters that can be accepted by API Gateway. A key is a method request parameter name matching the pattern of ` + "`" + `` + "`" + `method.request.{location}.{name}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `location` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `querystring` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `path` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `header` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `) or optional (` + "`" + `` + "`" + `false` + "`" + `` + "`" + `). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "bool"
        ]
      },
      "request_validator_id": {
        "computed": true,
        "description": "The identifier of a RequestValidator for request validation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_id": {
        "description": "The Resource identifier for the MethodResponse resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rest_api_id": {
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Method` + "`" + `` + "`" + ` resource creates API Gateway methods that define the parameters and body that clients must send in their requests.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayMethodSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayMethod), &result)
	return &result
}
