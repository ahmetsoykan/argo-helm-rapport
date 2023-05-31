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
- User can set the charts repository OK - Repo Pkg
- User can set the chart type as private and public OK - Repo Pkg
- If the chart is private, username and password has to be set too OK - Repo Pkg
- Tool detects all argo apps that has this chart in it OK - Detects Pkg
- User can set the charts to be monitored OK - Watch Pkg
- User pass the versions to tool via the pipeline OK - Watch Pkg
- Tool authenticates the chart repositories OK - Helm Pkg
- Tool templates the chart with the given informations in the argo app definition to a single file
    - helm render OK
    - helm render + kustomize patch
- Tool compares two different rendered template OK
- Print out the differences OK
- Combine some unneccessary functions

# Usage
- To start to use this tool see the example commands below:
    ## chart repository authentication
    - ./main repo add -n stable --host https://charts.helm.sh/stable
    ## this is the repository that we will check from argo applications
    - ./main watch chart -n nginx-ingress
    ## search under the the app for apps folder and its values.yaml
    - ./main detect -p ./internals/detect/example-folder
    ## templates the files
    - ./main chart render
    ## compare the files
    - ./main chart compare
