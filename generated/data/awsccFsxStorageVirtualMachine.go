package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxStorageVirtualMachine = `{
  "block": {
    "attributes": {
      "active_directory_configuration": {
        "computed": true,
        "description": "Describes the Microsoft Active Directory configuration to which the SVM is joined, if applicable.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "net_bios_name": {
              "computed": true,
              "description": "The NetBIOS name of the Active Directory computer object that will be created for your SVM. Minimum values: 1. Maximum values: 15. Pattern: ` + "`" + `^[^\\u0000\\u0085\\u2028\\u2029\\r\\n]{1,255}$` + "`" + `",
              "description_kind": "plain",
              "type": "string"
            },
            "self_managed_active_directory_configuration": {
              "computed": true,
              "description": "The configuration that Amazon FSx uses to join the ONTAP storage virtual machine (SVM) to your self-managed (including on-premises) Microsoft Active Directory directory.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dns_ips": {
                    "computed": true,
                    "description": "A list of IP addresses of DNS servers or domain controllers in the self-managed AD directory. Minimum: 1. Maximum: 3.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "domain_join_service_account_secret": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret containing the self-managed Active Directory domain join service account credentials. Minimum: 64. Maximum: 1024.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "domain_name": {
                    "computed": true,
                    "description": "The fully qualified domain name of the self-managed AD directory, such as corp.example.com. Minimum: 1. Maximum: 255. Pattern: ` + "`" + `^[^\\u0000\\u0085\\u2028\\u2029\\r\\n]{1,255}$` + "`" + `",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "file_system_administrators_group": {
                    "computed": true,
                    "description": "The name of the domain group whose members are granted administrative privileges for the file system. Minimum: 1. Maximum: 256. Pattern: ` + "`" + `^[^\\u0000\\u0085\\u2028\\u2029\\r\\n]{1,256}$` + "`" + `",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "organizational_unit_distinguished_name": {
                    "computed": true,
                    "description": "The fully qualified distinguished name of the organizational unit within your self-managed AD directory. Minimum: 1. Maximum: 2000. Pattern: ` + "`" + `^[^\\u0000\\u0085\\u2028\\u2029\\r\\n]{1,2000}$` + "`" + `",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "password": {
                    "computed": true,
                    "description": "The password for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain. Minimum: 1. Maximum: 256. Pattern: ` + "`" + `^.{1,256}$` + "`" + `",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_name": {
                    "computed": true,
                    "description": "The user name for the service account on your self-managed AD domain that Amazon FSx will use to join to your AD domain. Minimum: 1. Maximum: 256. Pattern: ` + "`" + `^[^\\u0000\\u0085\\u2028\\u2029\\r\\n]{1,256}$` + "`" + `",
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
      "file_system_id": {
        "computed": true,
        "description": "Specifies the FSx for ONTAP file system on which to create the SVM.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the SVM. Pattern: ` + "`" + `^[^\\u0000\\u0085\\u2028\\u2029\\r\\n]{1,47}$` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for a given resource. Min Length: 8, Max Length 512",
        "description_kind": "plain",
        "type": "string"
      },
      "root_volume_security_style": {
        "computed": true,
        "description": "The security style of the root volume of the SVM. Allowed values: UNIX | NTFS | MIXED.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_virtual_machine_id": {
        "computed": true,
        "description": "The SVM's system generated unique ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "svm_admin_password": {
        "computed": true,
        "description": "Specifies the password to use when logging on to the SVM using a secure shell (SSH) connection to the SVM's management endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of Tag values, with a maximum of 50 elements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "uuid": {
        "computed": true,
        "description": "The SVM's UUID (universally unique identifier).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::FSx::StorageVirtualMachine",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFsxStorageVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxStorageVirtualMachine), &result)
	return &result
}
