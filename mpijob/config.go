package mpijob

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/kubeflow/mpi-operator/pkg/apis/kubeflow/v1alpha1"
	"gopkg.in/yaml.v2"
	"k8s.io/api/core/v1"
)

type Config struct {
	key    string
	MPIJob *v1alpha1.MPIJob
}

const (
	Kind          = "MPIJob"
	GroupName     = "kubeflow.org"
	GroupVersion  = "v1alpha1"
	ContainerName = "mpi-container"
)

func NewConfig(key string) *Config {
	config := new(Config)
	config.key = key
	config.MPIJob = new(v1alpha1.MPIJob)
	config.MPIJob.Kind = Kind
	config.MPIJob.APIVersion = filepath.Join(GroupName, GroupVersion)

	return config
}

func GetPodTemplateSpec(b []byte) (*v1.PodTemplateSpec, error) {
	pts := new(v1.PodTemplateSpec)
	if err := json.Unmarshal(b, pts); err != nil {
		return nil, err
	}

	for i := range pts.Spec.Containers {
		pts.Spec.Containers[i].Name = ContainerName
	}

	return pts, nil
}

func (c *Config) Yaml() error {
	if c == nil || c.key == "" {
		return fmt.Errorf("invalid key or nil Config struct")
	}

	b, err := json.Marshal(c.MPIJob)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})

	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	b, err = yaml.Marshal(m)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.key+".yaml", b, 0644)
}
