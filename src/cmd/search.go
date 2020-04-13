package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/yibozhuang/dockerhub-go/src/dockerhub"
)

var page uint32
var limit uint32

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Docker Hub for an image",
	Long: `Search Docker Hub (https://hub.docker.com)
for images by the name specified. For example:

dockerhub search mysql

By default it will only display the first 10 entries from the search.
You can specify the starting page or increase the number of entries per page

dockerhub search mysql --limit=100
dockerhub search mysql --page=2 --limit=100
`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := cmd.Flags().GetUint32("page")
		limit, _ := cmd.Flags().GetUint32("limit")
		search := args[0]

		output, err := dockerhub.Search(search, page, limit)
		if err != nil {
			er(err)
		}

		DisplaySearchResults(output)
	},
}

func init() {
	searchCmd.Flags().Uint32VarP(&page, "page", "p", 1, "The starting page number")
	searchCmd.Flags().Uint32VarP(&limit, "limit", "l", 10, "The number of search output to display")

	rootCmd.AddCommand(searchCmd)
}

// DisplaySearchResults outputs the search result to the user in a tabuler format
func DisplaySearchResults(results dockerhub.SearchAPIResponse) {
	const RFC3339FullDate = "2006-01-02"

	headers := []string{"Name", "Publisher", "Pull Count", "Type", "Updated At"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	table.SetRowLine(true)
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, image := range results.Images {
		info := []string{image.Name, image.Publisher.Name, image.PullCount, image.Type, image.UpdatedAt.Format(RFC3339FullDate)}
		table.Append(info)
	}

	table.Render()
}
