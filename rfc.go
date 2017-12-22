package radius

const (
	UserName_Type                       byte = 1
	UserPassword_Type                   byte = 2
	CHAPPassword_Type                   byte = 3
	NASIPAddress_Type                   byte = 4
	NASPort_Type                        byte = 5
	ServiceType_Type                    byte = 6
	FramedProtocol_Type                 byte = 7
	FramedIPAddress_Type                byte = 8
	FramedIPNetmask_Type                byte = 9
	FramedRouting_Type                  byte = 10
	FilterId_Type                       byte = 11
	FramedMTU_Type                      byte = 12
	FramedCompression_Type              byte = 13
	LoginIPHost_Type                    byte = 14
	LoginService_Type                   byte = 15
	LoginTCPPort_Type                   byte = 16
	ReplyMessage_Type                   byte = 18
	CallbackNumber_Type                 byte = 19
	CallbackId_Type                     byte = 20
	FramedRoute_Type                    byte = 22
	FramedIPXNetwork_Type               byte = 23
	State_Type                          byte = 24
	Class_Type                          byte = 25
	SessionTimeout_Type                 byte = 27
	IdleTimeout_Type                    byte = 28
	TerminationAction_Type              byte = 29
	CalledStationId_Type                byte = 30
	CallingStationId_Type               byte = 31
	NASIdentifier_Type                  byte = 32
	ProxyState_Type                     byte = 33
	LoginLATService_Type                byte = 34
	LoginLATNode_Type                   byte = 35
	LoginLATGroup_Type                  byte = 36
	FramedAppleTalkLink_Type            byte = 37
	FramedAppleTalkNetwork_Type         byte = 38
	FramedAppleTalkZone_Type            byte = 39
	AcctStatusType_Type                 byte = 40
	AcctDelayTime_Type                  byte = 41
	AcctInputOctets_Type                byte = 42
	AcctOutputOctets_Type               byte = 43
	AcctSessionId_Type                  byte = 44
	AcctAuthentic_Type                  byte = 45
	AcctSessionTime_Type                byte = 46
	AcctInputPackets_Type               byte = 47
	AcctOutputPackets_Type              byte = 48
	AcctTerminateCause_Type             byte = 49
	AcctMultiSessionId_Type             byte = 50
	AcctLinkCount_Type                  byte = 51
	AcctInputGigawords_Type             byte = 52
	AcctOutputGigawords_Type            byte = 53
	EventTimestamp_Type                 byte = 55
	EgressVLANID_Type                   byte = 56
	IngressFilters_Type                 byte = 57
	EgressVLANName_Type                 byte = 58
	UserPriorityTable_Type              byte = 59
	CHAPChallenge_Type                  byte = 60
	NASPortType_Type                    byte = 61
	PortLimit_Type                      byte = 62
	LoginLATPort_Type                   byte = 63
	TunnelType_Type                     byte = 64
	TunnelMediumType_Type               byte = 65
	TunnelClientEndpoint_Type           byte = 66
	TunnelServerEndpoint_Type           byte = 67
	AcctTunnelConnection_Type           byte = 68
	TunnelPassword_Type                 byte = 69
	ARAPPassword_Type                   byte = 70
	ARAPFeatures_Type                   byte = 71
	ARAPZoneAccess_Type                 byte = 72
	ARAPSecurity_Type                   byte = 73
	ARAPSecurityData_Type               byte = 74
	PasswordRetry_Type                  byte = 75
	Prompt_Type                         byte = 76
	ConnectInfo_Type                    byte = 77
	ConfigurationToken_Type             byte = 78
	MessageAuthenticator_Type           byte = 80
	TunnelPrivateGroupId_Type           byte = 81
	TunnelAssignmentId_Type             byte = 82
	TunnelPreference_Type               byte = 83
	ARAPChallengeResponse_Type          byte = 84
	AcctInterimInterval_Type            byte = 85
	AcctTunnelPacketsLost_Type          byte = 86
	NASPortId_Type                      byte = 87
	FramedPool_Type                     byte = 88
	ChargeableUserIdentity_Type         byte = 89
	TunnelClientAuthId_Type             byte = 90
	TunnelServerAuthId_Type             byte = 91
	NASFilterRule_Type                  byte = 92
	OriginatingLineInfo_Type            byte = 94
	NASIPv6Address_Type                 byte = 95
	FramedInterfaceId_Type              byte = 96
	FramedIPv6Prefix_Type               byte = 97
	LoginIPv6Host_Type                  byte = 98
	FramedIPv6Route_Type                byte = 99
	FramedIPv6Pool_Type                 byte = 100
	ErrorCause_Type                     byte = 101
	EAPKeyName_Type                     byte = 102
	DigestResponse_Type                 byte = 103
	DigestRealm_Type                    byte = 104
	DigestNonce_Type                    byte = 105
	DigestResponseAuth_Type             byte = 106
	DigestNextnonce_Type                byte = 107
	DigestMethod_Type                   byte = 108
	DigestURI_Type                      byte = 109
	DigestQop_Type                      byte = 110
	DigestAlgorithm_Type                byte = 111
	DigestEntityBodyHash_Type           byte = 112
	DigestCNonce_Type                   byte = 113
	DigestNonceCount_Type               byte = 114
	DigestUsername_Type                 byte = 115
	DigestOpaque_Type                   byte = 116
	DigestAuthParam_Type                byte = 117
	DigestAKAAuts_Type                  byte = 118
	DigestDomain_Type                   byte = 119
	DigestStale_Type                    byte = 120
	DigestHA1_Type                      byte = 121
	SIPAOR_Type                         byte = 122
	DelegatedIPv6Prefix_Type            byte = 123
	MIP6FeatureVector_Type              byte = 124
	MIP6HomeLinkPrefix_Type             byte = 125
	OperatorName_Type                   byte = 126
	LocationInformation_Type            byte = 127
	LocationData_Type                   byte = 128
	BasicLocationPolicyRules_Type       byte = 129
	ExtendedLocationPolicyRules_Type    byte = 130
	LocationCapable_Type                byte = 131
	RequestedLocationInfo_Type          byte = 132
	FramedManagement_Type               byte = 133
	ManagementTransportProtection_Type  byte = 134
	ManagementPolicyId_Type             byte = 135
	ManagementPrivilegeLevel_Type       byte = 136
	PKMConfigSettings_Type              byte = 139
	PKMCryptosuiteList_Type             byte = 140
	PKMSAID_Type                        byte = 141
	PKMSADescriptor_Type                byte = 142
	PKMAuthKey_Type                     byte = 143
	DSLiteTunnelName_Type               byte = 144
	MobileNodeIdentifier_Type           byte = 145
	ServiceSelection_Type               byte = 146
	PMIP6HomeLMAIPv6Address_Type        byte = 147
	PMIP6VisitedLMAIPv6Address_Type     byte = 148
	PMIP6HomeLMAIPv4Address_Type        byte = 149
	PMIP6VisitedLMAIPv4Address_Type     byte = 150
	PMIP6HomeHNPrefix_Type              byte = 151
	PMIP6VisitedHNPrefix_Type           byte = 152
	PMIP6HomeInterfaceID_Type           byte = 153
	PMIP6VisitedInterfaceID_Type        byte = 154
	PMIP6HomeDHCP4ServerAddress_Type    byte = 157
	PMIP6VisitedDHCP4ServerAddress_Type byte = 158
	PMIP6HomeDHCP6ServerAddress_Type    byte = 159
	PMIP6VisitedDHCP6ServerAddress_Type byte = 160
	PMIP6HomeIPv4Gateway_Type           byte = 161
	PMIP6VisitedIPv4Gateway_Type        byte = 162
	EAPLowerLayer_Type                  byte = 163
	GSSAcceptorServiceName_Type         byte = 164
	GSSAcceptorHostName_Type            byte = 165
	GSSAcceptorServiceSpecifics_Type    byte = 166
	GSSAcceptorRealmName_Type           byte = 167
	FramedIPv6Address_Type              byte = 168
	DNSServerIPv6Address_Type           byte = 169
	RouteIPv6Information_Type           byte = 170
	DelegatedIPv6PrefixPool_Type        byte = 171
	StatefulIPv6AddressPool_Type        byte = 172
	AllowedCalledStationId_Type         byte = 174
	EAPPeerId_Type                      byte = 175
	EAPServerId_Type                    byte = 176
	MobilityDomainId_Type               byte = 177
	PreauthTimeout_Type                 byte = 178
	NetworkIdName_Type                  byte = 179
	WLANHESSID_Type                     byte = 181
	WLANVenueInfo_Type                  byte = 182
	WLANVenueLanguage_Type              byte = 183
	WLANVenueName_Type                  byte = 184
	WLANReasonCode_Type                 byte = 185
	WLANPairwiseCipher_Type             byte = 186
	WLANGroupCipher_Type                byte = 187
	WLANAKMSuite_Type                   byte = 188
	WLANGroupMgmtCipher_Type            byte = 189
	WLANRFBand_Type                     byte = 190
)

