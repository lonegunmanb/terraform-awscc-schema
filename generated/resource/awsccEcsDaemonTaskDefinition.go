package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsDaemonTaskDefinition = `{
  "block": {
    "attributes": {
      "container_definitions": {
        "computed": true,
        "description": "A list of container definitions in JSON format that describe the containers that make up the daemon task.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "command": {
              "computed": true,
              "description": "The command that's passed to the container.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "computed": true,
              "description": "The number of ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` units reserved for the container.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "depends_on": {
              "computed": true,
              "description": "The dependencies defined for container startup and shutdown. A container can contain multiple dependencies on other containers in a task definition.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description": "The dependency condition of the container. The following are the available conditions and their behavior:\n  +  ` + "`" + `` + "`" + `START` + "`" + `` + "`" + ` - This condition emulates the behavior of links and volumes today. It validates that a dependent container is started before permitting other containers to start.\n  +  ` + "`" + `` + "`" + `COMPLETE` + "`" + `` + "`" + ` - This condition validates that a dependent container runs to completion (exits) before permitting other containers to start. This can be useful for nonessential containers that run a script and then exit. This condition can't be set on an essential container.\n  +  ` + "`" + `` + "`" + `SUCCESS` + "`" + `` + "`" + ` - This condition is the same as ` + "`" + `` + "`" + `COMPLETE` + "`" + `` + "`" + `, but it also requires that the container exits with a ` + "`" + `` + "`" + `zero` + "`" + `` + "`" + ` status. This condition can't be set on an essential container.\n  +  ` + "`" + `` + "`" + `HEALTHY` + "`" + `` + "`" + ` - This condition validates that the dependent container passes its Docker health check before permitting other containers to start. This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_name": {
                    "computed": true,
                    "description": "The name of a container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "entry_point": {
              "computed": true,
              "description": "The entry point that's passed to the container.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "environment": {
              "computed": true,
              "description": "The environment variables to pass to a container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the key-value pair. For environment variables, this is the name of the environment variable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value of the key-value pair. For environment variables, this is the value of the environment variable.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "environment_files": {
              "computed": true,
              "description": "A list of files containing the environment variables to pass to a container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "The file type to use. Environment files are objects in Amazon S3. The only supported value is ` + "`" + `` + "`" + `s3` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "essential": {
              "computed": true,
              "description": "If the ` + "`" + `` + "`" + `essential` + "`" + `` + "`" + ` parameter of a container is marked as ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, and that container fails or stops for any reason, all other containers that are part of the task are stopped.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "firelens_configuration": {
              "computed": true,
              "description": "The FireLens configuration for the container. This is used to specify and configure a log router for container logs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "options": {
                    "computed": true,
                    "description": "The options to use when configuring the log router. This field is optional and can be used to specify a custom configuration file or to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event. If specified, the syntax to use is ` + "`" + `` + "`" + `\"options\":{\"enable-ecs-log-metadata\":\"true|false\",\"config-file-type:\"s3|file\",\"config-file-value\":\"arn:aws:s3:::mybucket/fluent.conf|filepath\"}` + "`" + `` + "`" + `. For more information, see [Creating a task definition that uses a FireLens configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html#firelens-taskdef) in the *Amazon Elastic Container Service Developer Guide*.\n  Tasks hosted on FARGATElong only support the ` + "`" + `` + "`" + `file` + "`" + `` + "`" + ` configuration file type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "type": {
                    "computed": true,
                    "description": "The log router to use. The valid values are ` + "`" + `` + "`" + `fluentd` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `fluentbit` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "health_check": {
              "computed": true,
              "description": "The container health check command and associated configuration parameters for the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "A string array representing the command that the container runs to determine if it is healthy. The string array must start with ` + "`" + `` + "`" + `CMD` + "`" + `` + "`" + ` to run the command arguments directly, or ` + "`" + `` + "`" + `CMD-SHELL` + "`" + `` + "`" + ` to run the command with the container's default shell. \n  When you use the AWS Management Console JSON panel, the CLIlong, or the APIs, enclose the list of commands in double quotes and brackets.\n  ` + "`" + `` + "`" + `[ \"CMD-SHELL\", \"curl -f http://localhost/ || exit 1\" ]` + "`" + `` + "`" + ` \n You don't include the double quotes and brackets when you use the AWS Management Console.\n  ` + "`" + `` + "`" + `CMD-SHELL, curl -f http://localhost/ || exit 1` + "`" + `` + "`" + ` \n An exit code of 0 indicates success, and non-zero exit code indicates failure. For more information, see ` + "`" + `` + "`" + `HealthCheck` + "`" + `` + "`" + ` in the docker container create command.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "interval": {
                    "computed": true,
                    "description": "The time period in seconds between each health check execution. You may specify between 5 and 300 seconds. The default value is 30 seconds. This value applies only when you specify a ` + "`" + `` + "`" + `command` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "retries": {
                    "computed": true,
                    "description": "The number of times to retry a failed health check before the container is considered unhealthy. You may specify between 1 and 10 retries. The default value is 3. This value applies only when you specify a ` + "`" + `` + "`" + `command` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_period": {
                    "computed": true,
                    "description": "The optional grace period to provide containers time to bootstrap before failed health checks count towards the maximum number of retries. You can specify between 0 and 300 seconds. By default, the ` + "`" + `` + "`" + `startPeriod` + "`" + `` + "`" + ` is off. This value applies only when you specify a ` + "`" + `` + "`" + `command` + "`" + `` + "`" + `. \n  If a health check succeeds within the ` + "`" + `` + "`" + `startPeriod` + "`" + `` + "`" + `, then the container is considered healthy and any subsequent failures count toward the maximum number of retries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout": {
                    "computed": true,
                    "description": "The time period in seconds to wait for a health check to succeed before it is considered a failure. You may specify between 2 and 60 seconds. The default value is 5. This value applies only when you specify a ` + "`" + `` + "`" + `command` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "image": {
              "computed": true,
              "description": "The image used to start the container. This string is passed directly to the Docker daemon. Images in the Docker Hub registry are available by default. Other repositories are specified with either ` + "`" + `` + "`" + `repository-url/image:tag` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `repository-url/image@digest` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "interactive": {
              "computed": true,
              "description": "When this parameter is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, you can deploy containerized applications that require ` + "`" + `` + "`" + `stdin` + "`" + `` + "`" + ` or a ` + "`" + `` + "`" + `tty` + "`" + `` + "`" + ` to be allocated.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "linux_parameters": {
              "computed": true,
              "description": "Linux-specific modifications that are applied to the container configuration, such as Linux kernel capabilities.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capabilities": {
                    "computed": true,
                    "description": "The Linux capabilities for the container that are added to or dropped from the default configuration provided by Docker.\n  For tasks that use the Fargate launch type, ` + "`" + `` + "`" + `capabilities` + "`" + `` + "`" + ` is supported for all platform versions but the ` + "`" + `` + "`" + `add` + "`" + `` + "`" + ` parameter is only supported if using platform version 1.4.0 or later.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "add": {
                          "computed": true,
                          "description": "The Linux capabilities for the container that have been added to the default configuration provided by Docker. This parameter maps to ` + "`" + `` + "`" + `CapAdd` + "`" + `` + "`" + ` in the docker container create command and the ` + "`" + `` + "`" + `--cap-add` + "`" + `` + "`" + ` option to docker run.\n  Tasks launched on FARGATElong only support adding the ` + "`" + `` + "`" + `SYS_PTRACE` + "`" + `` + "`" + ` kernel capability.\n  Valid values: ` + "`" + `` + "`" + `\"ALL\" | \"AUDIT_CONTROL\" | \"AUDIT_WRITE\" | \"BLOCK_SUSPEND\" | \"CHOWN\" | \"DAC_OVERRIDE\" | \"DAC_READ_SEARCH\" | \"FOWNER\" | \"FSETID\" | \"IPC_LOCK\" | \"IPC_OWNER\" | \"KILL\" | \"LEASE\" | \"LINUX_IMMUTABLE\" | \"MAC_ADMIN\" | \"MAC_OVERRIDE\" | \"MKNOD\" | \"NET_ADMIN\" | \"NET_BIND_SERVICE\" | \"NET_BROADCAST\" | \"NET_RAW\" | \"SETFCAP\" | \"SETGID\" | \"SETPCAP\" | \"SETUID\" | \"SYS_ADMIN\" | \"SYS_BOOT\" | \"SYS_CHROOT\" | \"SYS_MODULE\" | \"SYS_NICE\" | \"SYS_PACCT\" | \"SYS_PTRACE\" | \"SYS_RAWIO\" | \"SYS_RESOURCE\" | \"SYS_TIME\" | \"SYS_TTY_CONFIG\" | \"SYSLOG\" | \"WAKE_ALARM\"` + "`" + `` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "drop": {
                          "computed": true,
                          "description": "The Linux capabilities for the container that have been removed from the default configuration provided by Docker. This parameter maps to ` + "`" + `` + "`" + `CapDrop` + "`" + `` + "`" + ` in the docker container create command and the ` + "`" + `` + "`" + `--cap-drop` + "`" + `` + "`" + ` option to docker run.\n Valid values: ` + "`" + `` + "`" + `\"ALL\" | \"AUDIT_CONTROL\" | \"AUDIT_WRITE\" | \"BLOCK_SUSPEND\" | \"CHOWN\" | \"DAC_OVERRIDE\" | \"DAC_READ_SEARCH\" | \"FOWNER\" | \"FSETID\" | \"IPC_LOCK\" | \"IPC_OWNER\" | \"KILL\" | \"LEASE\" | \"LINUX_IMMUTABLE\" | \"MAC_ADMIN\" | \"MAC_OVERRIDE\" | \"MKNOD\" | \"NET_ADMIN\" | \"NET_BIND_SERVICE\" | \"NET_BROADCAST\" | \"NET_RAW\" | \"SETFCAP\" | \"SETGID\" | \"SETPCAP\" | \"SETUID\" | \"SYS_ADMIN\" | \"SYS_BOOT\" | \"SYS_CHROOT\" | \"SYS_MODULE\" | \"SYS_NICE\" | \"SYS_PACCT\" | \"SYS_PTRACE\" | \"SYS_RAWIO\" | \"SYS_RESOURCE\" | \"SYS_TIME\" | \"SYS_TTY_CONFIG\" | \"SYSLOG\" | \"WAKE_ALARM\"` + "`" + `` + "`" + `",
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
                  },
                  "devices": {
                    "computed": true,
                    "description": "Any host devices to expose to the container. This parameter maps to ` + "`" + `` + "`" + `Devices` + "`" + `` + "`" + ` in the docker container create command and the ` + "`" + `` + "`" + `--device` + "`" + `` + "`" + ` option to docker run.\n  If you're using tasks that use the Fargate launch type, the ` + "`" + `` + "`" + `devices` + "`" + `` + "`" + ` parameter isn't supported.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_path": {
                          "computed": true,
                          "description": "The path inside the container at which to expose the host device.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "host_path": {
                          "computed": true,
                          "description": "The path for the device on the host container instance.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "permissions": {
                          "computed": true,
                          "description": "The explicit permissions to provide to the container for the device. By default, the container has permissions for ` + "`" + `` + "`" + `read` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `write` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `mknod` + "`" + `` + "`" + ` for the device.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "init_process_enabled": {
                    "computed": true,
                    "description": "Run an ` + "`" + `` + "`" + `init` + "`" + `` + "`" + ` process inside the container that forwards signals and reaps processes. This parameter maps to the ` + "`" + `` + "`" + `--init` + "`" + `` + "`" + ` option to docker run. This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "tmpfs": {
                    "computed": true,
                    "description": "The container path, mount options, and size (in MiB) of the tmpfs mount. This parameter maps to the ` + "`" + `` + "`" + `--tmpfs` + "`" + `` + "`" + ` option to docker run.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_path": {
                          "computed": true,
                          "description": "The absolute file path where the tmpfs volume is to be mounted.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "mount_options": {
                          "computed": true,
                          "description": "The list of tmpfs volume mount options.\n Valid values: ` + "`" + `` + "`" + `\"defaults\" | \"ro\" | \"rw\" | \"suid\" | \"nosuid\" | \"dev\" | \"nodev\" | \"exec\" | \"noexec\" | \"sync\" | \"async\" | \"dirsync\" | \"remount\" | \"mand\" | \"nomand\" | \"atime\" | \"noatime\" | \"diratime\" | \"nodiratime\" | \"bind\" | \"rbind\" | \"unbindable\" | \"runbindable\" | \"private\" | \"rprivate\" | \"shared\" | \"rshared\" | \"slave\" | \"rslave\" | \"relatime\" | \"norelatime\" | \"strictatime\" | \"nostrictatime\" | \"mode\" | \"uid\" | \"gid\" | \"nr_inodes\" | \"nr_blocks\" | \"mpol\"` + "`" + `` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "size": {
                          "computed": true,
                          "description": "The maximum size (in MiB) of the tmpfs volume.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "log_configuration": {
              "computed": true,
              "description": "The log configuration specification for the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_driver": {
                    "computed": true,
                    "description": "The log driver to use for the container.\n For tasks on FARGATElong, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For tasks hosted on Amazon EC2 instances, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `fluentd` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `gelf` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `json-file` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `journald` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `syslog` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For more information about using the ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` log driver, see [Send Amazon ECS logs to CloudWatch](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html) in the *Amazon Elastic Container Service Developer Guide*.\n For more information about using the ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` log driver, see [Send Amazon ECS logs to an service or Partner](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html).\n  If you have a custom driver that isn't listed, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you would like to have included. However, we don't currently provide support for running modified copies of this software.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The configuration options to send to the log driver.\n The options you can specify depend on the log driver. Some of the options you can specify when you use the ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` log driver to route logs to Amazon CloudWatch include the following:\n  + awslogs-create-group Required: No Specify whether you want the log group to be created automatically. If this option isn't specified, it defaults to false. Your IAM policy must include the logs:CreateLogGroup permission before you attempt to use awslogs-create-group. + awslogs-region Required: Yes Specify the Region that the awslogs log driver is to send your Docker logs to. You can choose to send all of your logs from clusters in different Regions to a single region in CloudWatch Logs. This is so that they're all visible in one location. Otherwise, you can separate them by Region for more granularity. Make sure that the specified log group exists in the Region that you specify with this option. + awslogs-group Required: Yes Make sure to specify a log group that the awslogs log driver sends its log streams to. + awslogs-stream-prefix Required: Yes, when using Fargate.Optional when using EC2. Use the awslogs-stream-prefix option to associate a log stream with the specified prefix, the container name, and the ID of the Amazon ECS task that the container belongs to. If you specify a prefix with this option, then the log stream takes the format prefix-name/container-name/ecs-task-id. If you don't specify a prefix with this option, then the log stream is named after the container ID that's assigned by the Docker daemon on the container instance. Because it's difficult to trace logs back to the container that sent them with just the Docker container ID (which is only available on the container instance), we recommend that you specify a prefix with this option. For Amazon ECS services, you can use the service name as the prefix. Doing so, you can trace log streams to the service that the container belongs to, the name of the container that sent them, and the ID of the task that the container belongs to. You must specify a stream-prefix for your logs to have your logs appear in the Log pane when using the Amazon ECS console. + awslogs-datetime-format Required: No This option defines a multiline start pattern in Python strftime format. A log message consists of a line that matches the pattern and any following lines that don’t match the pattern. The matched line is the delimiter between log messages. One example of a use case for using this format is for parsing output such as a stack dump, which might otherwise be logged in multiple entries. The correct pattern allows it to be captured in a single entry. For more information, see awslogs-datetime-format. You cannot configure both the awslogs-datetime-format and awslogs-multiline-pattern options. Multiline logging performs regular expression parsing and matching of all log messages. This might have a negative impact on logging performance. + awslogs-multiline-pattern Required: No This option defines a multiline start pattern that uses a regular expression. A log message consists of a line that matches the pattern and any following lines that don’t match the pattern. The matched line is the delimiter between log messages. For more information, see awslogs-multiline-pattern. This option is ignored if awslogs-datetime-format is also configured. You cannot configure both the awslogs-datetime-format and awslogs-multiline-pattern options. Multiline logging performs regular expression parsing and matching of all log messages. This might have a negative impact on logging performance. \n The following options apply to all supported log drivers.\n  + mode Required: No Valid values: non-blocking | blocking This option defines the delivery mode of log messages from the container to the log driver specified using logDriver. The delivery mode you choose affects application availability when the flow of logs from container is interrupted. If you use the blocking mode and the flow of logs is interrupted, calls from container code to write to the stdout and stderr streams will block. The logging thread of the application will block as a result. This may cause the application to become unresponsive and lead to container healthcheck failure. If you use the non-blocking mode, the container's logs are instead stored in an in-memory intermediate buffer configured with the max-buffer-size option. This prevents the application from becoming unresponsive when logs cannot be sent. We recommend using this mode if you want to ensure service availability and are okay with some log loss. For more information, see Preventing log loss with non-blocking mode in the awslogs container log driver. You can set a default mode for all containers in a specific Region by using the defaultLogDriverMode account setting. If you don't specify the mode option or configure the account setting, Amazon ECS will default to the non-blocking mode. For more information about the account setting, see Default log driver mode in the Amazon Elastic Container Service Developer Guide. On June 25, 2025, Amazon ECS changed the default log driver mode from blocking to non-blocking to prioritize task availability over logging. To continue using the blocking mode after this change, do one of the following: Set the mode option in your container definition's logConfiguration as blocking. Set the defaultLogDriverMode account setting to blocking. + max-buffer-size Required: No Default value: 10m When non-blocking mode is used, the max-buffer-size log option controls the size of the buffer that's used for intermediate message storage. Make sure to specify an adequate buffer size based on your application. When the buffer fills up, further logs cannot be stored. Logs that cannot be stored are lost. \n To route logs using the ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + ` log router, you need to specify a ` + "`" + `` + "`" + `splunk-token` + "`" + `` + "`" + ` and a ` + "`" + `` + "`" + `splunk-url` + "`" + `` + "`" + `.\n When you use the ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` log router to route logs to an AWS Service or AWS Partner Network destination for log storage and analytics, you can set the ` + "`" + `` + "`" + `log-driver-buffer-limit` + "`" + `` + "`" + ` option to limit the number of events that are buffered in memory, before being sent to the log router container. It can help to resolve potential log loss issue because high throughput might result in memory running out for the buffer inside of Docker.\n Other options you can specify when using ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` to route logs depend on the destination. When you export logs to Amazon Data Firehose, you can specify the AWS Region with ` + "`" + `` + "`" + `region` + "`" + `` + "`" + ` and a name for the log stream with ` + "`" + `` + "`" + `delivery_stream` + "`" + `` + "`" + `.\n When you export logs to Amazon Kinesis Data Streams, you can specify an AWS Region with ` + "`" + `` + "`" + `region` + "`" + `` + "`" + ` and a data stream name with ` + "`" + `` + "`" + `stream` + "`" + `` + "`" + `.\n  When you export logs to Amazon OpenSearch Service, you can specify options like ` + "`" + `` + "`" + `Name` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Host` + "`" + `` + "`" + ` (OpenSearch Service endpoint without protocol), ` + "`" + `` + "`" + `Port` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Index` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Aws_auth` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Aws_region` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Suppress_Type_Name` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `tls` + "`" + `` + "`" + `. For more information, see [Under the hood: FireLens for Amazon ECS Tasks](https://docs.aws.amazon.com/containers/under-the-hood-firelens-for-amazon-ecs-tasks/).\n When you export logs to Amazon S3, you can specify the bucket using the ` + "`" + `` + "`" + `bucket` + "`" + `` + "`" + ` option. You can also specify ` + "`" + `` + "`" + `region` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `total_file_size` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `upload_timeout` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `use_put_object` + "`" + `` + "`" + ` as options.\n This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "secret_options": {
                    "computed": true,
                    "description": "The secrets to pass to the log configuration. For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the secret.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value_from": {
                          "computed": true,
                          "description": "The secret to expose to the container. The supported values are either the full ARN of the ASMlong secret or the full ARN of the parameter in the SSM Parameter Store.\n For information about the require IAMlong permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide*.\n  If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.",
                          "description_kind": "plain",
                          "optional": true,
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
            "memory": {
              "computed": true,
              "description": "The amount (in MiB) of memory to present to the container. If the container attempts to exceed the memory specified here, the container is killed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory_reservation": {
              "computed": true,
              "description": "The soft limit (in MiB) of memory to reserve for the container.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mount_points": {
              "computed": true,
              "description": "The mount points for data volumes in your container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_path": {
                    "computed": true,
                    "description": "The path on the container to mount the host volume at.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "read_only": {
                    "computed": true,
                    "description": "If this value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the container has read-only access to the volume. If this value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, then the container can write to the volume. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_volume": {
                    "computed": true,
                    "description": "The name of the volume to mount. Must be a volume name referenced in the ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` parameter of task definition ` + "`" + `` + "`" + `volume` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "name": {
              "computed": true,
              "description": "The name of the container. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "privileged": {
              "computed": true,
              "description": "When this parameter is true, the container is given elevated privileges on the host container instance (similar to the ` + "`" + `` + "`" + `root` + "`" + `` + "`" + ` user).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pseudo_terminal": {
              "computed": true,
              "description": "When this parameter is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, a TTY is allocated.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "readonly_root_filesystem": {
              "computed": true,
              "description": "When this parameter is true, the container is given read-only access to its root file system.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "repository_credentials": {
              "computed": true,
              "description": "The private repository authentication credentials to use.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "credentials_parameter": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the secret containing the private repository credentials.\n  When you use the Amazon ECS API, CLI, or AWS SDK, if the secret exists in the same Region as the task that you're launching then you can use either the full ARN or the name of the secret. When you use the AWS Management Console, you must specify the full ARN of the secret.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "restart_policy": {
              "computed": true,
              "description": "The restart policy for the container. When you set up a restart policy, Amazon ECS can restart the container without needing to replace the task.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ignored_exit_codes": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "restart_attempt_period": {
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
            "secrets": {
              "computed": true,
              "description": "The secrets to pass to the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the secret.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value_from": {
                    "computed": true,
                    "description": "The secret to expose to the container. The supported values are either the full ARN of the ASMlong secret or the full ARN of the parameter in the SSM Parameter Store.\n For information about the require IAMlong permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide*.\n  If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "start_timeout": {
              "computed": true,
              "description": "Time duration (in seconds) to wait before giving up on resolving dependencies for a container.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stop_timeout": {
              "computed": true,
              "description": "Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "system_controls": {
              "computed": true,
              "description": "A list of namespaced kernel parameters to set in the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "namespace": {
                    "computed": true,
                    "description": "The namespaced kernel parameter to set a ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` for.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The namespaced kernel parameter to set a ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` for.\n Valid IPC namespace values: ` + "`" + `` + "`" + `\"kernel.msgmax\" | \"kernel.msgmnb\" | \"kernel.msgmni\" | \"kernel.sem\" | \"kernel.shmall\" | \"kernel.shmmax\" | \"kernel.shmmni\" | \"kernel.shm_rmid_forced\"` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Sysctls` + "`" + `` + "`" + ` that start with ` + "`" + `` + "`" + `\"fs.mqueue.*\"` + "`" + `` + "`" + `\n Valid network namespace values: ` + "`" + `` + "`" + `Sysctls` + "`" + `` + "`" + ` that start with ` + "`" + `` + "`" + `\"net.*\"` + "`" + `` + "`" + `. Only namespaced ` + "`" + `` + "`" + `Sysctls` + "`" + `` + "`" + ` that exist within the container starting with \"net.* are accepted.\n All of these values are supported by Fargate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "ulimits": {
              "computed": true,
              "description": "A list of ` + "`" + `` + "`" + `ulimits` + "`" + `` + "`" + ` to set in the container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hard_limit": {
                    "computed": true,
                    "description": "The hard limit for the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + ` type. The value can be specified in bytes, seconds, or as a count, depending on the ` + "`" + `` + "`" + `type` + "`" + `` + "`" + ` of the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description": "The ` + "`" + `` + "`" + `type` + "`" + `` + "`" + ` of the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "soft_limit": {
                    "computed": true,
                    "description": "The soft limit for the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + ` type. The value can be specified in bytes, seconds, or as a count, depending on the ` + "`" + `` + "`" + `type` + "`" + `` + "`" + ` of the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "user": {
              "computed": true,
              "description": "The user to use inside the container.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "working_directory": {
              "computed": true,
              "description": "The working directory to run commands inside the container in.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "cpu": {
        "computed": true,
        "description": "The number of CPU units used by the daemon task.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "daemon_task_definition_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make Amazon Web Services API calls on your behalf.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "family": {
        "computed": true,
        "description": "The name of a family that this daemon task definition is registered to.",
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
      "ipc_mode": {
        "computed": true,
        "description": "The IPC namespace mode for the daemon. The valid values are ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `shared` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `.\n If ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` is specified or no value is provided, the daemon runs with its own IPC namespace, isolated from other tasks. If ` + "`" + `` + "`" + `shared` + "`" + `` + "`" + ` is specified, the daemon joins the host IPC namespace, making it accessible to non-daemon tasks that use ` + "`" + `` + "`" + `ipcMode: \"host\"` + "`" + `` + "`" + ` or other daemons that use ` + "`" + `` + "`" + `ipcMode: \"shared\"` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory": {
        "computed": true,
        "description": "The amount of memory (in MiB) used by the daemon task.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pid_mode": {
        "computed": true,
        "description": "The PID namespace mode for the daemon. The valid values are ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `shared` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `.\n If ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` is specified or no value is provided, the daemon runs with its own PID namespace, isolated from other tasks. If ` + "`" + `` + "`" + `shared` + "`" + `` + "`" + ` is specified, the daemon joins the host PID namespace, making it accessible to non-daemon tasks that use ` + "`" + `` + "`" + `pidMode: \"host\"` + "`" + `` + "`" + ` or other daemons that use ` + "`" + `` + "`" + `pidMode: \"shared\"` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "task_role_arn": {
        "computed": true,
        "description": "The short name or full Amazon Resource Name (ARN) of the IAM role that grants containers in the daemon task permission to call Amazon Web Services APIs on your behalf.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volumes": {
        "computed": true,
        "description": "The list of data volume definitions for the daemon task.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "host": {
              "computed": true,
              "description": "This parameter is specified when you use bind mount host volumes. The contents of the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` parameter determine whether your bind mount host volume persists on the host container instance and where it's stored. If the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.\n Windows containers can mount whole directories on the same drive as ` + "`" + `` + "`" + `$env:ProgramData` + "`" + `` + "`" + `. Windows containers can't mount directories on a different drive, and mount point can't be across drives. For example, you can mount ` + "`" + `` + "`" + `C:\\my\\path:C:\\my\\path` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `D:\\:D:\\` + "`" + `` + "`" + `, but not ` + "`" + `` + "`" + `D:\\my\\path:C:\\my\\path` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `D:\\:C:\\my\\path` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_path": {
                    "computed": true,
                    "description": "When the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` parameter is used, specify a ` + "`" + `` + "`" + `sourcePath` + "`" + `` + "`" + ` to declare the path on the host container instance that's presented to the container. If this parameter is empty, then the Docker daemon has assigned a host path for you. If the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` parameter contains a ` + "`" + `` + "`" + `sourcePath` + "`" + `` + "`" + ` file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the ` + "`" + `` + "`" + `sourcePath` + "`" + `` + "`" + ` value doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.\n If you're using the Fargate launch type, the ` + "`" + `` + "`" + `sourcePath` + "`" + `` + "`" + ` parameter is not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "name": {
              "computed": true,
              "description": "The name of the volume. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.\n When using a volume configured at launch, the ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` is required and must also be specified as the volume name in the ` + "`" + `` + "`" + `ServiceVolumeConfiguration` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `TaskVolumeConfiguration` + "`" + `` + "`" + ` parameter when creating your service or standalone task.\n For all other types of volumes, this name is referenced in the ` + "`" + `` + "`" + `sourceVolume` + "`" + `` + "`" + ` parameter of the ` + "`" + `` + "`" + `mountPoints` + "`" + `` + "`" + ` object in the container definition.\n When a volume is using the ` + "`" + `` + "`" + `efsVolumeConfiguration` + "`" + `` + "`" + `, the name is required.\n When a volume is using the ` + "`" + `` + "`" + `s3filesVolumeConfiguration` + "`" + `` + "`" + `, the name is required.",
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
    "description": "The details of a daemon task definition. A daemon task definition is a template that describes the containers that form a daemon. Daemons deploy cross-cutting software agents independently across your Amazon ECS infrastructure.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsDaemonTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsDaemonTaskDefinition), &result)
	return &result
}
