# Instrução para inicializar o WTS Lead

* Configure o .env utilizando o .env_example como referência.
* Execute esse comando na pasta base do projeto:
```powershell
cd code/frontend && npm i -y
```
* Abra uma nova aba do terminal e execute:
```powershell
cd code/backend && go mod tidy
```

## Inicializar servidor
Comando para inicializar o servidor na pasta base
```powershell
cd code/backend && go run .
```