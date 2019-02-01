package commands

import (
	"encoding/json"
  "fmt"
	"github.com/spf13/cobra"
	"time"
)

var Cookie string
var DryRun bool
var Site string
var Url string

func init() {
	RootCmd.Flags().StringP("cookie", "c", "", "jira cloud session cookie")
	RootCmd.MarkFlagRequired("cookie")
	RootCmd.Flags().StringP("site", "s", "", "jira cloud site id")
	RootCmd.MarkFlagRequired("site")
	RootCmd.Flags().StringP("url", "u", "", "jira cloud url")
	RootCmd.MarkFlagRequired("url")
	RootCmd.Flags().BoolP("dryrun", "d", false, "enable dryrun")
}

var RootCmd = &cobra.Command{
	Use: "jiraAdmin",
	Run: func(cmd *cobra.Command, args []string) {
		cookie := cmd.Flag("cookie").Value.String()
		url := cmd.Flag("url").Value.String()
		siteID := cmd.Flag("site").Value.String()
		dryRunEnabled, derr := cmd.Flags().GetBool("dryrun")
		if derr != nil {
			fmt.Println("error:", derr)
		}

		now := time.Now()
		allUsers := MakeRequest("GET", url+"/rest/api/3/user/search?startAt=0&maxResults=2000&username=%", cookie)
		var users []userObject
		err := json.Unmarshal(allUsers, &users)
		if err != nil {
			fmt.Println("error:", err)
		}
		for k := range users {
			userAdminInterface := MakeRequest("GET", url+"/gateway/api/adminhub/um/site/"+siteID+"/users/"+users[k].AccountID, cookie)
			var userAdmin userObjectAdmin
			err2 := json.Unmarshal(userAdminInterface, &userAdmin)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			timeDiff := now.Sub(userAdmin.Presence).Hours()
			if (timeDiff > 1440 && userAdmin.System != true) {
				fmt.Printf("User: %s \n", userAdmin.DisplayName)
				fmt.Printf("LastLog: %s \n", userAdmin.Presence)
				fmt.Printf("Email: %s \n", userAdmin.Email)
				fmt.Printf("id: %s \n", userAdmin.AccountID)
				if dryRunEnabled != true {
					fmt.Printf("Deactivating...\n")
					if userAdmin.Presence.IsZero() {
						MakeRequest("DELETE", url+"/gateway/api/adminhub/um/site/"+siteID+"/users/"+userAdmin.AccountID, cookie)
					}
					MakeRequest("POST", url+"/gateway/api/adminhub/um/site/"+siteID+"/users/"+userAdmin.AccountID+"/deactivate", cookie)
					fmt.Printf("\n")
				}
			}
		}
	},
}