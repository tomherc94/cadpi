<h1>Cluster de Alto Desempenho para Processamento de Imagens (CADPI)</h1>

<p>Instalar dependências:<br>
$sudo apt-get install virtualbox<br>
$sudo apt-get install vagrant</p>


<p>Abra um terminal na pasta do projeto</p>
$mkdir masterInput masterOutput

<p>Copie as imagens para a pasta masterInput</p>
$cp -R <base de imagens>/* ./masterInput

<p>Suba as máquinas virtuais (VMs)</p>
$vagrant up

<p>Acesse a VM master</p>
$vagrant ssh master

<p>Execute o algoritmo para processar as imagens em cluster</p>
$go run master.go
