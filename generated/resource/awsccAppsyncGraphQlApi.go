package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncGraphQlApi = `{
  "block": {
    "attributes": {
      "additional_authentication_providers": {
        "computed": true,
        "description": "A list of additional authentication providers for the GraphqlApi API.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_type": {
              "computed": true,
              "description": "The authentication type for API key, AWS Identity and Access Management, OIDC, Amazon Cognito user pools, or AWS Lambda.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lambda_authorizer_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorizer_result_ttl_in_seconds": {
                    "computed": true,
                    "description": "The number of seconds a response should be cached for.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "authorizer_uri": {
                    "computed": true,
                    "description": "The ARN of the Lambda function to be called for authorization.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "identity_validation_expression": {
                    "computed": true,
                    "description": "A regular expression for validation of tokens before the Lambda function is called.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "open_id_connect_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_ttl": {
                    "computed": true,
                    "description": "The number of milliseconds that a token is valid after being authenticated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "client_id": {
                    "computed": true,
                    "description": "The client identifier of the Relying party at the OpenID identity provider.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iat_ttl": {
                    "computed": true,
                    "description": "The number of milliseconds that a token is valid after it's issued to a user.\n\n",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "issuer": {
                    "computed": true,
                    "description": "The issuer for the OIDC configuration. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "user_pool_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_id_client_regex": {
                    "computed": true,
                    "description": "A regular expression for validating the incoming Amazon Cognito user pool app client ID. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "aws_region": {
                    "computed": true,
                    "description": "The AWS Region in which the user pool was created.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_pool_id": {
                    "computed": true,
                    "description": "The user pool ID",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "api_id": {
        "computed": true,
        "description": "Unique AWS AppSync GraphQL API identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "api_type": {
        "computed": true,
        "description": "The value that indicates whether the GraphQL API is a standard API (GRAPHQL) or merged API (MERGED).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the API key",
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_type": {
        "description": "Security configuration for your GraphQL API",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enhanced_metrics_config": {
        "computed": true,
        "description": "Enables and controls the enhanced metrics feature. Enhanced metrics emit granular data on API usage and performance such as AppSync request and error counts, latency, and cache hits/misses. All enhanced metric data is sent to your CloudWatch account, and you can configure the types of data that will be sent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_source_level_metrics_behavior": {
              "computed": true,
              "description": "Controls how data source metrics will be emitted to CloudWatch. Data source metrics include:\n\n",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "operation_level_metrics_config": {
              "computed": true,
              "description": "Controls how operation metrics will be emitted to CloudWatch. Operation metrics include:\n\n",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resolver_level_metrics_behavior": {
              "computed": true,
              "description": "Controls how resolver metrics will be emitted to CloudWatch. Resolver metrics include:\n\n",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "environment_variables": {
        "computed": true,
        "description": "A map containing the list of resources with their properties and environment variables.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "graph_ql_dns": {
        "computed": true,
        "description": "The fully qualified domain name (FQDN) of the endpoint URL of your GraphQL API.",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_ql_endpoint_arn": {
        "computed": true,
        "description": "The GraphQL endpoint ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_ql_url": {
        "computed": true,
        "description": "The Endpoint URL of your GraphQL API.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "introspection_config": {
        "computed": true,
        "description": "Sets the value of the GraphQL API to enable (ENABLED) or disable (DISABLED) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lambda_authorizer_config": {
        "computed": true,
        "description": "A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode. Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorizer_result_ttl_in_seconds": {
              "computed": true,
              "description": "The number of seconds a response should be cached for.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "authorizer_uri": {
              "computed": true,
              "description": "The ARN of the Lambda function to be called for authorization.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identity_validation_expression": {
              "computed": true,
              "description": "A regular expression for validation of tokens before the Lambda function is called.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "log_config": {
        "computed": true,
        "description": "The Amazon CloudWatch Logs configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_role_arn": {
              "computed": true,
              "description": "The service role that AWS AppSync will assume to publish to Amazon CloudWatch Logs in your account.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "exclude_verbose_content": {
              "computed": true,
              "description": "Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging level.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "field_log_level": {
              "computed": true,
              "description": "The field logging level. Values can be NONE, ERROR, INFO, DEBUG, or ALL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "merged_api_execution_role_arn": {
        "computed": true,
        "description": "The AWS Identity and Access Management service role ARN for a merged API. ",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The API name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "open_id_connect_config": {
        "computed": true,
        "description": "The OpenID Connect configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_ttl": {
              "computed": true,
              "description": "The number of milliseconds that a token is valid after being authenticated.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "client_id": {
              "computed": true,
              "description": "The client identifier of the Relying party at the OpenID identity provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "iat_ttl": {
              "computed": true,
              "description": "The number of milliseconds that a token is valid after it's issued to a user.\n\n",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "issuer": {
              "computed": true,
              "description": "The issuer for the OIDC configuration. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "owner_contact": {
        "computed": true,
        "description": "The owner contact information for an API resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_depth_limit": {
        "computed": true,
        "description": "The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "realtime_dns": {
        "computed": true,
        "description": "The fully qualified domain name (FQDN) of the real-time endpoint URL of your GraphQL API.",
        "description_kind": "plain",
        "type": "string"
      },
      "realtime_url": {
        "computed": true,
        "description": "The GraphQL API real-time endpoint URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "resolver_count_limit": {
        "computed": true,
        "description": "The maximum number of resolvers that can be invoked in a single request.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this GraphQL API.\n\n",
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
      "user_pool_config": {
        "computed": true,
        "description": "Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint.\n\n",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_id_client_regex": {
              "computed": true,
              "description": "A regular expression for validating the incoming Amazon Cognito user pool app client ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "aws_region": {
              "computed": true,
              "description": "The AWS Region in which the user pool was created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_action": {
              "computed": true,
              "description": "The action that you want your GraphQL API to take when a request that uses Amazon Cognito user pool authentication doesn't match the Amazon Cognito user pool configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_pool_id": {
              "computed": true,
              "description": "The user pool ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "visibility": {
        "computed": true,
        "description": "Sets the scope of the GraphQL API to public (GLOBAL) or private (PRIVATE). By default, the scope is set to Global if no value is provided.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "xray_enabled": {
        "computed": true,
        "description": "A flag indicating whether to use AWS X-Ray tracing for this GraphqlApi.\n\n",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource Type definition for AWS::AppSync::GraphQLApi",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncGraphQlApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncGraphQlApi), &result)
	return &result
}
