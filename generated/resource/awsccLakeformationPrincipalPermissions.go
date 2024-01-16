package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "permissions_with_grant_option": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "principal": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_lake_principal_identifier": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "principal_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_cells_filter": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "data_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "resource_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "database": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "lf_tag": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tag_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tag_values": {
                    "description_kind": "plain",
                    "required": true,
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
            "lf_tag_policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "expression": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tag_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tag_values": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
                  },
                  "resource_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_wildcard": {
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
            "table_with_columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "column_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
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
          "nesting_mode": "single"
        },
        "required": true
      },
      "resource_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A resource schema representing a Lake Formation Permission.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLakeformationPrincipalPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationPrincipalPermissions), &result)
	return &result
}
