package tag

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vsphere_tag", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "tag"
	r.References["tagcategory"] = config.Reference{
            Type: "github.com/AitorLeon89/provider-vsphere/apis/tagcategory/v1alpha1.id",
        }
    })
}
