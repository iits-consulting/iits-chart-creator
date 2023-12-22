# iits-chart-creator

`iits-chart-creator` is a powerful cli helm tool designed to streamline and automate the process of creating and managing Helm charts for iits projects.

Current features:
- ingress
- _helpers.tpl


## Usage

```shell
helm plugin install https://github.com/iits-consulting/iits-chart-creator
#We would recommend to put this also into your .bash_aliases
alias updateChartCreator="helm plugin update helm-chart-creator && helm iits-chart-creator -v"
```