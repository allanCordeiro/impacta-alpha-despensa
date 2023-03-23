# Projeto Despensa :: backend

Serviço WEB desenvolvido em Golang com a responsabilidade de envio/disponibilização de dados de estoque de determinada 
despensa.

Projeto do grupo `Alpha` da turma do 5o semestre em ADS da Faculdade Impacta.

### Para subir o projeto localmente:

- Criar um arquivo `.env` à partir do `env.example`
- Incluir a URL do banco de dados postgres na variável `DB_URL`
- Executar `docker-compose up -d` na raiz do backend
- Os endpoints deverão estar disponíveis no http://localhost:8000/api/stock

### Testes de integração
WIP

### Documentação
https://despensa.onrender.com/swagger/

### Ambiente produtivo
Para título de apresentação do projeto os endpoints estão disponíveis em https://despensa.onrender.com/api/stock