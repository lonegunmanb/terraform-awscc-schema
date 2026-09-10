package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxVolume = `{
  "block": {
    "attributes": {
      "backup_id": {
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      },
      "ontap_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aggregate_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aggregates": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "constituents_per_aggregate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "copy_tags_to_backups": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "junction_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ontap_volume_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "security_style": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "size_in_bytes": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "size_in_megabytes": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "snaplock_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "audit_log_volume": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "autocommit_period": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "privileged_delete": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "retention_period": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "default_retention": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "maximum_retention": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "minimum_retention": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "snaplock_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "volume_append_mode_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "snapshot_policy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "storage_efficiency_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "storage_virtual_machine_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tiering_policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cooling_period": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "volume_style": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "open_zfs_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_tags_to_snapshots": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "data_compression_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nfs_exports": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "clients": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "options": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "options": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "origin_snapshot": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "copy_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "snapshot_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "parent_volume_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "read_only": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "record_size_ki_b": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "storage_capacity_quota_gi_b": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "storage_capacity_reservation_gi_b": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "user_and_group_quotas": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "storage_capacity_quota_gi_b": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_arn": {
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::FSx::Volume",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFsxVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxVolume), &result)
	return &result
}
