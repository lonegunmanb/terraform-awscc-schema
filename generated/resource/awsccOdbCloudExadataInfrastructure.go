package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOdbCloudExadataInfrastructure = `{
  "block": {
    "attributes": {
      "activated_storage_count": {
        "computed": true,
        "description": "The number of storage servers requested for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "additional_storage_count": {
        "computed": true,
        "description": "The number of storage servers requested for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "availability_zone": {
        "computed": true,
        "description": "The name of the Availability Zone (AZ) where the Exadata infrastructure is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The AZ ID of the AZ where the Exadata infrastructure is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "available_storage_size_in_g_bs": {
        "computed": true,
        "description": "The amount of available storage, in gigabytes (GB), for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "cloud_exadata_infrastructure_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_exadata_infrastructure_id": {
        "computed": true,
        "description": "The unique identifier for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_count": {
        "computed": true,
        "description": "The number of database servers for the Exadata infrastructure.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "compute_model": {
        "computed": true,
        "description": "The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores that are allocated to the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "customer_contacts_to_send_to_oci": {
        "computed": true,
        "description": "The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "The email address of the contact.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "data_storage_size_in_t_bs": {
        "computed": true,
        "description": "The size of the Exadata infrastructure's data disk group, in terabytes (TB).",
        "description_kind": "plain",
        "type": "number"
      },
      "database_server_type": {
        "computed": true,
        "description": "The database server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_node_storage_size_in_g_bs": {
        "computed": true,
        "description": "The size of the Exadata infrastructure's local node storage, in gigabytes (GB).",
        "description_kind": "plain",
        "type": "number"
      },
      "db_server_ids": {
        "computed": true,
        "description": "The list of database server identifiers for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "db_server_version": {
        "computed": true,
        "description": "The software version of the database servers (dom0) in the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The user-friendly name for the Exadata infrastructure.",
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
      "max_cpu_count": {
        "computed": true,
        "description": "The total number of CPU cores available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_data_storage_in_t_bs": {
        "computed": true,
        "description": "The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_db_node_storage_size_in_g_bs": {
        "computed": true,
        "description": "The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_memory_in_g_bs": {
        "computed": true,
        "description": "The total amount of memory, in gigabytes (GB), that's available on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_g_bs": {
        "computed": true,
        "description": "The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor for the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The HTTPS link to the Exadata infrastructure in OCI.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "shape": {
        "computed": true,
        "description": "The model name of the Exadata infrastructure.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_count": {
        "computed": true,
        "description": "The number of storage servers that are activated for the Exadata infrastructure.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_server_type": {
        "computed": true,
        "description": "The storage server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_server_version": {
        "computed": true,
        "description": "The software version of the storage servers on the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the Exadata Infrastructure.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and \".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "total_storage_size_in_g_bs": {
        "computed": true,
        "description": "The total amount of storage, in gigabytes (GB), on the the Exadata infrastructure.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "The AWS::ODB::CloudExadataInfrastructure resource creates an Exadata Infrastructure",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOdbCloudExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOdbCloudExadataInfrastructure), &result)
	return &result
}
