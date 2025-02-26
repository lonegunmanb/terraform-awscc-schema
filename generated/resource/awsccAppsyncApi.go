package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncApi = `{
  "block": {
    "attributes": {
      "api_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AppSync Api",
        "description_kind": "plain",
        "type": "string"
      },
      "api_id": {
        "computed": true,
        "description": "The unique identifier for the AppSync Api generated by the service",
        "description_kind": "plain",
        "type": "string"
      },
      "dns": {
        "computed": true,
        "description": "A map of DNS names for the AppSync API.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "http": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "realtime": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "event_config": {
        "computed": true,
        "description": "The configuration for an Event Api",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_providers": {
              "computed": true,
              "description": "A list of auth providers for the AppSync API.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_type": {
                    "computed": true,
                    "description": "Security configuration for your AppSync API.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cognito_config": {
                    "computed": true,
                    "description": "Optional authorization configuration for using Amazon Cognito user pools with your API endpoint.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_id_client_regex": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "aws_region": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user_pool_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lambda_authorizer_config": {
                    "computed": true,
                    "description": "A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode. Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorizer_result_ttl_in_seconds": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "authorizer_uri": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "identity_validation_expression": {
                          "computed": true,
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
                    "description": "The OpenID Connect configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auth_ttl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "client_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iat_ttl": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "issuer": {
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "connection_auth_modes": {
              "computed": true,
              "description": "A list of auth modes for the AppSync API.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_type": {
                    "computed": true,
                    "description": "Security configuration for your AppSync API.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "default_publish_auth_modes": {
              "computed": true,
              "description": "A list of auth modes for the AppSync API.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_type": {
                    "computed": true,
                    "description": "Security configuration for your AppSync API.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "default_subscribe_auth_modes": {
              "computed": true,
              "description": "A list of auth modes for the AppSync API.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_type": {
                    "computed": true,
                    "description": "Security configuration for your AppSync API.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "log_config": {
              "computed": true,
              "description": "The log config for the AppSync API.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_logs_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "log_level": {
                    "computed": true,
                    "description": "Logging level for the AppSync API.",
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
      "name": {
        "description": "The name of the AppSync API.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_contact": {
        "computed": true,
        "description": "The owner contact information for an API resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this AppSync API.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag. You can specify a maximum of 128 characters for a tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AppSync Api",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncApi), &result)
	return &result
}
