```mermaid

sequenceDiagram
	actor User
		User  ->>  main : aws-env go run main.go start
		main  ->>  app.hub : Execute()
		app.hub  ->>  app.hub : NewApp()

```
