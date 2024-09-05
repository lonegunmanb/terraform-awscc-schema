package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserviceDomain = `{
  "block": {
    "attributes": {
      "access_policies": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "advanced_options": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "advanced_security_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "anonymous_auth_disable_date": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "anonymous_auth_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "internal_user_database_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "jwt_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "public_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "roles_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subject_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "master_user_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "master_user_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "master_user_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "master_user_password": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "saml_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "idp": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "entity_id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metadata_content": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "master_backend_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "master_user_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "roles_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "session_timeout_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "subject_key": {
                    "computed": true,
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
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cold_storage_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "dedicated_master_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "dedicated_master_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "dedicated_master_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "multi_az_with_standby_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "warm_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "warm_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "warm_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "zone_awareness_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_count": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "zone_awareness_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cognito_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "identity_pool_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "user_pool_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_endpoint_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_endpoint": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "custom_endpoint_certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "custom_endpoint_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enforce_https": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "tls_security_policy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_endpoint_v2": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "domain_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ebs_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "iops": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "throughput": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "volume_size": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "volume_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "encryption_at_rest_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "engine_version": {
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
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_publishing_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_log_group_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "map"
        }
      },
      "node_to_node_encryption_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "off_peak_window_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "off_peak_window": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "window_start_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hours": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "minutes": {
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "service_software_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "automated_update_date": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "cancellable": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "current_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "new_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "optional_deployment": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "update_available": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "update_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "skip_shard_migration_wait": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "snapshot_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "automated_snapshot_start_hour": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "software_update_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_software_update_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this Domain.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::OpenSearchService::Domain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserviceDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserviceDomain), &result)
	return &result
}
