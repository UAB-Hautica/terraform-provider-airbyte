// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

// FileType - The file type you want to sync. Currently only 'csv' and 'json' files are supported.
type FileType string

const (
	FileTypeCsv  FileType = "csv"
	FileTypeJSON FileType = "json"
)

func (e FileType) ToPointer() *FileType {
	return &e
}

func (e *FileType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		fallthrough
	case "json":
		*e = FileType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileType: %v", v)
	}
}

type SourceSftpBulkUpdate struct {
	// Sync only the most recent file for the configured folder path and file pattern
	FileMostRecent *bool `default:"false" json:"file_most_recent"`
	// The regular expression to specify files for sync in a chosen Folder Path
	FilePattern *string `default:"" json:"file_pattern"`
	// The file type you want to sync. Currently only 'csv' and 'json' files are supported.
	FileType *FileType `default:"csv" json:"file_type"`
	// The directory to search files for sync
	FolderPath *string `default:"" json:"folder_path"`
	// The server host address
	Host string `json:"host"`
	// OS-level password for logging into the jump server host
	Password *string `json:"password,omitempty"`
	// The server port
	Port *int64 `default:"22" json:"port"`
	// The private key
	PrivateKey *string `json:"private_key,omitempty"`
	// The separator used in the CSV files. Define None if you want to use the Sniffer functionality
	Separator *string `default:"," json:"separator"`
	// The date from which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// The name of the stream or table you want to create
	StreamName string `json:"stream_name"`
	// The server user
	Username string `json:"username"`
}

func (s SourceSftpBulkUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSftpBulkUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSftpBulkUpdate) GetFileMostRecent() *bool {
	if o == nil {
		return nil
	}
	return o.FileMostRecent
}

func (o *SourceSftpBulkUpdate) GetFilePattern() *string {
	if o == nil {
		return nil
	}
	return o.FilePattern
}

func (o *SourceSftpBulkUpdate) GetFileType() *FileType {
	if o == nil {
		return nil
	}
	return o.FileType
}

func (o *SourceSftpBulkUpdate) GetFolderPath() *string {
	if o == nil {
		return nil
	}
	return o.FolderPath
}

func (o *SourceSftpBulkUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceSftpBulkUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceSftpBulkUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceSftpBulkUpdate) GetPrivateKey() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKey
}

func (o *SourceSftpBulkUpdate) GetSeparator() *string {
	if o == nil {
		return nil
	}
	return o.Separator
}

func (o *SourceSftpBulkUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceSftpBulkUpdate) GetStreamName() string {
	if o == nil {
		return ""
	}
	return o.StreamName
}

func (o *SourceSftpBulkUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
