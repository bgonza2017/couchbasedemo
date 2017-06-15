package types

import (

)

const (

)

var (

)

// Right record
type Right struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Disabled bool   `protobuf:"varint,2,opt,name=disabled" json:"disabled,omitempty"`
}
// Service record
type Service struct {
	ID     string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Rights []Right `protobuf:"bytes,3,rep,name=rights" json:"rights,omitempty"`
}

// Account record
type Account struct {
	AccountID     	string `protobuf:"bytes,1,opt,name=AccountID" json:"accountid,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Disabled 		string `protobuf:"varint,2,opt,name=disabled" json:"disabled,omitempty"`
}

