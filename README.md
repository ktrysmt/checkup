# Checkup

Checkup is E2E testing tool made by Golang.  
You can write testcases easily that Syntax is YAML format.

## Usage

```
$ checkup -r http://localhost:4444/wd/hub testcase.yaml
```

also use remote yamlfile...
 
```
$ checkup -r http://localhost:4444/wd/hub testcase.yaml https://dl.dropboxusercontent.com/s/8zm9smw00oc/test.yml
```
 
## Install

```
$ go get github.com/aqafiam/checkup
```

## Author

[aqafiam](https://github.com/aqafiam)

