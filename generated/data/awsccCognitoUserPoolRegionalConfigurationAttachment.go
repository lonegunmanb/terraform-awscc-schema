package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolRegionalConfigurationAttachment = `{
  "block": {
    "attributes": {
      "email_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration_set": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "email_sending_account": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "from": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "reply_to_email_address": {
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
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lambda_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "create_auth_challenge": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "custom_email_sender": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "lambda_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "custom_message": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "custom_sms_sender": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "lambda_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "define_auth_challenge": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "inbound_federation": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "lambda_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "post_authentication": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "post_confirmation": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "pre_authentication": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "pre_sign_up": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "pre_token_generation": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "pre_token_generation_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "lambda_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "user_migration": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "verify_auth_challenge_response": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sms_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "external_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sns_caller_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sns_region": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the replica. Set to ACTIVE or INACTIVE.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolRegionalConfigurationAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolRegionalConfigurationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolRegionalConfigurationAttachment), &result)
	return &result
}
