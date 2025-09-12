package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoicePhoneNumber = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "When set to true the sender ID can't be deleted. By default this is set to false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "iso_country_code": {
        "computed": true,
        "description": "The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.",
        "description_kind": "plain",
        "type": "string"
      },
      "mandatory_keywords": {
        "computed": true,
        "description": "A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords \"HELP\" and \"STOP\" are mandatory keywords",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "help": {
              "computed": true,
              "description": "A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords \"HELP\" and \"STOP\" are mandatory keywords",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stop": {
              "computed": true,
              "description": "A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords \"HELP\" and \"STOP\" are mandatory keywords",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message": {
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
      "number_capabilities": {
        "computed": true,
        "description": "Indicates if the phone number will be used for text messages, voice messages, or both.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "number_type": {
        "computed": true,
        "description": "The type of phone number to request.",
        "description_kind": "plain",
        "type": "string"
      },
      "opt_out_list_name": {
        "computed": true,
        "description": "The name of the OptOutList to associate with the phone number. You can use the OptOutListName or OptOutListArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "optional_keywords": {
        "computed": true,
        "description": "A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "keyword": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "phone_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "phone_number_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "self_managed_opt_outs_enabled": {
        "computed": true,
        "description": "By default this is set to false. When an end recipient sends a message that begins with HELP or STOP to one of your dedicated numbers, AWS End User Messaging SMS and Voice automatically replies with a customizable message and adds the end recipient to the OptOutList. When set to true you're responsible for responding to HELP and STOP requests. You're also responsible for tracking and honoring opt-out requests.",
        "description_kind": "plain",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "two_way": {
        "computed": true,
        "description": "When you set up two-way SMS, you can receive incoming messages from your customers. When one of your customers sends a message to your phone number, the message body is sent to an Amazon SNS topic or Amazon Connect for processing.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "channel_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the two way channel.",
              "description_kind": "plain",
              "type": "string"
            },
            "channel_role": {
              "computed": true,
              "description": "An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.",
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description": "By default this is set to false. When set to true you can receive incoming text messages from your end recipients.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::SMSVOICE::PhoneNumber",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSmsvoicePhoneNumberSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoicePhoneNumber), &result)
	return &result
}
