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
        "description": "The source of the API key for metering requests according to a usage plan. Valid values are: ` + "`" + `` + "`" + `HEADER` + "`" + `` + "`" + ` to read the API key from the ` + "`" + `` + "`" + `X-API-Key` + "`" + `` + "`" + ` header of a request. ` + "`" + `` + "`" + `AUTHORIZER` + "`" + `` + "`" + ` to read the API key from the ` + "`" + `` + "`" + `UsageIdentifierKey` + "`" + `` + "`" + ` from a custom authorizer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "binary_media_types": {
        "computed": true,
        "description": "The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.",
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
        "description": "The ID of the RestApi that you want to clone from.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the RestApi.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_execute_api_endpoint": {
        "computed": true,
        "description": "Specifies whether clients can invoke your API by using the default ` + "`" + `` + "`" + `execute-api` + "`" + `` + "`" + ` endpoint. By default, clients can invoke your API with the default ` + "`" + `` + "`" + `https://{api_id}.execute-api.{region}.amazonaws.com` + "`" + `` + "`" + ` endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint_configuration": {
        "computed": true,
        "description": "A list of the endpoint types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the ` + "`" + `` + "`" + `Parameters` + "`" + `` + "`" + ` property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "types": {
              "computed": true,
              "description": "A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is ` + "`" + `` + "`" + `\"EDGE\"` + "`" + `` + "`" + `. For a regional API and its custom domain name, the endpoint type is ` + "`" + `` + "`" + `REGIONAL` + "`" + `` + "`" + `. For a private API, the endpoint type is ` + "`" + `` + "`" + `PRIVATE` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_endpoint_ids": {
              "computed": true,
              "description": "A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes. It is only supported for ` + "`" + `` + "`" + `PRIVATE` + "`" + `` + "`" + ` endpoint type.",
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
        "description": "A query parameter to indicate whether to rollback the API update (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `) or not (` + "`" + `` + "`" + `false` + "`" + `` + "`" + `) when a warning is encountered. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.",
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
        "description": "A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.",
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
        "description": "Custom header parameters as part of the request. For example, to exclude DocumentationParts from an imported API, set ` + "`" + `` + "`" + `ignore=documentation` + "`" + `` + "`" + ` as a ` + "`" + `` + "`" + `parameters` + "`" + `` + "`" + ` value, as in the AWS CLI command of ` + "`" + `` + "`" + `aws apigateway import-rest-api --parameters ignore=documentation --body 'file:///path/to/imported-api-body.json'` + "`" + `` + "`" + `.",
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
        "description": "The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `. The tag value can be up to 256 characters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
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
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::RestApi` + "`" + `` + "`" + ` resource creates a REST API. For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) in the *Amazon API Gateway REST API Reference*.\n On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://docs.aws.amazon.com/https://www.openapis.org/), becoming the foundation of the OpenAPI Specification.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayRestApiSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayRestApi), &result)
	return &result
}
