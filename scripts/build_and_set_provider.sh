#!/bin/bash

# This script build this provider and install it into terraform plugin directory.
# please set build task in .vscode/tasks.json like following.

<< COMMENTOUT
{
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
    "version": "1.0.0",
    "tasks": [
        {
            "label": "go build",
            "type": "shell",
            "command": "/bin/bash",
            "args": [
                "-x",
                "${workspaceRoot}/scripts/build_and_set_provider.sh",
                "~/dev//terraform/.terraform/plugins/darwin_amd64" <-- Change here depending on your terraform path.
            ],
            "group": {
                "kind": "build",
                "isDefault": true
            }
        }
    ]
}
COMMENTOUT

go build
cp terraform-provider-ecl "$1"