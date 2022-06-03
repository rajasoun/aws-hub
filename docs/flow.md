# Key Flows

**[Mermaid Sequence Diagram Reference](https://mermaid-js.github.io/mermaid/#/sequenceDiagram)**


## Starting Server

Open Terminal and execute.
> aws-env is wrapper to aws-vault tool for securely stoding AWS Credentials Locally 

```
aws-env go run main.go start 
```

```mermaid
sequenceDiagram
    autonumber
    actor User
    User->>main : aws-env go run main.go start 
    main->>app.hub: Execute()
    opt app
        app.hub->>app.hub: hub:= NewApp()
        opt app.hub
            app.hub->>app.flag:hub.setUpFlags()
            app.hub->>app.cmd: hub.setUpCommands() 
            app.cmd->>app.cmd: cmd.StartCommandHandler
        end
        app.hub->>app.hub: app.app.setUpOutput()
        app.hub->>app.cmd: app.cli.Run(args)  
        Note right of app.cmd: invokes StartCommandHandler
        app.cmd->>app.server: server.Start
         Note right of app.server: Server Started! 
    end
```