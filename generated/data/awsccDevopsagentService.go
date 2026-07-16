package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentService = `{
  "block": {
    "attributes": {
      "accessible_resources": {
        "computed": true,
        "description": "List of accessible resources for this service",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "map",
            "string"
          ]
        ]
      },
      "additional_service_details": {
        "computed": true,
        "description": "Additional details specific to the service type",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure_identity": {
              "computed": true,
              "description": "Azure Identity service details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description": "Azure AD application client ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tenant_id": {
                    "computed": true,
                    "description": "Azure AD tenant ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "web_identity_role_arn": {
                    "computed": true,
                    "description": "ARN of the IAM role for web identity token exchange",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "web_identity_token_audiences": {
                    "computed": true,
                    "description": "List of audiences for the web identity token",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "dynatrace": {
              "computed": true,
              "description": "Dynatrace service details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_urn": {
                    "computed": true,
                    "description": "Dynatrace resource account URN",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "git_lab": {
              "computed": true,
              "description": "GitLab service details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_id": {
                    "computed": true,
                    "description": "Optional GitLab group ID for group-level access tokens",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_url": {
                    "computed": true,
                    "description": "GitLab instance URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token_type": {
                    "computed": true,
                    "description": "Type of GitLab access token",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server": {
              "computed": true,
              "description": "MCP server details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_header": {
                    "computed": true,
                    "description": "API key header name if using API key authentication",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_method": {
                    "computed": true,
                    "description": "MCP server authorization method",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_grafana": {
              "computed": true,
              "description": "Grafana MCP server details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_method": {
                    "computed": true,
                    "description": "MCP server authorization method",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_new_relic": {
              "computed": true,
              "description": "New Relic service details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description": "New Relic account ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional user description",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "New Relic region",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_sig_v4": {
              "computed": true,
              "description": "SigV4-authenticated MCP server details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_headers": {
                    "computed": true,
                    "description": "Custom headers for the SigV4 MCP server",
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "The MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mcp_role_arn": {
                    "computed": true,
                    "description": "IAM role ARN for SigV4 signing. Absent when no dedicated role is configured.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "AWS region for SigV4 signing",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "Deprecated - use McpRoleArn instead. IAM role ARN for SigV4 signing",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "service": {
                    "computed": true,
                    "description": "AWS service name for SigV4 signing",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_splunk": {
              "computed": true,
              "description": "MCP server details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "api_key_header": {
                    "computed": true,
                    "description": "API key header name if using API key authentication",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_method": {
                    "computed": true,
                    "description": "MCP server authorization method",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "pager_duty": {
              "computed": true,
              "description": "PagerDuty service details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "scopes": {
                    "computed": true,
                    "description": "The scopes assigned to the service",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_now": {
              "computed": true,
              "description": "ServiceNow service details returned after registration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_url": {
                    "computed": true,
                    "description": "ServiceNow instance URL",
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
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the KMS key to use for encryption.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_details": {
        "computed": true,
        "description": "Service-specific configuration details for create operation",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "azure_identity": {
              "computed": true,
              "description": "Azure Identity service configuration for federated identity",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_id": {
                    "computed": true,
                    "description": "Azure AD application client ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tenant_id": {
                    "computed": true,
                    "description": "Azure AD tenant ID",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "web_identity_role_arn": {
                    "computed": true,
                    "description": "ARN of the IAM role for web identity token exchange",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "web_identity_token_audiences": {
                    "computed": true,
                    "description": "List of audiences for the web identity token",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "dynatrace": {
              "computed": true,
              "description": "Dynatrace service configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_urn": {
                    "computed": true,
                    "description": "Dynatrace resource account URN",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_config": {
                    "computed": true,
                    "description": "Dynatrace OAuth authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "o_auth_client_credentials": {
                          "computed": true,
                          "description": "OAuth client credentials",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "client_id": {
                                "computed": true,
                                "description": "OAuth client ID",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_name": {
                                "computed": true,
                                "description": "User friendly OAuth client name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_secret": {
                                "computed": true,
                                "description": "OAuth client secret",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "exchange_parameters": {
                                "computed": true,
                                "description": "OAuth token exchange parameters",
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
                  }
                },
                "nesting_mode": "single"
              }
            },
            "git_lab": {
              "computed": true,
              "description": "GitLab service configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_id": {
                    "computed": true,
                    "description": "Optional GitLab group ID for group-level access tokens",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_url": {
                    "computed": true,
                    "description": "GitLab instance URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token_type": {
                    "computed": true,
                    "description": "Type of GitLab access token",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token_value": {
                    "computed": true,
                    "description": "GitLab access token value",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server": {
              "computed": true,
              "description": "MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "MCP server authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_key": {
                          "computed": true,
                          "description": "API key authentication details",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "api_key_header": {
                                "computed": true,
                                "description": "HTTP header name to send the API key",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "api_key_name": {
                                "computed": true,
                                "description": "User friendly API key name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "api_key_value": {
                                "computed": true,
                                "description": "API key value",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "o_auth_client_credentials": {
                          "computed": true,
                          "description": "MCP server OAuth client credentials configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "client_id": {
                                "computed": true,
                                "description": "OAuth client ID",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_name": {
                                "computed": true,
                                "description": "User friendly OAuth client name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_secret": {
                                "computed": true,
                                "description": "OAuth client secret",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "exchange_parameters": {
                                "computed": true,
                                "description": "OAuth token exchange parameters",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "exchange_url": {
                                "computed": true,
                                "description": "OAuth token exchange URL",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "scopes": {
                                "computed": true,
                                "description": "OAuth scopes",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_grafana": {
              "computed": true,
              "description": "Grafana MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "Grafana MCP server authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bearer_token": {
                          "computed": true,
                          "description": "Bearer token authentication details",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "authorization_header": {
                                "computed": true,
                                "description": "HTTP header name to send the bearer token",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "token_name": {
                                "computed": true,
                                "description": "User friendly bearer token name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "token_value": {
                                "computed": true,
                                "description": "Bearer token value",
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
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_new_relic": {
              "computed": true,
              "description": "New Relic service configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "New Relic authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_key": {
                          "computed": true,
                          "description": "New Relic API key configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "account_id": {
                                "computed": true,
                                "description": "New Relic Account ID",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "alert_policy_ids": {
                                "computed": true,
                                "description": "List of alert policy IDs",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "api_key": {
                                "computed": true,
                                "description": "New Relic User API Key",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "application_ids": {
                                "computed": true,
                                "description": "List of monitored APM application IDs",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "entity_guids": {
                                "computed": true,
                                "description": "List of globally unique IDs for New Relic resources",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "region": {
                                "computed": true,
                                "description": "New Relic region",
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
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_sig_v4": {
              "computed": true,
              "description": "SigV4-authenticated MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "SigV4 authorization configuration for MCP server",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "custom_headers": {
                          "computed": true,
                          "description": "Custom headers for the SigV4 MCP server",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "mcp_role_arn": {
                          "computed": true,
                          "description": "IAM role ARN to assume for SigV4 signing. Optional - when omitted, credentials are resolved at runtime via a monitor account association.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "region": {
                          "computed": true,
                          "description": "AWS region for SigV4 signing. Use '*' for SigV4a multi-region signing.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "Deprecated - use McpRoleArn instead. IAM role ARN to assume for SigV4 signing",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "service": {
                          "computed": true,
                          "description": "AWS service name for SigV4 signing",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mcp_server_splunk": {
              "computed": true,
              "description": "Splunk MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "MCP server splunk authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bearer_token": {
                          "computed": true,
                          "description": "Bearer token authentication details",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "authorization_header": {
                                "computed": true,
                                "description": "HTTP header name to send the bearer token",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "token_name": {
                                "computed": true,
                                "description": "User friendly bearer token name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "token_value": {
                                "computed": true,
                                "description": "Bearer token value",
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
                  "description": {
                    "computed": true,
                    "description": "Optional description for the MCP server",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "MCP server name",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "pager_duty": {
              "computed": true,
              "description": "PagerDuty service configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "PagerDuty OAuth authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "o_auth_client_credentials": {
                          "computed": true,
                          "description": "OAuth client credentials",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "client_id": {
                                "computed": true,
                                "description": "OAuth client ID",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_name": {
                                "computed": true,
                                "description": "User friendly OAuth client name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_secret": {
                                "computed": true,
                                "description": "OAuth client secret",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "exchange_parameters": {
                                "computed": true,
                                "description": "OAuth token exchange parameters",
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
                  "scopes": {
                    "computed": true,
                    "description": "PagerDuty scopes",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_now": {
              "computed": true,
              "description": "ServiceNow service configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "ServiceNow OAuth authorization configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "o_auth_client_credentials": {
                          "computed": true,
                          "description": "OAuth client credentials",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "client_id": {
                                "computed": true,
                                "description": "OAuth client ID",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_name": {
                                "computed": true,
                                "description": "User friendly OAuth client name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "client_secret": {
                                "computed": true,
                                "description": "OAuth client secret",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "exchange_parameters": {
                                "computed": true,
                                "description": "OAuth token exchange parameters",
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
                  "instance_url": {
                    "computed": true,
                    "description": "ServiceNow instance URL",
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
      "service_id": {
        "computed": true,
        "description": "The unique identifier of the service",
        "description_kind": "plain",
        "type": "string"
      },
      "service_type": {
        "computed": true,
        "description": "The type of service being registered",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DevOpsAgent::Service",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDevopsagentServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentService), &result)
	return &result
}
