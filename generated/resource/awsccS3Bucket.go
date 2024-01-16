package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3Bucket = `{
  "block": {
    "attributes": {
      "accelerate_configuration": {
        "computed": true,
        "description": "Configuration for the transfer acceleration state.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acceleration_status": {
              "description": "Configures the transfer acceleration state for an Amazon S3 bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "access_control": {
        "computed": true,
        "description": "A canned access control list (ACL) that grants predefined permissions to the bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "analytics_configurations": {
        "computed": true,
        "description": "The configuration and any analyses for the analytics filter of an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "description": "The ID that identifies the analytics configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description": "The prefix that an object must have to be included in the analytics results.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_class_analysis": {
              "description": "Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_export": {
                    "computed": true,
                    "description": "Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "destination": {
                          "description": "Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_account_id": {
                                "computed": true,
                                "description": "The account ID that owns the destination S3 bucket. ",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "bucket_arn": {
                                "description": "The Amazon Resource Name (ARN) of the bucket to which data is exported.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "format": {
                                "description": "Specifies the file format used when exporting data to Amazon S3.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "prefix": {
                                "computed": true,
                                "description": "The prefix to use when exporting data. The prefix is prepended to all results.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "output_schema_version": {
                          "description": "The version of the output schema to use when exporting data.",
                          "description_kind": "plain",
                          "required": true,
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
              "required": true
            },
            "tag_filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_encryption": {
        "computed": true,
        "description": "Specifies default encryption for a bucket using server-side encryption with either Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_side_encryption_configuration": {
              "description": "Specifies the default server-side-encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_key_enabled": {
                    "computed": true,
                    "description": "Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Setting the BucketKeyEnabled element to true causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "server_side_encryption_by_default": {
                    "computed": true,
                    "description": "Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kms_master_key_id": {
                          "computed": true,
                          "description": "\"KMSMasterKeyID\" can only be used when you set the value of SSEAlgorithm as aws:kms or aws:kms:dsse.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sse_algorithm": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "bucket_name": {
        "computed": true,
        "description": "A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cors_configuration": {
        "computed": true,
        "description": "Rules that define cross-origin resource sharing of objects in this bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cors_rules": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_headers": {
                    "computed": true,
                    "description": "Headers that are specified in the Access-Control-Request-Headers header.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_methods": {
                    "description": "An HTTP method that you allow the origin to execute.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_origins": {
                    "description": "One or more origins you want customers to be able to access the bucket from.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exposed_headers": {
                    "computed": true,
                    "description": "One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "id": {
                    "computed": true,
                    "description": "A unique identifier for this rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_age": {
                    "computed": true,
                    "description": "The time in seconds that your browser is to cache the preflight response for the specified resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "domain_name": {
        "computed": true,
        "description": "The IPv4 DNS name of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "dual_stack_domain_name": {
        "computed": true,
        "description": "The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "intelligent_tiering_configurations": {
        "computed": true,
        "description": "Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "description": "The ID used to identify the S3 Intelligent-Tiering configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description": "An object key name prefix that identifies the subset of objects to which the rule applies.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "description": "Specifies the status of the configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tag_filters": {
              "computed": true,
              "description": "A container for a key-value pair.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "tierings": {
              "description": "Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration. At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: ARCHIVE_ACCESS and DEEP_ARCHIVE_ACCESS.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_tier": {
                    "description": "S3 Intelligent-Tiering access tier. See Storage class for automatically optimizing frequently and infrequently accessed objects for a list of access tiers in the S3 Intelligent-Tiering storage class.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "days": {
                    "description": "The number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier. The minimum number of days specified for Archive Access tier must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to 2 years (730 days).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "inventory_configurations": {
        "computed": true,
        "description": "The inventory configuration for an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination": {
              "description": "Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_account_id": {
                    "computed": true,
                    "description": "The account ID that owns the destination S3 bucket. ",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "bucket_arn": {
                    "description": "The Amazon Resource Name (ARN) of the bucket to which data is exported.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "format": {
                    "description": "Specifies the file format used when exporting data to Amazon S3.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "The prefix to use when exporting data. The prefix is prepended to all results.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "enabled": {
              "description": "Specifies whether the inventory is enabled or disabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "id": {
              "description": "The ID used to identify the inventory configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "included_object_versions": {
              "description": "Object versions to include in the inventory list.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "optional_fields": {
              "computed": true,
              "description": "Contains the optional fields that are included in the inventory results.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "prefix": {
              "computed": true,
              "description": "The prefix that is prepended to all inventory results.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_frequency": {
              "description": "Specifies the schedule for generating inventory results.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "lifecycle_configuration": {
        "computed": true,
        "description": "Rules that define how Amazon S3 manages objects during their lifetime.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "description": "A lifecycle rule for individual objects in an Amazon S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "abort_incomplete_multipart_upload": {
                    "computed": true,
                    "description": "Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days_after_initiation": {
                          "description": "Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "expiration_date": {
                    "computed": true,
                    "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expiration_in_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "expired_object_delete_marker": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "noncurrent_version_expiration": {
                    "computed": true,
                    "description": "Container for the expiration rule that describes when noncurrent objects are expired. If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 expire noncurrent object versions at a specific period in the object's lifetime",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "newer_noncurrent_versions": {
                          "computed": true,
                          "description": "Specified the number of newer noncurrent and current versions that must exists before performing the associated action",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "noncurrent_days": {
                          "description": "Specified the number of days an object is noncurrent before Amazon S3 can perform the associated action",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "noncurrent_version_expiration_in_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "noncurrent_version_transition": {
                    "computed": true,
                    "description": "Container for the transition rule that describes when noncurrent objects transition to the STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER_IR, GLACIER, or DEEP_ARCHIVE storage class. If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 transition noncurrent object versions to the STANDARD_IA, ONEZONE_IA, INTELLIGENT_TIERING, GLACIER_IR, GLACIER, or DEEP_ARCHIVE storage class at a specific period in the object's lifetime.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "newer_noncurrent_versions": {
                          "computed": true,
                          "description": "Specified the number of newer noncurrent and current versions that must exists before performing the associated action",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_class": {
                          "description": "The class of storage used to store the object.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transition_in_days": {
                          "description": "Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "noncurrent_version_transitions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "newer_noncurrent_versions": {
                          "computed": true,
                          "description": "Specified the number of newer noncurrent and current versions that must exists before performing the associated action",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_class": {
                          "description": "The class of storage used to store the object.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transition_in_days": {
                          "description": "Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "object_size_greater_than": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "object_size_less_than": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "tag_filters": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "transition": {
                    "computed": true,
                    "description": "You must specify at least one of \"TransitionDate\" and \"TransitionInDays\"",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_class": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transition_date": {
                          "computed": true,
                          "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "transition_in_days": {
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
                  "transitions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_class": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transition_date": {
                          "computed": true,
                          "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "transition_in_days": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "logging_configuration": {
        "computed": true,
        "description": "Settings that define where logs are stored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_bucket_name": {
              "computed": true,
              "description": "The name of an Amazon S3 bucket where Amazon S3 store server access log files. You can store log files in any bucket that you own. By default, logs are stored in the bucket where the LoggingConfiguration property is defined.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_file_prefix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_object_key_format": {
              "computed": true,
              "description": "Describes the key format for server access log file in the target bucket. You can choose between SimplePrefix and PartitionedPrefix.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "partitioned_prefix": {
                    "computed": true,
                    "description": "This format appends a time based prefix to the given log file prefix for delivering server access log file.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_date_source": {
                          "computed": true,
                          "description": "Date Source for creating a partitioned prefix. This can be event time or delivery time.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "simple_prefix": {
                    "computed": true,
                    "description": "This format defaults the prefix to the given log file prefix for delivering server access log file.",
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
      "metrics_configurations": {
        "computed": true,
        "description": "Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_point_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_filters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "notification_configuration": {
        "computed": true,
        "description": "Configuration that defines how Amazon S3 handles bucket notifications.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_bridge_configuration": {
              "computed": true,
              "description": "Describes the Amazon EventBridge notification configuration for an Amazon S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event_bridge_enabled": {
                    "computed": true,
                    "description": "Specifies whether to send notifications to Amazon EventBridge when events occur in an Amazon S3 bucket.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "lambda_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event": {
                    "description": "The Amazon S3 bucket event for which to invoke the AWS Lambda function.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The filtering rules that determine which objects invoke the AWS Lambda function.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_key": {
                          "description": "A container for object key name prefix and suffix filtering rules.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rules": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "function": {
                    "description": "The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon S3 invokes when the specified event type occurs.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "queue_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event": {
                    "description": "The Amazon S3 bucket event about which you want to publish messages to Amazon SQS.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The filtering rules that determine which objects trigger notifications.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_key": {
                          "description": "A container for object key name prefix and suffix filtering rules.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rules": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "queue": {
                    "description": "The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "topic_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event": {
                    "description": "The Amazon S3 bucket event about which to send notifications.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The filtering rules that determine for which objects to send notifications.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_key": {
                          "description": "A container for object key name prefix and suffix filtering rules.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rules": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "topic": {
                    "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.",
                    "description_kind": "plain",
                    "required": true,
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
      "object_lock_configuration": {
        "computed": true,
        "description": "Places an Object Lock configuration on the specified bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "object_lock_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule": {
              "computed": true,
              "description": "The Object Lock rule in place for the specified object.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_retention": {
                    "computed": true,
                    "description": "The default retention period that you want to apply to new objects placed in the specified bucket.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "years": {
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
      "object_lock_enabled": {
        "computed": true,
        "description": "Indicates whether this bucket has an Object Lock configuration enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ownership_controls": {
        "computed": true,
        "description": "Specifies the container element for object ownership rules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "object_ownership": {
                    "computed": true,
                    "description": "Specifies an object ownership rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "public_access_block_configuration": {
        "computed": true,
        "description": "Configuration that defines how Amazon S3 handles public access.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "block_public_policy": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.\nEnabling this setting doesn't affect existing bucket policies.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ignore_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.\nEnabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "restrict_public_buckets": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "regional_domain_name": {
        "computed": true,
        "description": "Returns the regional domain name of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_configuration": {
        "computed": true,
        "description": "Configuration for replicating objects in an S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role": {
              "description": "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rules": {
              "description": "A container for one or more replication rules.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delete_marker_replication": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
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
                  "destination": {
                    "description": "Specifies which Amazon S3 bucket to store replicated objects in and their storage class.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_control_translation": {
                          "computed": true,
                          "description": "Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "owner": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "account": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "encryption_configuration": {
                          "computed": true,
                          "description": "Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "replica_kms_key_id": {
                                "description": "Specifies the ID (Key ARN or Alias ARN) of the customer managed customer master key (CMK) stored in AWS Key Management Service (KMS) for the destination bucket.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "metrics": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "event_threshold": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "minutes": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "status": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "replication_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "time": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "minutes": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "required": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "storage_class": {
                          "computed": true,
                          "description": "The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "filter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "and": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prefix": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "tag_filters": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "key": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "description_kind": "plain",
                                      "required": true,
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
                        "prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tag_filter": {
                          "computed": true,
                          "description": "Tags to use to identify a subset of objects for an Amazon S3 bucket.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description_kind": "plain",
                                "required": true,
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
                  "id": {
                    "computed": true,
                    "description": "A unique identifier for the rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "An object key name prefix that identifies the object or objects to which the rule applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "source_selection_criteria": {
                    "computed": true,
                    "description": "A container that describes additional filters for identifying the source objects that you want to replicate.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "replica_modifications": {
                          "computed": true,
                          "description": "A filter that you can specify for selection for modifications on replicas.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status": {
                                "description": "Specifies whether Amazon S3 replicates modifications on replicas.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "sse_kms_encrypted_objects": {
                          "computed": true,
                          "description": "A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status": {
                                "description": "Specifies whether Amazon S3 replicates objects created with server-side encryption using a customer master key (CMK) stored in AWS Key Management Service.",
                                "description_kind": "plain",
                                "required": true,
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
                  "status": {
                    "description": "Specifies whether the rule is enabled.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "versioning_configuration": {
        "computed": true,
        "description": "Describes the versioning state of an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "The versioning state of the bucket.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "website_configuration": {
        "computed": true,
        "description": "Specifies website configuration parameters for an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "error_document": {
              "computed": true,
              "description": "The name of the error document for the website.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "index_document": {
              "computed": true,
              "description": "The name of the index document for the website.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "redirect_all_requests_to": {
              "computed": true,
              "description": "Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host_name": {
                    "description": "Name of the host where requests are redirected.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "Protocol to use when redirecting requests. The default is the protocol that is used in the original request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "routing_rules": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "redirect_rule": {
                    "description": "Container for redirect information. You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "host_name": {
                          "computed": true,
                          "description": "The host name to use in the redirect request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "http_redirect_code": {
                          "computed": true,
                          "description": "The HTTP redirect code to use on the response. Not required if one of the siblings is present.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "protocol": {
                          "computed": true,
                          "description": "Protocol to use when redirecting requests. The default is the protocol that is used in the original request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replace_key_prefix_with": {
                          "computed": true,
                          "description": "The object key prefix to use in the redirect request.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "replace_key_with": {
                          "computed": true,
                          "description": "The specific object key to use in the redirect request.d",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "routing_rule_condition": {
                    "computed": true,
                    "description": "A container for describing a condition that must be met for the specified redirect to apply.You must specify at least one of HttpErrorCodeReturnedEquals and KeyPrefixEquals",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "http_error_code_returned_equals": {
                          "computed": true,
                          "description": "The HTTP error code when the redirect is applied. ",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_prefix_equals": {
                          "computed": true,
                          "description": "The object key name prefix when the redirect is applied.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "website_url": {
        "computed": true,
        "description": "The Amazon S3 website endpoint for the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3::Bucket",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3BucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3Bucket), &result)
	return &result
}
