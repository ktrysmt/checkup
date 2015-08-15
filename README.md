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

## Format

- Type YAML Format at the file.
- Use XPATH to Access DOM-Element

## Example

```
testcase:
- name: sample test
  browser: firefox
  steps:
    - get: https://code.google.com/p/selenium/wiki/Buck
    - saveScreenshot: sele1.png
    - setElementSelected: 
        target: //*[@id="can"]/option[5]
    - saveScreenshot: sele2.png
```
 
## Install

```
$ go get github.com/aqafiam/checkup
```

## Feature

- Support CSS-Selector
- Support to select CSS-Selector or XPATH
- Add more example source
- write Reference for how to use
- make steps about StoreXXX command
- Use the color in STDOUT for readability

## License

MIT

## Author

[aqafiam](https://github.com/aqafiam)

