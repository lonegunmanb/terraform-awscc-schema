package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightDataSource = `{
  "block": {
    "attributes": {
      "alternate_data_source_parameters": {
        "computed": true,
        "description": "\u003cp\u003eA set of alternate data source parameters that you want to share for the credentials\n            stored with this data source. The credentials are applied in tandem with the data source\n            parameters when you copy a data source by using a create or update request. The API\n            operation compares the \u003ccode\u003eDataSourceParameters\u003c/code\u003e structure that's in the request\n            with the structures in the \u003ccode\u003eAlternateDataSourceParameters\u003c/code\u003e allow list. If the\n            structures are an exact match, the request is allowed to use the credentials from this\n            existing data source. If the \u003ccode\u003eAlternateDataSourceParameters\u003c/code\u003e list is null,\n            the \u003ccode\u003eCredentials\u003c/code\u003e originally used with this \u003ccode\u003eDataSourceParameters\u003c/code\u003e\n            are automatically allowed.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_elasticsearch_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for OpenSearch.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "description": "\u003cp\u003eThe OpenSearch domain.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "amazon_open_search_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for OpenSearch.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "description": "\u003cp\u003eThe OpenSearch domain.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "athena_parameters": {
              "computed": true,
              "description": "\u003cp\u003eParameters for Amazon Athena.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to override an account-wide role for a specific Athena data source. For example, say an account administrator has turned off all Athena access with an account-wide role. The administrator can then use \u003ccode\u003eRoleArn\u003c/code\u003e to bypass the account-wide role and allow Athena access for the single Athena data source that is specified in the structure, even if the account-wide role forbidding Athena access is still active.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "work_group": {
                    "computed": true,
                    "description": "\u003cp\u003eThe workgroup that Amazon Athena uses.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "aurora_parameters": {
              "computed": true,
              "description": "\u003cp\u003eParameters for Amazon Aurora.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "aurora_postgre_sql_parameters": {
              "computed": true,
              "description": "\u003cp\u003eParameters for Amazon Aurora PostgreSQL-Compatible Edition.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eThe Amazon Aurora PostgreSQL database to connect to.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eThe Amazon Aurora PostgreSQL-Compatible host to connect to.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port that Amazon Aurora PostgreSQL is listening on.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "databricks_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters that are required to connect to a Databricks data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host": {
                    "description": "\u003cp\u003eThe host name of the Databricks data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port for the Databricks data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "sql_endpoint_path": {
                    "description": "\u003cp\u003eThe HTTP path of the Databricks data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "maria_db_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for MariaDB.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "my_sql_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for MySQL.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "oracle_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Oracle.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eThe database.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eAn Oracle host.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "postgre_sql_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for PostgreSQL.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "presto_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Presto.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "description": "\u003cp\u003eCatalog.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "rds_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Amazon RDS.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "instance_id": {
                    "description": "\u003cp\u003eInstance ID.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "redshift_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Amazon Redshift. The \u003ccode\u003eClusterId\u003c/code\u003e field can be blank if\n            \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e are both set. The \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e fields can be blank if the \u003ccode\u003eClusterId\u003c/code\u003e field is set.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_id": {
                    "computed": true,
                    "description": "\u003cp\u003eCluster ID. This field can be blank if the \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e are\n            provided.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "computed": true,
                    "description": "\u003cp\u003eHost. This field can be blank if \u003ccode\u003eClusterId\u003c/code\u003e is provided.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iam_parameters": {
                    "computed": true,
                    "description": "\u003cp\u003eA structure that grants Amazon QuickSight access to your cluster and make a call to the \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e API. For more information on the \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e API, see \u003ca href=\"https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html\"\u003e\n               \u003ccode\u003eGetClusterCredentials\u003c/code\u003e\n            \u003c/a\u003e.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_create_database_user": {
                          "computed": true,
                          "description": "\u003cp\u003eAutomatically creates a database user. If your database doesn't have a \u003ccode\u003eDatabaseUser\u003c/code\u003e, set this parameter to \u003ccode\u003eTrue\u003c/code\u003e. If there is no \u003ccode\u003eDatabaseUser\u003c/code\u003e, Amazon QuickSight can't connect to your cluster. The \u003ccode\u003eRoleArn\u003c/code\u003e that you use for this operation must grant access to \u003ccode\u003eredshift:CreateClusterUser\u003c/code\u003e to successfully create the user.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "database_groups": {
                          "computed": true,
                          "description": "\u003cp\u003eA list of groups whose permissions will be granted to Amazon QuickSight to access the cluster. These permissions are combined with the permissions granted to Amazon QuickSight by the \u003ccode\u003eDatabaseUser\u003c/code\u003e. If you choose to include this parameter, the \u003ccode\u003eRoleArn\u003c/code\u003e must grant access to \u003ccode\u003eredshift:JoinGroup\u003c/code\u003e.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "database_user": {
                          "computed": true,
                          "description": "\u003cp\u003eThe user whose permissions and group memberships will be used by Amazon QuickSight to access the cluster. If this user already exists in your database, Amazon QuickSight is granted the same permissions that the user has. If the user doesn't exist, set the value of \u003ccode\u003eAutoCreateDatabaseUser\u003c/code\u003e to \u003ccode\u003eTrue\u003c/code\u003e to create a new user with PUBLIC permissions.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to allow Amazon QuickSight to call \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e on your cluster. The calling principal must have \u003ccode\u003eiam:PassRole\u003c/code\u003e access to pass the role to Amazon QuickSight. The role's trust policy must allow the Amazon QuickSight service principal to assume the role.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "identity_center_configuration": {
                    "computed": true,
                    "description": "\u003cp\u003eThe parameters for an IAM Identity Center configuration.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_identity_propagation": {
                          "computed": true,
                          "description": "\u003cp\u003eA Boolean option that controls whether Trusted Identity Propagation should be used.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort. This field can be blank if the \u003ccode\u003eClusterId\u003c/code\u003e is provided.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for S3.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "manifest_file_location": {
                    "description": "\u003cp\u003eAmazon S3 manifest file location.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "description": "\u003cp\u003eAmazon S3 bucket.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
                          "description": "\u003cp\u003eAmazon S3 key that identifies an object.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to override an account-wide role for a specific S3 data source. For example, say an account administrator has turned off all S3 access with an account-wide role. The administrator can then use \u003ccode\u003eRoleArn\u003c/code\u003e to bypass the account-wide role and allow S3 access for the single S3 data source that is specified in the structure, even if the account-wide role forbidding S3 access is still active.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "snowflake_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Snowflake.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "warehouse": {
                    "description": "\u003cp\u003eWarehouse.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "spark_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Spark.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sql_server_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for SQL Server.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "starburst_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters that are required to connect to a Starburst data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "description": "\u003cp\u003eThe catalog name for the Starburst data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eThe host name of the Starburst data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port for the Starburst data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "product_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "teradata_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Teradata.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "trino_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters that are required to connect to a Trino data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "description": "\u003cp\u003eThe catalog name for the Trino data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eThe host name of the Trino data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port for the Trino data source.\u003c/p\u003e",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the data source.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "\u003cp\u003eThe time that this data source was created.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "credentials": {
        "computed": true,
        "description": "\u003cp\u003eData source credentials. This is a variant type structure. For this structure to be\n            valid, only one of the attributes can be non-null.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_source_arn": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of a data source that has the credential pair that you\n            want to use. When \u003ccode\u003eCopySourceArn\u003c/code\u003e is not null, the credential pair from the\n            data source in the ARN is used as the credentials for the\n            \u003ccode\u003eDataSourceCredentials\u003c/code\u003e structure.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "credential_pair": {
              "computed": true,
              "description": "\u003cp\u003eThe combination of user name and password that are used as credentials.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alternate_data_source_parameters": {
                    "computed": true,
                    "description": "\u003cp\u003eA set of alternate data source parameters that you want to share for these\n            credentials. The credentials are applied in tandem with the data source parameters when\n            you copy a data source by using a create or update request. The API operation compares\n            the \u003ccode\u003eDataSourceParameters\u003c/code\u003e structure that's in the request with the\n            structures in the \u003ccode\u003eAlternateDataSourceParameters\u003c/code\u003e allow list. If the\n            structures are an exact match, the request is allowed to use the new data source with\n            the existing credentials. If the \u003ccode\u003eAlternateDataSourceParameters\u003c/code\u003e list is\n            null, the \u003ccode\u003eDataSourceParameters\u003c/code\u003e originally used with these\n                \u003ccode\u003eCredentials\u003c/code\u003e is automatically allowed.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "amazon_elasticsearch_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for OpenSearch.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "domain": {
                                "description": "\u003cp\u003eThe OpenSearch domain.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "amazon_open_search_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for OpenSearch.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "domain": {
                                "description": "\u003cp\u003eThe OpenSearch domain.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "athena_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eParameters for Amazon Athena.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "role_arn": {
                                "computed": true,
                                "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to override an account-wide role for a specific Athena data source. For example, say an account administrator has turned off all Athena access with an account-wide role. The administrator can then use \u003ccode\u003eRoleArn\u003c/code\u003e to bypass the account-wide role and allow Athena access for the single Athena data source that is specified in the structure, even if the account-wide role forbidding Athena access is still active.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "work_group": {
                                "computed": true,
                                "description": "\u003cp\u003eThe workgroup that Amazon Athena uses.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "aurora_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eParameters for Amazon Aurora.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "aurora_postgre_sql_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eParameters for Amazon Aurora PostgreSQL-Compatible Edition.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eThe Amazon Aurora PostgreSQL database to connect to.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eThe Amazon Aurora PostgreSQL-Compatible host to connect to.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003eThe port that Amazon Aurora PostgreSQL is listening on.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "databricks_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters that are required to connect to a Databricks data source.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "host": {
                                "description": "\u003cp\u003eThe host name of the Databricks data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003eThe port for the Databricks data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "sql_endpoint_path": {
                                "description": "\u003cp\u003eThe HTTP path of the Databricks data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "maria_db_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for MariaDB.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "my_sql_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for MySQL.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "oracle_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Oracle.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eThe database.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eAn Oracle host.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003eThe port.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "postgre_sql_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for PostgreSQL.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "presto_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Presto.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "catalog": {
                                "description": "\u003cp\u003eCatalog.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "rds_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Amazon RDS.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "instance_id": {
                                "description": "\u003cp\u003eInstance ID.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "redshift_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Amazon Redshift. The \u003ccode\u003eClusterId\u003c/code\u003e field can be blank if\n            \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e are both set. The \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e fields can be blank if the \u003ccode\u003eClusterId\u003c/code\u003e field is set.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cluster_id": {
                                "computed": true,
                                "description": "\u003cp\u003eCluster ID. This field can be blank if the \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e are\n            provided.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "computed": true,
                                "description": "\u003cp\u003eHost. This field can be blank if \u003ccode\u003eClusterId\u003c/code\u003e is provided.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "iam_parameters": {
                                "computed": true,
                                "description": "\u003cp\u003eA structure that grants Amazon QuickSight access to your cluster and make a call to the \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e API. For more information on the \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e API, see \u003ca href=\"https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html\"\u003e\n               \u003ccode\u003eGetClusterCredentials\u003c/code\u003e\n            \u003c/a\u003e.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "auto_create_database_user": {
                                      "computed": true,
                                      "description": "\u003cp\u003eAutomatically creates a database user. If your database doesn't have a \u003ccode\u003eDatabaseUser\u003c/code\u003e, set this parameter to \u003ccode\u003eTrue\u003c/code\u003e. If there is no \u003ccode\u003eDatabaseUser\u003c/code\u003e, Amazon QuickSight can't connect to your cluster. The \u003ccode\u003eRoleArn\u003c/code\u003e that you use for this operation must grant access to \u003ccode\u003eredshift:CreateClusterUser\u003c/code\u003e to successfully create the user.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "database_groups": {
                                      "computed": true,
                                      "description": "\u003cp\u003eA list of groups whose permissions will be granted to Amazon QuickSight to access the cluster. These permissions are combined with the permissions granted to Amazon QuickSight by the \u003ccode\u003eDatabaseUser\u003c/code\u003e. If you choose to include this parameter, the \u003ccode\u003eRoleArn\u003c/code\u003e must grant access to \u003ccode\u003eredshift:JoinGroup\u003c/code\u003e.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "database_user": {
                                      "computed": true,
                                      "description": "\u003cp\u003eThe user whose permissions and group memberships will be used by Amazon QuickSight to access the cluster. If this user already exists in your database, Amazon QuickSight is granted the same permissions that the user has. If the user doesn't exist, set the value of \u003ccode\u003eAutoCreateDatabaseUser\u003c/code\u003e to \u003ccode\u003eTrue\u003c/code\u003e to create a new user with PUBLIC permissions.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "role_arn": {
                                      "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to allow Amazon QuickSight to call \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e on your cluster. The calling principal must have \u003ccode\u003eiam:PassRole\u003c/code\u003e access to pass the role to Amazon QuickSight. The role's trust policy must allow the Amazon QuickSight service principal to assume the role.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "identity_center_configuration": {
                                "computed": true,
                                "description": "\u003cp\u003eThe parameters for an IAM Identity Center configuration.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "enable_identity_propagation": {
                                      "computed": true,
                                      "description": "\u003cp\u003eA Boolean option that controls whether Trusted Identity Propagation should be used.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort. This field can be blank if the \u003ccode\u003eClusterId\u003c/code\u003e is provided.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "s3_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for S3.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "manifest_file_location": {
                                "description": "\u003cp\u003eAmazon S3 manifest file location.\u003c/p\u003e",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "bucket": {
                                      "description": "\u003cp\u003eAmazon S3 bucket.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "description": "\u003cp\u003eAmazon S3 key that identifies an object.\u003c/p\u003e",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              },
                              "role_arn": {
                                "computed": true,
                                "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to override an account-wide role for a specific S3 data source. For example, say an account administrator has turned off all S3 access with an account-wide role. The administrator can then use \u003ccode\u003eRoleArn\u003c/code\u003e to bypass the account-wide role and allow S3 access for the single S3 data source that is specified in the structure, even if the account-wide role forbidding S3 access is still active.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "snowflake_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Snowflake.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "warehouse": {
                                "description": "\u003cp\u003eWarehouse.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "spark_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Spark.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "sql_server_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for SQL Server.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "starburst_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters that are required to connect to a Starburst data source.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "catalog": {
                                "description": "\u003cp\u003eThe catalog name for the Starburst data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eThe host name of the Starburst data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003eThe port for the Starburst data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "product_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "teradata_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters for Teradata.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "database": {
                                "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eHost.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003ePort.\u003c/p\u003e",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "trino_parameters": {
                          "computed": true,
                          "description": "\u003cp\u003eThe parameters that are required to connect to a Trino data source.\u003c/p\u003e",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "catalog": {
                                "description": "\u003cp\u003eThe catalog name for the Trino data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "host": {
                                "description": "\u003cp\u003eThe host name of the Trino data source.\u003c/p\u003e",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
                                "description": "\u003cp\u003eThe port for the Trino data source.\u003c/p\u003e",
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "password": {
                    "description": "\u003cp\u003ePassword.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "\u003cp\u003eUser name.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "secret_arn": {
              "computed": true,
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the secret associated with the data source in Amazon Secrets Manager.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "data_source_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_source_parameters": {
        "computed": true,
        "description": "\u003cp\u003eThe parameters that Amazon QuickSight uses to connect to your underlying data source.\n            This is a variant type structure. For this structure to be valid, only one of the\n            attributes can be non-null.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "amazon_elasticsearch_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for OpenSearch.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "description": "\u003cp\u003eThe OpenSearch domain.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "amazon_open_search_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for OpenSearch.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "description": "\u003cp\u003eThe OpenSearch domain.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "athena_parameters": {
              "computed": true,
              "description": "\u003cp\u003eParameters for Amazon Athena.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to override an account-wide role for a specific Athena data source. For example, say an account administrator has turned off all Athena access with an account-wide role. The administrator can then use \u003ccode\u003eRoleArn\u003c/code\u003e to bypass the account-wide role and allow Athena access for the single Athena data source that is specified in the structure, even if the account-wide role forbidding Athena access is still active.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "work_group": {
                    "computed": true,
                    "description": "\u003cp\u003eThe workgroup that Amazon Athena uses.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "aurora_parameters": {
              "computed": true,
              "description": "\u003cp\u003eParameters for Amazon Aurora.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "aurora_postgre_sql_parameters": {
              "computed": true,
              "description": "\u003cp\u003eParameters for Amazon Aurora PostgreSQL-Compatible Edition.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eThe Amazon Aurora PostgreSQL database to connect to.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eThe Amazon Aurora PostgreSQL-Compatible host to connect to.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port that Amazon Aurora PostgreSQL is listening on.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "databricks_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters that are required to connect to a Databricks data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host": {
                    "description": "\u003cp\u003eThe host name of the Databricks data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port for the Databricks data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "sql_endpoint_path": {
                    "description": "\u003cp\u003eThe HTTP path of the Databricks data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "maria_db_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for MariaDB.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "my_sql_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for MySQL.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "oracle_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Oracle.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eThe database.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eAn Oracle host.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "postgre_sql_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for PostgreSQL.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "presto_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Presto.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "description": "\u003cp\u003eCatalog.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "rds_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Amazon RDS.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "instance_id": {
                    "description": "\u003cp\u003eInstance ID.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "redshift_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Amazon Redshift. The \u003ccode\u003eClusterId\u003c/code\u003e field can be blank if\n            \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e are both set. The \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e fields can be blank if the \u003ccode\u003eClusterId\u003c/code\u003e field is set.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_id": {
                    "computed": true,
                    "description": "\u003cp\u003eCluster ID. This field can be blank if the \u003ccode\u003eHost\u003c/code\u003e and \u003ccode\u003ePort\u003c/code\u003e are\n            provided.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "computed": true,
                    "description": "\u003cp\u003eHost. This field can be blank if \u003ccode\u003eClusterId\u003c/code\u003e is provided.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iam_parameters": {
                    "computed": true,
                    "description": "\u003cp\u003eA structure that grants Amazon QuickSight access to your cluster and make a call to the \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e API. For more information on the \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e API, see \u003ca href=\"https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html\"\u003e\n               \u003ccode\u003eGetClusterCredentials\u003c/code\u003e\n            \u003c/a\u003e.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "auto_create_database_user": {
                          "computed": true,
                          "description": "\u003cp\u003eAutomatically creates a database user. If your database doesn't have a \u003ccode\u003eDatabaseUser\u003c/code\u003e, set this parameter to \u003ccode\u003eTrue\u003c/code\u003e. If there is no \u003ccode\u003eDatabaseUser\u003c/code\u003e, Amazon QuickSight can't connect to your cluster. The \u003ccode\u003eRoleArn\u003c/code\u003e that you use for this operation must grant access to \u003ccode\u003eredshift:CreateClusterUser\u003c/code\u003e to successfully create the user.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "database_groups": {
                          "computed": true,
                          "description": "\u003cp\u003eA list of groups whose permissions will be granted to Amazon QuickSight to access the cluster. These permissions are combined with the permissions granted to Amazon QuickSight by the \u003ccode\u003eDatabaseUser\u003c/code\u003e. If you choose to include this parameter, the \u003ccode\u003eRoleArn\u003c/code\u003e must grant access to \u003ccode\u003eredshift:JoinGroup\u003c/code\u003e.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "database_user": {
                          "computed": true,
                          "description": "\u003cp\u003eThe user whose permissions and group memberships will be used by Amazon QuickSight to access the cluster. If this user already exists in your database, Amazon QuickSight is granted the same permissions that the user has. If the user doesn't exist, set the value of \u003ccode\u003eAutoCreateDatabaseUser\u003c/code\u003e to \u003ccode\u003eTrue\u003c/code\u003e to create a new user with PUBLIC permissions.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to allow Amazon QuickSight to call \u003ccode\u003eredshift:GetClusterCredentials\u003c/code\u003e on your cluster. The calling principal must have \u003ccode\u003eiam:PassRole\u003c/code\u003e access to pass the role to Amazon QuickSight. The role's trust policy must allow the Amazon QuickSight service principal to assume the role.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "identity_center_configuration": {
                    "computed": true,
                    "description": "\u003cp\u003eThe parameters for an IAM Identity Center configuration.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_identity_propagation": {
                          "computed": true,
                          "description": "\u003cp\u003eA Boolean option that controls whether Trusted Identity Propagation should be used.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort. This field can be blank if the \u003ccode\u003eClusterId\u003c/code\u003e is provided.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "s3_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for S3.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "manifest_file_location": {
                    "description": "\u003cp\u003eAmazon S3 manifest file location.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "description": "\u003cp\u003eAmazon S3 bucket.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
                          "description": "\u003cp\u003eAmazon S3 key that identifies an object.\u003c/p\u003e",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "\u003cp\u003eUse the \u003ccode\u003eRoleArn\u003c/code\u003e structure to override an account-wide role for a specific S3 data source. For example, say an account administrator has turned off all S3 access with an account-wide role. The administrator can then use \u003ccode\u003eRoleArn\u003c/code\u003e to bypass the account-wide role and allow S3 access for the single S3 data source that is specified in the structure, even if the account-wide role forbidding S3 access is still active.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "snowflake_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Snowflake.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "warehouse": {
                    "description": "\u003cp\u003eWarehouse.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "spark_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Spark.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sql_server_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for SQL Server.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "starburst_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters that are required to connect to a Starburst data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "description": "\u003cp\u003eThe catalog name for the Starburst data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eThe host name of the Starburst data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port for the Starburst data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "product_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "teradata_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters for Teradata.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database": {
                    "description": "\u003cp\u003eDatabase.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eHost.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003ePort.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "trino_parameters": {
              "computed": true,
              "description": "\u003cp\u003eThe parameters that are required to connect to a Trino data source.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog": {
                    "description": "\u003cp\u003eThe catalog name for the Trino data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "host": {
                    "description": "\u003cp\u003eThe host name of the Trino data source.\u003c/p\u003e",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "\u003cp\u003eThe port for the Trino data source.\u003c/p\u003e",
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
      "error_info": {
        "computed": true,
        "description": "\u003cp\u003eError information for the data source creation or update.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "message": {
              "computed": true,
              "description": "\u003cp\u003eError message.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "\u003cp\u003eThe last time that this data source was updated.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n         \u003cul\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n               \u003cp\u003eThe ARN of an Amazon Web Services account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across Amazon Web Services accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource": {
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
      "ssl_properties": {
        "computed": true,
        "description": "\u003cp\u003eSecure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your\n            underlying data source.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disable_ssl": {
              "computed": true,
              "description": "\u003cp\u003eA Boolean option to control whether SSL should be disabled.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "\u003cp\u003eTag key.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "\u003cp\u003eTag value.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_connection_properties": {
        "computed": true,
        "description": "\u003cp\u003eVPC connection properties.\u003c/p\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "vpc_connection_arn": {
              "description": "\u003cp\u003eThe Amazon Resource Name (ARN) for the VPC connection.\u003c/p\u003e",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Definition of the AWS::QuickSight::DataSource Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightDataSource), &result)
	return &result
}
