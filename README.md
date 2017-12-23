# radius
a Go (golang) RADIUS client and server implementation

# The datatypes currently supported are:

|type       | description
|-----------|-------------------
|string     | ASCII string
|ipaddr     | IPv4 address
|date       | 32 bits UNIX
|octets     | arbitrary binary data
|ipv6addr   | 16 octets in network byte order
|integer    | 32 bits unsigned number
|signed     | 32 bits signed number
|short      | 16 bits unsigned number
|byte       | 8 bits unsigned number

# Example

## Parse packet

```go
p, err := radius.Parse(request, secret)
if err != nil {
	t.Fatal(err)
}
// Get a attribute by name
if p.GetByName("User-Name").(string) != "nemo" {
	t.Fatal("expecting User-Name = nemo")
}
// Get a attribute by type
if p.Get(radius.AttrUserName).(string) != "nemo" {
	t.Fatal("expecting User-Name = nemo")
}

```

## Build attributes with tags

```go
attr := make(Attributes)
attr.Set(AttrTunnelPrivateGroupId.WithTag(1), []byte("test1"))
attr.Set(AttrTunnelPrivateGroupId.WithTag(2), []byte("test2"))
if attr.Get(AttrTunnelPrivateGroupId.WithTag(1)).(string) != "test1" {
	t.Fail()
}
if attr.Get(AttrTunnelPrivateGroupId.WithTag(2)).(string) != "test2" {
	t.Fail()
}
```

# Builtin Dictionaries

* dictionary.rfc2865
* dictionary.rfc2866
* dictionary.rfc2867
* dictionary.rfc2868
* dictionary.rfc2869
* dictionary.rfc3162
* dictionary.rfc3576
* dictionary.rfc3580
* dictionary.rfc4072
* dictionary.rfc4372
* dictionary.rfc4603
* dictionary.rfc4675
* dictionary.rfc4679 (partial)
* dictionary.rfc4818
* dictionary.rfc4849
* dictionary.rfc5090
* dictionary.rfc5176
* dictionary.rfc5447
* dictionary.rfc5580
* dictionary.rfc5607
* dictionary.rfc5904
* dictionary.rfc6519
* dictionary.rfc6572 (partial)
* dictionary.rfc6677
* dictionary.rfc6911
* dictionary.rfc6930 (partial)
* dictionary.rfc7055
* dictionary.rfc7155
* dictionary.rfc7268
* dictionary.rfc7930
