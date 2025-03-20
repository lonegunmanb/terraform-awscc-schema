package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3Bucket = `{
  "block": {
    "attributes": {
      "accelerate_configuration": {
        "computed": true,
        "description": "Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acceleration_status": {
              "computed": true,
              "description": "Specifies the transfer acceleration status of the bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "access_control": {
        "computed": true,
        "description": "This is a legacy property, and it is not recommended for most use cases. A majority of modern use cases in Amazon S3 no longer require the use of ACLs, and we recommend that you keep ACLs disabled. For more information, see [Controlling object ownership](https://docs.aws.amazon.com//AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide*.\n  A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide*.\n  S3 buckets are created with ACLs disabled by default. Therefore, unless you explicitly set the [AWS::S3::OwnershipControls](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html) property to enable ACLs, your resource will fail to deploy with any value other than Private. Use cases requiring ACLs are uncommon.\n  The majority of access control configurations can be successfully and more easily achieved with bucket policies. For more information, see [AWS::S3::BucketPolicy](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html). For examples of common policy configurations, including S3 Server Access Logs buckets and more, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "analytics_configurations": {
        "computed": true,
        "description": "Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The ID that identifies the analytics configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description": "The prefix that an object must have to be included in the analytics results.",
              "description_kind": "plain",
              "type": "string"
            },
            "storage_class_analysis": {
              "computed": true,
              "description": "Contains data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes.",
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
                          "computed": true,
                          "description": "The place to store the data for an analysis.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket_account_id": {
                                "computed": true,
                                "description": "The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not validated before exporting data.\n   Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "bucket_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the bucket to which data is exported.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "format": {
                                "computed": true,
                                "description": "Specifies the file format used when exporting data to Amazon S3.\n  *Allowed values*: ` + "`" + `` + "`" + `CSV` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `ORC` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `Parquet` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prefix": {
                                "computed": true,
                                "description": "The prefix to use when exporting data. The prefix is prepended to all results.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "output_schema_version": {
                          "computed": true,
                          "description": "The version of the output schema to use when exporting data. Must be ` + "`" + `` + "`" + `V_1` + "`" + `` + "`" + `.",
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
            "tag_filters": {
              "computed": true,
              "description": "The tags to use when evaluating an analytics filter.\n The analytics only includes objects that meet the filter's criteria. If no filter is specified, all of the contents of the bucket are included in the analysis.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The tag value.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "arn": {
        "computed": true,
        "description": "the Amazon Resource Name (ARN) of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_encryption": {
        "computed": true,
        "description": "Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3), AWS KMS-managed keys (SSE-KMS), or dual-layer server-side encryption with KMS-managed keys (DSSE-KMS). For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_side_encryption_configuration": {
              "computed": true,
              "description": "Specifies the default server-side-encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_key_enabled": {
                    "computed": true,
                    "description": "Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Setting the ` + "`" + `` + "`" + `BucketKeyEnabled` + "`" + `` + "`" + ` element to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.\n For more information, see [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
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
                          "description": "AWS Key Management Service (KMS) customer managed key ID to use for the default encryption. \n   +   *General purpose buckets* - This parameter is allowed if and only if ` + "`" + `` + "`" + `SSEAlgorithm` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `aws:kms` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `aws:kms:dsse` + "`" + `` + "`" + `.\n  +   *Directory buckets* - This parameter is allowed if and only if ` + "`" + `` + "`" + `SSEAlgorithm` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `aws:kms` + "`" + `` + "`" + `.\n  \n  You can specify the key ID, key alias, or the Amazon Resource Name (ARN) of the KMS key.\n  +  Key ID: ` + "`" + `` + "`" + `1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + ` \n  +  Key ARN: ` + "`" + `` + "`" + `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + ` \n  +  Key Alias: ` + "`" + `` + "`" + `alias/alias-name` + "`" + `` + "`" + ` \n  \n If you are using encryption with cross-account or AWS service operations, you must use a fully qualified KMS key ARN. For more information, see [Using encryption for cross-account operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html#bucket-encryption-update-bucket-policy).\n   +   *General purpose buckets* - If you're specifying a customer managed KMS key, we recommend using a fully qualified KMS key ARN. If you use a KMS key alias instead, then KMS resolves the key within the requester?s account. This behavior can result in data that's encrypted with a KMS key that belongs to the requester, and not the bucket owner. Also, if you use a key ID, you can run into a LogDestination undeliverable error when creating a VPC flow log. \n  +   *Directory buckets* - When you specify an [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk) for encryption in your directory bucket, only use the key ID or key ARN. The key alias format of the KMS key isn't supported.\n  \n   Amazon S3 only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *Key Management Service Developer Guide*.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sse_algorithm": {
                          "computed": true,
                          "description": "Server-side encryption algorithm to use for the default encryption.\n  For directory buckets, there are only two supported values for server-side encryption: ` + "`" + `` + "`" + `AES256` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `aws:kms` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "bucket_name": {
        "computed": true,
        "description": "A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html). For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html) in the *Amazon S3 User Guide*. \n  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.",
        "description_kind": "plain",
        "type": "string"
      },
      "cors_configuration": {
        "computed": true,
        "description": "Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cors_rules": {
              "computed": true,
              "description": "A set of origins and methods (cross-origin access that you want to allow). You can add up to 100 rules to the configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_headers": {
                    "computed": true,
                    "description": "Headers that are specified in the ` + "`" + `` + "`" + `Access-Control-Request-Headers` + "`" + `` + "`" + ` header. These headers are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns any requested headers that are allowed.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_methods": {
                    "computed": true,
                    "description": "An HTTP method that you allow the origin to run.\n  *Allowed values*: ` + "`" + `` + "`" + `GET` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `PUT` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `HEAD` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `POST` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `DELETE` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_origins": {
                    "computed": true,
                    "description": "One or more origins you want customers to be able to access the bucket from.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "exposed_headers": {
                    "computed": true,
                    "description": "One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript ` + "`" + `` + "`" + `XMLHttpRequest` + "`" + `` + "`" + ` object).",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "id": {
                    "computed": true,
                    "description": "A unique identifier for this rule. The value must be no more than 255 characters.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "max_age": {
                    "computed": true,
                    "description": "The time in seconds that your browser is to cache the preflight response for the specified resource.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dual_stack_domain_name": {
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
      "intelligent_tiering_configurations": {
        "computed": true,
        "description": "Defines how Amazon S3 handles Intelligent-Tiering storage.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The ID used to identify the S3 Intelligent-Tiering configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description": "An object key name prefix that identifies the subset of objects to which the rule applies.",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "Specifies the status of the configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "tag_filters": {
              "computed": true,
              "description": "A container for a key-value pair.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The tag value.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "tierings": {
              "computed": true,
              "description": "Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration. At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: ` + "`" + `` + "`" + `ARCHIVE_ACCESS` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `DEEP_ARCHIVE_ACCESS` + "`" + `` + "`" + `.\n  You only need Intelligent Tiering Configuration enabled on a bucket if you want to automatically move objects stored in the Intelligent-Tiering storage class to Archive Access or Deep Archive Access tiers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_tier": {
                    "computed": true,
                    "description": "S3 Intelligent-Tiering access tier. See [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) for a list of access tiers in the S3 Intelligent-Tiering storage class.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "days": {
                    "computed": true,
                    "description": "The number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier. The minimum number of days specified for Archive Access tier must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to 2 years (730 days).",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "inventory_configurations": {
        "computed": true,
        "description": "Specifies the inventory configuration for an Amazon S3 bucket. For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination": {
              "computed": true,
              "description": "Contains information about where to publish the inventory results.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_account_id": {
                    "computed": true,
                    "description": "The account ID that owns the destination S3 bucket. If no account ID is provided, the owner is not validated before exporting data.\n   Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "bucket_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the bucket to which data is exported.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "format": {
                    "computed": true,
                    "description": "Specifies the file format used when exporting data to Amazon S3.\n  *Allowed values*: ` + "`" + `` + "`" + `CSV` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `ORC` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `Parquet` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "The prefix to use when exporting data. The prefix is prepended to all results.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "enabled": {
              "computed": true,
              "description": "Specifies whether the inventory is enabled or disabled. If set to ` + "`" + `` + "`" + `True` + "`" + `` + "`" + `, an inventory list is generated. If set to ` + "`" + `` + "`" + `False` + "`" + `` + "`" + `, no inventory list is generated.",
              "description_kind": "plain",
              "type": "bool"
            },
            "id": {
              "computed": true,
              "description": "The ID used to identify the inventory configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "included_object_versions": {
              "computed": true,
              "description": "Object versions to include in the inventory list. If set to ` + "`" + `` + "`" + `All` + "`" + `` + "`" + `, the list includes all the object versions, which adds the version-related fields ` + "`" + `` + "`" + `VersionId` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `IsLatest` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `DeleteMarker` + "`" + `` + "`" + ` to the list. If set to ` + "`" + `` + "`" + `Current` + "`" + `` + "`" + `, the list does not contain these version-related fields.",
              "description_kind": "plain",
              "type": "string"
            },
            "optional_fields": {
              "computed": true,
              "description": "Contains the optional fields that are included in the inventory results.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "prefix": {
              "computed": true,
              "description": "Specifies the inventory filter prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "schedule_frequency": {
              "computed": true,
              "description": "Specifies the schedule for generating inventory results.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "lifecycle_configuration": {
        "computed": true,
        "description": "Specifies the lifecycle configuration for objects in an Amazon S3 bucket. For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "A lifecycle rule for individual objects in an Amazon S3 bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "abort_incomplete_multipart_upload": {
                    "computed": true,
                    "description": "Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3 bucket.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days_after_initiation": {
                          "computed": true,
                          "description": "Specifies the number of days after which Amazon S3 stops an incomplete multipart upload.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "expiration_date": {
                    "computed": true,
                    "description": "Indicates when objects are deleted from Amazon S3 and Amazon S3 Glacier. The date value must be in ISO 8601 format. The time is always midnight UTC. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expiration_in_days": {
                    "computed": true,
                    "description": "Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon S3 Glacier. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "expired_object_delete_marker": {
                    "computed": true,
                    "description": "Indicates whether Amazon S3 will remove a delete marker without any noncurrent versions. If set to true, the delete marker will be removed if there are no noncurrent versions. This cannot be specified with ` + "`" + `` + "`" + `ExpirationInDays` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ExpirationDate` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `TagFilters` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "id": {
                    "computed": true,
                    "description": "Unique identifier for the rule. The value can't be longer than 255 characters.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "noncurrent_version_expiration": {
                    "computed": true,
                    "description": "Specifies when noncurrent object versions expire. Upon expiration, S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that S3 delete noncurrent object versions at a specific period in the object's lifetime.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "newer_noncurrent_versions": {
                          "computed": true,
                          "description": "Specifies how many noncurrent versions S3 will retain. If there are this many more recent noncurrent versions, S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "noncurrent_days": {
                          "computed": true,
                          "description": "Specifies the number of days an object is noncurrent before S3 can perform the associated action. For information about the noncurrent days calculations, see [How Amazon S3 Calculates When an Object Became Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "noncurrent_version_expiration_in_days": {
                    "computed": true,
                    "description": "(Deprecated.) For buckets with versioning enabled (or suspended), specifies the time, in days, between when a new version of the object is uploaded to the bucket and when old versions of the object expire. When object versions expire, Amazon S3 permanently deletes them. If you specify a transition and expiration time, the expiration time must be later than the transition time.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "noncurrent_version_transition": {
                    "computed": true,
                    "description": "(Deprecated.) For buckets with versioning enabled (or suspended), specifies when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the ` + "`" + `` + "`" + `NoncurrentVersionTransitions` + "`" + `` + "`" + ` property.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "newer_noncurrent_versions": {
                          "computed": true,
                          "description": "Specifies how many noncurrent versions S3 will retain. If there are this many more recent noncurrent versions, S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "storage_class": {
                          "computed": true,
                          "description": "The class of storage used to store the object.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "transition_in_days": {
                          "computed": true,
                          "description": "Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action. For information about the noncurrent days calculations, see [How Amazon S3 Calculates How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "noncurrent_version_transitions": {
                    "computed": true,
                    "description": "For buckets with versioning enabled (or suspended), one or more transition rules that specify when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the ` + "`" + `` + "`" + `NoncurrentVersionTransition` + "`" + `` + "`" + ` property.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "newer_noncurrent_versions": {
                          "computed": true,
                          "description": "Specifies how many noncurrent versions S3 will retain. If there are this many more recent noncurrent versions, S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "storage_class": {
                          "computed": true,
                          "description": "The class of storage used to store the object.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "transition_in_days": {
                          "computed": true,
                          "description": "Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action. For information about the noncurrent days calculations, see [How Amazon S3 Calculates How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "object_size_greater_than": {
                    "computed": true,
                    "description": "Specifies the minimum object size in bytes for this rule to apply to. Objects must be larger than this value in bytes. For more information about size based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "object_size_less_than": {
                    "computed": true,
                    "description": "Specifies the maximum object size in bytes for this rule to apply to. Objects must be smaller than this value in bytes. For more information about sized based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "Object key prefix that identifies one or more objects to which this rule applies.\n  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "If ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + `, the rule is currently being applied. If ` + "`" + `` + "`" + `Disabled` + "`" + `` + "`" + `, the rule is not currently being applied.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "tag_filters": {
                    "computed": true,
                    "description": "Tags to use to identify a subset of objects to which the lifecycle rule applies.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The tag key.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The tag value.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "transition": {
                    "computed": true,
                    "description": "(Deprecated.) Specifies when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the ` + "`" + `` + "`" + `Transitions` + "`" + `` + "`" + ` property.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_class": {
                          "computed": true,
                          "description": "The storage class to which you want the object to transition.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "transition_date": {
                          "computed": true,
                          "description": "Indicates when objects are transitioned to the specified storage class. The date value must be in ISO 8601 format. The time is always midnight UTC.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "transition_in_days": {
                          "computed": true,
                          "description": "Indicates the number of days after creation when objects are transitioned to the specified storage class. If the specified storage class is ` + "`" + `` + "`" + `INTELLIGENT_TIERING` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `GLACIER_IR` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `GLACIER` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `DEEP_ARCHIVE` + "`" + `` + "`" + `, valid values are ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` or positive integers. If the specified storage class is ` + "`" + `` + "`" + `STANDARD_IA` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ONEZONE_IA` + "`" + `` + "`" + `, valid values are positive integers greater than ` + "`" + `` + "`" + `30` + "`" + `` + "`" + `. Be aware that some storage classes have a minimum storage duration and that you're charged for transitioning objects before their minimum storage duration. For more information, see [Constraints and considerations for transitions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html#lifecycle-configuration-constraints) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "transitions": {
                    "computed": true,
                    "description": "One or more transition rules that specify when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the ` + "`" + `` + "`" + `Transition` + "`" + `` + "`" + ` property.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_class": {
                          "computed": true,
                          "description": "The storage class to which you want the object to transition.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "transition_date": {
                          "computed": true,
                          "description": "Indicates when objects are transitioned to the specified storage class. The date value must be in ISO 8601 format. The time is always midnight UTC.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "transition_in_days": {
                          "computed": true,
                          "description": "Indicates the number of days after creation when objects are transitioned to the specified storage class. If the specified storage class is ` + "`" + `` + "`" + `INTELLIGENT_TIERING` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `GLACIER_IR` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `GLACIER` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `DEEP_ARCHIVE` + "`" + `` + "`" + `, valid values are ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` or positive integers. If the specified storage class is ` + "`" + `` + "`" + `STANDARD_IA` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ONEZONE_IA` + "`" + `` + "`" + `, valid values are positive integers greater than ` + "`" + `` + "`" + `30` + "`" + `` + "`" + `. Be aware that some storage classes have a minimum storage duration and that you're charged for transitioning objects before their minimum storage duration. For more information, see [Constraints and considerations for transitions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-transition-general-considerations.html#lifecycle-configuration-constraints) in the *Amazon S3 User Guide*.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "transition_default_minimum_object_size": {
              "computed": true,
              "description": "Indicates which default minimum object size behavior is applied to the lifecycle configuration.\n  This parameter applies to general purpose buckets only. It isn't supported for directory bucket lifecycle configurations.\n   +   ` + "`" + `` + "`" + `all_storage_classes_128K` + "`" + `` + "`" + ` - Objects smaller than 128 KB will not transition to any storage class by default.\n  +   ` + "`" + `` + "`" + `varies_by_storage_class` + "`" + `` + "`" + ` - Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes will prevent transitions smaller than 128 KB. \n  \n To customize the minimum object size for any transition you can add a filter that specifies a custom ` + "`" + `` + "`" + `ObjectSizeGreaterThan` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ObjectSizeLessThan` + "`" + `` + "`" + ` in the body of your transition rule. Custom filters always take precedence over the default transition behavior.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "logging_configuration": {
        "computed": true,
        "description": "Settings that define where logs are stored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_bucket_name": {
              "computed": true,
              "description": "The name of the bucket where Amazon S3 should store server access log files. You can store log files in any bucket that you own. By default, logs are stored in the bucket where the ` + "`" + `` + "`" + `LoggingConfiguration` + "`" + `` + "`" + ` property is defined.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_file_prefix": {
              "computed": true,
              "description": "A prefix for all log object keys. If you store log files from multiple Amazon S3 buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_object_key_format": {
              "computed": true,
              "description": "Amazon S3 key format for log objects. Only one format, either PartitionedPrefix or SimplePrefix, is allowed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "partitioned_prefix": {
                    "computed": true,
                    "description": "Amazon S3 keys for log objects are partitioned in the following format:\n  ` + "`" + `` + "`" + `[DestinationPrefix][SourceAccountId]/[SourceRegion]/[SourceBucket]/[YYYY]/[MM]/[DD]/[YYYY]-[MM]-[DD]-[hh]-[mm]-[ss]-[UniqueString]` + "`" + `` + "`" + ` \n PartitionedPrefix defaults to EventTime delivery when server access logs are delivered.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "partition_date_source": {
                          "computed": true,
                          "description": "Specifies the partition date source for the partitioned prefix. ` + "`" + `` + "`" + `PartitionDateSource` + "`" + `` + "`" + ` can be ` + "`" + `` + "`" + `EventTime` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DeliveryTime` + "`" + `` + "`" + `.\n For ` + "`" + `` + "`" + `DeliveryTime` + "`" + `` + "`" + `, the time in the log file names corresponds to the delivery time for the log files. \n  For ` + "`" + `` + "`" + `EventTime` + "`" + `` + "`" + `, The logs delivered are for a specific day only. The year, month, and day correspond to the day on which the event occurred, and the hour, minutes and seconds are set to 00 in the key.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "simple_prefix": {
                    "computed": true,
                    "description": "This format defaults the prefix to the given log file prefix for delivering server access log file.",
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
      "metadata_table_configuration": {
        "computed": true,
        "description": "The metadata table configuration of an S3 general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html) and [Setting up permissions for configuring metadata tables](https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_tables_destination": {
              "computed": true,
              "description": "The destination information for the metadata table configuration. The destination table bucket must be in the same Region and AWS-account as the general purpose bucket. The specified metadata table name must be unique within the ` + "`" + `` + "`" + `aws_s3_metadata` + "`" + `` + "`" + ` namespace in the destination table bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "table_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) for the metadata table in the metadata table configuration. The specified metadata table name must be unique within the ` + "`" + `` + "`" + `aws_s3_metadata` + "`" + `` + "`" + ` namespace in the destination table bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_bucket_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) for the table bucket that's specified as the destination in the metadata table configuration. The destination table bucket must be in the same Region and AWS-account as the general purpose bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "The name for the metadata table in your metadata table configuration. The specified metadata table name must be unique within the ` + "`" + `` + "`" + `aws_s3_metadata` + "`" + `` + "`" + ` namespace in the destination table bucket.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_namespace": {
                    "computed": true,
                    "description": "The table bucket namespace for the metadata table in your metadata table configuration. This value is always ` + "`" + `` + "`" + `aws_s3_metadata` + "`" + `` + "`" + `.",
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
      "metrics_configurations": {
        "computed": true,
        "description": "Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_point_arn": {
              "computed": true,
              "description": "The access point that was used while performing operations on the object. The metrics configuration only includes objects that meet the filter's criteria.",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "The ID used to identify the metrics configuration. This can be any value you choose that helps you identify your metrics configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "prefix": {
              "computed": true,
              "description": "The prefix that an object must have to be included in the metrics results.",
              "description_kind": "plain",
              "type": "string"
            },
            "tag_filters": {
              "computed": true,
              "description": "Specifies a list of tag filters to use as a metrics configuration filter. The metrics configuration includes only objects that meet the filter's criteria.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The tag value.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "notification_configuration": {
        "computed": true,
        "description": "Configuration that defines how Amazon S3 handles bucket notifications.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_bridge_configuration": {
              "computed": true,
              "description": "Enables delivery of events to Amazon EventBridge.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event_bridge_enabled": {
                    "computed": true,
                    "description": "Enables delivery of events to Amazon EventBridge.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "lambda_configurations": {
              "computed": true,
              "description": "Describes the LAMlong functions to invoke and the events for which to invoke them.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event": {
                    "computed": true,
                    "description": "The Amazon S3 bucket event for which to invoke the LAMlong function. For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The filtering rules that determine which objects invoke the AWS Lambda function. For example, you can create a filter so that only image files with a ` + "`" + `` + "`" + `.jpg` + "`" + `` + "`" + ` extension invoke the function when they are added to the Amazon S3 bucket.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_key": {
                          "computed": true,
                          "description": "A container for object key name prefix and suffix filtering rules.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rules": {
                                "computed": true,
                                "description": "A list of containers for the key-value pair that defines the criteria for the filter rule.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description": "The object key name prefix or suffix identifying one or more objects to which the filtering rule applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The value that the filter searches for in object key names.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "set"
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
                  "function": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the LAMlong function that Amazon S3 invokes when the specified event type occurs.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "queue_configurations": {
              "computed": true,
              "description": "The Amazon Simple Queue Service queues to publish messages to and the events for which to publish messages.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event": {
                    "computed": true,
                    "description": "The Amazon S3 bucket event about which you want to publish messages to Amazon SQS. For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The filtering rules that determine which objects trigger notifications. For example, you can create a filter so that Amazon S3 sends notifications only when image files with a ` + "`" + `` + "`" + `.jpg` + "`" + `` + "`" + ` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/user-guide/notification-how-to-filtering.html) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_key": {
                          "computed": true,
                          "description": "A container for object key name prefix and suffix filtering rules.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rules": {
                                "computed": true,
                                "description": "A list of containers for the key-value pair that defines the criteria for the filter rule.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description": "The object key name prefix or suffix identifying one or more objects to which the filtering rule applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The value that the filter searches for in object key names.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "set"
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
                  "queue": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type. FIFO queues are not allowed when enabling an SQS queue as the event notification destination.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "topic_configurations": {
              "computed": true,
              "description": "The topic to which notifications are sent and the events for which notifications are generated.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event": {
                    "computed": true,
                    "description": "The Amazon S3 bucket event about which to send notifications. For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The filtering rules that determine for which objects to send notifications. For example, you can create a filter so that Amazon S3 sends notifications only when image files with a ` + "`" + `` + "`" + `.jpg` + "`" + `` + "`" + ` extension are added to the bucket.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_key": {
                          "computed": true,
                          "description": "A container for object key name prefix and suffix filtering rules.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "rules": {
                                "computed": true,
                                "description": "A list of containers for the key-value pair that defines the criteria for the filter rule.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description": "The object key name prefix or suffix identifying one or more objects to which the filtering rule applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide*.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The value that the filter searches for in object key names.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "set"
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
                  "topic": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.",
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
      "object_lock_configuration": {
        "computed": true,
        "description": "This operation is not supported for directory buckets.\n  Places an Object Lock configuration on the specified bucket. The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html). \n   +  The ` + "`" + `` + "`" + `DefaultRetention` + "`" + `` + "`" + ` settings require both a mode and a period.\n  +  The ` + "`" + `` + "`" + `DefaultRetention` + "`" + `` + "`" + ` period can be either ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + ` but you must select one. You cannot specify ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + ` at the same time.\n  +  You can enable Object Lock for new or existing buckets. For more information, see [Configuring Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "object_lock_enabled": {
              "computed": true,
              "description": "Indicates whether this bucket has an Object Lock configuration enabled. Enable ` + "`" + `` + "`" + `ObjectLockEnabled` + "`" + `` + "`" + ` when you apply ` + "`" + `` + "`" + `ObjectLockConfiguration` + "`" + `` + "`" + ` to a bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "rule": {
              "computed": true,
              "description": "Specifies the Object Lock rule for the specified object. Enable this rule when you apply ` + "`" + `` + "`" + `ObjectLockConfiguration` + "`" + `` + "`" + ` to a bucket. If Object Lock is turned on, bucket settings require both ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` and a period of either ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + `. You cannot specify ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + ` at the same time. For more information, see [ObjectLockRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html) and [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_retention": {
                    "computed": true,
                    "description": "The default Object Lock retention mode and period that you want to apply to new objects placed in the specified bucket. If Object Lock is turned on, bucket settings require both ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` and a period of either ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + `. You cannot specify ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + ` at the same time. For more information about allowable values for mode and period, see [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days": {
                          "computed": true,
                          "description": "The number of days that you want to specify for the default retention period. If Object Lock is turned on, you must specify ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` and specify either ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "mode": {
                          "computed": true,
                          "description": "The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. If Object Lock is turned on, you must specify ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` and specify either ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "years": {
                          "computed": true,
                          "description": "The number of years that you want to specify for the default retention period. If Object Lock is turned on, you must specify ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` and specify either ` + "`" + `` + "`" + `Days` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Years` + "`" + `` + "`" + `.",
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
      "object_lock_enabled": {
        "computed": true,
        "description": "Indicates whether this bucket has an Object Lock configuration enabled. Enable ` + "`" + `` + "`" + `ObjectLockEnabled` + "`" + `` + "`" + ` when you apply ` + "`" + `` + "`" + `ObjectLockConfiguration` + "`" + `` + "`" + ` to a bucket.",
        "description_kind": "plain",
        "type": "bool"
      },
      "ownership_controls": {
        "computed": true,
        "description": "Configuration that defines how Amazon S3 handles Object Ownership rules.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "Specifies the container element for Object Ownership rules.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "object_ownership": {
                    "computed": true,
                    "description": "Specifies an object ownership rule.",
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
      "public_access_block_configuration": {
        "computed": true,
        "description": "Configuration that defines how Amazon S3 handles public access.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket. Setting this element to ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + ` causes the following behavior:\n  +  PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.\n  +  PUT Object calls fail if the request includes a public ACL.\n  +  PUT Bucket calls fail if the request includes a public ACL.\n  \n Enabling this setting doesn't affect existing policies or ACLs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "block_public_policy": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should block public bucket policies for this bucket. Setting this element to ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + ` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. \n Enabling this setting doesn't affect existing bucket policies.",
              "description_kind": "plain",
              "type": "bool"
            },
            "ignore_public_acls": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting this element to ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + ` causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.\n Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
              "description_kind": "plain",
              "type": "bool"
            },
            "restrict_public_buckets": {
              "computed": true,
              "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + ` restricts access to this bucket to only AWS-service principals and authorized users within this account if the bucket has a public policy.\n Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "regional_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "replication_configuration": {
        "computed": true,
        "description": "Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the ` + "`" + `` + "`" + `VersioningConfiguration` + "`" + `` + "`" + ` property.\n Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAMlong (IAM) role that Amazon S3 assumes when replicating objects. For more information, see [How to Set Up Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-how-setup.html) in the *Amazon S3 User Guide*.",
              "description_kind": "plain",
              "type": "string"
            },
            "rules": {
              "computed": true,
              "description": "A container for one or more replication rules. A replication configuration must have at least one rule and can contain a maximum of 1,000 rules.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delete_marker_replication": {
                    "computed": true,
                    "description": "Specifies whether Amazon S3 replicates delete markers. If you specify a ` + "`" + `` + "`" + `Filter` + "`" + `` + "`" + ` in your replication configuration, you must also include a ` + "`" + `` + "`" + `DeleteMarkerReplication` + "`" + `` + "`" + ` element. If your ` + "`" + `` + "`" + `Filter` + "`" + `` + "`" + ` includes a ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` element, the ` + "`" + `` + "`" + `DeleteMarkerReplication` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `Status` + "`" + `` + "`" + ` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config). \n For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html). \n  If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "status": {
                          "computed": true,
                          "description": "Indicates whether to replicate delete markers. Disabled by default.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "destination": {
                    "computed": true,
                    "description": "A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_control_translation": {
                          "computed": true,
                          "description": "Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS-account that owns the destination bucket. If this is not specified in the replication configuration, the replicas are owned by same AWS-account that owns the source object.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "owner": {
                                "computed": true,
                                "description": "Specifies the replica ownership. For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) in the *Amazon S3 API Reference*.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "account": {
                          "computed": true,
                          "description": "Destination bucket owner account ID. In a cross-account scenario, if you direct Amazon S3 to change replica ownership to the AWS-account that owns the destination bucket by specifying the ` + "`" + `` + "`" + `AccessControlTranslation` + "`" + `` + "`" + ` property, this is the account ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional Configuration: Change Replica Owner](https://docs.aws.amazon.com/AmazonS3/latest/dev/crr-change-owner.html) in the *Amazon S3 User Guide*.\n If you specify the ` + "`" + `` + "`" + `AccessControlTranslation` + "`" + `` + "`" + ` property, the ` + "`" + `` + "`" + `Account` + "`" + `` + "`" + ` property is required.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bucket": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "encryption_configuration": {
                          "computed": true,
                          "description": "Specifies encryption-related information.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "replica_kms_key_id": {
                                "computed": true,
                                "description": "Specifies the ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket. Amazon S3 uses this key to encrypt replica objects. Amazon S3 only supports symmetric encryption KMS keys. For more information, see [Asymmetric keys in KMS](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *Key Management Service Developer Guide*.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "metrics": {
                          "computed": true,
                          "description": "A container specifying replication metrics-related settings enabling replication metrics and events.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "event_threshold": {
                                "computed": true,
                                "description": "A container specifying the time threshold for emitting the ` + "`" + `` + "`" + `s3:Replication:OperationMissedThreshold` + "`" + `` + "`" + ` event.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "minutes": {
                                      "computed": true,
                                      "description": "Contains an integer specifying time in minutes. \n  Valid value: 15",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "status": {
                                "computed": true,
                                "description": "Specifies whether the replication metrics are enabled.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "replication_time": {
                          "computed": true,
                          "description": "A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a ` + "`" + `` + "`" + `Metrics` + "`" + `` + "`" + ` block.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status": {
                                "computed": true,
                                "description": "Specifies whether the replication time is enabled.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "time": {
                                "computed": true,
                                "description": "A container specifying the time by which replication should be complete for all objects and operations on objects.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "minutes": {
                                      "computed": true,
                                      "description": "Contains an integer specifying time in minutes. \n  Valid value: 15",
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
                        "storage_class": {
                          "computed": true,
                          "description": "The storage class to use when replicating objects, such as S3 Standard or reduced redundancy. By default, Amazon S3 uses the storage class of the source object to create the object replica. \n For valid values, see the ` + "`" + `` + "`" + `StorageClass` + "`" + `` + "`" + ` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference*.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "filter": {
                    "computed": true,
                    "description": "A filter that identifies the subset of objects to which the replication rule applies. A ` + "`" + `` + "`" + `Filter` + "`" + `` + "`" + ` must specify exactly one ` + "`" + `` + "`" + `Prefix` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `TagFilter` + "`" + `` + "`" + `, or an ` + "`" + `` + "`" + `And` + "`" + `` + "`" + ` child element. The use of the filter field indicates that this is a V2 replication configuration. This field isn't supported in a V1 replication configuration.\n  V1 replication configuration only supports filtering by key prefix. To filter using a V1 replication configuration, add the ` + "`" + `` + "`" + `Prefix` + "`" + `` + "`" + ` directly as a child element of the ` + "`" + `` + "`" + `Rule` + "`" + `` + "`" + ` element.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "and": {
                          "computed": true,
                          "description": "A container for specifying rule filters. The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter. For example: \n  +  If you specify both a ` + "`" + `` + "`" + `Prefix` + "`" + `` + "`" + ` and a ` + "`" + `` + "`" + `TagFilter` + "`" + `` + "`" + `, wrap these filters in an ` + "`" + `` + "`" + `And` + "`" + `` + "`" + ` tag.\n  +  If you specify a filter based on multiple tags, wrap the ` + "`" + `` + "`" + `TagFilter` + "`" + `` + "`" + ` elements in an ` + "`" + `` + "`" + `And` + "`" + `` + "`" + ` tag.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prefix": {
                                "computed": true,
                                "description": "An object key name prefix that identifies the subset of objects to which the rule applies.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "tag_filters": {
                                "computed": true,
                                "description": "An array of tags containing key and value pairs.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "key": {
                                      "computed": true,
                                      "description": "The tag key.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The tag value.",
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
                        "prefix": {
                          "computed": true,
                          "description": "An object key name prefix that identifies the subset of objects to which the rule applies.\n  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tag_filter": {
                          "computed": true,
                          "description": "A container for specifying a tag key and value. \n The rule applies only to objects that have the tag in their tag set.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description": "The tag key.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The tag value.",
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
                  "id": {
                    "computed": true,
                    "description": "A unique identifier for the rule. The maximum value is 255 characters. If you don't specify a value, AWS CloudFormation generates a random ID. When using a V2 replication configuration this property is capitalized as \"ID\".",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description": "An object key name prefix that identifies the object or objects to which the rule applies. The maximum prefix length is 1,024 characters. To include all objects in a bucket, specify an empty string. To filter using a V1 replication configuration, add the ` + "`" + `` + "`" + `Prefix` + "`" + `` + "`" + ` directly as a child element of the ` + "`" + `` + "`" + `Rule` + "`" + `` + "`" + ` element.\n  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "priority": {
                    "computed": true,
                    "description": "The priority indicates which rule has precedence whenever two or more replication rules conflict. Amazon S3 will attempt to replicate objects according to all replication rules. However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority. The higher the number, the higher the priority. \n For more information, see [Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the *Amazon S3 User Guide*.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "source_selection_criteria": {
                    "computed": true,
                    "description": "A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects.",
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
                                "computed": true,
                                "description": "Specifies whether Amazon S3 replicates modifications on replicas.\n  *Allowed values*: ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `Disabled` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "sse_kms_encrypted_objects": {
                          "computed": true,
                          "description": "A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "status": {
                                "computed": true,
                                "description": "Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.",
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
                  "status": {
                    "computed": true,
                    "description": "Specifies whether the rule is enabled.",
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
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Name of the object key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "versioning_configuration": {
        "computed": true,
        "description": "Enables multiple versions of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.\n  When you enable versioning on a bucket for the first time, it might take a short amount of time for the change to be fully propagated. We recommend that you wait for 15 minutes after enabling versioning before issuing write operations (` + "`" + `` + "`" + `PUT` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DELETE` + "`" + `` + "`" + `) on objects in the bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "The versioning state of the bucket.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "website_configuration": {
        "computed": true,
        "description": "Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "error_document": {
              "computed": true,
              "description": "The name of the error document for the website.",
              "description_kind": "plain",
              "type": "string"
            },
            "index_document": {
              "computed": true,
              "description": "The name of the index document for the website.",
              "description_kind": "plain",
              "type": "string"
            },
            "redirect_all_requests_to": {
              "computed": true,
              "description": "The redirect behavior for every request to this bucket's website endpoint.\n  If you specify this property, you can't specify any other property.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "host_name": {
                    "computed": true,
                    "description": "Name of the host where requests are redirected.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "Protocol to use when redirecting requests. The default is the protocol that is used in the original request.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "routing_rules": {
              "computed": true,
              "description": "Rules that define when a redirect is applied and the redirect behavior.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "redirect_rule": {
                    "computed": true,
                    "description": "Container for redirect information. You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "host_name": {
                          "computed": true,
                          "description": "The host name to use in the redirect request.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "http_redirect_code": {
                          "computed": true,
                          "description": "The HTTP redirect code to use on the response. Not required if one of the siblings is present.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "protocol": {
                          "computed": true,
                          "description": "Protocol to use when redirecting requests. The default is the protocol that is used in the original request.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "replace_key_prefix_with": {
                          "computed": true,
                          "description": "The object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix ` + "`" + `` + "`" + `docs/` + "`" + `` + "`" + ` (objects in the ` + "`" + `` + "`" + `docs/` + "`" + `` + "`" + ` folder) to ` + "`" + `` + "`" + `documents/` + "`" + `` + "`" + `, you can set a condition block with ` + "`" + `` + "`" + `KeyPrefixEquals` + "`" + `` + "`" + ` set to ` + "`" + `` + "`" + `docs/` + "`" + `` + "`" + ` and in the Redirect set ` + "`" + `` + "`" + `ReplaceKeyPrefixWith` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `/documents` + "`" + `` + "`" + `. Not required if one of the siblings is present. Can be present only if ` + "`" + `` + "`" + `ReplaceKeyWith` + "`" + `` + "`" + ` is not provided.\n  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "replace_key_with": {
                          "computed": true,
                          "description": "The specific object key to use in the redirect request. For example, redirect request to ` + "`" + `` + "`" + `error.html` + "`" + `` + "`" + `. Not required if one of the siblings is present. Can be present only if ` + "`" + `` + "`" + `ReplaceKeyPrefixWith` + "`" + `` + "`" + ` is not provided.\n  Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "routing_rule_condition": {
                    "computed": true,
                    "description": "A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the ` + "`" + `` + "`" + `/docs` + "`" + `` + "`" + ` folder, redirect to the ` + "`" + `` + "`" + `/documents` + "`" + `` + "`" + ` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "http_error_code_returned_equals": {
                          "computed": true,
                          "description": "The HTTP error code when the redirect is applied. In the event of an error, if the error code equals this value, then the specified redirect is applied.\n Required when parent element ` + "`" + `` + "`" + `Condition` + "`" + `` + "`" + ` is specified and sibling ` + "`" + `` + "`" + `KeyPrefixEquals` + "`" + `` + "`" + ` is not specified. If both are specified, then both must be true for the redirect to be applied.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key_prefix_equals": {
                          "computed": true,
                          "description": "The object key name prefix when the redirect is applied. For example, to redirect requests for ` + "`" + `` + "`" + `ExamplePage.html` + "`" + `` + "`" + `, the key prefix will be ` + "`" + `` + "`" + `ExamplePage.html` + "`" + `` + "`" + `. To redirect request for all pages with the prefix ` + "`" + `` + "`" + `docs/` + "`" + `` + "`" + `, the key prefix will be ` + "`" + `` + "`" + `/docs` + "`" + `` + "`" + `, which identifies all objects in the docs/ folder.\n Required when the parent element ` + "`" + `` + "`" + `Condition` + "`" + `` + "`" + ` is specified and sibling ` + "`" + `` + "`" + `HttpErrorCodeReturnedEquals` + "`" + `` + "`" + ` is not specified. If both conditions are specified, both must be true for the redirect to be applied.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "website_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3::Bucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3BucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3Bucket), &result)
	return &result
}
