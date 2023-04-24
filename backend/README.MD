# Projeto Despensa :: backend

Serviço WEB desenvolvido em Golang com a responsabilidade de envio/disponibilização de dados de estoque de determinada 
despensa.

Projeto do grupo `Alpha` da turma do 5o semestre em ADS da Faculdade Impacta.

### Para subir o projeto localmente:

O arquivo `docker-compose.yaml` possui uma instância local de um banco de dados Postgres. Para executar manualmente:

- Executar `docker-compose up -d` na raiz do backend
- Os endpoints deverão estar disponíveis no http://localhost:8000/
- O swagger com os endpoints e respectivos contratos estão em http://localhost:8000/swagger/

### Testes de integração
WIP

### Documentação
https://despensa.onrender.com/swagger/

### Ambiente produtivo
Para título de apresentação do projeto os endpoints estão disponíveis em https://despensa.onrender.com/api/stock