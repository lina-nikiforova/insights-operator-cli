/*
Copyright © 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package restapi

import (
	"bytes"
	"encoding/json"
	"github.com/redhatinsighs/insights-operator-cli/types"
	"net/url"
)

// RestAPI is a structure representing instance of REST API
type RestAPI struct {
	controllerURL string
}

// NewRestAPI constructs instance of REST API
func NewRestAPI(controllerURL string) RestAPI {
	return RestAPI{
		controllerURL: controllerURL,
	}
}

// ReadListOfClusters reads list of clusters via the REST API
func (api RestAPI) ReadListOfClusters() ([]types.Cluster, error) {
	clusters := []types.Cluster{}

	url := api.controllerURL + APIPrefix + "client/cluster"
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

// ReadListOfTriggers reads list of triggers via the REST API
func (api RestAPI) ReadListOfTriggers() ([]types.Trigger, error) {
	var triggers []types.Trigger
	url := api.controllerURL + APIPrefix + "client/trigger"
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &triggers)
	if err != nil {
		return nil, err
	}
	return triggers, nil
}

// ReadTriggerByID reads trigger identified by its ID via the REST API
func (api RestAPI) ReadTriggerByID(triggerID string) (*types.Trigger, error) {
	var trigger types.Trigger
	url := api.controllerURL + APIPrefix + "client/trigger/" + triggerID
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &trigger)
	if err != nil {
		return nil, err
	}
	return &trigger, nil
}

// ReadListOfConfigurationProfiles reads list of configuration profiles via the REST API
func (api RestAPI) ReadListOfConfigurationProfiles() ([]types.ConfigurationProfile, error) {
	profiles := []types.ConfigurationProfile{}

	url := api.controllerURL + APIPrefix + "client/profile"
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

// ReadListOfConfigurations reads list of configuration via the REST API
func (api RestAPI) ReadListOfConfigurations() ([]types.ClusterConfiguration, error) {
	configurations := []types.ClusterConfiguration{}

	url := api.controllerURL + APIPrefix + "client/configuration"
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &configurations)
	if err != nil {
		return nil, err
	}
	return configurations, nil
}

// ReadConfigurationProfile access the REST API endpoint to read selected configuration profile
func (api RestAPI) ReadConfigurationProfile(profileID string) (*types.ConfigurationProfile, error) {
	var profile types.ConfigurationProfile
	url := api.controllerURL + APIPrefix + "client/profile/" + profileID
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// ReadClusterConfigurationByID access the REST API endpoint to read cluster configuration for cluster defined by its ID
func (api RestAPI) ReadClusterConfigurationByID(configurationID string) (*string, error) {
	url := api.controllerURL + APIPrefix + "client/configuration/" + configurationID
	body, err := performReadRequest(url)
	if err != nil {
		return nil, err
	}

	str := string(body)
	return &str, nil
}

// EnableClusterConfiguration access the REST API endpoint to enable existing cluster configuration
func (api RestAPI) EnableClusterConfiguration(configurationID string) error {
	url := api.controllerURL + APIPrefix + "client/configuration/" + configurationID + "/enable"
	err := performWriteRequest(url, "PUT", nil)
	return err
}

// DisableClusterConfiguration access the REST API endpoint to disable existing cluster configuration
func (api RestAPI) DisableClusterConfiguration(configurationID string) error {
	url := api.controllerURL + APIPrefix + "client/configuration/" + configurationID + "/disable"
	err := performWriteRequest(url, "PUT", nil)
	return err
}

// DeleteClusterConfiguration access the REST API endpoint to delete existing cluster configuration
func (api RestAPI) DeleteClusterConfiguration(configurationID string) error {
	url := api.controllerURL + APIPrefix + "client/configuration/" + configurationID
	err := performWriteRequest(url, "DELETE", nil)
	return err
}

// DeleteCluster access the REST API endpoint to delete/deregister existing cluster
func (api RestAPI) DeleteCluster(clusterID string) error {
	url := api.controllerURL + APIPrefix + "client/cluster/" + clusterID
	err := performWriteRequest(url, "DELETE", nil)
	return err
}

// DeleteConfigurationProfile access the REST API endpoint to delete existing configuration profile
func (api RestAPI) DeleteConfigurationProfile(profileID string) error {
	url := api.controllerURL + APIPrefix + "client/profile/" + profileID
	err := performWriteRequest(url, "DELETE", nil)
	return err
}

// AddCluster access the REST API endpoint to add/register new cluster
func (api RestAPI) AddCluster(id string, name string) error {
	query := id + "/" + name
	url := api.controllerURL + APIPrefix + "client/cluster/" + query
	err := performWriteRequest(url, "POST", nil)
	return err
}

// AddConfigurationProfile access the REST API endpoint to add new configuration profile
func (api RestAPI) AddConfigurationProfile(username string, description string, configuration []byte) error {
	query := "username=" + url.QueryEscape(username) + "&description=" + url.QueryEscape(description)
	url := api.controllerURL + APIPrefix + "client/profile?" + query
	err := performWriteRequest(url, "POST", bytes.NewReader(configuration))
	return err
}

// AddClusterConfiguration access the REST API endpoint to add new cluster configuration
func (api RestAPI) AddClusterConfiguration(username string, cluster string, reason string, description string, configuration []byte) error {
	query := "username=" + url.QueryEscape(username) + "&reason=" + url.QueryEscape(reason) + "&description=" + url.QueryEscape(description)
	url := api.controllerURL + APIPrefix + "client/cluster/" + url.PathEscape(cluster) + "/configuration/create?" + query
	err := performWriteRequest(url, "POST", bytes.NewReader(configuration))
	return err
}

// AddTrigger access the REST API endpoint to add/register new trigger
func (api RestAPI) AddTrigger(username string, clusterName string, reason string, link string) error {
	query := "username=" + url.QueryEscape(username) + "&reason=" + url.QueryEscape(reason) + "&link=" + url.QueryEscape(link)
	url := api.controllerURL + APIPrefix + "client/cluster/" + url.PathEscape(clusterName) + "/trigger/must-gather?" + query
	err := performWriteRequest(url, "POST", nil)
	return err
}

// DeleteTrigger access the REST API endpoint to delete the selected trigger
func (api RestAPI) DeleteTrigger(triggerID string) error {
	url := api.controllerURL + APIPrefix + "client/trigger/" + triggerID
	err := performWriteRequest(url, "DELETE", nil)
	return err
}

// ActivateTrigger access the REST API endpoint to activate the selected trigger
func (api RestAPI) ActivateTrigger(triggerID string) error {
	url := api.controllerURL + APIPrefix + "client/trigger/" + triggerID + "/activate"
	err := performWriteRequest(url, "PUT", nil)
	return err
}

// DeactivateTrigger access the REST API endpoint to deactivate the selected trigger
func (api RestAPI) DeactivateTrigger(triggerID string) error {
	url := api.controllerURL + APIPrefix + "client/trigger/" + triggerID + "/deactivate"
	err := performWriteRequest(url, "PUT", nil)
	return err
}
