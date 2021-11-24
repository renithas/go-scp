//using yaml package , unmarshall a yaml file and if one key is provided in the command line, get the value from the yaml file
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	//"gopkg.in/yaml.v3"
)

var yamlFile string = "kube.yaml"

//var mdata = make(map[interface{}]interface{})

type T struct {
	Version  string `yaml:"apiVersion,omitempty"`
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name string `yaml:"name,omitempty"`
	} `yaml:"metadata"`
	Spec struct {
		Replicas int `yaml:"replicas"`
		Selector struct {
			MatchLabels struct {
				App string `yaml:"app"`
			} `yaml:"matchLabels"`
		} `yaml:"selector"`

		Template struct {
			Metadata struct {
				Labels struct {
					App string `yaml:"app"`
				} `yaml:"labels"`
			} `yaml:"metadata"`
			Spec struct {
				Containers []struct {
					Name            string `yaml:"name"`
					Image           string `yaml:"image"`
					ImagePullPolicy string `yaml:"imagePullPolicy"`
					Ports           []struct {
						ContainerPort string `yaml:"containerPort"`
					} `yaml:"ports"`

					LivenessProbe struct {
						HttpGet struct {
							Path   string `yaml:"path"`
							Port   int    `yaml:"port"`
							Scheme string `yaml:"scheme"`
						} `yaml:"httpGet"`

						InitialDelaySeconds int `yaml:"initialDelaySeconds"`
						PeriodSeconds       int `yaml:"periodSeconds"`
						TimeoutSeconds      int `yaml:"timeoutSeconds"`
					} `yaml:"livenessProbe"`
					readinessProbe struct {
						HttpGet struct {
							Path   string `yaml:"path"`
							Port   int    `yaml:"port"`
							Scheme string `yaml:"scheme"`
						} `yaml:"httpGet"`
						initialDelaySeconds int `yaml:"initialDelaySeconds"`

						TimeoutSeconds int `yaml:"timeoutSeconds"`
					} `yaml:"readinessProbe"`
				} `yaml:"containers"`
			} `yaml:"spec"`
		} `yaml:"template"`
	} `yaml:"spec"`
}

var t T

func main() {
	yFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		//fmt.Printf("No Yaml File is found with the name %s to read!", yamlFile)
		log.Fatal(err)
		return
	}

	err1 := yaml.Unmarshal(yFile, &t)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("LivenessProbe Port is :", t.Spec.Template.Spec.Containers[0].LivenessProbe.HttpGet.Port)
	fmt.Println("Image Name is :", t.Spec.Template.Spec.Containers[0].Image)
	fmt.Println("Label is :", t.Spec.Template.Metadata.Labels.App)
	//fmt.Println(t.Version)
}
