# Checkup

<a class="repo-badge" href="https://badge.fury.io/for/go/github.com/ktrysmt/checkup"><img src="https://badge.fury.io/go/github.com%2Fktrysmt%2Fcheckup.svg" alt="Github.com%2fktrysmt%2fcheckup"></a>
<a class="repo-badge" href="https://godoc.org/github.com/ktrysmt/checkup"><img src="https://godoc.org/github.com/ktrysmt/checkup?status.svg" alt="Checkup?status"></a>
<a class="repo-badge" href="https://goreportcard.com/report/github.com/ktrysmt/checkup"><img src="https://goreportcard.com/badge/github.com/ktrysmt/checkup" alt="Checkup"></a>

> Checkup is E2E testing tool made by Golang.  
> You can write testcases easily that Syntax is YAML format.

## Usage

```
$ checkup -r http://localhost:4444/wd/hub testcase.yaml
```

also use remote yamlfile...
 
```
$ checkup -r http://localhost:4444/wd/hub https://dl.dropboxusercontent.com/s/8zm9smw00oc/test.yml
```

## Format

- Type YAML Format at the file.
- Use XPATH or CSS-Selector to Access the DOM-Element.

## Example

```
testcase:
- name: sample test
  browser: firefox
  selector: xpath # Use "xpath" or "css selector", default is "xpath"
  steps:
    - get: https://code.google.com/p/selenium/wiki/Buck
    - saveScreenshot: sele1.png
    - setElementSelected: 
        target: //*[@id="can"]/option[5]
    - saveScreenshot: sele2.png
```
 
## Install

```
$ go get github.com/ktrysmt/checkup
```

## Future

[ ] Support CSS-Selector
[ ] Support to select CSS-Selector or XPATH
[ ] Add more example source
[ ] write Reference for how to use
[ ] make steps about StoreXXX command
[ ] Use the color in STDOUT for readability

## (Optional) Use Selenium-Hub

Run docker to use Selenium-Hub easily in your localhost.

```
docker ps -aq | xargs docker rm # remove trush at first ...
docker run -d -p 4444:4444 --name selenium-hub selenium/hub
docker run -d --link selenium-hub:hub selenium/node-chrome-debug
docker run -d --link selenium-hub:hub selenium/node-firefox-debug
```

## License

[MIT](./LICENSE)

## Author

[ktrysmt](https://github.com/ktrysmt)

