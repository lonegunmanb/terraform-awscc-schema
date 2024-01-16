package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotAccountAuditConfiguration = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
        "description_kind": "plain",
        "type": "string"
      },
      "audit_check_configurations": {
        "computed": true,
        "description": "Specifies which audit checks are enabled and disabled for this account.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authenticated_cognito_role_overly_permissive_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ca_certificate_expiring_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ca_certificate_key_quality_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "conflicting_client_ids_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "device_certificate_expiring_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "device_certificate_key_quality_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "device_certificate_shared_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "intermediate_ca_revoked_for_active_device_certificates_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "io_t_policy_potential_mis_configuration_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "iot_policy_overly_permissive_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "iot_role_alias_allows_access_to_unused_services_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "iot_role_alias_overly_permissive_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "logging_disabled_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "revoked_ca_certificate_still_active_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "revoked_device_certificate_still_active_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "unauthenticated_cognito_role_overly_permissive_check": {
              "computed": true,
              "description": "The configuration for a specific audit check.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
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
      "audit_notification_target_configurations": {
        "computed": true,
        "description": "Information about the targets to which audit notifications are sent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "True if notifications to the target are enabled.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The ARN of the role that grants permission to send notifications to the target.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_arn": {
                    "computed": true,
                    "description": "The ARN of the target (SNS topic) to which audit notifications are sent.",
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::AccountAuditConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotAccountAuditConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotAccountAuditConfiguration), &result)
	return &result
}
