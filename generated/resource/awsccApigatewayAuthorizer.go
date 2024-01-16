package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayAuthorizer = `{
  "block": {
    "attributes": {
      "auth_type": {
        "computed": true,
        "description": "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorizer_credentials": {
        "computed": true,
        "description": "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorizer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_result_ttl_in_seconds": {
        "computed": true,
        "description": "The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "authorizer_uri": {
        "computed": true,
        "description": "Specifies the authorizer's Uniform Resource Identifier (URI). For ` + "`" + `` + "`" + `TOKEN` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `REQUEST` + "`" + `` + "`" + ` authorizers, this must be a well-formed Lambda function URI, for example, ` + "`" + `` + "`" + `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations` + "`" + `` + "`" + `. In general, the URI has this form ` + "`" + `` + "`" + `arn:aws:apigateway:{region}:lambda:path/{service_api}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `{region}` + "`" + `` + "`" + ` is the same as the region hosting the Lambda function, ` + "`" + `` + "`" + `path` + "`" + `` + "`" + ` indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ` + "`" + `` + "`" + `/` + "`" + `` + "`" + `. For Lambda functions, this is usually of the form ` + "`" + `` + "`" + `/2015-03-31/functions/[FunctionARN]/invocations` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_source": {
        "computed": true,
        "description": "The identity source for which authorization is requested. For a ` + "`" + `` + "`" + `TOKEN` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` authorizer, this is required and specifies the request header mapping expression for the custom header holding the authorization token submitted by the client. For example, if the token header name is ` + "`" + `` + "`" + `Auth` + "`" + `` + "`" + `, the header mapping expression is ` + "`" + `` + "`" + `method.request.header.Auth` + "`" + `` + "`" + `. For the ` + "`" + `` + "`" + `REQUEST` + "`" + `` + "`" + ` authorizer, this is required when authorization caching is enabled. The value is a comma-separated string of one or more mapping expressions of the specified request parameters. For example, if an ` + "`" + `` + "`" + `Auth` + "`" + `` + "`" + ` header, a ` + "`" + `` + "`" + `Name` + "`" + `` + "`" + ` query string parameter are defined as identity sources, this value is ` + "`" + `` + "`" + `method.request.header.Auth, method.request.querystring.Name` + "`" + `` + "`" + `. These parameters will be used to derive the authorization caching key and to perform runtime validation of the ` + "`" + `` + "`" + `REQUEST` + "`" + `` + "`" + ` authorizer by verifying all of the identity-related request parameters are present, not null and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized response without calling the Lambda function. The valid value is a string of comma-separated mapping expressions of the specified request parameters. When the authorization caching is not enabled, this property is optional.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_validation_expression": {
        "computed": true,
        "description": "A validation expression for the incoming identity token. For ` + "`" + `` + "`" + `TOKEN` + "`" + `` + "`" + ` authorizers, this value is a regular expression. For ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` authorizers, API Gateway will match the ` + "`" + `` + "`" + `aud` + "`" + `` + "`" + ` field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the ` + "`" + `` + "`" + `REQUEST` + "`" + `` + "`" + ` authorizer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the authorizer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_ar_ns": {
        "computed": true,
        "description": "A list of the Amazon Cognito user pool ARNs for the ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` authorizer. Each element is of this format: ` + "`" + `` + "`" + `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}` + "`" + `` + "`" + `. For a ` + "`" + `` + "`" + `TOKEN` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `REQUEST` + "`" + `` + "`" + ` authorizer, this is not defined.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "rest_api_id": {
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The authorizer type. Valid values are ` + "`" + `` + "`" + `TOKEN` + "`" + `` + "`" + ` for a Lambda function using a single authorization token submitted in a custom header, ` + "`" + `` + "`" + `REQUEST` + "`" + `` + "`" + ` for a Lambda function using incoming request parameters, and ` + "`" + `` + "`" + `COGNITO_USER_POOLS` + "`" + `` + "`" + ` for using an Amazon Cognito user pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Authorizer` + "`" + `` + "`" + ` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayAuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayAuthorizer), &result)
	return &result
}
