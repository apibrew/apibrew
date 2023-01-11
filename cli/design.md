# General Purpose
cli tool is for multiple purposes
1. interacting with data-handler from cli
2. reading writing resources/records from cli using different formats
3. doing backup/restore
4. testing connections on data-sources
5. etc.


# Design
**Main file**: dhctl

## All operations are crud like operations

```
dhctl get <type> <id/name> <other fields> <flags>
dhctl create <type> <id/name> <other fields> <flags>
dhctl update <type> <id/name> <other fields> <flags>
dhctl delete <type> <id/name> <other fields> <flags>
```

## File based operations
```
dhctl apply -f <file>
dhctl create -f <file>
dhctl update -f <file>
dhctl delete -f <file>
```

## Additional operations
```
dhctl backup -o filename --flags
dhctl restore -o filename --flags
dhctl status data-source/ds-1
```

## Configuration
```
~/dhctl.yml

resource: Config
clusters:
- cluster:
    proxy-url: http://proxy.example.org:3128
    server: https://k8s.example.org/k8s/clusters/c-xxyyzz
  name: development

users:
- name: developer

contexts:
- context:
  name: development
```


