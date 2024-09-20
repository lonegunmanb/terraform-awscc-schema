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
        "description": "The identifier for the GLUDC. By default, the account ID. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.",
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
        "description": "The permissions granted or revoked.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "permissions_with_grant_option": {
        "description": "Indicates the ability to grant permissions (as a subset of permissions granted).",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "principal": {
        "description": "The principal to be granted a permission.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_lake_principal_identifier": {
              "computed": true,
              "description": "An identifier for the LFlong principal.",
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
        "description": "The resource to be granted or revoked permissions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog": {
              "computed": true,
              "description": "The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your LFlong environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_cells_filter": {
              "computed": true,
              "description": "A data cell filter.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "database_name": {
                    "computed": true,
                    "description": "A database in the GLUDC.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name given by the user to the data filter cell.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_catalog_id": {
                    "computed": true,
                    "description": "The ID of the catalog to which the table belongs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "The name of the table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "data_location": {
              "computed": true,
              "description": "The location of an Amazon S3 path where permissions are granted or revoked.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The identifier for the GLUDC where the location is registered with LFlong.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) that uniquely identifies the data location resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "database": {
              "computed": true,
              "description": "The database for the resource. Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The identifier for the Data Catalog. By default, it is the account ID of the caller.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the database resource. Unique to the Data Catalog.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "lf_tag": {
              "computed": true,
              "description": "The LF-tag key and values attached to a resource.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The identifier for the GLUDC where the location is registered with GLUDC.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_key": {
                    "computed": true,
                    "description": "The key-name for the LF-tag.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_values": {
                    "computed": true,
                    "description": "A list of possible values for the corresponding ` + "`" + `` + "`" + `TagKey` + "`" + `` + "`" + ` of an LF-tag key-value pair.",
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
            "lf_tag_policy": {
              "computed": true,
              "description": "A list of LF-tag conditions that define a resource's LF-tag policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The identifier for the GLUDC. The GLUDC is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your LFlong environment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expression": {
                    "computed": true,
                    "description": "A list of LF-tag conditions that apply to the resource's LF-tag policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tag_key": {
                          "computed": true,
                          "description": "The key-name for the LF-tag.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tag_values": {
                          "computed": true,
                          "description": "A list of possible values of the corresponding ` + "`" + `` + "`" + `TagKey` + "`" + `` + "`" + ` of an LF-tag key-value pair.",
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
                    "optional": true
                  },
                  "resource_type": {
                    "computed": true,
                    "description": "The resource type for which the LF-tag policy applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "table": {
              "computed": true,
              "description": "The table for the resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The identifier for the Data Catalog. By default, it is the account ID of the caller.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "The name of the database for the table. Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the table.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_wildcard": {
                    "computed": true,
                    "description": "A wildcard object representing every table under a database.\n At least one of ` + "`" + `` + "`" + `TableResource$Name` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `TableResource$TableWildcard` + "`" + `` + "`" + ` is required.",
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
              "description": "The table with columns for the resource. A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The identifier for the GLUDC where the location is registered with LFlong.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "column_names": {
                    "computed": true,
                    "description": "The list of column names for the table. At least one of ` + "`" + `` + "`" + `ColumnNames` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ColumnWildcard` + "`" + `` + "`" + ` is required.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "column_wildcard": {
                    "computed": true,
                    "description": "A wildcard specified by a ` + "`" + `` + "`" + `ColumnWildcard` + "`" + `` + "`" + ` object. At least one of ` + "`" + `` + "`" + `ColumnNames` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ColumnWildcard` + "`" + `` + "`" + ` is required.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "excluded_column_names": {
                          "computed": true,
                          "description": "Excludes column names. Any column with this name will be excluded.",
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
                    "computed": true,
                    "description": "The name of the database for the table with columns resource. Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the table resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.",
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
        "required": true
      },
      "resource_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::LakeFormation::PrincipalPermissions` + "`" + `` + "`" + ` resource represents the permissions that a principal has on a GLUDC resource (such as GLUlong databases or GLUlong tables). When you create a ` + "`" + `` + "`" + `PrincipalPermissions` + "`" + `` + "`" + ` resource, the permissions are granted via the LFlong ` + "`" + `` + "`" + `GrantPermissions` + "`" + `` + "`" + ` API operation. When you delete a ` + "`" + `` + "`" + `PrincipalPermissions` + "`" + `` + "`" + ` resource, the permissions on principal-resource pair are revoked via the LFlong ` + "`" + `` + "`" + `RevokePermissions` + "`" + `` + "`" + ` API operation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLakeformationPrincipalPermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationPrincipalPermissions), &result)
	return &result
}
