package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

func init() {
	nakedType := meta.Static(naked.CDROM{})

	cdrom := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.ID(),
			fields.Name(),
			fields.Description(),
			fields.DisplayOrder(),
			fields.Tags(),
			fields.Availability(),
			fields.Scope(),
			fields.StorageClass(),
			fields.Storage(),
			fields.Icon(),
			fields.CreatedAt(),
			fields.ModifiedAt(),
		},
	}

	ftpServer := &schema.Model{
		Name:      "FTPServer",
		NakedType: meta.Static(naked.OpeningFTPServer{}),
		Fields: []*schema.FieldDesc{
			fields.HostName(),
			fields.IPAddress(),
			fields.User(),
			fields.Password(),
		},
	}

	createParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.SizeMB(),
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	updateParam := &schema.Model{
		Fields: []*schema.FieldDesc{
			fields.Name(),
			fields.Description(),
			fields.Tags(),
			fields.IconID(),
		},
	}

	Resources.DefineWith("CDROM", func(r *schema.Resource) {
		r.Operations(
			// find
			r.DefineOperationFind(nakedType, findParameter, cdrom),
			// create
			r.DefineOperationCreate(nakedType, createParam, cdrom).
				ResponseEnvelope(&schema.EnvelopePayloadDesc{ // TODO エンベロープとResultを同時に定義できないか?
					PayloadName: "FTPServer",
					PayloadType: meta.Static(naked.OpeningFTPServer{}),
				}).
				ResultWithSourceField("FTPServer", ftpServer),
			// read
			r.DefineOperationRead(nakedType, cdrom),
			// update
			r.DefineOperationUpdate(nakedType, updateParam, cdrom),
			// delete
			r.DefineOperationDelete(),
		)
	})
}
