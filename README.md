# Cloud Foundry Buildpack Usage CLI Plugin

Cloud Foundry plugin extension to view all buildpacks used in across a Cloud Foundry or in specific organizations and spaces.

## Install

```
$ go get github.com/csterwa/cf_buildpack_usage_cmd
$ cf install-plugin $GOPATH/bin/cf_buildpacks_usage_cmd
```

## Usage

```
$ cf buildpack-usage

13 buildpacks found across 110 app deployments

Buildpacks Used
----------------

Node.js
PHP
Ruby
https://github.com/cloudfoundry/java-buildpack.git
https://github.com/cloudfoundry/php-buildpack.git
https://github.com/csterwa/cf-meteor-buildpack.git
https://github.com/dmikusa-pivotal/cf-php-build-pack.git
java-buildpack=v2.6.1-https://github.com/cloudfoundry/java-buildpack.git#2d92e70 java-main open-jdk-jre=1.8.0_40
java-buildpack=v2.6.1-https://github.com/cloudfoundry/java-buildpack.git#2d92e70 open-jdk-jre=1.8.0_31 spring-auto-reconfiguration=1.7.0_RELEASE tomcat-access-logging-support=2.4.0_RELEASE tomcat-instance=8.0.20 tomcat-lifecycle-support=2.4.0_RELEASE t...
java-buildpack=v2.6.1-https://github.com/cloudfoundry/java-buildpack.git#2d92e70 open-jdk-jre=1.8.0_40 spring-auto-reconfiguration=1.7.0_RELEASE tomcat-access-logging-support=2.4.0_RELEASE tomcat-instance=8.0.20 tomcat-lifecycle-support=2.4.0_RELEASE t...
java-buildpack=v2.6.1-https://github.com/cloudfoundry/java-buildpack.git#2d92e70 open-jdk-jre=1.8.0_40 tomcat-access-logging-support=2.4.0_RELEASE tomcat-instance=8.0.20 tomcat-lifecycle-support=2.4.0_RELEASE tomcat-logging-support=2.4.0_RELEASE
nodejs_buildpack
```

## Uninstall

```
$ cf uninstall-plugin CliBuildpackUsage
```
