//Este código lê o dockerhub e reescreve o arquivo typeMongoDBVersionTag.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//https://mholt.github.io/json-to-go/

const (
	KTagName = "RabbitMQ"
)

type DockerHubResults struct {
	Creator             int               `json:"creator"`
	ID                  int               `json:"id"`
	ImageID             interface{}       `json:"image_id"`
	Images              []DockerHubImages `json:"images"`
	LastUpdated         time.Time         `json:"last_updated"`
	LastUpdater         int               `json:"last_updater"`
	LastUpdaterUsername string            `json:"last_updater_username"`
	Name                string            `json:"name"`
	Repository          int               `json:"repository"`
	FullSize            int               `json:"full_size"`
	V2                  bool              `json:"v2"`
}

type DockerHubImages struct {
	Architecture string      `json:"architecture"`
	Features     string      `json:"features"`
	Variant      interface{} `json:"variant"`
	Digest       string      `json:"digest"`
	Os           string      `json:"os"`
	OsFeatures   string      `json:"os_features"`
	OsVersion    string      `json:"os_version"`
	Size         int64       `json:"size"`
}

type DockerHubPageJSon struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous interface{}        `json:"previous"`
	Results  []DockerHubResults `json:"results"`
}

func main() {
	var data = populate()
	var file string
	var varName = strings.ToLower(KTagName[:1]) + KTagName[1:]

	file += "package factoryContainer" + KTagName + "\n\n"
	file += "type " + KTagName + "VersionTag int\n\n"
	file += "const (\n"
	for k, line := range data {
		if k == 0 {
			file += fmt.Sprintf("\tK"+KTagName+"VersionTag_%v "+KTagName+"VersionTag = iota\n", replace(line))
		} else {
			file += fmt.Sprintf("\tK"+KTagName+"VersionTag_%v\n", replace(line))
		}
	}
	file += ")\n\n"

	file += "func(el " + KTagName + "VersionTag)String() string {\n" +
		"\treturn " + varName + "VersionTags[el]" +
		"}\n\n"

	file += "var " + varName + "VersionTags = [...]string{\n"
	file += fmt.Sprintf("\t\"%v\",\n", strings.Join(data, "\",\n\t\""))
	file += "}\n"

	ioutil.WriteFile("./type"+KTagName+"VersionTag.go", []byte(file), os.ModePerm)
}

func populate() []string {
	var ret = make([]string, 0)

	for page := 1; page != 10; page += 1 {
		pageData := decode(page)

		for _, resultsData := range pageData.Results {
			ret = append(ret, resultsData.Name)
		}

		if pageData.Next == "" {
			return ret
		}
	}

	return ret
}

func replace(value string) string {
	value = strings.ReplaceAll(value, ".", "_")
	value = strings.ReplaceAll(value, "-", "_")

	return value
}

func decode(page int) DockerHubPageJSon {
	var jsonData []byte
	var pageData DockerHubPageJSon
	response, err := http.Get(
		fmt.Sprintf("https://hub.docker.com/v2/repositories/library/rabbitmq/tags/?page_size=100&page=%v", page),
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &pageData)
	if err != nil {
		panic(err)
	}

	return pageData
}
