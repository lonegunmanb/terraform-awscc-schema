package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcorePaymentCredentialProvider = `{
  "block": {
    "attributes": {
      "created_time": {
        "computed": true,
        "description": "The timestamp when the credential provider was created",
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the payment credential provider",
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_vendor": {
        "description": "Supported vendor types for payment providers",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the credential provider was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Unique name for the payment credential provider",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_configuration_input": {
        "computed": true,
        "description": "Provider configuration input containing secrets for creation/update",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "coinbase_cdp_configuration": {
              "computed": true,
              "description": "Coinbase CDP configuration with API credentials",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_id": {
                    "computed": true,
                    "description": "The Coinbase CDP API key ID",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "api_key_secret": {
                    "computed": true,
                    "description": "The Coinbase CDP API key secret",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "api_key_secret_config": {
                    "computed": true,
                    "description": "A reference to a customer-provided secret stored in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "json_key": {
                          "computed": true,
                          "description": "The JSON key within the secret that contains the credential value",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_id": {
                          "computed": true,
                          "description": "The ID or ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "api_key_secret_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "wallet_secret": {
                    "computed": true,
                    "description": "The Coinbase CDP wallet secret",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "wallet_secret_config": {
                    "computed": true,
                    "description": "A reference to a customer-provided secret stored in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "json_key": {
                          "computed": true,
                          "description": "The JSON key within the secret that contains the credential value",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_id": {
                          "computed": true,
                          "description": "The ID or ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "wallet_secret_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "stripe_privy_configuration": {
              "computed": true,
              "description": "Stripe Privy configuration with credentials",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_id": {
                    "computed": true,
                    "description": "The app ID provided by Privy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "app_secret": {
                    "computed": true,
                    "description": "The app secret provided by Privy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "app_secret_config": {
                    "computed": true,
                    "description": "A reference to a customer-provided secret stored in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "json_key": {
                          "computed": true,
                          "description": "The JSON key within the secret that contains the credential value",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_id": {
                          "computed": true,
                          "description": "The ID or ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "app_secret_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "authorization_id": {
                    "computed": true,
                    "description": "The authorization ID for the Stripe Privy integration",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "authorization_private_key": {
                    "computed": true,
                    "description": "The authorization private key for the Stripe Privy integration",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "authorization_private_key_config": {
                    "computed": true,
                    "description": "A reference to a customer-provided secret stored in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "json_key": {
                          "computed": true,
                          "description": "The JSON key within the secret that contains the credential value",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_id": {
                          "computed": true,
                          "description": "The ID or ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "authorization_private_key_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
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
      "provider_configuration_output": {
        "computed": true,
        "description": "Provider configuration output containing secret ARNs (no raw secrets)",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "coinbase_cdp_configuration": {
              "computed": true,
              "description": "Coinbase CDP configuration output with secret ARNs",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_id": {
                    "computed": true,
                    "description": "The Coinbase CDP API key ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "api_key_secret_arn": {
                    "computed": true,
                    "description": "Contains information about a secret in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description": "The ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "api_key_secret_json_key": {
                    "computed": true,
                    "description": "The JSON key within the secret that contains the API key secret value",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "api_key_secret_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "wallet_secret_arn": {
                    "computed": true,
                    "description": "Contains information about a secret in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description": "The ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "wallet_secret_json_key": {
                    "computed": true,
                    "description": "The JSON key within the secret that contains the wallet secret value",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "wallet_secret_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stripe_privy_configuration": {
              "computed": true,
              "description": "Stripe Privy configuration output with secret ARNs",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_id": {
                    "computed": true,
                    "description": "The app ID provided by Privy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "app_secret_arn": {
                    "computed": true,
                    "description": "Contains information about a secret in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description": "The ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "app_secret_json_key": {
                    "computed": true,
                    "description": "The JSON key within the secret that contains the app secret value",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "app_secret_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_id": {
                    "computed": true,
                    "description": "The authorization ID for the Stripe Privy integration",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_private_key_arn": {
                    "computed": true,
                    "description": "Contains information about a secret in AWS Secrets Manager",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_arn": {
                          "computed": true,
                          "description": "The ARN of the secret in AWS Secrets Manager",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "authorization_private_key_json_key": {
                    "computed": true,
                    "description": "The JSON key within the secret that contains the authorization private key value",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_private_key_source": {
                    "computed": true,
                    "description": "The source of the secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.",
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
      "tags": {
        "computed": true,
        "description": "Tags for the payment credential provider",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::PaymentCredentialProvider",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcorePaymentCredentialProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcorePaymentCredentialProvider), &result)
	return &result
}
