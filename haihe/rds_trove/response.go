package rds_trove

import "time"

type ResDatastore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}


type ResInstance struct {
	Id        string       `json:"id"`
	Status    string       `json:"status"`
	Updated   string       `json:"updated"`
	Created   string       `json:"created"`
	Name      string       `json:"name"`
	TenantId  string       `json:"tenant_id"`
	Datastore ResDatastore `json:"datastore"`
}

type CreateInstanceResponse struct {
	Instance ResInstance `json:"instance"`
}

type InstanceConf struct {
	Configuration map[string]interface{} `json:"configuration"`
}

type GetInstanceConfResponse struct {
	Instance InstanceConf `json:"instance"`
}

type Database struct {
	Name string `json:"name"`
}
type User struct {
	Host      string     `json:"host"`
	Name      string     `json:"name"`
	Databases []Database `json:"databases"`
}
type GetUsersResponse struct {
	Users []User `json:"users"`
}

type GetDatabasesResponse struct {
	Databases []Database `json:"databases"`
}

type Conf struct {
	DatastoreName        string `json:"datastore_name"`
	Updated              string `json:"updated"`
	Name                 string `json:"name"`
	Created              string `json:"created"`
	DatastoreVersionName string `json:"datastore_version_name"`
	Id                   string `json:"id"`
	DatastoreVersionId   string `json:"datastore_version_id"`
	Description          string `json:"description"`
}

type ConfValues struct {
	Conf
	Values map[string]interface{} `json:"values"`
}

type GetConfGroupsResponse struct {
	Configurations []Conf `json:"configurations"`
}

type ConfInstance struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetConfInstancesResponse struct {
	Instances []ConfInstance `json:"instances"`
}

type GetConfItmesResponse struct {
	Configuration ConfValues `json:"configuration"`
}

type CreateConfGroupResponse struct {
	Configuration ConfValues `json:"configuration"`
}

type DiskUsageResponse struct {
	Used int `json:"used"`
	Total int `json:"total"`
}

type BackupResponse struct {
	Id string `json:"id"`
	FromLsn string `json:"fromLsn"`
	ToLsn string `json:"toLsn"`
	BinlogPos string `json:"binlogPos"`
	BinlogFilename string `json:"binlogFilename"`
	SubDir string `json:"subDir"`
	BackupAt time.Time `json:"backupAt"`
}