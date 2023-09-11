package host

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vsphere_host", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "host"

        r.References["datacenter"] = config.Reference{
                Type: "github.com/AitorLeon89/provider-vsphere/apis/datacenter/v1alpha1.Datacenter",
        }
    })
}
