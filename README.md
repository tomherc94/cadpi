<h1>Cluster de Alto Desempenho para Processamento de Imagens (CADPI)</h1>

<p><b>Faça o clone do projeto<b><br>
$git clone https://github.com/tomherc94/cadpi.git</p>

<p>Instalar dependências:<br>
$sudo apt-get install virtualbox<br>
$sudo apt-get install vagrant</p>


<p>Abra um terminal na pasta do projeto<br>
$mkdir masterInput masterOutput</p>

<p>Copie as imagens para a pasta masterInput<br>
$cp -R "suasImagens"/* ./masterInput</p>

<p>Suba as máquinas virtuais (VMs)<br>
$vagrant up</p>

<p>Acesse a VM master<br>
$vagrant ssh master</p>

<p>Execute o algoritmo para processar as imagens em cluster<br>
$go run master.go</p>

<p>Ao final do processamento as imagens modificadas estarão em masterOutput</p>
