package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendraDataSource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_document_enrichment_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "inline_configurations": {
              "computed": true,
              "description": "List of InlineCustomDocumentEnrichmentConfigurations",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_document_attribute_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_on_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "document_content_deletion": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "target": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_document_attribute_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "target_document_attribute_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "target_document_attribute_value_deletion": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "post_extraction_hook_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_document_attribute_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_on_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "pre_extraction_hook_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invocation_condition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "condition_document_attribute_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "condition_on_value": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "date_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "long_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "string_list_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "string_value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "operator": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "lambda_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "role_arn": {
              "computed": true,
              "description": "Role ARN",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_source_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "confluence_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attachment_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "attachment_field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "crawl_attachments": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "blog_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "blog_field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "page_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "page_field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "server_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "space_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crawl_archived_spaces": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "crawl_personal_spaces": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "exclude_spaces": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "include_spaces": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "space_field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_group_ids": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "subnet_ids": {
                          "computed": true,
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
            "database_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "acl_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "allowed_groups_column_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "column_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "change_detecting_columns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "document_data_column_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "document_id_column_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "document_title_column_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "connection_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_host": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "secret_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "database_engine_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sql_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "query_identifiers_enclosing_option": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "vpc_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_group_ids": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "subnet_ids": {
                          "computed": true,
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
            "google_drive_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "exclude_mime_types": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclude_shared_drives": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclude_user_accounts": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "field_mappings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_source_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "date_field_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "one_drive_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "disable_local_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "field_mappings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_source_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "date_field_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "one_drive_users": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "one_drive_user_list": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "one_drive_user_s3_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "key": {
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
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tenant_domain": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "s3_configuration": {
              "computed": true,
              "description": "S3 data source configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_control_list_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "documents_metadata_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "inclusion_prefixes": {
                    "computed": true,
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
            "salesforce_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "chatter_feed_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "document_data_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "document_title_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "include_filter_types": {
                          "computed": true,
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
                  "crawl_attachments": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "exclude_attachment_file_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "include_attachment_file_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "knowledge_article_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "custom_knowledge_article_type_configurations": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "document_data_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "document_title_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "field_mappings": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_source_field_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "date_field_format": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "index_field_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "included_states": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "standard_knowledge_article_type_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "document_data_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "document_title_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "field_mappings": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "data_source_field_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "date_field_format": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "index_field_name": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
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
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "server_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "standard_object_attachment_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "document_title_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "standard_object_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "document_data_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "document_title_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_now_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "knowledge_article_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crawl_attachments": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "document_data_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "document_title_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "exclude_attachment_file_patterns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "filter_query": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "include_attachment_file_patterns": {
                          "computed": true,
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
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "service_catalog_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crawl_attachments": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "document_data_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "document_title_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "exclude_attachment_file_patterns": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "field_mappings": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_source_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "date_field_format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "index_field_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "include_attachment_file_patterns": {
                          "computed": true,
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
                  "service_now_build_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "share_point_configuration": {
              "computed": true,
              "description": "SharePoint configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawl_attachments": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "disable_local_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "document_title_field_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "field_mappings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_source_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "date_field_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "share_point_version": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_certificate_s3_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "urls": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "use_change_log": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "vpc_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "security_group_ids": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "subnet_ids": {
                          "computed": true,
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
            "web_crawler_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "basic_authentication": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "credentials": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "host": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "crawl_depth": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_content_size_per_page_in_mega_bytes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_links_per_page": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "max_urls_per_minute_crawl_rate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "proxy_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "credentials": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "host": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "url_exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "url_inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "urls": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "seed_url_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "seed_urls": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "web_crawler_mode": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "site_maps_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "site_maps": {
                                "computed": true,
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
                  }
                },
                "nesting_mode": "single"
              }
            },
            "work_docs_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawl_comments": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "exclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "field_mappings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_source_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "date_field_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "index_field_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "inclusion_patterns": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "organization_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "use_change_log": {
                    "computed": true,
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
      "data_source_id": {
        "computed": true,
        "description": "ID of data source",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of data source",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_id": {
        "computed": true,
        "description": "ID of Index",
        "description_kind": "plain",
        "type": "string"
      },
      "language_code": {
        "computed": true,
        "description": "The code for a language.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of data source",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Role ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "Schedule",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for labeling the data source",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "Data source type",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Kendra::DataSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKendraDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraDataSource), &result)
	return &result
}
