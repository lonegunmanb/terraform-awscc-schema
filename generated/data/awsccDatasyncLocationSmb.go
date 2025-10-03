package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationSmb = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "authentication_type": {
        "computed": true,
        "description": "The authentication mode used to determine identity of user.",
        "description_kind": "plain",
        "type": "string"
      },
      "cmk_secret_config": {
        "computed": true,
        "description": "Specifies configuration information for a DataSync-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "Specifies the ARN for the customer-managed AWS KMS key used to encrypt the secret specified for SecretArn. DataSync provides this key to AWS Secrets Manager.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "Specifies the ARN for an AWS Secrets Manager secret, managed by DataSync.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "custom_secret_config": {
        "computed": true,
        "description": "Specifies configuration information for a customer-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secret_access_role_arn": {
              "computed": true,
              "description": "Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for SecretArn.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "Specifies the ARN for a customer created AWS Secrets Manager secret.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "dns_ip_addresses": {
        "computed": true,
        "description": "Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to. This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "domain": {
        "computed": true,
        "description": "The name of the Windows domain that the SMB server belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kerberos_keytab": {
        "computed": true,
        "description": "The Base64 string representation of the Keytab file. Specifies your Kerberos key table (keytab) file, which includes mappings between your service principal name (SPN) and encryption keys. To avoid task execution errors, make sure that the SPN in the keytab file matches exactly what you specify for KerberosPrincipal and in your krb5.conf file.",
        "description_kind": "plain",
        "type": "string"
      },
      "kerberos_krb_5_conf": {
        "computed": true,
        "description": "The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket. Specifies a Kerberos configuration file (krb5.conf) that defines your Kerberos realm configuration. To avoid task execution errors, make sure that the service principal name (SPN) in the krb5.conf file matches exactly what you specify for KerberosPrincipal and in your keytab file.",
        "description_kind": "plain",
        "type": "string"
      },
      "kerberos_principal": {
        "computed": true,
        "description": "Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server. SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the SMB location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the SMB location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "managed_secret_config": {
        "computed": true,
        "description": "Specifies configuration information for a DataSync-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secret_arn": {
              "computed": true,
              "description": "Specifies the ARN for an AWS Secrets Manager secret.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "mount_options": {
        "computed": true,
        "description": "The mount options used by DataSync to access the SMB server.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "version": {
              "computed": true,
              "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "password": {
        "computed": true,
        "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
        "description_kind": "plain",
        "type": "string"
      },
      "server_hostname": {
        "computed": true,
        "description": "The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server.",
        "description_kind": "plain",
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination",
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
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "user": {
        "computed": true,
        "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataSync::LocationSMB",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationSmbSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationSmb), &result)
	return &result
}
