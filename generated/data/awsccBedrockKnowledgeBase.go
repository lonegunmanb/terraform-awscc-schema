package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockKnowledgeBase = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time at which the knowledge base was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_reasons": {
        "computed": true,
        "description": "A list of reasons that the API operation on the knowledge base failed.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "knowledge_base_arn": {
        "computed": true,
        "description": "The ARN of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_configuration": {
        "computed": true,
        "description": "Contains details about the embeddings model used for the knowledge base.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kendra_knowledge_base_configuration": {
              "computed": true,
              "description": "Configurations for a Kendra knowledge base",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kendra_index_arn": {
                    "computed": true,
                    "description": "Arn of a Kendra index",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sql_knowledge_base_configuration": {
              "computed": true,
              "description": "Configurations for a SQL knowledge base",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "redshift_configuration": {
                    "computed": true,
                    "description": "Configurations for a Redshift knowledge base",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "query_engine_configuration": {
                          "computed": true,
                          "description": "Configurations for Redshift query engine",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "provisioned_configuration": {
                                "computed": true,
                                "description": "Configurations for provisioned Redshift query engine",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "auth_configuration": {
                                      "computed": true,
                                      "description": "Configurations for Redshift query engine provisioned auth setup",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "database_user": {
                                            "computed": true,
                                            "description": "Redshift database user",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "type": {
                                            "computed": true,
                                            "description": "Provisioned Redshift auth type",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "username_password_secret_arn": {
                                            "computed": true,
                                            "description": "Arn of a SecretsManager Secret",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "cluster_identifier": {
                                      "computed": true,
                                      "description": "Redshift cluster identifier",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "serverless_configuration": {
                                "computed": true,
                                "description": "Configurations for serverless Redshift query engine",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "auth_configuration": {
                                      "computed": true,
                                      "description": "Configurations for Redshift query engine serverless auth setup",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "type": {
                                            "computed": true,
                                            "description": "Serverless Redshift auth type",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "username_password_secret_arn": {
                                            "computed": true,
                                            "description": "Arn of a SecretsManager Secret",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "workgroup_arn": {
                                      "computed": true,
                                      "description": "Workgroup arn",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "type": {
                                "computed": true,
                                "description": "Redshift query engine type",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "query_generation_configuration": {
                          "computed": true,
                          "description": "Configurations for generating Redshift engine queries",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "execution_timeout_seconds": {
                                "computed": true,
                                "description": "Max query execution timeout",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "generation_context": {
                                "computed": true,
                                "description": "Context used to improve query generation",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "curated_queries": {
                                      "computed": true,
                                      "description": "List of example queries and results",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "natural_language": {
                                            "computed": true,
                                            "description": "Question for the curated query",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "sql": {
                                            "computed": true,
                                            "description": "Answer for the curated query",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    },
                                    "tables": {
                                      "computed": true,
                                      "description": "List of tables used for Redshift query generation context",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "columns": {
                                            "computed": true,
                                            "description": "List of Redshift query generation columns",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "description": {
                                                  "computed": true,
                                                  "description": "Description for the attached entity",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "inclusion": {
                                                  "computed": true,
                                                  "description": "Include or Exclude status for an entity",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "name": {
                                                  "computed": true,
                                                  "description": "Query generation column name",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "description": {
                                            "computed": true,
                                            "description": "Description for the attached entity",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "inclusion": {
                                            "computed": true,
                                            "description": "Include or Exclude status for an entity",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "name": {
                                            "computed": true,
                                            "description": "Query generation table name. Must follow three-part notation",
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
                        "storage_configurations": {
                          "computed": true,
                          "description": "List of configurations for available Redshift query engine storage types",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aws_data_catalog_configuration": {
                                "computed": true,
                                "description": "Configurations for Redshift query engine AWS Data Catalog backed storage",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "table_names": {
                                      "computed": true,
                                      "description": "List of table names in AWS Data Catalog. Must follow two part notation",
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
                              "redshift_configuration": {
                                "computed": true,
                                "description": "Configurations for Redshift query engine Redshift backed storage",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "database_name": {
                                      "computed": true,
                                      "description": "Redshift database name",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "type": {
                                "computed": true,
                                "description": "Redshift query engine storage type",
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
                  "type": {
                    "computed": true,
                    "description": "SQL query engine type",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description": "The type of a knowledge base.",
              "description_kind": "plain",
              "type": "string"
            },
            "vector_knowledge_base_configuration": {
              "computed": true,
              "description": "Contains details about the model used to create vector embeddings for the knowledge base.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "embedding_model_arn": {
                    "computed": true,
                    "description": "The ARN of the model used to create vector embeddings for the knowledge base.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "embedding_model_configuration": {
                    "computed": true,
                    "description": "The embeddings model configuration details for the vector model used in Knowledge Base.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bedrock_embedding_model_configuration": {
                          "computed": true,
                          "description": "The vector configuration details for the Bedrock embeddings model.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "dimensions": {
                                "computed": true,
                                "description": "The dimensions details for the vector configuration used on the Bedrock embeddings model.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "embedding_data_type": {
                                "computed": true,
                                "description": "The data type for the vectors when using a model to convert text into vector embeddings.",
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
                  "supplemental_data_storage_configuration": {
                    "computed": true,
                    "description": "Configurations for supplemental data storage.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "supplemental_data_storage_locations": {
                          "computed": true,
                          "description": "List of supplemental data storage locations.",
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
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "supplemental_data_storage_location_type": {
                                "computed": true,
                                "description": "Supplemental data storage location type.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "knowledge_base_id": {
        "computed": true,
        "description": "The unique identifier of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role with permissions to invoke API operations on the knowledge base. The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of a knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_configuration": {
        "computed": true,
        "description": "The vector store service in which the knowledge base is stored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mongo_db_atlas_configuration": {
              "computed": true,
              "description": "Contains the storage configuration of the knowledge base in MongoDb Atlas Cloud.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "collection_name": {
                    "computed": true,
                    "description": "Name of the collection within MongoDB Atlas.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "credentials_secret_arn": {
                    "computed": true,
                    "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon Mongo database.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "Name of the database within MongoDB Atlas.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MongoDB Atlas endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint_service_name": {
                    "computed": true,
                    "description": "MongoDB Atlas endpoint service name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "field_mapping": {
                    "computed": true,
                    "description": "Contains the names of the fields to which to map information about the vector store.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metadata_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "vector_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "vector_index_name": {
                    "computed": true,
                    "description": "Name of a MongoDB Atlas index.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "neptune_analytics_configuration": {
              "computed": true,
              "description": "Contains the configurations to use Neptune Analytics as Vector Store.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "field_mapping": {
                    "computed": true,
                    "description": "A mapping of Bedrock Knowledge Base fields to Neptune Analytics fields.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metadata_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "graph_arn": {
                    "computed": true,
                    "description": "ARN for Neptune Analytics graph database.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "opensearch_serverless_configuration": {
              "computed": true,
              "description": "Contains the storage configuration of the knowledge base in Amazon OpenSearch Service.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "collection_arn": {
                    "computed": true,
                    "description": "The ARN of the OpenSearch Service vector store.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "field_mapping": {
                    "computed": true,
                    "description": "A mapping of Bedrock Knowledge Base fields to OpenSearch Serverless field names",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metadata_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "vector_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "vector_index_name": {
                    "computed": true,
                    "description": "The name of the vector store.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "pinecone_configuration": {
              "computed": true,
              "description": "Contains the storage configuration of the knowledge base in Pinecone.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_string": {
                    "computed": true,
                    "description": "The endpoint URL for your index management page.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "credentials_secret_arn": {
                    "computed": true,
                    "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "field_mapping": {
                    "computed": true,
                    "description": "Contains the names of the fields to which to map information about the vector store.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metadata_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "namespace": {
                    "computed": true,
                    "description": "The namespace to be used to write new data to your database.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "rds_configuration": {
              "computed": true,
              "description": "Contains details about the storage configuration of the knowledge base in Amazon RDS. For more information, see Create a vector index in Amazon RDS.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credentials_secret_arn": {
                    "computed": true,
                    "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "The name of your Amazon RDS database.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "field_mapping": {
                    "computed": true,
                    "description": "Contains the names of the fields to which to map information about the vector store.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metadata_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores metadata about the vector store.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "primary_key_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the ID for each entry.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "vector_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "The ARN of the vector store.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "The name of the table in the database.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description": "The storage type of a knowledge base.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "The time at which the knowledge base was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::KnowledgeBase",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockKnowledgeBaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockKnowledgeBase), &result)
	return &result
}
