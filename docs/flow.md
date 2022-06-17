# Key Flows

**[Mermaid Sequence Diagram Reference](https://mermaid-js.github.io/mermaid/#/sequenceDiagram)**


## Starting Server

Open Terminal and execute.

> aws-env is wrapper to aws-vault tool for securely stoding AWS Credentials Locally

![Overview of Command Line App - aws-hub](images/aws-hub.png)

```
aws-env go run main.go start
```

```mermaid
sequenceDiagram
    actor User
    User ->> main : aws-env go run main.go start
    main ->> app.hub: Execute()
    opt app
        app.hub ->> app.hub: hub:= New()
        opt app.hub
            app.hub  ->> + app.flag: hub.setUpFlags()
            app.hub  ->>    app.cmd: hub.setUpCommands()
            app.cmd  ->> + app.cmd: cmd.StartCommand
        end
        app.hub ->> app.hub: app.app.setUpOutput()
        app.hub ->> app.cmd: app.cli.Run(args)
        Note right of app.cmd: invokes StartCommand
        app.cmd ->> app.server: server.Start
        Note right of app.server: Server Started!
    end
```

Test Execution For the Flow

```
go test -timeout 5s  -coverprofile=coverage/coverage.out  github.com/rajasoun/aws-hub/app/... -v
```

or

Formated Output

```
gotestsum --format testname -- -coverprofile=coverage/coverage.out github.com/rajasoun/aws-hub/app/...
```
