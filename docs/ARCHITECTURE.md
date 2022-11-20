# Framework Architecture

## Main Problems

- Need to create a monitor for potentially thousands of websites.
- If we need to monitor potentially thousands of websites, each website will require a different HTTP request flow to scrape job application information.
  - Is there a way we can automate this?
  - How can we manage all of these request flows?
    - Thousands of modules...?
    - Need to maintain an internal list of supported websites.
    - Then, when starting monitor tasks, iterate through the list of websites and inject the "fetch `[]JobApplications`" callback into the task to use.
    - https://medium.com/@vicky.kurniawan/go-call-a-function-from-string-name-30b41dcb9e12
- What do we do if a request flow suddenly fails?
  - Need to be able to detect bad request flows
  - Then throw a webhook to our discord channel to notify ourselves
- How can we monitor user growth?
  - Grafana? --> should be on backend
- What job application information should we specifically scrape?
  - SWE/CS positions only?
  - Location?
  - Skills (need to compile by hand)
- Need to be able to detect when responses to our monitor requests change quickly
- How should we send potentially thousands of HTTP requests efficiently?
  - If we have 1000+ websites to monitor, then if we space it so that we send ~4/second, then we can effectively monitor 1k websites with relatively low burst over a 5 minute span
    - And then we can repeat it pretty easily
- Error detection?
  - Ideally, should unit test
  - And when encounters any error that we don't know about, then shutdown
    - If 500/501, then keep repeating until website is back up
    - Set a timer and alert if the website doesn't come back up after 20 mintues
  - **Need to be able restart a monitor task through an API request**
- Need to create an API to manage monitor

```
Hierarchy

MonitorsManager
    MonitorTask

MonitorsManager should be able to start/stop monitor tasks.
    Let each MonitorTask be its own context within MonitorsManager

MonitorsManager should know what websites are currently being monitored and which are not.
- Should be able to call MonitorsManager.start(website) to start a monitor task for the website if it doesn't exist already
- Should also be able to call MonitorsManager.stop(website) to stop the monitor task for the website if it doesn't exist already
- There should only be one MonitorTask for each website
    - Manage a map of pointers to MonitorTasks to remotely stop them if possible.
```
