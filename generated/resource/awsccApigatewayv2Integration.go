package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2Integration = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The API identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_id": {
        "computed": true,
        "description": "The ID of the VPC link for a private integration. Supported only for HTTP APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_type": {
        "computed": true,
        "description": "The type of the network connection to the integration endpoint. Specify INTERNET for connections through the public routable internet or VPC_LINK for private connections between API Gateway and resources in a VPC. The default value is INTERNET.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_handling_strategy": {
        "computed": true,
        "description": "Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "credentials_arn": {
        "computed": true,
        "description": "Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS services, don't specify this parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the integration.",
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
      "integration_id": {
        "computed": true,
        "description": "The integration ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_method": {
        "computed": true,
        "description": "Specifies the integration's HTTP method type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_subtype": {
        "computed": true,
        "description": "Supported only for HTTP API AWS_PROXY integrations. Specifies the AWS service action to invoke.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integration_type": {
        "description": "The integration type of an integration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_uri": {
        "computed": true,
        "description": "For a Lambda integration, specify the URI of a Lambda function. For an HTTP integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "passthrough_behavior": {
        "computed": true,
        "description": "Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource. There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "payload_format_version": {
        "computed": true,
        "description": "Specifies the format of the payload sent to an integration. Required for HTTP APIs. For HTTP APIs, supported values for Lambda proxy integrations are 1.0 and 2.0 For all other integrations, 1.0 is the only supported value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_parameters": {
        "computed": true,
        "description": "A key-value map specifying parameters.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "request_templates": {
        "computed": true,
        "description": "A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "response_parameters": {
        "computed": true,
        "description": "Parameters that transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "response_parameters": {
              "computed": true,
              "description": "list of response parameters",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "map"
        },
        "optional": true
      },
      "template_selection_expression": {
        "computed": true,
        "description": "The template selection expression for the integration. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout_in_millis": {
        "computed": true,
        "description": "Custom timeout between 50 and 29000 milliseconds for WebSocket APIs and between 50 and 30000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tls_config": {
        "computed": true,
        "description": "The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS protocol. Supported only for HTTP APIs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_name_to_verify": {
              "computed": true,
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
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2IntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Integration), &result)
	return &result
}
