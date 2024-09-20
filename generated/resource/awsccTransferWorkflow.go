package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferWorkflow = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A textual description for the workflow.",
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
      "on_exception_steps": {
        "computed": true,
        "description": "Specifies the steps (actions) to take if any errors are encountered during execution of the workflow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_step_details": {
              "computed": true,
              "description": "Details for a step that performs a file copy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_file_location": {
                    "computed": true,
                    "description": "Specifies the location for the file being copied. Only applicable for the Copy type of workflow steps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_file_location": {
                          "computed": true,
                          "description": "Specifies the details for a S3 file.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description": "Specifies the S3 bucket that contains the file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "computed": true,
                                "description": "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
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
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "overwrite_existing": {
                    "computed": true,
                    "description": "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "custom_step_details": {
              "computed": true,
              "description": "Details for a step that invokes a lambda function.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target": {
                    "computed": true,
                    "description": "The ARN for the lambda function that is being called.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "computed": true,
                    "description": "Timeout, in seconds, for the step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "decrypt_step_details": {
              "computed": true,
              "description": "Details for a step that performs a file decryption.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_file_location": {
                    "computed": true,
                    "description": "Specifies the location for the file being decrypted. Only applicable for the Decrypt type of workflow steps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "efs_file_location": {
                          "computed": true,
                          "description": "Specifies the details for an EFS file.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "file_system_id": {
                                "computed": true,
                                "description": "Specifies the EFS filesystem that contains the file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "computed": true,
                                "description": "The name assigned to the file when it was created in EFS. You use the object path to retrieve the object.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "s3_file_location": {
                          "computed": true,
                          "description": "Specifies the details for a S3 file.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description": "Specifies the S3 bucket that contains the file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "computed": true,
                                "description": "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
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
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "overwrite_existing": {
                    "computed": true,
                    "description": "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Specifies which encryption method to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "delete_step_details": {
              "computed": true,
              "description": "Details for a step that deletes the file.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "tag_step_details": {
              "computed": true,
              "description": "Details for a step that creates one or more tags.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "Array that contains from 1 to 10 key/value pairs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The name assigned to the tag that you create.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value that corresponds to the key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
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
      "steps": {
        "description": "Specifies the details for the steps that are in the specified workflow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_step_details": {
              "computed": true,
              "description": "Details for a step that performs a file copy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_file_location": {
                    "computed": true,
                    "description": "Specifies the location for the file being copied. Only applicable for the Copy type of workflow steps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_file_location": {
                          "computed": true,
                          "description": "Specifies the details for a S3 file.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description": "Specifies the S3 bucket that contains the file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "computed": true,
                                "description": "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
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
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "overwrite_existing": {
                    "computed": true,
                    "description": "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "custom_step_details": {
              "computed": true,
              "description": "Details for a step that invokes a lambda function.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target": {
                    "computed": true,
                    "description": "The ARN for the lambda function that is being called.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "computed": true,
                    "description": "Timeout, in seconds, for the step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "decrypt_step_details": {
              "computed": true,
              "description": "Details for a step that performs a file decryption.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_file_location": {
                    "computed": true,
                    "description": "Specifies the location for the file being decrypted. Only applicable for the Decrypt type of workflow steps.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "efs_file_location": {
                          "computed": true,
                          "description": "Specifies the details for an EFS file.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "file_system_id": {
                                "computed": true,
                                "description": "Specifies the EFS filesystem that contains the file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "path": {
                                "computed": true,
                                "description": "The name assigned to the file when it was created in EFS. You use the object path to retrieve the object.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "s3_file_location": {
                          "computed": true,
                          "description": "Specifies the details for a S3 file.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description": "Specifies the S3 bucket that contains the file.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key": {
                                "computed": true,
                                "description": "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
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
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "overwrite_existing": {
                    "computed": true,
                    "description": "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Specifies which encryption method to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "delete_step_details": {
              "computed": true,
              "description": "Details for a step that deletes the file.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "tag_step_details": {
              "computed": true,
              "description": "Details for a step that creates one or more tags.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the step, used as an identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "source_file_location": {
                    "computed": true,
                    "description": "Specifies which file to use as input to the workflow step.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "computed": true,
                    "description": "Array that contains from 1 to 10 key/value pairs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The name assigned to the tag that you create.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value that corresponds to the key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name assigned to the tag that you create.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "workflow_id": {
        "computed": true,
        "description": "A unique identifier for the workflow.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Transfer::Workflow",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTransferWorkflowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferWorkflow), &result)
	return &result
}
