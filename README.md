# argo-helm-rapport
This project aims to work as a pipeline tool to show manifest changes between the helm chart versions. 

# Requirements
- Works as a pipeline tool
- No vulnerabilites
- Detects the charts in the repository
- Render charts and templates to a file
- Compare changes
- Reasonable fast

# User Story Mapping
- User can set the charts repository OK - [ Repo Pkg ]
- User can set the chart type as private and public OK - [ Repo Pkg ]
- If the chart is private, username and password has to set OK - [ Repo Pkg ]
- Tool detects all argo apps that has this chart in it OK - [ Detects Pkg ]
- User can set the charts to be monitored OK - [ Watch Pkg ] 
- Tool authenticates the chart repositories OK - [ Chart Pkg ]
- Tool templates the chart with the given informations in the argo app definition to a single file
    - Charts can only have two version OK [ Chart pkg ]
    - helm render OK [ Chart pkg ]
    - When the versions are the same, ignore rendering and comparing OK [ Chart pkg ]
    - combine two values.yaml into one OK [ Chart pkg]
    - When the combined values are the same, ignore rendering and comparing OK [ Chart pkg ]
    - compare combined values yamls for the same version OK [ Detect and Chart pkg]
    - helm render + kustomize patch OK [ Chart pkg]
    - Render should fail if the chart repository is unknown OK [ Chart pkg ]
    - Compare should fail if there's no file mapping OK [ Chart pkg]
- Tool compares two different rendered template OK - [ Chart pkg]
- Print out the differences OK [ Chart pkg ]


# Usage
- To start to use this tool see the example commands below:
    - chart repository authentication
        ##### ./main repo add -n stable --host https://charts.helm.sh/stable
    - this is the repository that we will check from argo applications
        ##### ./main watch chart -n nginx-ingress
    - search under the the app for apps folder and its values.yaml (run x2 for your feature branch and master)
        ##### ./main detect -p ./cmd/detect/example-folder
    - templates the files
        ##### ./main chart render
    - compare the files
        ##### ./main chart compare
