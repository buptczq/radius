from pyrad import dictionary
from pyrad import tools
import os

attrMap = {
    "string"     : "AttributeString",
    "octets"     : "AttributeOctets",
    "ipaddr"     : "AttributeIPAddr",
    "date"       : "AttributeDate",
    "integer"    : "AttributeInteger",
    "signed"     : "AttributeSigned",
    "short"      : "AttributeShort",
    "ipv6addr"   : "AttributeIPv6Addr",
    "ipv6prefix" : "AttributeIPv6Prefix",
    "ifid"       : "AttributeIFID",
    "integer64"  : "AttributeInteger64",
}

tmpl = '''Builtin.MustRegister("{name}", {oid}, {code}, {tag}, {encrypt}, {type})'''

output = []
const_attr = []
const_vendor = []
const_value = []

def makeKey(oid, code):
    return oid << 32 | code

dicts = []
for _, _, files in os.walk('.'):
    for filename in  files:
        if filename.startswith('dictionary.'):
            if filename.find('rfc')>0:
                dicts.append(filename)
d = dictionary.Dictionary(*dicts)

for oid, name in d.vendors.backward.items():
    if oid != 0:
        const_vendor.append( (oid , name) )

for k in d.attributes:
    attr = d.attributes[k]
    oid = d.vendors[attr.vendor]
    key = makeKey(oid, attr.code)

    if attr.type == 'integer':
        for vname, value in attr.values.forward.items():
            newName = attr.name.replace('-','') + '_' + vname.replace('-','').replace('.','')
            const_value.append(( key, tools.DecodeInteger(value), newName))

    newName = 'Attr' + attr.name.replace('-','')
    output.append((key, tmpl.format(
            name = attr.name,
            oid = oid,
            code = attr.code,
            tag = 'true' if attr.has_tag else 'false',
            encrypt = attr.encrypt,
            type = attrMap.get(attr.type, 'AttributeUnknown')
            ), "%s AttributeKey = %s" % (newName, key) ))


output = sorted(output, key=lambda x:x[0])
const_vendor = sorted(const_vendor, key=lambda x:x[0])
const_value = sorted(const_value, key=lambda x:(x[0],x[1]))
print "package radius\n\n"
print "const ("
for i in output:
    if i[2]:
        print '\t',i[2]
for i in const_vendor:
    print '\tVendor%s uint32 = %s' % (i[1].replace('-', ''), i[0])
for i in const_value:
    print '\t%s uint32 = %s' % (i[2], i[1])
print ")\n"

print "func init() {"
print "\tbuiltinOnce.Do(initDictionary)"
for i in output:
    print '\t',i[1]
print "}"
