package mpijob

import (
	"testing"

	"k8s.io/api/core/v1"
)

func TestNewConfig(t *testing.T) {
	key := "mpi-job"
	config := NewConfig(key)

	job := config.MPIJob
	job.Name = key
	job.Spec.Replicas = new(int32)
	*(job.Spec.Replicas) = 3

	job.Spec.WorkerPodTemplateSpec = v1.PodTemplateSpec{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "tensorflow",
					Image: "gcr.io/tf-on-k8s-dogfood/tf_sample:dc944ff",
				},
			},
		},
	}

	if err := config.Yaml(); err != nil {
		t.Fatal(err)
	}
}
