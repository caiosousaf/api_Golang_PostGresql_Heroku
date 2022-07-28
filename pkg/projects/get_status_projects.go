package projetos

import (
	"github.com/caiosousaf/api_Golang_PostGresql_Heroku/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetStatusProjects(c *gin.Context) {

	status := c.Param("status")

	var projeto []models.Projeto

	if result := h.DB.Raw("select * from projetos where status = ?", status).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &projeto)

}





/*body {
  font-family: Arial, Helvetica, sans-serif;
  background-image: linear-gradient(
    to right,
    rgb(105, 190, 239),
    rgb(40, 58, 68)
  );
}
.box {
  color: white;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: rgba(0, 0, 0, 0.6);
  padding: 15px;
  border-radius: 15px;
  width: 20%;
}
fieldset {
  border: 3px solid rgb(147, 230, 118);
}
legend {
  border: 1px solid dodgerblue;
  padding: 10px;
  text-align: center;
  background-color: dodgerblue;
  border-radius: 8px;
}
.inputBox {
  position: relative;
}
.inputUser {
  background: none;
  border: none;
  border-bottom: 1px solid white;
  outline: none;
  color: white;
  font-size: 15px;
  width: 100%;
  letter-spacing: 2px;
}
.labelInput {
  position: absolute;
  top: 0px;
  left: 0px;
  pointer-events: none;
  transition: 0.5s;
}
.inputUser:focus ~ .labelInput,
 .labelInput {
  top: -20px;
  font-size: 12px;
  color: dodgerblue;
}

#calcular {
  background-image: linear-gradient(
    to right,
    rgb(0, 92, 197),
    rgb(90, 20, 220)
  );
  width: 100%;
  border: none;
  padding: 15px;
  color: white;
  font-size: 15px;
  cursor: pointer;
  border-radius: 10px;
}
#calcular:hover {
  background-image: linear-gradient(
    to right,
    rgb(0, 80, 172),
    rgb(80, 19, 195)
  );
}

.result {
  display: flex;
  margin-top: 20px;
  align-items: center;
  width: 300px;
  height: 150px;
  border-radius: 5px;
  font: italic 1.5rem serif;
  color: dodgerblue;
  padding: 20px;
  box-sizing: border-box;
  user-select: none;
}

















<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>IMC</title>
    <link rel="stylesheet" href="style.css">
  </head>
  <body>
      <script src="./imc.js"></script>
    <div class="box">

        <fieldset>
            <legend><b>Calculadora - IMC</b></legend>
       <br>
        <div class="inputBox">
            <input
              type="text"
              name="nome"
              id="nome"
              class="inputUser"
  
            />
            <label for="nome" class="labelInput">Nome: </label>
        </div>
        <br /> <br />

        <div class="inputBox">
            
            <input type="number" name="altura" id="altura" class="inputUser">
            <label for="altura" class="labelInput">Altura</label>
        </div>
        <br /> <br />
        <div class="input">
            <div class="inputBox">
                <input
                  type="number"
                  name="telefone"
                  id="peso"
                  class="inputUser"

                />
                <label for="telefone" class="labelInput">Peso</label>
        </div>
        

    <button id='calcular' onclick="imc()">Calcular</button>
    <div class="result" id='resultado'></div>
    </div>
  </body>
</html>





























function imc () {
    const nome = document.getElementById('nome').value;
    const altura = document.getElementById('altura').value;
    const peso = document.getElementById('peso').value;
    const resultado = document.getElementById('resultado');
 
    if (nome !== '' && altura !== '' && peso !== '') {

        const valorIMC = (peso / (altura * altura)).toFixed(1);

        let classificacao = '';

        if (valorIMC < 18.5){
            classificacao = 'abaixo do peso.';
        }else if (valorIMC < 25) {
            classificacao = 'com peso ideal. Parabéns!!!';
        }else if (valorIMC < 30){
            classificacao = 'levemente acima do peso.';
        }else if (valorIMC < 35){
            classificacao = 'com obesidade grau I.';
        }else if (valorIMC < 40){
            classificacao = 'com obesidade grau II';
        }else {
            classificacao = 'com obesidade grau III. Cuidado!!';
        }

        resultado.textContent = `${nome} seu IMC é ${valorIMC} e você está ${classificacao}`;
        
    }else {
        resultado.textContent = 'Preencha todos os campos!!!';
    }

}


*/

