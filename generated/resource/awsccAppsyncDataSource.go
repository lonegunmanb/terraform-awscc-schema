package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncDataSource = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "Unique AWS AppSync GraphQL API identifier where this data source will be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "data_source_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the API key, such as arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the data source.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dynamo_db_config": {
        "computed": true,
        "description": "AWS Region and TableName for an Amazon DynamoDB table in your account.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_region": {
              "computed": true,
              "description": "The AWS Region.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delta_sync_config": {
              "computed": true,
              "description": "The DeltaSyncConfig for a versioned datasource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "base_table_ttl": {
                    "computed": true,
                    "description": "The number of minutes that an Item is stored in the data source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "delta_sync_table_name": {
                    "computed": true,
                    "description": "The Delta Sync table name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "delta_sync_table_ttl": {
                    "computed": true,
                    "description": "The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "table_name": {
              "computed": true,
              "description": "The table name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_caller_credentials": {
              "computed": true,
              "description": "Set to TRUE to use AWS Identity and Access Management with this data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "versioned": {
              "computed": true,
              "description": "Set to TRUE to use Conflict Detection and Resolution with this data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "elasticsearch_config": {
        "computed": true,
        "description": "AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.\nAs of September 2021, Amazon Elasticsearch Service is Amazon OpenSearch Service. This property is deprecated. For new data sources, use OpenSearchServiceConfig to specify an OpenSearch Service data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_region": {
              "computed": true,
              "description": "The AWS Region.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "endpoint": {
              "computed": true,
              "description": "The endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "event_bridge_config": {
        "computed": true,
        "description": "ARN for the EventBridge bus.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_bus_arn": {
              "computed": true,
              "description": "ARN for the EventBridge bus.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "http_config": {
        "computed": true,
        "description": "Endpoints for an HTTP data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorization_config": {
              "computed": true,
              "description": "The authorization configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_type": {
                    "computed": true,
                    "description": "The authorization type that the HTTP endpoint requires.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "aws_iam_config": {
                    "computed": true,
                    "description": "The AWS Identity and Access Management settings.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "signing_region": {
                          "computed": true,
                          "description": "The signing Region for AWS Identity and Access Management authorization.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "signing_service_name": {
                          "computed": true,
                          "description": "The signing service name for AWS Identity and Access Management authorization.",
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
            "endpoint": {
              "computed": true,
              "description": "The endpoint.",
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
      "lambda_config": {
        "computed": true,
        "description": "An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda function that exists in the current account or in another account.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lambda_function_arn": {
              "computed": true,
              "description": "The ARN for the Lambda function.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "metrics_config": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Friendly name for you to identify your AppSync data source after creation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "open_search_service_config": {
        "computed": true,
        "description": "AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_region": {
              "computed": true,
              "description": "The AWS Region.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "endpoint": {
              "computed": true,
              "description": "The endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "relational_database_config": {
        "computed": true,
        "description": "Relational Database configuration of the relational database data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rds_http_endpoint_config": {
              "computed": true,
              "description": "Information about the Amazon RDS resource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_region": {
                    "computed": true,
                    "description": "AWS Region for RDS HTTP endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "aws_secret_store_arn": {
                    "computed": true,
                    "description": "The ARN for database credentials stored in AWS Secrets Manager.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "Logical database name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "db_cluster_identifier": {
                    "computed": true,
                    "description": "Amazon RDS cluster Amazon Resource Name (ARN).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schema": {
                    "computed": true,
                    "description": "Logical schema name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "relational_database_source_type": {
              "computed": true,
              "description": "The type of relational data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "service_role_arn": {
        "computed": true,
        "description": "The AWS Identity and Access Management service role ARN for the data source. The system assumes this role when accessing the data source.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of the data source.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppSync::DataSource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncDataSource), &result)
	return &result
}
