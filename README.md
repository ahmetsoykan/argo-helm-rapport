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
- User can set the charts repository OK - Repo
- User can set the chart type as private and public OK - Repo
- If the chart is private, username and password has to be set too OK - Repo
- Tool detects all argo apps that has this chart in it OK - Detects
- User can set the charts to be monitored OK - Watch
- User pass the versions to tool via the pipeline OK - Watch
- Tool authenticates the chart repositories
- Tool templates the chart with the given informations in the argo app definition to a single file
    - helm render
    - helm render + kustomize patch
- Tool compares two different rendered template
- Print out the differences
