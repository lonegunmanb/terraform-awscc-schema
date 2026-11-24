package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAgentregistryRegistryRecord = `{
  "block": {
    "attributes": {
      "artifact": {
        "computed": true,
        "description": "The skill package that the registry stores for a skill record that synchronizes a skill package. The package's download URL isn't returned.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sha_256": {
              "computed": true,
              "description": "The SHA-256 checksum of the skill package, as 64 lowercase hexadecimal characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
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
      "custom_metadata": {
        "computed": true,
        "description": "The custom metadata to attach to the registry record. Each key must match a property defined in the registry's custom metadata schema. Values can be strings (maximum 128 characters) or native JSON booleans (true or false).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the registry record.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "descriptors": {
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
                    "optional": true,
                    "type": "string"
                  },
                  "data_schema_version": {
                    "computed": true,
                    "description": "Version of the descriptor type schema.",
                    "description_kind": "plain",
                    "optional": true,
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
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "role_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the IAM role.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "service": {
                                                  "computed": true,
                                                  "description": "The SigV4 signing service name.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
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
                                                  "optional": true,
                                                  "type": [
                                                    "map",
                                                    "string"
                                                  ]
                                                },
                                                "grant_type": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "provider_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the OAuth credential provider.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "scopes": {
                                                  "computed": true,
                                                  "description": "OAuth scopes to request.",
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
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "credential_provider_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "url": {
                                "computed": true,
                                "description": "URL source for descriptor content.",
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
                                "optional": true,
                                "type": "string"
                              },
                              "data_schema_version": {
                                "computed": true,
                                "description": "Version of the descriptor type schema.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "source": {
                                "computed": true,
                                "description": "The source that the registry synchronizes agent skills content from: a SKILL.md document in SkillMd.Source, or a skill package in SkillPackage.Source.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "from_url": {
                                      "computed": true,
                                      "description": "URL-based source of agent skills content. With an s3:// URL, the registry reads the Amazon S3 object when the record is created, and again whenever the URL, the credential provider or the bucket owner changes. Each read replaces the SKILL.md Data and sets the record's Name and Description from the name and description in the SKILL.md frontmatter, so set Name and Description to the same values and don't set SkillMd Data. An https:// URL is a reference URL that the registry doesn't synchronize from: where Amazon S3 skill synchronization is available, the registry accepts an https:// URL only on a record that already stores it, with the content provided inline in SkillMd Data.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "credential_provider_configurations": {
                                            "computed": true,
                                            "description": "The credential provider that the registry uses to read the source. For an Amazon S3 source, specify an IAM role in your account that the agent-registry.amazonaws.com service principal can assume and that can read the object. Omit it to read a public object, and set SourceConfiguration.S3.BucketOwnerAccountId instead.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "credential_provider": {
                                                  "computed": true,
                                                  "description": "The credential provider details.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "iam_credential_provider": {
                                                        "computed": true,
                                                        "description": "The IAM role that the registry assumes to read the source. For an Amazon S3 source, omit Service or set it to s3, and omit Region or set it to the registry's Region.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "region": {
                                                              "computed": true,
                                                              "description": "The SigV4 signing region.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "role_arn": {
                                                              "computed": true,
                                                              "description": "The ARN of the IAM role.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "service": {
                                                              "computed": true,
                                                              "description": "The SigV4 signing service name.",
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
                                                "credential_provider_type": {
                                                  "computed": true,
                                                  "description": "The type of the credential provider.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "url": {
                                            "computed": true,
                                            "description": "The URL of the source. For synchronization, specify an s3://\u003cbucket\u003e/\u003ckey\u003e URL that names one object in a bucket in the same AWS Region as the registry: a SKILL.md object for SkillMd, or a skill package ZIP object for SkillPackage.",
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
                        "skill_package": {
                          "computed": true,
                          "description": "The source that the registry synchronizes the full skill package from. You can't set it together with SkillMd.Source or SkillMd Data.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "source": {
                                "computed": true,
                                "description": "The source of the skill package, as an s3:// URL of a ZIP object.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "from_url": {
                                      "computed": true,
                                      "description": "URL-based source of agent skills content. With an s3:// URL, the registry reads the Amazon S3 object when the record is created, and again whenever the URL, the credential provider or the bucket owner changes. Each read replaces the SKILL.md Data and sets the record's Name and Description from the name and description in the SKILL.md frontmatter, so set Name and Description to the same values and don't set SkillMd Data. An https:// URL is a reference URL that the registry doesn't synchronize from: where Amazon S3 skill synchronization is available, the registry accepts an https:// URL only on a record that already stores it, with the content provided inline in SkillMd Data.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "credential_provider_configurations": {
                                            "computed": true,
                                            "description": "The credential provider that the registry uses to read the source. For an Amazon S3 source, specify an IAM role in your account that the agent-registry.amazonaws.com service principal can assume and that can read the object. Omit it to read a public object, and set SourceConfiguration.S3.BucketOwnerAccountId instead.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "credential_provider": {
                                                  "computed": true,
                                                  "description": "The credential provider details.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "iam_credential_provider": {
                                                        "computed": true,
                                                        "description": "The IAM role that the registry assumes to read the source. For an Amazon S3 source, omit Service or set it to s3, and omit Region or set it to the registry's Region.",
                                                        "description_kind": "plain",
                                                        "nested_type": {
                                                          "attributes": {
                                                            "region": {
                                                              "computed": true,
                                                              "description": "The SigV4 signing region.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "role_arn": {
                                                              "computed": true,
                                                              "description": "The ARN of the IAM role.",
                                                              "description_kind": "plain",
                                                              "optional": true,
                                                              "type": "string"
                                                            },
                                                            "service": {
                                                              "computed": true,
                                                              "description": "The SigV4 signing service name.",
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
                                                "credential_provider_type": {
                                                  "computed": true,
                                                  "description": "The type of the credential provider.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "url": {
                                            "computed": true,
                                            "description": "The URL of the source. For synchronization, specify an s3://\u003cbucket\u003e/\u003ckey\u003e URL that names one object in a bucket in the same AWS Region as the registry: a SKILL.md object for SkillMd, or a skill package ZIP object for SkillPackage.",
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
                        "source_configuration": {
                          "computed": true,
                          "description": "The provider-specific settings of the record's skill source.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3": {
                                "computed": true,
                                "description": "The settings of an Amazon S3 source.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "bucket_owner_account_id": {
                                      "computed": true,
                                      "description": "The ID of the AWS account that owns the source bucket. The registry checks that the bucket belongs to this account each time it reads the source. If you omit it for a source that uses an IAM role, the registry expects the bucket in the role's account. A source without a credential provider, which the registry reads as a public object, requires it.",
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
                  "data": {
                    "computed": true,
                    "description": "Descriptor payload data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_schema_version": {
                    "computed": true,
                    "description": "Version of the descriptor type schema.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                                "optional": true,
                                "type": "string"
                              },
                              "data_schema_version": {
                                "computed": true,
                                "description": "Version of the tools descriptor schema.",
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
                  "data": {
                    "computed": true,
                    "description": "Descriptor payload data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_schema_version": {
                    "computed": true,
                    "description": "Version of the descriptor type schema.",
                    "description_kind": "plain",
                    "optional": true,
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
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "role_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the IAM role.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "service": {
                                                  "computed": true,
                                                  "description": "The SigV4 signing service name.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
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
                                                  "optional": true,
                                                  "type": [
                                                    "map",
                                                    "string"
                                                  ]
                                                },
                                                "grant_type": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "provider_arn": {
                                                  "computed": true,
                                                  "description": "The ARN of the OAuth credential provider.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "scopes": {
                                                  "computed": true,
                                                  "description": "OAuth scopes to request.",
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
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "credential_provider_type": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "url": {
                                "computed": true,
                                "description": "URL source for descriptor content.",
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
      "display_name": {
        "computed": true,
        "description": "The human-readable display name of the registry record.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the registry record.",
        "description_kind": "plain",
        "required": true,
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
        "description": "The type of the registry record.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "record_version": {
        "computed": true,
        "description": "The version of the registry record.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The lifecycle status of the registry record. CREATE_PENDING_AUTH and UPDATE_PENDING_AUTH mean that the record waits for a user to authorize access to a GitHub source. CloudFormation doesn't support GitHub sources, so a create or update that reaches one of these states fails.",
        "description_kind": "plain",
        "type": "string"
      },
      "synchronized_source": {
        "computed": true,
        "description": "The source that the registry last synchronized the record's skill content from, and the version that it read. Absent until the first successful synchronization.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3": {
              "computed": true,
              "description": "The Amazon S3 source that the registry read.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "e_tag": {
                    "computed": true,
                    "description": "The entity tag (ETag) of the object that the registry read, without quotes, when the object has no version ID. Not recorded for a SKILL.md source on a registry that is encrypted with a customer managed key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "synchronized_at": {
                    "computed": true,
                    "description": "The timestamp when the registry read the source.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version_id": {
                    "computed": true,
                    "description": "The version ID of the object that the registry read. Absent when the bucket isn't versioned.",
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
        "description": "Tags to assign to the registry record.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the registry record was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::AgentRegistry::RegistryRecord Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAgentregistryRegistryRecordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAgentregistryRegistryRecord), &result)
	return &result
}
