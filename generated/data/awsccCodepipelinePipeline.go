package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodepipelinePipeline = `{
  "block": {
    "attributes": {
      "artifact_store": {
        "computed": true,
        "description": "The S3 bucket where artifacts for the pipeline are stored.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_key": {
              "computed": true,
              "description": "Represents information about the key used to encrypt data in the artifact store, such as an AWS Key Management Service (AWS KMS) key",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "The ID used to identify the key. For an AWS KMS key, you can use the key ID, the key ARN, or the alias ARN.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of encryption key, such as an AWS KMS key. When creating or updating a pipeline, the value must be set to 'KMS'.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "location": {
              "computed": true,
              "description": "The S3 bucket used for storing the artifacts for a pipeline. You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of the artifact store, such as S3.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "artifact_stores": {
        "computed": true,
        "description": "A mapping of artifactStore objects and their corresponding AWS Regions. There must be an artifact store for the pipeline Region and for each cross-region action in the pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "artifact_store": {
              "computed": true,
              "description": "The S3 bucket where artifacts for the pipeline are stored.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encryption_key": {
                    "computed": true,
                    "description": "Represents information about the key used to encrypt data in the artifact store, such as an AWS Key Management Service (AWS KMS) key",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "id": {
                          "computed": true,
                          "description": "The ID used to identify the key. For an AWS KMS key, you can use the key ID, the key ARN, or the alias ARN.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The type of encryption key, such as an AWS KMS key. When creating or updating a pipeline, the value must be set to 'KMS'.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "location": {
                    "computed": true,
                    "description": "The S3 bucket used for storing the artifacts for a pipeline. You can specify the name of an S3 bucket but not a folder in the bucket. A folder to contain the pipeline artifacts is created for you based on the name of the pipeline. You can use any S3 bucket in the same AWS Region as the pipeline to store your pipeline artifacts.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of the artifact store, such as S3.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "region": {
              "computed": true,
              "description": "The action declaration's AWS Region, such as us-east-1.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "disable_inbound_stage_transitions": {
        "computed": true,
        "description": "Represents the input of a DisableStageTransition action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reason": {
              "computed": true,
              "description": "The reason given to the user that a stage is disabled, such as waiting for manual approval or manual tests. This message is displayed in the pipeline console UI.",
              "description_kind": "plain",
              "type": "string"
            },
            "stage_name": {
              "computed": true,
              "description": "The name of the stage where you want to disable the inbound or outbound transition of artifacts.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "execution_mode": {
        "computed": true,
        "description": "The method that the pipeline will use to handle multiple executions. The default mode is SUPERSEDED.",
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
        "description": "The name of the pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_type": {
        "computed": true,
        "description": "CodePipeline provides the following pipeline types, which differ in characteristics and price, so that you can tailor your pipeline features and cost to the needs of your applications.",
        "description_kind": "plain",
        "type": "string"
      },
      "restart_execution_on_update": {
        "computed": true,
        "description": "Indicates whether to rerun the CodePipeline pipeline after you update it.",
        "description_kind": "plain",
        "type": "bool"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for CodePipeline to use to either perform actions with no actionRoleArn, or to use to assume roles for actions with an actionRoleArn",
        "description_kind": "plain",
        "type": "string"
      },
      "stages": {
        "computed": true,
        "description": "Represents information about a stage and its definition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action_type_id": {
                    "computed": true,
                    "description": "Represents information about an action type.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "category": {
                          "computed": true,
                          "description": "A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Valid categories are limited to one of the values below.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "owner": {
                          "computed": true,
                          "description": "The creator of the action being called. There are three valid values for the Owner field in the action category section within your pipeline structure: AWS, ThirdParty, and Custom.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "provider": {
                          "computed": true,
                          "description": "The provider of the service being called by the action. Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of CodeDeploy, which would be specified as CodeDeploy.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description": "A string that describes the action version.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "configuration": {
                    "computed": true,
                    "description": "The action's configuration. These are key-value pairs that specify input values for an action.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_artifacts": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the artifact to be worked on (for example, \"My App\").",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "name": {
                    "computed": true,
                    "description": "The action declaration's name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "namespace": {
                    "computed": true,
                    "description": "The variable namespace associated with the action. All variables produced as output by this action fall under this namespace.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_artifacts": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the output of an artifact, such as \"My App\".",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "region": {
                    "computed": true,
                    "description": "The action declaration's AWS Region, such as us-east-1.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The ARN of the IAM service role that performs the declared action. This is assumed through the roleArn for the pipeline.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "run_order": {
                    "computed": true,
                    "description": "The order in which actions are run.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "timeout_in_minutes": {
                    "computed": true,
                    "description": "A timeout duration in minutes that can be applied against the ActionType?s default timeout value specified in Quotas for AWS CodePipeline. This attribute is available only to the manual approval ActionType.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "blockers": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "Reserved for future use.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Reserved for future use.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "name": {
              "computed": true,
              "description": "The name of the stage.",
              "description_kind": "plain",
              "type": "string"
            },
            "on_failure": {
              "computed": true,
              "description": "The method to use when a stage has not completed successfully",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "result": {
                    "computed": true,
                    "description": "The specified result for when the failure conditions are met, such as rolling back the stage",
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
      },
      "tags": {
        "computed": true,
        "description": "Specifies the tags applied to the pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "triggers": {
        "computed": true,
        "description": "The trigger configuration specifying a type of event, such as Git tags, that starts the pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "git_configuration": {
              "computed": true,
              "description": "A type of trigger configuration for Git-based source actions.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "pull_request": {
                    "computed": true,
                    "description": "The field where the repository event that will start the pipeline is specified as pull requests.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "branches": {
                          "computed": true,
                          "description": "The Git repository branches specified as filter criteria to start the pipeline.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "excludes": {
                                "computed": true,
                                "description": "The list of patterns of Git branches that, when a commit is pushed, are to be excluded from starting the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "includes": {
                                "computed": true,
                                "description": "The list of patterns of Git branches that, when a commit is pushed, are to be included as criteria that starts the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "events": {
                          "computed": true,
                          "description": "The field that specifies which pull request events to filter on (opened, updated, closed) for the trigger configuration.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "file_paths": {
                          "computed": true,
                          "description": "The Git repository file paths specified as filter criteria to start the pipeline.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "excludes": {
                                "computed": true,
                                "description": "The list of patterns of Git repository file paths that, when a commit is pushed, are to be excluded from starting the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "includes": {
                                "computed": true,
                                "description": "The list of patterns of Git repository file paths that, when a commit is pushed, are to be included as criteria that starts the pipeline.",
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
                      "nesting_mode": "list"
                    }
                  },
                  "push": {
                    "computed": true,
                    "description": "The field where the repository event that will start the pipeline, such as pushing Git tags, is specified with details.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "branches": {
                          "computed": true,
                          "description": "The Git repository branches specified as filter criteria to start the pipeline.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "excludes": {
                                "computed": true,
                                "description": "The list of patterns of Git branches that, when a commit is pushed, are to be excluded from starting the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "includes": {
                                "computed": true,
                                "description": "The list of patterns of Git branches that, when a commit is pushed, are to be included as criteria that starts the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "file_paths": {
                          "computed": true,
                          "description": "The Git repository file paths specified as filter criteria to start the pipeline.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "excludes": {
                                "computed": true,
                                "description": "The list of patterns of Git repository file paths that, when a commit is pushed, are to be excluded from starting the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "includes": {
                                "computed": true,
                                "description": "The list of patterns of Git repository file paths that, when a commit is pushed, are to be included as criteria that starts the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "tags": {
                          "computed": true,
                          "description": "The Git tags specified as filter criteria for whether a Git tag repository event will start the pipeline.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "excludes": {
                                "computed": true,
                                "description": "The list of patterns of Git tags that, when pushed, are to be excluded from starting the pipeline.",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "includes": {
                                "computed": true,
                                "description": "The list of patterns of Git tags that, when pushed, are to be included as criteria that starts the pipeline.",
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
                      "nesting_mode": "list"
                    }
                  },
                  "source_action_name": {
                    "computed": true,
                    "description": "The name of the pipeline source action where the trigger configuration, such as Git tags, is specified. The trigger configuration will start the pipeline upon the specified change only.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "provider_type": {
              "computed": true,
              "description": "The source provider for the event, such as connections configured for a repository with Git tags, for the specified trigger configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "variables": {
        "computed": true,
        "description": "A list that defines the pipeline variables for a pipeline resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9@\\-_]+.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description": "The value of a pipeline-level variable.",
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "The description of a pipeline-level variable. It's used to add additional context about the variable, and not being used at time when pipeline executes.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of a pipeline-level variable.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "version": {
        "computed": true,
        "description": "The version of the pipeline.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CodePipeline::Pipeline",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodepipelinePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodepipelinePipeline), &result)
	return &result
}
