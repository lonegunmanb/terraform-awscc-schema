package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftserverlessWorkgroup = `{
  "block": {
    "attributes": {
      "base_capacity": {
        "computed": true,
        "description": "The base compute capacity of the workgroup in Redshift Processing Units (RPUs).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "config_parameters": {
        "computed": true,
        "description": "A list of parameters to set for finer control over a database. Available options are datestyle, enable_user_activity_logging, query_group, search_path, max_query_execution_time, and require_ssl.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameter_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameter_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description": "The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_capacity": {
        "computed": true,
        "description": "The max compute capacity of the workgroup in Redshift Processing Units (RPUs).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "namespace_name": {
        "computed": true,
        "description": "The namespace the workgroup is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "price_performance_target": {
        "computed": true,
        "description": "A property that represents the price performance target settings for the workgroup.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "level": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "publicly_accessible": {
        "computed": true,
        "description": "A value that specifies whether the workgroup can be accessible from a public network.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "recovery_point_id": {
        "computed": true,
        "description": "The identifier of the recovery point to restore the namespace from. When this resource is first created, the namespace is restored from this recovery point. On subsequent updates, a restore occurs only when RecoveryPointId changes from its previous value. If the value is unchanged or removed, no restore takes place and existing data is preserved.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "A list of security group IDs to associate with the workgroup.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "snapshot_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the snapshot to restore the namespace from. Specify either SnapshotArn or SnapshotName, but not both. When this resource is first created, the namespace is restored from this snapshot. On subsequent updates, a restore occurs only when SnapshotArn changes from its previous value. If the value is unchanged or removed, no restore takes place and existing data is preserved.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_name": {
        "computed": true,
        "description": "The name of the snapshot to restore the namespace from. Because snapshot names are unique only within an account, also specify SnapshotOwnerAccount when restoring from a snapshot owned by a different account. Specify either SnapshotName or SnapshotArn, but not both. When this resource is first created, the namespace is restored from this snapshot. On subsequent updates, a restore occurs only when SnapshotName or SnapshotOwnerAccount changes from its previous value. If both values are unchanged or SnapshotName is removed, no restore takes place and existing data is preserved.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_owner_account": {
        "computed": true,
        "description": "The AWS account ID that owns the snapshot. Required when restoring from a snapshot shared by another account. Used in combination with SnapshotName. On updates, changing this value while SnapshotName is set triggers a restore from the newly referenced snapshot. If the value is unchanged, no restore takes place and existing data is preserved.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "A list of subnet IDs the workgroup is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The map of the key-value pairs used to tag the workgroup.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "track_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workgroup": {
        "computed": true,
        "description": "Definition for workgroup resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "config_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "parameter_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "creation_date": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "endpoint": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "vpc_endpoints": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "network_interfaces": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "availability_zone": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "network_interface_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "private_ip_address": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "subnet_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "vpc_endpoint_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "vpc_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "enhanced_vpc_routing": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "max_capacity": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "namespace_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "price_performance_target": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "level": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "publicly_accessible": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "track_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "workgroup_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "workgroup_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "workgroup_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "workgroup_name": {
        "description": "The name of the workgroup.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::RedshiftServerless::Workgroup Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftserverlessWorkgroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftserverlessWorkgroup), &result)
	return &result
}
