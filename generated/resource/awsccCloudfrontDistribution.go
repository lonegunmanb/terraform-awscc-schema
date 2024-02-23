package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontDistribution = `{
  "block": {
    "attributes": {
      "distribution_config": {
        "description": "The distribution's configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aliases": {
              "computed": true,
              "description": "A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cache_behaviors": {
              "computed": true,
              "description": "A complex type that contains zero or more ` + "`" + `` + "`" + `CacheBehavior` + "`" + `` + "`" + ` elements.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_methods": {
                    "computed": true,
                    "description": "A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. There are three choices:\n  +  CloudFront forwards only ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + ` requests.\n  +  CloudFront forwards only ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `OPTIONS` + "`" + `` + "`" + ` requests.\n  +  CloudFront forwards ` + "`" + `` + "`" + `GET, HEAD, OPTIONS, PUT, PATCH, POST` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `DELETE` + "`" + `` + "`" + ` requests.\n  \n If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cache_policy_id": {
                    "computed": true,
                    "description": "The unique identifier of the cache policy that is attached to this cache behavior. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n A ` + "`" + `` + "`" + `CacheBehavior` + "`" + `` + "`" + ` must include either a ` + "`" + `` + "`" + `CachePolicyId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ForwardedValues` + "`" + `` + "`" + `. We recommend that you use a ` + "`" + `` + "`" + `CachePolicyId` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cached_methods": {
                    "computed": true,
                    "description": "A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods. There are two choices:\n  +  CloudFront caches responses to ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + ` requests.\n  +  CloudFront caches responses to ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `OPTIONS` + "`" + `` + "`" + ` requests.\n  \n If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "compress": {
                    "computed": true,
                    "description": "Whether you want CloudFront to automatically compress certain files for this cache behavior. If so, specify true; if not, specify false. For more information, see [Serving Compressed Files](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "default_ttl": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use the ` + "`" + `` + "`" + `DefaultTTL` + "`" + `` + "`" + ` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n The default amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin does not add HTTP headers such as ` + "`" + `` + "`" + `Cache-Control max-age` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Cache-Control s-maxage` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Expires` + "`" + `` + "`" + ` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "field_level_encryption_id": {
                    "computed": true,
                    "description": "The value of ` + "`" + `` + "`" + `ID` + "`" + `` + "`" + ` for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for this cache behavior.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "forwarded_values": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. For more information, see [Working with policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-policies.html) in the *Amazon CloudFront Developer Guide*.\n If you want to include values in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-r",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cookies": {
                          "computed": true,
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see [How CloudFront Forwards, Caches, and Logs C",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward": {
                                "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send cookies to the origin but not include them in the cache key, use origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n Specifies which cookies to forward to the origin for this cache behavior: all, none, or the list of cookies specified in the ` + "`" + `` + "`" + `WhitelistedNames` + "`" + `` + "`" + ` complex type.\n Amazon S3 doesn't process cookies. When the cache behavior is forw",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "whitelisted_names": {
                                "computed": true,
                                "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n Required if you specify ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` for the value of ` + "`" + `` + "`" + `Forward` + "`" + `` + "`" + `. A complex type that specifies how many different cookies you want CloudFront to forward to the origin for this cache behavior and, if you want to forward se",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "headers": {
                          "computed": true,
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include headers in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send headers to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n A complex type that specifies the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + `, if any, that you want CloudFront to forward to the origin for this cache behavior (whitelisted headers). For the headers that you specify, CloudFront also caches separate versio",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "query_string": {
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include query strings in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send query strings to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n Indicates whether you want CloudFront to forward query strings to the origin that is associated with this cache behavior and cache based on the query string parameters. CloudFront behavior depends on the value of",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "query_string_cache_keys": {
                          "computed": true,
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include query strings in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send query strings to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n A complex type that contains information about the query string parameters that you want CloudFront to use for caching for this cache behavior.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "function_associations": {
                    "computed": true,
                    "description": "A list of CloudFront functions that are associated with this cache behavior. CloudFront functions must be published to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage to associate them with a cache behavior.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description": "The event type of the function, either ` + "`" + `` + "`" + `viewer-request` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `viewer-response` + "`" + `` + "`" + `. You cannot use origin-facing event types (` + "`" + `` + "`" + `origin-request` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `origin-response` + "`" + `` + "`" + `) with a CloudFront function.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "function_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the function.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "lambda_function_associations": {
                    "computed": true,
                    "description": "A complex type that contains zero or more Lambda@Edge function associations for a cache behavior.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description": "Specifies the event type that triggers a Lambda@Edge function invocation. You can specify the following values:\n  +   ` + "`" + `` + "`" + `viewer-request` + "`" + `` + "`" + `: The function executes when CloudFront receives a request from a viewer and before it checks to see whether the requested object is in the edge cache.\n  +   ` + "`" + `` + "`" + `origin-request` + "`" + `` + "`" + `: The function executes only when CloudFront sends a request to your origin. When the requested object is in the edge cache, the function doesn't execute.\n  +   ` + "`" + `` + "`" + `origin-response` + "`" + `` + "`" + `: The function executes after CloudFront receives a response from the origin and before it caches the object in the response. When the requested object is in the edge cache, the function doesn't execute.\n  +   ` + "`" + `` + "`" + `viewer-response` + "`" + `` + "`" + `: The function executes before CloudFront returns the requested object to the viewer. The function executes regardless of whether the object was already in the edge cache.\n If the origin returns an HTTP status code other than HTTP 200 (OK), the function doesn't execute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "include_body": {
                          "computed": true,
                          "description": "A flag that allows a Lambda@Edge function to have read access to the body content. For more information, see [Accessing the Request Body by Choosing the Include Body Option](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "lambda_function_arn": {
                          "computed": true,
                          "description": "The ARN of the Lambda@Edge function. You must specify the ARN of a function version; you can't specify an alias or $LATEST.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "max_ttl": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use the ` + "`" + `` + "`" + `MaxTTL` + "`" + `` + "`" + ` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n The maximum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin adds HTTP headers such as ` + "`" + `` + "`" + `Cache-Control max-age` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Cache-Control s-maxage` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Expires` + "`" + `` + "`" + ` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_ttl": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use the ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + ` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n You must specify ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + ` if you configure CloudFront to forward all headers to your origin (under ` + "`" + `` + "`" + `He",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_request_policy_id": {
                    "computed": true,
                    "description": "The unique identifier of the origin request policy that is attached to this cache behavior. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path_pattern": {
                    "description": "The pattern (for example, ` + "`" + `` + "`" + `images/*.jpg` + "`" + `` + "`" + `) that specifies which requests to apply the behavior to. When CloudFront receives a viewer request, the requested path is compared with path patterns in the order in which cache behaviors are listed in the distribution.\n  You can optionally include a slash (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `) at the beginning of the path pattern. For example, ` + "`" + `` + "`" + `/images/*.jpg` + "`" + `` + "`" + `. CloudFront behavior is the same with or without the leading ` + "`" + `` + "`" + `/` + "`" + `` + "`" + `.\n  The path pattern for the default cache behavior is ` + "`" + `` + "`" + `*` + "`" + `` + "`" + ` and cannot be changed. If the request for an object does not match the path pattern for any cache behaviors, CloudFront applies the behavior in the default cache behavior.\n For more information, see [Path Pattern](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesPathPattern) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "realtime_log_config_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior. For more information, see [Real-time logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "response_headers_policy_id": {
                    "computed": true,
                    "description": "The identifier for a response headers policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "smooth_streaming": {
                    "computed": true,
                    "description": "Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `; if not, specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. If you specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `SmoothStreaming` + "`" + `` + "`" + `, you can still distribute other content using this cache behavior if the content matches the value of ` + "`" + `` + "`" + `PathPattern` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "target_origin_id": {
                    "description": "The value of ` + "`" + `` + "`" + `ID` + "`" + `` + "`" + ` for the origin that you want CloudFront to route requests to when they match this cache behavior.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "trusted_key_groups": {
                    "computed": true,
                    "description": "A list of key groups that CloudFront can use to validate signed URLs or signed cookies.\n When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with a private key whose corresponding public key is in the key group. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "trusted_signers": {
                    "computed": true,
                    "description": "We recommend using ` + "`" + `` + "`" + `TrustedKeyGroups` + "`" + `` + "`" + ` instead of ` + "`" + `` + "`" + `TrustedSigners` + "`" + `` + "`" + `.\n  A list of AWS-account IDs whose public keys CloudFront can use to validate signed URLs or signed cookies.\n When a cache behavior contains trusted signers, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with the private key of a CloudFront key pair in the trusted signer's AWS-account. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "viewer_protocol_policy": {
                    "description": "The protocol that viewers can use to access the files in the origin specified by ` + "`" + `` + "`" + `TargetOriginId` + "`" + `` + "`" + ` when a request matches the path pattern in ` + "`" + `` + "`" + `PathPattern` + "`" + `` + "`" + `. You can specify the following options:\n  +   ` + "`" + `` + "`" + `allow-all` + "`" + `` + "`" + `: Viewers can use HTTP or HTTPS.\n  +   ` + "`" + `` + "`" + `redirect-to-https` + "`" + `` + "`" + `: If a viewer submits an HTTP request, CloudFront returns an HTTP status code of 301 (Moved Permanently) to the viewer along with the HTTPS URL. The viewer then resubmits the request using the new URL.\n  +   ` + "`" + `` + "`" + `https-only` + "`" + `` + "`" + `: If a viewer sends an HTTP request, CloudFront returns an HTTP status code of 403 (Forbidden).\n  \n For more information about requiring the HTTPS protocol, see [Requiring HTTPS Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-viewers-to-cloudfront.html) in the *Amazon CloudFront Developer Guide*.\n  The only way to guarantee that viewers retrieve an object that was fetched from the origin using HTTPS is never to use any other protocol",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "cnames": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "comment": {
              "computed": true,
              "description": "A comment to describe the distribution. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "continuous_deployment_policy_id": {
              "computed": true,
              "description": "The identifier of a continuous deployment policy. For more information, see ` + "`" + `` + "`" + `CreateContinuousDeploymentPolicy` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_error_responses": {
              "computed": true,
              "description": "A complex type that controls the following:\n  +  Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.\n  +  How long CloudFront caches HTTP status codes in the 4xx and 5xx range.\n  \n For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "error_caching_min_ttl": {
                    "computed": true,
                    "description": "The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ` + "`" + `` + "`" + `ErrorCode` + "`" + `` + "`" + `. When this time period has elapsed, CloudFront queries your origin to see whether the problem that caused the error has been resolved and the requested object is now available.\n For more information, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "error_code": {
                    "description": "The HTTP status code for which you want to specify a custom error page and/or a caching duration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "response_code": {
                    "computed": true,
                    "description": "The HTTP status code that you want CloudFront to return to the viewer along with the custom error page. There are a variety of reasons that you might want CloudFront to return a status code different from the status code that your origin returned to CloudFront, for example:\n  +  Some Internet devices (some firewalls and corporate proxies, for example) intercept HTTP 4xx and 5xx and prevent the response from being returned to the viewer. If you substitute ` + "`" + `` + "`" + `200` + "`" + `` + "`" + `, the response typically won't be intercepted.\n  +  If you don't care about distinguishing among different client errors or server errors, you can specify ` + "`" + `` + "`" + `400` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `500` + "`" + `` + "`" + ` as the ` + "`" + `` + "`" + `ResponseCode` + "`" + `` + "`" + ` for all 4xx or 5xx errors.\n  +  You might want to return a ` + "`" + `` + "`" + `200` + "`" + `` + "`" + ` status code (OK) and static website so your customers don't know that your website is down.\n  \n If you specify a value for ` + "`" + `` + "`" + `ResponseCode` + "`" + `` + "`" + `, you must also specify a value for ` + "`" + `` + "`" + `ResponsePagePath` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "response_page_path": {
                    "computed": true,
                    "description": "The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the HTTP status code specified by ` + "`" + `` + "`" + `ErrorCode` + "`" + `` + "`" + `, for example, ` + "`" + `` + "`" + `/4xx-errors/403-forbidden.html` + "`" + `` + "`" + `. If you want to store your objects and your custom error pages in different locations, your distribution must include a cache behavior for which the following is true:\n  +  The value of ` + "`" + `` + "`" + `PathPattern` + "`" + `` + "`" + ` matches the path to your custom error messages. For example, suppose you saved custom error pages for 4xx errors in an Amazon S3 bucket in a directory named ` + "`" + `` + "`" + `/4xx-errors` + "`" + `` + "`" + `. Your distribution must include a cache behavior for which the path pattern routes requests for your custom error pages to that location, for example, ` + "`" + `` + "`" + `/4xx-errors/*` + "`" + `` + "`" + `.\n  +  The value of ` + "`" + `` + "`" + `TargetOriginId` + "`" + `` + "`" + ` specifies the value of the ` + "`" + `` + "`" + `ID` + "`" + `` + "`" + ` element for the origin that contains your custom error pages.\n  \n If you specify a value for ` + "`" + `` + "`" + `ResponsePagePath` + "`" + `` + "`" + `, you must also specify a value for ` + "`" + `` + "`" + `ResponseCode` + "`" + `` + "`" + `.\n We recommend ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "custom_origin": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "http_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "https_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_protocol_policy": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "origin_ssl_protocols": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "default_cache_behavior": {
              "description": "A complex type that describes the default cache behavior if you don't specify a ` + "`" + `` + "`" + `CacheBehavior` + "`" + `` + "`" + ` element or if files don't match any of the values of ` + "`" + `` + "`" + `PathPattern` + "`" + `` + "`" + ` in ` + "`" + `` + "`" + `CacheBehavior` + "`" + `` + "`" + ` elements. You must create exactly one default cache behavior.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_methods": {
                    "computed": true,
                    "description": "A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. There are three choices:\n  +  CloudFront forwards only ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + ` requests.\n  +  CloudFront forwards only ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `OPTIONS` + "`" + `` + "`" + ` requests.\n  +  CloudFront forwards ` + "`" + `` + "`" + `GET, HEAD, OPTIONS, PUT, PATCH, POST` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `DELETE` + "`" + `` + "`" + ` requests.\n  \n If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cache_policy_id": {
                    "computed": true,
                    "description": "The unique identifier of the cache policy that is attached to the default cache behavior. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n A ` + "`" + `` + "`" + `DefaultCacheBehavior` + "`" + `` + "`" + ` must include either a ` + "`" + `` + "`" + `CachePolicyId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ForwardedValues` + "`" + `` + "`" + `. We recommend that you use a ` + "`" + `` + "`" + `CachePolicyId` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cached_methods": {
                    "computed": true,
                    "description": "A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods. There are two choices:\n  +  CloudFront caches responses to ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + ` requests.\n  +  CloudFront caches responses to ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `OPTIONS` + "`" + `` + "`" + ` requests.\n  \n If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "compress": {
                    "computed": true,
                    "description": "Whether you want CloudFront to automatically compress certain files for this cache behavior. If so, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `; if not, specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. For more information, see [Serving Compressed Files](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "default_ttl": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use the ` + "`" + `` + "`" + `DefaultTTL` + "`" + `` + "`" + ` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n The default amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin does not add HTTP headers such as ` + "`" + `` + "`" + `Cache-Control max-age` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Cache-Control s-maxage` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Expires` + "`" + `` + "`" + ` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "field_level_encryption_id": {
                    "computed": true,
                    "description": "The value of ` + "`" + `` + "`" + `ID` + "`" + `` + "`" + ` for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "forwarded_values": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. For more information, see [Working with policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-policies.html) in the *Amazon CloudFront Developer Guide*.\n If you want to include values in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-r",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cookies": {
                          "computed": true,
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see [How CloudFront Forwards, Caches, and Logs C",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "forward": {
                                "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send cookies to the origin but not include them in the cache key, use origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n Specifies which cookies to forward to the origin for this cache behavior: all, none, or the list of cookies specified in the ` + "`" + `` + "`" + `WhitelistedNames` + "`" + `` + "`" + ` complex type.\n Amazon S3 doesn't process cookies. When the cache behavior is forw",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "whitelisted_names": {
                                "computed": true,
                                "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include cookies in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send cookies to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n Required if you specify ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + ` for the value of ` + "`" + `` + "`" + `Forward` + "`" + `` + "`" + `. A complex type that specifies how many different cookies you want CloudFront to forward to the origin for this cache behavior and, if you want to forward se",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "headers": {
                          "computed": true,
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include headers in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send headers to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n A complex type that specifies the ` + "`" + `` + "`" + `Headers` + "`" + `` + "`" + `, if any, that you want CloudFront to forward to the origin for this cache behavior (whitelisted headers). For the headers that you specify, CloudFront also caches separate versio",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "query_string": {
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include query strings in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send query strings to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n Indicates whether you want CloudFront to forward query strings to the origin that is associated with this cache behavior and cache based on the query string parameters. CloudFront behavior depends on the value of",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        },
                        "query_string_cache_keys": {
                          "computed": true,
                          "description": "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field.\n If you want to include query strings in the cache key, use a cache policy. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the *Amazon CloudFront Developer Guide*.\n If you want to send query strings to the origin but not include them in the cache key, use an origin request policy. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the *Amazon CloudFront Developer Guide*.\n A complex type that contains information about the query string parameters that you want CloudFront to use for caching for this cache behavior.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "function_associations": {
                    "computed": true,
                    "description": "A list of CloudFront functions that are associated with this cache behavior. CloudFront functions must be published to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage to associate them with a cache behavior.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description": "The event type of the function, either ` + "`" + `` + "`" + `viewer-request` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `viewer-response` + "`" + `` + "`" + `. You cannot use origin-facing event types (` + "`" + `` + "`" + `origin-request` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `origin-response` + "`" + `` + "`" + `) with a CloudFront function.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "function_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the function.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "lambda_function_associations": {
                    "computed": true,
                    "description": "A complex type that contains zero or more Lambda@Edge function associations for a cache behavior.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description": "Specifies the event type that triggers a Lambda@Edge function invocation. You can specify the following values:\n  +   ` + "`" + `` + "`" + `viewer-request` + "`" + `` + "`" + `: The function executes when CloudFront receives a request from a viewer and before it checks to see whether the requested object is in the edge cache.\n  +   ` + "`" + `` + "`" + `origin-request` + "`" + `` + "`" + `: The function executes only when CloudFront sends a request to your origin. When the requested object is in the edge cache, the function doesn't execute.\n  +   ` + "`" + `` + "`" + `origin-response` + "`" + `` + "`" + `: The function executes after CloudFront receives a response from the origin and before it caches the object in the response. When the requested object is in the edge cache, the function doesn't execute.\n  +   ` + "`" + `` + "`" + `viewer-response` + "`" + `` + "`" + `: The function executes before CloudFront returns the requested object to the viewer. The function executes regardless of whether the object was already in the edge cache.\n If the origin returns an HTTP status code other than HTTP 200 (OK), the function doesn't execute.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "include_body": {
                          "computed": true,
                          "description": "A flag that allows a Lambda@Edge function to have read access to the body content. For more information, see [Accessing the Request Body by Choosing the Include Body Option](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "lambda_function_arn": {
                          "computed": true,
                          "description": "The ARN of the Lambda@Edge function. You must specify the ARN of a function version; you can't specify an alias or $LATEST.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "max_ttl": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use the ` + "`" + `` + "`" + `MaxTTL` + "`" + `` + "`" + ` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n The maximum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. The value that you specify applies only when your origin adds HTTP headers such as ` + "`" + `` + "`" + `Cache-Control max-age` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Cache-Control s-maxage` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Expires` + "`" + `` + "`" + ` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_ttl": {
                    "computed": true,
                    "description": "This field is deprecated. We recommend that you use the ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + ` field in a cache policy instead of this field. For more information, see [Creating cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) or [Using the managed cache policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html) in the *Amazon CloudFront Developer Guide*.\n The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront forwards another request to your origin to determine whether the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n You must specify ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `MinTTL` + "`" + `` + "`" + ` if you configure CloudFront to forward all headers to your origin (under ` + "`" + `` + "`" + `He",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "origin_request_policy_id": {
                    "computed": true,
                    "description": "The unique identifier of the origin request policy that is attached to the default cache behavior. For more information, see [Creating origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) or [Using the managed origin request policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-origin-request-policies.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "realtime_log_config_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior. For more information, see [Real-time logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "response_headers_policy_id": {
                    "computed": true,
                    "description": "The identifier for a response headers policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "smooth_streaming": {
                    "computed": true,
                    "description": "Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `; if not, specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. If you specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `SmoothStreaming` + "`" + `` + "`" + `, you can still distribute other content using this cache behavior if the content matches the value of ` + "`" + `` + "`" + `PathPattern` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "target_origin_id": {
                    "description": "The value of ` + "`" + `` + "`" + `ID` + "`" + `` + "`" + ` for the origin that you want CloudFront to route requests to when they use the default cache behavior.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "trusted_key_groups": {
                    "computed": true,
                    "description": "A list of key groups that CloudFront can use to validate signed URLs or signed cookies.\n When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with a private key whose corresponding public key is in the key group. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "trusted_signers": {
                    "computed": true,
                    "description": "We recommend using ` + "`" + `` + "`" + `TrustedKeyGroups` + "`" + `` + "`" + ` instead of ` + "`" + `` + "`" + `TrustedSigners` + "`" + `` + "`" + `.\n  A list of AWS-account IDs whose public keys CloudFront can use to validate signed URLs or signed cookies.\n When a cache behavior contains trusted signers, CloudFront requires signed URLs or signed cookies for all requests that match the cache behavior. The URLs or cookies must be signed with the private key of a CloudFront key pair in a trusted signer's AWS-account. The signed URL or cookie contains information about which public key CloudFront should use to verify the signature. For more information, see [Serving private content](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "viewer_protocol_policy": {
                    "description": "The protocol that viewers can use to access the files in the origin specified by ` + "`" + `` + "`" + `TargetOriginId` + "`" + `` + "`" + ` when a request matches the path pattern in ` + "`" + `` + "`" + `PathPattern` + "`" + `` + "`" + `. You can specify the following options:\n  +   ` + "`" + `` + "`" + `allow-all` + "`" + `` + "`" + `: Viewers can use HTTP or HTTPS.\n  +   ` + "`" + `` + "`" + `redirect-to-https` + "`" + `` + "`" + `: If a viewer submits an HTTP request, CloudFront returns an HTTP status code of 301 (Moved Permanently) to the viewer along with the HTTPS URL. The viewer then resubmits the request using the new URL.\n  +   ` + "`" + `` + "`" + `https-only` + "`" + `` + "`" + `: If a viewer sends an HTTP request, CloudFront returns an HTTP status code of 403 (Forbidden).\n  \n For more information about requiring the HTTPS protocol, see [Requiring HTTPS Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-viewers-to-cloudfront.html) in the *Amazon CloudFront Developer Guide*.\n  The only way to guarantee that viewers retrieve an object that was fetched from the origin using HTTPS is never to use any other protocol",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "default_root_object": {
              "computed": true,
              "description": "The object that you want CloudFront to request from your origin (for example, ` + "`" + `` + "`" + `index.html` + "`" + `` + "`" + `) when a viewer requests the root URL for your distribution (` + "`" + `` + "`" + `https://www.example.com` + "`" + `` + "`" + `) instead of an object in your distribution (` + "`" + `` + "`" + `https://www.example.com/product-description.html` + "`" + `` + "`" + `). Specifying a default root object avoids exposing the contents of your distribution.\n Specify only the object name, for example, ` + "`" + `` + "`" + `index.html` + "`" + `` + "`" + `. Don't add a ` + "`" + `` + "`" + `/` + "`" + `` + "`" + ` before the object name.\n If you don't want to specify a default root object when you create a distribution, include an empty ` + "`" + `` + "`" + `DefaultRootObject` + "`" + `` + "`" + ` element.\n To delete the default root object from an existing distribution, update the distribution configuration and include an empty ` + "`" + `` + "`" + `DefaultRootObject` + "`" + `` + "`" + ` element.\n To replace the default root object, update the distribution configuration and specify the new object.\n For more information about the default root object, see [Creating a Default Root Object](https://docs.aws.amazon.com/AmazonCloudFront/latest/D",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description": "From this field, you can enable or disable the selected distribution.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "http_version": {
              "computed": true,
              "description": "(Optional) Specify the maximum HTTP version(s) that you want viewers to use to communicate with CF. The default value for new distributions is ` + "`" + `` + "`" + `http1.1` + "`" + `` + "`" + `.\n For viewers and CF to use HTTP/2, viewers must support TLSv1.2 or later, and must support Server Name Indication (SNI).\n For viewers and CF to use HTTP/3, viewers must support TLSv1.3 and Server Name Indication (SNI). CF supports HTTP/3 connection migration to allow the viewer to switch networks without losing connection. For more information about connection migration, see [Connection Migration](https://docs.aws.amazon.com/https://www.rfc-editor.org/rfc/rfc9000.html#name-connection-migration) at RFC 9000. For more information about supported TLSv1.3 ciphers, see [Supported protocols and ciphers between viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/secure-connections-supported-viewer-protocols-ciphers.html).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ipv6_enabled": {
              "computed": true,
              "description": "If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. If you specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, CloudFront responds to IPv6 DNS requests with the DNS response code ` + "`" + `` + "`" + `NOERROR` + "`" + `` + "`" + ` and with no IP addresses. This allows viewers to submit a second request, for an IPv4 address for your distribution.\n In general, you should enable IPv6 if you have users on IPv6 networks who want to access your content. However, if you're using signed URLs or signed cookies to restrict access to your content, and if you're using a custom policy that includes the ` + "`" + `` + "`" + `IpAddress` + "`" + `` + "`" + ` parameter to restrict the IP addresses that can access your content, don't enable IPv6. If you want to restrict access to some content by IP address and not restrict access to other content (or restrict access but not by IP address), you can create two distributions. For more information, see [Creating a Signed URL Using a Custom Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/Devel",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "logging": {
              "computed": true,
              "description": "A complex type that controls whether access logs are written for the distribution.\n For more information about logging, see [Access Logs](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/AccessLogs.html) in the *Amazon CloudFront Developer Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "description": "The Amazon S3 bucket to store the access logs in, for example, ` + "`" + `` + "`" + `myawslogbucket.s3.amazonaws.com` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "include_cookies": {
                    "computed": true,
                    "description": "Specifies whether you want CloudFront to include cookies in access logs, specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `IncludeCookies` + "`" + `` + "`" + `. If you choose to include cookies in logs, CloudFront logs all cookies regardless of how you configure the cache behaviors for this distribution. If you don't want to include cookies when you create a distribution or if you want to disable include cookies for an existing distribution, specify ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` for ` + "`" + `` + "`" + `IncludeCookies` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "An optional string that you want CloudFront to prefix to the access log ` + "`" + `` + "`" + `filenames` + "`" + `` + "`" + ` for this distribution, for example, ` + "`" + `` + "`" + `myprefix/` + "`" + `` + "`" + `. If you want to enable logging, but you don't want to specify a prefix, you still must include an empty ` + "`" + `` + "`" + `Prefix` + "`" + `` + "`" + ` element in the ` + "`" + `` + "`" + `Logging` + "`" + `` + "`" + ` element.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "origin_groups": {
              "computed": true,
              "description": "A complex type that contains information about origin groups for this distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "items": {
                    "computed": true,
                    "description": "The items (origin groups) in a distribution.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "failover_criteria": {
                          "description": "A complex type that contains information about the failover criteria for an origin group.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status_codes": {
                                "description": "The status codes that, when returned from the primary origin, will trigger CloudFront to failover to the second origin.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "items": {
                                      "description": "The items (status codes) for an origin group.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "number"
                                      ]
                                    },
                                    "quantity": {
                                      "description": "The number of status codes.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "id": {
                          "description": "The origin group's ID.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "members": {
                          "description": "A complex type that contains information about the origins in an origin group.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "items": {
                                "description": "Items (origins) in an origin group.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "origin_id": {
                                      "description": "The ID for an origin in an origin group.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              },
                              "quantity": {
                                "description": "The number of origins in an origin group.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "quantity": {
                    "description": "The number of origin groups.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "origins": {
              "computed": true,
              "description": "A complex type that contains information about origins for this distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_attempts": {
                    "computed": true,
                    "description": "The number of times that CloudFront attempts to connect to the origin. The minimum number is 1, the maximum is 3, and the default (if you don't specify otherwise) is 3.\n For a custom origin (including an Amazon S3 bucket that's configured with static website hosting), this value also specifies the number of times that CloudFront attempts to get a response from the origin, in the case of an [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout).\n For more information, see [Origin Connection Attempts](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-attempts) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "connection_timeout": {
                    "computed": true,
                    "description": "The number of seconds that CloudFront waits when trying to establish a connection to the origin. The minimum timeout is 1 second, the maximum is 10 seconds, and the default (if you don't specify otherwise) is 10 seconds.\n For more information, see [Origin Connection Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#origin-connection-timeout) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "custom_origin_config": {
                    "computed": true,
                    "description": "Use this type to specify an origin that is not an Amazon S3 bucket, with one exception. If the Amazon S3 bucket is configured with static website hosting, use this type. If the Amazon S3 bucket is not configured with static website hosting, use the ` + "`" + `` + "`" + `S3OriginConfig` + "`" + `` + "`" + ` type instead.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "http_port": {
                          "computed": true,
                          "description": "The HTTP port that CloudFront uses to connect to the origin. Specify the HTTP port that the origin listens on.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "https_port": {
                          "computed": true,
                          "description": "The HTTPS port that CloudFront uses to connect to the origin. Specify the HTTPS port that the origin listens on.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "origin_keepalive_timeout": {
                          "computed": true,
                          "description": "Specifies how long, in seconds, CloudFront persists its connection to the origin. The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 5 seconds.\n For more information, see [Origin Keep-alive Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "origin_protocol_policy": {
                          "description": "Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin. Valid values are:\n  +   ` + "`" + `` + "`" + `http-only` + "`" + `` + "`" + ` ? CloudFront always uses HTTP to connect to the origin.\n  +   ` + "`" + `` + "`" + `match-viewer` + "`" + `` + "`" + ` ? CloudFront connects to the origin using the same protocol that the viewer used to connect to CloudFront.\n  +   ` + "`" + `` + "`" + `https-only` + "`" + `` + "`" + ` ? CloudFront always uses HTTPS to connect to the origin.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "origin_read_timeout": {
                          "computed": true,
                          "description": "Specifies how long, in seconds, CloudFront waits for a response from the origin. This is also known as the *origin response timeout*. The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 30 seconds.\n For more information, see [Origin Response Timeout](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "origin_ssl_protocols": {
                          "computed": true,
                          "description": "Specifies the minimum SSL/TLS protocol that CloudFront uses when connecting to your origin over HTTPS. Valid values include ` + "`" + `` + "`" + `SSLv3` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `TLSv1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `TLSv1.1` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `TLSv1.2` + "`" + `` + "`" + `.\n For more information, see [Minimum Origin SSL Protocol](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginSSLProtocols) in the *Amazon CloudFront Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "domain_name": {
                    "description": "The domain name for the origin.\n For more information, see [Origin Domain Name](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesDomainName) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "id": {
                    "description": "A unique identifier for the origin. This value must be unique within the distribution.\n Use this value to specify the ` + "`" + `` + "`" + `TargetOriginId` + "`" + `` + "`" + ` in a ` + "`" + `` + "`" + `CacheBehavior` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DefaultCacheBehavior` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "origin_access_control_id": {
                    "computed": true,
                    "description": "The unique identifier of an origin access control for this origin.\n For more information, see [Restricting access to an Amazon S3 origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "origin_custom_headers": {
                    "computed": true,
                    "description": "A list of HTTP header names and values that CloudFront adds to the requests that it sends to the origin.\n For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/add-origin-custom-headers.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header_name": {
                          "description": "The name of a header that you want CloudFront to send to your origin. For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/forward-custom-headers.html) in the *Amazon CloudFront Developer Guide*.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_value": {
                          "description": "The value for the header that you specified in the ` + "`" + `` + "`" + `HeaderName` + "`" + `` + "`" + ` field.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "origin_path": {
                    "computed": true,
                    "description": "An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin.\n For more information, see [Origin Path](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginPath) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "origin_shield": {
                    "computed": true,
                    "description": "CloudFront Origin Shield. Using Origin Shield can help reduce the load on your origin.\n For more information, see [Using Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the *Amazon CloudFront Developer Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enabled": {
                          "computed": true,
                          "description": "A flag that specifies whether Origin Shield is enabled.\n When it's enabled, CloudFront routes all requests through Origin Shield, which can help protect your origin. When it's disabled, CloudFront might send requests directly to your origin from multiple edge locations or regional edge caches.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "origin_shield_region": {
                          "computed": true,
                          "description": "The AWS-Region for Origin Shield.\n Specify the AWS-Region that has the lowest latency to your origin. To specify a region, use the region code, not the region name. For example, specify the US East (Ohio) region as ` + "`" + `` + "`" + `us-east-2` + "`" + `` + "`" + `.\n When you enable CloudFront Origin Shield, you must specify the AWS-Region for Origin Shield. For the list of AWS-Regions that you can specify, and for help choosing the best Region for your origin, see [Choosing the for Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html#choose-origin-shield-region) in the *Amazon CloudFront Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "s3_origin_config": {
                    "computed": true,
                    "description": "Use this type to specify an origin that is an Amazon S3 bucket that is not configured with static website hosting. To specify any other type of origin, including an Amazon S3 bucket that is configured with static website hosting, use the ` + "`" + `` + "`" + `CustomOriginConfig` + "`" + `` + "`" + ` type instead.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "origin_access_identity": {
                          "computed": true,
                          "description": "The CloudFront origin access identity to associate with the origin. Use an origin access identity to configure the origin so that viewers can *only* access objects in an Amazon S3 bucket through CloudFront. The format of the value is:\n origin-access-identity/cloudfront/*ID-of-origin-access-identity* \n where ` + "`" + `` + "`" + `ID-of-origin-access-identity` + "`" + `` + "`" + ` is the value that CloudFront returned in the ` + "`" + `` + "`" + `ID` + "`" + `` + "`" + ` element when you created the origin access identity.\n If you want viewers to be able to access objects using either the CloudFront URL or the Amazon S3 URL, specify an empty ` + "`" + `` + "`" + `OriginAccessIdentity` + "`" + `` + "`" + ` element.\n To delete the origin access identity from an existing distribution, update the distribution configuration and include an empty ` + "`" + `` + "`" + `OriginAccessIdentity` + "`" + `` + "`" + ` element.\n To replace the origin access identity, update the distribution configuration and specify the new origin access identity.\n For more information about the origin access identity, see [Serving Private Content through CloudFront](https://d",
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
            "price_class": {
              "computed": true,
              "description": "The price class that corresponds with the maximum price that you want to pay for CloudFront service. If you specify ` + "`" + `` + "`" + `PriceClass_All` + "`" + `` + "`" + `, CloudFront responds to requests for your objects from all CloudFront edge locations.\n If you specify a price class other than ` + "`" + `` + "`" + `PriceClass_All` + "`" + `` + "`" + `, CloudFront serves your objects from the CloudFront edge location that has the lowest latency among the edge locations in your price class. Viewers who are in or near regions that are excluded from your specified price class may encounter slower performance.\n For more information about price classes, see [Choosing the Price Class for a CloudFront Distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PriceClass.html) in the *Amazon CloudFront Developer Guide*. For information about CloudFront pricing, including how price classes (such as Price Class 100) map to CloudFront regions, see [Amazon CloudFront Pricing](https://docs.aws.amazon.com/cloudfront/pricing/).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "restrictions": {
              "computed": true,
              "description": "A complex type that identifies ways in which you want to restrict distribution of your content.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "geo_restriction": {
                    "description": "A complex type that controls the countries in which your content is distributed. CF determines the location of your users using ` + "`" + `` + "`" + `MaxMind` + "`" + `` + "`" + ` GeoIP databases. To disable geo restriction, remove the [Restrictions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-distributionconfig.html#cfn-cloudfront-distribution-distributionconfig-restrictions) property from your stack template.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "locations": {
                          "computed": true,
                          "description": "A complex type that contains a ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` element for each country in which you want CloudFront either to distribute your content (` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + `) or not distribute your content (` + "`" + `` + "`" + `blacklist` + "`" + `` + "`" + `).\n The ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` element is a two-letter, uppercase country code for a country that you want to include in your ` + "`" + `` + "`" + `blacklist` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + `. Include one ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` element for each country.\n CloudFront and ` + "`" + `` + "`" + `MaxMind` + "`" + `` + "`" + ` both use ` + "`" + `` + "`" + `ISO 3166` + "`" + `` + "`" + ` country codes. For the current list of countries and the corresponding codes, see ` + "`" + `` + "`" + `ISO 3166-1-alpha-2` + "`" + `` + "`" + ` code on the *International Organization for Standardization* website. You can also refer to the country list on the CloudFront console, which includes both country names and codes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "restriction_type": {
                          "description": "The method that you want to use to restrict distribution of your content by country:\n  +   ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `: No geo restriction is enabled, meaning access to content is not restricted by client geo location.\n  +   ` + "`" + `` + "`" + `blacklist` + "`" + `` + "`" + `: The ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` elements specify the countries in which you don't want CloudFront to distribute your content.\n  +   ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + `: The ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` elements specify the countries in which you want CloudFront to distribute your content.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_origin": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "origin_access_identity": {
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
            "staging": {
              "computed": true,
              "description": "A Boolean that indicates whether this is a staging distribution. When this value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, this is a staging distribution. When this value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, this is not a staging distribution.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "viewer_certificate": {
              "computed": true,
              "description": "A complex type that determines the distribution's SSL/TLS configuration for communicating with viewers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "acm_certificate_arn": {
                    "computed": true,
                    "description": "In CloudFormation, this field name is ` + "`" + `` + "`" + `AcmCertificateArn` + "`" + `` + "`" + `. Note the different capitalization.\n  If the distribution uses ` + "`" + `` + "`" + `Aliases` + "`" + `` + "`" + ` (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in [(ACM)](https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html), provide the Amazon Resource Name (ARN) of the ACM certificate. CloudFront only supports ACM certificates in the US East (N. Virginia) Region (` + "`" + `` + "`" + `us-east-1` + "`" + `` + "`" + `).\n If you specify an ACM certificate ARN, you must also specify values for ` + "`" + `` + "`" + `MinimumProtocolVersion` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SSLSupportMethod` + "`" + `` + "`" + `. (In CloudFormation, the field name is ` + "`" + `` + "`" + `SslSupportMethod` + "`" + `` + "`" + `. Note the different capitalization.)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cloudfront_default_certificate": {
                    "computed": true,
                    "description": "If the distribution uses the CloudFront domain name such as ` + "`" + `` + "`" + `d111111abcdef8.cloudfront.net` + "`" + `` + "`" + `, set this field to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.\n If the distribution uses ` + "`" + `` + "`" + `Aliases` + "`" + `` + "`" + ` (alternate domain names or CNAMEs), omit this field and specify values for the following fields:\n  +   ` + "`" + `` + "`" + `AcmCertificateArn` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `IamCertificateId` + "`" + `` + "`" + ` (specify a value for one, not both) \n  +   ` + "`" + `` + "`" + `MinimumProtocolVersion` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `SslSupportMethod` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "iam_certificate_id": {
                    "computed": true,
                    "description": "In CloudFormation, this field name is ` + "`" + `` + "`" + `IamCertificateId` + "`" + `` + "`" + `. Note the different capitalization.\n  If the distribution uses ` + "`" + `` + "`" + `Aliases` + "`" + `` + "`" + ` (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in [(IAM)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html), provide the ID of the IAM certificate.\n If you specify an IAM certificate ID, you must also specify values for ` + "`" + `` + "`" + `MinimumProtocolVersion` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SSLSupportMethod` + "`" + `` + "`" + `. (In CloudFormation, the field name is ` + "`" + `` + "`" + `SslSupportMethod` + "`" + `` + "`" + `. Note the different capitalization.)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "minimum_protocol_version": {
                    "computed": true,
                    "description": "If the distribution uses ` + "`" + `` + "`" + `Aliases` + "`" + `` + "`" + ` (alternate domain names or CNAMEs), specify the security policy that you want CloudFront to use for HTTPS connections with viewers. The security policy determines two settings:\n  +  The minimum SSL/TLS protocol that CloudFront can use to communicate with viewers.\n  +  The ciphers that CloudFront can use to encrypt the content that it returns to viewers.\n  \n For more information, see [Security Policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValues-security-policy) and [Supported Protocols and Ciphers Between Viewers and CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/secure-connections-supported-viewer-protocols-ciphers.html#secure-connections-supported-ciphers) in the *Amazon CloudFront Developer Guide*.\n  On the CloudFront console, this setting is called *Security Policy*.\n  When you're using SNI only (you set ` + "`" + `` + "`" + `SSLSupportMethod` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `sni-onl",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ssl_support_method": {
                    "computed": true,
                    "description": "In CloudFormation, this field name is ` + "`" + `` + "`" + `SslSupportMethod` + "`" + `` + "`" + `. Note the different capitalization.\n  If the distribution uses ` + "`" + `` + "`" + `Aliases` + "`" + `` + "`" + ` (alternate domain names or CNAMEs), specify which viewers the distribution accepts HTTPS connections from.\n  +   ` + "`" + `` + "`" + `sni-only` + "`" + `` + "`" + ` ? The distribution accepts HTTPS connections from only viewers that support [server name indication (SNI)](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Server_Name_Indication). This is recommended. Most browsers and clients support SNI.\n  +   ` + "`" + `` + "`" + `vip` + "`" + `` + "`" + ` ? The distribution accepts HTTPS connections from all viewers including those that don't support SNI. This is not recommended, and results in additional monthly charges from CloudFront.\n  +   ` + "`" + `` + "`" + `static-ip` + "`" + `` + "`" + ` - Do not specify this value unless your distribution has been enabled for this feature by the CloudFront team. If you have a use case that requires static IP addresses for a distribution, contact CloudFront through the [Center](https://docs.aws.amazon.com/support/home).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "web_acl_id": {
              "computed": true,
              "description": "A unique identifier that specifies the WAF web ACL, if any, to associate with this distribution. To specify a web ACL created using the latest version of WAF, use the ACL ARN, for example ` + "`" + `` + "`" + `arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/473e64fd-f30b-4765-81a0-62ad96dd167a` + "`" + `` + "`" + `. To specify a web ACL created using WAF Classic, use the ACL ID, for example ` + "`" + `` + "`" + `473e64fd-f30b-4765-81a0-62ad96dd167a` + "`" + `` + "`" + `.\n  WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests that are forwarded to CloudFront, and lets you control access to your content. Based on conditions that you specify, such as the IP addresses that requests originate from or the values of query strings, CloudFront responds to requests either with the requested content or with an HTTP 403 status code (Forbidden). You can also configure CloudFront to return a custom error page when a request is blocked. For more information about WAF, see the [Developer Guide](https://docs.aws.amazon.com/waf/latest",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A complex type that contains zero or more ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` elements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string that contains ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` key.\n The string length should be between 1 and 128 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A string that contains an optional ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` value.\n The string length should be between 0 and 256 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontDistributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontDistribution), &result)
	return &result
}
