/*
 * Copyright (C) 2020 Synopsys, Inc.
 *
 *  Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 *  with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 *  under the License.
 */

package bdba

import (
	"fmt"

	"github.com/blackducksoftware/synopsysctl/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// HelmValuesFromCobraFlags is a type for converting synopsysctl flags
// to Helm Chart fields and values
// args: map of helm chart field to value
type HelmValuesFromCobraFlags struct {
	args     map[string]interface{}
	flagTree FlagTree
}

// FlagTree is a set of fields needed to configure the BDBA Helm Chart
type FlagTree struct {
	Version string `json:"version"`

	ClusterDomain string `json:"clusterDomain"`

	// Storage
	PGStorageClass        string `json:"pgStorageClass"`
	PGSize                string `json:"pgSize"`
	PGExistingClaim       string `json:"pgExistingClaim"`
	MinioStorageClass     string `json:"minioStorageClass"`
	MinioSize             string `json:"minioSize"`
	MinioExistingClaim    string `json:"minioExistingClaim"`
	RabbitMQStorageClass  string `json:"rabbitmqStorageClass"`
	RabbitMQSize          string `json:"rabbitmqSize"`
	RabbitMQExistingClaim string `json:"rabbitmqExistingClaim"`

	// Licensing
	LicensingUsername string `json:"licensingUsername"`
	LicensingPassword string `json:"licensingPassword"`
	LicensingUpstream string `json:"licensingUpstream"`

	// Rabbitmq configuration

	// Web frontend configuration
	SessionCookieAge int    `json:"sessionCookieAge"`
	FrontendReplicas int    `json:"frontendReplicas"`
	HideLicenses     bool   `json:"hideLicenses"`
	OfflineMode      bool   `json:"offlineMode"`
	AdminEmail       string `json:"adminEmail"`
	ErrorAdminEmail  string `json:"errorAdminEmail"`
	RootURL          string `json:"rootURL"`

	// SMTP configuration
	EmailEnabled      bool   `json:"emailEnabled"`
	EmailSMTPHost     string `json:"emailSMTPHost"`
	EmailSMTPPort     int    `json:"emailSMTPPort"`
	EmailSMTPUser     string `json:"emailSMTPUser"`
	EmailSMTPPassword string `json:"emailSMTPPassword"`
	EmailFrom         string `json:"emailFrom"`
	EmailSecurity     string `json:"emailSecurity"`
	EmailVerify       bool   `json:"emailVerify"`

	// LDAP Authentication
	LDAPEnabled              bool   `json:"ldapEnabled"`
	LDAPServerURI            string `json:"ldapServerURI"`
	LDAPUserDNTemplate       string `json:"ldapUserDNTemplate"`
	LDAPBindAsAuthenticating bool   `json:"ldapBindAsAuthenticating"`
	LDAPBindDN               string `json:"ldapBindDN"`
	LDAPBindPassword         string `json:"ldapBindPassword"`
	LDAPStartTLS             bool   `json:"ldapStartTLS"`
	LDAPVerify               bool   `json:"ldapVerify"`
	LDAPRootCASecret         string `json:"ldapRootCASecret"`
	LDAPRequireGroup         string `json:"ldapRequireGroup"`
	LDAPUserSearch           string `json:"ldapUserSearch"`
	LDAPUserSearchScope      string `json:"ldapUserSearchScope"`
	LDAPGroupSearch          string `json:"ldapGroupSearch"`
	LDAPGroupSearchScope     string `json:"ldapGroupSearchScope"`
	LDAPNestedSearch         bool   `json:"ldapNestedSearch"`

	// Logging
	DisableFrontendLogging bool `json:"disableFrontendLogging"`
	DisableWorkerLogging   bool `json:"disableWorkerLogging"`

	// Worker scaling
	WorkerReplicas    int `json:"workerReplicas"`
	WorkerConcurrency int `json:"workerConcurrency"` // TODO: Patcher

	// Networking and security
	RootCASecret string `json:"rootCASecret"`
	HTTPProxy    string `json:"httpProxy"`
	HTTPNoProxy  string `json:"httpNoProxy"`

	// Minio
	MinioMode string `json:"minioMode"`

	// Ingress
	IngressEnabled       bool   `json:"ingressEnabled"`
	IngressHost          string `json:"ingressHost"`
	IngressTLSEnabled    bool   `json:"ingressTLSEnabled"`
	IngressTLSSecretName string `json:"ingressTLSSecretName"`

	// External PostgreSQL
	ExternalPG             bool   `json:ExternalPg`
	ExternalPGHost         string `json:"ExternalPgHost"`
	ExternalPGPort         int    `json:"ExternalPgPort"`
	ExternalPGUser         string `json:"ExternalPgUser"`
	ExternalPGPassword     string `json:"ExternalPgUser"`
	ExternalPGDataBase     string `json:"ExternalPgDataBase"`
	ExternalPGSSLMode      string `json:"ExternalPgSSLMode"`
	ExternalPGRootCASecret string `json:"ExternalPgRootCASecret"`
	ExternalPGClientSecret string `json:"ExternalPgClientSecret"`

	// Secrets
	PGPassword       string `json:"pgPassword"`
	PGExistingSecret string `json:"pgExistingSecret"`
}

// NewHelmValuesFromCobraFlags returns an initialized HelmValuesFromCobraFlags
func NewHelmValuesFromCobraFlags() *HelmValuesFromCobraFlags {
	return &HelmValuesFromCobraFlags{
		args:     make(map[string]interface{}, 0),
		flagTree: FlagTree{},
	}
}

// GetArgs returns the map of helm chart fields to values
func (ctl *HelmValuesFromCobraFlags) GetArgs() map[string]interface{} {
	return ctl.args
}

// SetArgs set the map to values
func (ctl *HelmValuesFromCobraFlags) SetArgs(args map[string]interface{}) {
	for key, value := range args {
		ctl.args[key] = value
	}
}

// AddCobraFlagsToCommand adds flags for the BDBA helm chart to the cmd
// master=true is used to add all flags for creating an instance
// master=false is used to add a subset of flags for updating an instance
func (ctl *HelmValuesFromCobraFlags) AddCobraFlagsToCommand(cmd *cobra.Command, master bool) {
	// [DEV NOTE:] please organize flags in order of importance and group related flags

	cmd.Flags().StringVar(&ctl.flagTree.Version, "version", ctl.flagTree.Version, "Version of BDBA you want to install [Example: 2020.03]")
	// if master {
	// 	cobra.MarkFlagRequired(cmd.Flags(), "version")
	// }

	cmd.Flags().StringVar(&ctl.flagTree.ClusterDomain, "cluster-domain", "cluster.local", "Kubernetes cluster domain")

	// Storage
	cmd.Flags().StringVar(&ctl.flagTree.PGStorageClass, "postgres-storage-class", "default", "Storage class for PostgreSQL")
	cmd.Flags().StringVar(&ctl.flagTree.PGSize, "postgres-size", "300Gi", "Persistent volument claim size for PostgreSQL")
	cmd.Flags().StringVar(&ctl.flagTree.PGExistingClaim, "postgres-existing-claim", ctl.flagTree.PGExistingClaim, "Existing claim to use for PostgreSQL")
	cmd.Flags().StringVar(&ctl.flagTree.MinioStorageClass, "minio-storage-class", "default", "Storage class for minio")
	cmd.Flags().StringVar(&ctl.flagTree.MinioSize, "minio-size", "300Gi", "Persistent volume claim size of minio")
	cmd.Flags().StringVar(&ctl.flagTree.MinioExistingClaim, "minio-existing-claim", ctl.flagTree.MinioExistingClaim, "Existing claim to use for minio")
	cmd.Flags().StringVar(&ctl.flagTree.RabbitMQStorageClass, "rabbitmq-storage-class", "default", "Storage class for RabbitMQ")
	cmd.Flags().StringVar(&ctl.flagTree.RabbitMQSize, "rabbitmq-size", "8Gi", "Persistent volument claim size for RabbitMQ")
	cmd.Flags().StringVar(&ctl.flagTree.RabbitMQExistingClaim, "rabbitmq-existing-claim", ctl.flagTree.RabbitMQExistingClaim, "Existing claim to use for RabbitMQ")

	// Licensing
	cmd.Flags().StringVar(&ctl.flagTree.LicensingUsername, "license-username", ctl.flagTree.LicensingUsername, "Username for licensing")
	cmd.Flags().StringVar(&ctl.flagTree.LicensingPassword, "license-password", ctl.flagTree.LicensingPassword, "Username for password")
	cmd.Flags().StringVar(&ctl.flagTree.LicensingUpstream, "license-upstream", "https://protecode-sc.com/", "Upstream server for data updates")

	// Web frontend configuration
	cmd.Flags().IntVar(&ctl.flagTree.SessionCookieAge, "session-cookie-age", 1209600, "Session cookie expiration time")
	cmd.Flags().IntVar(&ctl.flagTree.FrontendReplicas, "frontend-replicas", 1, "Number of web application replicas")
	cmd.Flags().BoolVar(&ctl.flagTree.HideLicenses, "hide-licenses", false, "Hide licensing information")
	cmd.Flags().BoolVar(&ctl.flagTree.OfflineMode, "enable-offline-mode", false, "Run in airgapped mode")
	cmd.Flags().StringVar(&ctl.flagTree.AdminEmail, "admin-email", "admin@bdba.local", "Administrator email address")
	cmd.Flags().StringVar(&ctl.flagTree.RootURL, "root-url", "http://bdba.local", "Root URL of application to be used in tempates")

	// SMTP configuration
	cmd.Flags().BoolVar(&ctl.flagTree.EmailEnabled, "enable-email", false, "Enable STMP to send emails")
	cmd.Flags().StringVar(&ctl.flagTree.EmailSMTPHost, "email-smtp-host", ctl.flagTree.EmailSMTPHost, "SMTP server address")
	cmd.Flags().IntVar(&ctl.flagTree.EmailSMTPPort, "email-smtp-port", 25, "SMTP server port")
	cmd.Flags().StringVar(&ctl.flagTree.EmailSMTPUser, "email-smtp-user", ctl.flagTree.EmailSMTPUser, "SMTP user")
	cmd.Flags().StringVar(&ctl.flagTree.EmailSMTPPassword, "email-smtp-password", ctl.flagTree.EmailSMTPPassword, "SMTP password")
	cmd.Flags().StringVar(&ctl.flagTree.EmailFrom, "email-from", "noreply@bdba.local", "Email sender address")
	cmd.Flags().StringVar(&ctl.flagTree.EmailSecurity, "email-security", "none", "Email security mode [none|ssl|starttls]")
	cmd.Flags().BoolVar(&ctl.flagTree.EmailVerify, "verify-email", false, "Verify SMTP server certificate")

	// LDAP Authentication
	cmd.Flags().BoolVar(&ctl.flagTree.LDAPEnabled, "enable-ldap", false, "Enable LDAP for authentication")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPServerURI, "ldap-server-uri", ctl.flagTree.LDAPServerURI, "LDAP server URI")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPUserDNTemplate, "ldap-user-dn-template", ctl.flagTree.LDAPUserDNTemplate, "LDAP user DN template")
	cmd.Flags().BoolVar(&ctl.flagTree.LDAPBindAsAuthenticating, "enable-ldap-bind-as-authenticating", false, "Bind as authenticating user")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPBindDN, "ldap-bind-dn", ctl.flagTree.LDAPBindDN, "Generic LDAP bind username (optional)")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPBindPassword, "ldap-bind-password", ctl.flagTree.LDAPBindPassword, "Generic LDAP bind password (optional)")
	cmd.Flags().BoolVar(&ctl.flagTree.LDAPStartTLS, "ldap-start-tls", false, "Enable start TLS for LDAP")
	cmd.Flags().BoolVar(&ctl.flagTree.LDAPVerify, "verify-ldap", false, "Verify LDAP server certificate")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPRootCASecret, "ldap-root-ca-secret", ctl.flagTree.LDAPRootCASecret, "Secret to use for LDAP root certificate")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPRequireGroup, "ldap-require-group", ctl.flagTree.LDAPRequireGroup, "LDAP group required to allow login")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPUserSearch, "ldap-user-search", ctl.flagTree.LDAPUserSearch, "Base DN for user branch")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPUserSearchScope, "ldap-user-search-scope", ctl.flagTree.LDAPUserSearchScope, "LDAP search filter for users")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPGroupSearch, "ldap-group-search", ctl.flagTree.LDAPGroupSearch, "Base DN for groups branch")
	cmd.Flags().StringVar(&ctl.flagTree.LDAPGroupSearchScope, "ldap-group-search-scope", ctl.flagTree.LDAPGroupSearchScope, "LDAP search filter for groups")
	cmd.Flags().BoolVar(&ctl.flagTree.LDAPNestedSearch, "enable-ldap-nested-search", false, "Enable nested search")

	// Logging
	cmd.Flags().BoolVar(&ctl.flagTree.DisableFrontendLogging, "disable-frontend-logging", false, "Disable log collection in web application pods")
	cmd.Flags().BoolVar(&ctl.flagTree.DisableWorkerLogging, "disable-worker-logging", false, "Disable log collection in scanner pods")

	// Worker scaling
	cmd.Flags().IntVar(&ctl.flagTree.WorkerReplicas, "worker-replicas", 1, "Number of worker replicas")
	cmd.Flags().IntVar(&ctl.flagTree.WorkerConcurrency, "worker-concurrency", 1, "Amount of concurrent workers per pod")

	// Minio
	cmd.Flags().StringVar(&ctl.flagTree.MinioMode, "minio-mode", "standalone", "Minio mode [standalone|distributed]")

	// Networking and security
	cmd.Flags().StringVar(&ctl.flagTree.RootCASecret, "root-ca-secret", ctl.flagTree.RootCASecret, "Additional root certificate")
	cmd.Flags().StringVar(&ctl.flagTree.HTTPProxy, "http-proxy", ctl.flagTree.HTTPProxy, "HTTP Proxy to use")
	cmd.Flags().StringVar(&ctl.flagTree.HTTPNoProxy, "http-no-proxy", ctl.flagTree.HTTPNoProxy, "Comma-separated list of domain extensions to omit proxy")

	// Ingress
	cmd.Flags().BoolVar(&ctl.flagTree.IngressEnabled, "enable-ingress", false, "Enable ingress")
	cmd.Flags().StringVar(&ctl.flagTree.IngressHost, "ingress-host", ctl.flagTree.IngressHost, "Hostname for ingress")
	cmd.Flags().BoolVar(&ctl.flagTree.IngressTLSEnabled, "enable-ingress-tls", false, "Enable TLS for ingress")
	cmd.Flags().StringVar(&ctl.flagTree.IngressTLSSecretName, "ingress-tls-secret", ctl.flagTree.IngressTLSSecretName, "TLS Secret to use for ingress")

	// External PG
	cmd.Flags().BoolVar(&ctl.flagTree.ExternalPG,
		"external-postgres", false, "Use external PostgreSQL")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGHost,
		"external-postgres-host", ctl.flagTree.ExternalPGHost, "Hostname for external postgresql database")
	cmd.Flags().IntVar(&ctl.flagTree.ExternalPGPort,
		"external-postgres-port", 5432, "Port for external PostgreSQL database")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGDataBase,
		"external-postgres-database", ctl.flagTree.ExternalPGDataBase, "Database for external PostgreSQL database")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGUser,
		"external-postgres-user", ctl.flagTree.ExternalPGUser, "User for external PostgreSQL database")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGPassword,
		"external-postgres-password", ctl.flagTree.ExternalPGPassword, "Password for external PostgreSQL database")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGSSLMode,
		"external-postgres-ssl-mode", "disable", "PostgreSQL SSL mode [disable|allow|prefer|require|verify-ca|verify-full]")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGClientSecret,
		"external-postgres-client-secret", ctl.flagTree.ExternalPGClientSecret, "Secret name for external PostgreSQL client certificate (TLS Secret)")
	cmd.Flags().StringVar(&ctl.flagTree.ExternalPGRootCASecret,
		"external-postgres-rootca-secret", ctl.flagTree.ExternalPGRootCASecret, "Secret name for external PostgreSQL root certificate")

	// Secrets
	cmd.Flags().StringVar(&ctl.flagTree.PGPassword, "postgres-password", "default", "PostgreSQL password")
	cmd.Flags().StringVar(&ctl.flagTree.PGExistingSecret, "postgres-secret", ctl.flagTree.PGExistingSecret, "Existing secret for PostgreSQL")

	cmd.Flags().SortFlags = false
}

