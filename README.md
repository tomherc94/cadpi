<h1>Cluster de Alto Desempenho para Processamento de Imagens (CADPI)</h1>

<b>Instalar dependências:</b>

$sudo apt-get install virtualbox

$sudo apt-get install vagrant


<b>Abra um terminal na pasta do projeto</b>

$mkdir masterInput masterOutput

$cp -R <base de imagens>/* ./masterInput

$vagrant up

$vagrant ssh master

$go run master.go
