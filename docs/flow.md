[Mermaid Sequence Diagram Reference](https://mermaid-js.github.io/mermaid/#/sequenceDiagram)

```mermaid
sequenceDiagram
    autonumber
    actor User
    User->>main : aws-env go run main.go start 
    main->>hub.app: Execute()
    opt app
        hub.app->>hub.app: app:= NewApp()
        hub.app->>hub.app: app.setUpCommand()
    end
    Note right of hub.server: Server Started!
    hub.app->>hub.server: app.cli.Run(args)   
```