package resource

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
        "optional": true,
        "type": "string"
      },
      "advanced_options": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "bool"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "iam_federation_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "roles_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject_key": {
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
            "internal_user_database_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "bool"
                  },
                  "public_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "roles_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject_key": {
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
            "master_user_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "master_user_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "master_user_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "master_user_password": {
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
            "saml_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "metadata_content": {
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
                  "master_backend_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "master_user_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "roles_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "session_timeout_minutes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "subject_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "aiml_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_vectors_engine": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "Whether to enable S3 vectors engine.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "dedicated_master_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "dedicated_master_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dedicated_master_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "instance_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "multi_az_with_standby_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "node_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "node_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "count": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "type": {
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
                  "node_type": {
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
            "warm_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "warm_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "warm_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "zone_awareness_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "cognito_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "identity_pool_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_pool_id": {
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
              "optional": true,
              "type": "string"
            },
            "custom_endpoint_certificate_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_endpoint_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enforce_https": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_security_policy": {
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
        "optional": true,
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
              "optional": true,
              "type": "bool"
            },
            "iops": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "throughput": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "volume_type": {
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
      "encryption_at_rest_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_id": {
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
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_center_options": {
        "computed": true,
        "description": "Options for configuring Identity Center",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled_api_access": {
              "computed": true,
              "description": "Whether Identity Center is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "identity_center_application_arn": {
              "computed": true,
              "description": "The ARN of the Identity Center application.",
              "description_kind": "plain",
              "type": "string"
            },
            "identity_center_instance_arn": {
              "computed": true,
              "description": "The ARN of the Identity Center instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identity_store_id": {
              "computed": true,
              "description": "The IdentityStoreId for Identity Center options.",
              "description_kind": "plain",
              "type": "string"
            },
            "roles_key": {
              "computed": true,
              "description": "The roles key for Identity Center options.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subject_key": {
              "computed": true,
              "description": "The subject key for Identity Center options.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "map"
        },
        "optional": true
      },
      "node_to_node_encryption_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "off_peak_window_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
        "optional": true,
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
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "software_update_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_software_update_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOpensearchserviceDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserviceDomain), &result)
	return &result
}
