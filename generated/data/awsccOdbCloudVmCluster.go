package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOdbCloudVmCluster = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructure_id": {
        "computed": true,
        "description": "The unique identifier of the Exadata infrastructure that this VM cluster belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_vm_cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_vm_cluster_id": {
        "computed": true,
        "description": "The unique identifier of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The name of the Grid Infrastructure (GI) cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "compute_model": {
        "computed": true,
        "description": "The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled.",
        "description_kind": "plain",
        "type": "string"
      },
      "cpu_core_count": {
        "computed": true,
        "description": "The number of CPU cores enabled on the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "data_collection_options": {
        "computed": true,
        "description": "The set of diagnostic collection options enabled for the VM cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "is_diagnostics_events_enabled": {
              "computed": true,
              "description": "Indicates whether diagnostic collection is enabled for the VM cluster.",
              "description_kind": "plain",
              "type": "bool"
            },
            "is_health_monitoring_enabled": {
              "computed": true,
              "description": "Indicates whether health monitoring is enabled for the VM cluster.",
              "description_kind": "plain",
              "type": "bool"
            },
            "is_incident_logs_enabled": {
              "computed": true,
              "description": "Indicates whether incident logs are enabled for the cloud VM cluster.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "data_storage_size_in_t_bs": {
        "computed": true,
        "description": "The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_node_storage_size_in_g_bs": {
        "computed": true,
        "description": "The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "db_servers": {
        "computed": true,
        "description": "The list of database servers for the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "disk_redundancy": {
        "computed": true,
        "description": "The type of redundancy configured for the VM cluster. NORMAL is 2-way redundancy. HIGH is 3-way redundancy.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The user-friendly name for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "gi_version": {
        "computed": true,
        "description": "The software version of the Oracle Grid Infrastructure (GI) for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "hostname": {
        "computed": true,
        "description": "The host name for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_local_backup_enabled": {
        "computed": true,
        "description": "Indicates whether database backups to local Exadata storage is enabled for the VM cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "is_sparse_diskgroup_enabled": {
        "computed": true,
        "description": "Indicates whether the VM cluster is configured with a sparse disk group.",
        "description_kind": "plain",
        "type": "bool"
      },
      "license_model": {
        "computed": true,
        "description": "The Oracle license model applied to the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "listener_port": {
        "computed": true,
        "description": "The port number configured for the listener on the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "memory_size_in_g_bs": {
        "computed": true,
        "description": "The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "node_count": {
        "computed": true,
        "description": "The number of nodes in the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description": "The HTTPS link to the VM cluster in OCI.",
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description": "The OCID of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "The unique identifier of the ODB network for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_dns_name": {
        "computed": true,
        "description": "The FQDN of the DNS record for the Single Client Access Name (SCAN) IP addresses that are associated with the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scan_ip_ids": {
        "computed": true,
        "description": "The OCID of the SCAN IP addresses that are associated with the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "scan_listener_port_tcp": {
        "computed": true,
        "description": "Property description not available.",
        "description_kind": "plain",
        "type": "number"
      },
      "shape": {
        "computed": true,
        "description": "The hardware model name of the Exadata infrastructure that's running the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "ssh_public_keys": {
        "computed": true,
        "description": "The public key portion of one or more key pairs used for SSH access to the VM cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "storage_size_in_g_bs": {
        "computed": true,
        "description": "The amount of local node storage, in gigabytes (GB), that's allocated to the VM cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "system_version": {
        "computed": true,
        "description": "The operating system version of the image chosen for the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the Vm Cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and \".",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "time_zone": {
        "computed": true,
        "description": "The time zone of the VM cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "vip_ids": {
        "computed": true,
        "description": "The virtual IP (VIP) addresses that are associated with the VM cluster. Oracle's Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the VM cluster to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::ODB::CloudVmCluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOdbCloudVmClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOdbCloudVmCluster), &result)
	return &result
}
