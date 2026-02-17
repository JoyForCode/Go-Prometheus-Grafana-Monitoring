package redfish

var BaseURL = "https://192.168.1.101"

var Power_Endpoints = []string{
	"/redfish/v1/Chassis/System.Embedded.1/Power/PowerControl",
}

var Temperature_Endpoints = []string{
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23SystemBoardInletTemp",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23SystemBoardExhaustTemp",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23CPU2Temp",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23CPU1Temp",
}

var Fan_Endpoints = []string{
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Fans/0x17%7C%7CFan.Embedded.1",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Fans/0x17%7C%7CFan.Embedded.2",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Fans/0x17%7C%7CFan.Embedded.3",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Fans/0x17%7C%7CFan.Embedded.4",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Fans/0x17%7C%7CFan.Embedded.5",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Fans/0x17%7C%7CFan.Embedded.6",
}
