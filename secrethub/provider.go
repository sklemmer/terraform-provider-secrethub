package secrethub

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/secrethub/secrethub-go/pkg/secrethub"
)

// Provider returns the ScretHub Terraform provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credential": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECRETHUB_CREDENTIAL", nil),
				Description: "Credential to use for SecretHub authentication. Can also be sourced from SECRETHUB_CREDENTIAL.",
			},
			"credential_passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SECRETHUB_CREDENTIAL_PASSPHRASE", nil),
				Description: "Passphrase to unlock the authentication passed in `credential`. Can also be sourced from SECRETHUB_CREDENTIAL_PASSPHRASE.",
			},
			"path_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The default value to prefix path values with. If set, paths for resources and data sources will be prefixed with the given prefix, allowing you to use relative paths instead. If left blank, every path must be absolute (namespace/repository/[dir/]secret_name).",
			},
		},
		ConfigureFunc: configureProvider,
		ResourcesMap: map[string]*schema.Resource{
			"secrethub_secret": resourceSecret(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"secrethub_secret": dataSourceSecret(),
		},
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	credRaw := d.Get("credential").(string)
	passphrase := d.Get("credential_passphrase").(string)

	cred, err := secrethub.NewCredential(credRaw, passphrase)
	if err != nil {
		return nil, err
	}

	client := secrethub.NewClient(cred, nil)
	pathPrefix := d.Get("path_prefix").(string)
	return providerMeta{&client, pathPrefix}, nil
}

type providerMeta struct {
	client     *secrethub.Client
	pathPrefix string
}
