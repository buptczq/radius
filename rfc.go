package radius

const (
	AttrUserName                           AttributeKey = 1
	AttrUserPassword                       AttributeKey = 2
	AttrCHAPPassword                       AttributeKey = 3
	AttrNASIPAddress                       AttributeKey = 4
	AttrNASPort                            AttributeKey = 5
	AttrServiceType                        AttributeKey = 6
	AttrFramedProtocol                     AttributeKey = 7
	AttrFramedIPAddress                    AttributeKey = 8
	AttrFramedIPNetmask                    AttributeKey = 9
	AttrFramedRouting                      AttributeKey = 10
	AttrFilterId                           AttributeKey = 11
	AttrFramedMTU                          AttributeKey = 12
	AttrFramedCompression                  AttributeKey = 13
	AttrLoginIPHost                        AttributeKey = 14
	AttrLoginService                       AttributeKey = 15
	AttrLoginTCPPort                       AttributeKey = 16
	AttrReplyMessage                       AttributeKey = 18
	AttrCallbackNumber                     AttributeKey = 19
	AttrCallbackId                         AttributeKey = 20
	AttrFramedRoute                        AttributeKey = 22
	AttrFramedIPXNetwork                   AttributeKey = 23
	AttrState                              AttributeKey = 24
	AttrClass                              AttributeKey = 25
	AttrSessionTimeout                     AttributeKey = 27
	AttrIdleTimeout                        AttributeKey = 28
	AttrTerminationAction                  AttributeKey = 29
	AttrCalledStationId                    AttributeKey = 30
	AttrCallingStationId                   AttributeKey = 31
	AttrNASIdentifier                      AttributeKey = 32
	AttrProxyState                         AttributeKey = 33
	AttrLoginLATService                    AttributeKey = 34
	AttrLoginLATNode                       AttributeKey = 35
	AttrLoginLATGroup                      AttributeKey = 36
	AttrFramedAppleTalkLink                AttributeKey = 37
	AttrFramedAppleTalkNetwork             AttributeKey = 38
	AttrFramedAppleTalkZone                AttributeKey = 39
	AttrAcctStatusType                     AttributeKey = 40
	AttrAcctDelayTime                      AttributeKey = 41
	AttrAcctInputOctets                    AttributeKey = 42
	AttrAcctOutputOctets                   AttributeKey = 43
	AttrAcctSessionId                      AttributeKey = 44
	AttrAcctAuthentic                      AttributeKey = 45
	AttrAcctSessionTime                    AttributeKey = 46
	AttrAcctInputPackets                   AttributeKey = 47
	AttrAcctOutputPackets                  AttributeKey = 48
	AttrAcctTerminateCause                 AttributeKey = 49
	AttrAcctMultiSessionId                 AttributeKey = 50
	AttrAcctLinkCount                      AttributeKey = 51
	AttrAcctInputGigawords                 AttributeKey = 52
	AttrAcctOutputGigawords                AttributeKey = 53
	AttrEventTimestamp                     AttributeKey = 55
	AttrEgressVLANID                       AttributeKey = 56
	AttrIngressFilters                     AttributeKey = 57
	AttrEgressVLANName                     AttributeKey = 58
	AttrUserPriorityTable                  AttributeKey = 59
	AttrCHAPChallenge                      AttributeKey = 60
	AttrNASPortType                        AttributeKey = 61
	AttrPortLimit                          AttributeKey = 62
	AttrLoginLATPort                       AttributeKey = 63
	AttrTunnelType                         AttributeKey = 64
	AttrTunnelMediumType                   AttributeKey = 65
	AttrTunnelClientEndpoint               AttributeKey = 66
	AttrTunnelServerEndpoint               AttributeKey = 67
	AttrAcctTunnelConnection               AttributeKey = 68
	AttrTunnelPassword                     AttributeKey = 69
	AttrARAPPassword                       AttributeKey = 70
	AttrARAPFeatures                       AttributeKey = 71
	AttrARAPZoneAccess                     AttributeKey = 72
	AttrARAPSecurity                       AttributeKey = 73
	AttrARAPSecurityData                   AttributeKey = 74
	AttrPasswordRetry                      AttributeKey = 75
	AttrPrompt                             AttributeKey = 76
	AttrConnectInfo                        AttributeKey = 77
	AttrConfigurationToken                 AttributeKey = 78
	AttrMessageAuthenticator               AttributeKey = 80
	AttrTunnelPrivateGroupId               AttributeKey = 81
	AttrTunnelAssignmentId                 AttributeKey = 82
	AttrTunnelPreference                   AttributeKey = 83
	AttrARAPChallengeResponse              AttributeKey = 84
	AttrAcctInterimInterval                AttributeKey = 85
	AttrAcctTunnelPacketsLost              AttributeKey = 86
	AttrNASPortId                          AttributeKey = 87
	AttrFramedPool                         AttributeKey = 88
	AttrChargeableUserIdentity             AttributeKey = 89
	AttrTunnelClientAuthId                 AttributeKey = 90
	AttrTunnelServerAuthId                 AttributeKey = 91
	AttrNASFilterRule                      AttributeKey = 92
	AttrOriginatingLineInfo                AttributeKey = 94
	AttrNASIPv6Address                     AttributeKey = 95
	AttrFramedInterfaceId                  AttributeKey = 96
	AttrFramedIPv6Prefix                   AttributeKey = 97
	AttrLoginIPv6Host                      AttributeKey = 98
	AttrFramedIPv6Route                    AttributeKey = 99
	AttrFramedIPv6Pool                     AttributeKey = 100
	AttrErrorCause                         AttributeKey = 101
	AttrEAPKeyName                         AttributeKey = 102
	AttrDigestResponse                     AttributeKey = 103
	AttrDigestRealm                        AttributeKey = 104
	AttrDigestNonce                        AttributeKey = 105
	AttrDigestResponseAuth                 AttributeKey = 106
	AttrDigestNextnonce                    AttributeKey = 107
	AttrDigestMethod                       AttributeKey = 108
	AttrDigestURI                          AttributeKey = 109
	AttrDigestQop                          AttributeKey = 110
	AttrDigestAlgorithm                    AttributeKey = 111
	AttrDigestEntityBodyHash               AttributeKey = 112
	AttrDigestCNonce                       AttributeKey = 113
	AttrDigestNonceCount                   AttributeKey = 114
	AttrDigestUsername                     AttributeKey = 115
	AttrDigestOpaque                       AttributeKey = 116
	AttrDigestAuthParam                    AttributeKey = 117
	AttrDigestAKAAuts                      AttributeKey = 118
	AttrDigestDomain                       AttributeKey = 119
	AttrDigestStale                        AttributeKey = 120
	AttrDigestHA1                          AttributeKey = 121
	AttrSIPAOR                             AttributeKey = 122
	AttrDelegatedIPv6Prefix                AttributeKey = 123
	AttrMIP6FeatureVector                  AttributeKey = 124
	AttrMIP6HomeLinkPrefix                 AttributeKey = 125
	AttrOperatorName                       AttributeKey = 126
	AttrLocationInformation                AttributeKey = 127
	AttrLocationData                       AttributeKey = 128
	AttrBasicLocationPolicyRules           AttributeKey = 129
	AttrExtendedLocationPolicyRules        AttributeKey = 130
	AttrLocationCapable                    AttributeKey = 131
	AttrRequestedLocationInfo              AttributeKey = 132
	AttrFramedManagement                   AttributeKey = 133
	AttrManagementTransportProtection      AttributeKey = 134
	AttrManagementPolicyId                 AttributeKey = 135
	AttrManagementPrivilegeLevel           AttributeKey = 136
	AttrPKMConfigSettings                  AttributeKey = 139
	AttrPKMCryptosuiteList                 AttributeKey = 140
	AttrPKMSAID                            AttributeKey = 141
	AttrPKMSADescriptor                    AttributeKey = 142
	AttrPKMAuthKey                         AttributeKey = 143
	AttrDSLiteTunnelName                   AttributeKey = 144
	AttrMobileNodeIdentifier               AttributeKey = 145
	AttrServiceSelection                   AttributeKey = 146
	AttrPMIP6HomeLMAIPv6Address            AttributeKey = 147
	AttrPMIP6VisitedLMAIPv6Address         AttributeKey = 148
	AttrPMIP6HomeLMAIPv4Address            AttributeKey = 149
	AttrPMIP6VisitedLMAIPv4Address         AttributeKey = 150
	AttrPMIP6HomeHNPrefix                  AttributeKey = 151
	AttrPMIP6VisitedHNPrefix               AttributeKey = 152
	AttrPMIP6HomeInterfaceID               AttributeKey = 153
	AttrPMIP6VisitedInterfaceID            AttributeKey = 154
	AttrPMIP6HomeDHCP4ServerAddress        AttributeKey = 157
	AttrPMIP6VisitedDHCP4ServerAddress     AttributeKey = 158
	AttrPMIP6HomeDHCP6ServerAddress        AttributeKey = 159
	AttrPMIP6VisitedDHCP6ServerAddress     AttributeKey = 160
	AttrPMIP6HomeIPv4Gateway               AttributeKey = 161
	AttrPMIP6VisitedIPv4Gateway            AttributeKey = 162
	AttrEAPLowerLayer                      AttributeKey = 163
	AttrGSSAcceptorServiceName             AttributeKey = 164
	AttrGSSAcceptorHostName                AttributeKey = 165
	AttrGSSAcceptorServiceSpecifics        AttributeKey = 166
	AttrGSSAcceptorRealmName               AttributeKey = 167
	AttrFramedIPv6Address                  AttributeKey = 168
	AttrDNSServerIPv6Address               AttributeKey = 169
	AttrRouteIPv6Information               AttributeKey = 170
	AttrDelegatedIPv6PrefixPool            AttributeKey = 171
	AttrStatefulIPv6AddressPool            AttributeKey = 172
	AttrAllowedCalledStationId             AttributeKey = 174
	AttrEAPPeerId                          AttributeKey = 175
	AttrEAPServerId                        AttributeKey = 176
	AttrMobilityDomainId                   AttributeKey = 177
	AttrPreauthTimeout                     AttributeKey = 178
	AttrNetworkIdName                      AttributeKey = 179
	AttrWLANHESSID                         AttributeKey = 181
	AttrWLANVenueInfo                      AttributeKey = 182
	AttrWLANVenueLanguage                  AttributeKey = 183
	AttrWLANVenueName                      AttributeKey = 184
	AttrWLANReasonCode                     AttributeKey = 185
	AttrWLANPairwiseCipher                 AttributeKey = 186
	AttrWLANGroupCipher                    AttributeKey = 187
	AttrWLANAKMSuite                       AttributeKey = 188
	AttrWLANGroupMgmtCipher                AttributeKey = 189
	AttrWLANRFBand                         AttributeKey = 190
	AttrADSLAgentCircuitId                 AttributeKey = 15294378541057
	AttrADSLAgentRemoteId                  AttributeKey = 15294378541058
	AttrActualDataRateUpstream             AttributeKey = 15294378541185
	AttrActualDataRateDownstream           AttributeKey = 15294378541186
	AttrMinimumDataRateUpstream            AttributeKey = 15294378541187
	AttrMinimumDataRateDownstream          AttributeKey = 15294378541188
	AttrAttainableDataRateUpstream         AttributeKey = 15294378541189
	AttrAttainableDataRateDownstream       AttributeKey = 15294378541190
	AttrMaximumDataRateUpstream            AttributeKey = 15294378541191
	AttrMaximumDataRateDownstream          AttributeKey = 15294378541192
	AttrMinimumDataRateUpstreamLowPower    AttributeKey = 15294378541193
	AttrMinimumDataRateDownstreamLowPower  AttributeKey = 15294378541194
	AttrMaximumInterleavingDelayUpstream   AttributeKey = 15294378541195
	AttrActualInterleavingDelayUpstream    AttributeKey = 15294378541196
	AttrMaximumInterleavingDelayDownstream AttributeKey = 15294378541197
	AttrActualInterleavingDelayDownstream  AttributeKey = 15294378541198
	AttrAccessLoopEncapsulation            AttributeKey = 15294378541200
	AttrIWFSession                         AttributeKey = 15294378541310
	VendorADSLForum                        uint32       = 3561
)

