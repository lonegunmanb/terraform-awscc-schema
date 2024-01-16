package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncResolver = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The AWS AppSync GraphQL API to which you want to attach this resolver.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "caching_config": {
        "computed": true,
        "description": "The caching configuration for the resolver.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "caching_keys": {
              "computed": true,
              "description": "The caching keys for a resolver that has caching activated. Valid values are entries from the $context.arguments, $context.source, and $context.identity maps.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ttl": {
              "description": "The TTL in seconds for a resolver that has caching activated. Valid values are 1-36.00 seconds.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "code": {
        "computed": true,
        "description": "The resolver code that contains the request and response functions. When code is used, the runtime is required.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "code_s3_location": {
        "computed": true,
        "description": "The Amazon S3 endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_name": {
        "computed": true,
        "description": "The resolver data source name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "field_name": {
        "description": "The GraphQL field on a type that invokes the resolver.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description": "The resolver type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_batch_size": {
        "computed": true,
        "description": "The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "pipeline_config": {
        "computed": true,
        "description": "Functions linked with the pipeline resolver.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "functions": {
              "computed": true,
              "description": "A list of Function objects.",
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
      "request_mapping_template": {
        "computed": true,
        "description": "Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_mapping_template_s3_location": {
        "computed": true,
        "description": "The location of a request mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resolver_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the resolver.",
        "description_kind": "plain",
        "type": "string"
      },
      "response_mapping_template": {
        "computed": true,
        "description": "The response mapping template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "response_mapping_template_s3_location": {
        "computed": true,
        "description": "The location of a response mapping template in an Amazon S3 bucket. Use this if you want to provision with a template file in Amazon S3 rather than embedding it in your CloudFormation template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "computed": true,
        "description": "Describes a runtime used by an AWS AppSync pipeline resolver or AWS AppSync function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The name of the runtime to use.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "runtime_version": {
              "description": "The version of the runtime to use.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sync_config": {
        "computed": true,
        "description": "The SyncConfig for a resolver attached to a versioned data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "conflict_detection": {
              "description": "The Conflict Detection strategy to use.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "conflict_handler": {
              "computed": true,
              "description": "The Conflict Resolution strategy to perform in the event of a conflict.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lambda_conflict_handler_config": {
              "computed": true,
              "description": "The LambdaConflictHandlerConfig when configuring LAMBDA as the Conflict Handler.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_conflict_handler_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.",
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
      "type_name": {
        "description": "The GraphQL type that invokes this resolver.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppSync::Resolver",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncResolverSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncResolver), &result)
	return &result
}
