// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vmware-tanzu/cloud-native-security-inspector/cnsi-scanner-trivy/pkg/harbor"
	"io"
	"net/http"
)

func main() {
	scan("")
	//getScanReport("467c76f4d47d91419d052425")
	//getScanReport("fa42a6100635ef789da5041f")
}

func scan(image string) (string, error) {
	url := "http://127.0.0.1:30003/api/v1/scan"
	contentType := "application/vnd.scanner.adapter.scan.request+json; version=1.0"
	postData := []byte(`{
	  "registry": {
		"url": "http://10.212.47.157",
		"authorization": "Bearer JWTTOKENGOESHERE"
	  },
	  "artifact": {
		"repository": "library/test-misconfig",
		"digest": "sha256:1dc4e5d3769b613ca32afb6457efe09cd0057e8949ec1e3e21de45c9221bc221"
	  }
    }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusAccepted {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		type scanResponse struct {
			Id string `json:"id"`
		}
		var scanResp scanResponse
		if err = json.Unmarshal(bodyBytes, &scanResp); err != nil {
			panic(err)
		}
		fmt.Println(scanResp.Id)
	}

	fmt.Println(resp.Status)
}

func getScanReport(id string) harbor.ScanReport {
	url := "http://127.0.0.1:30003/api/v1/scan/" + id + "/report"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var scanReport harbor.ScanReport
	err = json.Unmarshal(body, &scanReport)
	if err != nil {
		panic(err)
	}
	fmt.Println(scanReport.Report)
	fmt.Println(scanReport.Report.Results)
	return scanReport
}
