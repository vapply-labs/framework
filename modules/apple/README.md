# Apple Careers Documentation

Main Endpoint: [https://jobs.apple.com/en-us/search?key=software%20data%20intern%20software%25252520engineering&location=united-states-USA&page=1&sort=newest](https://jobs.apple.com/en-us/search?key=software%20data%20intern%20software%25252520engineering&location=united-states-USA&page=4&sort=newest)

## Notes

- "Newest" on the Apple Careers website ignores the filters you set...
- Jobs seem to be posted at mass in batches.
  - Take up around 2 pages
- Each job displays:
  - Role Name
  - Location
  - Date Posted
  - Department

The best approach for Apple seems to be:

1. Detect a batch of new jobs with [https://jobs.apple.com/en-us/search?key=software%20data%20intern%20software%25252520engineering&location=united-states-USA&page=1&sort=newest](https://jobs.apple.com/en-us/search?key=software%20data%20intern%20software%25252520engineering&location=united-states-USA&page=1&sort=newest)
   1. For each page, parse all of the jobs on the page with the filters.
2. Keep iterating through pages until the the data changes.
3. Return the jobs.
