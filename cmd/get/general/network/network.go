package network

import (
	"github.com/spf13/cobra"
)

// NetCmd - Root for all organization CLI commands.
var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "",
	Long:    "Entrypoint for network subcommands.",
}

// init - Entrypoint for NetCmd sub-commands.
func init() {
	NetCmd.AddCommand(alertconfig)
	NetCmd.AddCommand(clients)
	NetCmd.AddCommand(clientpolicy)
	NetCmd.AddCommand(clientsplashauthorization)
	NetCmd.AddCommand(devices)
	NetCmd.AddCommand(firmwareupgrades)
	NetCmd.AddCommand(floorplans)
	NetCmd.AddCommand(floorplan)
	NetCmd.AddCommand(grouppolicies)
	NetCmd.AddCommand(grouppolicy)
	NetCmd.AddCommand(merakiauthusers)
	NetCmd.AddCommand(merakiauthuser)
	NetCmd.AddCommand(mqttbrokers)
	NetCmd.AddCommand(mqttbroker)
	NetCmd.AddCommand(netflow)
	NetCmd.AddCommand(channelutilization)
	NetCmd.AddCommand(piirequests)
	NetCmd.AddCommand(piirequest)
	NetCmd.AddCommand(smdevices)
	NetCmd.AddCommand(smowners)
	NetCmd.AddCommand(settings)
	NetCmd.AddCommand(snmp)
	NetCmd.AddCommand(syslog)
	NetCmd.AddCommand(trafficanalysis)
	NetCmd.AddCommand(trafficshaping)
	NetCmd.AddCommand(dscp)
	NetCmd.AddCommand(httpservers)
	NetCmd.AddCommand(httpserver)
	NetCmd.AddCommand(webhooktest)
	NetCmd.AddCommand(bluetoothclients)
	NetCmd.AddCommand(bluetoothclient)
	NetCmd.AddCommand(clienttraffichistory)
	NetCmd.AddCommand(clientusagehistory)
	NetCmd.AddCommand(networkclients)
	NetCmd.AddCommand(clientidentifier)
	NetCmd.AddCommand(environmentalevents)
	NetCmd.AddCommand(events)
	NetCmd.AddCommand(splashloginattempts)
	NetCmd.AddCommand(traffic)
}
