package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

type monitorsDef struct{}

var monitors = &monitorsDef{}

func (m *monitorsDef) cpuTimeModel() *schema.Model {
	return &schema.Model{
		Name: "CPUTimeActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorCPUTimeValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "CPU",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) diskModel() *schema.Model {
	return &schema.Model{
		Name: "DiskActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorDiskValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "Disk",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) interfaceModel() *schema.Model {
	return &schema.Model{
		Name: "InterfaceActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorInterfaceValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "Interface",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) routerModel() *schema.Model {
	return &schema.Model{
		Name: "RouterActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorRouterValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "Router",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) databaseModel() *schema.Model {
	return &schema.Model{
		Name: "DatabaseActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorDatabaseValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "Database",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) freeDiskSizeModel() *schema.Model {
	return &schema.Model{
		Name: "FreeDiskSizeActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorFreeDiskSizeValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "FreeDiskSize",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) responseTimeSecModel() *schema.Model {
	return &schema.Model{
		Name: "ResponseTimeSecActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorResponseTimeSecValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "ResponseTimeSec",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) linkModel() *schema.Model {
	return &schema.Model{
		Name: "LinkActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorLinkValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "Link",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}

func (m *monitorsDef) connectionModel() *schema.Model {
	return &schema.Model{
		Name: "ConnectionActivity",
		Fields: []*schema.FieldDesc{
			{
				Name:     "Values",
				Type:     meta.Static([]naked.MonitorConnectionValue{}),
				ReadOnly: true,
				Tags: &schema.FieldTags{
					MapConv: "Connection",
				},
			},
		},
		NakedType: meta.Static(naked.MonitorValues{}),
	}
}