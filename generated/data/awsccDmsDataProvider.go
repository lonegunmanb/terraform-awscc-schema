package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsDataProvider = `{
  "block": {
    "attributes": {
      "data_provider_arn": {
        "computed": true,
        "description": "The data provider ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_provider_creation_time": {
        "computed": true,
        "description": "The data provider creation time.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_provider_identifier": {
        "computed": true,
        "description": "The property describes an identifier for the data provider. It is used for describing/deleting/modifying can be name/arn",
        "description_kind": "plain",
        "type": "string"
      },
      "data_provider_name": {
        "computed": true,
        "description": "The property describes a name to identify the data provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The optional description of the data provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description": "The property describes a data engine for the data provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "exact_settings": {
        "computed": true,
        "description": "The property describes the exact settings which can be modified",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "settings": {
        "computed": true,
        "description": "The property identifies the exact type of settings for the data provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "doc_db_settings": {
              "computed": true,
              "description": "DocDbSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "maria_db_settings": {
              "computed": true,
              "description": "MariaDbSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "microsoft_sql_server_settings": {
              "computed": true,
              "description": "MicrosoftSqlServerSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "mongo_db_settings": {
              "computed": true,
              "description": "MongoDbSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "auth_mechanism": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "auth_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "auth_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "my_sql_settings": {
              "computed": true,
              "description": "MySqlSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "oracle_settings": {
              "computed": true,
              "description": "OracleSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "asm_server": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "secrets_manager_oracle_asm_access_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secrets_manager_oracle_asm_secret_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secrets_manager_security_db_encryption_access_role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secrets_manager_security_db_encryption_secret_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "postgre_sql_settings": {
              "computed": true,
              "description": "PostgreSqlSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ssl_mode": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "redshift_settings": {
              "computed": true,
              "description": "RedshiftSettings property identifier.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "server_name": {
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
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DMS::DataProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsDataProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsDataProvider), &result)
	return &result
}
