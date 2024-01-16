package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailInstance = `{
  "block": {
    "attributes": {
      "add_ons": {
        "computed": true,
        "description": "An array of objects representing the add-ons to enable for the new instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "add_on_type": {
              "computed": true,
              "description": "The add-on type",
              "description_kind": "plain",
              "type": "string"
            },
            "auto_snapshot_add_on_request": {
              "computed": true,
              "description": "An object that represents additional parameters when enabling or modifying the automatic snapshot add-on",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "snapshot_time_of_day": {
                    "computed": true,
                    "description": "The daily time when an automatic snapshot will be created.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "status": {
              "computed": true,
              "description": "Status of the Addon",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
        "description_kind": "plain",
        "type": "string"
      },
      "blueprint_id": {
        "computed": true,
        "description": "The ID for a virtual private server image (e.g., app_wordpress_4_4 or app_lamp_7_0 ). Use the get blueprints operation to return a list of available images (or blueprints ).",
        "description_kind": "plain",
        "type": "string"
      },
      "bundle_id": {
        "computed": true,
        "description": "The bundle of specification information for your virtual private server (or instance ), including the pricing plan (e.g., micro_1_0 ).",
        "description_kind": "plain",
        "type": "string"
      },
      "hardware": {
        "computed": true,
        "description": "Hardware of the Instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu_count": {
              "computed": true,
              "description": "CPU count of the Instance.",
              "description_kind": "plain",
              "type": "number"
            },
            "disks": {
              "computed": true,
              "description": "Disks attached to the Instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attached_to": {
                    "computed": true,
                    "description": "Instance attached to the disk.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "attachment_state": {
                    "computed": true,
                    "description": "Attachment state of the disk.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "disk_name": {
                    "computed": true,
                    "description": "The names to use for your new Lightsail disk.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "iops": {
                    "computed": true,
                    "description": "IOPS of disk.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "is_system_disk": {
                    "computed": true,
                    "description": "Is the Attached disk is the system disk of the Instance.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "path": {
                    "computed": true,
                    "description": "Path of the disk attached to the instance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "size_in_gb": {
                    "computed": true,
                    "description": "Size of the disk attached to the Instance.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "ram_size_in_gb": {
              "computed": true,
              "description": "RAM Size of the Instance.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_name": {
        "computed": true,
        "description": "The names to use for your new Lightsail instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_static_ip": {
        "computed": true,
        "description": "Is the IP Address of the Instance is the static IP",
        "description_kind": "plain",
        "type": "bool"
      },
      "key_pair_name": {
        "computed": true,
        "description": "The name of your key pair.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "Location of a resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.",
              "description_kind": "plain",
              "type": "string"
            },
            "region_name": {
              "computed": true,
              "description": "The Region Name in which to create your instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "networking": {
        "computed": true,
        "description": "Networking of the Instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "monthly_transfer": {
              "computed": true,
              "description": "Monthly Transfer of the Instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "gb_per_month_allocated": {
                    "computed": true,
                    "description": "GbPerMonthAllocated of the Instance.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "ports": {
              "computed": true,
              "description": "Ports to the Instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_direction": {
                    "computed": true,
                    "description": "Access Direction for Protocol of the Instance(inbound/outbound).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "access_from": {
                    "computed": true,
                    "description": "Access From Protocol of the Instance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "access_type": {
                    "computed": true,
                    "description": "Access Type Protocol of the Instance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cidr_list_aliases": {
                    "computed": true,
                    "description": "cidr List Aliases",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cidrs": {
                    "computed": true,
                    "description": "cidrs",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "common_name": {
                    "computed": true,
                    "description": "CommonName for Protocol of the Instance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "from_port": {
                    "computed": true,
                    "description": "From Port of the Instance.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "ipv_6_cidrs": {
                    "computed": true,
                    "description": "IPv6 Cidrs",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "protocol": {
                    "computed": true,
                    "description": "Port Protocol of the Instance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "to_port": {
                    "computed": true,
                    "description": "To Port of the Instance.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "private_ip_address": {
        "computed": true,
        "description": "Private IP Address of the Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "public_ip_address": {
        "computed": true,
        "description": "Public IP Address of the Instance",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "Resource type of Lightsail instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "ssh_key_name": {
        "computed": true,
        "description": "SSH Key Name of the  Lightsail instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Current State of the Instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code": {
              "computed": true,
              "description": "Status code of the Instance.",
              "description_kind": "plain",
              "type": "number"
            },
            "name": {
              "computed": true,
              "description": "Status code of the Instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "support_code": {
        "computed": true,
        "description": "Support code to help identify any issues",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "user_data": {
        "computed": true,
        "description": "A launch script you can create that configures a server with additional user data. For example, you might want to run apt-get -y update.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description": "Username of the  Lightsail instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lightsail::Instance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLightsailInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailInstance), &result)
	return &result
}
