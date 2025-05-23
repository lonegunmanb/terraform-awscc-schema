package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontResponseHeadersPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "response_headers_policy_config": {
        "computed": true,
        "description": "A response headers policy configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "A comment to describe the response headers policy.\n The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "cors_config": {
              "computed": true,
              "description": "A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_control_allow_credentials": {
                    "computed": true,
                    "description": "A Boolean that CloudFront uses as the value for the ` + "`" + `` + "`" + `Access-Control-Allow-Credentials` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Access-Control-Allow-Credentials` + "`" + `` + "`" + ` HTTP response header, see [Access-Control-Allow-Credentials](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "access_control_allow_headers": {
                    "computed": true,
                    "description": "A list of HTTP header names that CloudFront includes as values for the ` + "`" + `` + "`" + `Access-Control-Allow-Headers` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Access-Control-Allow-Headers` + "`" + `` + "`" + ` HTTP response header, see [Access-Control-Allow-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
                          "computed": true,
                          "description": "The list of HTTP header names. You can specify ` + "`" + `` + "`" + `*` + "`" + `` + "`" + ` to allow all headers.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "access_control_allow_methods": {
                    "computed": true,
                    "description": "A list of HTTP methods that CloudFront includes as values for the ` + "`" + `` + "`" + `Access-Control-Allow-Methods` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Access-Control-Allow-Methods` + "`" + `` + "`" + ` HTTP response header, see [Access-Control-Allow-Methods](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
                          "computed": true,
                          "description": "The list of HTTP methods. Valid values are:\n  +   ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DELETE` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `OPTIONS` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `PATCH` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `POST` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `PUT` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `ALL` + "`" + `` + "`" + ` \n  \n ` + "`" + `` + "`" + `ALL` + "`" + `` + "`" + ` is a special value that includes all of the listed HTTP methods.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "access_control_allow_origins": {
                    "computed": true,
                    "description": "A list of origins (domain names) that CloudFront can use as the value for the ` + "`" + `` + "`" + `Access-Control-Allow-Origin` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Access-Control-Allow-Origin` + "`" + `` + "`" + ` HTTP response header, see [Access-Control-Allow-Origin](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
                          "computed": true,
                          "description": "The list of origins (domain names). You can specify ` + "`" + `` + "`" + `*` + "`" + `` + "`" + ` to allow all origins.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "access_control_expose_headers": {
                    "computed": true,
                    "description": "A list of HTTP headers that CloudFront includes as values for the ` + "`" + `` + "`" + `Access-Control-Expose-Headers` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Access-Control-Expose-Headers` + "`" + `` + "`" + ` HTTP response header, see [Access-Control-Expose-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "items": {
                          "computed": true,
                          "description": "The list of HTTP headers. You can specify ` + "`" + `` + "`" + `*` + "`" + `` + "`" + ` to expose all headers.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "access_control_max_age_sec": {
                    "computed": true,
                    "description": "A number that CloudFront uses as the value for the ` + "`" + `` + "`" + `Access-Control-Max-Age` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Access-Control-Max-Age` + "`" + `` + "`" + ` HTTP response header, see [Access-Control-Max-Age](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "origin_override": {
                    "computed": true,
                    "description": "A Boolean that determines whether CloudFront overrides HTTP response headers received from the origin with the ones specified in this response headers policy.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "custom_headers_config": {
              "computed": true,
              "description": "A configuration for a set of custom HTTP response headers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "items": {
                    "computed": true,
                    "description": "The list of HTTP response headers and their values.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "computed": true,
                          "description": "The HTTP response header name.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides a response header with the same name received from the origin with the header specified here.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value for the HTTP response header.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "name": {
              "computed": true,
              "description": "A name to identify the response headers policy.\n The name must be unique for response headers policies in this AWS-account.",
              "description_kind": "plain",
              "type": "string"
            },
            "remove_headers_config": {
              "computed": true,
              "description": "A configuration for a set of HTTP headers to remove from the HTTP response.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "items": {
                    "computed": true,
                    "description": "The list of HTTP header names.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "computed": true,
                          "description": "The HTTP header name.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_headers_config": {
              "computed": true,
              "description": "A configuration for a set of security-related HTTP response headers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_security_policy": {
                    "computed": true,
                    "description": "The policy directives and their values that CloudFront includes as values for the ` + "`" + `` + "`" + `Content-Security-Policy` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Content-Security-Policy` + "`" + `` + "`" + ` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_security_policy": {
                          "computed": true,
                          "description": "The policy directives and their values that CloudFront includes as values for the ` + "`" + `` + "`" + `Content-Security-Policy` + "`" + `` + "`" + ` HTTP response header.\n For more information about the ` + "`" + `` + "`" + `Content-Security-Policy` + "`" + `` + "`" + ` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides the ` + "`" + `` + "`" + `Content-Security-Policy` + "`" + `` + "`" + ` HTTP response header received from the origin with the one specified in this response headers policy.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "content_type_options": {
                    "computed": true,
                    "description": "Determines whether CloudFront includes the ` + "`" + `` + "`" + `X-Content-Type-Options` + "`" + `` + "`" + ` HTTP response header with its value set to ` + "`" + `` + "`" + `nosniff` + "`" + `` + "`" + `.\n For more information about the ` + "`" + `` + "`" + `X-Content-Type-Options` + "`" + `` + "`" + ` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides the ` + "`" + `` + "`" + `X-Content-Type-Options` + "`" + `` + "`" + ` HTTP response header received from the origin with the one specified in this response headers policy.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "frame_options": {
                    "computed": true,
                    "description": "Determines whether CloudFront includes the ` + "`" + `` + "`" + `X-Frame-Options` + "`" + `` + "`" + ` HTTP response header and the header's value.\n For more information about the ` + "`" + `` + "`" + `X-Frame-Options` + "`" + `` + "`" + ` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "frame_option": {
                          "computed": true,
                          "description": "The value of the ` + "`" + `` + "`" + `X-Frame-Options` + "`" + `` + "`" + ` HTTP response header. Valid values are ` + "`" + `` + "`" + `DENY` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SAMEORIGIN` + "`" + `` + "`" + `.\n For more information about these values, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides the ` + "`" + `` + "`" + `X-Frame-Options` + "`" + `` + "`" + ` HTTP response header received from the origin with the one specified in this response headers policy.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "referrer_policy": {
                    "computed": true,
                    "description": "Determines whether CloudFront includes the ` + "`" + `` + "`" + `Referrer-Policy` + "`" + `` + "`" + ` HTTP response header and the header's value.\n For more information about the ` + "`" + `` + "`" + `Referrer-Policy` + "`" + `` + "`" + ` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides the ` + "`" + `` + "`" + `Referrer-Policy` + "`" + `` + "`" + ` HTTP response header received from the origin with the one specified in this response headers policy.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "referrer_policy": {
                          "computed": true,
                          "description": "Determines whether CloudFront includes the ` + "`" + `` + "`" + `Referrer-Policy` + "`" + `` + "`" + ` HTTP response header and the header's value.\n For more information about the ` + "`" + `` + "`" + `Referrer-Policy` + "`" + `` + "`" + ` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "strict_transport_security": {
                    "computed": true,
                    "description": "Determines whether CloudFront includes the ` + "`" + `` + "`" + `Strict-Transport-Security` + "`" + `` + "`" + ` HTTP response header and the header's value.\n For more information about the ` + "`" + `` + "`" + `Strict-Transport-Security` + "`" + `` + "`" + ` HTTP response header, see [Security headers](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-response-headers-policies.html#understanding-response-headers-policies-security) in the *Amazon CloudFront Developer Guide* and [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_control_max_age_sec": {
                          "computed": true,
                          "description": "A number that CloudFront uses as the value for the ` + "`" + `` + "`" + `max-age` + "`" + `` + "`" + ` directive in the ` + "`" + `` + "`" + `Strict-Transport-Security` + "`" + `` + "`" + ` HTTP response header.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "include_subdomains": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront includes the ` + "`" + `` + "`" + `includeSubDomains` + "`" + `` + "`" + ` directive in the ` + "`" + `` + "`" + `Strict-Transport-Security` + "`" + `` + "`" + ` HTTP response header.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides the ` + "`" + `` + "`" + `Strict-Transport-Security` + "`" + `` + "`" + ` HTTP response header received from the origin with the one specified in this response headers policy.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "preload": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront includes the ` + "`" + `` + "`" + `preload` + "`" + `` + "`" + ` directive in the ` + "`" + `` + "`" + `Strict-Transport-Security` + "`" + `` + "`" + ` HTTP response header.",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "xss_protection": {
                    "computed": true,
                    "description": "Determines whether CloudFront includes the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` HTTP response header and the header's value.\n For more information about the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "mode_block": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront includes the ` + "`" + `` + "`" + `mode=block` + "`" + `` + "`" + ` directive in the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` header.\n For more information about this directive, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "override": {
                          "computed": true,
                          "description": "A Boolean that determines whether CloudFront overrides the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` HTTP response header received from the origin with the one specified in this response headers policy.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "protection": {
                          "computed": true,
                          "description": "A Boolean that determines the value of the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` HTTP response header. When this setting is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the value of the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` header is ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `. When this setting is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, the value of the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` header is ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n For more information about these settings, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "report_uri": {
                          "computed": true,
                          "description": "A reporting URI, which CloudFront uses as the value of the ` + "`" + `` + "`" + `report` + "`" + `` + "`" + ` directive in the ` + "`" + `` + "`" + `X-XSS-Protection` + "`" + `` + "`" + ` header.\n You cannot specify a ` + "`" + `` + "`" + `ReportUri` + "`" + `` + "`" + ` when ` + "`" + `` + "`" + `ModeBlock` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.\n For more information about using a reporting URL, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.",
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
            "server_timing_headers_config": {
              "computed": true,
              "description": "A configuration for enabling the ` + "`" + `` + "`" + `Server-Timing` + "`" + `` + "`" + ` header in HTTP responses sent from CloudFront.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "A Boolean that determines whether CloudFront adds the ` + "`" + `` + "`" + `Server-Timing` + "`" + `` + "`" + ` header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "sampling_rate": {
                    "computed": true,
                    "description": "A number 0?100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the ` + "`" + `` + "`" + `Server-Timing` + "`" + `` + "`" + ` header to. When you set the sampling rate to 100, CloudFront adds the ` + "`" + `` + "`" + `Server-Timing` + "`" + `` + "`" + ` header to the HTTP response for every request that matches the cache behavior that this response headers policy is attached to. When you set it to 50, CloudFront adds the header to 50% of the responses for requests that match the cache behavior. You can set the sampling rate to any number 0?100 with up to four decimal places.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "response_headers_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::ResponseHeadersPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontResponseHeadersPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontResponseHeadersPolicy), &result)
	return &result
}
