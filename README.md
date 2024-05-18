# dddebug
dddebug é um pacote Go simples que fornece uma funcionalidade de depuração semelhante ao dd() do Laravel. Ele é projetado para facilitar a visualização e depuração de dados durante o desenvolvimento, exibindo dados formatados e encerrando a execução do programa.


Instalação
Para instalar o pacote mydebug, execute o seguinte comando:

go get github.com/seuusuario/dddebug

Uso
Importe o pacote dddebug em seu projeto e utilize a função DD para depurar seus dados. A função DD imprime os dados fornecidos em formato JSON com indentação e encerra o programa imediatamente.

Exemplo
Aqui está um exemplo de como usar o pacote mydebug:

package main

import (
    "github.com/seuusuario/mydebug"
)

func main() {
    data := map[string]interface{}{
        "name": "John Doe",
        "age":  30,
        "address": map[string]string{
            "street": "123 Main St",
            "city":   "Springfield",
        },
    }

    mydebug.DD(data) // Isso imprimirá os dados e encerrará o programa
}

Resultado Esperado
Ao executar o exemplo acima, o programa imprimirá:

{
  "name": "John Doe",
  "age": 30,
  "address": {
    "street": "123 Main St",
    "city": "Springfield"
  }

Contribuição
Contribuições são bem-vindas! Se você encontrar um bug ou tiver uma sugestão de melhoria, sinta-se à vontade para abrir uma issue ou enviar um pull request.

Faça um fork do repositório
Crie uma nova branch (git checkout -b feature/nova-funcionalidade)
Commit suas alterações (git commit -am 'Adiciona nova funcionalidade')
Faça o push para a branch (git push origin feature/nova-funcionalidade)
Abra um Pull Request
Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.
