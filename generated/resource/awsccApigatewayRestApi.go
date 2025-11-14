package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayRestApi = `{
  "block": {
    "attributes": {
      "api_key_source_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "binary_media_types": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "body": {
        "computed": true,
        "description": "An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "body_s3_location": {
        "computed": true,
        "description": "The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description": "The name of the S3 bucket where the OpenAPI file is stored.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "e_tag": {
              "computed": true,
              "description": "The Amazon S3 ETag (a file checksum) of the OpenAPI file. If you don't specify a value, API Gateway skips ETag validation of your OpenAPI file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key": {
              "computed": true,
              "description": "The file name of the OpenAPI file (Amazon S3 object name).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "For versioning-enabled buckets, a specific version of the OpenAPI file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "clone_from": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_execute_api_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint_configuration": {
        "computed": true,
        "description": "A list of the endpoint types and IP address types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ` + "`" + `` + "`" + `Parameters` + "`" + `` + "`" + ` property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "types": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_endpoint_ids": {
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
        "optional": true
      },
      "fail_on_warnings": {
        "computed": true,
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
      "minimum_compression_size": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "mode": {
        "computed": true,
        "description": "This property applies only when you use OpenAPI to define your REST API. The ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` determines how API Gateway handles resource updates.\n Valid values are ` + "`" + `` + "`" + `overwrite` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `merge` + "`" + `` + "`" + `. \n For ` + "`" + `` + "`" + `overwrite` + "`" + `` + "`" + `, the new API definition replaces the existing one. The existing API identifier remains unchanged.\n  For ` + "`" + `` + "`" + `merge` + "`" + `` + "`" + `, the new API definition is merged with the existing API.\n If you don't specify this property, a default value is chosen. For REST APIs created before March 29, 2021, the default is ` + "`" + `` + "`" + `overwrite` + "`" + `` + "`" + `. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API. \n Use the default mode to define top-level ` + "`" + `` + "`" + `RestApi` + "`" + `` + "`" + ` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "A policy document that contains the permissions for the ` + "`" + `` + "`" + `RestApi` + "`" + `` + "`" + ` resource. To set the ARN for the policy, use the ` + "`" + `` + "`" + `!Join` + "`" + `` + "`" + ` intrinsic function with ` + "`" + `` + "`" + `\"\"` + "`" + `` + "`" + ` as delimiter and values of ` + "`" + `` + "`" + `\"execute-api:/\"` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `\"*\"` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "root_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::RestApi` + "`" + `` + "`" + ` resource creates a REST API. For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) in the *Amazon API Gateway REST API Reference*.\n  On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/), becoming the foundation of the OpenAPI Specification.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayRestApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayRestApi), &result)
	return &result
}
