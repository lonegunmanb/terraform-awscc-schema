package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolRiskConfigurationAttachment = `{
  "block": {
    "attributes": {
      "account_takeover_risk_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "high_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "notify": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "low_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "notify": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "medium_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "notify": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "notify_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "block_email": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "html_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "subject": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "from": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mfa_email": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "html_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "subject": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "no_action_email": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "html_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "subject": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_body": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "reply_to": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_arn": {
                    "computed": true,
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
      "client_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compromised_credentials_risk_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "event_filter": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "risk_exception_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "blocked_ip_range_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "skipped_ip_range_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolRiskConfigurationAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolRiskConfigurationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolRiskConfigurationAttachment), &result)
	return &result
}
