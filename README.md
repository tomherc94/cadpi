<h1>Cluster de Alto Desempenho para Processamento de Imagens (CADPI)</h1>

<p><b>Vídeo:</b> https://youtu.be/bpgFuUg8L-s</p>

<p><b>Requisitos mínimos:</b><br>
8GB de RAM e processador com 6 núcleos.</p>

<p><b>Faça o clone do projeto</b><br>
$git clone https://github.com/tomherc94/cadpi.git</p>

<p><b>Instalar dependências:</b><br>
$sudo apt-get install virtualbox<br>
$sudo apt-get install vagrant</p>


<p><b>Abra um terminal na pasta do projeto</b><br>
$mkdir masterInput masterOutput</p>

<p><b>Copie as imagens para a pasta masterInput</b><br>
$cp -R "suasImagens"/* ./masterInput</p>

<p><b>Suba as máquinas virtuais (VMs)</b><br>
$vagrant up</p>

<p><b>Acesse a VM master</b><br>
$vagrant ssh master</p>

<p><b>Execute o algoritmo para processar as imagens em cluster</b><br>
$go run master.go</p>

<p><b>Ao final do processamento as imagens modificadas estarão em masterOutput</b></p>
