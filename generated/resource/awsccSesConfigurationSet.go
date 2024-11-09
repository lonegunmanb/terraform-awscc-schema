package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesConfigurationSet = `{
  "block": {
    "attributes": {
      "delivery_options": {
        "computed": true,
        "description": "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_delivery_seconds": {
              "computed": true,
              "description": "Specifies the maximum time until which SES will retry sending emails",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sending_pool_name": {
              "computed": true,
              "description": "The name of the dedicated IP pool to associate with the configuration set.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tls_policy": {
              "computed": true,
              "description": "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require , messages are only delivered if a TLS connection can be established. If the value is Optional , messages can be delivered in plain text if a TLS connection can't be established.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      "name": {
        "computed": true,
        "description": "The name of the configuration set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reputation_options": {
        "computed": true,
        "description": "An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reputation_metrics_enabled": {
              "computed": true,
              "description": "If true , tracking of reputation metrics is enabled for the configuration set. If false , tracking of reputation metrics is disabled for the configuration set.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sending_options": {
        "computed": true,
        "description": "An object that defines whether or not Amazon SES can send email that you send using the configuration set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sending_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "suppression_options": {
        "computed": true,
        "description": "An object that contains information about the suppression list preferences for your account.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "suppressed_reasons": {
              "computed": true,
              "description": "A list that contains the reasons that email addresses are automatically added to the suppression list for your account.",
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
      "tracking_options": {
        "computed": true,
        "description": "An object that defines the open and click tracking options for emails that you send using the configuration set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_redirect_domain": {
              "computed": true,
              "description": "The domain to use for tracking open and click events.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "vdm_options": {
        "computed": true,
        "description": "An object that contains Virtual Deliverability Manager (VDM) settings for this configuration set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dashboard_options": {
              "computed": true,
              "description": "Preferences regarding the Dashboard feature.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "engagement_metrics": {
                    "computed": true,
                    "description": "Whether emails sent with this configuration set have engagement tracking enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "guardian_options": {
              "computed": true,
              "description": "Preferences regarding the Guardian feature.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "optimized_shared_delivery": {
                    "computed": true,
                    "description": "Whether emails sent with this configuration set have optimized delivery algorithm enabled.",
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
    "description": "Resource schema for AWS::SES::ConfigurationSet.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesConfigurationSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesConfigurationSet), &result)
	return &result
}