// CheckValuesFromFlags returns an error if a value set by a flag is invalid
func (ctl *HelmValuesFromCobraFlags) CheckValuesFromFlags(flagset *pflag.FlagSet) error {
	if flagset.Lookup("enable-email").Value.String() == "true" {
		if !flagset.Lookup("email-smtp-host").Changed {
			return fmt.Errorf("--email-smtp-host my be set for email")
		}
	}

	if flagset.Lookup("enable-offline-mode").Value.String() == "false" {
		if !flagset.Lookup("license-username").Changed {
			return fmt.Errorf("--license-username must be set if not in offline mode")
		}

		if !flagset.Lookup("license-password").Changed {
			return fmt.Errorf("--license-password must be set if not in offline mode")
		}
	}

	if flagset.Lookup("enable-ldap").Value.String() == "true" {
		if !flagset.Lookup("ldap-server-uri").Changed {
			return fmt.Errorf("--ldap-server-uri must be set for LDAP")
		}

		if !flagset.Lookup("ldap-user-dn-template").Changed {
			return fmt.Errorf("--ldap-user-dn-template must be set for LDAP")
		}

		if flagset.Lookup("ldap-require-group").Changed {
			if !flagset.Lookup("ldap-user-search").Changed {
				return fmt.Errorf("--ldap-user-search must be set for --ldap-require-group")
			}

			if !flagset.Lookup("ldap-user-search-scope").Changed {
				return fmt.Errorf("--ldap-user-search-scope must be set for --ldap-require-group")
			}

			if !flagset.Lookup("ldap-group-search").Changed {
				return fmt.Errorf("--ldap-group-search must be set for --ldap-require-group")
			}

			if !flagset.Lookup("ldap-group-search-scope").Changed {
				return fmt.Errorf("--ldap-group-search-scope must be set for --ldap-require-group")
			}
		}

		if flagset.Lookup("ldap-bind-dn").Changed {
			if !flagset.Lookup("ldap-bind-password").Changed {
				return fmt.Errorf("--ldap-bind-password must be set for --ldap-bind-dn")
			}
		}

		if flagset.Lookup("enable-ingress").Value.String() == "true" {
			if !flagset.Lookup("ingress-host").Changed {
				return fmt.Errorf("--ingress-host must be set for ingress")
			}

			if flagset.Lookup("enable-ingress-tls").Value.String() == "true" {
				if !flagset.Lookup("ingress-tls-secret").Changed {
					return fmt.Errorf("--ingress-tls-secret must be set for TLS-enabled ingress")
				}
			}
		}

		if flagset.Lookup("external-postgres").Value.String() == "false" {
			if !flagset.Lookup("postgres-password").Changed && !flagset.Lookup("postgres-secret").Changed {
				return fmt.Errorf("--postgres-password or --postgres-secret must be provided for internal postgres")
			}

			if flagset.Lookup("postgres-password").Changed && flagset.Lookup("postgres-secret").Changed {
				return fmt.Errorf("--postgres-password and --postgres-secret are mutually exclusive")
			}

		}

		if flagset.Lookup("external-postgres").Value.String() == "true" {
			if !flagset.Lookup("external-postgres-host").Changed {
				return fmt.Errorf("--external-postgres-host must be provided for external postgresql")
			}

			if !flagset.Lookup("external-postgres-database").Changed {
				return fmt.Errorf("--external-postgres-database must be provided for external postgresql")
			}

			if !flagset.Lookup("external-postgres-password").Changed && !flagset.Lookup("external-postgres-client-secret").Changed {
				return fmt.Errorf("--external-postgres-password or --external-postgres-client-secret must be provided for external postgres")
			}

		}

	}

	return nil
}

