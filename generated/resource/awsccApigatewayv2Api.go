package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2Api = `{
  "block": {
    "attributes": {
      "api_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_key_selection_expression": {
        "computed": true,
        "description": "An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "base_path": {
        "computed": true,
        "description": "Specifies how to interpret the base path of the API during import. Valid values are ` + "`" + `` + "`" + `ignore` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `prepend` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `split` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `ignore` + "`" + `` + "`" + `. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "body": {
        "computed": true,
        "description": "The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ` + "`" + `` + "`" + `Body` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `BodyS3Location` + "`" + `` + "`" + `. If you specify a ` + "`" + `` + "`" + `Body` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `BodyS3Location` + "`" + `` + "`" + `, don't specify CloudFormation resources such as ` + "`" + `` + "`" + `AWS::ApiGatewayV2::Authorizer` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `AWS::ApiGatewayV2::Route` + "`" + `` + "`" + `. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "body_s3_location": {
        "computed": true,
        "description": "The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ` + "`" + `` + "`" + `Body` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `BodyS3Location` + "`" + `` + "`" + `. If you specify a ` + "`" + `` + "`" + `Body` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `BodyS3Location` + "`" + `` + "`" + `, don't specify CloudFormation resources such as ` + "`" + `` + "`" + `AWS::ApiGatewayV2::Authorizer` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `AWS::ApiGatewayV2::Route` + "`" + `` + "`" + `. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": "The S3 bucket that contains the OpenAPI definition to import. Required if you specify a ` + "`" + `` + "`" + `BodyS3Location` + "`" + `` + "`" + ` for an API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "etag": {
              "computed": true,
              "description": "The Etag of the S3 object.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "computed": true,
              "description": "The key of the S3 object. Required if you specify a ` + "`" + `` + "`" + `BodyS3Location` + "`" + `` + "`" + ` for an API.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the S3 object.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cors_configuration": {
        "computed": true,
        "description": "A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_credentials": {
              "computed": true,
              "description": "Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_headers": {
              "computed": true,
              "description": "Represents a collection of allowed headers. Supported only for HTTP APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_methods": {
              "computed": true,
              "description": "Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_origins": {
              "computed": true,
              "description": "Represents a collection of allowed origins. Supported only for HTTP APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "expose_headers": {
              "computed": true,
              "description": "Represents a collection of exposed headers. Supported only for HTTP APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_age": {
              "computed": true,
              "description": "The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "credentials_arn": {
        "computed": true,
        "description": "This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ` + "`" + `` + "`" + `arn:aws:iam::*:user/*` + "`" + `` + "`" + `. To use resource-based permissions on supported AWS services, specify ` + "`" + `` + "`" + `null` + "`" + `` + "`" + `. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_execute_api_endpoint": {
        "computed": true,
        "description": "Specifies whether clients can invoke your API by using the default ` + "`" + `` + "`" + `execute-api` + "`" + `` + "`" + ` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "disable_schema_validation": {
        "computed": true,
        "description": "Avoid validating models when creating a deployment. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fail_on_warnings": {
        "computed": true,
        "description": "Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the API. Required unless you specify an OpenAPI definition for ` + "`" + `` + "`" + `Body` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `S3BodyLocation` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "protocol_type": {
        "computed": true,
        "description": "The API protocol. Valid values are ` + "`" + `` + "`" + `WEBSOCKET` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `HTTP` + "`" + `` + "`" + `. Required unless you specify an OpenAPI definition for ` + "`" + `` + "`" + `Body` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `S3BodyLocation` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_key": {
        "computed": true,
        "description": "This property is part of quick create. If you don't specify a ` + "`" + `` + "`" + `routeKey` + "`" + `` + "`" + `, a default route of ` + "`" + `` + "`" + `$default` + "`" + `` + "`" + ` is created. The ` + "`" + `` + "`" + `$default` + "`" + `` + "`" + ` route acts as a catch-all for any request made to your API, for a particular stage. The ` + "`" + `` + "`" + `$default` + "`" + `` + "`" + ` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "route_selection_expression": {
        "computed": true,
        "description": "The route selection expression for the API. For HTTP APIs, the ` + "`" + `` + "`" + `routeSelectionExpression` + "`" + `` + "`" + ` must be ` + "`" + `` + "`" + `${request.method} ${request.path}` + "`" + `` + "`" + `. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target": {
        "computed": true,
        "description": "This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "A version identifier for the API.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGatewayV2::Api` + "`" + `` + "`" + ` resource creates an API. WebSocket APIs and HTTP APIs are supported. For more information about WebSocket APIs, see [About WebSocket APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html) in the *API Gateway Developer Guide*. For more information about HTTP APIs, see [HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html) in the *API Gateway Developer Guide.*",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2ApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Api), &result)
	return &result
}
