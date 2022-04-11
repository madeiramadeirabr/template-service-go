# Guidelines para projetos
Neste documento você encontra orientações para o seu projeto ficar de acordo com as `guidelines` 
definidas pelo time de arquitetura, com o objetivo de deixar o seu projeto bem completo.

## Tópicos
* Github templates
* Github actions/workflows
* Github docs
* Sonar
* OpenAPI/Swagger
* RESTful e HATEOS
* Healthcheck
* Docker 
* Editorconfig 
* Serverless
* Logs e monitoramento
* Scripts de automação
* Scripts de configuração de ambiente
* Terraform
* Testes
* Arquivos de CI/CD para AWS
* Meta-arquivos




## Github templates
Incluir templates padronizados para:
* Pull Request
* Issues
  * Bug Report
  * Documentation Request
  * Feature Request

## Github actions/workflows
Incluir scripts para as seguintes ações:
* Lint
* Unit Test
* Component Test
* Sonar
* Versioning

## Github docs 
Incluir arquivos focados ao uso da comunidade:
* README.md
* CHANGELOG.md
* CODE_OF_CONDUCT.md
* CONTRIBUTING.md
* LICENSE.md

### README.md
É importante que este arquivo contenha os seguintes passos no mesmo:
* Descrição breve
* Requesitos
* Funcionalidades
* Instalação
* Execução
* Execução de testes
* Execução ferramentas para desenvolvimento (opcional)

### Outros arquivos
Seguir o padrão das arquiteturas de referência, você pode copiar os mesmos.

## Sonar
O projeto deve possuir o arquivo de configuração do sonar e action para análise.
* sonar.properties
* .github/workflows/sonar.yml

## OpenAPI/Swagger
Quando o projeto for uma API, é necessario que o mesmo tenho a documentação no padrão OpenAPI.

Considerar o uso de:
* UI
* Schemas
* Rotas

## RESTful e HATEOS
Quando o projeto for uma API, é desejável que o mesmo implemente as definições do padrão RESTful.
Melhor ainda se puder aplicar conceitos de HATEOS.

Para mais detalhes ver:
* [MadeiraMadeira - Guidelines RESTful e HATEOS](https://madeiramadeira.atlassian.net/wiki/spaces/CAR/pages/2244149708/WIP+-+Guidelines+-+RESTful+e+HATEOS)
* [Designing-a-Beautiful-REST%2BJSON-API.pdf](https://docs.huihoo.com/apache/apachecon/us2014/Designing-a-Beautiful-REST%2BJSON-API.pdf)
* [HTTP Methods for RESTful Services](https://www.restapitutorial.com/lessons/httpmethods.html#:~:text=The%20primary%20or%20most%2Dcommonly,but%20are%20utilized%20less%20frequently.)
* [RESTful Web Services Resources](https://www.restapitutorial.com/resources.html)
* [REST-API-Design-Filtering-Sorting-and-Pagination](https://www.moesif.com/blog/technical/api-design/REST-API-Design-Filtering-Sorting-and-Pagination/)
* [HTTP Status Dogs](https://httpstatusdogs.com/)

## Healthcheck
Quando o projeto for uma API, é requerido que o mesmo implemente um endpoint de `healthcheck`, é recomendado que
o projeto aplique o padrão definido da documentação da guideline, para que o mesmo seja um endpoint inteligente.

Para mais detalhes ver:
* [MadeiraMadeira - Guideline de Healthcheck](https://madeiramadeira.atlassian.net/wiki/spaces/CAR/pages/2226749441/Guidelines+para+projetos#Health-Check)
* [Microsoft - Monitoramento de integridade](https://docs.microsoft.com/pt-br/dotnet/architecture/microservices/implement-resilient-applications/monitor-app-health)
* [Microsoft - Exemplo com ASP.NET Core](https://docs.microsoft.com/pt-br/aspnet/core/host-and-deploy/health-checks?view=aspnetcore-6.0)
* [Testfully - Artigo Health Check](https://testfully.io/blog/api-health-check-monitoring/)

## Docker
Arquivos de docker devem estar na pasta docker, sendo organizados por contexto, exemplos:
* docker/
  * php/
    * Dockerfile
    * entrypoint.sh
  * nginx/
    * logs/*
    * Dockerfile
    * app.conf
    * nginx.conf
  * python/
    * Dockerfile
    * entrypoint.sh

## Editorconfig
É de suma importância que o projeto possua um arquivo de configuração universal para que independente da ferramenta que 
venha a utilizar, o projeto não sofra alterações não desejadas em formatação de arquivos, tipo de quebra de linha etc. 

## Serverless
Quando aplicável ao projeto deixar configurado na raiz do projeto seus respectivos arquivos.

## Logs e monitoramento
É recomendável que o projeto faça de uma interface de log para prover informações de execução do projeto.
Também é recomendável que o mesmo envie logs para a NewRelic e que o projeto esteja instrumentado na mesma.

## Scripts de automação
Os scripts de automação de execução de tarefas de desenvolvimento do projeto devem estar na pasta `scripts`;

## Scripts de configuração de ambiente
É recomendável que os arquivos de configuração de ambiente estejam salvos na pasta `env/`. 
Apenas o arquivo de desenvolvimento (com apontamentos para recursos locais via Docker) e exemplo de arquivo de integração, 
demais arquivos de configuração não devem ser versionados.

Exemplo:
```
./env/development.env
./env/integration.env.example
```
Ou:
```
./env/.env.development
./env/.env.integration.example
```
## Terraform
Os arquivos de Terraform se presentes no projeto deverão estar na pasta `infrastructure`;

## Testes
É muito recomendado que o projeto possua testes, principalmente que sigam a abordagem de contexto, como testes de
componente, unidade e integração. Para projetos que sejam voltados para front-end é interessante que tenhamos o contexto
de usabilidade e teste de componentes da aplicação.

Para mais detalhes ver:
* [Martin Fowler - Testando Microsserviços](https://martinfowler.com/articles/microservice-testing/)

## Arquivos de CI/CD para AWS
Quando o projeto estiver devidamente configurado, é ideal que o mesmo possua os arquivos de CI/CD para AWS.
Estes arquivos devem estar focados principalmente nas ações voltadas para o CD.
Futuramente as tarefas focadas no CI vão ser realizadas via Github.

Exemplo:
```
buildspec.yaml
appspec.yaml
```

## Meta-arquivos
### Project resource
Criar um arquivo de metadados do projeto chamado `.projectrc`.
Este arquivo irá conter dados do projeto como nome, versão e docker network, 
este poderá ser utilizado para outras finalidades de execução de `scripts de automação, 
Além da integração com ferramentas de DX (Developer Experience).

Exemplo:

```dotenv
APP_NAME=project-name-here
APP_VERSION=1.0.0
NETWORK_NAME=docker-network-name-here
```

### dockerignore
Arquivo com as referências de pastas e arquivos que devem ser ignorados pelo docker durante a cópia 
de conteúdo da pasta do projeto.

### gitignore
Arquivo com as referências de pastas e arquivos que devem ser ignorados pelo git durante o desenvolvimento
do projeto.

### docker-compose.yml
Arquivo com configurações para gestão de containers no ambiente de desenvolvimento.


## Referências
* [Guidelines para projetos](https://madeiramadeira.atlassian.net/wiki/spaces/CAR/pages/2226749441/Guidelines+para+projetos)

