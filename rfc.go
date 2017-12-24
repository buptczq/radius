package radius

const (
	AttrUserName                                                     AttributeKey = 1
	AttrUserPassword                                                 AttributeKey = 2
	AttrCHAPPassword                                                 AttributeKey = 3
	AttrNASIPAddress                                                 AttributeKey = 4
	AttrNASPort                                                      AttributeKey = 5
	AttrServiceType                                                  AttributeKey = 6
	AttrFramedProtocol                                               AttributeKey = 7
	AttrFramedIPAddress                                              AttributeKey = 8
	AttrFramedIPNetmask                                              AttributeKey = 9
	AttrFramedRouting                                                AttributeKey = 10
	AttrFilterId                                                     AttributeKey = 11
	AttrFramedMTU                                                    AttributeKey = 12
	AttrFramedCompression                                            AttributeKey = 13
	AttrLoginIPHost                                                  AttributeKey = 14
	AttrLoginService                                                 AttributeKey = 15
	AttrLoginTCPPort                                                 AttributeKey = 16
	AttrReplyMessage                                                 AttributeKey = 18
	AttrCallbackNumber                                               AttributeKey = 19
	AttrCallbackId                                                   AttributeKey = 20
	AttrFramedRoute                                                  AttributeKey = 22
	AttrFramedIPXNetwork                                             AttributeKey = 23
	AttrState                                                        AttributeKey = 24
	AttrClass                                                        AttributeKey = 25
	AttrSessionTimeout                                               AttributeKey = 27
	AttrIdleTimeout                                                  AttributeKey = 28
	AttrTerminationAction                                            AttributeKey = 29
	AttrCalledStationId                                              AttributeKey = 30
	AttrCallingStationId                                             AttributeKey = 31
	AttrNASIdentifier                                                AttributeKey = 32
	AttrProxyState                                                   AttributeKey = 33
	AttrLoginLATService                                              AttributeKey = 34
	AttrLoginLATNode                                                 AttributeKey = 35
	AttrLoginLATGroup                                                AttributeKey = 36
	AttrFramedAppleTalkLink                                          AttributeKey = 37
	AttrFramedAppleTalkNetwork                                       AttributeKey = 38
	AttrFramedAppleTalkZone                                          AttributeKey = 39
	AttrAcctStatusType                                               AttributeKey = 40
	AttrAcctDelayTime                                                AttributeKey = 41
	AttrAcctInputOctets                                              AttributeKey = 42
	AttrAcctOutputOctets                                             AttributeKey = 43
	AttrAcctSessionId                                                AttributeKey = 44
	AttrAcctAuthentic                                                AttributeKey = 45
	AttrAcctSessionTime                                              AttributeKey = 46
	AttrAcctInputPackets                                             AttributeKey = 47
	AttrAcctOutputPackets                                            AttributeKey = 48
	AttrAcctTerminateCause                                           AttributeKey = 49
	AttrAcctMultiSessionId                                           AttributeKey = 50
	AttrAcctLinkCount                                                AttributeKey = 51
	AttrAcctInputGigawords                                           AttributeKey = 52
	AttrAcctOutputGigawords                                          AttributeKey = 53
	AttrEventTimestamp                                               AttributeKey = 55
	AttrEgressVLANID                                                 AttributeKey = 56
	AttrIngressFilters                                               AttributeKey = 57
	AttrEgressVLANName                                               AttributeKey = 58
	AttrUserPriorityTable                                            AttributeKey = 59
	AttrCHAPChallenge                                                AttributeKey = 60
	AttrNASPortType                                                  AttributeKey = 61
	AttrPortLimit                                                    AttributeKey = 62
	AttrLoginLATPort                                                 AttributeKey = 63
	AttrTunnelType                                                   AttributeKey = 64
	AttrTunnelMediumType                                             AttributeKey = 65
	AttrTunnelClientEndpoint                                         AttributeKey = 66
	AttrTunnelServerEndpoint                                         AttributeKey = 67
	AttrAcctTunnelConnection                                         AttributeKey = 68
	AttrTunnelPassword                                               AttributeKey = 69
	AttrARAPPassword                                                 AttributeKey = 70
	AttrARAPFeatures                                                 AttributeKey = 71
	AttrARAPZoneAccess                                               AttributeKey = 72
	AttrARAPSecurity                                                 AttributeKey = 73
	AttrARAPSecurityData                                             AttributeKey = 74
	AttrPasswordRetry                                                AttributeKey = 75
	AttrPrompt                                                       AttributeKey = 76
	AttrConnectInfo                                                  AttributeKey = 77
	AttrConfigurationToken                                           AttributeKey = 78
	AttrMessageAuthenticator                                         AttributeKey = 80
	AttrTunnelPrivateGroupId                                         AttributeKey = 81
	AttrTunnelAssignmentId                                           AttributeKey = 82
	AttrTunnelPreference                                             AttributeKey = 83
	AttrARAPChallengeResponse                                        AttributeKey = 84
	AttrAcctInterimInterval                                          AttributeKey = 85
	AttrAcctTunnelPacketsLost                                        AttributeKey = 86
	AttrNASPortId                                                    AttributeKey = 87
	AttrFramedPool                                                   AttributeKey = 88
	AttrChargeableUserIdentity                                       AttributeKey = 89
	AttrTunnelClientAuthId                                           AttributeKey = 90
	AttrTunnelServerAuthId                                           AttributeKey = 91
	AttrNASFilterRule                                                AttributeKey = 92
	AttrOriginatingLineInfo                                          AttributeKey = 94
	AttrNASIPv6Address                                               AttributeKey = 95
	AttrFramedInterfaceId                                            AttributeKey = 96
	AttrFramedIPv6Prefix                                             AttributeKey = 97
	AttrLoginIPv6Host                                                AttributeKey = 98
	AttrFramedIPv6Route                                              AttributeKey = 99
	AttrFramedIPv6Pool                                               AttributeKey = 100
	AttrErrorCause                                                   AttributeKey = 101
	AttrEAPKeyName                                                   AttributeKey = 102
	AttrDigestResponse                                               AttributeKey = 103
	AttrDigestRealm                                                  AttributeKey = 104
	AttrDigestNonce                                                  AttributeKey = 105
	AttrDigestResponseAuth                                           AttributeKey = 106
	AttrDigestNextnonce                                              AttributeKey = 107
	AttrDigestMethod                                                 AttributeKey = 108
	AttrDigestURI                                                    AttributeKey = 109
	AttrDigestQop                                                    AttributeKey = 110
	AttrDigestAlgorithm                                              AttributeKey = 111
	AttrDigestEntityBodyHash                                         AttributeKey = 112
	AttrDigestCNonce                                                 AttributeKey = 113
	AttrDigestNonceCount                                             AttributeKey = 114
	AttrDigestUsername                                               AttributeKey = 115
	AttrDigestOpaque                                                 AttributeKey = 116
	AttrDigestAuthParam                                              AttributeKey = 117
	AttrDigestAKAAuts                                                AttributeKey = 118
	AttrDigestDomain                                                 AttributeKey = 119
	AttrDigestStale                                                  AttributeKey = 120
	AttrDigestHA1                                                    AttributeKey = 121
	AttrSIPAOR                                                       AttributeKey = 122
	AttrDelegatedIPv6Prefix                                          AttributeKey = 123
	AttrMIP6FeatureVector                                            AttributeKey = 124
	AttrMIP6HomeLinkPrefix                                           AttributeKey = 125
	AttrOperatorName                                                 AttributeKey = 126
	AttrLocationInformation                                          AttributeKey = 127
	AttrLocationData                                                 AttributeKey = 128
	AttrBasicLocationPolicyRules                                     AttributeKey = 129
	AttrExtendedLocationPolicyRules                                  AttributeKey = 130
	AttrLocationCapable                                              AttributeKey = 131
	AttrRequestedLocationInfo                                        AttributeKey = 132
	AttrFramedManagement                                             AttributeKey = 133
	AttrManagementTransportProtection                                AttributeKey = 134
	AttrManagementPolicyId                                           AttributeKey = 135
	AttrManagementPrivilegeLevel                                     AttributeKey = 136
	AttrPKMConfigSettings                                            AttributeKey = 139
	AttrPKMCryptosuiteList                                           AttributeKey = 140
	AttrPKMSAID                                                      AttributeKey = 141
	AttrPKMSADescriptor                                              AttributeKey = 142
	AttrPKMAuthKey                                                   AttributeKey = 143
	AttrDSLiteTunnelName                                             AttributeKey = 144
	AttrMobileNodeIdentifier                                         AttributeKey = 145
	AttrServiceSelection                                             AttributeKey = 146
	AttrPMIP6HomeLMAIPv6Address                                      AttributeKey = 147
	AttrPMIP6VisitedLMAIPv6Address                                   AttributeKey = 148
	AttrPMIP6HomeLMAIPv4Address                                      AttributeKey = 149
	AttrPMIP6VisitedLMAIPv4Address                                   AttributeKey = 150
	AttrPMIP6HomeHNPrefix                                            AttributeKey = 151
	AttrPMIP6VisitedHNPrefix                                         AttributeKey = 152
	AttrPMIP6HomeInterfaceID                                         AttributeKey = 153
	AttrPMIP6VisitedInterfaceID                                      AttributeKey = 154
	AttrPMIP6HomeDHCP4ServerAddress                                  AttributeKey = 157
	AttrPMIP6VisitedDHCP4ServerAddress                               AttributeKey = 158
	AttrPMIP6HomeDHCP6ServerAddress                                  AttributeKey = 159
	AttrPMIP6VisitedDHCP6ServerAddress                               AttributeKey = 160
	AttrPMIP6HomeIPv4Gateway                                         AttributeKey = 161
	AttrPMIP6VisitedIPv4Gateway                                      AttributeKey = 162
	AttrEAPLowerLayer                                                AttributeKey = 163
	AttrGSSAcceptorServiceName                                       AttributeKey = 164
	AttrGSSAcceptorHostName                                          AttributeKey = 165
	AttrGSSAcceptorServiceSpecifics                                  AttributeKey = 166
	AttrGSSAcceptorRealmName                                         AttributeKey = 167
	AttrFramedIPv6Address                                            AttributeKey = 168
	AttrDNSServerIPv6Address                                         AttributeKey = 169
	AttrRouteIPv6Information                                         AttributeKey = 170
	AttrDelegatedIPv6PrefixPool                                      AttributeKey = 171
	AttrStatefulIPv6AddressPool                                      AttributeKey = 172
	AttrAllowedCalledStationId                                       AttributeKey = 174
	AttrEAPPeerId                                                    AttributeKey = 175
	AttrEAPServerId                                                  AttributeKey = 176
	AttrMobilityDomainId                                             AttributeKey = 177
	AttrPreauthTimeout                                               AttributeKey = 178
	AttrNetworkIdName                                                AttributeKey = 179
	AttrWLANHESSID                                                   AttributeKey = 181
	AttrWLANVenueInfo                                                AttributeKey = 182
	AttrWLANVenueLanguage                                            AttributeKey = 183
	AttrWLANVenueName                                                AttributeKey = 184
	AttrWLANReasonCode                                               AttributeKey = 185
	AttrWLANPairwiseCipher                                           AttributeKey = 186
	AttrWLANGroupCipher                                              AttributeKey = 187
	AttrWLANAKMSuite                                                 AttributeKey = 188
	AttrWLANGroupMgmtCipher                                          AttributeKey = 189
	AttrWLANRFBand                                                   AttributeKey = 190
	AttrADSLAgentCircuitId                                           AttributeKey = 15294378541057
	AttrADSLAgentRemoteId                                            AttributeKey = 15294378541058
	AttrActualDataRateUpstream                                       AttributeKey = 15294378541185
	AttrActualDataRateDownstream                                     AttributeKey = 15294378541186
	AttrMinimumDataRateUpstream                                      AttributeKey = 15294378541187
	AttrMinimumDataRateDownstream                                    AttributeKey = 15294378541188
	AttrAttainableDataRateUpstream                                   AttributeKey = 15294378541189
	AttrAttainableDataRateDownstream                                 AttributeKey = 15294378541190
	AttrMaximumDataRateUpstream                                      AttributeKey = 15294378541191
	AttrMaximumDataRateDownstream                                    AttributeKey = 15294378541192
	AttrMinimumDataRateUpstreamLowPower                              AttributeKey = 15294378541193
	AttrMinimumDataRateDownstreamLowPower                            AttributeKey = 15294378541194
	AttrMaximumInterleavingDelayUpstream                             AttributeKey = 15294378541195
	AttrActualInterleavingDelayUpstream                              AttributeKey = 15294378541196
	AttrMaximumInterleavingDelayDownstream                           AttributeKey = 15294378541197
	AttrActualInterleavingDelayDownstream                            AttributeKey = 15294378541198
	AttrAccessLoopEncapsulation                                      AttributeKey = 15294378541200
	AttrIWFSession                                                   AttributeKey = 15294378541310
	VendorADSLForum                                                  uint32       = 3561
	ServiceType_LoginUser                                            uint32       = 1
	ServiceType_FramedUser                                           uint32       = 2
	ServiceType_CallbackLoginUser                                    uint32       = 3
	ServiceType_CallbackFramedUser                                   uint32       = 4
	ServiceType_OutboundUser                                         uint32       = 5
	ServiceType_AdministrativeUser                                   uint32       = 6
	ServiceType_NASPromptUser                                        uint32       = 7
	ServiceType_AuthenticateOnly                                     uint32       = 8
	ServiceType_CallbackNASPrompt                                    uint32       = 9
	ServiceType_CallCheck                                            uint32       = 10
	ServiceType_CallbackAdministrative                               uint32       = 11
	ServiceType_AuthorizeOnly                                        uint32       = 17
	ServiceType_FramedManagement                                     uint32       = 18
	FramedProtocol_PPP                                               uint32       = 1
	FramedProtocol_SLIP                                              uint32       = 2
	FramedProtocol_ARAP                                              uint32       = 3
	FramedProtocol_GandalfSLML                                       uint32       = 4
	FramedProtocol_XylogicsIPXSLIP                                   uint32       = 5
	FramedProtocol_X75Synchronous                                    uint32       = 6
	FramedRouting_None                                               uint32       = 0
	FramedRouting_Broadcast                                          uint32       = 1
	FramedRouting_Listen                                             uint32       = 2
	FramedRouting_BroadcastListen                                    uint32       = 3
	FramedCompression_None                                           uint32       = 0
	FramedCompression_VanJacobsonTCPIP                               uint32       = 1
	FramedCompression_IPXHeaderCompression                           uint32       = 2
	FramedCompression_StacLZS                                        uint32       = 3
	LoginService_Telnet                                              uint32       = 0
	LoginService_Rlogin                                              uint32       = 1
	LoginService_TCPClear                                            uint32       = 2
	LoginService_PortMaster                                          uint32       = 3
	LoginService_LAT                                                 uint32       = 4
	LoginService_X25PAD                                              uint32       = 5
	LoginService_X25T3POS                                            uint32       = 6
	LoginService_TCPClearQuiet                                       uint32       = 8
	LoginTCPPort_Telnet                                              uint32       = 23
	LoginTCPPort_Rlogin                                              uint32       = 513
	LoginTCPPort_Rsh                                                 uint32       = 514
	TerminationAction_Default                                        uint32       = 0
	TerminationAction_RADIUSRequest                                  uint32       = 1
	AcctStatusType_Start                                             uint32       = 1
	AcctStatusType_Stop                                              uint32       = 2
	AcctStatusType_InterimUpdate                                     uint32       = 3
	AcctStatusType_Alive                                             uint32       = 3
	AcctStatusType_AccountingOn                                      uint32       = 7
	AcctStatusType_AccountingOff                                     uint32       = 8
	AcctStatusType_TunnelStart                                       uint32       = 9
	AcctStatusType_TunnelStop                                        uint32       = 10
	AcctStatusType_TunnelReject                                      uint32       = 11
	AcctStatusType_TunnelLinkStart                                   uint32       = 12
	AcctStatusType_TunnelLinkStop                                    uint32       = 13
	AcctStatusType_TunnelLinkReject                                  uint32       = 14
	AcctStatusType_Failed                                            uint32       = 15
	AcctAuthentic_RADIUS                                             uint32       = 1
	AcctAuthentic_Local                                              uint32       = 2
	AcctAuthentic_Remote                                             uint32       = 3
	AcctAuthentic_Diameter                                           uint32       = 4
	AcctTerminateCause_UserRequest                                   uint32       = 1
	AcctTerminateCause_LostCarrier                                   uint32       = 2
	AcctTerminateCause_LostService                                   uint32       = 3
	AcctTerminateCause_IdleTimeout                                   uint32       = 4
	AcctTerminateCause_SessionTimeout                                uint32       = 5
	AcctTerminateCause_AdminReset                                    uint32       = 6
	AcctTerminateCause_AdminReboot                                   uint32       = 7
	AcctTerminateCause_PortError                                     uint32       = 8
	AcctTerminateCause_NASError                                      uint32       = 9
	AcctTerminateCause_NASRequest                                    uint32       = 10
	AcctTerminateCause_NASReboot                                     uint32       = 11
	AcctTerminateCause_PortUnneeded                                  uint32       = 12
	AcctTerminateCause_PortPreempted                                 uint32       = 13
	AcctTerminateCause_PortSuspended                                 uint32       = 14
	AcctTerminateCause_ServiceUnavailable                            uint32       = 15
	AcctTerminateCause_Callback                                      uint32       = 16
	AcctTerminateCause_UserError                                     uint32       = 17
	AcctTerminateCause_HostRequest                                   uint32       = 18
	AcctTerminateCause_SupplicantRestart                             uint32       = 19
	AcctTerminateCause_ReauthenticationFailure                       uint32       = 20
	AcctTerminateCause_PortReinit                                    uint32       = 21
	AcctTerminateCause_PortDisabled                                  uint32       = 22
	IngressFilters_Enabled                                           uint32       = 1
	IngressFilters_Disabled                                          uint32       = 2
	NASPortType_Async                                                uint32       = 0
	NASPortType_Sync                                                 uint32       = 1
	NASPortType_ISDN                                                 uint32       = 2
	NASPortType_ISDNV120                                             uint32       = 3
	NASPortType_ISDNV110                                             uint32       = 4
	NASPortType_Virtual                                              uint32       = 5
	NASPortType_PIAFS                                                uint32       = 6
	NASPortType_HDLCClearChannel                                     uint32       = 7
	NASPortType_X25                                                  uint32       = 8
	NASPortType_X75                                                  uint32       = 9
	NASPortType_G3Fax                                                uint32       = 10
	NASPortType_SDSL                                                 uint32       = 11
	NASPortType_ADSLCAP                                              uint32       = 12
	NASPortType_ADSLDMT                                              uint32       = 13
	NASPortType_IDSL                                                 uint32       = 14
	NASPortType_Ethernet                                             uint32       = 15
	NASPortType_xDSL                                                 uint32       = 16
	NASPortType_Cable                                                uint32       = 17
	NASPortType_WirelessOther                                        uint32       = 18
	NASPortType_Wireless80211                                        uint32       = 19
	NASPortType_TokenRing                                            uint32       = 20
	NASPortType_FDDI                                                 uint32       = 21
	NASPortType_PPPoA                                                uint32       = 30
	NASPortType_PPPoEoA                                              uint32       = 31
	NASPortType_PPPoEoE                                              uint32       = 32
	NASPortType_PPPoEoVLAN                                           uint32       = 33
	NASPortType_PPPoEoQinQ                                           uint32       = 34
	TunnelType_PPTP                                                  uint32       = 1
	TunnelType_L2F                                                   uint32       = 2
	TunnelType_L2TP                                                  uint32       = 3
	TunnelType_ATMP                                                  uint32       = 4
	TunnelType_VTP                                                   uint32       = 5
	TunnelType_AH                                                    uint32       = 6
	TunnelType_IP                                                    uint32       = 7
	TunnelType_MINIP                                                 uint32       = 8
	TunnelType_ESP                                                   uint32       = 9
	TunnelType_GRE                                                   uint32       = 10
	TunnelType_DVS                                                   uint32       = 11
	TunnelType_IPinIP                                                uint32       = 12
	TunnelType_VLAN                                                  uint32       = 13
	TunnelMediumType_IP                                              uint32       = 1
	TunnelMediumType_IPv4                                            uint32       = 1
	TunnelMediumType_IPv6                                            uint32       = 2
	TunnelMediumType_NSAP                                            uint32       = 3
	TunnelMediumType_HDLC                                            uint32       = 4
	TunnelMediumType_BBN1822                                         uint32       = 5
	TunnelMediumType_IEEE802                                         uint32       = 6
	TunnelMediumType_E163                                            uint32       = 7
	TunnelMediumType_E164                                            uint32       = 8
	TunnelMediumType_F69                                             uint32       = 9
	TunnelMediumType_X121                                            uint32       = 10
	TunnelMediumType_IPX                                             uint32       = 11
	TunnelMediumType_Appletalk                                       uint32       = 12
	TunnelMediumType_DecNetIV                                        uint32       = 13
	TunnelMediumType_BanyanVines                                     uint32       = 14
	TunnelMediumType_E164NSAP                                        uint32       = 15
	ARAPZoneAccess_DefaultZone                                       uint32       = 1
	ARAPZoneAccess_ZoneFilterInclusive                               uint32       = 2
	ARAPZoneAccess_ZoneFilterExclusive                               uint32       = 4
	Prompt_NoEcho                                                    uint32       = 0
	Prompt_Echo                                                      uint32       = 1
	ErrorCause_ResidualContextRemoved                                uint32       = 201
	ErrorCause_InvalidEAPPacket                                      uint32       = 202
	ErrorCause_UnsupportedAttribute                                  uint32       = 401
	ErrorCause_MissingAttribute                                      uint32       = 402
	ErrorCause_NASIdentificationMismatch                             uint32       = 403
	ErrorCause_InvalidRequest                                        uint32       = 404
	ErrorCause_UnsupportedService                                    uint32       = 405
	ErrorCause_UnsupportedExtension                                  uint32       = 406
	ErrorCause_InvalidAttributeValue                                 uint32       = 407
	ErrorCause_AdministrativelyProhibited                            uint32       = 501
	ErrorCause_ProxyRequestNotRoutable                               uint32       = 502
	ErrorCause_SessionContextNotFound                                uint32       = 503
	ErrorCause_SessionContextNotRemovable                            uint32       = 504
	ErrorCause_ProxyProcessingError                                  uint32       = 505
	ErrorCause_ResourcesUnavailable                                  uint32       = 506
	ErrorCause_RequestInitiated                                      uint32       = 507
	ErrorCause_MultipleSessionSelectionUnsupported                   uint32       = 508
	LocationCapable_CivicLocation                                    uint32       = 1
	LocationCapable_GeoLocation                                      uint32       = 2
	LocationCapable_UsersLocation                                    uint32       = 4
	LocationCapable_NASLocation                                      uint32       = 8
	RequestedLocationInfo_CivicLocation                              uint32       = 1
	RequestedLocationInfo_GeoLocation                                uint32       = 2
	RequestedLocationInfo_UsersLocation                              uint32       = 4
	RequestedLocationInfo_NASLocation                                uint32       = 8
	RequestedLocationInfo_FutureRequests                             uint32       = 16
	RequestedLocationInfo_None                                       uint32       = 32
	FramedManagement_SNMP                                            uint32       = 1
	FramedManagement_WebBased                                        uint32       = 2
	FramedManagement_Netconf                                         uint32       = 3
	FramedManagement_FTP                                             uint32       = 4
	FramedManagement_TFTP                                            uint32       = 5
	FramedManagement_SFTP                                            uint32       = 6
	FramedManagement_RCP                                             uint32       = 7
	FramedManagement_SCP                                             uint32       = 8
	ManagementTransportProtection_NoProtection                       uint32       = 1
	ManagementTransportProtection_IntegrityProtection                uint32       = 2
	ManagementTransportProtection_IntegrityConfidentialityProtection uint32       = 3
	EAPLowerLayer_WiredIEEE8021X                                     uint32       = 1
	EAPLowerLayer_IEEE8021XNoPreauth                                 uint32       = 2
	EAPLowerLayer_IEEE8021XPreauth                                   uint32       = 3
	EAPLowerLayer_IEEE80216e                                         uint32       = 4
	EAPLowerLayer_IKEv2                                              uint32       = 5
	EAPLowerLayer_PPP                                                uint32       = 6
	EAPLowerLayer_PANANoPreauth                                      uint32       = 7
	EAPLowerLayer_GSSAPI                                             uint32       = 8
	EAPLowerLayer_PANAPreauth                                        uint32       = 9
)

