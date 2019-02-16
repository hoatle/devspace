package cloud

import "github.com/covexo/devspace/pkg/devspace/docker"

// GetFirstPublicRegistry retrieves the first public registry
func (p *Provider) GetFirstPublicRegistry() (string, error) {
	registries, err := p.GetRegistries()
	if err != nil {
		return "", err
	}

	registryURL := ""
	for _, registry := range registries {
		if registry.OwnerID == nil {
			registryURL = registry.URL
			break
		}
	}

	return registryURL, nil
}

// LoginIntoRegistries logs the user into the user docker registries
func (p *Provider) LoginIntoRegistries() error {
	registries, err := p.GetRegistries()
	if err != nil {
		return err
	}

	// We don't want the minikube client to login into the registry
	client, err := docker.NewClient(false)
	if err != nil {
		return err
	}

	// Get account name
	accountName, err := p.GetAccountName()
	if err != nil {
		return err
	}

	for _, registry := range registries {
		// Login
		_, err = docker.Login(client, registry.URL, accountName, p.Token, true, true)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoginIntoRegistry logs the user into the user docker registry
func (p *Provider) LoginIntoRegistry(name string) error {
	// We don't want the minikube client to login into the registry
	client, err := docker.NewClient(false)
	if err != nil {
		return err
	}

	// Get account name
	accountName, err := p.GetAccountName()
	if err != nil {
		return err
	}

	// Get account name
	_, err = docker.Login(client, name, accountName, p.Token, true, true)
	if err != nil {
		return err
	}

	return nil
}