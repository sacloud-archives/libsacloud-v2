// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-envelope'; DO NOT EDIT

package sacloud

import (
	"github.com/sacloud/libsacloud-v2/sacloud/naked"
)

// NoteFindRequestEnvelope is envelop of API request
type NoteFindRequestEnvelope struct {
	Count   int                    `json:",omitempty"`
	From    int                    `json:",omitempty"`
	Sort    []string               `json:",omitempty"`
	Filter  map[string]interface{} `json:",omitempty"`
	Include []string               `json:",omitempty"`
	Exclude []string               `json:",omitempty"`
}

// NoteFindResponseEnvelope is envelop of API response
type NoteFindResponseEnvelope struct {
	Total int `json:",omitempty"` // トータル件数
	From  int `json:",omitempty"` // ページング開始ページ
	Count int `json:",omitempty"` // 件数

	Notes []*naked.Note `json:",omitempty"`
}

// NoteCreateRequestEnvelope is envelop of API request
type NoteCreateRequestEnvelope struct {
	Note *naked.Note `json:",omitempty"`
}

// NoteCreateResponseEnvelope is envelop of API response
type NoteCreateResponseEnvelope struct {
	IsOk    bool `json:"is_ok,omitempty"` // is_ok項目
	Success bool `json:",omitempty"`      // success項目

	Note *naked.Note `json:",omitempty"`
}

// NoteReadResponseEnvelope is envelop of API response
type NoteReadResponseEnvelope struct {
	IsOk    bool `json:"is_ok,omitempty"` // is_ok項目
	Success bool `json:",omitempty"`      // success項目

	Note *naked.Note `json:",omitempty"`
}

// NoteUpdateRequestEnvelope is envelop of API request
type NoteUpdateRequestEnvelope struct {
	Note *naked.Note `json:",omitempty"`
}

// NoteUpdateResponseEnvelope is envelop of API response
type NoteUpdateResponseEnvelope struct {
	IsOk    bool `json:"is_ok,omitempty"` // is_ok項目
	Success bool `json:",omitempty"`      // success項目

	Note *naked.Note `json:",omitempty"`
}

// ZoneFindRequestEnvelope is envelop of API request
type ZoneFindRequestEnvelope struct {
	Count   int                    `json:",omitempty"`
	From    int                    `json:",omitempty"`
	Sort    []string               `json:",omitempty"`
	Filter  map[string]interface{} `json:",omitempty"`
	Include []string               `json:",omitempty"`
	Exclude []string               `json:",omitempty"`
}

// ZoneFindResponseEnvelope is envelop of API response
type ZoneFindResponseEnvelope struct {
	Total int `json:",omitempty"` // トータル件数
	From  int `json:",omitempty"` // ページング開始ページ
	Count int `json:",omitempty"` // 件数

	Zones []*naked.Zone `json:",omitempty"`
}

// ZoneReadResponseEnvelope is envelop of API response
type ZoneReadResponseEnvelope struct {
	IsOk    bool `json:"is_ok,omitempty"` // is_ok項目
	Success bool `json:",omitempty"`      // success項目

	Zone *naked.Zone `json:",omitempty"`
}
