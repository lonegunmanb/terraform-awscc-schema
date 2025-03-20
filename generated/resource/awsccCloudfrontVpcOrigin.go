package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontVpcOrigin = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
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
              "computed": true,
              "description": "A string that contains ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` key.\n The string length should be between 1 and 128 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string that contains an optional ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` value.\n The string length should be between 0 and 256 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_origin_endpoint_config": {
        "description": "The VPC origin endpoint configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "description": "The ARN of the CloudFront VPC origin endpoint configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "http_port": {
              "computed": true,
              "description": "The HTTP port for the CloudFront VPC origin endpoint configuration. The default value is ` + "`" + `` + "`" + `80` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "https_port": {
              "computed": true,
              "description": "The HTTPS port of the CloudFront VPC origin endpoint configuration. The default value is ` + "`" + `` + "`" + `443` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description": "The name of the CloudFront VPC origin endpoint configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "origin_protocol_policy": {
              "computed": true,
              "description": "The origin protocol policy for the CloudFront VPC origin endpoint configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "origin_ssl_protocols": {
              "computed": true,
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
        "required": true
      },
      "vpc_origin_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "An Amazon CloudFront VPC origin.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontVpcOriginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontVpcOrigin), &result)
	return &result
}
