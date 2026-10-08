package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaWebFunctionEndpoint = `{
  "block": {
    "attributes": {
      "auth_type": {
        "description": "The authentication type for the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The domain name of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_name": {
        "description": "The name of the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_type": {
        "description": "The type of the endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "description": "The name of the web function this endpoint belongs to. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
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
      "regional_endpoints": {
        "computed": true,
        "description": "Per-region endpoint status and configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_type": {
              "computed": true,
              "description": "The authentication type for the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "domain_name": {
              "computed": true,
              "description": "The domain name of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "revision_weights": {
              "computed": true,
              "description": "The revision weights for the endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "revision_id": {
                    "computed": true,
                    "description": "The revision identifier.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "weight": {
                    "computed": true,
                    "description": "The traffic weight for this revision.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "scaling_config": {
              "computed": true,
              "description": "The scaling configuration for the endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_environments": {
                    "computed": true,
                    "description": "The maximum number of concurrent execution environments for the endpoint. This optional limit further constrains the endpoint's scaling. When omitted, the endpoint's scaling is limited only by your account's vCPU quota.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "state": {
              "computed": true,
              "description": "The current state of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "state_reason": {
              "computed": true,
              "description": "The reason for the current state of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "throttle_config": {
              "computed": true,
              "description": "The throttling configuration for the endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rate_limit": {
                    "computed": true,
                    "description": "The maximum request rate per second for the endpoint, up to a maximum of 10000. This optional limit further constrains the endpoint's request rate. When omitted, the endpoint's request rate is limited only by your account's rate limit quota. Specify 0 to reject all new requests. Other supported values are 100 through 1000 in increments of 100, and 2000 through 10000 in increments of 1000. Supported values can vary by Region; if you specify an unsupported value, the error lists the values available in that Region.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "update_status": {
              "computed": true,
              "description": "The status of the most recent update to the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "update_status_reason": {
              "computed": true,
              "description": "The reason for the current update status of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      },
      "regions": {
        "computed": true,
        "description": "The list of AWS Regions for the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "revision_weights": {
        "computed": true,
        "description": "List of revision routing entries. 1 or 2 entries. With 1 entry, weight must be 100. With 2 entries, weights must sum to 100.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "revision_id": {
              "computed": true,
              "description": "The revision identifier.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description": "The traffic weight for this revision.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "scaling_config": {
        "computed": true,
        "description": "The scaling configuration for the endpoint. Optionally constrains how many concurrent execution environments the endpoint can use, in addition to your account's vCPU quota.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_environments": {
              "computed": true,
              "description": "The maximum number of concurrent execution environments for the endpoint. This optional limit further constrains the endpoint's scaling. When omitted, the endpoint's scaling is limited only by your account's vCPU quota.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "state": {
        "computed": true,
        "description": "The current state of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description": "The reason for the endpoint's current state.",
        "description_kind": "plain",
        "type": "string"
      },
      "throttle_config": {
        "computed": true,
        "description": "The throttling configuration for the endpoint. Optionally constrains the request rate that the endpoint accepts, in addition to your account's rate limit quota.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rate_limit": {
              "computed": true,
              "description": "The maximum request rate per second for the endpoint, up to a maximum of 10000. This optional limit further constrains the endpoint's request rate. When omitted, the endpoint's request rate is limited only by your account's rate limit quota. Specify 0 to reject all new requests. Other supported values are 100 through 1000 in increments of 100, and 2000 through 10000 in increments of 1000. Supported values can vary by Region; if you specify an unsupported value, the error lists the values available in that Region.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "update_status": {
        "computed": true,
        "description": "The status of the last update operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_status_reason": {
        "computed": true,
        "description": "The reason for the update status.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the endpoint was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::WebFunctionEndpoint. An endpoint exposes a Lambda web function over HTTPS and routes traffic to one or more revisions. The endpoint type determines how traffic is served and routed across Regions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaWebFunctionEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaWebFunctionEndpoint), &result)
	return &result
}
