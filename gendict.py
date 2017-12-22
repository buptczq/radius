from pyrad import dictionary
import sys
import os



attrMap = {
    "string"     : "AttributeString",
    "octets"     : "AttributeOctets",
    "ipaddr"     : "AttributeIPAddr",
    "date"       : "AttributeDate",
    "integer"    : "AttributeInteger",
    "ipv6addr"   : "AttributeIPv6Addr",
    "ipv6prefix" : "AttributeIPv6Prefix",
    "ifid"       : "AttributeIFID",
    "integer64"  : "AttributeInteger64",
}

fileTmpl = '''
package radius


const (
%s
)

func init() {
	builtinOnce.Do(initDictionary)
%s
}
'''

tmpl = '''Builtin.MustRegister("{name}", {code}, {type})'''
tmplVendor = '''Builtin.MustRegisterVendor("{name}", {oid}, {code}, {type})'''

output = []
consts = []

dicts = []
for _, _, files in os.walk('.'):
    for filename in  files:
        if filename.startswith('dictionary.'):
            if filename.find('rfc')>0:
                dicts.append(filename)
d = dictionary.Dictionary(*dicts)
for k in d.attributes:
    attr = d.attributes[k]
    oid = d.vendors[attr.vendor]
    # print attr.vendor
    # for vname in attr.values.forward:
        # print vname

    newName = attr.name.replace('-','') + '_Type'
    if oid == 0:
        output.append((attr.code, tmpl.format(
                name = attr.name,
                code = newName,
                type = attrMap.get(attr.type, 'AttributeUnknown')
                ), "%s byte = %s" % (newName, attr.code) ))
    else:
        output.append((attr.code, tmplVendor.format(
                name = attr.name,
                oid = oid,
                code = attr.code,
                type = attrMap.get(attr.type, 'AttributeUnknown')
                ), ""))


output = sorted(output, key=lambda x:x[0])
print "package radius\n\n"
print "const ("
for i in output:
    if i[2]:
        print '\t',i[2]
print ")\n"

print "func init() {"
print "\tbuiltinOnce.Do(initDictionary)"
for i in output:
    print '\t',i[1]
print "}"