func init() {
	builtinOnce.Do(initDictionary)
	Builtin.MustRegister("User-Name", UserName_Type, AttributeString)
	Builtin.MustRegisterVendor("ADSL-Agent-Circuit-Id", 3561, 1, AttributeOctets)
	Builtin.MustRegister("User-Password", UserPassword_Type, AttributeString)
	Builtin.MustRegisterVendor("ADSL-Agent-Remote-Id", 3561, 2, AttributeOctets)
	Builtin.MustRegister("CHAP-Password", CHAPPassword_Type, AttributeOctets)
	Builtin.MustRegister("NAS-IP-Address", NASIPAddress_Type, AttributeIPAddr)
	Builtin.MustRegister("NAS-Port", NASPort_Type, AttributeInteger)
	Builtin.MustRegister("Service-Type", ServiceType_Type, AttributeInteger)
	Builtin.MustRegister("Framed-Protocol", FramedProtocol_Type, AttributeInteger)
	Builtin.MustRegister("Framed-IP-Address", FramedIPAddress_Type, AttributeIPAddr)
	Builtin.MustRegister("Framed-IP-Netmask", FramedIPNetmask_Type, AttributeIPAddr)
	Builtin.MustRegister("Framed-Routing", FramedRouting_Type, AttributeInteger)
	Builtin.MustRegister("Filter-Id", FilterId_Type, AttributeString)
	Builtin.MustRegister("Framed-MTU", FramedMTU_Type, AttributeInteger)
	Builtin.MustRegister("Framed-Compression", FramedCompression_Type, AttributeInteger)
	Builtin.MustRegister("Login-IP-Host", LoginIPHost_Type, AttributeIPAddr)
	Builtin.MustRegister("Login-Service", LoginService_Type, AttributeInteger)
	Builtin.MustRegister("Login-TCP-Port", LoginTCPPort_Type, AttributeInteger)
	Builtin.MustRegister("Reply-Message", ReplyMessage_Type, AttributeString)
	Builtin.MustRegister("Callback-Number", CallbackNumber_Type, AttributeString)
	Builtin.MustRegister("Callback-Id", CallbackId_Type, AttributeString)
	Builtin.MustRegister("Framed-Route", FramedRoute_Type, AttributeString)
	Builtin.MustRegister("Framed-IPX-Network", FramedIPXNetwork_Type, AttributeIPAddr)
	Builtin.MustRegister("State", State_Type, AttributeOctets)
	Builtin.MustRegister("Class", Class_Type, AttributeOctets)
	Builtin.MustRegister("Session-Timeout", SessionTimeout_Type, AttributeInteger)
	Builtin.MustRegister("Idle-Timeout", IdleTimeout_Type, AttributeInteger)
	Builtin.MustRegister("Termination-Action", TerminationAction_Type, AttributeInteger)
	Builtin.MustRegister("Called-Station-Id", CalledStationId_Type, AttributeString)
	Builtin.MustRegister("Calling-Station-Id", CallingStationId_Type, AttributeString)
	Builtin.MustRegister("NAS-Identifier", NASIdentifier_Type, AttributeString)
	Builtin.MustRegister("Proxy-State", ProxyState_Type, AttributeOctets)
	Builtin.MustRegister("Login-LAT-Service", LoginLATService_Type, AttributeString)
	Builtin.MustRegister("Login-LAT-Node", LoginLATNode_Type, AttributeString)
	Builtin.MustRegister("Login-LAT-Group", LoginLATGroup_Type, AttributeOctets)
	Builtin.MustRegister("Framed-AppleTalk-Link", FramedAppleTalkLink_Type, AttributeInteger)
	Builtin.MustRegister("Framed-AppleTalk-Network", FramedAppleTalkNetwork_Type, AttributeInteger)
	Builtin.MustRegister("Framed-AppleTalk-Zone", FramedAppleTalkZone_Type, AttributeString)
	Builtin.MustRegister("Acct-Status-Type", AcctStatusType_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Delay-Time", AcctDelayTime_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Input-Octets", AcctInputOctets_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Output-Octets", AcctOutputOctets_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Session-Id", AcctSessionId_Type, AttributeString)
	Builtin.MustRegister("Acct-Authentic", AcctAuthentic_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Session-Time", AcctSessionTime_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Input-Packets", AcctInputPackets_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Output-Packets", AcctOutputPackets_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Terminate-Cause", AcctTerminateCause_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Multi-Session-Id", AcctMultiSessionId_Type, AttributeString)
	Builtin.MustRegister("Acct-Link-Count", AcctLinkCount_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Input-Gigawords", AcctInputGigawords_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Output-Gigawords", AcctOutputGigawords_Type, AttributeInteger)
	Builtin.MustRegister("Event-Timestamp", EventTimestamp_Type, AttributeDate)
	Builtin.MustRegister("Egress-VLANID", EgressVLANID_Type, AttributeInteger)
	Builtin.MustRegister("Ingress-Filters", IngressFilters_Type, AttributeInteger)
	Builtin.MustRegister("Egress-VLAN-Name", EgressVLANName_Type, AttributeString)
	Builtin.MustRegister("User-Priority-Table", UserPriorityTable_Type, AttributeOctets)
	Builtin.MustRegister("CHAP-Challenge", CHAPChallenge_Type, AttributeOctets)
	Builtin.MustRegister("NAS-Port-Type", NASPortType_Type, AttributeInteger)
	Builtin.MustRegister("Port-Limit", PortLimit_Type, AttributeInteger)
	Builtin.MustRegister("Login-LAT-Port", LoginLATPort_Type, AttributeString)
	Builtin.MustRegister("Tunnel-Type", TunnelType_Type, AttributeInteger)
	Builtin.MustRegister("Tunnel-Medium-Type", TunnelMediumType_Type, AttributeInteger)
	Builtin.MustRegister("Tunnel-Client-Endpoint", TunnelClientEndpoint_Type, AttributeString)
	Builtin.MustRegister("Tunnel-Server-Endpoint", TunnelServerEndpoint_Type, AttributeString)
	Builtin.MustRegister("Acct-Tunnel-Connection", AcctTunnelConnection_Type, AttributeString)
	Builtin.MustRegister("Tunnel-Password", TunnelPassword_Type, AttributeString)
	Builtin.MustRegister("ARAP-Password", ARAPPassword_Type, AttributeOctets)
	Builtin.MustRegister("ARAP-Features", ARAPFeatures_Type, AttributeOctets)
	Builtin.MustRegister("ARAP-Zone-Access", ARAPZoneAccess_Type, AttributeInteger)
	Builtin.MustRegister("ARAP-Security", ARAPSecurity_Type, AttributeInteger)
	Builtin.MustRegister("ARAP-Security-Data", ARAPSecurityData_Type, AttributeString)
	Builtin.MustRegister("Password-Retry", PasswordRetry_Type, AttributeInteger)
	Builtin.MustRegister("Prompt", Prompt_Type, AttributeInteger)
	Builtin.MustRegister("Connect-Info", ConnectInfo_Type, AttributeString)
	Builtin.MustRegister("Configuration-Token", ConfigurationToken_Type, AttributeString)
	Builtin.MustRegister("Message-Authenticator", MessageAuthenticator_Type, AttributeOctets)
	Builtin.MustRegister("Tunnel-Private-Group-Id", TunnelPrivateGroupId_Type, AttributeString)
	Builtin.MustRegister("Tunnel-Assignment-Id", TunnelAssignmentId_Type, AttributeString)
	Builtin.MustRegister("Tunnel-Preference", TunnelPreference_Type, AttributeInteger)
	Builtin.MustRegister("ARAP-Challenge-Response", ARAPChallengeResponse_Type, AttributeOctets)
	Builtin.MustRegister("Acct-Interim-Interval", AcctInterimInterval_Type, AttributeInteger)
	Builtin.MustRegister("Acct-Tunnel-Packets-Lost", AcctTunnelPacketsLost_Type, AttributeInteger)
	Builtin.MustRegister("NAS-Port-Id", NASPortId_Type, AttributeString)
	Builtin.MustRegister("Framed-Pool", FramedPool_Type, AttributeString)
	Builtin.MustRegister("Chargeable-User-Identity", ChargeableUserIdentity_Type, AttributeOctets)
	Builtin.MustRegister("Tunnel-Client-Auth-Id", TunnelClientAuthId_Type, AttributeString)
	Builtin.MustRegister("Tunnel-Server-Auth-Id", TunnelServerAuthId_Type, AttributeString)
	Builtin.MustRegister("NAS-Filter-Rule", NASFilterRule_Type, AttributeString)
	Builtin.MustRegister("Originating-Line-Info", OriginatingLineInfo_Type, AttributeOctets)
	Builtin.MustRegister("NAS-IPv6-Address", NASIPv6Address_Type, AttributeIPv6Addr)
	Builtin.MustRegister("Framed-Interface-Id", FramedInterfaceId_Type, AttributeIFID)
	Builtin.MustRegister("Framed-IPv6-Prefix", FramedIPv6Prefix_Type, AttributeIPv6Prefix)
	Builtin.MustRegister("Login-IPv6-Host", LoginIPv6Host_Type, AttributeIPv6Addr)
	Builtin.MustRegister("Framed-IPv6-Route", FramedIPv6Route_Type, AttributeString)
	Builtin.MustRegister("Framed-IPv6-Pool", FramedIPv6Pool_Type, AttributeString)
	Builtin.MustRegister("Error-Cause", ErrorCause_Type, AttributeInteger)
	Builtin.MustRegister("EAP-Key-Name", EAPKeyName_Type, AttributeOctets)
	Builtin.MustRegister("Digest-Response", DigestResponse_Type, AttributeString)
	Builtin.MustRegister("Digest-Realm", DigestRealm_Type, AttributeString)
	Builtin.MustRegister("Digest-Nonce", DigestNonce_Type, AttributeString)
	Builtin.MustRegister("Digest-Response-Auth", DigestResponseAuth_Type, AttributeString)
	Builtin.MustRegister("Digest-Nextnonce", DigestNextnonce_Type, AttributeString)
	Builtin.MustRegister("Digest-Method", DigestMethod_Type, AttributeString)
	Builtin.MustRegister("Digest-URI", DigestURI_Type, AttributeString)
	Builtin.MustRegister("Digest-Qop", DigestQop_Type, AttributeString)
	Builtin.MustRegister("Digest-Algorithm", DigestAlgorithm_Type, AttributeString)
	Builtin.MustRegister("Digest-Entity-Body-Hash", DigestEntityBodyHash_Type, AttributeString)
	Builtin.MustRegister("Digest-CNonce", DigestCNonce_Type, AttributeString)
	Builtin.MustRegister("Digest-Nonce-Count", DigestNonceCount_Type, AttributeString)
	Builtin.MustRegister("Digest-Username", DigestUsername_Type, AttributeString)
	Builtin.MustRegister("Digest-Opaque", DigestOpaque_Type, AttributeString)
	Builtin.MustRegister("Digest-Auth-Param", DigestAuthParam_Type, AttributeString)
	Builtin.MustRegister("Digest-AKA-Auts", DigestAKAAuts_Type, AttributeString)
	Builtin.MustRegister("Digest-Domain", DigestDomain_Type, AttributeString)
	Builtin.MustRegister("Digest-Stale", DigestStale_Type, AttributeString)
	Builtin.MustRegister("Digest-HA1", DigestHA1_Type, AttributeString)
	Builtin.MustRegister("SIP-AOR", SIPAOR_Type, AttributeString)
	Builtin.MustRegister("Delegated-IPv6-Prefix", DelegatedIPv6Prefix_Type, AttributeIPv6Prefix)
	Builtin.MustRegister("MIP6-Feature-Vector", MIP6FeatureVector_Type, AttributeOctets)
	Builtin.MustRegister("MIP6-Home-Link-Prefix", MIP6HomeLinkPrefix_Type, AttributeIPv6Prefix)
	Builtin.MustRegister("Operator-Name", OperatorName_Type, AttributeString)
	Builtin.MustRegister("Location-Information", LocationInformation_Type, AttributeOctets)
	Builtin.MustRegister("Location-Data", LocationData_Type, AttributeOctets)
	Builtin.MustRegisterVendor("Actual-Data-Rate-Upstream", 3561, 129, AttributeInteger)
	Builtin.MustRegister("Basic-Location-Policy-Rules", BasicLocationPolicyRules_Type, AttributeOctets)
	Builtin.MustRegisterVendor("Actual-Data-Rate-Downstream", 3561, 130, AttributeInteger)
	Builtin.MustRegister("Extended-Location-Policy-Rules", ExtendedLocationPolicyRules_Type, AttributeOctets)
	Builtin.MustRegisterVendor("Minimum-Data-Rate-Upstream", 3561, 131, AttributeInteger)
	Builtin.MustRegister("Location-Capable", LocationCapable_Type, AttributeInteger)
	Builtin.MustRegister("Requested-Location-Info", RequestedLocationInfo_Type, AttributeInteger)
	Builtin.MustRegisterVendor("Minimum-Data-Rate-Downstream", 3561, 132, AttributeInteger)
	Builtin.MustRegister("Framed-Management", FramedManagement_Type, AttributeInteger)
	Builtin.MustRegisterVendor("Attainable-Data-Rate-Upstream", 3561, 133, AttributeInteger)
	Builtin.MustRegister("Management-Transport-Protection", ManagementTransportProtection_Type, AttributeInteger)
	Builtin.MustRegisterVendor("Attainable-Data-Rate-Downstream", 3561, 134, AttributeInteger)
	Builtin.MustRegisterVendor("Maximum-Data-Rate-Upstream", 3561, 135, AttributeInteger)
	Builtin.MustRegister("Management-Policy-Id", ManagementPolicyId_Type, AttributeString)
	Builtin.MustRegisterVendor("Maximum-Data-Rate-Downstream", 3561, 136, AttributeInteger)
	Builtin.MustRegister("Management-Privilege-Level", ManagementPrivilegeLevel_Type, AttributeInteger)
	Builtin.MustRegisterVendor("Minimum-Data-Rate-Upstream-Low-Power", 3561, 137, AttributeInteger)
	Builtin.MustRegisterVendor("Minimum-Data-Rate-Downstream-Low-Power", 3561, 138, AttributeInteger)
	Builtin.MustRegisterVendor("Maximum-Interleaving-Delay-Upstream", 3561, 139, AttributeInteger)
	Builtin.MustRegister("PKM-Config-Settings", PKMConfigSettings_Type, AttributeOctets)
	Builtin.MustRegister("PKM-Cryptosuite-List", PKMCryptosuiteList_Type, AttributeOctets)
	Builtin.MustRegisterVendor("Actual-Interleaving-Delay-Upstream", 3561, 140, AttributeInteger)
	Builtin.MustRegister("PKM-SAID", PKMSAID_Type, AttributeUnknown)
	Builtin.MustRegisterVendor("Maximum-Interleaving-Delay-Downstream", 3561, 141, AttributeInteger)
	Builtin.MustRegister("PKM-SA-Descriptor", PKMSADescriptor_Type, AttributeOctets)
	Builtin.MustRegisterVendor("Actual-Interleaving-Delay-Downstream", 3561, 142, AttributeInteger)
	Builtin.MustRegister("PKM-Auth-Key", PKMAuthKey_Type, AttributeOctets)
	Builtin.MustRegisterVendor("Access-Loop-Encapsulation", 3561, 144, AttributeOctets)
	Builtin.MustRegister("DS-Lite-Tunnel-Name", DSLiteTunnelName_Type, AttributeString)
	Builtin.MustRegister("Mobile-Node-Identifier", MobileNodeIdentifier_Type, AttributeOctets)
	Builtin.MustRegister("Service-Selection", ServiceSelection_Type, AttributeString)
	Builtin.MustRegister("PMIP6-Home-LMA-IPv6-Address", PMIP6HomeLMAIPv6Address_Type, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Visited-LMA-IPv6-Address", PMIP6VisitedLMAIPv6Address_Type, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Home-LMA-IPv4-Address", PMIP6HomeLMAIPv4Address_Type, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Visited-LMA-IPv4-Address", PMIP6VisitedLMAIPv4Address_Type, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Home-HN-Prefix", PMIP6HomeHNPrefix_Type, AttributeIPv6Prefix)
	Builtin.MustRegister("PMIP6-Visited-HN-Prefix", PMIP6VisitedHNPrefix_Type, AttributeIPv6Prefix)
	Builtin.MustRegister("PMIP6-Home-Interface-ID", PMIP6HomeInterfaceID_Type, AttributeIFID)
	Builtin.MustRegister("PMIP6-Visited-Interface-ID", PMIP6VisitedInterfaceID_Type, AttributeIFID)
	Builtin.MustRegister("PMIP6-Home-DHCP4-Server-Address", PMIP6HomeDHCP4ServerAddress_Type, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Visited-DHCP4-Server-Address", PMIP6VisitedDHCP4ServerAddress_Type, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Home-DHCP6-Server-Address", PMIP6HomeDHCP6ServerAddress_Type, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Visited-DHCP6-Server-Address", PMIP6VisitedDHCP6ServerAddress_Type, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Home-IPv4-Gateway", PMIP6HomeIPv4Gateway_Type, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Visited-IPv4-Gateway", PMIP6VisitedIPv4Gateway_Type, AttributeIPAddr)
	Builtin.MustRegister("EAP-Lower-Layer", EAPLowerLayer_Type, AttributeInteger)
	Builtin.MustRegister("GSS-Acceptor-Service-Name", GSSAcceptorServiceName_Type, AttributeString)
	Builtin.MustRegister("GSS-Acceptor-Host-Name", GSSAcceptorHostName_Type, AttributeString)
	Builtin.MustRegister("GSS-Acceptor-Service-Specifics", GSSAcceptorServiceSpecifics_Type, AttributeString)
	Builtin.MustRegister("GSS-Acceptor-Realm-Name", GSSAcceptorRealmName_Type, AttributeString)
	Builtin.MustRegister("Framed-IPv6-Address", FramedIPv6Address_Type, AttributeIPv6Addr)
	Builtin.MustRegister("DNS-Server-IPv6-Address", DNSServerIPv6Address_Type, AttributeIPv6Addr)
	Builtin.MustRegister("Route-IPv6-Information", RouteIPv6Information_Type, AttributeIPv6Prefix)
	Builtin.MustRegister("Delegated-IPv6-Prefix-Pool", DelegatedIPv6PrefixPool_Type, AttributeString)
	Builtin.MustRegister("Stateful-IPv6-Address-Pool", StatefulIPv6AddressPool_Type, AttributeString)
	Builtin.MustRegister("Allowed-Called-Station-Id", AllowedCalledStationId_Type, AttributeString)
	Builtin.MustRegister("EAP-Peer-Id", EAPPeerId_Type, AttributeOctets)
	Builtin.MustRegister("EAP-Server-Id", EAPServerId_Type, AttributeOctets)
	Builtin.MustRegister("Mobility-Domain-Id", MobilityDomainId_Type, AttributeInteger)
	Builtin.MustRegister("Preauth-Timeout", PreauthTimeout_Type, AttributeInteger)
	Builtin.MustRegister("Network-Id-Name", NetworkIdName_Type, AttributeOctets)
	Builtin.MustRegister("WLAN-HESSID", WLANHESSID_Type, AttributeString)
	Builtin.MustRegister("WLAN-Venue-Info", WLANVenueInfo_Type, AttributeInteger)
	Builtin.MustRegister("WLAN-Venue-Language", WLANVenueLanguage_Type, AttributeOctets)
	Builtin.MustRegister("WLAN-Venue-Name", WLANVenueName_Type, AttributeString)
	Builtin.MustRegister("WLAN-Reason-Code", WLANReasonCode_Type, AttributeInteger)
	Builtin.MustRegister("WLAN-Pairwise-Cipher", WLANPairwiseCipher_Type, AttributeInteger)
	Builtin.MustRegister("WLAN-Group-Cipher", WLANGroupCipher_Type, AttributeInteger)
	Builtin.MustRegister("WLAN-AKM-Suite", WLANAKMSuite_Type, AttributeInteger)
	Builtin.MustRegister("WLAN-Group-Mgmt-Cipher", WLANGroupMgmtCipher_Type, AttributeInteger)
	Builtin.MustRegister("WLAN-RF-Band", WLANRFBand_Type, AttributeInteger)
	Builtin.MustRegisterVendor("IWF-Session", 3561, 254, AttributeOctets)
}
