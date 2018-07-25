// using this as a guide to show me how to contact api, etc https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//ralph blah
type ralph struct {
	Accessories           string `json:"accessories"`
	Assetholders          string `json:"assetholders"`
	Assetmodels           string `json:"assetmodels"`
	BackOfficeAssets      string `json:"back-office-assets"`
	BaseObjectClusters    string `json:"base-object-clusters"`
	BaseObjects           string `json:"base-objects"`
	BaseObjectsLicences   string `json:"base-objects-licences"`
	BaseObjectsSupports   string `json:"base-objects-supports"`
	BudgetInfo            string `json:"budget-info"`
	BusinessSegments      string `json:"business-segments"`
	Categories            string `json:"categories"`
	CloudFlavors          string `json:"cloud-flavors"`
	CloudHosts            string `json:"cloud-hosts"`
	CloudProjects         string `json:"cloud-projects"`
	CloudProviders        string `json:"cloud-providers"`
	ClusterTypes          string `json:"cluster-types"`
	Clusters              string `json:"clusters"`
	ConfigurationClasses  string `json:"configuration-classes"`
	ConfigurationModules  string `json:"configuration-modules"`
	CustomFields          string `json:"custom-fields"`
	DataCenterAssets      string `json:"data-center-assets"`
	DataCenters           string `json:"data-centers"`
	Databases             string `json:"databases"`
	DcHosts               string `json:"dc-hosts"`
	Disks                 string `json:"disks"`
	DNSServers            string `json:"dns-servers"`
	Domains               string `json:"domains"`
	Environments          string `json:"environments"`
	Ethernets             string `json:"ethernets"`
	FibreChannelCards     string `json:"fibre-channel-cards"`
	Groups                string `json:"groups"`
	Ipaddresses           string `json:"ipaddresses"`
	LicenceTypes          string `json:"licence-types"`
	LicenceUsers          string `json:"licence-users"`
	Licences              string `json:"licences"`
	Manufacturers         string `json:"manufacturers"`
	Memory                string `json:"memory"`
	NetworkEnvironments   string `json:"network-environments"`
	NetworkKinds          string `json:"network-kinds"`
	Networks              string `json:"networks"`
	OfficeInfrastructures string `json:"office-infrastructures"`
	Processors            string `json:"processors"`
	ProfitCenters         string `json:"profit-centers"`
	RackAccessories       string `json:"rack-accessories"`
	Racks                 string `json:"racks"`
	Regions               string `json:"regions"`
	SecurityScans         string `json:"security-scans"`
	ServerRooms           string `json:"server-rooms"`
	Services              string `json:"services"`
	ServicesEnvironments  string `json:"services-environments"`
	Software              string `json:"software"`
	SupportTypes          string `json:"support-types"`
	Supports              string `json:"supports"`
	Teams                 string `json:"teams"`
	Transitions           string `json:"transitions"`
	TransitionsAction     string `json:"transitions-action"`
	TransitionsJob        string `json:"transitions-job"`
	TransitionsModel      string `json:"transitions-model"`
	Users                 string `json:"users"`
	Vips                  string `json:"vips"`
	VirtualServerTypes    string `json:"virtual-server-types"`
	VirtualServers        string `json:"virtual-servers"`
	Vulnerabilities       string `json:"vulnerabilities"`
	Warehouses            string `json:"warehouses"`
}

func callralphAPI() {
	apikey := "YOUR_API_KEY"
	//queryescape so it can be placed inside the url
	apiURL := url.QueryEscape(apikey)

	//api request url curl -X GET https://<YOUR-RALPH-URL>/api/ -H 'Authorization: Token <YOUR-TOKEN>'

	url := fmt.Sprintf("YOUR_RALPH_URL", apiURL)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record ralph

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

}
