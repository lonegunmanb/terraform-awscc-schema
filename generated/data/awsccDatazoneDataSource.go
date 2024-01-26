package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneDataSource = `{
  "block": {
    "attributes": {
      "asset_forms_input": {
        "computed": true,
        "description": "The metadata forms that are to be attached to the assets that this data source works with.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description": "The content of the metadata form.",
              "description_kind": "plain",
              "type": "string"
            },
            "form_name": {
              "computed": true,
              "description": "The name of the metadata form.",
              "description_kind": "plain",
              "type": "string"
            },
            "type_identifier": {
              "computed": true,
              "description": "The ID of the metadata form type.",
              "description_kind": "plain",
              "type": "string"
            },
            "type_revision": {
              "computed": true,
              "description": "The revision of the metadata form type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "configuration": {
        "computed": true,
        "description": "Configuration of the data source. It can be set to either glueRunConfiguration or redshiftRunConfiguration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "glue_run_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_access_role": {
                    "computed": true,
                    "description": "The data access role included in the configuration details of the AWS Glue data source.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "relational_filter_configurations": {
                    "computed": true,
                    "description": "The relational filter configurations included in the configuration details of the AWS Glue data source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_name": {
                          "computed": true,
                          "description": "The database name specified in the relational filter configuration for the data source.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "filter_expressions": {
                          "computed": true,
                          "description": "The filter expressions specified in the relational filter configuration for the data source.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "type": {
                                "computed": true,
                                "description": "The search filter expression type.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "schema_name": {
                          "computed": true,
                          "description": "The schema name specified in the relational filter configuration for the data source.",
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
            "redshift_run_configuration": {
              "computed": true,
              "description": "The configuration details of the Amazon Redshift data source.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_access_role": {
                    "computed": true,
                    "description": "The data access role included in the configuration details of the Amazon Redshift data source.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "redshift_credential_configuration": {
                    "computed": true,
                    "description": "The details of the credentials required to access an Amazon Redshift cluster.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "secret_manager_arn": {
                          "computed": true,
                          "description": "The ARN of a secret manager for an Amazon Redshift cluster.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "redshift_storage": {
                    "computed": true,
                    "description": "The details of the Amazon Redshift storage as part of the configuration of an Amazon Redshift data source run.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "redshift_cluster_source": {
                          "computed": true,
                          "description": "The name of an Amazon Redshift cluster.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cluster_name": {
                                "computed": true,
                                "description": "The name of an Amazon Redshift cluster.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "redshift_serverless_source": {
                          "computed": true,
                          "description": "The details of the Amazon Redshift Serverless workgroup storage.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "workgroup_name": {
                                "computed": true,
                                "description": "The name of the Amazon Redshift Serverless workgroup.",
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
                  "relational_filter_configurations": {
                    "computed": true,
                    "description": "The relational filter configurations included in the configuration details of the Amazon Redshift data source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_name": {
                          "computed": true,
                          "description": "The database name specified in the relational filter configuration for the data source.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "filter_expressions": {
                          "computed": true,
                          "description": "The filter expressions specified in the relational filter configuration for the data source.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "type": {
                                "computed": true,
                                "description": "The search filter expression type.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "schema_name": {
                          "computed": true,
                          "description": "The schema name specified in the relational filter configuration for the data source.",
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
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the data source was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain where the data source is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone domain where the data source is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_setting": {
        "computed": true,
        "description": "Specifies whether the data source is enabled.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_identifier": {
        "computed": true,
        "description": "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_run_asset_count": {
        "computed": true,
        "description": "The number of assets created by the data source during its last run.",
        "description_kind": "plain",
        "type": "number"
      },
      "last_run_at": {
        "computed": true,
        "description": "The timestamp that specifies when the data source was last run.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_run_status": {
        "computed": true,
        "description": "The status of the last run of this data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone project to which the data source is added.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_identifier": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone project in which you want to add the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "publish_on_import": {
        "computed": true,
        "description": "Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.",
        "description_kind": "plain",
        "type": "bool"
      },
      "recommendation": {
        "computed": true,
        "description": "Specifies whether the business name generation is to be enabled for this data source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_business_name_generation": {
              "computed": true,
              "description": "Specifies whether automatic business name generation is to be enabled or not as part of the recommendation configuration.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "schedule": {
        "computed": true,
        "description": "The schedule of the data source runs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule": {
              "computed": true,
              "description": "The schedule of the data source runs.",
              "description_kind": "plain",
              "type": "string"
            },
            "timezone": {
              "computed": true,
              "description": "The timezone of the data source run.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp of when this data source was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::DataSource",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneDataSource), &result)
	return &result
}
