package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueCatalog = `{
  "block": {
    "attributes": {
      "allow_full_table_external_data_access": {
        "computed": true,
        "description": "Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "catalog_id": {
        "computed": true,
        "description": "The ID of the catalog.",
        "description_kind": "plain",
        "type": "string"
      },
      "catalog_properties": {
        "computed": true,
        "description": "A structure that specifies data lake access properties and other custom properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_properties": {
              "computed": true,
              "description": "Additional key-value properties for the catalog.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "data_lake_access_properties": {
              "computed": true,
              "description": "Data lake access properties for the catalog.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allow_full_table_external_data_access": {
                    "computed": true,
                    "description": "Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "catalog_type": {
                    "computed": true,
                    "description": "Specifies a federated catalog type for the native catalog resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_lake_access": {
                    "computed": true,
                    "description": "Turns on or off data lake access for Apache Spark applications that access Amazon Redshift databases in the Data Catalog from any non-Redshift engine.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "data_transfer_role": {
                    "computed": true,
                    "description": "A role that will be assumed by Glue for transferring data into/out of the staging bucket during a query.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key": {
                    "computed": true,
                    "description": "An encryption key that will be used for the staging bucket that will be created along with the catalog.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "managed_workgroup_name": {
                    "computed": true,
                    "description": "The name of the managed workgroup associated with the catalog.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "managed_workgroup_status": {
                    "computed": true,
                    "description": "The status of the managed workgroup.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "redshift_database_name": {
                    "computed": true,
                    "description": "The name of the Redshift database.",
                    "description_kind": "plain",
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
      "create_database_default_permissions": {
        "computed": true,
        "description": "An array of PrincipalPermissions objects for default database permissions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "permissions": {
              "computed": true,
              "description": "The permissions that are granted to the principal.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description": "The Lake Formation principal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_lake_principal_identifier": {
                    "computed": true,
                    "description": "An identifier for the Lake Formation principal.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "create_table_default_permissions": {
        "computed": true,
        "description": "An array of PrincipalPermissions objects for default table permissions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "permissions": {
              "computed": true,
              "description": "The permissions that are granted to the principal.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description": "The Lake Formation principal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_lake_principal_identifier": {
                    "computed": true,
                    "description": "An identifier for the Lake Formation principal.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "create_time": {
        "computed": true,
        "description": "The time at which the catalog was created.",
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "A description of the catalog.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "federated_catalog": {
        "computed": true,
        "description": "A FederatedCatalog structure that references an entity outside the Glue Data Catalog.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_name": {
              "computed": true,
              "description": "The name of the connection to an external data source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identifier": {
              "computed": true,
              "description": "A unique identifier for the federated catalog.",
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
      "name": {
        "description": "The name of the catalog to create.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "overwrite_child_resource_permissions_with_default": {
        "computed": true,
        "description": "Specifies whether to overwrite child resource permissions with the default permissions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "A map of key-value pairs that define parameters and properties of the catalog.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the catalog.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_redshift_catalog": {
        "computed": true,
        "description": "A structure that describes a target catalog for resource linking.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the catalog resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the catalog was last updated.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Creates a catalog in the Glue Data Catalog.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueCatalog), &result)
	return &result
}
