package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodedeployDeploymentGroup = `{
  "block": {
    "attributes": {
      "alarm_configuration": {
        "computed": true,
        "description": "Information about the Amazon CloudWatch alarms that are associated with the deployment group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarms": {
              "computed": true,
              "description": "A list of alarms configured for the deployment or deployment group. A maximum of 10 alarms can be added.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the alarm. Maximum length is 255 characters. Each alarm name can be used only once in a list of alarms.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "enabled": {
              "computed": true,
              "description": "Indicates whether the alarm configuration is enabled.",
              "description_kind": "plain",
              "type": "bool"
            },
            "ignore_poll_alarm_failure": {
              "computed": true,
              "description": "Indicates whether a deployment should continue if information about the current state of alarms cannot be retrieved from Amazon CloudWatch. The default value is false.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "application_name": {
        "computed": true,
        "description": "The name of an existing CodeDeploy application to associate this deployment group with.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_rollback_configuration": {
        "computed": true,
        "description": "Information about the automatic rollback configuration that is associated with the deployment group. If you specify this property, don't specify the Deployment property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Indicates whether a defined automatic rollback configuration is currently enabled.",
              "description_kind": "plain",
              "type": "bool"
            },
            "events": {
              "computed": true,
              "description": "The event type or types that trigger a rollback.",
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
      "auto_scaling_groups": {
        "computed": true,
        "description": "A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created. Duplicates are not allowed.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "blue_green_deployment_configuration": {
        "computed": true,
        "description": "Information about blue/green deployment options for a deployment group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deployment_ready_option": {
              "computed": true,
              "description": "Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action_on_timeout": {
                    "computed": true,
                    "description": "Information about when to reroute traffic from an original environment to a replacement environment in a blue/green deployment. CONTINUE_DEPLOYMENT: Register new instances with the load balancer immediately after the new application revision is installed on the instances in the replacement environment. STOP_DEPLOYMENT: Do not register new instances with a load balancer unless traffic rerouting is started using ContinueDeployment . If traffic rerouting is not started before the end of the specified wait period, the deployment status is changed to Stopped.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "wait_time_in_minutes": {
                    "computed": true,
                    "description": "The number of minutes to wait before the status of a blue/green deployment is changed to Stopped if rerouting is not started manually. Applies only to the STOP_DEPLOYMENT option for actionOnTimeout.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "green_fleet_provisioning_option": {
              "computed": true,
              "description": "Information about how instances are provisioned for a replacement environment in a blue/green deployment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "The method used to add instances to a replacement environment.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "terminate_blue_instances_on_deployment_success": {
              "computed": true,
              "description": "Information about whether to terminate instances in the original fleet during a blue/green deployment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "The action to take on instances in the original environment after a successful blue/green deployment.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "termination_wait_time_in_minutes": {
                    "computed": true,
                    "description": "For an Amazon EC2 deployment, the number of minutes to wait after a successful blue/green deployment before terminating instances from the original environment. For an Amazon ECS deployment, the number of minutes before deleting the original (blue) task set. During an Amazon ECS deployment, CodeDeploy shifts traffic from the original (blue) task set to a replacement (green) task set. The maximum setting is 2880 minutes (2 days).",
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
      "deployment": {
        "computed": true,
        "description": "The application revision to deploy to this deployment group. If you specify this property, your target application revision is deployed as soon as the provisioning process is complete. If you specify this property, don't specify the AutoRollbackConfiguration property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description of the deployment.",
              "description_kind": "plain",
              "type": "string"
            },
            "ignore_application_stop_failures": {
              "computed": true,
              "description": "If true, then if an ApplicationStop, BeforeBlockTraffic, or AfterBlockTraffic deployment lifecycle event to an instance fails, then the deployment continues to the next deployment lifecycle event. If false or not specified, then if a lifecycle event fails during a deployment to an instance, that deployment fails. If deployment to that instance is part of an overall deployment and the number of healthy hosts is not less than the minimum number of healthy hosts, then a deployment to the next instance is attempted.",
              "description_kind": "plain",
              "type": "bool"
            },
            "revision": {
              "computed": true,
              "description": "Information about the location of stored application artifacts and the service from which to retrieve them.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "git_hub_location": {
                    "computed": true,
                    "description": "Specifies the location of an application revision that is stored in GitHub.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "commit_id": {
                          "computed": true,
                          "description": "The SHA1 commit ID of the GitHub commit that represents the bundled artifacts for the application revision.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "repository": {
                          "computed": true,
                          "description": "The GitHub account and repository pair that stores the application revision to be deployed.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "revision_type": {
                    "computed": true,
                    "description": "The type of application revision.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_location": {
                    "computed": true,
                    "description": "Information about the location of application artifacts stored in Amazon S3.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description": "The name of the Amazon S3 bucket where the application revision is stored.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "bundle_type": {
                          "computed": true,
                          "description": "The file type of the application revision.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "e_tag": {
                          "computed": true,
                          "description": "The ETag of the Amazon S3 object that represents the bundled artifacts for the application revision. If the ETag is not specified as an input parameter, ETag validation of the object is skipped.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
                          "computed": true,
                          "description": "The name of the Amazon S3 object that represents the bundled artifacts for the application revision.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description": "A specific version of the Amazon S3 object that represents the bundled artifacts for the application revision. If the version is not specified, the system uses the most recent version by default.",
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
      "deployment_config_name": {
        "computed": true,
        "description": "A deployment configuration name or a predefined configuration name. With predefined configurations, you can deploy application revisions to one instance at a time (CodeDeployDefault.OneAtATime), half of the instances at a time (CodeDeployDefault.HalfAtATime), or all the instances at once (CodeDeployDefault.AllAtOnce).",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_group_name": {
        "computed": true,
        "description": "A name for the deployment group. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment group name.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_style": {
        "computed": true,
        "description": "Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer. If you specify this property with a blue/green deployment type, don't specify the AutoScalingGroups, LoadBalancerInfo, or Deployment properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deployment_option": {
              "computed": true,
              "description": "Indicates whether to route deployment traffic behind a load balancer.",
              "description_kind": "plain",
              "type": "string"
            },
            "deployment_type": {
              "computed": true,
              "description": "Indicates whether to run an in-place or blue/green deployment.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "ec_2_tag_filters": {
        "computed": true,
        "description": "The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group. CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed. You can specify EC2TagFilters or Ec2TagSet, but not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag filter key.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The tag filter type.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag filter value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "ec_2_tag_set": {
        "computed": true,
        "description": "Information about groups of tags applied to Amazon EC2 instances. Use when the deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same call as ec2TagFilter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ec_2_tag_set_list": {
              "computed": true,
              "description": "The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group. CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ec_2_tag_group": {
                    "computed": true,
                    "description": "A list that contains other lists of Amazon EC2 instance tag groups. For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The tag filter key.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The tag filter type.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The tag filter value.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "ecs_services": {
        "computed": true,
        "description": "The target Amazon ECS services in the deployment group. This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format \u003cclustername\u003e:\u003cservicename\u003e.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cluster_name": {
              "computed": true,
              "description": "The name of the cluster that the Amazon ECS service is associated with.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_name": {
              "computed": true,
              "description": "The name of the target Amazon ECS service.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "load_balancer_info": {
        "computed": true,
        "description": "Information about the load balancer to use in a deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "elb_info_list": {
              "computed": true,
              "description": "An array that contains information about the load balancers to use for load balancing in a deployment. If you're using Classic Load Balancers, specify those load balancers in this array.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "For blue/green deployments, the name of the load balancer that is used to route traffic from original instances to replacement instances in a blue/green deployment. For in-place deployments, the name of the load balancer that instances are deregistered from so they are not serving traffic during a deployment, and then re-registered with after the deployment is complete.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "target_group_info_list": {
              "computed": true,
              "description": "An array that contains information about the target groups to use for load balancing in a deployment. If you're using Application Load Balancers and Network Load Balancers, specify their associated target groups in this array.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "For blue/green deployments, the name of the target group that instances in the original environment are deregistered from, and instances in the replacement environment registered with. For in-place deployments, the name of the target group that instances are deregistered from, so they are not serving traffic during a deployment, and then re-registered with after the deployment completes. No duplicates allowed.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "target_group_pair_info_list": {
              "computed": true,
              "description": "The target group pair information. This is an array of TargeGroupPairInfo objects with a maximum size of one.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "prod_traffic_route": {
                    "computed": true,
                    "description": "The path used by a load balancer to route production traffic when an Amazon ECS deployment is complete.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "listener_arns": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of one listener. The listener identifies the route between a target group and a load balancer. This is an array of strings with a maximum size of one.",
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
                  "target_groups": {
                    "computed": true,
                    "description": "One pair of target groups. One is associated with the original task set. The second is associated with the task set that serves traffic after the deployment is complete.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "For blue/green deployments, the name of the target group that instances in the original environment are deregistered from, and instances in the replacement environment registered with. For in-place deployments, the name of the target group that instances are deregistered from, so they are not serving traffic during a deployment, and then re-registered with after the deployment completes. No duplicates allowed.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "test_traffic_route": {
                    "computed": true,
                    "description": "An optional path used by a load balancer to route test traffic after an Amazon ECS deployment. Validation can occur while test traffic is served during a deployment.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "listener_arns": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of one listener. The listener identifies the route between a target group and a load balancer. This is an array of strings with a maximum size of one.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "on_premises_instance_tag_filters": {
        "computed": true,
        "description": "The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group. CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed. You can specify OnPremisesInstanceTagFilters or OnPremisesInstanceTagSet, but not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The on-premises instance tag filter key.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The on-premises instance tag filter type",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The on-premises instance tag filter value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "on_premises_tag_set": {
        "computed": true,
        "description": "Information about groups of tags applied to on-premises instances. The deployment group includes only on-premises instances identified by all the tag groups. You can specify OnPremisesInstanceTagFilters or OnPremisesInstanceTagSet, but not both.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "on_premises_tag_set_list": {
              "computed": true,
              "description": "A list that contains other lists of on-premises instance tag groups. For an instance to be included in the deployment group, it must be identified by all of the tag groups in the list.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_premises_tag_group": {
                    "computed": true,
                    "description": "Information about groups of on-premises instance tags.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The on-premises instance tag filter key.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The on-premises instance tag filter type",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The on-premises instance tag filter value.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "outdated_instances_strategy": {
        "computed": true,
        "description": "Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision. If this option is set to UPDATE or is unspecified, CodeDeploy initiates one or more 'auto-update outdated instances' deployments to apply the deployed application revision to the new Amazon EC2 instances. If this option is set to IGNORE, CodeDeploy does not initiate a deployment to update the new Amazon EC2 instances. This may result in instances having different revisions.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_role_arn": {
        "computed": true,
        "description": "A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf. For more information, see 'Create a Service Role for AWS CodeDeploy' in the AWS CodeDeploy User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them. Each tag consists of a key and an optional value, both of which you define.",
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
      "termination_hook_enabled": {
        "computed": true,
        "description": "Indicates whether the deployment group was configured to have CodeDeploy install a termination hook into an Auto Scaling group.",
        "description_kind": "plain",
        "type": "bool"
      },
      "trigger_configurations": {
        "computed": true,
        "description": "Information about triggers associated with the deployment group. Duplicates are not allowed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "trigger_events": {
              "computed": true,
              "description": "The event type or types that trigger notifications.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "trigger_name": {
              "computed": true,
              "description": "The name of the notification trigger.",
              "description_kind": "plain",
              "type": "string"
            },
            "trigger_target_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic through which notifications about deployment or instance events are sent.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeDeploy::DeploymentGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodedeployDeploymentGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodedeployDeploymentGroup), &result)
	return &result
}
