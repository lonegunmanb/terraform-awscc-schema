package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsTaskDefinition = `{
  "block": {
    "attributes": {
      "container_definitions": {
        "computed": true,
        "description": "A list of container definitions in JSON format that describe the different containers that make up your task. For more information about container definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "command": {
              "computed": true,
              "description": "The command that's passed to the container. This parameter maps to ` + "`" + `` + "`" + `Cmd` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `COMMAND` + "`" + `` + "`" + ` parameter to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). For more information, see [https://docs.docker.com/engine/reference/builder/#cmd](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#cmd). If there are multiple arguments, each argument is a separated string in the array.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "computed": true,
              "description": "The number of ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` units reserved for the container. This parameter maps to ` + "`" + `` + "`" + `CpuShares` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--cpu-shares` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n This field is optional for tasks using the Fargate launch type, and the only requirement is that the total amount of CPU reserved for all containers within a task be lower than the task-level ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` value.\n  You can determine the number of CPU units that are available per EC2 instance type by multiplying the vCPUs listed for that instance type on the [Amazon EC2 Instances](https://docs.aws.amazon.com/ec2/instance-types/) detail page by 1,024.\n  Linux containers share unallocated CPU units with other containers on the container instance with the same ratio as their allocated amount. For example, if you run a single-container task on a single-core instance type with 512 CPU units specified for that container, and that's the only task running on the container instance, that container could use the full 1,024 CPU unit share at any given time. However, if you launched another copy of the same task on that container instance, each task is guaranteed a minimum of 512 CPU units when needed. Moreover, each container could float to higher CPU usage if the other container was not using it. If both tasks were 100% active all of the time, they would be limited to 512 CPU units.\n On Linux container instances, the Docker daemon on the container instance uses the CPU value to calculate the relative CPU share ratios for running containers. For more information, see [CPU share constraint](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#cpu-share-constraint) in the Docker documentation. The minimum valid CPU share value that the Linux kernel allows is 2. However, the CPU parameter isn't required, and you can use CPU values below 2 in your container definitions. For CPU values below 2 (including null), the behavior varies based on your Amazon ECS container agent version:\n  +   *Agent versions less than or equal to 1.1.0:* Null and zero CPU values are passed to Docker as 0, which Docker then converts to 1,024 CPU shares. CPU values of 1 are passed to Docker as 1, which the Linux kernel converts to two CPU shares.\n  +   *Agent versions greater than or equal to 1.2.0:* Null, zero, and CPU values of 1 are passed to Docker as 2.\n  \n On Windows container instances, the CPU limit is enforced as an absolute limit, or a quota. Windows containers only have access to the specified amount of CPU that's described in the task definition. A null or zero CPU value is passed to Docker as ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `, which Windows interprets as 1% of one CPU.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "credential_specs": {
              "computed": true,
              "description": "A list of ARNs in SSM or Amazon S3 to a credential spec (` + "`" + `` + "`" + `CredSpec` + "`" + `` + "`" + `) file that configures the container for Active Directory authentication. We recommend that you use this parameter instead of the ` + "`" + `` + "`" + `dockerSecurityOptions` + "`" + `` + "`" + `. The maximum number of ARNs is 1.\n There are two formats for each ARN.\n  + credentialspecdomainless:MyARN You use credentialspecdomainless:MyARN to provide a CredSpec with an additional section for a secret in . You provide the login credentials to the domain in the secret. Each task that runs on any container instance can join different domains. You can use this format without joining the container instance to a domain. + credentialspec:MyARN You use credentialspec:MyARN to provide a CredSpec for a single domain. You must join the container instance to the domain before you start any tasks that use this task definition. \n In both formats, replace ` + "`" + `` + "`" + `MyARN` + "`" + `` + "`" + ` with the ARN in SSM or Amazon S3.\n If you provide a ` + "`" + `` + "`" + `credentialspecdomainless:MyARN` + "`" + `` + "`" + `, the ` + "`" + `` + "`" + `credspec` + "`" + `` + "`" + ` must provide a ARN in ASMlong for a secret containing the username, password, and the domain to connect to. For better security, the instance isn't joined to the domain for domainless authentication. Other applications on the instance can't use the domainless credentials. You can use this parameter to run tasks on the same instance, even it the tasks need to join different domains. For more information, see [Using gMSAs for Windows Containers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows-gmsa.html) and [Using gMSAs for Linux Containers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/linux-gmsa.html).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "depends_on": {
              "computed": true,
              "description": "The dependencies defined for container startup and shutdown. A container can contain multiple dependencies. When a dependency is defined for container startup, for container shutdown it is reversed.\n For tasks using the EC2 launch type, the container instances require at least version 1.26.0 of the container agent to turn on container dependencies. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide*. If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + ` package. If your container instances are launched from version ` + "`" + `` + "`" + `20190301` + "`" + `` + "`" + ` or later, then they contain the required versions of the container agent and ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + `. For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide*.\n For tasks using the Fargate launch type, the task or service requires the following platforms:\n  +  Linux platform version ` + "`" + `` + "`" + `1.3.0` + "`" + `` + "`" + ` or later.\n  +  Windows platform version ` + "`" + `` + "`" + `1.0.0` + "`" + `` + "`" + ` or later.\n  \n If the task definition is used in a blue/green deployment that uses [AWS::CodeDeploy::DeploymentGroup BlueGreenDeploymentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-bluegreendeploymentconfiguration.html), the ` + "`" + `` + "`" + `dependsOn` + "`" + `` + "`" + ` parameter is not supported. For more information see [Issue #680](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-coverage-roadmap/issues/680) on the on the GitHub website.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "condition": {
                    "computed": true,
                    "description": "The dependency condition of the container. The following are the available conditions and their behavior:\n  +   ` + "`" + `` + "`" + `START` + "`" + `` + "`" + ` - This condition emulates the behavior of links and volumes today. It validates that a dependent container is started before permitting other containers to start.\n  +   ` + "`" + `` + "`" + `COMPLETE` + "`" + `` + "`" + ` - This condition validates that a dependent container runs to completion (exits) before permitting other containers to start. This can be useful for nonessential containers that run a script and then exit. This condition can't be set on an essential container.\n  +   ` + "`" + `` + "`" + `SUCCESS` + "`" + `` + "`" + ` - This condition is the same as ` + "`" + `` + "`" + `COMPLETE` + "`" + `` + "`" + `, but it also requires that the container exits with a ` + "`" + `` + "`" + `zero` + "`" + `` + "`" + ` status. This condition can't be set on an essential container.\n  +   ` + "`" + `` + "`" + `HEALTHY` + "`" + `` + "`" + ` - This condition validates that the dependent container passes its Docker health check before permitting other containers to start. This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.",
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
            "disable_networking": {
              "computed": true,
              "description": "When this parameter is true, networking is off within the container. This parameter maps to ` + "`" + `` + "`" + `NetworkDisabled` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/).\n  This parameter is not supported for Windows containers.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dns_search_domains": {
              "computed": true,
              "description": "A list of DNS search domains that are presented to the container. This parameter maps to ` + "`" + `` + "`" + `DnsSearch` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--dns-search` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  This parameter is not supported for Windows containers.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dns_servers": {
              "computed": true,
              "description": "A list of DNS servers that are presented to the container. This parameter maps to ` + "`" + `` + "`" + `Dns` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--dns` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  This parameter is not supported for Windows containers.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "docker_labels": {
              "computed": true,
              "description": "A key/value map of labels to add to the container. This parameter maps to ` + "`" + `` + "`" + `Labels` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--label` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "docker_security_options": {
              "computed": true,
              "description": "A list of strings to provide custom configuration for multiple security systems. For more information about valid values, see [Docker Run Security Configuration](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). This field isn't valid for containers in tasks using the Fargate launch type.\n For Linux tasks on EC2, this parameter can be used to reference custom labels for SELinux and AppArmor multi-level security systems.\n For any tasks on EC2, this parameter can be used to reference a credential spec file that configures a container for Active Directory authentication. For more information, see [Using gMSAs for Windows Containers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows-gmsa.html) and [Using gMSAs for Linux Containers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/linux-gmsa.html) in the *Amazon Elastic Container Service Developer Guide*.\n This parameter maps to ` + "`" + `` + "`" + `SecurityOpt` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--security-opt` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  The Amazon ECS container agent running on a container instance must register with the ` + "`" + `` + "`" + `ECS_SELINUX_CAPABLE=true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ECS_APPARMOR_CAPABLE=true` + "`" + `` + "`" + ` environment variables before containers placed on that instance can use these security options. For more information, see [Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide*.\n  For more information about valid values, see [Docker Run Security Configuration](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). \n Valid values: \"no-new-privileges\" | \"apparmor:PROFILE\" | \"label:value\" | \"credentialspec:CredentialSpecFilePath\"",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "entry_point": {
              "computed": true,
              "description": "Early versions of the Amazon ECS container agent don't properly handle ` + "`" + `` + "`" + `entryPoint` + "`" + `` + "`" + ` parameters. If you have problems using ` + "`" + `` + "`" + `entryPoint` + "`" + `` + "`" + `, update your container agent or enter your commands and arguments as ` + "`" + `` + "`" + `command` + "`" + `` + "`" + ` array items instead.\n  The entry point that's passed to the container. This parameter maps to ` + "`" + `` + "`" + `Entrypoint` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--entrypoint` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). For more information, see [https://docs.docker.com/engine/reference/builder/#entrypoint](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/builder/#entrypoint).",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "environment": {
              "computed": true,
              "description": "The environment variables to pass to a container. This parameter maps to ` + "`" + `` + "`" + `Env` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--env` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  We don't recommend that you use plaintext environment variables for sensitive information, such as credential data.",
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
                "nesting_mode": "set"
              },
              "optional": true
            },
            "environment_files": {
              "computed": true,
              "description": "A list of files containing the environment variables to pass to a container. This parameter maps to the ` + "`" + `` + "`" + `--env-file` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n You can specify up to ten environment files. The file must have a ` + "`" + `` + "`" + `.env` + "`" + `` + "`" + ` file extension. Each line in an environment file contains an environment variable in ` + "`" + `` + "`" + `VARIABLE=VALUE` + "`" + `` + "`" + ` format. Lines beginning with ` + "`" + `` + "`" + `#` + "`" + `` + "`" + ` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/).\n If there are environment variables specified using the ` + "`" + `` + "`" + `environment` + "`" + `` + "`" + ` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying Environment Variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide*.",
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
              "description": "If the ` + "`" + `` + "`" + `essential` + "`" + `` + "`" + ` parameter of a container is marked as ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, and that container fails or stops for any reason, all other containers that are part of the task are stopped. If the ` + "`" + `` + "`" + `essential` + "`" + `` + "`" + ` parameter of a container is marked as ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, its failure doesn't affect the rest of the containers in a task. If this parameter is omitted, a container is assumed to be essential.\n All tasks must have at least one essential container. If you have an application that's composed of multiple containers, group containers that are used for a common purpose into components, and separate the different components into multiple task definitions. For more information, see [Application Architecture](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/application_architecture.html) in the *Amazon Elastic Container Service Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "extra_hosts": {
              "computed": true,
              "description": "A list of hostnames and IP address mappings to append to the ` + "`" + `` + "`" + `/etc/hosts` + "`" + `` + "`" + ` file on the container. This parameter maps to ` + "`" + `` + "`" + `ExtraHosts` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--add-host` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  This parameter isn't supported for Windows containers or tasks that use the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hostname": {
                    "computed": true,
                    "description": "The hostname to use in the ` + "`" + `` + "`" + `/etc/hosts` + "`" + `` + "`" + ` entry.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ip_address": {
                    "computed": true,
                    "description": "The IP address to use in the ` + "`" + `` + "`" + `/etc/hosts` + "`" + `` + "`" + ` entry.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "firelens_configuration": {
              "computed": true,
              "description": "The FireLens configuration for the container. This is used to specify and configure a log router for container logs. For more information, see [Custom Log Routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "options": {
                    "computed": true,
                    "description": "The options to use when configuring the log router. This field is optional and can be used to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event.\n  If specified, valid option keys are:\n  +   ` + "`" + `` + "`" + `enable-ecs-log-metadata` + "`" + `` + "`" + `, which can be ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `config-file-type` + "`" + `` + "`" + `, which can be ` + "`" + `` + "`" + `s3` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `file` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `config-file-value` + "`" + `` + "`" + `, which is either an S3 ARN or a file path",
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
              "description": "The container health check command and associated configuration parameters for the container. This parameter maps to ` + "`" + `` + "`" + `HealthCheck` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `HEALTHCHECK` + "`" + `` + "`" + ` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "A string array representing the command that the container runs to determine if it is healthy. The string array must start with ` + "`" + `` + "`" + `CMD` + "`" + `` + "`" + ` to run the command arguments directly, or ` + "`" + `` + "`" + `CMD-SHELL` + "`" + `` + "`" + ` to run the command with the container's default shell. \n  When you use the AWS Management Console JSON panel, the CLIlong, or the APIs, enclose the list of commands in double quotes and brackets.\n  ` + "`" + `` + "`" + `[ \"CMD-SHELL\", \"curl -f http://localhost/ || exit 1\" ]` + "`" + `` + "`" + ` \n You don't include the double quotes and brackets when you use the AWS Management Console.\n  ` + "`" + `` + "`" + `CMD-SHELL, curl -f http://localhost/ || exit 1` + "`" + `` + "`" + ` \n An exit code of 0 indicates success, and non-zero exit code indicates failure. For more information, see ` + "`" + `` + "`" + `HealthCheck` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "interval": {
                    "computed": true,
                    "description": "The time period in seconds between each health check execution. You may specify between 5 and 300 seconds. The default value is 30 seconds.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "retries": {
                    "computed": true,
                    "description": "The number of times to retry a failed health check before the container is considered unhealthy. You may specify between 1 and 10 retries. The default value is 3.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_period": {
                    "computed": true,
                    "description": "The optional grace period to provide containers time to bootstrap before failed health checks count towards the maximum number of retries. You can specify between 0 and 300 seconds. By default, the ` + "`" + `` + "`" + `startPeriod` + "`" + `` + "`" + ` is off.\n  If a health check succeeds within the ` + "`" + `` + "`" + `startPeriod` + "`" + `` + "`" + `, then the container is considered healthy and any subsequent failures count toward the maximum number of retries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout": {
                    "computed": true,
                    "description": "The time period in seconds to wait for a health check to succeed before it is considered a failure. You may specify between 2 and 60 seconds. The default value is 5.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "hostname": {
              "computed": true,
              "description": "The hostname to use for your container. This parameter maps to ` + "`" + `` + "`" + `Hostname` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--hostname` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  The ` + "`" + `` + "`" + `hostname` + "`" + `` + "`" + ` parameter is not supported if you're using the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image": {
              "description": "The image used to start a container. This string is passed directly to the Docker daemon. By default, images in the Docker Hub registry are available. Other repositories are specified with either ` + "`" + `` + "`" + `repository-url/image:tag` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `repository-url/image@digest` + "`" + `` + "`" + `. Up to 255 letters (uppercase and lowercase), numbers, hyphens, underscores, colons, periods, forward slashes, and number signs are allowed. This parameter maps to ` + "`" + `` + "`" + `Image` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `IMAGE` + "`" + `` + "`" + ` parameter of [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  +  When a new task starts, the Amazon ECS container agent pulls the latest version of the specified image and tag for the container to use. However, subsequent updates to a repository image aren't propagated to already running tasks.\n  +  Images in Amazon ECR repositories can be specified by either using the full ` + "`" + `` + "`" + `registry/repository:tag` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `registry/repository@digest` + "`" + `` + "`" + `. For example, ` + "`" + `` + "`" + `012345678910.dkr.ecr.\u003cregion-name\u003e.amazonaws.com/\u003crepository-name\u003e:latest` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `012345678910.dkr.ecr.\u003cregion-name\u003e.amazonaws.com/\u003crepository-name\u003e@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE` + "`" + `` + "`" + `. \n  +  Images in official repositories on Docker Hub use a single name (for example, ` + "`" + `` + "`" + `ubuntu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `mongo` + "`" + `` + "`" + `).\n  +  Images in other repositories on Docker Hub are qualified with an organization name (for example, ` + "`" + `` + "`" + `amazon/amazon-ecs-agent` + "`" + `` + "`" + `).\n  +  Images in other online repositories are qualified further by a domain name (for example, ` + "`" + `` + "`" + `quay.io/assemblyline/ubuntu` + "`" + `` + "`" + `).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "interactive": {
              "computed": true,
              "description": "When this parameter is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, you can deploy containerized applications that require ` + "`" + `` + "`" + `stdin` + "`" + `` + "`" + ` or a ` + "`" + `` + "`" + `tty` + "`" + `` + "`" + ` to be allocated. This parameter maps to ` + "`" + `` + "`" + `OpenStdin` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--interactive` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "links": {
              "computed": true,
              "description": "The ` + "`" + `` + "`" + `links` + "`" + `` + "`" + ` parameter allows containers to communicate with each other without the need for port mappings. This parameter is only supported if the network mode of a task definition is ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + `. The ` + "`" + `` + "`" + `name:internalName` + "`" + `` + "`" + ` construct is analogous to ` + "`" + `` + "`" + `name:alias` + "`" + `` + "`" + ` in Docker links. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. For more information about linking Docker containers, go to [Legacy container links](https://docs.aws.amazon.com/https://docs.docker.com/network/links/) in the Docker documentation. This parameter maps to ` + "`" + `` + "`" + `Links` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--link` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  This parameter is not supported for Windows containers.\n   Containers that are collocated on a single container instance may be able to communicate with each other without requiring links or host port mappings. Network isolation is achieved on the container instance using security groups and VPC settings.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "linux_parameters": {
              "computed": true,
              "description": "Linux-specific modifications that are applied to the container, such as Linux kernel capabilities. For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).\n  This parameter is not supported for Windows containers.",
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
                          "description": "The Linux capabilities for the container that have been added to the default configuration provided by Docker. This parameter maps to ` + "`" + `` + "`" + `CapAdd` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--cap-add` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  Tasks launched on FARGATElong only support adding the ` + "`" + `` + "`" + `SYS_PTRACE` + "`" + `` + "`" + ` kernel capability.\n  Valid values: ` + "`" + `` + "`" + `\"ALL\" | \"AUDIT_CONTROL\" | \"AUDIT_WRITE\" | \"BLOCK_SUSPEND\" | \"CHOWN\" | \"DAC_OVERRIDE\" | \"DAC_READ_SEARCH\" | \"FOWNER\" | \"FSETID\" | \"IPC_LOCK\" | \"IPC_OWNER\" | \"KILL\" | \"LEASE\" | \"LINUX_IMMUTABLE\" | \"MAC_ADMIN\" | \"MAC_OVERRIDE\" | \"MKNOD\" | \"NET_ADMIN\" | \"NET_BIND_SERVICE\" | \"NET_BROADCAST\" | \"NET_RAW\" | \"SETFCAP\" | \"SETGID\" | \"SETPCAP\" | \"SETUID\" | \"SYS_ADMIN\" | \"SYS_BOOT\" | \"SYS_CHROOT\" | \"SYS_MODULE\" | \"SYS_NICE\" | \"SYS_PACCT\" | \"SYS_PTRACE\" | \"SYS_RAWIO\" | \"SYS_RESOURCE\" | \"SYS_TIME\" | \"SYS_TTY_CONFIG\" | \"SYSLOG\" | \"WAKE_ALARM\"` + "`" + `` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "drop": {
                          "computed": true,
                          "description": "The Linux capabilities for the container that have been removed from the default configuration provided by Docker. This parameter maps to ` + "`" + `` + "`" + `CapDrop` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--cap-drop` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n Valid values: ` + "`" + `` + "`" + `\"ALL\" | \"AUDIT_CONTROL\" | \"AUDIT_WRITE\" | \"BLOCK_SUSPEND\" | \"CHOWN\" | \"DAC_OVERRIDE\" | \"DAC_READ_SEARCH\" | \"FOWNER\" | \"FSETID\" | \"IPC_LOCK\" | \"IPC_OWNER\" | \"KILL\" | \"LEASE\" | \"LINUX_IMMUTABLE\" | \"MAC_ADMIN\" | \"MAC_OVERRIDE\" | \"MKNOD\" | \"NET_ADMIN\" | \"NET_BIND_SERVICE\" | \"NET_BROADCAST\" | \"NET_RAW\" | \"SETFCAP\" | \"SETGID\" | \"SETPCAP\" | \"SETUID\" | \"SYS_ADMIN\" | \"SYS_BOOT\" | \"SYS_CHROOT\" | \"SYS_MODULE\" | \"SYS_NICE\" | \"SYS_PACCT\" | \"SYS_PTRACE\" | \"SYS_RAWIO\" | \"SYS_RESOURCE\" | \"SYS_TIME\" | \"SYS_TTY_CONFIG\" | \"SYSLOG\" | \"WAKE_ALARM\"` + "`" + `` + "`" + `",
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
                    "description": "Any host devices to expose to the container. This parameter maps to ` + "`" + `` + "`" + `Devices` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--device` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  If you're using tasks that use the Fargate launch type, the ` + "`" + `` + "`" + `devices` + "`" + `` + "`" + ` parameter isn't supported.",
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
                            "set",
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
                    "description": "Run an ` + "`" + `` + "`" + `init` + "`" + `` + "`" + ` process inside the container that forwards signals and reaps processes. This parameter maps to the ` + "`" + `` + "`" + `--init` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). This parameter requires version 1.25 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "max_swap": {
                    "computed": true,
                    "description": "The total amount of swap memory (in MiB) a container can use. This parameter will be translated to the ` + "`" + `` + "`" + `--memory-swap` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration) where the value would be the sum of the container memory plus the ` + "`" + `` + "`" + `maxSwap` + "`" + `` + "`" + ` value.\n If a ` + "`" + `` + "`" + `maxSwap` + "`" + `` + "`" + ` value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` is specified, the container will not use swap. Accepted values are ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` or any positive integer. If the ` + "`" + `` + "`" + `maxSwap` + "`" + `` + "`" + ` parameter is omitted, the container will use the swap configuration for the container instance it is running on. A ` + "`" + `` + "`" + `maxSwap` + "`" + `` + "`" + ` value must be set for the ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` parameter to be used.\n  If you're using tasks that use the Fargate launch type, the ` + "`" + `` + "`" + `maxSwap` + "`" + `` + "`" + ` parameter isn't supported.\n If you're using tasks on Amazon Linux 2023 the ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` parameter isn't supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "shared_memory_size": {
                    "computed": true,
                    "description": "The value for the size (in MiB) of the ` + "`" + `` + "`" + `/dev/shm` + "`" + `` + "`" + ` volume. This parameter maps to the ` + "`" + `` + "`" + `--shm-size` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  If you are using tasks that use the Fargate launch type, the ` + "`" + `` + "`" + `sharedMemorySize` + "`" + `` + "`" + ` parameter is not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "swappiness": {
                    "computed": true,
                    "description": "This allows you to tune a container's memory swappiness behavior. A ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` will cause swapping to not happen unless absolutely necessary. A ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` value of ` + "`" + `` + "`" + `100` + "`" + `` + "`" + ` will cause pages to be swapped very aggressively. Accepted values are whole numbers between ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `100` + "`" + `` + "`" + `. If the ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` parameter is not specified, a default value of ` + "`" + `` + "`" + `60` + "`" + `` + "`" + ` is used. If a value is not specified for ` + "`" + `` + "`" + `maxSwap` + "`" + `` + "`" + ` then this parameter is ignored. This parameter maps to the ` + "`" + `` + "`" + `--memory-swappiness` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  If you're using tasks that use the Fargate launch type, the ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` parameter isn't supported.\n If you're using tasks on Amazon Linux 2023 the ` + "`" + `` + "`" + `swappiness` + "`" + `` + "`" + ` parameter isn't supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "tmpfs": {
                    "computed": true,
                    "description": "The container path, mount options, and size (in MiB) of the tmpfs mount. This parameter maps to the ` + "`" + `` + "`" + `--tmpfs` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  If you're using tasks that use the Fargate launch type, the ` + "`" + `` + "`" + `tmpfs` + "`" + `` + "`" + ` parameter isn't supported.",
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
                          "description": "The maximum size (in MiB) of the tmpfs volume.",
                          "description_kind": "plain",
                          "required": true,
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
              "description": "The log configuration specification for the container.\n This parameter maps to ` + "`" + `` + "`" + `LogConfig` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--log-driver` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/). By default, containers use the same logging driver that the Docker daemon uses. However, the container may use a different logging driver than the Docker daemon by specifying a log driver with this parameter in the container definition. To use a different logging driver for a container, the log system must be configured properly on the container instance (or on a different log server for remote logging options). For more information on the options for different supported log drivers, see [Configure logging drivers](https://docs.aws.amazon.com/https://docs.docker.com/engine/admin/logging/overview/) in the Docker documentation.\n  Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon (shown in the [LogConfiguration](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_LogConfiguration.html) data type). Additional log drivers may be available in future releases of the Amazon ECS container agent.\n  This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + ` \n  The Amazon ECS container agent running on a container instance must register the logging drivers available on that instance with the ` + "`" + `` + "`" + `ECS_AVAILABLE_LOGGING_DRIVERS` + "`" + `` + "`" + ` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS Container Agent Configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_driver": {
                    "description": "The log driver to use for the container.\n For tasks on FARGATElong, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For tasks hosted on Amazon EC2 instances, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `fluentd` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `gelf` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `json-file` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `journald` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `logentries` + "`" + `` + "`" + `,` + "`" + `` + "`" + `syslog` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For more information about using the ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` log driver, see [Using the awslogs log driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html) in the *Amazon Elastic Container Service Developer Guide*.\n For more information about using the ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` log driver, see [Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide*.\n  If you have a custom driver that isn't listed, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you would like to have included. However, we don't currently provide support for running modified copies of this software.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The configuration options to send to the log driver. This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + `",
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
                          "description": "The name of the secret.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value_from": {
                          "description": "The secret to expose to the container. The supported values are either the full ARN of the ASMlong secret or the full ARN of the parameter in the SSM Parameter Store.\n For information about the require IAMlong permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide*.\n  If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.",
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
            "memory": {
              "computed": true,
              "description": "The amount (in MiB) of memory to present to the container. If your container attempts to exceed the memory specified here, the container is killed. The total amount of memory reserved for all containers within a task must be lower than the task ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` value, if one is specified. This parameter maps to ` + "`" + `` + "`" + `Memory` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--memory` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n If using the Fargate launch type, this parameter is optional.\n If using the EC2 launch type, you must specify either a task-level memory value or a container-level memory value. If you specify both a container-level ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + ` value, ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` must be greater than ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + `. If you specify ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + `, then that value is subtracted from the available memory resources for the container instance where the container is placed. Otherwise, the value of ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` is used.\n The Docker 20.10.0 or later daemon reserves a minimum of 6 MiB of memory for a container, so you should not specify fewer than 6 MiB of memory for your containers.\n The Docker 19.03.13-ce or earlier daemon reserves a minimum of 4 MiB of memory for a container, so you should not specify fewer than 4 MiB of memory for your containers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory_reservation": {
              "computed": true,
              "description": "The soft limit (in MiB) of memory to reserve for the container. When system memory is under heavy contention, Docker attempts to keep the container memory to this soft limit. However, your container can consume more memory when it needs to, up to either the hard limit specified with the ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` parameter (if applicable), or all of the available memory on the container instance, whichever comes first. This parameter maps to ` + "`" + `` + "`" + `MemoryReservation` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--memory-reservation` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n If a task-level memory value is not specified, you must specify a non-zero integer for one or both of ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + ` in a container definition. If you specify both, ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` must be greater than ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + `. If you specify ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + `, then that value is subtracted from the available memory resources for the container instance where the container is placed. Otherwise, the value of ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` is used.\n For example, if your container normally uses 128 MiB of memory, but occasionally bursts to 256 MiB of memory for short periods of time, you can set a ` + "`" + `` + "`" + `memoryReservation` + "`" + `` + "`" + ` of 128 MiB, and a ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` hard limit of 300 MiB. This configuration would allow the container to only reserve 128 MiB of memory from the remaining resources on the container instance, but also allow the container to consume more memory resources when needed.\n The Docker 20.10.0 or later daemon reserves a minimum of 6 MiB of memory for a container. So, don't specify less than 6 MiB of memory for your containers. \n The Docker 19.03.13-ce or earlier daemon reserves a minimum of 4 MiB of memory for a container. So, don't specify less than 4 MiB of memory for your containers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mount_points": {
              "computed": true,
              "description": "The mount points for data volumes in your container.\n This parameter maps to ` + "`" + `` + "`" + `Volumes` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--volume` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n Windows containers can mount whole directories on the same drive as ` + "`" + `` + "`" + `$env:ProgramData` + "`" + `` + "`" + `. Windows containers can't mount directories on a different drive, and mount point can't be across drives.",
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
              "description": "The name of a container. If you're linking multiple containers together in a task definition, the ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` of one container can be entered in the ` + "`" + `` + "`" + `links` + "`" + `` + "`" + ` of another container to connect the containers. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. This parameter maps to ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--name` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port_mappings": {
              "computed": true,
              "description": "The list of port mappings for the container. Port mappings allow containers to access ports on the host container instance to send or receive traffic.\n For task definitions that use the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode, you should only specify the ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + `. The ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + ` can be left blank or it must be the same value as the ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + `.\n Port mappings on Windows use the ` + "`" + `` + "`" + `NetNAT` + "`" + `` + "`" + ` gateway address rather than ` + "`" + `` + "`" + `localhost` + "`" + `` + "`" + `. There is no loopback for port mappings on Windows, so you cannot access a container's mapped port from the host itself. \n This parameter maps to ` + "`" + `` + "`" + `PortBindings` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--publish` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/). If the network mode of a task definition is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, then you can't specify port mappings. If the network mode of a task definition is set to ` + "`" + `` + "`" + `host` + "`" + `` + "`" + `, then host ports must either be undefined or they must match the container port in the port mapping.\n  After a task reaches the ` + "`" + `` + "`" + `RUNNING` + "`" + `` + "`" + ` status, manual and automatic host and container port assignments are visible in the *Network Bindings* section of a container description for a selected task in the Amazon ECS console. The assignments are also visible in the ` + "`" + `` + "`" + `networkBindings` + "`" + `` + "`" + ` section [DescribeTasks](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeTasks.html) responses.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_protocol": {
                    "computed": true,
                    "description": "The application protocol that's used for the port mapping. This parameter only applies to Service Connect. We recommend that you set this parameter to be consistent with the protocol that your application uses. If you set this parameter, Amazon ECS adds protocol-specific connection handling to the Service Connect proxy. If you set this parameter, Amazon ECS adds protocol-specific telemetry in the Amazon ECS console and CloudWatch.\n If you don't set a value for this parameter, then TCP is used. However, Amazon ECS doesn't add protocol-specific telemetry for TCP.\n  ` + "`" + `` + "`" + `appProtocol` + "`" + `` + "`" + ` is immutable in a Service Connect service. Updating this field requires a service deletion and redeployment.\n Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_port": {
                    "computed": true,
                    "description": "The port number on the container that's bound to the user-specified or automatically assigned host port.\n If you use containers in a task with the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode, specify the exposed ports using ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + `.\n If you use containers in a task with the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` network mode and you specify a container port and not a host port, your container automatically receives a host port in the ephemeral port range. For more information, see ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + `. Port mappings that are automatically assigned in this way do not count toward the 100 reserved ports limit of a container instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "container_port_range": {
                    "computed": true,
                    "description": "The port number range on the container that's bound to the dynamically mapped host port range. \n The following rules apply when you specify a ` + "`" + `` + "`" + `containerPortRange` + "`" + `` + "`" + `:\n  +  You must use either the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` network mode or the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode.\n  +  This parameter is available for both the EC2 and FARGATElong launch types.\n  +  This parameter is available for both the Linux and Windows operating systems.\n  +  The container instance must have at least version 1.67.0 of the container agent and at least version 1.67.0-1 of the ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + ` package \n  +  You can specify a maximum of 100 port ranges per container.\n  +  You do not specify a ` + "`" + `` + "`" + `hostPortRange` + "`" + `` + "`" + `. The value of the ` + "`" + `` + "`" + `hostPortRange` + "`" + `` + "`" + ` is set as follows:\n  +  For containers in a task with the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode, the ` + "`" + `` + "`" + `hostPortRange` + "`" + `` + "`" + ` is set to the same value as the ` + "`" + `` + "`" + `containerPortRange` + "`" + `` + "`" + `. This is a static mapping strategy.\n  +  For containers in a task with the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` network mode, the Amazon ECS agent finds open host ports from the default ephemeral range and passes it to docker to bind them to the container ports.\n  \n  +  The ` + "`" + `` + "`" + `containerPortRange` + "`" + `` + "`" + ` valid values are between 1 and 65535.\n  +  A port can only be included in one port mapping per container.\n  +  You cannot specify overlapping port ranges.\n  +  The first port in the range must be less than last port in the range.\n  +  Docker recommends that you turn off the docker-proxy in the Docker daemon config file when you have a large number of ports.\n For more information, see [Issue #11185](https://docs.aws.amazon.com/https://github.com/moby/moby/issues/11185) on the Github website.\n For information about how to turn off the docker-proxy in the Docker daemon config file, see [Docker daemon](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/bootstrap_container_instance.html#bootstrap_docker_daemon) in the *Amazon ECS Developer Guide*.\n  \n You can call [DescribeTasks](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeTasks.html) to view the ` + "`" + `` + "`" + `hostPortRange` + "`" + `` + "`" + ` which are the host ports that are bound to the container ports.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "host_port": {
                    "computed": true,
                    "description": "The port number on the container instance to reserve for your container.\n If you specify a ` + "`" + `` + "`" + `containerPortRange` + "`" + `` + "`" + `, leave this field empty and the value of the ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + ` is set as follows:\n  +  For containers in a task with the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode, the ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + ` is set to the same value as the ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + `. This is a static mapping strategy.\n  +  For containers in a task with the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` network mode, the Amazon ECS agent finds open ports on the host and automatically binds them to the container ports. This is a dynamic mapping strategy.\n  \n If you use containers in a task with the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode, the ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + ` can either be left blank or set to the same value as the ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + `.\n If you use containers in a task with the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` network mode, you can specify a non-reserved host port for your container port mapping, or you can omit the ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + ` (or set it to ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `) while specifying a ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + ` and your container automatically receives a port in the ephemeral port range for your container instance operating system and Docker version.\n The default ephemeral port range for Docker version 1.6.0 and later is listed on the instance under ` + "`" + `` + "`" + `/proc/sys/net/ipv4/ip_local_port_range` + "`" + `` + "`" + `. If this kernel parameter is unavailable, the default ephemeral port range from 49153 through 65535 (Linux) or 49152 through 65535 (Windows) is used. Do not attempt to specify a host port in the ephemeral port range as these are reserved for automatic assignment. In general, ports below 32768 are outside of the ephemeral port range.\n The default reserved ports are 22 for SSH, the Docker ports 2375 and 2376, and the Amazon ECS container agent ports 51678-51680. Any host port that was previously specified in a running task is also reserved while the task is running. That is, after a task stops, the host port is released. The current reserved ports are displayed in the ` + "`" + `` + "`" + `remainingResources` + "`" + `` + "`" + ` of [DescribeContainerInstances](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DescribeContainerInstances.html) output. A container instance can have up to 100 reserved ports at a time. This number includes the default reserved ports. Automatically assigned ports aren't included in the 100 reserved ports quota.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name that's used for the port mapping. This parameter only applies to Service Connect. This parameter is the name that you use in the ` + "`" + `` + "`" + `serviceConnectConfiguration` + "`" + `` + "`" + ` of a service. The name can include up to 64 characters. The characters can include lowercase letters, numbers, underscores (_), and hyphens (-). The name can't start with a hyphen.\n For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "The protocol used for the port mapping. Valid values are ` + "`" + `` + "`" + `tcp` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `udp` + "`" + `` + "`" + `. The default is ` + "`" + `` + "`" + `tcp` + "`" + `` + "`" + `. ` + "`" + `` + "`" + `protocol` + "`" + `` + "`" + ` is immutable in a Service Connect service. Updating this field requires a service deletion and redeployment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "privileged": {
              "computed": true,
              "description": "When this parameter is true, the container is given elevated privileges on the host container instance (similar to the ` + "`" + `` + "`" + `root` + "`" + `` + "`" + ` user). This parameter maps to ` + "`" + `` + "`" + `Privileged` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--privileged` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  This parameter is not supported for Windows containers or tasks run on FARGATElong.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pseudo_terminal": {
              "computed": true,
              "description": "When this parameter is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, a TTY is allocated. This parameter maps to ` + "`" + `` + "`" + `Tty` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--tty` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "readonly_root_filesystem": {
              "computed": true,
              "description": "When this parameter is true, the container is given read-only access to its root file system. This parameter maps to ` + "`" + `` + "`" + `ReadonlyRootfs` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--read-only` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  This parameter is not supported for Windows containers.",
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
            "resource_requirements": {
              "computed": true,
              "description": "The type and amount of a resource to assign to a container. The only supported resource is a GPU.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "description": "The type of resource to assign to a container. The supported values are ` + "`" + `` + "`" + `GPU` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `InferenceAccelerator` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value for the specified resource type.\n If the ` + "`" + `` + "`" + `GPU` + "`" + `` + "`" + ` type is used, the value is the number of physical ` + "`" + `` + "`" + `GPUs` + "`" + `` + "`" + ` the Amazon ECS container agent reserves for the container. The number of GPUs that's reserved for all containers in a task can't exceed the number of available GPUs on the container instance that the task is launched on.\n If the ` + "`" + `` + "`" + `InferenceAccelerator` + "`" + `` + "`" + ` type is used, the ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` matches the ` + "`" + `` + "`" + `deviceName` + "`" + `` + "`" + ` for an [InferenceAccelerator](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_InferenceAccelerator.html) specified in a task definition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "secrets": {
              "computed": true,
              "description": "The secrets to pass to the container. For more information, see [Specifying Sensitive Data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide*.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "description": "The name of the secret.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value_from": {
                    "description": "The secret to expose to the container. The supported values are either the full ARN of the ASMlong secret or the full ARN of the parameter in the SSM Parameter Store.\n For information about the require IAMlong permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide*.\n  If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "start_timeout": {
              "computed": true,
              "description": "Time duration (in seconds) to wait before giving up on resolving dependencies for a container. For example, you specify two containers in a task definition with containerA having a dependency on containerB reaching a ` + "`" + `` + "`" + `COMPLETE` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `SUCCESS` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `HEALTHY` + "`" + `` + "`" + ` status. If a ` + "`" + `` + "`" + `startTimeout` + "`" + `` + "`" + ` value is specified for containerB and it doesn't reach the desired status within that time then containerA gives up and not start. This results in the task transitioning to a ` + "`" + `` + "`" + `STOPPED` + "`" + `` + "`" + ` state.\n  When the ` + "`" + `` + "`" + `ECS_CONTAINER_START_TIMEOUT` + "`" + `` + "`" + ` container agent configuration variable is used, it's enforced independently from this start timeout value.\n  For tasks using the Fargate launch type, the task or service requires the following platforms:\n  +  Linux platform version ` + "`" + `` + "`" + `1.3.0` + "`" + `` + "`" + ` or later.\n  +  Windows platform version ` + "`" + `` + "`" + `1.0.0` + "`" + `` + "`" + ` or later.\n  \n For tasks using the EC2 launch type, your container instances require at least version ` + "`" + `` + "`" + `1.26.0` + "`" + `` + "`" + ` of the container agent to use a container start timeout value. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide*. If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version ` + "`" + `` + "`" + `1.26.0-1` + "`" + `` + "`" + ` of the ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + ` package. If your container instances are launched from version ` + "`" + `` + "`" + `20190301` + "`" + `` + "`" + ` or later, then they contain the required versions of the container agent and ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + `. For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide*.\n The valid values are 2-120 seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stop_timeout": {
              "computed": true,
              "description": "Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.\n For tasks using the Fargate launch type, the task or service requires the following platforms:\n  +  Linux platform version ` + "`" + `` + "`" + `1.3.0` + "`" + `` + "`" + ` or later.\n  +  Windows platform version ` + "`" + `` + "`" + `1.0.0` + "`" + `` + "`" + ` or later.\n  \n The max stop timeout value is 120 seconds and if the parameter is not specified, the default value of 30 seconds is used.\n For tasks that use the EC2 launch type, if the ` + "`" + `` + "`" + `stopTimeout` + "`" + `` + "`" + ` parameter isn't specified, the value set for the Amazon ECS container agent configuration variable ` + "`" + `` + "`" + `ECS_CONTAINER_STOP_TIMEOUT` + "`" + `` + "`" + ` is used. If neither the ` + "`" + `` + "`" + `stopTimeout` + "`" + `` + "`" + ` parameter or the ` + "`" + `` + "`" + `ECS_CONTAINER_STOP_TIMEOUT` + "`" + `` + "`" + ` agent configuration variable are set, then the default values of 30 seconds for Linux containers and 30 seconds on Windows containers are used. Your container instances require at least version 1.26.0 of the container agent to use a container stop timeout value. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide*. If you're using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + ` package. If your container instances are launched from version ` + "`" + `` + "`" + `20190301` + "`" + `` + "`" + ` or later, then they contain the required versions of the container agent and ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + `. For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide*.\n The valid values are 2-120 seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "system_controls": {
              "computed": true,
              "description": "A list of namespaced kernel parameters to set in the container. This parameter maps to ` + "`" + `` + "`" + `Sysctls` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--sysctl` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration). For example, you can configure ` + "`" + `` + "`" + `net.ipv4.tcp_keepalive_time` + "`" + `` + "`" + ` setting to maintain longer lived connections.",
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
                    "description": "The namespaced kernel parameter to set a ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` for.\n Valid IPC namespace values: ` + "`" + `` + "`" + `\"kernel.msgmax\" | \"kernel.msgmnb\" | \"kernel.msgmni\" | \"kernel.sem\" | \"kernel.shmall\" | \"kernel.shmmax\" | \"kernel.shmmni\" | \"kernel.shm_rmid_forced\"` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `Sysctls` + "`" + `` + "`" + ` that start with ` + "`" + `` + "`" + `\"fs.mqueue.*\"` + "`" + `` + "`" + ` \n Valid network namespace values: ` + "`" + `` + "`" + `Sysctls` + "`" + `` + "`" + ` that start with ` + "`" + `` + "`" + `\"net.*\"` + "`" + `` + "`" + ` \n All of these values are supported by Fargate.",
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
              "description": "A list of ` + "`" + `` + "`" + `ulimits` + "`" + `` + "`" + ` to set in the container. This parameter maps to ` + "`" + `` + "`" + `Ulimits` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--ulimit` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/). Valid naming values are displayed in the [Ulimit](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Ulimit.html) data type. This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + ` \n  This parameter is not supported for Windows containers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hard_limit": {
                    "description": "The hard limit for the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + ` type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "name": {
                    "description": "The ` + "`" + `` + "`" + `type` + "`" + `` + "`" + ` of the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "soft_limit": {
                    "description": "The soft limit for the ` + "`" + `` + "`" + `ulimit` + "`" + `` + "`" + ` type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "user": {
              "computed": true,
              "description": "The user to use inside the container. This parameter maps to ` + "`" + `` + "`" + `User` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--user` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).\n  When running tasks using the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode, don't run containers using the root user (UID 0). We recommend using a non-root user for better security.\n  You can specify the ` + "`" + `` + "`" + `user` + "`" + `` + "`" + ` using the following formats. If specifying a UID or GID, you must specify it as a positive integer.\n  +   ` + "`" + `` + "`" + `user` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `user:group` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `uid` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `uid:gid` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `user:gid` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `uid:group` + "`" + `` + "`" + ` \n  \n  This parameter is not supported for Windows containers.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "volumes_from": {
              "computed": true,
              "description": "Data volumes to mount from another container. This parameter maps to ` + "`" + `` + "`" + `VolumesFrom` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--volumes-from` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_only": {
                    "computed": true,
                    "description": "If this value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the container has read-only access to the volume. If this value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, then the container can write to the volume. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_container": {
                    "computed": true,
                    "description": "The name of another container within the same task definition to mount volumes from.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "working_directory": {
              "computed": true,
              "description": "The working directory to run commands inside the container in. This parameter maps to ` + "`" + `` + "`" + `WorkingDir` + "`" + `` + "`" + ` in the [Create a container](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `--workdir` + "`" + `` + "`" + ` option to [docker run](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#security-configuration).",
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
        "description": "The number of ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` units used by the task. If you use the EC2 launch type, this field is optional. Any value can be used. If you use the Fargate launch type, this field is required. You must use one of the following values. The value that you choose determines your range of valid values for the ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` parameter.\n The CPU units cannot be less than 1 vCPU when you use Windows containers on Fargate.\n  +  256 (.25 vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)\n  +  512 (.5 vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)\n  +  1024 (1 vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)\n  +  2048 (2 vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)\n  +  4096 (4 vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB)\n  +  8192 (8 vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 16 GB and 60 GB in 4 GB increments\n This option requires Linux platform ` + "`" + `` + "`" + `1.4.0` + "`" + `` + "`" + ` or later.\n  +  16384 (16vCPU) - Available ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + ` values: 32GB and 120 GB in 8 GB increments\n This option requires Linux platform ` + "`" + `` + "`" + `1.4.0` + "`" + `` + "`" + ` or later.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ephemeral_storage": {
        "computed": true,
        "description": "The ephemeral storage settings to use for tasks run with the task definition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "size_in_gi_b": {
              "computed": true,
              "description": "The total amount, in GiB, of ephemeral storage to set for the task. The minimum supported value is ` + "`" + `` + "`" + `20` + "`" + `` + "`" + ` GiB and the maximum supported value is ` + "`" + `` + "`" + `200` + "`" + `` + "`" + ` GiB.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make AWS API calls on your behalf. The task execution IAM role is required depending on the requirements of your task. For more information, see [Amazon ECS task execution IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_execution_IAM_role.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "family": {
        "computed": true,
        "description": "The name of a family that this task definition is registered to. Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.\n A family groups multiple versions of a task definition. Amazon ECS gives the first task definition that you registered to a family a revision number of 1. Amazon ECS gives sequential revision numbers to each task definition that you add.\n  To use revision numbers when you update a task definition, specify this property. If you don't specify a value, CFNlong generates a new task definition each time that you update it.",
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
      "inference_accelerators": {
        "computed": true,
        "description": "The Elastic Inference accelerators to use for the containers in the task.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "device_name": {
              "computed": true,
              "description": "The Elastic Inference accelerator device name. The ` + "`" + `` + "`" + `deviceName` + "`" + `` + "`" + ` must also be referenced in a container definition as a [ResourceRequirement](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ResourceRequirement.html).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "device_type": {
              "computed": true,
              "description": "The Elastic Inference accelerator type to use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "ipc_mode": {
        "computed": true,
        "description": "The IPC resource namespace to use for the containers in the task. The valid values are ` + "`" + `` + "`" + `host` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `task` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `. If ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` is specified, then all containers within the tasks that specified the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` IPC mode on the same container instance share the same IPC resources with the host Amazon EC2 instance. If ` + "`" + `` + "`" + `task` + "`" + `` + "`" + ` is specified, all containers within the specified task share the same IPC resources. If ` + "`" + `` + "`" + `none` + "`" + `` + "`" + ` is specified, then IPC resources within the containers of a task are private and not shared with other containers in a task or on the container instance. If no value is specified, then the IPC resource namespace sharing depends on the Docker daemon setting on the container instance. For more information, see [IPC settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#ipc-settings---ipc) in the *Docker run reference*.\n If the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` IPC mode is used, be aware that there is a heightened risk of undesired IPC namespace expose. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/).\n If you are setting namespaced kernel parameters using ` + "`" + `` + "`" + `systemControls` + "`" + `` + "`" + ` for the containers in the task, the following will apply to your IPC resource namespace. For more information, see [System Controls](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) in the *Amazon Elastic Container Service Developer Guide*.\n  +  For tasks that use the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` IPC mode, IPC namespace related ` + "`" + `` + "`" + `systemControls` + "`" + `` + "`" + ` are not supported.\n  +  For tasks that use the ` + "`" + `` + "`" + `task` + "`" + `` + "`" + ` IPC mode, IPC namespace related ` + "`" + `` + "`" + `systemControls` + "`" + `` + "`" + ` will apply to all containers within a task.\n  \n  This parameter is not supported for Windows containers or tasks run on FARGATElong.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "memory": {
        "computed": true,
        "description": "The amount (in MiB) of memory used by the task.\n If your tasks runs on Amazon EC2 instances, you must specify either a task-level memory value or a container-level memory value. This field is optional and any value can be used. If a task-level memory value is specified, the container-level memory value is optional. For more information regarding container-level memory and memory reservation, see [ContainerDefinition](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html).\n If your tasks runs on FARGATElong, this field is required. You must use one of the following values. The value you choose determines your range of valid values for the ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` parameter.\n  +  512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 256 (.25 vCPU)\n  +  1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 512 (.5 vCPU)\n  +  2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 1024 (1 vCPU)\n  +  Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 2048 (2 vCPU)\n  +  Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 4096 (4 vCPU)\n  +  Between 16 GB and 60 GB in 4 GB increments - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 8192 (8 vCPU)\n This option requires Linux platform ` + "`" + `` + "`" + `1.4.0` + "`" + `` + "`" + ` or later.\n  +  Between 32GB and 120 GB in 8 GB increments - Available ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` values: 16384 (16 vCPU)\n This option requires Linux platform ` + "`" + `` + "`" + `1.4.0` + "`" + `` + "`" + ` or later.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_mode": {
        "computed": true,
        "description": "The Docker networking mode to use for the containers in the task. The valid values are ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `host` + "`" + `` + "`" + `. If no network mode is specified, the default is ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + `.\n For Amazon ECS tasks on Fargate, the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode is required. For Amazon ECS tasks on Amazon EC2 Linux instances, any network mode can be used. For Amazon ECS tasks on Amazon EC2 Windows instances, ` + "`" + `` + "`" + `\u003cdefault\u003e` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` can be used. If the network mode is set to ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `, you cannot specify port mappings in your container definitions, and the tasks containers do not have external connectivity. The ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network modes offer the highest networking performance for containers because they use the EC2 network stack instead of the virtualized network stack provided by the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` mode.\n With the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network modes, exposed container ports are mapped directly to the corresponding host port (for the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode) or the attached elastic network interface port (for the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode), so you cannot take advantage of dynamic host port mappings. \n  When using the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode, you should not run containers using the root user (UID 0). It is considered best practice to use a non-root user.\n  If the network mode is ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + `, the task is allocated an elastic network interface, and you must specify a NetworkConfiguration value when you create a service or run a task with the task definition. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide*.\n If the network mode is ` + "`" + `` + "`" + `host` + "`" + `` + "`" + `, you cannot run multiple instantiations of the same task on a single container instance when port mappings are used.\n For more information, see [Network settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#network-settings) in the *Docker run reference*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pid_mode": {
        "computed": true,
        "description": "The process namespace to use for the containers in the task. The valid values are ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `task` + "`" + `` + "`" + `. On Fargate for Linux containers, the only valid value is ` + "`" + `` + "`" + `task` + "`" + `` + "`" + `. For example, monitoring sidecars might need ` + "`" + `` + "`" + `pidMode` + "`" + `` + "`" + ` to access information about other containers running in the same task.\n If ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` is specified, all containers within the tasks that specified the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` PID mode on the same container instance share the same process namespace with the host Amazon EC2 instance.\n If ` + "`" + `` + "`" + `task` + "`" + `` + "`" + ` is specified, all containers within the specified task share the same process namespace.\n If no value is specified, the default is a private namespace for each container. For more information, see [PID settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#pid-settings---pid) in the *Docker run reference*.\n If the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` PID mode is used, there's a heightened risk of undesired process namespace exposure. For more information, see [Docker security](https://docs.aws.amazon.com/https://docs.docker.com/engine/security/security/).\n  This parameter is not supported for Windows containers.\n   This parameter is only supported for tasks that are hosted on FARGATElong if the tasks are using platform version ` + "`" + `` + "`" + `1.4.0` + "`" + `` + "`" + ` or later (Linux). This isn't supported for Windows containers on Fargate.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "placement_constraints": {
        "computed": true,
        "description": "An array of placement constraint objects to use for tasks.\n  This parameter isn't supported for tasks run on FARGATElong.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expression": {
              "computed": true,
              "description": "A cluster query language expression to apply to the constraint. For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "The type of constraint. The ` + "`" + `` + "`" + `MemberOf` + "`" + `` + "`" + ` constraint restricts selection to be from a group of valid candidates.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "proxy_configuration": {
        "computed": true,
        "description": "The configuration details for the App Mesh proxy.\n Your Amazon ECS container instances require at least version 1.26.0 of the container agent and at least version 1.26.0-1 of the ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + ` package to use a proxy configuration. If your container instances are launched from the Amazon ECS optimized AMI version ` + "`" + `` + "`" + `20190301` + "`" + `` + "`" + ` or later, they contain the required versions of the container agent and ` + "`" + `` + "`" + `ecs-init` + "`" + `` + "`" + `. For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "description": "The name of the container that will serve as the App Mesh proxy.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "proxy_configuration_properties": {
              "computed": true,
              "description": "The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified as key-value pairs.\n  +   ` + "`" + `` + "`" + `IgnoredUID` + "`" + `` + "`" + ` - (Required) The user ID (UID) of the proxy container as defined by the ` + "`" + `` + "`" + `user` + "`" + `` + "`" + ` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If ` + "`" + `` + "`" + `IgnoredGID` + "`" + `` + "`" + ` is specified, this field can be empty.\n  +   ` + "`" + `` + "`" + `IgnoredGID` + "`" + `` + "`" + ` - (Required) The group ID (GID) of the proxy container as defined by the ` + "`" + `` + "`" + `user` + "`" + `` + "`" + ` parameter in a container definition. This is used to ensure the proxy ignores its own traffic. If ` + "`" + `` + "`" + `IgnoredUID` + "`" + `` + "`" + ` is specified, this field can be empty.\n  +   ` + "`" + `` + "`" + `AppPorts` + "`" + `` + "`" + ` - (Required) The list of ports that the application uses. Network traffic to these ports is forwarded to the ` + "`" + `` + "`" + `ProxyIngressPort` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `ProxyEgressPort` + "`" + `` + "`" + `.\n  +   ` + "`" + `` + "`" + `ProxyIngressPort` + "`" + `` + "`" + ` - (Required) Specifies the port that incoming traffic to the ` + "`" + `` + "`" + `AppPorts` + "`" + `` + "`" + ` is directed to.\n  +   ` + "`" + `` + "`" + `ProxyEgressPort` + "`" + `` + "`" + ` - (Required) Specifies the port that outgoing traffic from the ` + "`" + `` + "`" + `AppPorts` + "`" + `` + "`" + ` is directed to.\n  +   ` + "`" + `` + "`" + `EgressIgnoredPorts` + "`" + `` + "`" + ` - (Required) The egress traffic going to the specified ports is ignored and not redirected to the ` + "`" + `` + "`" + `ProxyEgressPort` + "`" + `` + "`" + `. It can be an empty list.\n  +   ` + "`" + `` + "`" + `EgressIgnoredIPs` + "`" + `` + "`" + ` - (Required) The egress traffic going to the specified IP addresses is ignored and not redirected to the ` + "`" + `` + "`" + `ProxyEgressPort` + "`" + `` + "`" + `. It can be an empty list.",
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
                "nesting_mode": "set"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description": "The proxy type. The only supported value is ` + "`" + `` + "`" + `APPMESH` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "requires_compatibilities": {
        "computed": true,
        "description": "The task launch types the task definition was validated against. The valid values are ` + "`" + `` + "`" + `EC2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `FARGATE` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `. For more information, see [Amazon ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "runtime_platform": {
        "computed": true,
        "description": "The operating system that your tasks definitions run on. A platform family is specified only for tasks using the Fargate launch type. \n When you specify a task definition in a service, this value must match the ` + "`" + `` + "`" + `runtimePlatform` + "`" + `` + "`" + ` value of the service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu_architecture": {
              "computed": true,
              "description": "The CPU architecture.\n You can run your Linux tasks on an ARM-based platform by setting the value to ` + "`" + `` + "`" + `ARM64` + "`" + `` + "`" + `. This option is available for tasks that run on Linux Amazon EC2 instance or Linux containers on Fargate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "operating_system_family": {
              "computed": true,
              "description": "The operating system.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to the task definition to help you categorize and organize them. Each tag consists of a key and an optional value. You define both of them.\n The following basic restrictions apply to tags:\n  +  Maximum number of tags per resource - 50\n  +  For each resource, each tag key must be unique, and each tag key can have only one value.\n  +  Maximum key length - 128 Unicode characters in UTF-8\n  +  Maximum value length - 256 Unicode characters in UTF-8\n  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.\n  +  Tag keys and values are case-sensitive.\n  +  Do not use ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AWS:` + "`" + `` + "`" + `, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.",
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
      "task_definition_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "task_role_arn": {
        "computed": true,
        "description": "The short name or full Amazon Resource Name (ARN) of the IAMlong role that grants containers in the task permission to call AWS APIs on your behalf. For more information, see [Amazon ECS Task Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide*.\n IAM roles for tasks on Windows require that the ` + "`" + `` + "`" + `-EnableTaskIAMRole` + "`" + `` + "`" + ` option is set when you launch the Amazon ECS-optimized Windows AMI. Your containers must also run some configuration code to use the feature. For more information, see [Windows IAM roles for tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/windows_task_IAM_roles.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volumes": {
        "computed": true,
        "description": "The list of data volume definitions for the task. For more information, see [Using data volumes in tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) in the *Amazon Elastic Container Service Developer Guide*.\n  The ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `sourcePath` + "`" + `` + "`" + ` parameters aren't supported for tasks run on FARGATElong.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configured_at_launch": {
              "computed": true,
              "description": "Indicates whether the volume should be configured at launch time. This is used to create Amazon EBS volumes for standalone tasks or tasks created as part of a service. Each task definition revision may only have one volume configured at launch in the volume configuration.\n To configure a volume at launch time, use this task definition revision and specify a ` + "`" + `` + "`" + `volumeConfigurations` + "`" + `` + "`" + ` object when calling the ` + "`" + `` + "`" + `CreateService` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `UpdateService` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `RunTask` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `StartTask` + "`" + `` + "`" + ` APIs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "docker_volume_configuration": {
              "computed": true,
              "description": "This parameter is specified when you use Docker volumes.\n Windows containers only support the use of the ` + "`" + `` + "`" + `local` + "`" + `` + "`" + ` driver. To use bind mounts, specify the ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` parameter instead.\n  Docker volumes aren't supported by tasks run on FARGATElong.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "autoprovision": {
                    "computed": true,
                    "description": "If this value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the Docker volume is created if it doesn't already exist.\n  This field is only used if the ` + "`" + `` + "`" + `scope` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `shared` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "driver": {
                    "computed": true,
                    "description": "The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement. If the driver was installed using the Docker plugin CLI, use ` + "`" + `` + "`" + `docker plugin ls` + "`" + `` + "`" + ` to retrieve the driver name from your container instance. If the driver was installed using another method, use Docker plugin discovery to retrieve the driver name. For more information, see [Docker plugin discovery](https://docs.aws.amazon.com/https://docs.docker.com/engine/extend/plugin_api/#plugin-discovery). This parameter maps to ` + "`" + `` + "`" + `Driver` + "`" + `` + "`" + ` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `xxdriver` + "`" + `` + "`" + ` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "driver_opts": {
                    "computed": true,
                    "description": "A map of Docker driver-specific options passed through. This parameter maps to ` + "`" + `` + "`" + `DriverOpts` + "`" + `` + "`" + ` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `xxopt` + "`" + `` + "`" + ` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "labels": {
                    "computed": true,
                    "description": "Custom metadata to add to your Docker volume. This parameter maps to ` + "`" + `` + "`" + `Labels` + "`" + `` + "`" + ` in the [Create a volume](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/#operation/VolumeCreate) section of the [Docker Remote API](https://docs.aws.amazon.com/https://docs.docker.com/engine/api/v1.35/) and the ` + "`" + `` + "`" + `xxlabel` + "`" + `` + "`" + ` option to [docker volume create](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/commandline/volume_create/).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "scope": {
                    "computed": true,
                    "description": "The scope for the Docker volume that determines its lifecycle. Docker volumes that are scoped to a ` + "`" + `` + "`" + `task` + "`" + `` + "`" + ` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as ` + "`" + `` + "`" + `shared` + "`" + `` + "`" + ` persist after the task stops.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "efs_volume_configuration": {
              "computed": true,
              "description": "This parameter is specified when you use an Amazon Elastic File System file system for task storage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description": "The authorization configuration details for the Amazon EFS file system.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_point_id": {
                          "computed": true,
                          "description": "The Amazon EFS access point ID to use. If an access point is specified, the root directory value specified in the ` + "`" + `` + "`" + `EFSVolumeConfiguration` + "`" + `` + "`" + ` must either be omitted or set to ` + "`" + `` + "`" + `/` + "`" + `` + "`" + ` which will enforce the path set on the EFS access point. If an access point is used, transit encryption must be on in the ` + "`" + `` + "`" + `EFSVolumeConfiguration` + "`" + `` + "`" + `. For more information, see [Working with Amazon EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html) in the *Amazon Elastic File System User Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "iam": {
                          "computed": true,
                          "description": "Determines whether to use the Amazon ECS task role defined in a task definition when mounting the Amazon EFS file system. If it is turned on, transit encryption must be turned on in the ` + "`" + `` + "`" + `EFSVolumeConfiguration` + "`" + `` + "`" + `. If this parameter is omitted, the default value of ` + "`" + `` + "`" + `DISABLED` + "`" + `` + "`" + ` is used. For more information, see [Using Amazon EFS access points](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#efs-volume-accesspoints) in the *Amazon Elastic Container Service Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "filesystem_id": {
                    "description": "The Amazon EFS file system ID to use.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "root_directory": {
                    "computed": true,
                    "description": "The directory within the Amazon EFS file system to mount as the root directory inside the host. If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying ` + "`" + `` + "`" + `/` + "`" + `` + "`" + ` will have the same effect as omitting this parameter.\n  If an EFS access point is specified in the ` + "`" + `` + "`" + `authorizationConfig` + "`" + `` + "`" + `, the root directory parameter must either be omitted or set to ` + "`" + `` + "`" + `/` + "`" + `` + "`" + ` which will enforce the path set on the EFS access point.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transit_encryption": {
                    "computed": true,
                    "description": "Determines whether to use encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be turned on if Amazon EFS IAM authorization is used. If this parameter is omitted, the default value of ` + "`" + `` + "`" + `DISABLED` + "`" + `` + "`" + ` is used. For more information, see [Encrypting data in transit](https://docs.aws.amazon.com/efs/latest/ug/encryption-in-transit.html) in the *Amazon Elastic File System User Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transit_encryption_port": {
                    "computed": true,
                    "description": "The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server. If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses. For more information, see [EFS mount helper](https://docs.aws.amazon.com/efs/latest/ug/efs-mount-helper.html) in the *Amazon Elastic File System User Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "fsx_windows_file_server_volume_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "credentials_parameter": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "domain": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "file_system_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "root_directory": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
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
              "description": "The name of the volume. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed.\n When using a volume configured at launch, the ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` is required and must also be specified as the volume name in the ` + "`" + `` + "`" + `ServiceVolumeConfiguration` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `TaskVolumeConfiguration` + "`" + `` + "`" + ` parameter when creating your service or standalone task.\n For all other types of volumes, this name is referenced in the ` + "`" + `` + "`" + `sourceVolume` + "`" + `` + "`" + ` parameter of the ` + "`" + `` + "`" + `mountPoints` + "`" + `` + "`" + ` object in the container definition.\n When a volume is using the ` + "`" + `` + "`" + `efsVolumeConfiguration` + "`" + `` + "`" + `, the name is required.",
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
    "description": "Registers a new task definition from the supplied ` + "`" + `` + "`" + `family` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `containerDefinitions` + "`" + `` + "`" + `. Optionally, you can add data volumes to your containers with the ` + "`" + `` + "`" + `volumes` + "`" + `` + "`" + ` parameter. For more information about task definition parameters and defaults, see [Amazon ECS Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon Elastic Container Service Developer Guide*.\n You can specify a role for your task with the ` + "`" + `` + "`" + `taskRoleArn` + "`" + `` + "`" + ` parameter. When you specify a role for a task, its containers can then use the latest versions of the CLI or SDKs to make API requests to the AWS services that are specified in the policy that's associated with the role. For more information, see [IAM Roles for Tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html) in the *Amazon Elastic Container Service Developer Guide*.\n You can specify a Docker networking mode for the containers in your task definition with the ` + "`" + `` + "`" + `networkMode` + "`" + `` + "`" + ` parameter. The available network modes correspond to those described in [Network settings](https://docs.aws.amazon.com/https://docs.docker.com/engine/reference/run/#/network-settings) in the Docker run reference. If you specify the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode, the task is allocated an elastic network interface, and you must specify a NetworkConfiguration when you create a service or run a task with the task definition. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide*.\n  In the following example or examples, the Authorization header contents (` + "`" + `` + "`" + `AUTHPARAMS` + "`" + `` + "`" + `) must be replaced with an AWS Signature Version 4 signature. For more information, see [Signature Version 4 Signing Process](https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html) in the *General Reference*.\n You only need to learn how to sign HTTP requests if you intend to create them manually. When you use the [](https://docs.aws.amazon.com/cli/) or one of the [SDKs](https://docs.aws.amazon.com/tools/) to make requests to AWS, these tools automatically sign the requests for you, with the access key that you specify when you configure the tools. When you use these tools, you don't have to sign requests yourself.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsTaskDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsTaskDefinition), &result)
	return &result
}
