# jiraAdmin

There are not documented ways to manage users as an admin in the Atlassian Cloud stack.  This is a tool written in go to deactivate users from Atlassian cloud instances if they have not accessed the instance in more than 60 days.

## Usage

```
Usage:
  jiraAdmin [flags]
 
Flags:
  -c, --cookie string   jira cloud session cookie
  -d, --dryrun          enable dryrun
  -h, --help            help for jiraAdmin
  -s, --site string     jira cloud site id
  -u, --url string      jira cloud url
  ```

Cookie, site, and url are all required parameters.  The cookie header can be found using developer tools looking at the entries within cookie, specifically the "cloud.session.token" entry.  The site id can be derived from your site adminstration page https://<site_url>/admin/s/<siteID>/users.  Finally the site url is just the base url of your instance https://<site_url>.

### Disclaimer
These are not documented APIs and may be subject to change at any time breaking this tool.
