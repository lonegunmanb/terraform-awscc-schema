package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAgentregistryRegistryRecord = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the registry record was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The identifier of the AWS account that created the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "descriptors": {
        "computed": true,
        "description": "The typed set of descriptors for a registry record. Exactly one descriptor field is populated based on the record type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "a2_a_agent_card": {
              "computed": true,
              "description": "The A2A agent card descriptor, populated when the record type is AGENT.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data": {
                    "computed": true,
                    "description": "Descriptor payload data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_schema_version": {
                    "computed": true,
                    "description": "Version of the descriptor type schema.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description": "The source configuration that defines where descriptor content is retrieved from.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from_url": {
                          "computed": true,
                          "description": "URL-based descriptor source configuration, with credential provider configurations for authenticated URL retrieval.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "credential_provider_configurations": {
                                "computed": true,
                                "description": "The credential providers used to authenticate when fetching descriptor content from the source URL.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "credential_provider": {
                                      "computed": true,
                                      "description": "The credential provider details. Specify exactly one member.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "iam_credential_provider": {
                                            "computed": true,
                                            "description": "IAM credential provider configuration.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "region": {
                                                  "computed": true,
                                                  "description": "The SigV4 signing region.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "role_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the IAM role.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "service": {
                                                  "computed": true,
                                                  "description": "The SigV4 signing service name.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "oauth_credential_provider": {
                                            "computed": true,
                                            "description": "OAuth credential provider configuration.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_parameters": {
                                                  "computed": true,
                                                  "description": "Additional custom parameters for the OAuth flow.",
                                                  "description_kind": "plain",
                                                  "type": [
                                                    "map",
                                                    "string"
                                                  ]
                                                },
                                                "grant_type": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "provider_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the OAuth credential provider.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "scopes": {
                                                  "computed": true,
                                                  "description": "OAuth scopes to request.",
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
                                    "credential_provider_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "url": {
                                "computed": true,
                                "description": "URL source for descriptor content.",
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
            "agent_skills_definition": {
              "computed": true,
              "description": "The agent skills definition descriptor, populated when the record type is SKILL.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "additional_data": {
                    "computed": true,
                    "description": "Additional data associated with an agent skills definition descriptor.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "skill_md": {
                          "computed": true,
                          "description": "Markdown-format descriptor containing an agent skills document.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data": {
                                "computed": true,
                                "description": "Descriptor payload data.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "data_schema_version": {
                                "computed": true,
                                "description": "Version of the descriptor type schema.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "source": {
                                "computed": true,
                                "description": "Source configuration for a SkillMd document. Unlike MCP/A2A sources, SkillMd does not support credential providers.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "from_url": {
                                      "computed": true,
                                      "description": "URL-based source for SkillMd content (sync is skipped; content is provided inline via Data).",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "url": {
                                            "computed": true,
                                            "description": "URL source for the SkillMd document.",
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "data": {
                    "computed": true,
                    "description": "Descriptor payload data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_schema_version": {
                    "computed": true,
                    "description": "Version of the descriptor type schema.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "agui": {
              "computed": true,
              "description": "The AG-UI (Agent-User Interaction) descriptor, populated for records detected from an AG-UI protocol source. This descriptor is source-only: its content is synchronized from the configured source URL rather than supplied inline.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source": {
                    "computed": true,
                    "description": "Source configuration for a source-only descriptor. Unlike mcpServer/a2aAgentCard sources, source-only descriptors do not support credential providers.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from_url": {
                          "computed": true,
                          "description": "URL-based source configuration for a source-only descriptor.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "url": {
                                "computed": true,
                                "description": "URL source for descriptor content.",
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
            "custom": {
              "computed": true,
              "description": "The custom descriptor, populated when the record type is CUSTOM.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data": {
                    "computed": true,
                    "description": "Descriptor payload data.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "http": {
              "computed": true,
              "description": "The HTTP descriptor, populated for records detected from an HTTP protocol source. This descriptor is source-only: its content is synchronized from the configured source URL rather than supplied inline.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source": {
                    "computed": true,
                    "description": "Source configuration for a source-only descriptor. Unlike mcpServer/a2aAgentCard sources, source-only descriptors do not support credential providers.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from_url": {
                          "computed": true,
                          "description": "URL-based source configuration for a source-only descriptor.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "url": {
                                "computed": true,
                                "description": "URL source for descriptor content.",
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
            "mcp_server": {
              "computed": true,
              "description": "The MCP server descriptor, populated when the record type is MCP.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "additional_data": {
                    "computed": true,
                    "description": "Additional data associated with an MCP server descriptor.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tools": {
                          "computed": true,
                          "description": "The MCP tools descriptor.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data": {
                                "computed": true,
                                "description": "Descriptor payload data.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "data_schema_version": {
                                "computed": true,
                                "description": "Version of the tools descriptor schema.",
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
                  "data": {
                    "computed": true,
                    "description": "Descriptor payload data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_schema_version": {
                    "computed": true,
                    "description": "Version of the descriptor type schema.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description": "The source configuration that defines where descriptor content is retrieved from.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from_url": {
                          "computed": true,
                          "description": "URL-based descriptor source configuration, with credential provider configurations for authenticated URL retrieval.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "credential_provider_configurations": {
                                "computed": true,
                                "description": "The credential providers used to authenticate when fetching descriptor content from the source URL.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "credential_provider": {
                                      "computed": true,
                                      "description": "The credential provider details. Specify exactly one member.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "iam_credential_provider": {
                                            "computed": true,
                                            "description": "IAM credential provider configuration.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "region": {
                                                  "computed": true,
                                                  "description": "The SigV4 signing region.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "role_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the IAM role.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "service": {
                                                  "computed": true,
                                                  "description": "The SigV4 signing service name.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "oauth_credential_provider": {
                                            "computed": true,
                                            "description": "OAuth credential provider configuration.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "custom_parameters": {
                                                  "computed": true,
                                                  "description": "Additional custom parameters for the OAuth flow.",
                                                  "description_kind": "plain",
                                                  "type": [
                                                    "map",
                                                    "string"
                                                  ]
                                                },
                                                "grant_type": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "provider_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the OAuth credential provider.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "scopes": {
                                                  "computed": true,
                                                  "description": "OAuth scopes to request.",
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
                                    "credential_provider_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "url": {
                                "computed": true,
                                "description": "URL source for descriptor content.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "display_name": {
        "computed": true,
        "description": "The human-readable display name of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "record_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "record_id": {
        "computed": true,
        "description": "The unique identifier of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "record_type": {
        "computed": true,
        "description": "The type of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "record_version": {
        "computed": true,
        "description": "The version of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the registry containing the record.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description": "The identifier of the registry in which to create the record. You can specify either the registry ID or the registry Amazon Resource Name (ARN). Use the ARN form to reference a registry shared from another account via AWS Resource Access Manager (RAM).",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The lifecycle status of the registry record.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the registry record.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the registry record was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AgentRegistry::RegistryRecord",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAgentregistryRegistryRecordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAgentregistryRegistryRecord), &result)
	return &result
}
