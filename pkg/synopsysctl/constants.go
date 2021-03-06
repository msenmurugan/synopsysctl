/*
Copyright (C) 2019 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package synopsysctl

import "fmt"

// DefaultMetricsImage is the Metrics image deployed with Synopsys Operator by default
const DefaultMetricsImage string = "docker.io/prom/prometheus:v2.1.0"

// DefaultOperatorNamespace is the default namespace of Synopsys Operator
const DefaultOperatorNamespace string = "synopsys-operator"

// Default Base Specs for Create
const defaultBaseAlertSpec string = "default"
const defaultBaseBlackDuckSpec string = "persistentStorageLatest"
const defaultBaseOpsSightSpec string = "default"

// busybox image
const defaultBusyBoxImage string = "docker.io/busybox:1.28"

// flag for all namespaces
const allNamespacesFlag string = "--all-namespaces"

// Flags for using native mode - doesn't deploy
var nativeFormat = "json"

const (
	clusterTypeKubernetes = "KUBERNETES"
	clusterTypeOpenshift  = "OPENSHIFT"
)

var nativeClusterType = clusterTypeKubernetes

var baseURL = "https://raw.githubusercontent.com/blackducksoftware/releases/master"

var baseChartRepository = "https://chartmuseum.cloudnative.sig-clops.synopsys.com"

// Alert

// AlertPostSuffix adds "-alert" to the end of the release
const AlertPostSuffix = "-alert"

var alertChartRepository = fmt.Sprintf("%s/charts/alert-helmchart-5.3.0.tgz", baseChartRepository)

// Black Duck
var blackduckChartRepository = fmt.Sprintf("%s/charts/blackduck-2020.4.0.tgz", baseChartRepository)

// Polaris
var polarisName = "polaris"
var polarisChartRepository = fmt.Sprintf("%s/charts/polaris-helmchart-2020.03.tgz", baseChartRepository)

// Polaris Reporting Constants
var polarisReportingName = "polaris-reporting"
var polarisReportingChartRepository = fmt.Sprintf("%s/charts/polaris-helmchart-reporting-2020.03.tgz", baseChartRepository)

// BDBA Constants
var bdbaName = "bdba"
var bdbaVersion = "2020.03"
var bdbaChartRepository = fmt.Sprintf("%s/charts/bdba-%s.tgz", baseChartRepository, bdbaVersion)
