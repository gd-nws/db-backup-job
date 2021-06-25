## Database Backup Job

A Golang script to fetch a database dump, and then upload
a zipped dump to S3. 

## Setup

Use the dev container provided in `./devcontainer/`. You will need to update
the run args in `./devcontainer.json` to update the network the container
should run in. This needs to match the network of your mongo db container.
```
"runArgs": [ 
  "--cap-add=SYS_PTRACE", 
  "--security-opt", 
  "seccomp=unconfined", 
  "--network=<YOUR_NETWORK>", 
]
```

### Environment
Copy and fill out the `.env.example` file.