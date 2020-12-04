package rds_trove

type Datastore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Instance struct {
	Name      string    `json:"name"`
	Datastore Datastore `json:"datastore"`
}

type CreateInstanceRequest struct {
	Instance Instance `json:"instance"`
}

type detailFlavor struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Cpu      int    `json:"cpu"`
	DiskSize int    `json:"diskSize"`
	Memory   int    `json:"memory"`
}

type PutInstanceConf struct {
	Configuration string       `json:"configuration,omitempty"`
	Flavor        detailFlavor `json:"flavor"`
}

type PutInstanceConfRequest struct {
	Instance PutInstanceConf `json:"instance"`
}

type CUser struct {
	Databases []map[string]string `json:"databases" binding:"required"`
	Name      string              `json:"name" binding:"required"`
	Password  string              `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Users []CUser `json:"users" binding:"required"`
}

type CDatabase struct {
	CharacterSet string `json:"character_set,omitempty"`
	Collate      string `json:"collate,omitempty"`
	Name         string `json:"name" bidding:"required"`
}

type CreateDatabaseRequest struct {
	Databases []CDatabase `json:"databases" bidding:"required"`
}

type CConf struct {
	Datastore   Datastore              `json:"datastore"`
	Values      map[string]interface{} `json:"values"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
}

type CreateConfGroupRequest struct {
	Configuration CConf `json:"configuration"`
}

type ConfItems struct {
	Values map[string]interface{} `json:"values"`
}

type AddConfItemsRequest struct {
	Configuration ConfItems `json:"configuration"`
}

type ResetConf struct {
	Values      map[string]interface{} `json:"values"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
}
type ResetConfRequest struct {
	Configuration ResetConf `json:"configuration"`
}

type BuildArgs struct {
	Path      string `json:"path" bidding:"required"`
	NeedStart bool   `json:"need_start" bidding:"required"`
	Step      int    `json:"step" bidding:"required"`
	Size      int    `json:"size" bidding:"required"`
	Port      string `json:"port" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
type BuildRequest struct {
	Build BuildArgs `json:"build"`
}

type StopRequest struct {
	Stop map[string]interface{} `json:"stop"`
}

type StartRequest struct {
	Start map[string]interface{} `json:"start"`
}

type RestartRequest struct {
	Restart map[string]interface{} `json:"restart"`
}

type resizeLvm struct {
	Path string `json:"path"`
	Size int    `json:"size"`
}

type ResizeLvmRequest struct {
	ResizeLvm resizeLvm `json:"resize_lvm"`
}

type DiskBody struct {
	Path string `json:"path"`
	Disk string `json:"disk"`
}

type MountRequest struct {
	MountDisk DiskBody `json:"mount_disk"`
}

type UmountRequest struct {
	UmountDisk DiskBody `json:"umount_disk"`
}

type GetDiskUsageRequest struct {
	GetDiskUsage DiskBody `json:"get_disk_usage"`
}

type Resize2FSRequest struct {
	Resize2FS DiskBody `json:"resize2fs"`
}

type BackupBody struct {
	Info string `json:"info"`
	LogPath string `json:"logPath"`
	BackupPath string `json:"backupPath"`
}

type BackupRequest struct {
	Backup BackupBody `json:"backup"`
}

type MergeIncBody struct {
	Backups []string `json:"backups"`
	LogPath string `json:"logPath"`
	BackupPath string `json:"backupPath"`
}

type MergeIncRequest struct {
	MergeInc2Full MergeIncBody `json:"merge_inc2full"`
}

type DeleteIncBody struct {
	Backups []string `json:"backups"`
	LogPath string `json:"logPath"`
}

type DeleteIncRequest struct {
	DeleteInc DeleteIncBody `json:"delete_inc"`
}

type CreateClusterArg struct {
	Name          string     `json:"name"`
	Id            string     `json:"id"`
	Datastore     Datastore  `json:"datastore"`
	Instances     []struct{} `json:"instances"`
	Configuration string     `json:"configuration"`
}
type CreateClusterBackupInfoArg struct {
	LogFile     string `json:"log_file"`
	LogPosition string `json:"log_position"`
}
type CreateClusterRequest struct {
	Cluster    CreateClusterArg           `json:"cluster"`
	BackupInfo CreateClusterBackupInfoArg `json:"backup_info"`
}
