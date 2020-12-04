package autoscaling

import "github.com/glumu/go-jewel/haihe/uitls"

type AutoscalingClient struct {
	utils.BaseClient
}

var (
	CreateGroupUrl         string = "/createGroup"
	DeleteGroupUrl         string = "/deleteGroup"
	CreateConfigurationUrl string = "/createConfiguration"
	DeleteConfigurationUrl string = "/deleteConfiguration"
	CreatePolicyUrl        string = "/createPolicy"
	DeletePolicyUrl        string = "/deletePolicy"
	ScaleHostInstanceUrl   string = "/scaleHostInstance"
	ListElasticLogUrl      string = "/listElasticLog"
	GetElasticLogUrl       string = "/getElasticLog"
)
