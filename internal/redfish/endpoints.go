package redfish

var power_endpoints = []string {
	"/redfish/v1/Chassis/System.Embedded.1/Power/PowerControl",
}

var temperature_endpoints = []string {
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23SystemBoardInletTemp",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23SystemBoardExhaustTemp",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23CPU2Temp",
	"/redfish/v1/Chassis/System.Embedded.1/Sensors/Temperatures/iDRAC.Embedded.1%23CPU1Temp",
}

var fan_endpoints = []string {

}