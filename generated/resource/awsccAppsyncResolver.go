package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncResolver = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The APSYlong GraphQL API to which you want to attach this resolver.",
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
              "description": "The caching keys for a resolver that has caching activated.\n Valid values are entries from the ` + "`" + `` + "`" + `$context.arguments` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `$context.source` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `$context.identity` + "`" + `` + "`" + ` maps.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "ttl": {
              "description": "The TTL in seconds for a resolver that has caching activated.\n Valid values are 1?3,600 seconds.",
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
        "description": "The ` + "`" + `` + "`" + `resolver` + "`" + `` + "`" + ` code that contains the request and response functions. When code is used, the ` + "`" + `` + "`" + `runtime` + "`" + `` + "`" + ` is required. The runtime value must be ` + "`" + `` + "`" + `APPSYNC_JS` + "`" + `` + "`" + `.",
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
        "description": "The resolver type.\n  +   *UNIT*: A UNIT resolver type. A UNIT resolver is the default resolver type. You can use a UNIT resolver to run a GraphQL query against a single data source.\n  +   *PIPELINE*: A PIPELINE resolver type. You can use a PIPELINE resolver to invoke a series of ` + "`" + `` + "`" + `Function` + "`" + `` + "`" + ` objects in a serial manner. You can use a pipeline resolver to run a GraphQL query against multiple data sources.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_batch_size": {
        "computed": true,
        "description": "The maximum number of resolver request inputs that will be sent to a single LAMlong function in a ` + "`" + `` + "`" + `BatchInvoke` + "`" + `` + "`" + ` operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "metrics_config": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pipeline_config": {
        "computed": true,
        "description": "Functions linked with the pipeline resolver.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "functions": {
              "computed": true,
              "description": "A list of ` + "`" + `` + "`" + `Function` + "`" + `` + "`" + ` objects.",
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
        "description": "The request mapping template.\n Request mapping templates are optional when using a Lambda data source. For all other data sources, a request mapping template is required.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_mapping_template_s3_location": {
        "computed": true,
        "description": "The location of a request mapping template in an S3 bucket. Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resolver_arn": {
        "computed": true,
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
        "description": "The location of a response mapping template in an S3 bucket. Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "runtime": {
        "computed": true,
        "description": "Describes a runtime used by an APSYlong resolver or APSYlong function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "description": "The ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` of the runtime to use. Currently, the only allowed value is ` + "`" + `` + "`" + `APPSYNC_JS` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "runtime_version": {
              "description": "The ` + "`" + `` + "`" + `version` + "`" + `` + "`" + ` of the runtime to use. Currently, the only allowed version is ` + "`" + `` + "`" + `1.0.0` + "`" + `` + "`" + `.",
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
        "description": "The ` + "`" + `` + "`" + `SyncConfig` + "`" + `` + "`" + ` for a resolver attached to a versioned data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "conflict_detection": {
              "description": "The Conflict Detection strategy to use.\n  +   *VERSION*: Detect conflicts based on object versions for this resolver.\n  +   *NONE*: Do not detect conflicts when invoking this resolver.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "conflict_handler": {
              "computed": true,
              "description": "The Conflict Resolution strategy to perform in the event of a conflict.\n  +   *OPTIMISTIC_CONCURRENCY*: Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.\n  +   *AUTOMERGE*: Resolve conflicts with the Automerge conflict resolution strategy.\n  +   *LAMBDA*: Resolve conflicts with an LAMlong function supplied in the ` + "`" + `` + "`" + `LambdaConflictHandlerConfig` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lambda_conflict_handler_config": {
              "computed": true,
              "description": "The ` + "`" + `` + "`" + `LambdaConflictHandlerConfig` + "`" + `` + "`" + ` when configuring ` + "`" + `` + "`" + `LAMBDA` + "`" + `` + "`" + ` as the Conflict Handler.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::AppSync::Resolver` + "`" + `` + "`" + ` resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html).\n  When you submit an update, CFNLong updates resources based on differences between what you submit and the stack's current template. To cause this resource to be updated you must change a property value for this resource in the CFNshort template. Changing the S3 file content without changing a property value will not result in an update operation.\n See [Update Behaviors of Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html) in the *User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncResolverSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncResolver), &result)
	return &result
}
