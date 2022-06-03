[Mermaid Sequence Diagram Reference](https://mermaid-js.github.io/mermaid/#/sequenceDiagram)

```mermaid
sequenceDiagram
    autonumber
    actor User
    User->>main : aws-env go run main.go start 
    main->>app.hub: Execute()
    opt app
        app.hub->>app.hub: app:= NewApp()
        app.hub->>app.hub: app.setUpCommand()
    end
    Note right of hub.server: Server Started!
    app.hub->>hub.server: app.cli.Run(args)   
```