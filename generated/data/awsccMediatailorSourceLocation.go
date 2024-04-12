package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorSourceLocation = `{
  "block": {
    "attributes": {
      "access_configuration": {
        "computed": true,
        "description": "\u003cp\u003eAccess configuration parameters.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_token_configuration": {
              "computed": true,
              "description": "\u003cp\u003eAWS Secrets Manager access token configuration parameters. For information about Secrets Manager access token authentication, see \u003ca href=\"https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html\"\u003eWorking with AWS Secrets Manager access token authentication\u003c/a\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header_name": {
                    "computed": true,
                    "description": "\u003cp\u003eThe name of the HTTP header used to supply the access token in requests to the source location.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the access token.\u003c/p\u003e",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_string_key": {
                    "computed": true,
                    "description": "\u003cp\u003eThe AWS Secrets Manager \u003ca href=\"https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html#SecretsManager-CreateSecret-request-SecretString.html\"\u003eSecretString\u003c/a\u003e key associated with the access token. MediaTailor uses the key to look up SecretString key and value pair containing the access token.\u003c/p\u003e",
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
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe ARN of the source location.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "default_segment_delivery_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe optional configuration for a server that serves segments. Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_url": {
              "computed": true,
              "description": "\u003cp\u003eThe hostname of the server that will be used to serve segments. This string must include the protocol, such as \u003cb\u003ehttps://\u003c/b\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "http_configuration": {
        "computed": true,
        "description": "\u003cp\u003eThe HTTP configuration for the source location.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_url": {
              "computed": true,
              "description": "\u003cp\u003eThe base URL for the source location host server. This string must include the protocol, such as \u003cb\u003ehttps://\u003c/b\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "segment_delivery_configurations": {
        "computed": true,
        "description": "\u003cp\u003eA list of the segment delivery configurations associated with this resource.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_url": {
              "computed": true,
              "description": "\u003cp\u003eThe base URL of the host or path of the segment delivery server that you're using to serve segments. This is typically a content delivery network (CDN). The URL can be absolute or relative. To use an absolute URL include the protocol, such as \u003ccode\u003ehttps://example.com/some/path\u003c/code\u003e. To use a relative URL specify the relative path, such as \u003ccode\u003e/some/path*\u003c/code\u003e.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "\u003cp\u003eA unique identifier used to distinguish between multiple segment delivery configurations in a source location.\u003c/p\u003e",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "source_location_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the source location.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::MediaTailor::SourceLocation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediatailorSourceLocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorSourceLocation), &result)
	return &result
}
