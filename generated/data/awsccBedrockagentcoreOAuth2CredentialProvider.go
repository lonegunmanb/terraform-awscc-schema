package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreOAuth2CredentialProvider = `{
  "block": {
    "attributes": {
      "callback_url": {
        "computed": true,
        "description": "The callback URL for the OAuth2 authorization flow",
        "description_kind": "plain",
        "type": "string"
      },
      "client_secret_arn": {
        "computed": true,
        "description": "The ARN of the client secret in AWS Secrets Manager",
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
      "created_time": {
        "computed": true,
        "description": "The timestamp when the credential provider was created",
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the OAuth2 credential provider",
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_vendor": {
        "computed": true,
        "description": "The vendor of the OAuth2 credential provider",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the credential provider was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the OAuth2 credential provider",
        "description_kind": "plain",
        "type": "string"
      },
      "oauth_2_provider_config_input": {
        "computed": true,
        "description": "The configuration settings for the OAuth2 provider",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "atlassian_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for an Atlassian OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "custom_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a custom OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description": "The client ID for the custom OAuth2 provider",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description": "The client secret for the custom OAuth2 provider",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "oauth_discovery": {
                    "computed": true,
                    "description": "Discovery information for an OAuth2 provider",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization_server_metadata": {
                          "computed": true,
                          "description": "Authorization server metadata for the OAuth2 provider",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "authorization_endpoint": {
                                "computed": true,
                                "description": "The authorization endpoint URL",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "issuer": {
                                "computed": true,
                                "description": "The issuer URL for the OAuth2 authorization server",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "response_types": {
                                "computed": true,
                                "description": "The supported response types",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "token_endpoint": {
                                "computed": true,
                                "description": "The token endpoint URL",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "discovery_url": {
                          "computed": true,
                          "description": "The discovery URL for the OAuth2 provider",
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
            "github_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a GitHub OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "google_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a Google OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "included_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a supported non-custom OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_endpoint": {
                    "computed": true,
                    "description": "OAuth2 authorization endpoint for your isolated OAuth2 application tenant",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "issuer": {
                    "computed": true,
                    "description": "Token issuer of your isolated OAuth2 application tenant",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token_endpoint": {
                    "computed": true,
                    "description": "OAuth2 token endpoint for your isolated OAuth2 application tenant",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "linkedin_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a LinkedIn OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "microsoft_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a Microsoft OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tenant_id": {
                    "computed": true,
                    "description": "The Microsoft Entra ID tenant ID",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "salesforce_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a Salesforce OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "slack_oauth_2_provider_config": {
              "computed": true,
              "description": "Input configuration for a Slack OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_secret": {
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
      "oauth_2_provider_config_output": {
        "computed": true,
        "description": "The output configuration for the OAuth2 provider",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "oauth_discovery": {
              "computed": true,
              "description": "Discovery information for an OAuth2 provider",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_server_metadata": {
                    "computed": true,
                    "description": "Authorization server metadata for the OAuth2 provider",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization_endpoint": {
                          "computed": true,
                          "description": "The authorization endpoint URL",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "issuer": {
                          "computed": true,
                          "description": "The issuer URL for the OAuth2 authorization server",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "response_types": {
                          "computed": true,
                          "description": "The supported response types",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "token_endpoint": {
                          "computed": true,
                          "description": "The token endpoint URL",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "discovery_url": {
                    "computed": true,
                    "description": "The discovery URL for the OAuth2 provider",
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
        "description": "Tags to assign to the OAuth2 credential provider",
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
    "description": "Data Source schema for AWS::BedrockAgentCore::OAuth2CredentialProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreOAuth2CredentialProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreOAuth2CredentialProvider), &result)
	return &result
}
