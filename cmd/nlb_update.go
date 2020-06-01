package cmd

import (
	"fmt"

	apiv2 "github.com/exoscale/egoscale/api/v2"
	"github.com/spf13/cobra"
)

var nlbUpdateCmd = &cobra.Command{
	Use:   "update <ID>",
	Short: "Update a Network Load Balancer",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmdExitOnUsageError(cmd, "missing arguments")
		}
		return cmdCheckRequiredFlags(cmd, []string{"zone"})
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var nlbID = args[0]

		zone, err := cmd.Flags().GetString("zone")
		if err != nil {
			return err
		}

		nlb, err := cs.GetNetworkLoadBalancer(
			apiv2.WithEndpoint(gContext, apiv2.NewReqEndpoint(gCurrentAccount.Environment, "")),
			zone,
			nlbID)
		if err != nil {
			return err
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if cmd.Flags().Changed("name") {
			nlb.Name = name
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		if cmd.Flags().Changed("description") {
			nlb.Description = description
		}

		if _, err := cs.UpdateNetworkLoadBalancer(
			apiv2.WithEndpoint(gContext, apiv2.NewReqEndpoint(gCurrentAccount.Environment, "")),
			zone,
			nlb); err != nil {
			return fmt.Errorf("unable to update Network Load Balancer: %s", err)
		}

		if !gQuiet {
			return output(showNLB(nlb.ID, zone))
		}

		return nil
	},
}

func init() {
	nlbUpdateCmd.Flags().StringP("zone", "z", "", "Network Load Balancer zone")
	nlbUpdateCmd.Flags().String("name", "", "service name")
	nlbUpdateCmd.Flags().String("description", "", "service description")
	nlbCmd.AddCommand(nlbUpdateCmd)
}