func init() {
	builtinOnce.Do(initDictionary)
	Builtin.MustRegister("User-Name", 0, 1, false, 0, AttributeString)
	Builtin.MustRegister("User-Password", 0, 2, false, 1, AttributeString)
	Builtin.MustRegister("CHAP-Password", 0, 3, false, 0, AttributeOctets)
	Builtin.MustRegister("NAS-IP-Address", 0, 4, false, 0, AttributeIPAddr)
	Builtin.MustRegister("NAS-Port", 0, 5, false, 0, AttributeInteger)
	Builtin.MustRegister("Service-Type", 0, 6, false, 0, AttributeInteger)
	Builtin.MustRegister("Framed-Protocol", 0, 7, false, 0, AttributeInteger)
	Builtin.MustRegister("Framed-IP-Address", 0, 8, false, 0, AttributeIPAddr)
	Builtin.MustRegister("Framed-IP-Netmask", 0, 9, false, 0, AttributeIPAddr)
	Builtin.MustRegister("Framed-Routing", 0, 10, false, 0, AttributeInteger)
	Builtin.MustRegister("Filter-Id", 0, 11, false, 0, AttributeString)
	Builtin.MustRegister("Framed-MTU", 0, 12, false, 0, AttributeInteger)
	Builtin.MustRegister("Framed-Compression", 0, 13, false, 0, AttributeInteger)
	Builtin.MustRegister("Login-IP-Host", 0, 14, false, 0, AttributeIPAddr)
	Builtin.MustRegister("Login-Service", 0, 15, false, 0, AttributeInteger)
	Builtin.MustRegister("Login-TCP-Port", 0, 16, false, 0, AttributeInteger)
	Builtin.MustRegister("Reply-Message", 0, 18, false, 0, AttributeString)
	Builtin.MustRegister("Callback-Number", 0, 19, false, 0, AttributeString)
	Builtin.MustRegister("Callback-Id", 0, 20, false, 0, AttributeString)
	Builtin.MustRegister("Framed-Route", 0, 22, false, 0, AttributeString)
	Builtin.MustRegister("Framed-IPX-Network", 0, 23, false, 0, AttributeIPAddr)
	Builtin.MustRegister("State", 0, 24, false, 0, AttributeOctets)
	Builtin.MustRegister("Class", 0, 25, false, 0, AttributeOctets)
	Builtin.MustRegister("Session-Timeout", 0, 27, false, 0, AttributeInteger)
	Builtin.MustRegister("Idle-Timeout", 0, 28, false, 0, AttributeInteger)
	Builtin.MustRegister("Termination-Action", 0, 29, false, 0, AttributeInteger)
	Builtin.MustRegister("Called-Station-Id", 0, 30, false, 0, AttributeString)
	Builtin.MustRegister("Calling-Station-Id", 0, 31, false, 0, AttributeString)
	Builtin.MustRegister("NAS-Identifier", 0, 32, false, 0, AttributeString)
	Builtin.MustRegister("Proxy-State", 0, 33, false, 0, AttributeOctets)
	Builtin.MustRegister("Login-LAT-Service", 0, 34, false, 0, AttributeString)
	Builtin.MustRegister("Login-LAT-Node", 0, 35, false, 0, AttributeString)
	Builtin.MustRegister("Login-LAT-Group", 0, 36, false, 0, AttributeOctets)
	Builtin.MustRegister("Framed-AppleTalk-Link", 0, 37, false, 0, AttributeInteger)
	Builtin.MustRegister("Framed-AppleTalk-Network", 0, 38, false, 0, AttributeInteger)
	Builtin.MustRegister("Framed-AppleTalk-Zone", 0, 39, false, 0, AttributeString)
	Builtin.MustRegister("Acct-Status-Type", 0, 40, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Delay-Time", 0, 41, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Input-Octets", 0, 42, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Output-Octets", 0, 43, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Session-Id", 0, 44, false, 0, AttributeString)
	Builtin.MustRegister("Acct-Authentic", 0, 45, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Session-Time", 0, 46, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Input-Packets", 0, 47, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Output-Packets", 0, 48, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Terminate-Cause", 0, 49, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Multi-Session-Id", 0, 50, false, 0, AttributeString)
	Builtin.MustRegister("Acct-Link-Count", 0, 51, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Input-Gigawords", 0, 52, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Output-Gigawords", 0, 53, false, 0, AttributeInteger)
	Builtin.MustRegister("Event-Timestamp", 0, 55, false, 0, AttributeDate)
	Builtin.MustRegister("Egress-VLANID", 0, 56, false, 0, AttributeInteger)
	Builtin.MustRegister("Ingress-Filters", 0, 57, false, 0, AttributeInteger)
	Builtin.MustRegister("Egress-VLAN-Name", 0, 58, false, 0, AttributeString)
	Builtin.MustRegister("User-Priority-Table", 0, 59, false, 0, AttributeOctets)
	Builtin.MustRegister("CHAP-Challenge", 0, 60, false, 0, AttributeOctets)
	Builtin.MustRegister("NAS-Port-Type", 0, 61, false, 0, AttributeInteger)
	Builtin.MustRegister("Port-Limit", 0, 62, false, 0, AttributeInteger)
	Builtin.MustRegister("Login-LAT-Port", 0, 63, false, 0, AttributeString)
	Builtin.MustRegister("Tunnel-Type", 0, 64, true, 0, AttributeInteger)
	Builtin.MustRegister("Tunnel-Medium-Type", 0, 65, true, 0, AttributeInteger)
	Builtin.MustRegister("Tunnel-Client-Endpoint", 0, 66, true, 0, AttributeString)
	Builtin.MustRegister("Tunnel-Server-Endpoint", 0, 67, true, 0, AttributeString)
	Builtin.MustRegister("Acct-Tunnel-Connection", 0, 68, false, 0, AttributeString)
	Builtin.MustRegister("Tunnel-Password", 0, 69, true, 2, AttributeString)
	Builtin.MustRegister("ARAP-Password", 0, 70, false, 0, AttributeOctets)
	Builtin.MustRegister("ARAP-Features", 0, 71, false, 0, AttributeOctets)
	Builtin.MustRegister("ARAP-Zone-Access", 0, 72, false, 0, AttributeInteger)
	Builtin.MustRegister("ARAP-Security", 0, 73, false, 0, AttributeInteger)
	Builtin.MustRegister("ARAP-Security-Data", 0, 74, false, 0, AttributeString)
	Builtin.MustRegister("Password-Retry", 0, 75, false, 0, AttributeInteger)
	Builtin.MustRegister("Prompt", 0, 76, false, 0, AttributeInteger)
	Builtin.MustRegister("Connect-Info", 0, 77, false, 0, AttributeString)
	Builtin.MustRegister("Configuration-Token", 0, 78, false, 0, AttributeString)
	Builtin.MustRegister("Message-Authenticator", 0, 80, false, 0, AttributeOctets)
	Builtin.MustRegister("Tunnel-Private-Group-Id", 0, 81, true, 0, AttributeString)
	Builtin.MustRegister("Tunnel-Assignment-Id", 0, 82, true, 0, AttributeString)
	Builtin.MustRegister("Tunnel-Preference", 0, 83, true, 0, AttributeInteger)
	Builtin.MustRegister("ARAP-Challenge-Response", 0, 84, false, 0, AttributeOctets)
	Builtin.MustRegister("Acct-Interim-Interval", 0, 85, false, 0, AttributeInteger)
	Builtin.MustRegister("Acct-Tunnel-Packets-Lost", 0, 86, false, 0, AttributeInteger)
	Builtin.MustRegister("NAS-Port-Id", 0, 87, false, 0, AttributeString)
	Builtin.MustRegister("Framed-Pool", 0, 88, false, 0, AttributeString)
	Builtin.MustRegister("Chargeable-User-Identity", 0, 89, false, 0, AttributeOctets)
	Builtin.MustRegister("Tunnel-Client-Auth-Id", 0, 90, true, 0, AttributeString)
	Builtin.MustRegister("Tunnel-Server-Auth-Id", 0, 91, true, 0, AttributeString)
	Builtin.MustRegister("NAS-Filter-Rule", 0, 92, false, 0, AttributeString)
	Builtin.MustRegister("Originating-Line-Info", 0, 94, false, 0, AttributeOctets)
	Builtin.MustRegister("NAS-IPv6-Address", 0, 95, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("Framed-Interface-Id", 0, 96, false, 0, AttributeIFID)
	Builtin.MustRegister("Framed-IPv6-Prefix", 0, 97, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegister("Login-IPv6-Host", 0, 98, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("Framed-IPv6-Route", 0, 99, false, 0, AttributeString)
	Builtin.MustRegister("Framed-IPv6-Pool", 0, 100, false, 0, AttributeString)
	Builtin.MustRegister("Error-Cause", 0, 101, false, 0, AttributeInteger)
	Builtin.MustRegister("EAP-Key-Name", 0, 102, false, 0, AttributeOctets)
	Builtin.MustRegister("Digest-Response", 0, 103, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Realm", 0, 104, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Nonce", 0, 105, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Response-Auth", 0, 106, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Nextnonce", 0, 107, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Method", 0, 108, false, 0, AttributeString)
	Builtin.MustRegister("Digest-URI", 0, 109, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Qop", 0, 110, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Algorithm", 0, 111, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Entity-Body-Hash", 0, 112, false, 0, AttributeString)
	Builtin.MustRegister("Digest-CNonce", 0, 113, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Nonce-Count", 0, 114, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Username", 0, 115, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Opaque", 0, 116, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Auth-Param", 0, 117, false, 0, AttributeString)
	Builtin.MustRegister("Digest-AKA-Auts", 0, 118, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Domain", 0, 119, false, 0, AttributeString)
	Builtin.MustRegister("Digest-Stale", 0, 120, false, 0, AttributeString)
	Builtin.MustRegister("Digest-HA1", 0, 121, false, 0, AttributeString)
	Builtin.MustRegister("SIP-AOR", 0, 122, false, 0, AttributeString)
	Builtin.MustRegister("Delegated-IPv6-Prefix", 0, 123, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegister("MIP6-Feature-Vector", 0, 124, false, 0, AttributeOctets)
	Builtin.MustRegister("MIP6-Home-Link-Prefix", 0, 125, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegister("Operator-Name", 0, 126, false, 0, AttributeString)
	Builtin.MustRegister("Location-Information", 0, 127, false, 0, AttributeOctets)
	Builtin.MustRegister("Location-Data", 0, 128, false, 0, AttributeOctets)
	Builtin.MustRegister("Basic-Location-Policy-Rules", 0, 129, false, 0, AttributeOctets)
	Builtin.MustRegister("Extended-Location-Policy-Rules", 0, 130, false, 0, AttributeOctets)
	Builtin.MustRegister("Location-Capable", 0, 131, false, 0, AttributeInteger)
	Builtin.MustRegister("Requested-Location-Info", 0, 132, false, 0, AttributeInteger)
	Builtin.MustRegister("Framed-Management", 0, 133, false, 0, AttributeInteger)
	Builtin.MustRegister("Management-Transport-Protection", 0, 134, false, 0, AttributeInteger)
	Builtin.MustRegister("Management-Policy-Id", 0, 135, false, 0, AttributeString)
	Builtin.MustRegister("Management-Privilege-Level", 0, 136, false, 0, AttributeInteger)
	Builtin.MustRegister("PKM-Config-Settings", 0, 139, false, 0, AttributeOctets)
	Builtin.MustRegister("PKM-Cryptosuite-List", 0, 140, false, 0, AttributeOctets)
	Builtin.MustRegister("PKM-SAID", 0, 141, false, 0, AttributeShort)
	Builtin.MustRegister("PKM-SA-Descriptor", 0, 142, false, 0, AttributeOctets)
	Builtin.MustRegister("PKM-Auth-Key", 0, 143, false, 0, AttributeOctets)
	Builtin.MustRegister("DS-Lite-Tunnel-Name", 0, 144, false, 0, AttributeString)
	Builtin.MustRegister("Mobile-Node-Identifier", 0, 145, false, 0, AttributeOctets)
	Builtin.MustRegister("Service-Selection", 0, 146, false, 0, AttributeString)
	Builtin.MustRegister("PMIP6-Home-LMA-IPv6-Address", 0, 147, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Visited-LMA-IPv6-Address", 0, 148, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Home-LMA-IPv4-Address", 0, 149, false, 0, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Visited-LMA-IPv4-Address", 0, 150, false, 0, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Home-HN-Prefix", 0, 151, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegister("PMIP6-Visited-HN-Prefix", 0, 152, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegister("PMIP6-Home-Interface-ID", 0, 153, false, 0, AttributeIFID)
	Builtin.MustRegister("PMIP6-Visited-Interface-ID", 0, 154, false, 0, AttributeIFID)
	Builtin.MustRegister("PMIP6-Home-DHCP4-Server-Address", 0, 157, false, 0, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Visited-DHCP4-Server-Address", 0, 158, false, 0, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Home-DHCP6-Server-Address", 0, 159, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Visited-DHCP6-Server-Address", 0, 160, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("PMIP6-Home-IPv4-Gateway", 0, 161, false, 0, AttributeIPAddr)
	Builtin.MustRegister("PMIP6-Visited-IPv4-Gateway", 0, 162, false, 0, AttributeIPAddr)
	Builtin.MustRegister("EAP-Lower-Layer", 0, 163, false, 0, AttributeInteger)
	Builtin.MustRegister("GSS-Acceptor-Service-Name", 0, 164, false, 0, AttributeString)
	Builtin.MustRegister("GSS-Acceptor-Host-Name", 0, 165, false, 0, AttributeString)
	Builtin.MustRegister("GSS-Acceptor-Service-Specifics", 0, 166, false, 0, AttributeString)
	Builtin.MustRegister("GSS-Acceptor-Realm-Name", 0, 167, false, 0, AttributeString)
	Builtin.MustRegister("Framed-IPv6-Address", 0, 168, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("DNS-Server-IPv6-Address", 0, 169, false, 0, AttributeIPv6Addr)
	Builtin.MustRegister("Route-IPv6-Information", 0, 170, false, 0, AttributeIPv6Prefix)
	Builtin.MustRegister("Delegated-IPv6-Prefix-Pool", 0, 171, false, 0, AttributeString)
	Builtin.MustRegister("Stateful-IPv6-Address-Pool", 0, 172, false, 0, AttributeString)
	Builtin.MustRegister("Allowed-Called-Station-Id", 0, 174, false, 0, AttributeString)
	Builtin.MustRegister("EAP-Peer-Id", 0, 175, false, 0, AttributeOctets)
	Builtin.MustRegister("EAP-Server-Id", 0, 176, false, 0, AttributeOctets)
	Builtin.MustRegister("Mobility-Domain-Id", 0, 177, false, 0, AttributeInteger)
	Builtin.MustRegister("Preauth-Timeout", 0, 178, false, 0, AttributeInteger)
	Builtin.MustRegister("Network-Id-Name", 0, 179, false, 0, AttributeOctets)
	Builtin.MustRegister("WLAN-HESSID", 0, 181, false, 0, AttributeString)
	Builtin.MustRegister("WLAN-Venue-Info", 0, 182, false, 0, AttributeInteger)
	Builtin.MustRegister("WLAN-Venue-Language", 0, 183, false, 0, AttributeOctets)
	Builtin.MustRegister("WLAN-Venue-Name", 0, 184, false, 0, AttributeString)
	Builtin.MustRegister("WLAN-Reason-Code", 0, 185, false, 0, AttributeInteger)
	Builtin.MustRegister("WLAN-Pairwise-Cipher", 0, 186, false, 0, AttributeInteger)
	Builtin.MustRegister("WLAN-Group-Cipher", 0, 187, false, 0, AttributeInteger)
	Builtin.MustRegister("WLAN-AKM-Suite", 0, 188, false, 0, AttributeInteger)
	Builtin.MustRegister("WLAN-Group-Mgmt-Cipher", 0, 189, false, 0, AttributeInteger)
	Builtin.MustRegister("WLAN-RF-Band", 0, 190, false, 0, AttributeInteger)
	Builtin.MustRegister("ADSL-Agent-Circuit-Id", 3561, 1, false, 0, AttributeOctets)
	Builtin.MustRegister("ADSL-Agent-Remote-Id", 3561, 2, false, 0, AttributeOctets)
	Builtin.MustRegister("Actual-Data-Rate-Upstream", 3561, 129, false, 0, AttributeInteger)
	Builtin.MustRegister("Actual-Data-Rate-Downstream", 3561, 130, false, 0, AttributeInteger)
	Builtin.MustRegister("Minimum-Data-Rate-Upstream", 3561, 131, false, 0, AttributeInteger)
	Builtin.MustRegister("Minimum-Data-Rate-Downstream", 3561, 132, false, 0, AttributeInteger)
	Builtin.MustRegister("Attainable-Data-Rate-Upstream", 3561, 133, false, 0, AttributeInteger)
	Builtin.MustRegister("Attainable-Data-Rate-Downstream", 3561, 134, false, 0, AttributeInteger)
	Builtin.MustRegister("Maximum-Data-Rate-Upstream", 3561, 135, false, 0, AttributeInteger)
	Builtin.MustRegister("Maximum-Data-Rate-Downstream", 3561, 136, false, 0, AttributeInteger)
	Builtin.MustRegister("Minimum-Data-Rate-Upstream-Low-Power", 3561, 137, false, 0, AttributeInteger)
	Builtin.MustRegister("Minimum-Data-Rate-Downstream-Low-Power", 3561, 138, false, 0, AttributeInteger)
	Builtin.MustRegister("Maximum-Interleaving-Delay-Upstream", 3561, 139, false, 0, AttributeInteger)
	Builtin.MustRegister("Actual-Interleaving-Delay-Upstream", 3561, 140, false, 0, AttributeInteger)
	Builtin.MustRegister("Maximum-Interleaving-Delay-Downstream", 3561, 141, false, 0, AttributeInteger)
	Builtin.MustRegister("Actual-Interleaving-Delay-Downstream", 3561, 142, false, 0, AttributeInteger)
	Builtin.MustRegister("Access-Loop-Encapsulation", 3561, 144, false, 0, AttributeOctets)
	Builtin.MustRegister("IWF-Session", 3561, 254, false, 0, AttributeOctets)
}
