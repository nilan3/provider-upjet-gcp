// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/upbound/provider-gcp/config/accessapproval"
	"github.com/upbound/provider-gcp/config/accesscontextmanager"
	"github.com/upbound/provider-gcp/config/apigee"
	"github.com/upbound/provider-gcp/config/beyondcorp"
	"github.com/upbound/provider-gcp/config/bigquery"
	"github.com/upbound/provider-gcp/config/bigtable"
	composer "github.com/upbound/provider-gcp/config/cloudcomposer"
	"github.com/upbound/provider-gcp/config/cloudfunctions"
	"github.com/upbound/provider-gcp/config/cloudplatform"
	"github.com/upbound/provider-gcp/config/cloudrun"
	"github.com/upbound/provider-gcp/config/cloudscheduler"
	"github.com/upbound/provider-gcp/config/cloudtasks"
	"github.com/upbound/provider-gcp/config/compute"
	"github.com/upbound/provider-gcp/config/container"
	"github.com/upbound/provider-gcp/config/containerattached"
	"github.com/upbound/provider-gcp/config/containeraws"
	"github.com/upbound/provider-gcp/config/containerazure"
	"github.com/upbound/provider-gcp/config/dataflow"
	"github.com/upbound/provider-gcp/config/dataplex"
	"github.com/upbound/provider-gcp/config/dataproc"
	"github.com/upbound/provider-gcp/config/dns"
	"github.com/upbound/provider-gcp/config/endpoints"
	"github.com/upbound/provider-gcp/config/firebaserules"
	"github.com/upbound/provider-gcp/config/gameservices"
	"github.com/upbound/provider-gcp/config/gkehub"
	"github.com/upbound/provider-gcp/config/healthcare"
	"github.com/upbound/provider-gcp/config/iap"
	"github.com/upbound/provider-gcp/config/identityplatform"
	"github.com/upbound/provider-gcp/config/kms"
	"github.com/upbound/provider-gcp/config/logging"
	"github.com/upbound/provider-gcp/config/monitoring"
	"github.com/upbound/provider-gcp/config/notebooks"
	"github.com/upbound/provider-gcp/config/oslogin"
	"github.com/upbound/provider-gcp/config/privateca"
	"github.com/upbound/provider-gcp/config/project"
	"github.com/upbound/provider-gcp/config/pubsub"
	"github.com/upbound/provider-gcp/config/redis"
	"github.com/upbound/provider-gcp/config/secretmanager"
	"github.com/upbound/provider-gcp/config/servicenetworking"
	"github.com/upbound/provider-gcp/config/sourcerepo"
	"github.com/upbound/provider-gcp/config/spanner"
	"github.com/upbound/provider-gcp/config/sql"
	"github.com/upbound/provider-gcp/config/storage"
	"github.com/upbound/provider-gcp/config/tags"
	"github.com/upbound/provider-gcp/config/tpu"
	"github.com/upbound/provider-gcp/config/vertexai"
	"github.com/upbound/provider-gcp/config/vpcaccess"
)

func init() {
	ProviderConfiguration.AddConfig(accessapproval.Configure)
	ProviderConfiguration.AddConfig(accesscontextmanager.Configure)
	ProviderConfiguration.AddConfig(apigee.Configure)
	ProviderConfiguration.AddConfig(bigtable.Configure)
	ProviderConfiguration.AddConfig(composer.Configure)
	ProviderConfiguration.AddConfig(cloudfunctions.Configure)
	ProviderConfiguration.AddConfig(cloudplatform.Configure)
	ProviderConfiguration.AddConfig(cloudrun.Configure)
	ProviderConfiguration.AddConfig(cloudscheduler.Configure)
	ProviderConfiguration.AddConfig(cloudtasks.Configure)
	ProviderConfiguration.AddConfig(containerattached.Configure)
	ProviderConfiguration.AddConfig(containeraws.Configure)
	ProviderConfiguration.AddConfig(containerazure.Configure)
	ProviderConfiguration.AddConfig(compute.Configure)
	ProviderConfiguration.AddConfig(container.Configure)
	ProviderConfiguration.AddConfig(dataflow.Configure)
	ProviderConfiguration.AddConfig(dataplex.Configure)
	ProviderConfiguration.AddConfig(dataproc.Configure)
	ProviderConfiguration.AddConfig(dns.Configure)
	ProviderConfiguration.AddConfig(endpoints.Configure)
	ProviderConfiguration.AddConfig(firebaserules.Configure)
	ProviderConfiguration.AddConfig(gameservices.Configure)
	ProviderConfiguration.AddConfig(iap.Configure)
	ProviderConfiguration.AddConfig(identityplatform.Configure)
	ProviderConfiguration.AddConfig(logging.Configure)
	ProviderConfiguration.AddConfig(kms.Configure)
	ProviderConfiguration.AddConfig(notebooks.Configure)
	ProviderConfiguration.AddConfig(privateca.Configure)
	ProviderConfiguration.AddConfig(oslogin.Configure)
	ProviderConfiguration.AddConfig(project.Configure)
	ProviderConfiguration.AddConfig(pubsub.Configure)
	ProviderConfiguration.AddConfig(redis.Configure)
	ProviderConfiguration.AddConfig(secretmanager.Configure)
	ProviderConfiguration.AddConfig(servicenetworking.Configure)
	ProviderConfiguration.AddConfig(sourcerepo.Configure)
	ProviderConfiguration.AddConfig(spanner.Configure)
	ProviderConfiguration.AddConfig(storage.Configure)
	ProviderConfiguration.AddConfig(sql.Configure)
	ProviderConfiguration.AddConfig(redis.Configure)
	ProviderConfiguration.AddConfig(bigquery.Configure)
	ProviderConfiguration.AddConfig(beyondcorp.Configure)
	ProviderConfiguration.AddConfig(vertexai.Configure)
	ProviderConfiguration.AddConfig(tags.Configure)
	ProviderConfiguration.AddConfig(tpu.Configure)
	ProviderConfiguration.AddConfig(vpcaccess.Configure)
	ProviderConfiguration.AddConfig(healthcare.Configure)
	ProviderConfiguration.AddConfig(gkehub.Configure)
	ProviderConfiguration.AddConfig(monitoring.Configure)
}
