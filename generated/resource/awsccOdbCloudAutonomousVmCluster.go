package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOdbCloudAutonomousVmCluster = `{
  "block": {
    "attributes": {
      "autonomous_data_storage_percentage": {
        "computed": true,
        "description": "The percentage of data storage currently in use for Autonomous Databases in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "autonomous_data_storage_size_in_t_bs": {
        "computed": true,
        "description": "The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "available_autonomous_data_storage_size_in_t_bs": {
        "computed": true,
        "description": "The available data storage space for Autonomous Databases in the Autonomous VM cluster, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "available_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs that you can create with the currently available storage.",
        "description_kind": "plain",
        "type": "number"
      },
      "available_cpus": {
        "computed": true,
        "description": "The number of CPU cores available for allocation to Autonomous Databases.",
        "description_kind": "plain",
        "type": "number"
      },
      "cloud_autonomous_vm_cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_autonomous_vm_cluster_id": {
        "computed": true,
        "description": "The unique identifier of the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_exadata_infrastructure_id": {
        "computed": true,
        "description": "The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compute_model": {
        "computed": true,
        "description": "The compute model of the Autonomous VM cluster: ECPU or OCPU.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "The total number of CPU cores in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "cpu_core_count_per_node": {
        "computed": true,
        "description": "The number of CPU cores enabled per node in the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cpu_percentage": {
        "computed": true,
        "description": "The percentage of total CPU cores currently in use in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_g_bs": {
        "computed": true,
        "description": "The total data storage allocated to the Autonomous VM cluster, in GB.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_storage_size_in_t_bs": {
        "computed": true,
        "description": "The total data storage allocated to the Autonomous VM cluster, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_node_storage_size_in_g_bs": {
        "computed": true,
        "description": "The local node storage allocated to the Autonomous VM cluster, in gigabytes (GB).",
        "description_kind": "plain",
        "type": "number"
      },
      "db_servers": {
        "computed": true,
        "description": "The list of database servers associated with the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "The user-provided description of the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain name for the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "exadata_storage_in_t_bs_lowest_scaled_value": {
        "computed": true,
        "description": "The minimum value to which you can scale down the Exadata storage, in TB.",
        "description_kind": "plain",
        "type": "number"
      },
      "hostname": {
        "computed": true,
        "description": "The hostname for the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_mtls_enabled_vm_cluster": {
        "computed": true,
        "description": "Indicates whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "license_model": {
        "computed": true,
        "description": "The Oracle license model that applies to the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description": "The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "days_of_week": {
              "computed": true,
              "description": "The days of the week when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "hours_of_day": {
              "computed": true,
              "description": "The hours of the day when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            },
            "lead_time_in_weeks": {
              "computed": true,
              "description": "The lead time in weeks before the maintenance window.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "months": {
              "computed": true,
              "description": "The months when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "preference": {
              "computed": true,
              "description": "The preference for the maintenance window scheduling.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weeks_of_month": {
              "computed": true,
              "description": "The weeks of the month when maintenance can be performed.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "max_acds_lowest_scaled_value": {
        "computed": true,
        "description": "The minimum value to which you can scale down the maximum number of Autonomous CDBs.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_per_oracle_compute_unit_in_g_bs": {
        "computed": true,
        "description": "The amount of memory allocated per Oracle Compute Unit, in GB.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "memory_size_in_g_bs": {
        "computed": true,
        "description": "The total amount of memory allocated to the Autonomous VM cluster, in gigabytes (GB).",
        "description_kind": "plain",
        "type": "number"
      },
      "node_count": {
        "computed": true,
        "description": "The number of database server nodes in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "non_provisionable_autonomous_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs that can't be provisioned because of resource constraints.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor associated with this Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The URL for accessing the OCI console page for this Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The Oracle Cloud Identifier (OCID) of the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "The unique identifier of the ODB network associated with this Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisionable_autonomous_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs that can be provisioned in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "provisioned_autonomous_container_databases": {
        "computed": true,
        "description": "The number of Autonomous CDBs currently provisioned in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "provisioned_cpus": {
        "computed": true,
        "description": "The number of CPU cores currently provisioned in the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "reclaimable_cpus": {
        "computed": true,
        "description": "The number of CPU cores that can be reclaimed from terminated or scaled-down Autonomous Databases.",
        "description_kind": "plain",
        "type": "number"
      },
      "reserved_cpus": {
        "computed": true,
        "description": "The number of CPU cores reserved for system operations and redundancy.",
        "description_kind": "plain",
        "type": "number"
      },
      "scan_listener_port_non_tls": {
        "computed": true,
        "description": "The SCAN listener port for non-TLS (TCP) protocol. The default is 1521.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scan_listener_port_tls": {
        "computed": true,
        "description": "The SCAN listener port for TLS (TCP) protocol. The default is 2484.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shape": {
        "computed": true,
        "description": "The shape of the Exadata infrastructure for the Autonomous VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the Autonomous VM cluster.",
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
      "time_zone": {
        "computed": true,
        "description": "The time zone of the Autonomous VM cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "total_container_databases": {
        "computed": true,
        "description": "The total number of Autonomous Container Databases that can be created with the allocated local storage.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "The AWS::ODB::CloudAutonomousVmCluster resource creates a Cloud Autonomous VM Cluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOdbCloudAutonomousVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOdbCloudAutonomousVmCluster), &result)
	return &result
}
