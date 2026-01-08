package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLexBotAlias = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the bot alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_alias_id": {
        "computed": true,
        "description": "Unique ID of resource",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_alias_locale_settings": {
        "computed": true,
        "description": "A list of bot alias locale settings to add to the bot alias.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bot_alias_locale_setting": {
              "computed": true,
              "description": "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code_hook_specification": {
                    "computed": true,
                    "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "lambda_code_hook": {
                          "computed": true,
                          "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "code_hook_interface_version": {
                                "computed": true,
                                "description": "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "lambda_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the Lambda function.",
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
                  "enabled": {
                    "computed": true,
                    "description": "Whether the Lambda code hook is enabled",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "locale_id": {
              "computed": true,
              "description": "A string used to identify the locale",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "bot_alias_name": {
        "computed": true,
        "description": "A unique identifier for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_alias_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bot_alias_tags": {
        "computed": true,
        "description": "A list of tags to add to the bot alias.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "bot_id": {
        "computed": true,
        "description": "Unique ID of resource",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_version": {
        "computed": true,
        "description": "The version of a bot.",
        "description_kind": "plain",
        "type": "string"
      },
      "conversation_log_settings": {
        "computed": true,
        "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audio_log_settings": {
              "computed": true,
              "description": "List of audio log settings",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "The location of audio log files collected when conversation logging is enabled for a bot.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_bucket": {
                          "computed": true,
                          "description": "Specifies an Amazon S3 bucket for logging audio conversations",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "kms_key_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_prefix": {
                                "computed": true,
                                "description": "The Amazon S3 key of the deployment package.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "s3_bucket_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
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
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "text_log_settings": {
              "computed": true,
              "description": "List of text log settings",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cloudwatch": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cloudwatch_log_group_arn": {
                                "computed": true,
                                "description": "A string used to identify the groupArn for the Cloudwatch Log Group",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "log_prefix": {
                                "computed": true,
                                "description": "A string containing the value for the Log Prefix",
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
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the bot alias. Use the description to help identify the bot alias in lists.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sentiment_analysis_settings": {
        "computed": true,
        "description": "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "detect_sentiment": {
              "computed": true,
              "description": "Enable to call Amazon Comprehend for Sentiment natively within Lex",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Lex::BotAlias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLexBotAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLexBotAlias), &result)
	return &result
}
