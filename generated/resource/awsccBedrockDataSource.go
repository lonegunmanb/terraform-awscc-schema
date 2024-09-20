package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockDataSource = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time at which the data source was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_deletion_policy": {
        "computed": true,
        "description": "The deletion policy for the data source.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_configuration": {
        "description": "Specifies a raw data source location to ingest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "confluence_configuration": {
              "computed": true,
              "description": "The configuration information to connect to Confluence as your data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawler_configuration": {
                    "computed": true,
                    "description": "The configuration of the Confluence content. For example, configuring specific types of Confluence content.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "filter_configuration": {
                          "computed": true,
                          "description": "The type of filtering that you want to apply to certain objects or content of the data source. For example, the PATTERN type is regular expression patterns you can apply to filter your content.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "pattern_object_filter": {
                                "computed": true,
                                "description": "The configuration of specific filters applied to your data source content. You can filter out or include certain content.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "filters": {
                                      "computed": true,
                                      "description": "Contains information",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "exclusion_filters": {
                                            "computed": true,
                                            "description": "A set of regular expression filter patterns for a type of object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "inclusion_filters": {
                                            "computed": true,
                                            "description": "A set of regular expression filter patterns for a type of object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "object_type": {
                                            "computed": true,
                                            "description": "The supported object type or content type of the data source.",
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
                              "type": {
                                "computed": true,
                                "description": "The crawl filter type.",
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
                  "source_configuration": {
                    "computed": true,
                    "description": "The endpoint information to connect to your Confluence data source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auth_type": {
                          "computed": true,
                          "description": "The supported authentication type to authenticate and connect to your Confluence instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "credentials_secret_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Confluence instance URL. For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see Confluence connection configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_type": {
                          "computed": true,
                          "description": "The supported host type, whether online/cloud or server/on-premises.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_url": {
                          "computed": true,
                          "description": "The Confluence host URL or instance URL.",
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
            "s3_configuration": {
              "computed": true,
              "description": "The configuration information to connect to Amazon S3 as your data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_arn": {
                    "computed": true,
                    "description": "The ARN of the bucket that contains the data source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_owner_account_id": {
                    "computed": true,
                    "description": "The account ID for the owner of the S3 bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "inclusion_prefixes": {
                    "computed": true,
                    "description": "A list of S3 prefixes that define the object containing the data sources.",
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
            "salesforce_configuration": {
              "computed": true,
              "description": "The configuration information to connect to Salesforce as your data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawler_configuration": {
                    "computed": true,
                    "description": "The configuration of filtering the Salesforce content. For example, configuring regular expression patterns to include or exclude certain content.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "filter_configuration": {
                          "computed": true,
                          "description": "The type of filtering that you want to apply to certain objects or content of the data source. For example, the PATTERN type is regular expression patterns you can apply to filter your content.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "pattern_object_filter": {
                                "computed": true,
                                "description": "The configuration of specific filters applied to your data source content. You can filter out or include certain content.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "filters": {
                                      "computed": true,
                                      "description": "Contains information",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "exclusion_filters": {
                                            "computed": true,
                                            "description": "A set of regular expression filter patterns for a type of object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "inclusion_filters": {
                                            "computed": true,
                                            "description": "A set of regular expression filter patterns for a type of object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "object_type": {
                                            "computed": true,
                                            "description": "The supported object type or content type of the data source.",
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
                              "type": {
                                "computed": true,
                                "description": "The crawl filter type.",
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
                  "source_configuration": {
                    "computed": true,
                    "description": "The endpoint information to connect to your Salesforce data source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auth_type": {
                          "computed": true,
                          "description": "The supported authentication type to authenticate and connect to your Salesforce instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "credentials_secret_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your Salesforce instance URL. For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see Salesforce connection configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_url": {
                          "computed": true,
                          "description": "The Salesforce host URL or instance URL.",
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
            "share_point_configuration": {
              "computed": true,
              "description": "The configuration information to connect to SharePoint as your data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawler_configuration": {
                    "computed": true,
                    "description": "The configuration of the SharePoint content. For example, configuring specific types of SharePoint content.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "filter_configuration": {
                          "computed": true,
                          "description": "The type of filtering that you want to apply to certain objects or content of the data source. For example, the PATTERN type is regular expression patterns you can apply to filter your content.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "pattern_object_filter": {
                                "computed": true,
                                "description": "The configuration of specific filters applied to your data source content. You can filter out or include certain content.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "filters": {
                                      "computed": true,
                                      "description": "Contains information",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "exclusion_filters": {
                                            "computed": true,
                                            "description": "A set of regular expression filter patterns for a type of object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "inclusion_filters": {
                                            "computed": true,
                                            "description": "A set of regular expression filter patterns for a type of object.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "string"
                                            ]
                                          },
                                          "object_type": {
                                            "computed": true,
                                            "description": "The supported object type or content type of the data source.",
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
                              "type": {
                                "computed": true,
                                "description": "The crawl filter type.",
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
                  "source_configuration": {
                    "computed": true,
                    "description": "The endpoint information to connect to your SharePoint data source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auth_type": {
                          "computed": true,
                          "description": "The supported authentication type to authenticate and connect to your SharePoint site/sites.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "credentials_secret_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your SharePoint site/sites. For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see SharePoint connection configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "domain": {
                          "computed": true,
                          "description": "The domain of your SharePoint instance or site URL/URLs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_type": {
                          "computed": true,
                          "description": "The supported host type, whether online/cloud or server/on-premises.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "site_urls": {
                          "computed": true,
                          "description": "A list of one or more SharePoint site URLs.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "tenant_id": {
                          "computed": true,
                          "description": "The identifier of your Microsoft 365 tenant.",
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
            "type": {
              "description": "The type of the data source location.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "web_configuration": {
              "computed": true,
              "description": "Configures a web data source location.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "crawler_configuration": {
                    "computed": true,
                    "description": "Configuration for the web crawler.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "crawler_limits": {
                          "computed": true,
                          "description": "Limit settings for the web crawler.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rate_limit": {
                                "computed": true,
                                "description": "Rate of web URLs retrieved per minute.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "exclusion_filters": {
                          "computed": true,
                          "description": "A set of regular expression filter patterns for a type of object.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "inclusion_filters": {
                          "computed": true,
                          "description": "A set of regular expression filter patterns for a type of object.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "scope": {
                          "computed": true,
                          "description": "The scope that a web crawl job will be restricted to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "source_configuration": {
                    "computed": true,
                    "description": "A web source configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "url_configuration": {
                          "computed": true,
                          "description": "A url configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "seed_urls": {
                                "computed": true,
                                "description": "A list of web urls.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "url": {
                                      "computed": true,
                                      "description": "A web url.",
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
      "data_source_id": {
        "computed": true,
        "description": "Identifier for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_status": {
        "computed": true,
        "description": "The status of a data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reasons": {
        "computed": true,
        "description": "The details of the failure reasons related to the data source.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_id": {
        "description": "The unique identifier of the knowledge base to which to add the data source.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the data source.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_side_encryption_configuration": {
        "computed": true,
        "description": "Contains details about the server-side encryption for the data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of the AWS KMS key used to encrypt the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time at which the knowledge base was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "vector_ingestion_configuration": {
        "computed": true,
        "description": "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "chunking_configuration": {
              "computed": true,
              "description": "Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "chunking_strategy": {
                    "computed": true,
                    "description": "Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for NONE, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fixed_size_chunking_configuration": {
                    "computed": true,
                    "description": "Configurations for when you choose fixed-size chunking. If you set the chunkingStrategy as NONE, exclude this field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_tokens": {
                          "computed": true,
                          "description": "The maximum number of tokens to include in a chunk.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "overlap_percentage": {
                          "computed": true,
                          "description": "The percentage of overlap between adjacent chunks of a data source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "hierarchical_chunking_configuration": {
                    "computed": true,
                    "description": "Configurations for when you choose hierarchical chunking. If you set the chunkingStrategy as NONE, exclude this field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "level_configurations": {
                          "computed": true,
                          "description": "Token settings for each layer.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_tokens": {
                                "computed": true,
                                "description": "The maximum number of tokens that a chunk can contain in this layer.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "overlap_tokens": {
                          "computed": true,
                          "description": "The number of tokens to repeat across chunks in the same layer.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "semantic_chunking_configuration": {
                    "computed": true,
                    "description": "Configurations for when you choose semantic chunking. If you set the chunkingStrategy as NONE, exclude this field.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "breakpoint_percentile_threshold": {
                          "computed": true,
                          "description": "The dissimilarity threshold for splitting chunks.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "buffer_size": {
                          "computed": true,
                          "description": "The buffer size.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_tokens": {
                          "computed": true,
                          "description": "The maximum number of tokens that a chunk can contain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "custom_transformation_configuration": {
              "computed": true,
              "description": "Settings for customizing steps in the data source content ingestion pipeline.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "intermediate_storage": {
                    "computed": true,
                    "description": "A location for storing content from data sources temporarily as it is processed by custom components in the ingestion pipeline.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_location": {
                          "computed": true,
                          "description": "An Amazon S3 location.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "uri": {
                                "computed": true,
                                "description": "The location's URI",
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
                  "transformations": {
                    "computed": true,
                    "description": "A list of Lambda functions that process documents.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "step_to_apply": {
                          "computed": true,
                          "description": "When the service applies the transformation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "transformation_function": {
                          "computed": true,
                          "description": "A Lambda function that processes documents.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "transformation_lambda_configuration": {
                                "computed": true,
                                "description": "A Lambda function that processes documents.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "lambda_arn": {
                                      "computed": true,
                                      "description": "The function's ARN identifier.",
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "parsing_configuration": {
              "computed": true,
              "description": "Settings for parsing document contents",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bedrock_foundation_model_configuration": {
                    "computed": true,
                    "description": "Settings for a foundation model used to parse documents for a data source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "model_arn": {
                          "computed": true,
                          "description": "The model's ARN.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "parsing_prompt": {
                          "computed": true,
                          "description": "Instructions for interpreting the contents of a document.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "parsing_prompt_text": {
                                "computed": true,
                                "description": "Instructions for interpreting the contents of a document.",
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
                  "parsing_strategy": {
                    "computed": true,
                    "description": "The parsing strategy for the data source.",
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
    "description": "Definition of AWS::Bedrock::DataSource Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockDataSource), &result)
	return &result
}
