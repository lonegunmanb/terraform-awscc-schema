package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticloadbalancingv2ListenerRule = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "The actions.\n The rule must include exactly one of the following types of actions: ` + "`" + `` + "`" + `forward` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `fixed-response` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `redirect` + "`" + `` + "`" + `, and it must be the last action to be performed. If the rule is for an HTTPS listener, it can also optionally include an authentication action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authenticate_cognito_config": {
              "computed": true,
              "description": "[HTTPS listeners] Information for using Amazon Cognito to authenticate users. Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `authenticate-cognito` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "computed": true,
                    "description": "The query parameters (up to 10) to include in the redirect request to the authorization endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description": "The behavior if the user is not authenticated. The following are possible values:\n  +  deny` + "`" + `` + "`" + `` + "`" + `` + "`" + ` - Return an HTTP 401 Unauthorized error.\n  +  allow` + "`" + `` + "`" + `` + "`" + `` + "`" + ` - Allow the request to be forwarded to the target.\n  +  authenticate` + "`" + `` + "`" + `` + "`" + `` + "`" + ` - Redirect the request to the IdP authorization endpoint. This is the default value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description": "The set of user claims to be requested from the IdP. The default is ` + "`" + `` + "`" + `openid` + "`" + `` + "`" + `.\n To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description": "The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description": "The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "user_pool_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the Amazon Cognito user pool.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_pool_client_id": {
                    "computed": true,
                    "description": "The ID of the Amazon Cognito user pool client.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_pool_domain": {
                    "computed": true,
                    "description": "The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "authenticate_oidc_config": {
              "computed": true,
              "description": "[HTTPS listeners] Information about an identity provider that is compliant with OpenID Connect (OIDC). Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `authenticate-oidc` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_request_extra_params": {
                    "computed": true,
                    "description": "The query parameters (up to 10) to include in the redirect request to the authorization endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "authorization_endpoint": {
                    "computed": true,
                    "description": "The authorization endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_id": {
                    "computed": true,
                    "description": "The OAuth 2.0 client identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description": "The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you are modifying a rule, you can omit this parameter if you set ` + "`" + `` + "`" + `UseExistingClientSecret` + "`" + `` + "`" + ` to true.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "issuer": {
                    "computed": true,
                    "description": "The OIDC issuer identifier of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "on_unauthenticated_request": {
                    "computed": true,
                    "description": "The behavior if the user is not authenticated. The following are possible values:\n  +  deny` + "`" + `` + "`" + `` + "`" + `` + "`" + ` - Return an HTTP 401 Unauthorized error.\n  +  allow` + "`" + `` + "`" + `` + "`" + `` + "`" + ` - Allow the request to be forwarded to the target.\n  +  authenticate` + "`" + `` + "`" + `` + "`" + `` + "`" + ` - Redirect the request to the IdP authorization endpoint. This is the default value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scope": {
                    "computed": true,
                    "description": "The set of user claims to be requested from the IdP. The default is ` + "`" + `` + "`" + `openid` + "`" + `` + "`" + `.\n To verify which scope values your IdP supports and how to separate multiple values, see the documentation for your IdP.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_cookie_name": {
                    "computed": true,
                    "description": "The name of the cookie used to maintain session information. The default is AWSELBAuthSessionCookie.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout": {
                    "computed": true,
                    "description": "The maximum duration of the authentication session, in seconds. The default is 604800 seconds (7 days).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "token_endpoint": {
                    "computed": true,
                    "description": "The token endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "use_existing_client_secret": {
                    "computed": true,
                    "description": "Indicates whether to use the existing client secret when modifying a rule. If you are creating a rule, you can omit this parameter or set it to false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "user_info_endpoint": {
                    "computed": true,
                    "description": "The user info endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the domain, and the path.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "fixed_response_config": {
              "computed": true,
              "description": "[Application Load Balancer] Information for creating an action that returns a custom HTTP response. Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `fixed-response` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_type": {
                    "computed": true,
                    "description": "The content type.\n Valid Values: text/plain | text/css | text/html | application/javascript | application/json",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "message_body": {
                    "computed": true,
                    "description": "The message.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
                    "computed": true,
                    "description": "The HTTP response code (2XX, 4XX, or 5XX).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "forward_config": {
              "computed": true,
              "description": "Information for creating an action that distributes requests among multiple target groups. Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `forward` + "`" + `` + "`" + `.\n If you specify both ` + "`" + `` + "`" + `ForwardConfig` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `TargetGroupArn` + "`" + `` + "`" + `, you can specify only one target group using ` + "`" + `` + "`" + `ForwardConfig` + "`" + `` + "`" + ` and it must be the same target group specified in ` + "`" + `` + "`" + `TargetGroupArn` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_group_stickiness_config": {
                    "computed": true,
                    "description": "Information about the target group stickiness for a rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "duration_seconds": {
                          "computed": true,
                          "description": "[Application Load Balancers] The time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days). You must specify this value when enabling target group stickiness.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Indicates whether target group stickiness is enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "target_groups": {
                    "computed": true,
                    "description": "Information about how traffic will be distributed between multiple target groups in a forward rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_group_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the target group.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weight": {
                          "computed": true,
                          "description": "The weight. The range is 0 to 999.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "jwt_validation_config": {
              "computed": true,
              "description": "[HTTPS listeners] Information for validating JWT access tokens in client requests. Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `jwt-validation` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "additional_claims": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "format": {
                          "computed": true,
                          "description": "The format of the claim value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the claim. You can't specify ` + "`" + `` + "`" + `exp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `iss` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `nbf` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `iat` + "`" + `` + "`" + ` because we validate them by default.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "values": {
                          "computed": true,
                          "description": "The claim value. The maximum size of the list is 10. Each value can be up to 256 characters in length. If the format is ` + "`" + `` + "`" + `space-separated-values` + "`" + `` + "`" + `, the values can't include spaces.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "issuer": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "jwks_endpoint": {
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
            "order": {
              "computed": true,
              "description": "The order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "redirect_config": {
              "computed": true,
              "description": "[Application Load Balancer] Information for creating a redirect action. Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `redirect` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host": {
                    "computed": true,
                    "description": "The hostname. This component is not percent-encoded. The hostname can contain #{host}.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "The absolute path, starting with the leading \"/\". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "The port. You can specify a value from 1 to 65535 or #{port}.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "The protocol. You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You can't redirect HTTPS to HTTP.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query": {
                    "computed": true,
                    "description": "The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading \"?\", as it is automatically added. You can specify any of the reserved keywords.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
                    "computed": true,
                    "description": "The HTTP redirect code. The redirect is either permanent (HTTP 301) or temporary (HTTP 302).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "target_group_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the target group. Specify only when ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `forward` + "`" + `` + "`" + ` and you want to route to a single target group. To route to multiple target groups, you must use ` + "`" + `` + "`" + `ForwardConfig` + "`" + `` + "`" + ` instead.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "The type of action.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "conditions": {
        "description": "The conditions.\n The rule can optionally include up to one of each of the following conditions: ` + "`" + `` + "`" + `http-request-method` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `host-header` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `path-pattern` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `source-ip` + "`" + `` + "`" + `. A rule can also optionally include one or more of each of the following conditions: ` + "`" + `` + "`" + `http-header` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `query-string` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field": {
              "computed": true,
              "description": "The field in the HTTP request. The following are the possible values:\n  +   ` + "`" + `` + "`" + `http-header` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `http-request-method` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `host-header` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `path-pattern` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `query-string` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `source-ip` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_header_config": {
              "computed": true,
              "description": "Information for a host header condition. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `host-header` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "regex_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "values": {
                    "computed": true,
                    "description": "The host names. The maximum length of each string is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character). You must include at least one \".\" character. You can include only alphabetical characters after the final \".\" character.\n If you specify multiple strings, the condition is satisfied if one of the strings matches the host name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "http_header_config": {
              "computed": true,
              "description": "Information for an HTTP header condition. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `http-header` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "http_header_name": {
                    "computed": true,
                    "description": "The name of the HTTP header field. The maximum size is 40 characters. The header name is case insensitive. The allowed characters are specified by RFC 7230. Wildcards are not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "regex_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "values": {
                    "computed": true,
                    "description": "The strings to compare against the value of the HTTP header. The maximum length of each string is 128 characters. The comparison strings are case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).\n If the same header appears multiple times in the request, we search them in order until a match is found.\n If you specify multiple strings, the condition is satisfied if one of the strings matches the value of the HTTP header. To require that all of the strings are a match, create one condition per string.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "http_request_method_config": {
              "computed": true,
              "description": "Information for an HTTP method condition. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `http-request-method` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description": "The name of the request method. The maximum length is 40 characters. The allowed characters are A-Z, hyphen (-), and underscore (_). The comparison is case sensitive. Wildcards are not supported; therefore, the method name must be an exact match.\n If you specify multiple strings, the condition is satisfied if one of the strings matches the HTTP request method. We recommend that you route GET and HEAD requests in the same way, because the response to a HEAD request may be cached.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "path_pattern_config": {
              "computed": true,
              "description": "Information for a path pattern condition. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `path-pattern` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "regex_values": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "values": {
                    "computed": true,
                    "description": "The path patterns to compare against the request URL. The maximum size of each string is 128 characters. The comparison is case sensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).\n If you specify multiple strings, the condition is satisfied if one of them matches the request URL. The path pattern is compared only to the path of the URL, not to its query string.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "query_string_config": {
              "computed": true,
              "description": "Information for a query string condition. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `query-string` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description": "The key/value pairs or values to find in the query string. The maximum length of each string is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, you must escape these characters in ` + "`" + `` + "`" + `Values` + "`" + `` + "`" + ` using a '\\' character.\n If you specify multiple key/value pairs or values, the condition is satisfied if one of them is found in the query string.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The key. You can omit the key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "regex_values": {
              "computed": true,
              "description": "The regular expressions to match against the condition field. The maximum length of each string is 128 characters. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `http-header` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `host-header` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `path-pattern` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "source_ip_config": {
              "computed": true,
              "description": "Information for a source IP condition. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `source-ip` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "values": {
                    "computed": true,
                    "description": "The source IP addresses, in CIDR format. You can use both IPv4 and IPv6 addresses. Wildcards are not supported.\n If you specify multiple addresses, the condition is satisfied if the source IP address of the request matches one of the CIDR blocks. This condition is not satisfied by the addresses in the X-Forwarded-For header.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "values": {
              "computed": true,
              "description": "The condition value. Specify only when ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `host-header` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `path-pattern` + "`" + `` + "`" + `. Alternatively, to specify multiple host names or multiple path patterns, use ` + "`" + `` + "`" + `HostHeaderConfig` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `PathPatternConfig` + "`" + `` + "`" + `.\n If ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `host-header` + "`" + `` + "`" + ` and you're not using ` + "`" + `` + "`" + `HostHeaderConfig` + "`" + `` + "`" + `, you can specify a single host name (for example, my.example.com). A host name is case insensitive, can be up to 128 characters in length, and can contain any of the following characters.\n  +  A-Z, a-z, 0-9\n  +  - .\n  +  * (matches 0 or more characters)\n  +  ? (matches exactly 1 character)\n  \n If ` + "`" + `` + "`" + `Field` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `path-pattern` + "`" + `` + "`" + ` and you're not using ` + "`" + `` + "`" + `PathPatternConfig` + "`" + `` + "`" + `, you can specify a single path pattern (for example, /img/*). A path pattern is case-sensitive, can be up to 128 characters in length, and can contain any of the following characters.\n  +  A-Z, a-z, 0-9\n  +  _ - . $ / ~ \" ' @ : +\n  +  \u0026 (using \u0026amp;)\n  +  * (matches 0 or more characters)\n  +  ? (matches exactly 1 character)",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "listener_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the listener.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description": "The rule priority. A listener can't have multiple rules with the same priority.\n If you try to reorder rules by updating their priorities, do not specify a new priority if an existing rule already uses this priority, as this can cause an error. If you need to reuse a priority with a different rule, you must remove it as a priority first, and then specify it in a subsequent update.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rule_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transforms": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "host_header_rewrite_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rewrites": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "regex": {
                          "computed": true,
                          "description": "The regular expression to match in the input string. The maximum length of the string is 1,024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replace": {
                          "computed": true,
                          "description": "The replacement string to use when rewriting the matched input. The maximum length of the string is 1,024 characters. You can specify capture groups in the regular expression (for example, $1 and $2).",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_rewrite_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rewrites": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "regex": {
                          "computed": true,
                          "description": "The regular expression to match in the input string. The maximum length of the string is 1,024 characters.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replace": {
                          "computed": true,
                          "description": "The replacement string to use when rewriting the matched input. The maximum length of the string is 1,024 characters. You can specify capture groups in the regular expression (for example, $1 and $2).",
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
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Specifies a listener rule. The listener must be associated with an Application Load Balancer. Each rule consists of a priority, one or more actions, and one or more conditions.\n For more information, see [Quotas for your Application Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-limits.html) in the *User Guide for Application Load Balancers*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticloadbalancingv2ListenerRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticloadbalancingv2ListenerRule), &result)
	return &result
}
