package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yibozhuang/dockerhub-go/src/dockerhub"
)

// tagsCmd represents the tags command
var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "Get list of tags for a particular image from Docker Hub",
	Long: `Get list of tags for a particular image from Dokcer Hub
(https://hub.docker.com). For example:

dockerhub tags mysql

By default it will only display the first 10 tags for the given image.
You can specify the starting page or increase the number of entries per page

dockerhub tags mysql --limit=100
dockerhub tags mysql --page=2 --limit=100
`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		page, _ := cmd.Flags().GetUint32("page")
		limit, _ := cmd.Flags().GetUint32("limit")
		image := args[0]

		output, err := dockerhub.Tags(image, page, limit)
		if err != nil {
			er(err)
		}

		DisplayTagsResults(output)
	},
}

func init() {
	tagsCmd.Flags().Uint32VarP(&page, "page", "p", 1, "The starting page number")
	tagsCmd.Flags().Uint32VarP(&limit, "limit", "l", 10, "The number of tags to display")

	rootCmd.AddCommand(tagsCmd)
}

// DisplayTagsResults outputs the tags list results to the user
func DisplayTagsResults(results dockerhub.TagAPIResponse) {
	const RFC3339FullDate = "2006-01-02"

	headers := []string{"Name", "Hash", "Updated At"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)

	table.SetRowLine(true)
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, tag := range results.Tags {
		hash := ""
		if len(tag.Details) > 0 {
			hash = tag.Details[0].Hash
		}

		info := []string{tag.Name, hash, tag.LastUpdated.Format(RFC3339FullDate)}
		table.Append(info)
	}

	table.Render()
}
