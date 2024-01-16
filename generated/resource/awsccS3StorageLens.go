package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3StorageLens = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_lens_configuration": {
        "description": "Specifies the details of Amazon S3 Storage Lens configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_level": {
              "description": "Account-level metrics configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "activity_metrics": {
                    "computed": true,
                    "description": "Enables activity metrics.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_enabled": {
                          "computed": true,
                          "description": "Specifies whether activity metrics are enabled or disabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "advanced_cost_optimization_metrics": {
                    "computed": true,
                    "description": "Enables advanced cost optimization metrics.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_enabled": {
                          "computed": true,
                          "description": "Specifies whether advanced cost optimization metrics are enabled or disabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "advanced_data_protection_metrics": {
                    "computed": true,
                    "description": "Enables advanced data protection metrics.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_enabled": {
                          "computed": true,
                          "description": "Specifies whether advanced data protection metrics are enabled or disabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "bucket_level": {
                    "description": "Bucket-level metrics configurations.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "activity_metrics": {
                          "computed": true,
                          "description": "Enables activity metrics.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "is_enabled": {
                                "computed": true,
                                "description": "Specifies whether activity metrics are enabled or disabled.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "advanced_cost_optimization_metrics": {
                          "computed": true,
                          "description": "Enables advanced cost optimization metrics.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "is_enabled": {
                                "computed": true,
                                "description": "Specifies whether advanced cost optimization metrics are enabled or disabled.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "advanced_data_protection_metrics": {
                          "computed": true,
                          "description": "Enables advanced data protection metrics.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "is_enabled": {
                                "computed": true,
                                "description": "Specifies whether advanced data protection metrics are enabled or disabled.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "detailed_status_codes_metrics": {
                          "computed": true,
                          "description": "Enables detailed status codes metrics.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "is_enabled": {
                                "computed": true,
                                "description": "Specifies whether detailed status codes metrics are enabled or disabled.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "prefix_level": {
                          "computed": true,
                          "description": "Prefix-level metrics configurations.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "storage_metrics": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "is_enabled": {
                                      "computed": true,
                                      "description": "Specifies whether prefix-level storage metrics are enabled or disabled.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "selection_criteria": {
                                      "computed": true,
                                      "description": "Selection criteria for prefix-level metrics.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "delimiter": {
                                            "computed": true,
                                            "description": "Delimiter to divide S3 key into hierarchy of prefixes.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "max_depth": {
                                            "computed": true,
                                            "description": "Max depth of prefixes of S3 key that Amazon S3 Storage Lens will analyze.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "min_storage_bytes_percentage": {
                                            "computed": true,
                                            "description": "The minimum storage bytes threshold for the prefixes to be included in the analysis.",
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
                                "required": true
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
                  "detailed_status_codes_metrics": {
                    "computed": true,
                    "description": "Enables detailed status codes metrics.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_enabled": {
                          "computed": true,
                          "description": "Specifies whether detailed status codes metrics are enabled or disabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "storage_lens_group_level": {
                    "computed": true,
                    "description": "Specifies the details of Amazon S3 Storage Lens Group configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_lens_group_selection_criteria": {
                          "computed": true,
                          "description": "Selection criteria for Storage Lens Group level metrics",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "exclude": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "include": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "set",
                                  "string"
                                ]
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
              "required": true
            },
            "aws_org": {
              "computed": true,
              "description": "The AWS Organizations ARN to use in the Amazon S3 Storage Lens configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "description": "The Amazon Resource Name (ARN) of the specified resource.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "data_export": {
              "computed": true,
              "description": "Specifies how Amazon S3 Storage Lens metrics should be exported.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_metrics": {
                    "computed": true,
                    "description": "CloudWatch metrics settings for the Amazon S3 Storage Lens metrics export.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_enabled": {
                          "description": "Specifies whether CloudWatch metrics are enabled or disabled.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "s3_bucket_destination": {
                    "computed": true,
                    "description": "S3 bucket destination settings for the Amazon S3 Storage Lens metrics export.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "account_id": {
                          "description": "The AWS account ID that owns the destination S3 bucket.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "arn": {
                          "description": "The ARN of the bucket to which Amazon S3 Storage Lens exports will be placed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "encryption": {
                          "computed": true,
                          "description": "Configures the server-side encryption for Amazon S3 Storage Lens report files with either S3-managed keys (SSE-S3) or KMS-managed keys (SSE-KMS).",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "ssekms": {
                                "computed": true,
                                "description": "AWS KMS server-side encryption.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "key_id": {
                                      "description": "The ARN of the KMS key to use for encryption.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "sses3": {
                                "computed": true,
                                "description": "S3 default server-side encryption.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "format": {
                          "description": "Specifies the file format to use when exporting Amazon S3 Storage Lens metrics export.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "output_schema_version": {
                          "description": "The version of the output schema to use when exporting Amazon S3 Storage Lens metrics.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "prefix": {
                          "computed": true,
                          "description": "The prefix to use for Amazon S3 Storage Lens export.",
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
            "exclude": {
              "computed": true,
              "description": "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "buckets": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "regions": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "id": {
              "description": "The ID that identifies the Amazon S3 Storage Lens configuration.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "include": {
              "computed": true,
              "description": "S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "buckets": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "regions": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "is_enabled": {
              "description": "Specifies whether the Amazon S3 Storage Lens configuration is enabled or disabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "storage_lens_arn": {
              "computed": true,
              "description": "The ARN for the Amazon S3 Storage Lens configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::S3::StorageLens resource is an Amazon S3 resource type that you can use to create Storage Lens configurations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3StorageLensSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3StorageLens), &result)
	return &result
}
