package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yibozhuang/dockerhub-go/src/dockerhub"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get detailed info of an image from Docker Hub",
	Long: `Get detailed info of an image name from Docker Hub
(https://hub.docker.com). For example:

dockerhub info mysql
dockerhub info prom/prometheus
`,
	Args: cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]

		output, err := dockerhub.Info(image)
		if err != nil {
			er(err)
		}

		DisplayImageInfo(output)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

// DisplayImageInfo outputs the image detail info to the user
func DisplayImageInfo(info dockerhub.DockerInfoImage) {
	if info.Name == "" {
		fmt.Println("No such image found")
	} else {
		const RFC3339FullDate = "2006-01-02"

		table := tablewriter.NewWriter(os.Stdout)

		table.SetAlignment(tablewriter.ALIGN_LEFT)

		table.Append([]string{"Name", info.Name})
		table.Append([]string{"NameSpace", info.NameSpace})
		table.Append([]string{"Type", info.Type})
		table.Append([]string{"Pull Count", strconv.Itoa(info.PullCount)})
		table.Append([]string{"Last Updated", info.LastUpdated.Format(RFC3339FullDate)})
		table.Append([]string{"Description", info.Description})

		table.Render()

		dockerhub.RenderMarkdown(info.FullDescription, os.Stdout)
	}
}
