/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/upbound/upjet/pkg/config"
	"context"
//	"fmt"
	"strings"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vsphere_license": config.IdentifierFromProvider,
        "vsphere_datacenter": config.NameAsIdentifier,
        "vsphere_host": config.IdentifierFromProvider,
        "vsphere_custom_attribute": config.NameAsIdentifier,
        "vsphere_folder": config.IdentifierFromProvider,
        "vsphere_content_library": config.IdentifierFromProvider,
        "vsphere_tag_category": config.NameAsIdentifier,
        "vsphere_tag": config.IdentifierFromProvider,
//        "vsphere_tag": config.TemplatedStringAsIdentifier("name", "'{\"category_name\": \"{{ .parameter.category_id }}\", \"tag_name\": \"{{ .externalName }}\"}'"),
        "vsphere_host_virtual_switch": config.TemplatedStringAsIdentifier("name", "tf-HostVirtualSwitch:{{ .parameter.host_system_id }}:{{ .externalName }}"),
        "vsphere_host_port_group": config.TemplatedStringAsIdentifier("name", "tf-HostPortGroup:{{ .parameter.host_system_id }}:{{ .externalName }}"),
//        "vsphere_distributed_virtual_switch": config.TemplatedStringAsIdentifier("name", "/{{ .parameter.datacenter_id }}/network/{{ .externalName }}"),
//        "vsphere_distributed_port_group": config.TemplatedStringAsIdentifier("name", "/{{ .parameter.datacenter_id }}/network/{{ .externalName }}"),
        "vsphere_distributed_virtual_switch": config.NameAsIdentifier,
        "vsphere_distributed_port_group": config.NameAsIdentifier,
        "vsphere_vnic": config.IdentifierFromProvider,
        "vsphere_role": config.IdentifierFromProvider,
        "vsphere_entity_permissions": config.IdentifierFromProvider,
        "vsphere_vmfs_datastore": config.IdentifierFromProvider,
        "vsphere_nas_datastore": config.IdentifierFromProvider,
        "vsphere_file": config.IdentifierFromProvider,
        "vsphere_vm_storage_policy": config.IdentifierFromProvider,
        "vsphere_datastore_cluster": config.TemplatedStringAsIdentifier("name","/{{ .parameter.datacenter_id }}/datastore/{{ .externalName }}"),
//        "vsphere_compute_cluster": config.TemplatedStringAsIdentifier("name","/{{ .parameter.datacenter_id }}/datastore/{{ .externalName }}"),
        "vsphere_compute_cluster": config.IdentifierFromProvider,
        "vsphere_resource_pool": config.IdentifierFromProvider,
        "vsphere_virtual_disk": config.ParameterAsIdentifier("vmdk_path"),
        "vsphere_vapp_entity": config.IdentifierFromProvider,
        "vsphere_virtual_machine": config.TemplatedStringAsIdentifier("name","/{{ .parameter.datacenter_id }}/vm/{{ .externalName }}"),
        "vsphere_virtual_machine_snapshot": config.IdentifierFromProvider,
        "vsphere_vapp_container": config.IdentifierFromProvider,
        "vsphere_content_library_item": config.IdentifierFromProvider,
        "vsphere_compute_cluster_host_group": config.IdentifierFromProvider,
        "vsphere_compute_cluster_vm_affinity_rule": config.IdentifierFromProvider,
        "vsphere_compute_cluster_vm_anti_affinity_rule": config.IdentifierFromProvider,
        "vsphere_compute_cluster_vm_dependency_rule": config.IdentifierFromProvider,
        "vsphere_compute_cluster_vm_group": config.IdentifierFromProvider,
	"vsphere_compute_cluster_vm_host_rule": config.IdentifierFromProvider,
        "vsphere_dpm_host_override": config.IdentifierFromProvider,
        "vsphere_drs_vm_override": config.IdentifierFromProvider,
        "vsphere_ha_vm_override": config.IdentifierFromProvider,
        "vsphere_datastore_cluster_vm_anti_affinity_rule": config.IdentifierFromProvider,
        "vsphere_storage_drs_vm_override": config.IdentifierFromProvider,
}



// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}



// FormattedIdentifierUserDefined is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, vpc_id:cluster_name where vpc_id comes from spec
// but cluster_name is a naming field we can use external name for.
func FormattedIdentifierUserDefined(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys)+1)
		for i, k := range keys {
			v, ok := parameters[k]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", k)
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be a string", k)
			}
			vals[i] = s
		}
		vals[len(vals)-1] = externalName
		return strings.Join(vals, separator), nil
	}
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		s, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}
		w := strings.Split(s, separator)
		return w[len(w)-1], nil
	}
	return e
}


// FormattedIdentifierFromProvider is a helper function to construct Terraform
// IDs that use elements from the parameters in a certain string format.
// It should be used in cases where all information in the ID is gathered from
// the spec and not user defined like name. For example, zone_id:vpc_id.
func FormattedIdentifierFromProvider(separator string, keys ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys))
		for i, key := range keys {
			val, ok := parameters[key]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", key)
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be string", key)
			}
			vals[i] = s
		}
		return strings.Join(vals, separator), nil
	}
	return e
}




// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
