// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-interfaces'; DO NOT EDIT

package sacloud

import (
	"context"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

/*************************************************
* CDROMAPI
*************************************************/

// CDROMAPI is interface for operate CDROM resource
type CDROMAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*CDROM, error)
	Create(ctx context.Context, zone string, param *CDROMCreateRequest) (*CDROM, *FTPServer, error)
	Read(ctx context.Context, zone string, id types.ID) (*CDROM, error)
	Update(ctx context.Context, zone string, id types.ID, param *CDROMUpdateRequest) (*CDROM, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	OpenFTP(ctx context.Context, zone string, id types.ID, openOption *OpenFTPParam) (*FTPServer, error)
	CloseFTP(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* NFSAPI
*************************************************/

// NFSAPI is interface for operate NFS resource
type NFSAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*NFS, error)
	Create(ctx context.Context, zone string, param *NFSCreateRequest) (*NFS, error)
	Read(ctx context.Context, zone string, id types.ID) (*NFS, error)
	Update(ctx context.Context, zone string, id types.ID, param *NFSUpdateRequest) (*NFS, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* NoteAPI
*************************************************/

// NoteAPI is interface for operate Note resource
type NoteAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Note, error)
	Create(ctx context.Context, zone string, param *NoteCreateRequest) (*Note, error)
	Read(ctx context.Context, zone string, id types.ID) (*Note, error)
	Update(ctx context.Context, zone string, id types.ID, param *NoteUpdateRequest) (*Note, error)
	Delete(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* SwitchAPI
*************************************************/

// SwitchAPI is interface for operate Switch resource
type SwitchAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Switch, error)
	Create(ctx context.Context, zone string, param *SwitchCreateRequest) (*Switch, error)
	Read(ctx context.Context, zone string, id types.ID) (*Switch, error)
	Update(ctx context.Context, zone string, id types.ID, param *SwitchUpdateRequest) (*Switch, error)
	Delete(ctx context.Context, zone string, id types.ID) error
	ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error
	DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error
}

/*************************************************
* ZoneAPI
*************************************************/

// ZoneAPI is interface for operate Zone resource
type ZoneAPI interface {
	Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Zone, error)
	Read(ctx context.Context, zone string, id types.ID) (*Zone, error)
}
