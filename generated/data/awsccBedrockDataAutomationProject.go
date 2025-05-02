package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockDataAutomationProject = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_output_configuration": {
        "computed": true,
        "description": "Custom output configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "blueprints": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "blueprint_arn": {
                    "computed": true,
                    "description": "ARN of a Blueprint",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "blueprint_stage": {
                    "computed": true,
                    "description": "Stage of the Blueprint",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "blueprint_version": {
                    "computed": true,
                    "description": "Blueprint Version",
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_encryption_context": {
        "computed": true,
        "description": "KMS encryption context",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "kms_key_id": {
        "computed": true,
        "description": "KMS key identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "override_configuration": {
        "computed": true,
        "description": "Override configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audio": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "modality_processing": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
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
            "document": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "modality_processing": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "splitter": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
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
            "image": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "modality_processing": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
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
            "modality_routing": {
              "computed": true,
              "description": "Modality routing configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "jpeg": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mov": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mp_4": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "png": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "video": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "modality_processing": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "project_arn": {
        "computed": true,
        "description": "ARN of a DataAutomationProject",
        "description_kind": "plain",
        "type": "string"
      },
      "project_description": {
        "computed": true,
        "description": "Description of the DataAutomationProject",
        "description_kind": "plain",
        "type": "string"
      },
      "project_name": {
        "computed": true,
        "description": "Name of the DataAutomationProject",
        "description_kind": "plain",
        "type": "string"
      },
      "project_stage": {
        "computed": true,
        "description": "Stage of the Project",
        "description_kind": "plain",
        "type": "string"
      },
      "standard_output_configuration": {
        "computed": true,
        "description": "Standard output configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "audio": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "extraction": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "category": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "types": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "generative_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "types": {
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
                "nesting_mode": "single"
              }
            },
            "document": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "extraction": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bounding_box": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "granularity": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "types": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "generative_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "output_format": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "additional_file_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "text_format": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "types": {
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
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "extraction": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bounding_box": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "category": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "types": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "generative_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "types": {
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
                "nesting_mode": "single"
              }
            },
            "video": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "extraction": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bounding_box": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "category": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "state": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "types": {
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
                      "nesting_mode": "single"
                    }
                  },
                  "generative_field": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "state": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "types": {
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
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Key for the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Bedrock::DataAutomationProject",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockDataAutomationProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockDataAutomationProject), &result)
	return &result
}