func init() {
	builtinOnce.Do(initDictionary)
	Builtin.MustRegisterEx("User-Name", 0, 1, false, 0, AttributeString)
	Builtin.MustRegisterEx("User-Password", 0, 2, false, 1, AttributeString)
	Builtin.MustRegisterEx("CHAP-Password", 0, 3, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("NAS-IP-Address", 0, 4, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("NAS-Port", 0, 5, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Service-Type", 0, 6, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Framed-Protocol", 0, 7, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Framed-IP-Address", 0, 8, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("Framed-IP-Netmask", 0, 9, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("Framed-Routing", 0, 10, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Filter-Id", 0, 11, false, 0, AttributeString)
	Builtin.MustRegisterEx("Framed-MTU", 0, 12, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Framed-Compression", 0, 13, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Login-IP-Host", 0, 14, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("Login-Service", 0, 15, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Login-TCP-Port", 0, 16, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Reply-Message", 0, 18, false, 0, AttributeString)
	Builtin.MustRegisterEx("Callback-Number", 0, 19, false, 0, AttributeString)
	Builtin.MustRegisterEx("Callback-Id", 0, 20, false, 0, AttributeString)
	Builtin.MustRegisterEx("Framed-Route", 0, 22, false, 0, AttributeString)
	Builtin.MustRegisterEx("Framed-IPX-Network", 0, 23, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("State", 0, 24, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Class", 0, 25, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Session-Timeout", 0, 27, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Idle-Timeout", 0, 28, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Termination-Action", 0, 29, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Called-Station-Id", 0, 30, false, 0, AttributeString)
	Builtin.MustRegisterEx("Calling-Station-Id", 0, 31, false, 0, AttributeString)
	Builtin.MustRegisterEx("NAS-Identifier", 0, 32, false, 0, AttributeString)
	Builtin.MustRegisterEx("Proxy-State", 0, 33, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Login-LAT-Service", 0, 34, false, 0, AttributeString)
	Builtin.MustRegisterEx("Login-LAT-Node", 0, 35, false, 0, AttributeString)
	Builtin.MustRegisterEx("Login-LAT-Group", 0, 36, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Framed-AppleTalk-Link", 0, 37, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Framed-AppleTalk-Network", 0, 38, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Framed-AppleTalk-Zone", 0, 39, false, 0, AttributeString)
	Builtin.MustRegisterEx("Acct-Status-Type", 0, 40, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Delay-Time", 0, 41, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Input-Octets", 0, 42, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Output-Octets", 0, 43, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Session-Id", 0, 44, false, 0, AttributeString)
	Builtin.MustRegisterEx("Acct-Authentic", 0, 45, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Session-Time", 0, 46, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Input-Packets", 0, 47, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Output-Packets", 0, 48, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Terminate-Cause", 0, 49, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Multi-Session-Id", 0, 50, false, 0, AttributeString)
	Builtin.MustRegisterEx("Acct-Link-Count", 0, 51, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Input-Gigawords", 0, 52, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Output-Gigawords", 0, 53, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Event-Timestamp", 0, 55, false, 0, AttributeDate)
	Builtin.MustRegisterEx("Egress-VLANID", 0, 56, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Ingress-Filters", 0, 57, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Egress-VLAN-Name", 0, 58, false, 0, AttributeString)
	Builtin.MustRegisterEx("User-Priority-Table", 0, 59, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("CHAP-Challenge", 0, 60, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("NAS-Port-Type", 0, 61, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Port-Limit", 0, 62, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Login-LAT-Port", 0, 63, false, 0, AttributeString)
	Builtin.MustRegisterEx("Tunnel-Type", 0, 64, true, 0, AttributeInteger)
	Builtin.MustRegisterEx("Tunnel-Medium-Type", 0, 65, true, 0, AttributeInteger)
	Builtin.MustRegisterEx("Tunnel-Client-Endpoint", 0, 66, true, 0, AttributeString)
	Builtin.MustRegisterEx("Tunnel-Server-Endpoint", 0, 67, true, 0, AttributeString)
	Builtin.MustRegisterEx("Acct-Tunnel-Connection", 0, 68, false, 0, AttributeString)
	Builtin.MustRegisterEx("Tunnel-Password", 0, 69, true, 2, AttributeString)
	Builtin.MustRegisterEx("ARAP-Password", 0, 70, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("ARAP-Features", 0, 71, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("ARAP-Zone-Access", 0, 72, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("ARAP-Security", 0, 73, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("ARAP-Security-Data", 0, 74, false, 0, AttributeString)
	Builtin.MustRegisterEx("Password-Retry", 0, 75, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Prompt", 0, 76, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Connect-Info", 0, 77, false, 0, AttributeString)
	Builtin.MustRegisterEx("Configuration-Token", 0, 78, false, 0, AttributeString)
	Builtin.MustRegisterEx("Message-Authenticator", 0, 80, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Tunnel-Private-Group-Id", 0, 81, true, 0, AttributeString)
	Builtin.MustRegisterEx("Tunnel-Assignment-Id", 0, 82, true, 0, AttributeString)
	Builtin.MustRegisterEx("Tunnel-Preference", 0, 83, true, 0, AttributeInteger)
	Builtin.MustRegisterEx("ARAP-Challenge-Response", 0, 84, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Acct-Interim-Interval", 0, 85, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Acct-Tunnel-Packets-Lost", 0, 86, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("NAS-Port-Id", 0, 87, false, 0, AttributeString)
	Builtin.MustRegisterEx("Framed-Pool", 0, 88, false, 0, AttributeString)
	Builtin.MustRegisterEx("Chargeable-User-Identity", 0, 89, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Tunnel-Client-Auth-Id", 0, 90, true, 0, AttributeString)
	Builtin.MustRegisterEx("Tunnel-Server-Auth-Id", 0, 91, true, 0, AttributeString)
	Builtin.MustRegisterEx("NAS-Filter-Rule", 0, 92, false, 0, AttributeString)
	Builtin.MustRegisterEx("Originating-Line-Info", 0, 94, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("NAS-IPv6-Address", 0, 95, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("Framed-Interface-Id", 0, 96, false, 0, AttributeIFID)
	Builtin.MustRegisterEx("Framed-IPv6-Prefix", 0, 97, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegisterEx("Login-IPv6-Host", 0, 98, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("Framed-IPv6-Route", 0, 99, false, 0, AttributeString)
	Builtin.MustRegisterEx("Framed-IPv6-Pool", 0, 100, false, 0, AttributeString)
	Builtin.MustRegisterEx("Error-Cause", 0, 101, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("EAP-Key-Name", 0, 102, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Digest-Response", 0, 103, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Realm", 0, 104, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Nonce", 0, 105, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Response-Auth", 0, 106, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Nextnonce", 0, 107, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Method", 0, 108, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-URI", 0, 109, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Qop", 0, 110, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Algorithm", 0, 111, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Entity-Body-Hash", 0, 112, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-CNonce", 0, 113, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Nonce-Count", 0, 114, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Username", 0, 115, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Opaque", 0, 116, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Auth-Param", 0, 117, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-AKA-Auts", 0, 118, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Domain", 0, 119, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-Stale", 0, 120, false, 0, AttributeString)
	Builtin.MustRegisterEx("Digest-HA1", 0, 121, false, 0, AttributeString)
	Builtin.MustRegisterEx("SIP-AOR", 0, 122, false, 0, AttributeString)
	Builtin.MustRegisterEx("Delegated-IPv6-Prefix", 0, 123, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegisterEx("MIP6-Feature-Vector", 0, 124, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("MIP6-Home-Link-Prefix", 0, 125, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegisterEx("Operator-Name", 0, 126, false, 0, AttributeString)
	Builtin.MustRegisterEx("Location-Information", 0, 127, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Location-Data", 0, 128, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Basic-Location-Policy-Rules", 0, 129, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Extended-Location-Policy-Rules", 0, 130, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Location-Capable", 0, 131, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Requested-Location-Info", 0, 132, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Framed-Management", 0, 133, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Management-Transport-Protection", 0, 134, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Management-Policy-Id", 0, 135, false, 0, AttributeString)
	Builtin.MustRegisterEx("Management-Privilege-Level", 0, 136, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("PKM-Config-Settings", 0, 139, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("PKM-Cryptosuite-List", 0, 140, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("PKM-SAID", 0, 141, false, 0, AttributeUnknown)
	Builtin.MustRegisterEx("PKM-SA-Descriptor", 0, 142, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("PKM-Auth-Key", 0, 143, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("DS-Lite-Tunnel-Name", 0, 144, false, 0, AttributeString)
	Builtin.MustRegisterEx("Mobile-Node-Identifier", 0, 145, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Service-Selection", 0, 146, false, 0, AttributeString)
	Builtin.MustRegisterEx("PMIP6-Home-LMA-IPv6-Address", 0, 147, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("PMIP6-Visited-LMA-IPv6-Address", 0, 148, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("PMIP6-Home-LMA-IPv4-Address", 0, 149, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("PMIP6-Visited-LMA-IPv4-Address", 0, 150, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("PMIP6-Home-HN-Prefix", 0, 151, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegisterEx("PMIP6-Visited-HN-Prefix", 0, 152, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegisterEx("PMIP6-Home-Interface-ID", 0, 153, false, 0, AttributeIFID)
	Builtin.MustRegisterEx("PMIP6-Visited-Interface-ID", 0, 154, false, 0, AttributeIFID)
	Builtin.MustRegisterEx("PMIP6-Home-DHCP4-Server-Address", 0, 157, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("PMIP6-Visited-DHCP4-Server-Address", 0, 158, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("PMIP6-Home-DHCP6-Server-Address", 0, 159, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("PMIP6-Visited-DHCP6-Server-Address", 0, 160, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("PMIP6-Home-IPv4-Gateway", 0, 161, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("PMIP6-Visited-IPv4-Gateway", 0, 162, false, 0, AttributeIPAddr)
	Builtin.MustRegisterEx("EAP-Lower-Layer", 0, 163, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("GSS-Acceptor-Service-Name", 0, 164, false, 0, AttributeString)
	Builtin.MustRegisterEx("GSS-Acceptor-Host-Name", 0, 165, false, 0, AttributeString)
	Builtin.MustRegisterEx("GSS-Acceptor-Service-Specifics", 0, 166, false, 0, AttributeString)
	Builtin.MustRegisterEx("GSS-Acceptor-Realm-Name", 0, 167, false, 0, AttributeString)
	Builtin.MustRegisterEx("Framed-IPv6-Address", 0, 168, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("DNS-Server-IPv6-Address", 0, 169, false, 0, AttributeIPv6Addr)
	Builtin.MustRegisterEx("Route-IPv6-Information", 0, 170, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegisterEx("Delegated-IPv6-Prefix-Pool", 0, 171, false, 0, AttributeString)
	Builtin.MustRegisterEx("Stateful-IPv6-Address-Pool", 0, 172, false, 0, AttributeString)
	Builtin.MustRegisterEx("Allowed-Called-Station-Id", 0, 174, false, 0, AttributeString)
	Builtin.MustRegisterEx("EAP-Peer-Id", 0, 175, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("EAP-Server-Id", 0, 176, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Mobility-Domain-Id", 0, 177, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Preauth-Timeout", 0, 178, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Network-Id-Name", 0, 179, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("WLAN-HESSID", 0, 181, false, 0, AttributeString)
	Builtin.MustRegisterEx("WLAN-Venue-Info", 0, 182, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("WLAN-Venue-Language", 0, 183, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("WLAN-Venue-Name", 0, 184, false, 0, AttributeString)
	Builtin.MustRegisterEx("WLAN-Reason-Code", 0, 185, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("WLAN-Pairwise-Cipher", 0, 186, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("WLAN-Group-Cipher", 0, 187, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("WLAN-AKM-Suite", 0, 188, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("WLAN-Group-Mgmt-Cipher", 0, 189, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("WLAN-RF-Band", 0, 190, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("ADSL-Agent-Circuit-Id", 3561, 1, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("ADSL-Agent-Remote-Id", 3561, 2, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("Actual-Data-Rate-Upstream", 3561, 129, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Actual-Data-Rate-Downstream", 3561, 130, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Minimum-Data-Rate-Upstream", 3561, 131, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Minimum-Data-Rate-Downstream", 3561, 132, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Attainable-Data-Rate-Upstream", 3561, 133, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Attainable-Data-Rate-Downstream", 3561, 134, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Maximum-Data-Rate-Upstream", 3561, 135, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Maximum-Data-Rate-Downstream", 3561, 136, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Minimum-Data-Rate-Upstream-Low-Power", 3561, 137, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Minimum-Data-Rate-Downstream-Low-Power", 3561, 138, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Maximum-Interleaving-Delay-Upstream", 3561, 139, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Actual-Interleaving-Delay-Upstream", 3561, 140, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Maximum-Interleaving-Delay-Downstream", 3561, 141, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Actual-Interleaving-Delay-Downstream", 3561, 142, false, 0, AttributeInteger)
	Builtin.MustRegisterEx("Access-Loop-Encapsulation", 3561, 144, false, 0, AttributeOctets)
	Builtin.MustRegisterEx("IWF-Session", 3561, 254, false, 0, AttributeOctets)
}
