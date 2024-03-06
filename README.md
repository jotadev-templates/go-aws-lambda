# Preparar ambiente
- Alterar arquivo /infra/variables.tf

  - Incluir conta AWS em `PROVIDE_AWS_ACCOUNT`


- Em terminal, criar variáveis de ambiente
  `PROVIDE_AWS_ACCOUNT=123456`
  `TF_DIR=/path`  (path completo da pasta "crm-lambda/infra")

# Provisionar recursos
- Na pasta app, rodar:
`make provida-infra`

# Destruir recursos
- Na pasta app, rodar:
  `make destroy-infra`

# Projeto desafio
  `projeto-desafio.drawio`
  