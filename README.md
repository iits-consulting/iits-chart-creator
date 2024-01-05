# iits-chart-creator

`iits-chart-creator` is a powerful cli helm tool designed to streamline and automate the process of creating and managing Helm charts for iits projects.

```shell
Available Commands:
  deployment      Creates a deployment deployment file
  helpers.tpl     Creates a helpers.tpl deployment file
  ingress         Creates a Ingress YAML deployment file
  service         Creates a service deployment file
  service-monitor Creates a service-monitor deployment file
  serviceaccount  Creates a serviceaccount deployment file
```

## Demo

[2024-01-05 16-44-27-cut-merged-1704470188632.webm](https://github.com/iits-consulting/iits-chart-creator/assets/19291722/5dead70c-ed39-4b9e-8b28-c144a6c56b9a)

## Install


```shell
helm plugin install https://github.com/iits-consulting/iits-chart-creator
#We would recommend to put this also into your .bash_aliases
alias updateCharter="helm plugin update iits-chart-creator && helm iits-chart-creator -v"
alias charter='helm iits-chart-creator infrastructure-charts'
```

### Local development
If you want to test it out locally execute the following

```shell
go build . && mv iits-chart-creator binaries/iits-chart-creator_linux_amd64_v1/ && helm plugin rm iits-chart-creator && helm plugin install .
```

## Usage


```shell
charter

Creates Infrastructure Charts

Usage:
  iits-chart-creator infrastructure-charts [command]

Available Commands:
  deployment      Creates a deployment deployment file
  helpers.tpl     Creates a helpers.tpl deployment file
  ingress         Creates a Ingress YAML deployment file
  service         Creates a service deployment file
  service-monitor Creates a service-monitor deployment file
  serviceaccount  Creates a serviceaccount deployment file

Flags:
  -h, --help   help for infrastructure-charts

Use "iits-chart-creator infrastructure-charts [command] --help" for more information about a command.
```
