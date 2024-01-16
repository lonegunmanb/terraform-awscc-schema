package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLakeformationPrincipalPermissions = `{
  "block": {
    "attributes": {
      "catalog": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "permissions_with_grant_option": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "principal": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_lake_principal_identifier": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "principal_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "data_cells_filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_catalog_id": {
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
            "data_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "database": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "lf_tag": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tag_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tag_values": {
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
            "lf_tag_policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tag_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tag_values": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "resource_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_wildcard": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "table_with_columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "column_wildcard": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "excluded_column_names": {
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
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
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
      "resource_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::LakeFormation::PrincipalPermissions",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLakeformationPrincipalPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationPrincipalPermissions), &result)
	return &result
}
