package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontContinuousDeploymentPolicy = `{
  "block": {
    "attributes": {
      "continuous_deployment_policy_config": {
        "computed": true,
        "description": "Contains the configuration for a continuous deployment policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "A Boolean that indicates whether this continuous deployment policy is enabled (in effect). When this value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, this policy is enabled and in effect. When this value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, this policy is not enabled and has no effect.",
              "description_kind": "plain",
              "type": "bool"
            },
            "single_header_policy_config": {
              "computed": true,
              "description": "This configuration determines which HTTP requests are sent to the staging distribution. If the HTTP request contains a header and value that matches what you specify here, the request is sent to the staging distribution. Otherwise the request is sent to the primary distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "header": {
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
                "nesting_mode": "single"
              }
            },
            "single_weight_policy_config": {
              "computed": true,
              "description": "This configuration determines the percentage of HTTP requests that are sent to the staging distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "session_stickiness_config": {
                    "computed": true,
                    "description": "Session stickiness provides the ability to define multiple requests from a single viewer as a single session. This prevents the potentially inconsistent experience of sending some of a given user's requests to your staging distribution, while others are sent to your primary distribution. Define the session duration using TTL values.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "idle_ttl": {
                          "computed": true,
                          "description": "The amount of time after which you want sessions to cease if no requests are received. Allowed values are 300?3600 seconds (5?60 minutes).",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "maximum_ttl": {
                          "computed": true,
                          "description": "The maximum amount of time to consider requests from the viewer as being part of the same session. Allowed values are 300?3600 seconds (5?60 minutes).",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "weight": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "staging_distribution_dns_names": {
              "computed": true,
              "description": "The CloudFront domain name of the staging distribution. For example: ` + "`" + `` + "`" + `d111111abcdef8.cloudfront.net` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "traffic_config": {
              "computed": true,
              "description": "Contains the parameters for routing production traffic from your primary to staging distributions.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "single_header_config": {
                    "computed": true,
                    "description": "Determines which HTTP requests are sent to the staging distribution.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "header": {
                          "computed": true,
                          "description": "The request header name that you want CloudFront to send to your staging distribution. The header must contain the prefix ` + "`" + `` + "`" + `aws-cf-cd-` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The request header value.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "single_weight_config": {
                    "computed": true,
                    "description": "Contains the percentage of traffic to send to the staging distribution.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "session_stickiness_config": {
                          "computed": true,
                          "description": "Session stickiness provides the ability to define multiple requests from a single viewer as a single session. This prevents the potentially inconsistent experience of sending some of a given user's requests to your staging distribution, while others are sent to your primary distribution. Define the session duration using TTL values.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "idle_ttl": {
                                "computed": true,
                                "description": "The amount of time after which you want sessions to cease if no requests are received. Allowed values are 300?3600 seconds (5?60 minutes).",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "maximum_ttl": {
                                "computed": true,
                                "description": "The maximum amount of time to consider requests from the viewer as being part of the same session. Allowed values are 300?3600 seconds (5?60 minutes).",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "weight": {
                          "computed": true,
                          "description": "The percentage of traffic to send to a staging distribution, expressed as a decimal number between 0 and 0.15. For example, a value of 0.10 means 10% of traffic is sent to the staging distribution.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of traffic configuration.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description": "The type of traffic configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "continuous_deployment_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
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
      }
    },
    "description": "Data Source schema for AWS::CloudFront::ContinuousDeploymentPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontContinuousDeploymentPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontContinuousDeploymentPolicy), &result)
	return &result
}
