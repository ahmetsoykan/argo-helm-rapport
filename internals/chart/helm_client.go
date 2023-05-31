package chart

import (
	helmclient "github.com/mittwald/go-helm-client"
)

type HelmClient struct {
	helmclient.Client
}

func NewHelmClient() (HelmClient, error) {

	opt := &helmclient.Options{
		Namespace:        "default",
		RepositoryCache:  "/tmp/.helmcache",
		RepositoryConfig: "/tmp/.helmrepo",
		Debug:            true,
	}

	client, err := helmclient.New(opt)
	if err != nil {
		return HelmClient{}, err
	}

	return HelmClient{client}, nil
}
