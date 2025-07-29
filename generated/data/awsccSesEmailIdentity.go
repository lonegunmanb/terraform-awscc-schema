package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesEmailIdentity = `{
  "block": {
    "attributes": {
      "configuration_set_attributes": {
        "computed": true,
        "description": "Used to associate a configuration set with an email identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration_set_name": {
              "computed": true,
              "description": "The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "dkim_attributes": {
        "computed": true,
        "description": "Used to enable or disable DKIM authentication for an email identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "signing_enabled": {
              "computed": true,
              "description": "Sets the DKIM signing configuration for the identity. When you set this value true, then the messages that are sent from the identity are signed using DKIM. If you set this value to false, your messages are sent without DKIM signing.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "dkim_dns_token_name_1": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_dns_token_name_2": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_dns_token_name_3": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_dns_token_value_1": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_dns_token_value_2": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_dns_token_value_3": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dkim_signing_attributes": {
        "computed": true,
        "description": "If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain_signing_private_key": {
              "computed": true,
              "description": "[Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.",
              "description_kind": "plain",
              "type": "string"
            },
            "domain_signing_selector": {
              "computed": true,
              "description": "[Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.",
              "description_kind": "plain",
              "type": "string"
            },
            "next_signing_key_length": {
              "computed": true,
              "description": "[Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "email_identity": {
        "computed": true,
        "description": "The email address or domain to verify.",
        "description_kind": "plain",
        "type": "string"
      },
      "feedback_attributes": {
        "computed": true,
        "description": "Used to enable or disable feedback forwarding for an identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email_forwarding_enabled": {
              "computed": true,
              "description": "If the value is true, you receive email notifications when bounce or complaint events occur",
              "description_kind": "plain",
              "type": "bool"
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
      "mail_from_attributes": {
        "computed": true,
        "description": "Used to enable or disable the custom Mail-From domain configuration for an email identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "behavior_on_mx_failure": {
              "computed": true,
              "description": "The action to take if the required MX record isn't found when you send an email. When you set this value to UseDefaultValue , the mail is sent using amazonses.com as the MAIL FROM domain. When you set this value to RejectMessage , the Amazon SES API v2 returns a MailFromDomainNotVerified error, and doesn't attempt to deliver the email.",
              "description_kind": "plain",
              "type": "string"
            },
            "mail_from_domain": {
              "computed": true,
              "description": "The custom MAIL FROM domain that you want the verified identity to use",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the email identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SES::EmailIdentity",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesEmailIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesEmailIdentity), &result)
	return &result
}
