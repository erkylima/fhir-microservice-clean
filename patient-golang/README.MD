# Patient-service
Microservice para cadastro e gestão de profissionais


## Arquitetura
### /docs
Documentação do sistema. Exemplo: documentação da api, godocs, etc...

### /app
Pasta a onde deve ficar o executavel do aplicativo: main.go.
Exemplo: ```/app/<app_name>```, ```/app/profissional_service```

### /interactor
Pasta a onde ficará os arquivos relacionados aos handlers http da aplicação: ```/interactor/<seus_handlers_files>/```.

### /model
Aqui temos todas as estruturas relacionadas a modelagem dos dado que serão persistidos.

### /provider
Pasta onde ficará os arquivos. Ex:  ```/service/create_patient.go``` ou ```/service/patient_db.go```.

### /service
Pasta onde ficará todos os services. Exemplo: ```/service/validade.go```, ```/service/isValid.go```, ```/service/checkToken.go```.

## Deploy

Para buildar a image docker

```sh
docker build . -t patient-service
```
