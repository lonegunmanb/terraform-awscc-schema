package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentAssociation = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "description": "The unique identifier of the AgentSpace",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "association_id": {
        "computed": true,
        "description": "The unique identifier of the association",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "description": "The configuration that directs how AgentSpace interacts with the given service",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws": {
              "computed": true,
              "description": "AWS association for 'monitor' account",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description": "AWS Account Id corresponding to provided resources",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "account_type": {
                    "computed": true,
                    "description": "Account Type 'monitor' for DevOpsAgent monitoring",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "assumable_role_arn": {
                    "computed": true,
                    "description": "Role ARN to be assumed by DevOpsAgent to operate on behalf of customer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resources": {
                    "computed": true,
                    "description": "List of AWS resources",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resource_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the resource",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_metadata": {
                          "computed": true,
                          "description": "Additional metadata for the resource",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_type": {
                          "computed": true,
                          "description": "Resource type",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "tags": {
                    "computed": true,
                    "description": "List of AWS tags as key-value pairs",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "dynatrace": {
              "computed": true,
              "description": "Dynatrace monitoring configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "env_id": {
                    "computed": true,
                    "description": "Dynatrace environment id",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resources": {
                    "computed": true,
                    "description": "List of Dynatrace resources to monitor",
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
            "event_channel": {
              "computed": true,
              "description": "EventChannelconfiguration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "git_hub": {
              "computed": true,
              "description": "GitHub repository integration configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "owner": {
                    "computed": true,
                    "description": "Repository owner",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "owner_type": {
                    "computed": true,
                    "description": "Type of repository owner",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "repo_id": {
                    "computed": true,
                    "description": "Associated Github repo ID",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "repo_name": {
                    "computed": true,
                    "description": "Associated Github repo name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "git_lab": {
              "computed": true,
              "description": "GitLab project integration configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "instance_identifier": {
                    "computed": true,
                    "description": "GitLab instance identifier",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_id": {
                    "computed": true,
                    "description": "GitLab numeric project ID",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "project_path": {
                    "computed": true,
                    "description": "Full GitLab project path (e.g., namespace/project-name)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "mcp_server": {
              "computed": true,
              "description": "MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "The description of the MCP server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the MCP server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tools": {
                    "computed": true,
                    "description": "List of MCP tools that can be used with the association",
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
            "mcp_server_datadog": {
              "computed": true,
              "description": "Datadog MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "The description of the MCP server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the MCP server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "mcp_server_new_relic": {
              "computed": true,
              "description": "NewRelic MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description": "New Relic Account ID",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL (e.g., https://mcp.newrelic.com/mcp/)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "mcp_server_splunk": {
              "computed": true,
              "description": "Splunk MCP server configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "description": {
                    "computed": true,
                    "description": "The description of the MCP server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MCP server endpoint URL",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the MCP server",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "service_now": {
              "computed": true,
              "description": "ServiceNow integration configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_webhook_updates": {
                    "computed": true,
                    "description": "When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "instance_id": {
                    "computed": true,
                    "description": "ServiceNow instance ID",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "slack": {
              "computed": true,
              "description": "Slack workspace integration configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "transmission_target": {
                    "computed": true,
                    "description": "Transmission targets for agent notifications",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "incident_response_target": {
                          "computed": true,
                          "description": "Destination for IncidentResponse agent.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "channel_id": {
                                "computed": true,
                                "description": "Slack channel ID",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "channel_name": {
                                "computed": true,
                                "description": "Slack channel name",
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
                  "workspace_id": {
                    "computed": true,
                    "description": "Associated Slack workspace ID",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "workspace_name": {
                    "computed": true,
                    "description": "Associated Slack workspace name",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "source_aws": {
              "computed": true,
              "description": "AWS association for 'source' account",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description": "AWS Account Id corresponding to provided resources",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "account_type": {
                    "computed": true,
                    "description": "Account Type 'source' for DevOpsAgent monitoring",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "assumable_role_arn": {
                    "computed": true,
                    "description": "Role ARN to be assumed by DevOpsAgent to operate on behalf of customer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resources": {
                    "computed": true,
                    "description": "List of AWS resources",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resource_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the resource",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_metadata": {
                          "computed": true,
                          "description": "Additional metadata for the resource",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_type": {
                          "computed": true,
                          "description": "Resource type",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "tags": {
                    "computed": true,
                    "description": "List of AWS tags as key-value pairs",
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
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the association was created",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "description": "The identifier for the associated service",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the association was last updated",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::DevOpsAgent::Association defining how the AgentSpace interacts with external services like GitHub, Slack, AWS accounts, and others.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsagentAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentAssociation), &result)
	return &result
}
