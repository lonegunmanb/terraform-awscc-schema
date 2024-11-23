package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomMessageTemplate = `{
  "block": {
    "attributes": {
      "channel_subtype": {
        "description": "The channel subtype this message template applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content": {
        "description": "The content of the message template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email_message_template_content": {
              "computed": true,
              "description": "The content of message template that applies to email channel subtype.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "body": {
                    "computed": true,
                    "description": "The body to use in email messages.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "html": {
                          "computed": true,
                          "description": "The message body, in HTML format, to use in email messages that are based on the message template. We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "plain_text": {
                          "computed": true,
                          "description": "The message body, in plain text format, to use in email messages that are based on the message template. We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content": {
                                "computed": true,
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
                  "headers": {
                    "computed": true,
                    "description": "The email headers to include in email messages.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the email header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value of the email header.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "subject": {
                    "computed": true,
                    "description": "The subject line, or title, to use in email messages.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sms_message_template_content": {
              "computed": true,
              "description": "The content of message template that applies to SMS channel subtype.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "body": {
                    "computed": true,
                    "description": "The body to use in SMS messages.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "plain_text": {
                          "computed": true,
                          "description": "The container of message template body.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content": {
                                "computed": true,
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "default_attributes": {
        "computed": true,
        "description": "An object that specifies the default values to use for variables in the message template. This object contains different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The corresponding value defines the default value for that variable.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "agent_attributes": {
              "computed": true,
              "description": "The agent attributes that are used with the message template.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "first_name": {
                    "computed": true,
                    "description": "The agent?s first name as entered in their Amazon Connect user account.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "last_name": {
                    "computed": true,
                    "description": "The agent?s last name as entered in their Amazon Connect user account.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "custom_attributes": {
              "computed": true,
              "description": "The custom attributes that are used with the message template.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "customer_profile_attributes": {
              "computed": true,
              "description": "The customer profile attributes that are used with the message template.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_number": {
                    "computed": true,
                    "description": "A unique account number that you have given to the customer.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "additional_information": {
                    "computed": true,
                    "description": "Any additional information relevant to the customer's profile.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "address_1": {
                    "computed": true,
                    "description": "The first line of a customer address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "address_2": {
                    "computed": true,
                    "description": "The second line of a customer address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "address_3": {
                    "computed": true,
                    "description": "The third line of a customer address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "address_4": {
                    "computed": true,
                    "description": "The fourth line of a customer address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_address_1": {
                    "computed": true,
                    "description": "The first line of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_address_2": {
                    "computed": true,
                    "description": "The second line of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_address_3": {
                    "computed": true,
                    "description": "The third line of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_address_4": {
                    "computed": true,
                    "description": "The fourth line of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_city": {
                    "computed": true,
                    "description": "The city of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_country": {
                    "computed": true,
                    "description": "The country of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_county": {
                    "computed": true,
                    "description": "The county of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_postal_code": {
                    "computed": true,
                    "description": "The postal code of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_province": {
                    "computed": true,
                    "description": "The province of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "billing_state": {
                    "computed": true,
                    "description": "The state of a customer?s billing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "birth_date": {
                    "computed": true,
                    "description": "The customer's birth date.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "business_email_address": {
                    "computed": true,
                    "description": "The customer's business email address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "business_name": {
                    "computed": true,
                    "description": "The name of the customer's business.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "business_phone_number": {
                    "computed": true,
                    "description": "The customer's business phone number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "city": {
                    "computed": true,
                    "description": "The city in which a customer lives.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "country": {
                    "computed": true,
                    "description": "The country in which a customer lives.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "county": {
                    "computed": true,
                    "description": "The county in which a customer lives.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom": {
                    "computed": true,
                    "description": "The custom attributes that are used with the message template.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "email_address": {
                    "computed": true,
                    "description": "The customer's email address, which has not been specified as a personal or business address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "first_name": {
                    "computed": true,
                    "description": "The customer's first name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "gender": {
                    "computed": true,
                    "description": "The customer's gender.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "home_phone_number": {
                    "computed": true,
                    "description": "The customer's home phone number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "last_name": {
                    "computed": true,
                    "description": "The customer's last name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_address_1": {
                    "computed": true,
                    "description": "The first line of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_address_2": {
                    "computed": true,
                    "description": "The second line of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_address_3": {
                    "computed": true,
                    "description": "The third line of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_address_4": {
                    "computed": true,
                    "description": "The fourth line of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_city": {
                    "computed": true,
                    "description": "The city of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_country": {
                    "computed": true,
                    "description": "The country of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_county": {
                    "computed": true,
                    "description": "The county of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_postal_code": {
                    "computed": true,
                    "description": "The postal code of a customer?s mailing address",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_province": {
                    "computed": true,
                    "description": "The province of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mailing_state": {
                    "computed": true,
                    "description": "The state of a customer?s mailing address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "middle_name": {
                    "computed": true,
                    "description": "The customer's middle name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "mobile_phone_number": {
                    "computed": true,
                    "description": "The customer's mobile phone number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "party_type": {
                    "computed": true,
                    "description": "The customer's party type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "phone_number": {
                    "computed": true,
                    "description": "The customer's phone number, which has not been specified as a mobile, home, or business number.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "postal_code": {
                    "computed": true,
                    "description": "The postal code of a customer address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "profile_arn": {
                    "computed": true,
                    "description": "The ARN of a customer profile.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "profile_id": {
                    "computed": true,
                    "description": "The unique identifier of a customer profile.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "province": {
                    "computed": true,
                    "description": "The province in which a customer lives.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_address_1": {
                    "computed": true,
                    "description": "The first line of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_address_2": {
                    "computed": true,
                    "description": "The second line of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_address_3": {
                    "computed": true,
                    "description": "The third line of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_address_4": {
                    "computed": true,
                    "description": "The fourth line of a customer?s shipping address",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_city": {
                    "computed": true,
                    "description": "The city of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_country": {
                    "computed": true,
                    "description": "The country of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_county": {
                    "computed": true,
                    "description": "The county of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_postal_code": {
                    "computed": true,
                    "description": "The postal code of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_province": {
                    "computed": true,
                    "description": "The province of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "shipping_state": {
                    "computed": true,
                    "description": "The state of a customer?s shipping address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "The state in which a customer lives.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "system_attributes": {
              "computed": true,
              "description": "The system attributes that are used with the message template.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "customer_endpoint": {
                    "computed": true,
                    "description": "The CustomerEndpoint attribute.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "address": {
                          "computed": true,
                          "description": "The customer's phone number if used with customerEndpoint, or the number the customer dialed to call your contact center if used with systemEndpoint.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the task.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "system_endpoint": {
                    "computed": true,
                    "description": "The SystemEndpoint attribute.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "address": {
                          "computed": true,
                          "description": "The customer's phone number if used with customerEndpoint, or the number the customer dialed to call your contact center if used with systemEndpoint.",
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of the message template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "grouping_configuration": {
        "computed": true,
        "description": "The configuration information of the user groups that the message template is accessible to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criteria": {
              "computed": true,
              "description": "The criteria used for grouping Amazon Q in Connect users.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
              "computed": true,
              "description": "The list of values that define different groups of Amazon Q in Connect users.",
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_arn": {
        "description": "The Amazon Resource Name (ARN) of the knowledge base to which the message template belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language": {
        "computed": true,
        "description": "The language code value for the language in which the message template is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "message_template_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the message template.",
        "description_kind": "plain",
        "type": "string"
      },
      "message_template_content_sha_256": {
        "computed": true,
        "description": "The content SHA256 of the message template.",
        "description_kind": "plain",
        "type": "string"
      },
      "message_template_id": {
        "computed": true,
        "description": "The unique identifier of the message template.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the message template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource. For example, { \"tags\": {\"key1\":\"value1\", \"key2\":\"value2\"} }.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::Wisdom::MessageTemplate Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWisdomMessageTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomMessageTemplate), &result)
	return &result
}
