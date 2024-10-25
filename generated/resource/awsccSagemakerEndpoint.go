package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerEndpoint = `{
  "block": {
    "attributes": {
      "deployment_config": {
        "computed": true,
        "description": "Specifies deployment configuration for updating the SageMaker endpoint. Includes rollback and update policies.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_rollback_configuration": {
              "computed": true,
              "description": "Configuration for automatic rollback if an error occurs during deployment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarms": {
                    "computed": true,
                    "description": "List of CloudWatch alarms to monitor during the deployment. If any alarm goes off, the deployment is rolled back.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_name": {
                          "computed": true,
                          "description": "The name of the CloudWatch alarm.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "blue_green_update_policy": {
              "computed": true,
              "description": "Configuration for blue-green update deployment policies.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum_execution_timeout_in_seconds": {
                    "computed": true,
                    "description": "The maximum time allowed for the blue/green update, in seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "termination_wait_in_seconds": {
                    "computed": true,
                    "description": "The wait time before terminating the old endpoint during a blue/green deployment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "traffic_routing_configuration": {
                    "computed": true,
                    "description": "The traffic routing configuration for the blue/green deployment.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "canary_size": {
                          "computed": true,
                          "description": "Specifies the size of the canary traffic in a canary deployment.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description": "Specifies whether the ` + "`" + `Value` + "`" + ` is an instance count or a capacity unit.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The value representing either the number of instances or the number of capacity units.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "linear_step_size": {
                          "computed": true,
                          "description": "Specifies the step size for linear traffic routing.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description": "Specifies whether the ` + "`" + `Value` + "`" + ` is an instance count or a capacity unit.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The value representing either the number of instances or the number of capacity units.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "type": {
                          "computed": true,
                          "description": "Specifies the type of traffic routing (e.g., 'AllAtOnce', 'Canary', 'Linear').",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "wait_interval_in_seconds": {
                          "computed": true,
                          "description": "Specifies the wait interval between traffic shifts, in seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "rolling_update_policy": {
              "computed": true,
              "description": "Configuration for rolling update deployment policies.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum_batch_size": {
                    "computed": true,
                    "description": "Specifies the maximum batch size for each rolling update.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "Specifies whether the ` + "`" + `Value` + "`" + ` is an instance count or a capacity unit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value representing either the number of instances or the number of capacity units.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "maximum_execution_timeout_in_seconds": {
                    "computed": true,
                    "description": "The maximum time allowed for the rolling update, in seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "rollback_maximum_batch_size": {
                    "computed": true,
                    "description": "The maximum batch size for rollback during an update failure.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "Specifies whether the ` + "`" + `Value` + "`" + ` is an instance count or a capacity unit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value representing either the number of instances or the number of capacity units.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "wait_interval_in_seconds": {
                    "computed": true,
                    "description": "The time to wait between steps during the rolling update, in seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_config_name": {
        "description": "The name of the endpoint configuration for the SageMaker endpoint. This is a required property.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the SageMaker endpoint. This name must be unique within an AWS Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "exclude_retained_variant_properties": {
        "computed": true,
        "description": "Specifies a list of variant properties that you want to exclude when updating an endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "variant_property_type": {
              "computed": true,
              "description": "The type of variant property (e.g., 'DesiredInstanceCount', 'DesiredWeight', 'DataCaptureConfig').",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "retain_all_variant_properties": {
        "computed": true,
        "description": "When set to true, retains all variant properties for an endpoint when it is updated.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "retain_deployment_config": {
        "computed": true,
        "description": "When set to true, retains the deployment configuration during endpoint updates.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
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
    "description": "Resource Type definition for AWS::SageMaker::Endpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerEndpoint), &result)
	return &result
}
