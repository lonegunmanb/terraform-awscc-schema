package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectoryserviceSimpleAd = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The alias for a directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_alias": {
        "computed": true,
        "description": "The name of the configuration set.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "Description for the directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
        "computed": true,
        "description": "The unique identifier for a directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_ip_addresses": {
        "computed": true,
        "description": "The IP addresses of the DNS servers for the directory, such as [ \"172.31.3.154\", \"172.31.63.203\" ].",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enable_sso": {
        "computed": true,
        "description": "Whether to enable single sign-on for a Simple Active Directory in AWS.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The fully qualified domain name for the AWS Managed Simple AD directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "The password for the default administrative user named Admin.",
        "description_kind": "plain",
        "type": "string"
      },
      "short_name": {
        "computed": true,
        "description": "The NetBIOS name for your domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "The size of the directory.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_settings": {
        "computed": true,
        "description": "VPC settings of the Simple AD directory server in AWS.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "subnet_ids": {
              "computed": true,
              "description": "The identifiers of the subnets for the directory servers. The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description": "The identifier of the VPC in which to create the directory.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DirectoryService::SimpleAD",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDirectoryserviceSimpleAdSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectoryserviceSimpleAd), &result)
	return &result
}