// GenerateHelmFlagsFromCobraFlags checks each flag in synopsysctl and updates the map to
// contain the corresponding helm chart field and value
func (ctl *HelmValuesFromCobraFlags) GenerateHelmFlagsFromCobraFlags(flagset *pflag.FlagSet) (map[string]interface{}, error) {
	err := ctl.CheckValuesFromFlags(flagset)
	if err != nil {
		return nil, err
	}
	flagset.VisitAll(ctl.AddHelmValueByCobraFlag)

	return ctl.args, nil
}

// AddHelmValueByCobraFlag adds the helm chart field and value based on the flag set
// in synopsysctl
func (ctl *HelmValuesFromCobraFlags) AddHelmValueByCobraFlag(f *pflag.Flag) {
	if f.Changed {
		log.Debugf("flag '%s': CHANGED", f.Name)
		switch f.Name {
		case "cluster-domain":
			util.SetHelmValueInMap(ctl.args, []string{"rabbitmq", "rabbitmq", "clustering", "k8s_domain"}, ctl.flagTree.ClusterDomain)
			util.SetHelmValueInMap(ctl.args, []string{"minio", "clusterDomain"}, ctl.flagTree.ClusterDomain)
		case "postgres-storage-class":
			util.SetHelmValueInMap(ctl.args, []string{"postgresql", "persistence", "storageClass"}, ctl.flagTree.PGStorageClass)
		case "postgres-size":
			util.SetHelmValueInMap(ctl.args, []string{"postgresql", "persistence", "size"}, ctl.flagTree.PGSize)
		case "postgres-existing-claim":
			util.SetHelmValueInMap(ctl.args, []string{"postgresql", "persistence", "existingClaim"}, ctl.flagTree.PGExistingClaim)
		case "minio-storage-class":
			util.SetHelmValueInMap(ctl.args, []string{"minio", "persistence", "storageClass"}, ctl.flagTree.MinioStorageClass)
		case "minio-size":
			util.SetHelmValueInMap(ctl.args, []string{"minio", "persistence", "size"}, ctl.flagTree.MinioSize)
		case "minio-existing-claim":
			util.SetHelmValueInMap(ctl.args, []string{"minio", "persistence", "existingClaim"}, ctl.flagTree.MinioExistingClaim)
		case "rabbitmq-storage-class":
			util.SetHelmValueInMap(ctl.args, []string{"rabbitmq", "persistence", "storageClass"}, ctl.flagTree.RabbitMQStorageClass)
		case "rabbitmq-size":
			util.SetHelmValueInMap(ctl.args, []string{"rabbitmq", "persistence", "size"}, ctl.flagTree.RabbitMQSize)
		case "rabbitmq-existing-claim":
			util.SetHelmValueInMap(ctl.args, []string{"rabbitmq", "persistence", "existingClaim"}, ctl.flagTree.RabbitMQExistingClaim)
		case "license-username":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "licensing", "username"}, ctl.flagTree.LicensingUsername)
		case "license-password":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "licensing", "password"}, ctl.flagTree.LicensingPassword)
		case "license-upstream":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "licensing", "upstream"}, ctl.flagTree.LicensingUpstream)
		case "session-cookie-age":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "web", "sessionCookieAge"}, ctl.flagTree.SessionCookieAge)
		case "frontend-replicas":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "web", "replicas"}, ctl.flagTree.FrontendReplicas)
		case "hide-licenses":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "web", "hideLicenses"}, ctl.flagTree.HideLicenses)
		case "enable-offline-mode":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "web", "offlineMode"}, ctl.flagTree.OfflineMode)
		case "admin-email":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "web", "admin"}, ctl.flagTree.AdminEmail)
		case "root-url":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "web", "rootURL"}, ctl.flagTree.RootURL)
		case "enable-email":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "enabled"}, ctl.flagTree.EmailEnabled)
		case "email-smtp-host":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "smtpHost"}, ctl.flagTree.EmailSMTPHost)
		case "email-smtp-port":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "smtpPort"}, ctl.flagTree.EmailSMTPPort)
		case "email-smtp-user":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "smtpUser"}, ctl.flagTree.EmailSMTPUser)
		case "email-smtp-password":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "smtpPassword"}, ctl.flagTree.EmailSMTPPassword)
		case "email-from":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "from"}, ctl.flagTree.EmailFrom)
		case "email-security":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "security"}, ctl.flagTree.EmailSecurity)
		case "verify-email":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "email", "verify"}, ctl.flagTree.EmailVerify)
		case "enable-ldap":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "enabled"}, ctl.flagTree.LDAPEnabled)
		case "ldap-server-uri":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "serverUri"}, ctl.flagTree.LDAPServerURI)
		case "ldap-user-dn-template":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "userDNTemplate"}, ctl.flagTree.LDAPUserDNTemplate)
		case "enable-ldap-bind-as-authenticating":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "bindAsAuthenticating"}, ctl.flagTree.LDAPBindAsAuthenticating)
		case "ldap-bind-dn":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "bindDN"}, ctl.flagTree.LDAPBindDN)
		case "ldap-bind-password":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "bindPassword"}, ctl.flagTree.LDAPBindPassword)
		case "ldap-start-tls":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "startTLS"}, ctl.flagTree.LDAPStartTLS)
		case "verify-ldap":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "verify"}, ctl.flagTree.LDAPVerify)
		case "ldap-root-ca-secret":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "rootCASecret"}, ctl.flagTree.LDAPRootCASecret)
		case "ldap-require-group":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "requireGroup"}, ctl.flagTree.LDAPRequireGroup)
		case "ldap-user-search":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "userSearch"}, ctl.flagTree.LDAPUserSearch)
		case "ldap-user-search-scope":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "userSearchScope"}, ctl.flagTree.LDAPUserSearchScope)
		case "ldap-group-search":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "groupSearch"}, ctl.flagTree.LDAPGroupSearch)
		case "ldap-group-search-scope":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "groupSearchScope"}, ctl.flagTree.LDAPGroupSearchScope)
		case "enable-ldap-nested-search":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "ldap", "nestedSearch"}, ctl.flagTree.LDAPNestedSearch)
		case "disable-frontend-logging":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "applicationLogging", "enabled"}, !ctl.flagTree.DisableFrontendLogging)
		case "disable-worker-logging":
			util.SetHelmValueInMap(ctl.args, []string{"worker", "applicationLogging", "enabled"}, !ctl.flagTree.DisableWorkerLogging)
		case "worker-replicas":
			util.SetHelmValueInMap(ctl.args, []string{"worker", "replicas"}, ctl.flagTree.WorkerReplicas)
		case "worker-concurrency":
			util.SetHelmValueInMap(ctl.args, []string{"worker", "concurrency"}, ctl.flagTree.WorkerConcurrency)
		case "minio-mode":
			util.SetHelmValueInMap(ctl.args, []string{"minio", "mode"}, ctl.flagTree.MinioMode)
		case "root-ca-secret":
			util.SetHelmValueInMap(ctl.args, []string{"rootCASecret"}, ctl.flagTree.RootCASecret)
		case "http-proxy":
			util.SetHelmValueInMap(ctl.args, []string{"httpProxy"}, ctl.flagTree.HTTPProxy)
		case "http-no-proxy":
			util.SetHelmValueInMap(ctl.args, []string{"httpNoProxy"}, ctl.flagTree.HTTPNoProxy)
		case "enable-ingress":
			util.SetHelmValueInMap(ctl.args, []string{"ingress", "enabled"}, ctl.flagTree.IngressEnabled)
		case "ingress-host":
			util.SetHelmValueInMap(ctl.args, []string{"ingress", "host"}, ctl.flagTree.IngressHost)
		case "enable-ingress-tls":
			util.SetHelmValueInMap(ctl.args, []string{"ingress", "tls", "enabled"}, ctl.flagTree.IngressTLSEnabled)
		case "ingress-tls-secret":
			util.SetHelmValueInMap(ctl.args, []string{"ingress", "tls", "secretName"}, ctl.flagTree.IngressTLSSecretName)
		case "external-postgres":
			util.SetHelmValueInMap(ctl.args, []string{"postgresql", "enabled"}, !ctl.flagTree.ExternalPG)
		case "external-postgres-host":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "postgresqlHost"}, ctl.flagTree.ExternalPGHost)
		case "external-postgres-port":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "postgresqlPort"}, ctl.flagTree.ExternalPGPort)
		case "external-postgres-database":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "postgresqlDatabase"}, ctl.flagTree.ExternalPGDataBase)
		case "external-postgres-user":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "postgresqlUsername"}, ctl.flagTree.ExternalPGUser)
		case "external-postgres-password":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "postgresqlPassword"}, ctl.flagTree.ExternalPGPassword)
		case "external-postgres-ssl-mode":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "postgresqlSslMode"}, ctl.flagTree.ExternalPGSSLMode)
		case "external-postgres-client-secret":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "clientSecretName"}, ctl.flagTree.ExternalPGClientSecret)
		case "external-postgres-rootca-secret":
			util.SetHelmValueInMap(ctl.args, []string{"frontend", "database", "rootCASecretName"}, ctl.flagTree.ExternalPGRootCASecret)
		case "postgres-password":
			util.SetHelmValueInMap(ctl.args, []string{"postgresql", "postgresqlPassword"}, ctl.flagTree.PGPassword)
		case "postgres-secret":
			util.SetHelmValueInMap(ctl.args, []string{"global", "postgresql", "existingSecret"}, ctl.flagTree.PGExistingSecret)

		default:
			log.Debugf("flag '%s': NOT FOUND", f.Name)
		}
	} else {
		log.Debugf("flag '%s': UNCHANGED", f.Name)
	}
}
