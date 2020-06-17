package cli

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/irismod/htlc/types"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	htlcQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the HTLC module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	htlcQueryCmd.AddCommand(flags.GetCommands(
		GetCmdQueryHTLC(queryRoute, cdc),
	)...)

	return htlcQueryCmd
}

// GetCmdQueryHTLC implements the query HTLC command.
func GetCmdQueryHTLC(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "htlc [hash-lock]",
		Short: "Query an HTLC",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details of an HTLC with the specified hash lock.

Example:
$ %s query htlc htlc <hash-lock>
`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			hashLock, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			bz, err := cdc.MarshalJSON(types.QueryHTLCParams{HashLock: hashLock})
			if err != nil {
				return err
			}

			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryHTLC)
			res, _, err := cliCtx.QueryWithData(route, bz)
			if err != nil {
				return err
			}

			var htlc types.HTLC
			if err := cdc.UnmarshalJSON(res, &htlc); err != nil {
				return err
			}

			return cliCtx.PrintOutput(htlc)
		},
	}
}
