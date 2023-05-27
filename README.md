# argo-helm-rapport
This project aims to work as a pipeline tool to show manifest changes between the helm chart versions. 

# Requirements
- Works as a pipeline tool
- No vulnarabilites
- Detect the charts in the repository
- Render charts for both of the version
- Compare changes
- It has to be as fast as possible

# User Story Mapping
- User can set the charts repository OK
- User can set the chart type as private and public OK
- If the chart is private, username and password has to be set too
- User can set the charts to be monitored
- Tool detects all argo apps that has this chart in it
- Tool authenticates the chart repositories, fails if it can not reach to it and then ignores that private ones
- User pass the versions to tool via the pipeline
- Tool templates the chart with the given informations in the argo app definition to a single file
- Tool compares two different rendered template
- Print out the differences
