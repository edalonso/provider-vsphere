/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/AitorLeon89/provider-vsphere/config/license"
        "github.com/AitorLeon89/provider-vsphere/config/datacenter"
        "github.com/AitorLeon89/provider-vsphere/config/host"
        "github.com/AitorLeon89/provider-vsphere/config/customattribute"
        "github.com/AitorLeon89/provider-vsphere/config/contentlibrary"
        "github.com/AitorLeon89/provider-vsphere/config/folder"
        "github.com/AitorLeon89/provider-vsphere/config/tagcategory"
        "github.com/AitorLeon89/provider-vsphere/config/tag"
        "github.com/AitorLeon89/provider-vsphere/config/hostvirtualswitch"
        "github.com/AitorLeon89/provider-vsphere/config/hostportgroup"
        "github.com/AitorLeon89/provider-vsphere/config/distributedvirtualswitch"
        "github.com/AitorLeon89/provider-vsphere/config/distributedportgroup"
        "github.com/AitorLeon89/provider-vsphere/config/vnic"
        "github.com/AitorLeon89/provider-vsphere/config/role"
        "github.com/AitorLeon89/provider-vsphere/config/entitypermissions"
        "github.com/AitorLeon89/provider-vsphere/config/vmfsdatastore"
        "github.com/AitorLeon89/provider-vsphere/config/nasdatastore"
        "github.com/AitorLeon89/provider-vsphere/config/file"
        "github.com/AitorLeon89/provider-vsphere/config/vmstoragepolicy"
        "github.com/AitorLeon89/provider-vsphere/config/datastorecluster"
        "github.com/AitorLeon89/provider-vsphere/config/computecluster"
        "github.com/AitorLeon89/provider-vsphere/config/resourcepool"
        "github.com/AitorLeon89/provider-vsphere/config/virtualdisk"
        "github.com/AitorLeon89/provider-vsphere/config/vappentity"
        "github.com/AitorLeon89/provider-vsphere/config/virtualmachine"
        "github.com/AitorLeon89/provider-vsphere/config/virtualmachinesnapshot"
        "github.com/AitorLeon89/provider-vsphere/config/vappcontainer"
        "github.com/AitorLeon89/provider-vsphere/config/contentlibraryitem"
        "github.com/AitorLeon89/provider-vsphere/config/computeclusterhostgroup"
        "github.com/AitorLeon89/provider-vsphere/config/computeclustervmaffinityrule"
        "github.com/AitorLeon89/provider-vsphere/config/computeclustervmantiaffinityrule"
        "github.com/AitorLeon89/provider-vsphere/config/computeclustervmdependencyrule"
        "github.com/AitorLeon89/provider-vsphere/config/computeclustervmgroup"
        "github.com/AitorLeon89/provider-vsphere/config/computeclustervmhostrule"
        "github.com/AitorLeon89/provider-vsphere/config/drsvmoverride"
        "github.com/AitorLeon89/provider-vsphere/config/dpmhostoverride"
        "github.com/AitorLeon89/provider-vsphere/config/havmoverride"
        "github.com/AitorLeon89/provider-vsphere/config/storagedrsvmoverride"
        "github.com/AitorLeon89/provider-vsphere/config/datastoreclustervmantiaffinityrule"
)

const (
	resourcePrefix = "vsphere"
	modulePath     = "github.com/AitorLeon89/provider-vsphere"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
                license.Configure,
                datacenter.Configure,
		host.Configure,
                customattribute.Configure,
                folder.Configure,
                contentlibrary.Configure,
		tagcategory.Configure,
                tag.Configure,
                hostvirtualswitch.Configure,
                hostportgroup.Configure,
                distributedvirtualswitch.Configure,
                distributedportgroup.Configure,
                vnic.Configure,
                role.Configure,
                entitypermissions.Configure,
                vmfsdatastore.Configure,
                nasdatastore.Configure,
                file.Configure,
                vmstoragepolicy.Configure,
                datastorecluster.Configure,
                computecluster.Configure,
                resourcepool.Configure,
                virtualdisk.Configure,
                vappentity.Configure,
                virtualmachine.Configure,
                virtualmachinesnapshot.Configure,
                vappcontainer.Configure,
                contentlibraryitem.Configure,
		computeclusterhostgroup.Configure,
                computeclustervmaffinityrule.Configure,
                computeclustervmantiaffinityrule.Configure,
		computeclustervmdependencyrule.Configure,
                computeclustervmgroup.Configure,
                computeclustervmhostrule.Configure,
                drsvmoverride.Configure,
                dpmhostoverride.Configure,
		havmoverride.Configure,
                storagedrsvmoverride.Configure,
		datastoreclustervmantiaffinityrule.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

