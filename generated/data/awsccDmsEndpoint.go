package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsEndpoint = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description": "The name of the endpoint database. For a MySQL source or target endpoint, don't specify DatabaseName. To migrate to a specific database, use this setting and targetDbType.",
        "description_kind": "plain",
        "type": "string"
      },
      "doc_db_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target DocumentDB endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "docs_to_investigate": {
              "computed": true,
              "description": "Indicates the number of documents to preview to determine the document organization. Use this setting when NestingLevel is set to \"one\".",
              "description_kind": "plain",
              "type": "number"
            },
            "extract_doc_id": {
              "computed": true,
              "description": "Specifies the document ID. Use this setting when NestingLevel is set to \"none\"",
              "description_kind": "plain",
              "type": "bool"
            },
            "nesting_level": {
              "computed": true,
              "description": "Specifies either document or table mode.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret. The role must allow the iam:PassRole action. SecretsManagerSecret has the value of the AWS Secrets Manager secret that allows access to the DocumentDB endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret. The role must allow the iam:PassRole action. SecretsManagerSecret has the value of the AWS Secrets Manager secret that allows access to the DocumentDB endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "dynamo_db_settings": {
        "computed": true,
        "description": "Settings in JSON format for the target Amazon DynamoDB endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "service_access_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) used by the service to access the IAM role. The role must allow the iam:PassRole action.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "elasticsearch_settings": {
        "computed": true,
        "description": "Settings in JSON format for the target OpenSearch endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_uri": {
              "computed": true,
              "description": "The endpoint for the OpenSearch cluster. AWS DMS uses HTTPS if a transport protocol (either HTTP or HTTPS) isn't specified.",
              "description_kind": "plain",
              "type": "string"
            },
            "error_retry_duration": {
              "computed": true,
              "description": "The maximum number of seconds for which DMS retries failed API requests to the OpenSearch cluster.",
              "description_kind": "plain",
              "type": "number"
            },
            "full_load_error_percentage": {
              "computed": true,
              "description": "The maximum percentage of records that can fail to be written before a full load operation stops.",
              "description_kind": "plain",
              "type": "number"
            },
            "service_access_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) used by the service to access the IAM role. The role must allow the iam:PassRole action.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "endpoint_arn": {
        "computed": true,
        "description": "The endpoint ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_identifier": {
        "computed": true,
        "description": "The database endpoint identifier. Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_type": {
        "computed": true,
        "description": "The type of endpoint. Valid values are source and target.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_name": {
        "computed": true,
        "description": "The type of engine for the endpoint, depending on the EndpointType value.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_id": {
        "computed": true,
        "description": "A value that can be used for cross-account validation.",
        "description_kind": "plain",
        "type": "string"
      },
      "extra_connection_attributes": {
        "computed": true,
        "description": "Additional attributes associated with the connection",
        "description_kind": "plain",
        "type": "string"
      },
      "gcp_my_sql_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source GCP MySQL endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "after_connect_script": {
              "computed": true,
              "description": "Specifies a script to run immediately after AWS DMS connects to the endpoint. The migration task continues running regardless if the SQL statement succeeds or fails.",
              "description_kind": "plain",
              "type": "string"
            },
            "clean_source_metadata_on_mismatch": {
              "computed": true,
              "description": "Adjusts the behavior of AWS DMS when migrating from an SQL Server source database that is hosted as part of an Always On availability group cluster. If you need AWS DMS to poll all the nodes in the Always On cluster for transaction backups, set this attribute to false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "database_name": {
              "computed": true,
              "description": "Database name for the endpoint. For a MySQL source or target endpoint, don't explicitly specify the database using the DatabaseName request parameter on either the CreateEndpoint or ModifyEndpoint API call. Specifying DatabaseName when you create or modify a MySQL endpoint replicates all the task tables to this single database. For MySQL endpoints, you specify the database only when you specify the schema in the table-mapping rules of the AWS DMS task.",
              "description_kind": "plain",
              "type": "string"
            },
            "events_poll_interval": {
              "computed": true,
              "description": "Specifies how often to check the binary log for new changes/events when the database is idle. The default is five seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_file_size": {
              "computed": true,
              "description": "Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.",
              "description_kind": "plain",
              "type": "number"
            },
            "parallel_load_threads": {
              "computed": true,
              "description": "Improves performance when loading data into the MySQL-compatible target database. Specifies how many threads to use to load the data into the MySQL-compatible target database. Setting a large number of threads can have an adverse effect on database performance, because a separate connection is required for each thread. The default is one.",
              "description_kind": "plain",
              "type": "number"
            },
            "password": {
              "computed": true,
              "description": "Endpoint connection password.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port used by the endpoint database.",
              "description_kind": "plain",
              "type": "number"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret. The role must allow the iam:PassRole action. SecretsManagerSecret has the value of the AWS Secrets Manager secret that allows access to the MySQL endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MySQL endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "server_name": {
              "computed": true,
              "description": "The MySQL host name.",
              "description_kind": "plain",
              "type": "string"
            },
            "server_timezone": {
              "computed": true,
              "description": "Specifies the time zone for the source MySQL database. Don't enclose time zones in single quotation marks.",
              "description_kind": "plain",
              "type": "string"
            },
            "username": {
              "computed": true,
              "description": "Specifies the time zone for the source MySQL database. Don't enclose time zones in single quotation marks.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "ibm_db_2_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source IBM Db2 LUW endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "current_lsn": {
              "computed": true,
              "description": "For ongoing replication (CDC), use CurrentLSN to specify a log sequence number (LSN) where you want the replication to start.",
              "description_kind": "plain",
              "type": "string"
            },
            "keep_csv_files": {
              "computed": true,
              "description": "If true, AWS DMS saves any .csv files to the Db2 LUW target that were used to replicate data. DMS uses these files for analysis and troubleshooting.",
              "description_kind": "plain",
              "type": "bool"
            },
            "load_timeout": {
              "computed": true,
              "description": "The amount of time (in milliseconds) before AWS DMS times out operations performed by DMS on the Db2 target. The default value is 1200 (20 minutes).",
              "description_kind": "plain",
              "type": "number"
            },
            "max_file_size": {
              "computed": true,
              "description": "Specifies the maximum size (in KB) of .csv files used to transfer data to Db2 LUW.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_k_bytes_per_read": {
              "computed": true,
              "description": "Maximum number of bytes per read, as a NUMBER value. The default is 64 KB.",
              "description_kind": "plain",
              "type": "number"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret. The role must allow the iam:PassRole action. SecretsManagerSecret has the value ofthe AWS Secrets Manager secret that allows access to the Db2 LUW endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the IBMDB2 endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "set_data_capture_changes": {
              "computed": true,
              "description": "Enables ongoing replication (CDC) as a BOOLEAN value. The default is true.",
              "description_kind": "plain",
              "type": "bool"
            },
            "write_buffer_size": {
              "computed": true,
              "description": "The size (in KB) of the in-memory file write buffer used when generating .csv files on the local disk on the DMS replication instance. The default value is 1024 (1 MB).",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kafka_settings": {
        "computed": true,
        "description": "Settings in JSON format for the target Apache Kafka endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "broker": {
              "computed": true,
              "description": "A comma-separated list of one or more broker locations in your Kafka cluster that host your Kafka instance. Specify each broker location in the form broker-hostname-or-ip:port ",
              "description_kind": "plain",
              "type": "string"
            },
            "include_control_details": {
              "computed": true,
              "description": "Shows detailed control information for table definition, column definition, and table and column changes in the Kafka message output. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_null_and_empty": {
              "computed": true,
              "description": "Include NULL and empty columns for records migrated to the endpoint. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_partition_value": {
              "computed": true,
              "description": "Shows the partition value within the Kafka message output unless the partition type is schema-table-type. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_table_alter_operations": {
              "computed": true,
              "description": "Includes any data definition language (DDL) operations that change the table in the control data, such as rename-table, drop-table, add-column, drop-column, and rename-column. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_transaction_details": {
              "computed": true,
              "description": "Provides detailed transaction information from the source database. This information includes a commit timestamp, a log position, and values for transaction_id, previous transaction_id, and transaction_record_id (the record offset within a transaction). The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "message_format": {
              "computed": true,
              "description": "The output format for the records created on the endpoint. The message format is JSON (default) or JSON_UNFORMATTED (a single line with no tab).",
              "description_kind": "plain",
              "type": "string"
            },
            "message_max_bytes": {
              "computed": true,
              "description": "The maximum size in bytes for records created on the endpoint The default is 1,000,000.",
              "description_kind": "plain",
              "type": "number"
            },
            "no_hex_prefix": {
              "computed": true,
              "description": "Set this optional parameter to true to avoid adding a '0x' prefix to raw data in hexadecimal format. For example, by default, AWS DMS adds a '0x' prefix to the LOB column type in hexadecimal format moving from an Oracle source to a Kafka target. Use the NoHexPrefix endpoint setting to enable migration of RAW data type columns without adding the '0x' prefix.",
              "description_kind": "plain",
              "type": "bool"
            },
            "partition_include_schema_table": {
              "computed": true,
              "description": "Prefixes schema and table names to partition values, when the partition type is primary-key-type.",
              "description_kind": "plain",
              "type": "bool"
            },
            "sasl_password": {
              "computed": true,
              "description": "The secure password that you created when you first set up your Amazon MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.",
              "description_kind": "plain",
              "type": "string"
            },
            "sasl_user_name": {
              "computed": true,
              "description": "The secure user name you created when you first set up your Amazon MSK cluster to validate a client identity and make an encrypted connection between server and client using SASL-SSL authentication.",
              "description_kind": "plain",
              "type": "string"
            },
            "security_protocol": {
              "computed": true,
              "description": "Set secure connection to a Kafka target endpoint using Transport Layer Security (TLS). Options include ssl-encryption, ssl-authentication, and sasl-ssl. sasl-ssl requires SaslUsername and SaslPassword.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_ca_certificate_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the private certificate authority (CA) cert that AWS DMS uses to securely connect to your Kafka target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_client_certificate_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the client certificate used to securely connect to a Kafka target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_client_key_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the client private key used to securely connect to a Kafka target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_client_key_password": {
              "computed": true,
              "description": "The password for the client private key used to securely connect to a Kafka target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "topic": {
              "computed": true,
              "description": "The topic to which you migrate the data. If you don't specify a topic, AWS DMS specifies \"kafka-default-topic\" as the migration topic.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kinesis_settings": {
        "computed": true,
        "description": "Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "include_control_details": {
              "computed": true,
              "description": "Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_null_and_empty": {
              "computed": true,
              "description": "Include NULL and empty columns for records migrated to the endpoint. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_partition_value": {
              "computed": true,
              "description": "Shows the partition value within the Kinesis message output, unless the partition type is schema-table-type. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_table_alter_operations": {
              "computed": true,
              "description": "Includes any data definition language (DDL) operations that change the table in the control data, such as rename-table, drop-table, add-column, drop-column, and rename-column. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "include_transaction_details": {
              "computed": true,
              "description": "Provides detailed transaction information from the source database.",
              "description_kind": "plain",
              "type": "bool"
            },
            "message_format": {
              "computed": true,
              "description": "The output format for the records created on the endpoint. The message format is JSON (default) or JSON_UNFORMATTED (a single line with no tab).",
              "description_kind": "plain",
              "type": "string"
            },
            "no_hex_prefix": {
              "computed": true,
              "description": "Set this optional parameter to true to avoid adding a '0x' prefix to raw data in hexadecimal format.",
              "description_kind": "plain",
              "type": "bool"
            },
            "partition_include_schema_table": {
              "computed": true,
              "description": "Prefixes schema and table names to partition values, when the partition type is primary-key-type.",
              "description_kind": "plain",
              "type": "bool"
            },
            "service_access_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the IAM role that AWS DMS uses to write to the Kinesis data stream. The role must allow the iam:PassRole action.",
              "description_kind": "plain",
              "type": "string"
            },
            "stream_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the Amazon Kinesis Data Streams endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kms_key_id": {
        "computed": true,
        "description": "An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.If you don't specify a value for the KmsKeyId parameter, AWS DMS uses your default encryption key.",
        "description_kind": "plain",
        "type": "string"
      },
      "microsoft_sql_server_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target Microsoft SQL Server endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bcp_packet_size": {
              "computed": true,
              "description": "The maximum size of the packets (in bytes) used to transfer data using BCP.",
              "description_kind": "plain",
              "type": "number"
            },
            "control_tables_file_group": {
              "computed": true,
              "description": "Specifies a file group for the AWS DMS internal tables.",
              "description_kind": "plain",
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description": "Database name for the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "force_lob_lookup": {
              "computed": true,
              "description": "Forces LOB lookup on inline LOB.",
              "description_kind": "plain",
              "type": "bool"
            },
            "password": {
              "computed": true,
              "description": "Endpoint connection password.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "Endpoint TCP port.",
              "description_kind": "plain",
              "type": "number"
            },
            "query_single_always_on_node": {
              "computed": true,
              "description": "Cleans and recreates table metadata information on the replication instance when a mismatch occurs. An example is a situation where running an alter DDL statement on a table might result in different information about the table cached in the replication instance.",
              "description_kind": "plain",
              "type": "bool"
            },
            "read_backup_only": {
              "computed": true,
              "description": "When this attribute is set to Y, AWS DMS only reads changes from transaction log backups and doesn't read from the active transaction log file during ongoing replication. Setting this parameter to Y enables you to control active transaction log file growth during full load and ongoing replication tasks. However, it can add some source latency to ongoing replication.",
              "description_kind": "plain",
              "type": "bool"
            },
            "safeguard_policy": {
              "computed": true,
              "description": "Use this attribute to minimize the need to access the backup log and enable AWS DMS to prevent truncation using one of the following two methods.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MicrosoftSQLServer endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "server_name": {
              "computed": true,
              "description": "Fully qualified domain name of the endpoint. For an Amazon RDS SQL Server instance, this is the output of DescribeDBInstances, in the Endpoint.Address field.",
              "description_kind": "plain",
              "type": "string"
            },
            "tlog_access_mode": {
              "computed": true,
              "description": "Indicates the mode used to fetch CDC data.",
              "description_kind": "plain",
              "type": "string"
            },
            "trim_space_in_char": {
              "computed": true,
              "description": "Use the TrimSpaceInChar source endpoint setting to right-trim data on CHAR and NCHAR data types during migration. Setting TrimSpaceInChar does not left-trim data. The default value is true.",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_bcp_full_load": {
              "computed": true,
              "description": "Use this to attribute to transfer data for full-load operations using BCP. When the target table contains an identity column that does not exist in the source table, you must disable the use BCP for loading table option.",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_third_party_backup_device": {
              "computed": true,
              "description": "When this attribute is set to Y, DMS processes third-party transaction log backups if they are created in native format.",
              "description_kind": "plain",
              "type": "bool"
            },
            "username": {
              "computed": true,
              "description": "Endpoint connection user name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "mongo_db_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source MongoDB endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_mechanism": {
              "computed": true,
              "description": "The authentication mechanism you use to access the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "auth_source": {
              "computed": true,
              "description": "The MongoDB database name. This setting isn't used when AuthType is set to \"no\".",
              "description_kind": "plain",
              "type": "string"
            },
            "auth_type": {
              "computed": true,
              "description": "The authentication type you use to access the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description": "The database name on the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "docs_to_investigate": {
              "computed": true,
              "description": "Indicates the number of documents to preview to determine the document organization. Use this setting when NestingLevel is set to \"one\".",
              "description_kind": "plain",
              "type": "string"
            },
            "extract_doc_id": {
              "computed": true,
              "description": "Specifies the document ID. Use this setting when NestingLevel is set to \"none\".",
              "description_kind": "plain",
              "type": "string"
            },
            "nesting_level": {
              "computed": true,
              "description": "Specifies either document or table mode.",
              "description_kind": "plain",
              "type": "string"
            },
            "password": {
              "computed": true,
              "description": "The password for the user account you use to access the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port value for the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "number"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MongoDB endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "server_name": {
              "computed": true,
              "description": "The name of the server on the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "username": {
              "computed": true,
              "description": "The user name you use to access the MongoDB source endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "my_sql_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target MySQL endpoin",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "after_connect_script": {
              "computed": true,
              "description": "Specifies a script to run immediately after AWS DMS connects to the endpoint. The migration task continues running regardless if the SQL statement succeeds or fails.",
              "description_kind": "plain",
              "type": "string"
            },
            "clean_source_metadata_on_mismatch": {
              "computed": true,
              "description": "Cleans and recreates table metadata information on the replication instance when a mismatch occurs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "events_poll_interval": {
              "computed": true,
              "description": "Specifies how often to check the binary log for new changes/events when the database is idle. The default is five seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_file_size": {
              "computed": true,
              "description": "Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.",
              "description_kind": "plain",
              "type": "number"
            },
            "parallel_load_threads": {
              "computed": true,
              "description": "Improves performance when loading data into the MySQL-compatible target database. Specifies how many threads to use to load the data into the MySQL-compatible target database.",
              "description_kind": "plain",
              "type": "number"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the MySQL endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "server_timezone": {
              "computed": true,
              "description": "Specifies the time zone for the source MySQL database.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_db_type": {
              "computed": true,
              "description": "Specifies where to migrate source tables on the target, either to a single database or multiple databases.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "neptune_settings": {
        "computed": true,
        "description": "Settings in JSON format for the target Amazon Neptune endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "error_retry_duration": {
              "computed": true,
              "description": "The number of milliseconds for AWS DMS to wait to retry a bulk-load of migrated graph data to the Neptune target database before raising an error. The default is 250.",
              "description_kind": "plain",
              "type": "number"
            },
            "iam_auth_enabled": {
              "computed": true,
              "description": "If you want IAM authorization enabled for this endpoint, set this parameter to true.",
              "description_kind": "plain",
              "type": "bool"
            },
            "max_file_size": {
              "computed": true,
              "description": "The maximum size in kilobytes of migrated graph data stored in a .csv file before AWS DMS bulk-loads the data to the Neptune target database.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_retry_count": {
              "computed": true,
              "description": "The number of times for AWS DMS to retry a bulk load of migrated graph data to the Neptune target database before raising an error. The default is 5.",
              "description_kind": "plain",
              "type": "number"
            },
            "s3_bucket_folder": {
              "computed": true,
              "description": "A folder path where you want AWS DMS to store migrated graph data in the S3 bucket specified by S3BucketName",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket_name": {
              "computed": true,
              "description": "The name of the Amazon S3 bucket where AWS DMS can temporarily store migrated graph data in .csv files before bulk-loading it to the Neptune target database.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_access_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the service role that you created for the Neptune target endpoint. The role must allow the iam:PassRole action.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "oracle_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target Oracle endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_alternate_directly": {
              "computed": true,
              "description": "Set this attribute to false in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.",
              "description_kind": "plain",
              "type": "bool"
            },
            "add_supplemental_logging": {
              "computed": true,
              "description": "Set this attribute to set up table-level supplemental logging for the Oracle database. This attribute enables PRIMARY KEY supplemental logging on all tables selected for a migration task.",
              "description_kind": "plain",
              "type": "bool"
            },
            "additional_archived_log_dest_id": {
              "computed": true,
              "description": "Set this attribute with ArchivedLogDestId in a primary/ standby setup",
              "description_kind": "plain",
              "type": "number"
            },
            "allow_select_nested_tables": {
              "computed": true,
              "description": "Set this attribute to true to enable replication of Oracle tables containing columns that are nested tables or defined types.",
              "description_kind": "plain",
              "type": "bool"
            },
            "archived_log_dest_id": {
              "computed": true,
              "description": "Specifies the ID of the destination for the archived redo logs.",
              "description_kind": "plain",
              "type": "number"
            },
            "archived_logs_only": {
              "computed": true,
              "description": "When this field is set to True, AWS DMS only accesses the archived redo logs",
              "description_kind": "plain",
              "type": "bool"
            },
            "asm_password": {
              "computed": true,
              "description": "For an Oracle source endpoint, your Oracle Automatic Storage Management (ASM) password.",
              "description_kind": "plain",
              "type": "string"
            },
            "asm_server": {
              "computed": true,
              "description": "For an Oracle source endpoint, your ASM server address.",
              "description_kind": "plain",
              "type": "string"
            },
            "asm_user": {
              "computed": true,
              "description": "For an Oracle source endpoint, your ASM user name.",
              "description_kind": "plain",
              "type": "string"
            },
            "char_length_semantics": {
              "computed": true,
              "description": "Specifies whether the length of a character column is in bytes or in characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "direct_path_no_log": {
              "computed": true,
              "description": "When set to true, this attribute helps to increase the commit rate on the Oracle target database by writing directly to tables and not writing a trail to database logs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "direct_path_parallel_load": {
              "computed": true,
              "description": "When set to true, this attribute specifies a parallel load when useDirectPathFullLoad is set to Y.",
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_homogenous_tablespace": {
              "computed": true,
              "description": "Set this attribute to enable homogenous tablespace replication and create existing tables or indexes under the same tablespace on the target.",
              "description_kind": "plain",
              "type": "bool"
            },
            "extra_archived_log_dest_ids": {
              "computed": true,
              "description": "Specifies the IDs of one more destinations for one or more archived redo logs.",
              "description_kind": "plain",
              "type": [
                "list",
                "number"
              ]
            },
            "fail_tasks_on_lob_truncation": {
              "computed": true,
              "description": "When set to true, this attribute causes a task to fail if the actual size of an LOB column is greater than the specified LobMaxSize.",
              "description_kind": "plain",
              "type": "bool"
            },
            "number_datatype_scale": {
              "computed": true,
              "description": "Specifies the number scale. You can select a scale up to 38, or you can select FLOAT. By default, the NUMBER data type is converted to precision 38, scale 10.",
              "description_kind": "plain",
              "type": "number"
            },
            "oracle_path_prefix": {
              "computed": true,
              "description": "Set this string attribute to the required value in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.",
              "description_kind": "plain",
              "type": "string"
            },
            "parallel_asm_read_threads": {
              "computed": true,
              "description": "Set this attribute to change the number of threads that DMS configures to perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).",
              "description_kind": "plain",
              "type": "number"
            },
            "read_ahead_blocks": {
              "computed": true,
              "description": "Set this attribute to change the number of read-ahead blocks that DMS configures to perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).",
              "description_kind": "plain",
              "type": "number"
            },
            "read_table_space_name": {
              "computed": true,
              "description": "When set to true, this attribute supports tablespace replication.",
              "description_kind": "plain",
              "type": "bool"
            },
            "replace_path_prefix": {
              "computed": true,
              "description": "Set this attribute to true in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.",
              "description_kind": "plain",
              "type": "bool"
            },
            "retry_interval": {
              "computed": true,
              "description": "Specifies the number of seconds that the system waits before resending a query.",
              "description_kind": "plain",
              "type": "number"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_oracle_asm_access_role_arn": {
              "computed": true,
              "description": "Required only if your Oracle endpoint uses Advanced Storage Manager (ASM).",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_oracle_asm_secret_id": {
              "computed": true,
              "description": "Required only if your Oracle endpoint uses Advanced Storage Manager (ASM).",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the Oracle endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "security_db_encryption": {
              "computed": true,
              "description": "For an Oracle source endpoint, the transparent data encryption (TDE) password required by AWM DMS to access Oracle redo logs encrypted by TDE using Binary Reader.",
              "description_kind": "plain",
              "type": "string"
            },
            "security_db_encryption_name": {
              "computed": true,
              "description": "For an Oracle source endpoint, the name of a key used for the transparent data encryption (TDE) of the columns and tablespaces in an Oracle source database that is encrypted using TDE.",
              "description_kind": "plain",
              "type": "string"
            },
            "spatial_data_option_to_geo_json_function_name": {
              "computed": true,
              "description": "Use this attribute to convert SDO_GEOMETRY to GEOJSON format. By default, DMS calls the SDO2GEOJSON custom function if present and accessible. Or you can create your own custom function that mimics the operation of SDOGEOJSON and set SpatialDataOptionToGeoJsonFunctionName to call it instead.",
              "description_kind": "plain",
              "type": "string"
            },
            "standby_delay_time": {
              "computed": true,
              "description": "Use this attribute to specify a time in minutes for the delay in standby sync.",
              "description_kind": "plain",
              "type": "number"
            },
            "use_alternate_folder_for_online": {
              "computed": true,
              "description": "Set this attribute to true in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_b_file": {
              "computed": true,
              "description": "Set this attribute to True to capture change data using the Binary Reader utility.",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_direct_path_full_load": {
              "computed": true,
              "description": "Set this attribute to True to have AWS DMS use a direct path full load.",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_logminer_reader": {
              "computed": true,
              "description": "Set this attribute to True to capture change data using the Oracle LogMiner utility (the default).",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_path_prefix": {
              "computed": true,
              "description": "Set this string attribute to the required value in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "password": {
        "computed": true,
        "description": "The password to be used to log in to the endpoint database.",
        "description_kind": "plain",
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port used by the endpoint database.",
        "description_kind": "plain",
        "type": "number"
      },
      "postgre_sql_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target PostgreSQL endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "after_connect_script": {
              "computed": true,
              "description": "For use with change data capture (CDC) only, this attribute has AWS DMS bypass foreign keys and user triggers to reduce the time it takes to bulk load data.",
              "description_kind": "plain",
              "type": "string"
            },
            "babelfish_database_name": {
              "computed": true,
              "description": "The Babelfish for Aurora PostgreSQL database name for the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "capture_ddls": {
              "computed": true,
              "description": "To capture DDL events, AWS DMS creates various artifacts in the PostgreSQL database when the task starts. You can later remove these artifacts.",
              "description_kind": "plain",
              "type": "bool"
            },
            "database_mode": {
              "computed": true,
              "description": "Specifies the default behavior of the replication's handling of PostgreSQL- compatible endpoints that require some additional configuration, such as Babelfish endpoints.",
              "description_kind": "plain",
              "type": "string"
            },
            "ddl_artifacts_schema": {
              "computed": true,
              "description": "The schema in which the operational DDL database artifacts are created.",
              "description_kind": "plain",
              "type": "string"
            },
            "execute_timeout": {
              "computed": true,
              "description": "Sets the client statement timeout for the PostgreSQL instance, in seconds. The default value is 60 seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "fail_tasks_on_lob_truncation": {
              "computed": true,
              "description": "When set to true, this value causes a task to fail if the actual size of a LOB column is greater than the specified LobMaxSize.",
              "description_kind": "plain",
              "type": "bool"
            },
            "heartbeat_enable": {
              "computed": true,
              "description": "The write-ahead log (WAL) heartbeat feature mimics a dummy transaction.",
              "description_kind": "plain",
              "type": "bool"
            },
            "heartbeat_frequency": {
              "computed": true,
              "description": "Sets the WAL heartbeat frequency (in minutes).",
              "description_kind": "plain",
              "type": "number"
            },
            "heartbeat_schema": {
              "computed": true,
              "description": "Sets the schema in which the heartbeat artifacts are created.",
              "description_kind": "plain",
              "type": "string"
            },
            "map_boolean_as_boolean": {
              "computed": true,
              "description": "When true, lets PostgreSQL migrate the boolean type as boolean.",
              "description_kind": "plain",
              "type": "bool"
            },
            "max_file_size": {
              "computed": true,
              "description": "Specifies the maximum size (in KB) of any .csv file used to transfer data to PostgreSQL.",
              "description_kind": "plain",
              "type": "number"
            },
            "plugin_name": {
              "computed": true,
              "description": "Specifies the plugin to use to create a replication slot.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the PostgreSQL endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "slot_name": {
              "computed": true,
              "description": "Sets the name of a previously created logical replication slot for a change data capture (CDC) load of the PostgreSQL source instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "redis_settings": {
        "computed": true,
        "description": "Settings in JSON format for the target Redis endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auth_password": {
              "computed": true,
              "description": "The password provided with the auth-role and auth-token options of the AuthType setting for a Redis target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "auth_type": {
              "computed": true,
              "description": "The type of authentication to perform when connecting to a Redis target.",
              "description_kind": "plain",
              "type": "string"
            },
            "auth_user_name": {
              "computed": true,
              "description": "The user name provided with the auth-role option of the AuthType setting for a Redis target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "Transmission Control Protocol (TCP) port for the endpoint.",
              "description_kind": "plain",
              "type": "number"
            },
            "server_name": {
              "computed": true,
              "description": "Fully qualified domain name of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_ca_certificate_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the certificate authority (CA) that DMS uses to connect to your Redis target endpoint.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_security_protocol": {
              "computed": true,
              "description": "The connection to a Redis target endpoint using Transport Layer Security (TLS). Valid values include plaintext and ssl-encryption.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "redshift_settings": {
        "computed": true,
        "description": "Settings in JSON format for the Amazon Redshift endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "accept_any_date": {
              "computed": true,
              "description": "A value that indicates to allow any date format, including invalid formats such as 00/00/00 00:00:00, to be loaded without generating an error. You can choose true or false (the default).",
              "description_kind": "plain",
              "type": "bool"
            },
            "after_connect_script": {
              "computed": true,
              "description": "Code to run after connecting. This parameter should contain the code itself, not the name of a file containing the code.",
              "description_kind": "plain",
              "type": "string"
            },
            "bucket_folder": {
              "computed": true,
              "description": "An S3 folder where the comma-separated-value (.csv) files are stored before being uploaded to the target Redshift cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "bucket_name": {
              "computed": true,
              "description": "The name of the intermediate S3 bucket used to store .csv files before uploading data to Redshift.",
              "description_kind": "plain",
              "type": "string"
            },
            "case_sensitive_names": {
              "computed": true,
              "description": "If Amazon Redshift is configured to support case sensitive schema names, set CaseSensitiveNames to true. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "comp_update": {
              "computed": true,
              "description": "If you set CompUpdate to true Amazon Redshift applies automatic compression if the table is empty.",
              "description_kind": "plain",
              "type": "bool"
            },
            "connection_timeout": {
              "computed": true,
              "description": "A value that sets the amount of time to wait (in milliseconds) before timing out, beginning from when you initially establish a connection.",
              "description_kind": "plain",
              "type": "number"
            },
            "date_format": {
              "computed": true,
              "description": "The date format that you are using.",
              "description_kind": "plain",
              "type": "string"
            },
            "empty_as_null": {
              "computed": true,
              "description": "A value that specifies whether AWS DMS should migrate empty CHAR and VARCHAR fields as NULL. A value of true sets empty CHAR and VARCHAR fields to null. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "encryption_mode": {
              "computed": true,
              "description": "The type of server-side encryption that you want to use for your data.",
              "description_kind": "plain",
              "type": "string"
            },
            "explicit_ids": {
              "computed": true,
              "description": "This setting is only valid for a full-load migration task. Set ExplicitIds to true to have tables with IDENTITY columns override their auto-generated values with explicit values loaded from the source data files used to populate the tables. The default is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "file_transfer_upload_streams": {
              "computed": true,
              "description": "The number of threads used to upload a single file. This parameter accepts a value from 1 through 64. It defaults to 10.",
              "description_kind": "plain",
              "type": "number"
            },
            "load_timeout": {
              "computed": true,
              "description": "The amount of time to wait (in milliseconds) before timing out of operations performed by AWS DMS on a Redshift cluster, such as Redshift COPY, INSERT, DELETE, and UPDATE.",
              "description_kind": "plain",
              "type": "number"
            },
            "map_boolean_as_boolean": {
              "computed": true,
              "description": "When true, lets Redshift migrate the boolean type as boolean. By default, Redshift migrates booleans as varchar(1). You must set this setting on both the source and target endpoints for it to take effect.",
              "description_kind": "plain",
              "type": "bool"
            },
            "max_file_size": {
              "computed": true,
              "description": "The maximum size (in KB) of any .csv file used to load data on an S3 bucket and transfer data to Amazon Redshift. It defaults to 1048576KB (1 GB).",
              "description_kind": "plain",
              "type": "number"
            },
            "remove_quotes": {
              "computed": true,
              "description": "A value that specifies to remove surrounding quotation marks from strings in the incoming data.",
              "description_kind": "plain",
              "type": "bool"
            },
            "replace_chars": {
              "computed": true,
              "description": "A value that specifies to replaces the invalid characters specified in ReplaceInvalidChars, substituting the specified characters instead. The default is \"?\".",
              "description_kind": "plain",
              "type": "string"
            },
            "replace_invalid_chars": {
              "computed": true,
              "description": "A list of characters that you want to replace. Use with ReplaceChars.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret.",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the Amazon Redshift endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            },
            "server_side_encryption_kms_key_id": {
              "computed": true,
              "description": "The AWS KMS key ID. If you are using SSE_KMS for the EncryptionMode, provide this key ID.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_access_role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM role that has access to the Amazon Redshift service. The role must allow the iam:PassRole action.",
              "description_kind": "plain",
              "type": "string"
            },
            "time_format": {
              "computed": true,
              "description": "The time format that you want to use. Valid values are auto (case-sensitive), 'timeformat_string', 'epochsecs', or 'epochmillisecs'.",
              "description_kind": "plain",
              "type": "string"
            },
            "trim_blanks": {
              "computed": true,
              "description": "A value that specifies to remove the trailing white space characters from a VARCHAR string.",
              "description_kind": "plain",
              "type": "bool"
            },
            "truncate_columns": {
              "computed": true,
              "description": "A value that specifies to truncate data in columns to the appropriate number of characters, so that the data fits in the column.",
              "description_kind": "plain",
              "type": "bool"
            },
            "write_buffer_size": {
              "computed": true,
              "description": "The size (in KB) of the in-memory file write buffer used when generating .csv files on the local disk at the DMS replication instance. The default value is 1000 (buffer size is 1000KB).",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_identifier": {
        "computed": true,
        "description": "A display name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target Amazon S3 endpoint",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_column_name": {
              "computed": true,
              "description": "An optional parameter that, when set to true or y, you can use to add column name information to the .csv output file.",
              "description_kind": "plain",
              "type": "bool"
            },
            "add_trailing_padding_character": {
              "computed": true,
              "description": "Use the S3 target endpoint setting AddTrailingPaddingCharacter to add padding on string data. The default value is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "bucket_folder": {
              "computed": true,
              "description": "An optional parameter to set a folder name in the S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "bucket_name": {
              "computed": true,
              "description": "The name of the S3 bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "canned_acl_for_objects": {
              "computed": true,
              "description": "A value that enables AWS DMS to specify a predefined (canned) access control list (ACL) for objects created in an Amazon S3 bucket as .csv or .parquet files.",
              "description_kind": "plain",
              "type": "string"
            },
            "cdc_inserts_and_updates": {
              "computed": true,
              "description": "A value that enables a change data capture (CDC) load to write INSERT and UPDATE operations to .csv or .parquet (columnar storage) output files.",
              "description_kind": "plain",
              "type": "bool"
            },
            "cdc_inserts_only": {
              "computed": true,
              "description": "A value that enables a change data capture (CDC) load to write only INSERT operations to .csv or columnar storage (.parquet) output files. By default (the false setting), the first field in a .csv or .parquet record contains the letter I (INSERT), U (UPDATE), or D (DELETE). These values indicate whether the row was inserted, updated, or deleted at the source database for a CDC load to the target.",
              "description_kind": "plain",
              "type": "bool"
            },
            "cdc_max_batch_interval": {
              "computed": true,
              "description": "Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3.",
              "description_kind": "plain",
              "type": "number"
            },
            "cdc_min_file_size": {
              "computed": true,
              "description": "Minimum file size, defined in kilobytes, to reach for a file output to Amazon S3.",
              "description_kind": "plain",
              "type": "number"
            },
            "cdc_path": {
              "computed": true,
              "description": "Specifies the folder path of CDC files. For an S3 source, this setting is required if a task captures change data; otherwise, it's optional.",
              "description_kind": "plain",
              "type": "string"
            },
            "compression_type": {
              "computed": true,
              "description": "An optional parameter. When set to GZIP it enables the service to compress the target files.",
              "description_kind": "plain",
              "type": "string"
            },
            "csv_delimiter": {
              "computed": true,
              "description": "The delimiter used to separate columns in the .csv file for both source and target. The default is a comma.",
              "description_kind": "plain",
              "type": "string"
            },
            "csv_no_sup_value": {
              "computed": true,
              "description": "This setting only applies if your Amazon S3 output files during a change data capture (CDC) load are written in .csv format.",
              "description_kind": "plain",
              "type": "string"
            },
            "csv_null_value": {
              "computed": true,
              "description": "An optional parameter that specifies how AWS DMS treats null values.",
              "description_kind": "plain",
              "type": "string"
            },
            "csv_row_delimiter": {
              "computed": true,
              "description": "The delimiter used to separate rows in the .csv file for both source and target.",
              "description_kind": "plain",
              "type": "string"
            },
            "data_format": {
              "computed": true,
              "description": "The format of the data that you want to use for output.",
              "description_kind": "plain",
              "type": "string"
            },
            "data_page_size": {
              "computed": true,
              "description": "The size of one data page in bytes. This parameter defaults to 1024 * 1024 bytes (1 MiB). This number is used for .parquet file format only.",
              "description_kind": "plain",
              "type": "number"
            },
            "date_partition_delimiter": {
              "computed": true,
              "description": "Specifies a date separating delimiter to use during folder partitioning. The default value is SLASH. Use this parameter when DatePartitionedEnabled is set to true.",
              "description_kind": "plain",
              "type": "string"
            },
            "date_partition_enabled": {
              "computed": true,
              "description": "When set to true, this parameter partitions S3 bucket folders based on transaction commit dates. The default value is false.",
              "description_kind": "plain",
              "type": "bool"
            },
            "date_partition_sequence": {
              "computed": true,
              "description": "Identifies the sequence of the date format to use during folder partitioning. The default value is YYYYMMDD. Use this parameter when DatePartitionedEnabled is set to true.",
              "description_kind": "plain",
              "type": "string"
            },
            "date_partition_timezone": {
              "computed": true,
              "description": "When creating an S3 target endpoint, set DatePartitionTimezone to convert the current UTC time into a specified time zone.",
              "description_kind": "plain",
              "type": "string"
            },
            "dict_page_size_limit": {
              "computed": true,
              "description": "The maximum size of an encoded dictionary page of a column",
              "description_kind": "plain",
              "type": "number"
            },
            "enable_statistics": {
              "computed": true,
              "description": "A value that enables statistics for Parquet pages and row groups.",
              "description_kind": "plain",
              "type": "bool"
            },
            "encoding_type": {
              "computed": true,
              "description": "The type of encoding that you're using.",
              "description_kind": "plain",
              "type": "string"
            },
            "encryption_mode": {
              "computed": true,
              "description": "The type of server-side encryption that you want to use for your data.",
              "description_kind": "plain",
              "type": "string"
            },
            "expected_bucket_owner": {
              "computed": true,
              "description": "To specify a bucket owner and prevent sniping, you can use the ExpectedBucketOwner endpoint setting.",
              "description_kind": "plain",
              "type": "string"
            },
            "external_table_definition": {
              "computed": true,
              "description": "The external table definition.",
              "description_kind": "plain",
              "type": "string"
            },
            "glue_catalog_generation": {
              "computed": true,
              "description": "When true, allows AWS Glue to catalog your S3 bucket. Creating an AWS Glue catalog lets you use Athena to query your data.",
              "description_kind": "plain",
              "type": "bool"
            },
            "ignore_header_rows": {
              "computed": true,
              "description": "When this value is set to 1, AWS DMS ignores the first row header in a .csv file. A value of 1 turns on the feature; a value of 0 turns off the feature.",
              "description_kind": "plain",
              "type": "number"
            },
            "include_op_for_full_load": {
              "computed": true,
              "description": "A value that enables a full load to write INSERT operations to the comma-separated value (.csv) output files only to indicate how the rows were added to the source database.",
              "description_kind": "plain",
              "type": "bool"
            },
            "max_file_size": {
              "computed": true,
              "description": "A value that specifies the maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load.",
              "description_kind": "plain",
              "type": "number"
            },
            "parquet_timestamp_in_millisecond": {
              "computed": true,
              "description": "A value that specifies the precision of any TIMESTAMP column values that are written to an Amazon S3 object file in .parquet format.",
              "description_kind": "plain",
              "type": "bool"
            },
            "parquet_version": {
              "computed": true,
              "description": "The version of the Apache Parquet format that you want to use: parquet_1_0 (the default) or parquet_2_0.",
              "description_kind": "plain",
              "type": "string"
            },
            "preserve_transactions": {
              "computed": true,
              "description": "If this setting is set to true, AWS DMS saves the transaction order for a change data capture (CDC) load on the Amazon S3 target specified by CdcPath.",
              "description_kind": "plain",
              "type": "bool"
            },
            "rfc_4180": {
              "computed": true,
              "description": "For an S3 source, when this value is set to true or y, each leading double quotation mark has to be followed by an ending double quotation mark.",
              "description_kind": "plain",
              "type": "bool"
            },
            "row_group_length": {
              "computed": true,
              "description": "The number of rows in a row group.",
              "description_kind": "plain",
              "type": "number"
            },
            "server_side_encryption_kms_key_id": {
              "computed": true,
              "description": "If you are using SSE_KMS for the EncryptionMode, provide the AWS KMS key ID. The key that you use needs an attached policy that enables IAM user permissions and allows use of the key.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_access_role_arn": {
              "computed": true,
              "description": "A required parameter that specifies the Amazon Resource Name (ARN) used by the service to access the IAM role.",
              "description_kind": "plain",
              "type": "string"
            },
            "timestamp_column_name": {
              "computed": true,
              "description": "A value that when nonblank causes AWS DMS to add a column with timestamp information to the endpoint data for an Amazon S3 target.",
              "description_kind": "plain",
              "type": "string"
            },
            "use_csv_no_sup_value": {
              "computed": true,
              "description": "This setting applies if the S3 output files during a change data capture (CDC) load are written in .csv format. If this setting is set to true for columns not included in the supplemental log, AWS DMS uses the value specified by CsvNoSupValue. If this setting isn't set or is set to false, AWS DMS uses the null value for these columns.",
              "description_kind": "plain",
              "type": "bool"
            },
            "use_task_start_time_for_full_load_timestamp": {
              "computed": true,
              "description": "When set to true, this parameter uses the task start time as the timestamp column value instead of the time data is written to target",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "server_name": {
        "computed": true,
        "description": "The name of the server where the endpoint database resides.",
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_mode": {
        "computed": true,
        "description": "The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is none.",
        "description_kind": "plain",
        "type": "string"
      },
      "sybase_settings": {
        "computed": true,
        "description": "Settings in JSON format for the source and target SAP ASE endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secrets_manager_access_role_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in SecretsManagerSecret",
              "description_kind": "plain",
              "type": "string"
            },
            "secrets_manager_secret_id": {
              "computed": true,
              "description": "The full ARN, partial ARN, or display name of the SecretsManagerSecret that contains the SAP SAE endpoint connection details.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "One or more tags to be assigned to the endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key is the required name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "username": {
        "computed": true,
        "description": "The user name to be used to log in to the endpoint database.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DMS::Endpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsEndpoint), &result)
	return &result
}
