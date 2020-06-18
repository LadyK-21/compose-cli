package cmd

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/docker/api/client"
	"github.com/docker/api/formatter"
)

// InspectCommand inspects into containers
func InspectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect",
		Short: "Inspect containers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInspect(cmd.Context(), args[0])
		},
	}

	return cmd
}

func runInspect(ctx context.Context, id string) error {
	c, err := client.New(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot connect to backend")
	}

	container, err := c.ContainerService().Inspect(ctx, id)
	if err != nil {
		return err
	}

	j, err := formatter.ToStandardJSON(container)
	if err != nil {
		return err
	}
	fmt.Println(j)

	return nil
}