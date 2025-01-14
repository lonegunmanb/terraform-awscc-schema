package resource

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
        "optional": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_arn": {
        "computed": true,
        "description": "The ARN of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_configuration": {
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "description": "The type of a knowledge base.",
              "description_kind": "plain",
              "required": true,
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
                    "optional": true,
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
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "supplemental_data_storage_location_type": {
                                "computed": true,
                                "description": "Supplemental data storage location type.",
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
        "required": true
      },
      "knowledge_base_id": {
        "computed": true,
        "description": "The unique identifier of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the knowledge base.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the IAM role with permissions to invoke API operations on the knowledge base. The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_",
        "description_kind": "plain",
        "required": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "credentials_secret_arn": {
                    "computed": true,
                    "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Amazon Mongo database.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "Name of the database within MongoDB Atlas.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "MongoDB Atlas endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "endpoint_service_name": {
                    "computed": true,
                    "description": "MongoDB Atlas endpoint service name.",
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "vector_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "vector_index_name": {
                    "computed": true,
                    "description": "Name of a MongoDB Atlas index.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "vector_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "vector_index_name": {
                    "computed": true,
                    "description": "The name of the vector store.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  },
                  "credentials_secret_arn": {
                    "computed": true,
                    "description": "The ARN of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.",
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "namespace": {
                    "computed": true,
                    "description": "The namespace to be used to write new data to your database.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "The name of your Amazon RDS database.",
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "primary_key_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the ID for each entry.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "text_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the raw text from your data. The text is split according to the chunking strategy you choose.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "vector_field": {
                          "computed": true,
                          "description": "The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "The ARN of the vector store.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "The name of the table in the database.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description": "The storage type of a knowledge base.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "optional": true,
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
    "description": "Definition of AWS::Bedrock::KnowledgeBase Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockKnowledgeBaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockKnowledgeBase), &result)
	return &result
}
