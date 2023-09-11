/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/AitorLeon89/provider-vsphere/internal/controller/computecluster/cluster"
	clusterhostgroup "github.com/AitorLeon89/provider-vsphere/internal/controller/computeclusterhostgroup/clusterhostgroup"
	clustervmaffinityrule "github.com/AitorLeon89/provider-vsphere/internal/controller/computeclustervmaffinityrule/clustervmaffinityrule"
	clustervmantiaffinityrule "github.com/AitorLeon89/provider-vsphere/internal/controller/computeclustervmantiaffinityrule/clustervmantiaffinityrule"
	clustervmdependencyrule "github.com/AitorLeon89/provider-vsphere/internal/controller/computeclustervmdependencyrule/clustervmdependencyrule"
	clustervmgroup "github.com/AitorLeon89/provider-vsphere/internal/controller/computeclustervmgroup/clustervmgroup"
	clustervmhostrule "github.com/AitorLeon89/provider-vsphere/internal/controller/computeclustervmhostrule/clustervmhostrule"
	library "github.com/AitorLeon89/provider-vsphere/internal/controller/contentlibrary/library"
	libraryitem "github.com/AitorLeon89/provider-vsphere/internal/controller/contentlibraryitem/libraryitem"
	attribute "github.com/AitorLeon89/provider-vsphere/internal/controller/customattribute/attribute"
	datacenter "github.com/AitorLeon89/provider-vsphere/internal/controller/datacenter/datacenter"
	clusterdatastorecluster "github.com/AitorLeon89/provider-vsphere/internal/controller/datastorecluster/cluster"
	clustervmantiaffinityruledatastoreclustervmantiaffinityrule "github.com/AitorLeon89/provider-vsphere/internal/controller/datastoreclustervmantiaffinityrule/clustervmantiaffinityrule"
	portgroup "github.com/AitorLeon89/provider-vsphere/internal/controller/distributedportgroup/portgroup"
	virtualswitch "github.com/AitorLeon89/provider-vsphere/internal/controller/distributedvirtualswitch/virtualswitch"
	hostoverride "github.com/AitorLeon89/provider-vsphere/internal/controller/dpmhostoverride/hostoverride"
	vmoverride "github.com/AitorLeon89/provider-vsphere/internal/controller/drsvmoverride/vmoverride"
	permissions "github.com/AitorLeon89/provider-vsphere/internal/controller/entitypermissions/permissions"
	file "github.com/AitorLeon89/provider-vsphere/internal/controller/file/file"
	folder "github.com/AitorLeon89/provider-vsphere/internal/controller/folder/folder"
	vmoverridehavmoverride "github.com/AitorLeon89/provider-vsphere/internal/controller/havmoverride/vmoverride"
	host "github.com/AitorLeon89/provider-vsphere/internal/controller/host/host"
	portgrouphostportgroup "github.com/AitorLeon89/provider-vsphere/internal/controller/hostportgroup/portgroup"
	virtualswitchhostvirtualswitch "github.com/AitorLeon89/provider-vsphere/internal/controller/hostvirtualswitch/virtualswitch"
	license "github.com/AitorLeon89/provider-vsphere/internal/controller/license/license"
	datastore "github.com/AitorLeon89/provider-vsphere/internal/controller/nasdatastore/datastore"
	providerconfig "github.com/AitorLeon89/provider-vsphere/internal/controller/providerconfig"
	pool "github.com/AitorLeon89/provider-vsphere/internal/controller/resourcepool/pool"
	role "github.com/AitorLeon89/provider-vsphere/internal/controller/role/role"
	drsvmoverride "github.com/AitorLeon89/provider-vsphere/internal/controller/storagedrsvmoverride/drsvmoverride"
	tag "github.com/AitorLeon89/provider-vsphere/internal/controller/tag/tag"
	category "github.com/AitorLeon89/provider-vsphere/internal/controller/tagcategory/category"
	container "github.com/AitorLeon89/provider-vsphere/internal/controller/vappcontainer/container"
	entity "github.com/AitorLeon89/provider-vsphere/internal/controller/vappentity/entity"
	disk "github.com/AitorLeon89/provider-vsphere/internal/controller/virtualdisk/disk"
	machine "github.com/AitorLeon89/provider-vsphere/internal/controller/virtualmachine/machine"
	machinesnapshot "github.com/AitorLeon89/provider-vsphere/internal/controller/virtualmachinesnapshot/machinesnapshot"
	datastorevmfsdatastore "github.com/AitorLeon89/provider-vsphere/internal/controller/vmfsdatastore/datastore"
	storagepolicy "github.com/AitorLeon89/provider-vsphere/internal/controller/vmstoragepolicy/storagepolicy"
	vnic "github.com/AitorLeon89/provider-vsphere/internal/controller/vnic/vnic"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		clusterhostgroup.Setup,
		clustervmaffinityrule.Setup,
		clustervmantiaffinityrule.Setup,
		clustervmdependencyrule.Setup,
		clustervmgroup.Setup,
		clustervmhostrule.Setup,
		library.Setup,
		libraryitem.Setup,
		attribute.Setup,
		datacenter.Setup,
		clusterdatastorecluster.Setup,
		clustervmantiaffinityruledatastoreclustervmantiaffinityrule.Setup,
		portgroup.Setup,
		virtualswitch.Setup,
		hostoverride.Setup,
		vmoverride.Setup,
		permissions.Setup,
		file.Setup,
		folder.Setup,
		vmoverridehavmoverride.Setup,
		host.Setup,
		portgrouphostportgroup.Setup,
		virtualswitchhostvirtualswitch.Setup,
		license.Setup,
		datastore.Setup,
		providerconfig.Setup,
		pool.Setup,
		role.Setup,
		drsvmoverride.Setup,
		tag.Setup,
		category.Setup,
		container.Setup,
		entity.Setup,
		disk.Setup,
		machine.Setup,
		machinesnapshot.Setup,
		datastorevmfsdatastore.Setup,
		storagepolicy.Setup,
		vnic.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
