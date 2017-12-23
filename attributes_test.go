package radius

import "testing"

func TestAttributes_AddRaw(t *testing.T) {
	attr := make(Attributes)
	attr.AddRaw(AttrUserName, []byte("username1"))
	attr.AddRaw(AttrUserName, []byte("username2"))
	attr.AddRaw(AttrADSLAgentRemoteId, []byte("id"))
	if string(attr.GetRaw(AttrUserName)) != "username1" {
		t.Fail()
	}
	if len(attr[byte(AttrUserName)]) != 2 {
		t.Fail()
	}
	if string(attr.GetRaw(AttrADSLAgentRemoteId)) != "id" {
		t.Fail()
	}
}

func TestAttributes_Del(t *testing.T) {
	attr := make(Attributes)
	attr.AddRaw(AttrUserName, []byte("username1"))
	attr.AddRaw(AttrUserName, []byte("username2"))
	attr.AddRaw(AttrADSLAgentRemoteId, []byte("id"))
	attr.AddRaw(AttrADSLAgentCircuitId, []byte("cid"))
	attr.Del(AttrUserName)
	attr.Del(AttrADSLAgentRemoteId)
	if attr.GetRaw(AttrUserName) != nil {
		t.Fail()
	}
	if attr.GetRaw(AttrADSLAgentRemoteId) != nil {
		t.Fail()
	}
	if string(attr.GetRaw(AttrADSLAgentCircuitId)) != "cid" {
		t.Fail()
	}
}

func TestAttributes_LookupRaw(t *testing.T) {
	attr := make(Attributes)
	attr.AddRaw(AttrUserName, []byte("username1"))
	if _, ok := attr.LookupRaw(AttrADSLAgentCircuitId); ok {
		t.Fail()
	}
}

func TestAttributes_LookupRaw_Malformation(t *testing.T) {
	attr := make(Attributes)
	attr[VENDOR_SPECIFIC] = append(attr[VENDOR_SPECIFIC], []byte{0x01})
	if _, ok := attr.LookupRaw(AttrADSLAgentCircuitId); ok {
		t.Fail()
	}
}

func TestAttributes_Set(t *testing.T) {
	attr := make(Attributes)
	attr.Set(AttrUserName, "username1")
	attr.Set(AttrUserName, "username2")
	attr.Set(AttrADSLAgentRemoteId, []byte("id1"))
	attr.Set(AttrADSLAgentRemoteId, []byte("id2"))
	if string(attr.GetRaw(AttrUserName)) != "username2" {
		t.Fail()
	}
	if len(attr[byte(AttrUserName)]) != 1 {
		t.Fail()
	}
	if string(attr.GetRaw(AttrADSLAgentRemoteId)) != "id2" {
		t.Fail()
	}
}

func TestAttributes_Get(t *testing.T) {
	attr := make(Attributes)
	attr.Set(AttrUserName, "username1")
	attr.Set(AttrADSLAgentRemoteId, []byte("id1"))
	if attr.Get(AttrUserName).(string) != "username1" {
		t.Fail()
	}
	if string(attr.Get(AttrADSLAgentRemoteId).([]byte)) != "id1" {
		t.Fail()
	}
}

func TestAttributes_GetString(t *testing.T) {
	attr := make(Attributes)
	attr.Set(AttrUserName, "username1")
	attr.Set(AttrADSLAgentRemoteId, []byte("id1"))
	attr.Set(AttrAcctDelayTime, 1800)
	if attr.GetString(AttrUserName) != "username1" {
		t.Fail()
	}
	if attr.GetString(AttrADSLAgentRemoteId) != "id1" {
		t.Fail()
	}
	if attr.GetString(AttrAcctDelayTime) != "1800" {
		t.Fail()
	}
}

func TestAttributes_SetByName_Error(t *testing.T) {
	attr := make(Attributes)
	if attr.SetByName("User-Name", 1232) == nil {
		t.Fail()
	}
	if attr.SetByName("Acct-Delay-Time", "test") == nil {
		t.Fail()
	}
}
