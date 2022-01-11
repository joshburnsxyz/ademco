package ademco

var (
	eventDescription string
)

func GetMessageDescription(eventCode string, qualifier bool) string {
	switch eventCode {
	// Medical
	case "100":
		eventDescription = "Medical"
	case "101":
		eventDescription = "Medical-Emergency (Pendant Transmitter)"
	case "102":
		eventDescription = "Medical-Failed Check-In"
	// Fire Alarms
	case "110":
		eventDescription = "FIRE"
	case "111":
		eventDescription = "Fire-SMOKE"
	case "112":
		eventDescription = "Fire-Combustion"
	case "113":
		eventDescription = "Fire-WATERFLOW"
	case "114":
		eventDescription = "Fire-Heat"
	case "115":
		eventDescription = "Fire-Pull Station"
	case "116":
		eventDescription = "Fire-Duct"
	case "117":
		eventDescription = "Fire-Flame"
	case "118":
		eventDescription = "Fire-Near Alarm"
	// Panic Alarms
	case "120":
		eventDescription = "Panic"
	case "121":
		eventDescription = "Panic-DURESS"
	case "122":
		eventDescription = "Panic-SILENT"
	case "123":
		eventDescription = "Panic-AUDIBLE"
	case "124":
		eventDescription = "Panic-Duress Access Granted"
	case "125":
		eventDescription = "Panic-Duress Egress Granted"
	// Burglar Alarms
	case "130":
		eventDescription = "Burglary"
	case "131":
		eventDescription = "Burglary-PERIMETER"
	case "132":
		eventDescription = "Burglary-INTERIOR"
	case "133":
		eventDescription = "Burglary-24 HR BURG (AUX)"
	case "134":
		eventDescription = "Burglary-ENTRY/EXIT"
	case "135":
		eventDescription = "Burglary-DAY/NIGHT"
	case "136":
		eventDescription = "Burglary-Outdoor"
	case "137":
		eventDescription = "Burglary-TAMPER"
	case "138":
		eventDescription = "Burglary-Near Alarm"
	case "139":
		eventDescription = "Burglary-Intrusion Verifier"
	// General Alarms
	case "140":
		eventDescription = "Alarm-General Alarm"
	case "141":
		eventDescription = "Alarm-Polling Loop Open"
	case "142":
		eventDescription = "Alarm-Polling Loop Short"
	case "143":
		eventDescription = "Alarm-Exp. Module Tamper"
	case "144":
		eventDescription = "Alarm-Sensor Tamper"
	case "145":
		eventDescription = "Alarm-Exp. Module Tamper"
	case "146":
		eventDescription = "Burglary-Silent Burglary"
	case "147":
		eventDescription = "Trouble-Sensor Super."
	// 24-Hour non-burglary
	case "150":
		eventDescription = "Alarm-24 Hr. Non-Burglary"
	case "151":
		eventDescription = "Alarm-Gas Detected"
	case "152":
		eventDescription = "Alarm-Refrigeration"
	case "153":
		eventDescription = "Alarm-Heating System"
	case "154":
		eventDescription = "Alarm-Water Leakage"
	case "155":
		eventDescription = "Trouble-Foil Break"
	case "156":
		eventDescription = "Trouble-Day Zone"
	case "157":
		eventDescription = "Alarm-Low Gas Level"
	case "158":
		eventDescription = "Alarm-High Temperature"
	case "159":
		eventDescription = "Alarm-Low Temperature"
	case "161":
		eventDescription = "Alarm-Air Flow"
	case "162":
		eventDescription = "Alarm-Carbon Monoxide"
	case "163":
		eventDescription = "Trouble-Tank Level"
	case "168":
		eventDescription = "Trouble-High Humidity"
	case "169":
		eventDescription = "Trouble-Low Humidity"
	// Fire Supervisory
	case "200":
		eventDescription = "Super.-Fire Supervisory"
	case "201":
		eventDescription = "Super.-Low Water Pressure"
	case "202":
		eventDescription = "Super.-Low CO2"
	case "203":
		eventDescription = "Super.-Gate Valve"
	case "204":
		eventDescription = "Super.-Low Water Level"
	case "205":
		eventDescription = "Super.-Pump Activation"
	case "206":
		eventDescription = "Super.-Pump Failure"
	// System Troubles
	case "300":
		eventDescription = "Trouble-System Trouble"
	case "301":
		eventDescription = "Trouble-AC Power"
	case "302":
		eventDescription = "Trouble-Low Battery"
	case "303":
		eventDescription = "Trouble-Bad RAM Checksum"
	case "304":
		eventDescription = "Trouble-Bad ROM Checksum"
	case "305":
		eventDescription = "Trouble-System Reset"
	case "306":
		eventDescription = "Trouble-Programming Changed"
	case "307":
		eventDescription = "Trouble-Self Test Failure"
	case "308":
		eventDescription = "Trouble-System Shutdown"
	case "309":
		eventDescription = "Trouble-Battery Test Failure"
	case "310":
		eventDescription = "Trouble-Ground Fault"
	case "311":
		eventDescription = "Trouble-Battery Missing"
	case "312":
		eventDescription = "Trouble-Power Supply Overcurrent"
	case "313":
		eventDescription = "Status-Engineer Reset"
	case "314":
		eventDescription = "Trouble-Primary Power Supply Fail"
	case "316":
		eventDescription = "Trouble-APL System Trouble"
	// Sounder / Relay Troubles
	case "320":
		eventDescription = "Trouble-Sounder/Relay"
	case "321":
		eventDescription = "Trouble-Siren #1"
	case "322":
		eventDescription = "Trouble-Siren #2"
	case "323":
		eventDescription = "Trouble-Alarm Relay"
	case "324":
		eventDescription = "Trouble-Trouble Relay"
	case "325":
		eventDescription = "Trouble-Reversing Relay"
	case "326":
		eventDescription = "Trouble-Notification Appliance Ckt #3"
	case "327":
		eventDescription = "Trouble-Notification Appliance Ckt #4"
	// System Peripheral Troubles
	case "R330":
		eventDescription = "Trouble-System Peripheral"
	case "331":
		eventDescription = "Trouble-Polling Loop Open"
	case "332":
		eventDescription = "Trouble-Polling Loop Short"
	case "333":
		eventDescription = "Trouble-Exp. Module Failure"
	case "334":
		eventDescription = "Trouble-Repeater Fail"
	case "335":
		eventDescription = "Trouble-Printer Paper Out"
	case "336":
		eventDescription = "Trouble-Local Printer"
	case "337":
		eventDescription = "Trouble-Exp. Module DC Loss"
	case "338":
		eventDescription = "Trouble-Exp. Module Low Batt"
	case "339":
		eventDescription = "Trouble-Exp. Module Reset"
	case "341":
		eventDescription = "Trouble-Exp. Module Tamper"
	case "342":
		eventDescription = "Trouble-Exp. Module AC Loss"
	case "343":
		eventDescription = "Trouble-Exp. Self-Test Fail"
	case "344":
		eventDescription = "Trouble-RF Reciever Jam Detection"
	case "345":
		eventDescription = "Trouble-AES Encryption"
	// Communications Troubles
	case "350":
		eventDescription = "Trouble-Communication"
	case "351":
		eventDescription = "Trouble-Phone line #1"
	case "352":
		eventDescription = "Trouble-Phone line #2"
	case "353":
		eventDescription = "Trouble-Radio Transmitter Comm Path"
	case "354":
		eventDescription = "Trouble-Fail to Communicate"
	case "E355":
		eventDescription = "Trouble-Radio Supervision"
	case "356":
		eventDescription = "Trouble-Central Radion Polling"
	case "357":
		eventDescription = "Trouble-Radio Xmitter. VSWR"
	// Protection Loop
	case "370":
		eventDescription = "Trouble-Protection Loop"
	case "371":
		eventDescription = "Trouble-Prot. Loop Open"
	case "372":
		eventDescription = "Trouble-Prot. Loop Short"
	case "373":
		eventDescription = "Trouble-Fire Loop"
	case "374":
		eventDescription = "Alarm-Exit Error"
	case "375":
		eventDescription = "Trouble-Panic Area Trouble"
	case "376":
		eventDescription = "Trouble-Hold-Up Trouble"
	case "377":
		eventDescription = "Trouble-Swinger Trouble"
	case "378":
		eventDescription = "Trouble-Cross Zone Trouble"
	// Sensor
	case "380":
		eventDescription = "Trouble-Sensor Trouble"
	case "381":
		eventDescription = "Trouble-RF Sensor Super."
	case "382":
		eventDescription = "Trouble-RPM Sensor Super."
	case "383":
		eventDescription = "Trouble-Sensor Tamper"
	case "384":
		eventDescription = "Trouble-RF Sensor Battery"
	case "385":
		eventDescription = "Trouble-Smoke Hi Sense"
	case "386":
		eventDescription = "Trouble-Smoke Lo Sense"
	case "-387":
		eventDescription = "Trouble-Intrusion Hi Sense"
	case "-388":
		eventDescription = "Trouble-Intrusion Lo Sense"
	case "389":
		eventDescription = "Trouble-Sensor Test Fail"
	case "391":
		eventDescription = "Trouble-Sensor Watch Fail"
	case "392":
		eventDescription = "Trouble-Drift Comp. Error"
	case "393":
		eventDescription = "Trouble-Maintenance Alert"
	// Open/Close - Refer to event qualifier, "E" = Open "R" = Close
	// qualifer stores the boolean value - false = "E" (open), true = "R" (close)
	case "400":
		if qualifier == true {
			eventDescription = "Closing"
		} else if qualifier == false {
			eventDescription = "Opening"
		}
	case "401":
		if qualifier == true {
			eventDescription = "Closing-User"
		} else if qualifier == false {
			eventDescription = "Opening-User"
		}
	case "402":
		if qualifier == true {
			eventDescription = "Closing-Group User"
		} else if qualifier == false {
			eventDescription = "N/A"
		}
	case "403":
		if qualifier == true {
			eventDescription = "Closing-Auto"
		} else if qualifier == false {
			eventDescription = "Opening-Auto"
		}
	case "404":
		if qualifier == true {
			eventDescription = "Closing-Late"
		} else if qualifier == false {
			eventDescription = "Opening-Late"
		}
	case "405":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "N/A"
		}
	case "406":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Opening-Cancel"
		}
	case "407":
		if qualifier == true {
			eventDescription = "Closing-Remote"
		} else if qualifier == false {
			eventDescription = "Opening-Remote"
		}
	case "408":
		if qualifier == true {
			eventDescription = "Closing-Quick Arm"
		} else if qualifier == false {
			eventDescription = "N/A"
		}
	case "409":
		if qualifier == true {
			eventDescription = "Closing-Keyswitch"
		} else if qualifier == false {
			eventDescription = "Opening-Keyswitch"
		}
	case "435":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "ACCESS-User"
		}
	case "436":
		if qualifier == true {
			eventDescription = ""
		} else if qualifier == false {
			eventDescription = "ACCESS-Irregular Access - User"
		}
	case "441":
		if qualifier == true {
			eventDescription = "Closing-Armed Stay"
		} else if qualifier == false {
			eventDescription = "Opening-Armed Stay"
		}
	case "442":
		if qualifier == true {
			eventDescription = "Closing-Keysw. Armed Stay"
		} else if qualifier == false {
			eventDescription = "Opening-Keysw. Armed Stay"
		}
	case "450":
		if qualifier == true {
			eventDescription = "Closing-Exception"
		} else if qualifier == false {
			eventDescription = "Opening-Exception"
		}
	case "451":
		if qualifier == true {
			eventDescription = "Closing-Early-User"
		} else if qualifier == false {
			eventDescription = "Opening-Early-User"
		}
	case "452":
		if qualifier == true {
			eventDescription = "Closing-Late-User"
		} else if qualifier == false {
			eventDescription = "Opening-Late-User"
		}
	case "453":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Trouble-Fail to open"
		}
	case "454":
		if qualifier == true {
			eventDescription = "Trouble-Fail to close"
		} else if qualifier == false {
			eventDescription = "N/A"
		}
	case "455":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Trouble-Auto Arm Fail"
		}
	case "456":
		if qualifier == true {
			eventDescription = "Closing-Partial arm"
		} else if qualifier == false {
			eventDescription = "N/A"
		}
	case "457":
		if qualifier == true {
			eventDescription = "Closing-Exit Error-User"
		} else if qualifier == false {
			eventDescription = "N/A"
		}
	case "458":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Opening-User on Prem"
		}
	case "459":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Trouble-Recent Close - User"
		}
	case "461":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Access-Wrong code entry"
		}
	case "462":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Access-Legal code entry"
		}
	case "463":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Status-Re Arm After Alarm - User"
		}
	case "464":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Status-Auto Arm Time Ext."
		}
	case "465":
		if qualifier == true {
			eventDescription = "N/A"
		} else if qualifier == false {
			eventDescription = "Status-Panic Alarm Reset"
		}
	case "466":
		if qualifier == true {
			eventDescription = "Access-Service off Prem."
		} else if qualifier == false {
			eventDescription = "Access-Service on Prem."
		}
	// Remote Access
	case "411":
		eventDescription = "Remote-Callback Requested"
	case "412":
		eventDescription = "Remote-Successful Access"
	case "413":
		eventDescription = "Remote-Unsuccessful Access"
	case "414":
		eventDescription = "Remote-System Shutdown"
	case "415":
		eventDescription = "Remote-Dialer Shutdown"
	case "416":
		eventDescription = "Remote-Successful Upload"
	// Access Control
	case "421":
		eventDescription = "Access-Access Denied"
	case "422":
		eventDescription = "Access-Access Granted"
	case "423":
		eventDescription = "Panic-Forced Access"
	case "424":
		eventDescription = "Access-Egress Denied"
	case "425":
		eventDescription = "Access-Egress Granted"
	case "426":
		eventDescription = "Access-Door Propped Open"
	case "427":
		eventDescription = "Access-ACS Point DSM Trouble"
	case "428":
		eventDescription = "Access-ACS Point RTE Trouble"
	case "429":
		eventDescription = "Access-ACS Prog. Entry"
	case "430":
		eventDescription = "Access-ACS Prog. Exit"
	case "431":
		eventDescription = "Access-ACS Threat Level Change"
	case "432":
		eventDescription = "Access-ACS Relay/Trigger Fail"
	case "433":
		eventDescription = "Access-ACS RTE Shunt"
	case "434":
		eventDescription = "Access-ACS DSM Shunt"
	// System Disables
	case "501":
		eventDescription = "Disable-Access Reader Disabled"
	// Sounder / Relay Disables
	case "520":
		eventDescription = "Disable-Sounder/Relay"
	case "521":
		eventDescription = "Disable-Bell/Siren #1"
	case "522":
		eventDescription = "Disable-Bell/Siren #2"
	case "523":
		eventDescription = "Disable-Alarm Relay"
	case "524":
		eventDescription = "Disable-Trouble Relay"
	case "525":
		eventDescription = "Disable-Reversing Relay"
	case "526":
		eventDescription = "Disable-Notification Appliance Ckt #3"
	case "527":
		eventDescription = "Disable-Notification Appliance Ckt #4"
	// System Peripheral Disables
	case "531":
		eventDescription = "Super.-Module Added"
	case "532":
		eventDescription = "Super.-Module Removed"
	// Communication Disables
	case "551":
		eventDescription = "Disable-Dialer Disabled"
	case "552":
		eventDescription = "Disable-Radio Disabled"
	case "553":
		eventDescription = "Disable-Up/download Disabled"
	// Bypasses
	case "570":
		eventDescription = "Bypass-Zone Bypass"
	case "571":
		eventDescription = "Bypass-Fire Bypass"
	case "572":
		eventDescription = "Bypass-24 Hour Bypass"
	case "573":
		eventDescription = "Bypass-Burg. Bypass"
	case "574":
		eventDescription = "Bypass-Group Bypass"
	case "575":
		eventDescription = "Bypass-Swinger Bypass"
	case "576":
		eventDescription = "Access-ACS Zone Shunt"
	case "577":
		eventDescription = "Access-ACS Point Bypass"
	case "578":
		eventDescription = "Bypass-Vault Bypass"
	case "579":
		eventDescription = "Bypass-Vent Zone Bypass"
	// Test / Misc
	case "601":
		eventDescription = "Test-Manual"
	case "602":
		eventDescription = "Test-Periodic"
	case "603":
		eventDescription = "Test-Periodic Radio"
	case "604":
		eventDescription = "Test-Fire Walk Test"
	case "605":
		eventDescription = "Test-Fire Walk Test Status Report"
	case "606":
		eventDescription = "Listen-Listen-in Acitve"
	case "607":
		eventDescription = "Test-Walk Test Mode"
	case "608":
		eventDescription = "Test-System Trouble Present"
	case "609":
		eventDescription = "Listen-Video Xmitter Active"
	case "611":
		eventDescription = "Test-Point Tested OK"
	case "612":
		eventDescription = "Test-Point Not Tested"
	case "613":
		eventDescription = "Test-Intrusion Zone Walk Test"
	case "614":
		eventDescription = "Test-Fire Zone Walk Test"
	case "615":
		eventDescription = "Test-Panic Zone Walk Test"
	case "616":
		eventDescription = "Test-Service Request"
	// Event Log
	case "621":
		eventDescription = "Trouble-Event Log Reset"
	case "622":
		eventDescription = "Trouble-Event Log 50% Full"
	case "623":
		eventDescription = "Trouble-Event Log 90% Full"
	case "624":
		eventDescription = "Trouble-Event Log Overflow"
	case "625":
		eventDescription = "Trouble-Time/Date Reset"
	case "626":
		eventDescription = "Trouble-Time/Date Invalid"
	case "627":
		eventDescription = "Trouble-Program mode Entry"
	case "628":
		eventDescription = "Trouble-Program mode Exit"
	// Scheduling
	case "630":
		eventDescription = "Trouble-Schedule Changed"
	case "631":
		eventDescription = "Trouble-Exception Schedule Changed"
	case "632":
		eventDescription = "Trouble-Access Schedule Changed"
	// Personnel Monitoring
	case "641":
		eventDescription = "Trouble-Senior Watch Trouble"
	case "642":
		eventDescription = "Status-Latch-key Supervision"
	default:
		eventDescription = "UNRECOGNIZED MESSAGE"
	}
	return eventDescription
}
