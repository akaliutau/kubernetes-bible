# Example of using Google Cloud Build

Like all Cloud Build steps, this consists of a set of YAML key-value pairs:
1) _id_ gives a human-friendly label to the build step.
2) _dir_ specifies the subdirectory of the Git repo to work in.
3) _name_ identifies the container to run for this step.
4) _entrypoint_ specifies the command to run in the container, if not the default.
5) _args_ gives the necessary arguments to the entrypoint command

Substitution Variables

In order to make Cloud Build pipeline files reusable and flexible we use variables, or what Cloud Build calls substitutions. 
Anything that begins with a $ will be substituted.
User-defined substitutions in Cloud Build must begin with an underscore and use only uppercase letters/numbers.
(must be created manually or via terraform at https://console.cloud.google.com/cloud-build/triggers)

In our example 2 vars are used:

1) _REGION  - it should match the GCP availability region with GKE cluster, the same for Artifact Registry repo.
2) _CLOUDSDK_CONTAINER_CLUSTER  - is the name of GKE cluster.

Note, the git commit should be tagged, since we are using TAG_NAME variable in pipeline.

Some other pre-defined vars are:

* PROJECT_ID