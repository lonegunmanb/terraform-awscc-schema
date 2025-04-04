package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotAccountAuditConfiguration = `{
  "block": {
    "attributes": {
      "account_id": {
        "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "audit_check_configurations": {
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "device_certificate_age_check": {
              "computed": true,
              "description": "A structure containing the configName and corresponding configValue for configuring DeviceCertAgeCheck.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration": {
                    "computed": true,
                    "description": "A structure containing the configName and corresponding configValue for configuring audit checks.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cert_age_threshold_in_days": {
                          "computed": true,
                          "description": "The configValue for configuring audit checks.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "device_certificate_expiring_check": {
              "computed": true,
              "description": "A structure containing the configName and corresponding configValue for configuring DeviceCertExpirationCheck.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "configuration": {
                    "computed": true,
                    "description": "A structure containing the configName and corresponding configValue for configuring audit checks.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cert_expiration_threshold_in_days": {
                          "computed": true,
                          "description": "The configValue for configuring audit checks.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "enabled": {
                    "computed": true,
                    "description": "True if the check is enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "bool"
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
                    "optional": true,
                    "type": "bool"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The ARN of the role that grants permission to send notifications to the target.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_arn": {
                    "computed": true,
                    "description": "The ARN of the target (SNS topic) to which audit notifications are sent.",
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotAccountAuditConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotAccountAuditConfiguration), &result)
	return &result
}
