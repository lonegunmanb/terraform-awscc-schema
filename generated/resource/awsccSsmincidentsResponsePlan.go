package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmincidentsResponsePlan = `{
  "block": {
    "attributes": {
      "actions": {
        "computed": true,
        "description": "The list of actions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ssm_automation": {
              "computed": true,
              "description": "The configuration to use when starting the SSM automation document.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "document_name": {
                    "description": "The document name to use when starting the SSM automation document.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "document_version": {
                    "computed": true,
                    "description": "The version of the document to use when starting the SSM automation document.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "dynamic_parameters": {
                    "computed": true,
                    "description": "The parameters with dynamic values to set when starting the SSM automation document.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "Value of the dynamic parameter to set when starting the SSM automation document.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "variable": {
                                "computed": true,
                                "description": "The variable types used as dynamic parameter value when starting the SSM automation document.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  },
                  "parameters": {
                    "computed": true,
                    "description": "The parameters to set when starting the SSM automation document.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  },
                  "role_arn": {
                    "description": "The role ARN to use when starting the SSM automation document.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "target_account": {
                    "computed": true,
                    "description": "The account type to use when starting the SSM automation document.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the response plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "chat_channel": {
        "computed": true,
        "description": "The chat channel configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "chatbot_sns": {
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
      "display_name": {
        "computed": true,
        "description": "The display name of the response plan.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engagements": {
        "computed": true,
        "description": "The list of engagements to use.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "incident_template": {
        "description": "The incident template configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dedupe_string": {
              "computed": true,
              "description": "The deduplication string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "impact": {
              "description": "The impact value.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "incident_tags": {
              "computed": true,
              "description": "Tags that get applied to incidents created by the StartIncident API action.",
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
                "nesting_mode": "set"
              },
              "optional": true
            },
            "notification_targets": {
              "computed": true,
              "description": "The list of notification targets.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sns_topic_arn": {
                    "computed": true,
                    "description": "The ARN of the Chatbot SNS topic.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "summary": {
              "computed": true,
              "description": "The summary string.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "title": {
              "description": "The title string.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "integrations": {
        "computed": true,
        "description": "The list of integrations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pager_duty_configuration": {
              "computed": true,
              "description": "The pagerDuty configuration to use when starting the incident.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "description": "The name of the pagerDuty configuration.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "pager_duty_incident_configuration": {
                    "description": "The pagerDuty incident configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "service_id": {
                          "description": "The pagerDuty serviceId.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "secret_id": {
                    "description": "The AWS secrets manager secretId storing the pagerDuty token.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "name": {
        "description": "The name of the response plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the response plan.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource type definition for AWS::SSMIncidents::ResponsePlan",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmincidentsResponsePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmincidentsResponsePlan), &result)
	return &result
}
