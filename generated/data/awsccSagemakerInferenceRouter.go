package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerInferenceRouter = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time when the Inference Router was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inference_router_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Inference Router.",
        "description_kind": "plain",
        "type": "string"
      },
      "inference_router_name": {
        "computed": true,
        "description": "The name of the Inference Router. Unique per account per Region, case-sensitive. If you don't specify a name, CloudFormation generates one.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time when the Inference Router was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_targets": {
        "computed": true,
        "description": "The routing targets of the Inference Router. An update replaces the full list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "async_inference_config": {
              "computed": true,
              "description": "Set this when the routing target's destinations are asynchronous inference endpoints. A routing target without it serves real-time requests.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_config": {
                    "computed": true,
                    "description": "How the Inference Router submits a request to a destination.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_time_in_backlog_seconds": {
                          "computed": true,
                          "description": "How long, in seconds, a request can wait for a destination to pick it up before it expires. The default is 900.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "output_config": {
                    "computed": true,
                    "description": "Where the results of asynchronous requests are written. Every destination of the routing target writes to these paths.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kms_key_id": {
                          "computed": true,
                          "description": "The AWS KMS key that encrypts the results. If any destination is in a different account or Region from the Inference Router, use a key ARN or an alias ARN.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "notification_config": {
                          "computed": true,
                          "description": "The SNS topics the Inference Router publishes a completion to. The notification carries the location of the result, not the result.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "error_topic": {
                                "computed": true,
                                "description": "The ARN of a standard (not FIFO) SNS topic to notify when a request fails. It can be in any Region.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "success_topic": {
                                "computed": true,
                                "description": "The ARN of a standard (not FIFO) SNS topic to notify when a request succeeds. It can be in any Region.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "s3_failure_path": {
                          "computed": true,
                          "description": "The Amazon S3 location for failed requests. It can be in any Region and account.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_output_path": {
                          "computed": true,
                          "description": "The Amazon S3 location for successful results. It can be in any Region and account.",
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
            "destinations": {
              "computed": true,
              "description": "The endpoints and inference components that serve this routing target.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "resource_arn": {
                    "computed": true,
                    "description": "The ARN of a SageMaker endpoint or inference component that serves the routing target. It can be in any Region and account within the Inference Router's partition.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The logical model name a caller passes at invoke time. Unique within an Inference Router.",
              "description_kind": "plain",
              "type": "string"
            },
            "routing_preference": {
              "computed": true,
              "description": "How the router prefers this routing target's destinations. LocalRegionFirst (the default) uses destinations in the router's Region first, then the others, each group in the order listed. InstancePoolOrder, for asynchronous routing targets only, tries destinations in order of the best instance pool they can serve from now. In both, degraded destinations are used after the others. If omitted, LocalRegionFirst is used and returned.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description": "Where the Inference Router's most recent configuration write has got to. Writes commit synchronously; propagation to the data plane is asynchronous.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the Inference Router.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Tag keys must be unique per resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::InferenceRouter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerInferenceRouterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerInferenceRouter), &result)
	return &result
}
