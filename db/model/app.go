package model

type AppStatus struct {
	EventID     string `gorm:"column:event_id;size:32;primary_key" json:"event_id"`
	Format      string `gorm:"column:format;size:32" json:"format"` // only rainbond-app/docker-compose
	SourceDir   string `gorm:"column:source_dir;size:255" json:"source_dir"`
	Apps        string `gorm:"column:apps"`
	Status      string `gorm:"column:status;size:32" json:"status"` // only exporting/importing/failed/success
	TarFileHref string `gorm:"column:tar_file_href;size:255" json:"tar_file_href"`
	Metadata    string `gorm:"column:metadata;type:text;" json:"metadata"`
}

//TableName 表名
func (t *AppStatus) TableName() string {
	return "app_status"
}

//AppBackup app backup info
type AppBackup struct {
	Model
	EventID  string `gorm:"column:event_id;size:32;" json:"event_id"`
	BackupID string `gorm:"column:backup_id;size:32;" json:"backup_id"`
	GroupID  string `gorm:"column:group_id;size:32;" json:"group_id"`
	//Status in starting,failed,success
	Status     string `gorm:"column:status;size:32" json:"status"`
	Version    string `gorm:"column:version;size:32" json:"version"`
	SourceDir  string `gorm:"column:source_dir;size:255" json:"source_dir"`
	BackupMode string `gorm:"column:backup_mode;size:32" json:"backup_mode"`
	BuckupSize int    `gorm:"column:backup_size" json:"backup_size"`
}

//TableName 表名
func (t *AppBackup) TableName() string {
	return "app_backup"
}
