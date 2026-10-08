package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectoryserviceMicrosoftAd = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The alias for a directory",
        "description_kind": "plain",
        "type": "string"
      },
      "create_alias": {
        "computed": true,
        "description": "Specifies an alias for a directory and assigns the alias to the directory. The alias is used to construct the access URL for the directory, such as http://\u003calias\u003e.awsapps.com. By default, AWS CloudFormation does not create an alias.",
        "description_kind": "plain",
        "type": "bool"
      },
      "directory_id": {
        "computed": true,
        "description": "Unique identifier for the Microsoft Active Directory",
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
      "edition": {
        "computed": true,
        "description": "AWS Managed Microsoft AD is available in two editions: Standard and Enterprise.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_sso": {
        "computed": true,
        "description": "Whether to enable single sign-on for a Microsoft Active Directory in AWS. Single sign-on allows users in your directory to access certain AWS services from a computer joined to the directory without having to enter their credentials separately. If you don't specify a value, AWS CloudFormation disables single sign-on by default.",
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
        "description": "The fully qualified domain name for the AWS Managed Microsoft AD directory, such as corp.example.com. This name will resolve inside your VPC only. It does not need to be publicly resolvable.",
        "description_kind": "plain",
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "The password for the default administrative user named Admin. If you need to change the password for the administrator account, see the ResetUserPassword API call in the AWS Directory Service API Reference.",
        "description_kind": "plain",
        "type": "string"
      },
      "short_name": {
        "computed": true,
        "description": "The NetBIOS name for your domain, such as CORP. If you don't specify a NetBIOS name, it will default to the first part of your directory DNS. For example, CORP for the directory DNS corp.example.com. ",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_settings": {
        "computed": true,
        "description": "Specifies the VPC settings of the Microsoft AD directory server in AWS.",
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
    "description": "Data Source schema for AWS::DirectoryService::MicrosoftAD",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDirectoryserviceMicrosoftAdSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectoryserviceMicrosoftAd), &result)
	return &result
}
