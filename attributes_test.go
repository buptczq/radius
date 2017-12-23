package radius

import (
	"testing"
)

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
	if attr.GetString(AttrADSLAgentRemoteId) != "0x696431" {
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

func TestAttributes_Tag(t *testing.T) {
	attr := make(Attributes)
	attr.AddRaw(AttrTunnelPrivateGroupId.WithTag(1), []byte{66})
	attr.AddRaw(AttrTunnelPrivateGroupId.WithTag(2), []byte{66})
	if len(attr[AttrTunnelPrivateGroupId.Type()]) != 2 {
		t.Fail()
	}
	attr.Del(AttrTunnelPrivateGroupId)
	if len(attr[AttrTunnelPrivateGroupId.Type()]) != 0 {
		t.Fail()
	}
	attr.AddRaw(AttrTunnelPrivateGroupId.WithTag(1), []byte{66})
	attr.AddRaw(AttrTunnelPrivateGroupId.WithTag(2), []byte{66})
	attr.Del(AttrTunnelPrivateGroupId.WithTag(1))
	if len(attr[AttrTunnelPrivateGroupId.Type()]) != 1 {
		t.Fail()
	}
}

func TestAttributes_AVPTag(t *testing.T) {
	Builtin.RegisterEx("Alc-Tunnel-Max-Sessions", 6527, 48, true, 0, AttributeInteger)
	key := MakeAttributeKey(6527, 0, 48)
	attr := make(Attributes)
	attr.AddRaw(key.WithTag(1), []byte{66})
	attr.AddRaw(key.WithTag(2), []byte{66})
	if len(attr[VENDOR_SPECIFIC]) != 2 {
		t.Fail()
	}
	attr.Del(key)
	if len(attr[VENDOR_SPECIFIC]) != 0 {
		t.Fail()
	}
	attr.AddRaw(key.WithTag(1), []byte{66})
	attr.AddRaw(key.WithTag(2), []byte{66})
	attr.Del(key.WithTag(1))
	if len(attr[VENDOR_SPECIFIC]) != 1 {
		t.Fail()
	}
}

func TestAttributes_Tag_2(t *testing.T) {
	attr := make(Attributes)
	attr.Set(AttrTunnelPrivateGroupId.WithTag(1), []byte("test1"))
	attr.Set(AttrTunnelPrivateGroupId.WithTag(2), []byte("test2"))
	if len(attr[AttrTunnelPrivateGroupId.Type()]) != 2 {
		t.Fail()
	}
	if attr.Get(AttrTunnelPrivateGroupId.WithTag(1)).(string) != "test1" {
		t.Fail()
	}
	if attr.Get(AttrTunnelPrivateGroupId.WithTag(2)).(string) != "test2" {
		t.Fail()
	}
}

func TestAttributes_AVPTag_2(t *testing.T) {
	Builtin.RegisterEx("Alc-Tunnel-Max-Sessions", 6527, 48, true, 0, AttributeInteger)
	key := MakeAttributeKey(6527, 0, 48)
	attr := make(Attributes)
	attr.Set(key.WithTag(1), 1)
	attr.Set(key.WithTag(2), 2)
	if attr.Get(key.WithTag(1)).(uint32) != 1 {
		t.Fail()
	}
	if attr.Get(key.WithTag(2)).(uint32) != 2 {
		t.Fail()
	}
	if attr.Get(key.WithTag(3)) != nil {
		t.Fail()
	}
}
