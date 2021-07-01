package config

// copy from go-lib
// TODO: refactor this code block

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//default config
const (
	DefaultApplication = "public"
	DefaultLabel       = "snapshot"
)

//model model
type model struct {
	Profile     string      `json:"profile"`
	Application string      `json:"application"`
	Label       string      `json:"label"`
	Key         string      `json:"key"`
	Data        interface{} `json:"data"`
}

func (p *model) getAppConfig(configCenterURL string) error {
	var api = "/configs?profile=" + p.Profile + "&application=" + p.Application + "&label=" + p.Label
	if p.Key != "" {
		api = api + "&key=" + p.Key
	}
	url := configCenterURL + api

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 200 {
		err = errors.New("read configcenter error")
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(body), p)
	if err != nil {
		return err
	}
	return err
}

func loadDataFromConfigCenter(configCenterURL, profile, serviceName string, cfg interface{}) {

	if err := getConfig(configCenterURL, profile, serviceName, "", cfg); err != nil {
	} else {
	}
}

func getConfig(configCenterURL, profile, serviceName, key string, cfg interface{}) error {
	publicConfig := newConf(profile, DefaultApplication, DefaultLabel, key, cfg)
	if err := publicConfig.getAppConfig(configCenterURL); err != nil {
		return err
	}

	appConfig := newConf(profile, serviceName, DefaultLabel, key, cfg)
	if err := appConfig.getAppConfig(configCenterURL); err != nil {
		return err
	}
	return nil
}

func newConf(profile, application, label, key string, cfg interface{}) *model {
	return &model{
		Profile:     profile,
		Application: application,
		Label:       label,
		Data:        cfg,
		Key:         key,
	}
}
