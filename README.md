# sandra_cli
Command Line Interface

Sandra CLI es una interfaz de línea de comandos o interfaz de línea de órdenes es un tipo de interfaz de usuario de computadora que permite a los usuarios dar instrucciones a algún programa informático o al sistema operativo por medio de una línea de texto simple. Esta herramienta de interfaz de comando que permite a los usuarios interactuar con la plataforma de Sandra Server para arquitecturas empresariales que estará operativa para su organización ampliando la relación entre sistemas de diferentes naturaleza, así como la colaboración entre sus tecnologías adyacentes.

##### SINOPSIS
sandra_cli [command] [flags] 

<p align="center">
   <img src="https://raw.githubusercontent.com/code-epic/sandra_cli/main/img/sandra_cli.jpg" width="500px;"/>
</p>

##### OPCIONES

El comando config permite generar el archivo de token para crear la instalacion de Sandra Server. podra acceder a nuestro sitio https://code-epic.com/register y loguearte con tu cuenta de Gmail o crear una luego acceder a la autenticacion y copiar el codigo SHA256 que te genera la plataforma.

config

```
--file key | crt la llave para descargar e iniciar Sandra Server
--add token que te genera la WEB de register
```

```sh
sandra_cli config --id sha256 --file crt --add {token.code-epic.com}
```

version

Ver la version del software 

```bash
sandra_cli --version
```


install 

- service: Este comando permite instalar *Sandra Service* e iniciar proceso dentro de systemctl 
- tools: los comando de lineas sandra_dwn, sandra_tcp, sandra_scanf
- data-base: podra instalar directamente las base de daton internas y externas
	
```sh
sandra_cli install --option service
sandra_cli install --option tools
sandra_cli install --option data-base
```


update 

- service: Este comando permite instalar *Sandra Service* e iniciar proceso dentro de systemctl 
- tools: los comando de lineas sandra_dwn, sandra_tcp, sandra_scanf
- data-base: podra instalar directamente las base de daton internas y externas

```sh
sandra_cli update --option service
sandra_cli update --option tools
sandra_cli update --option data-base
```

newProject create

```
-t --type web | movil
-l --language angular | react | vuejs
-n --name nombre del producto
-a --author firma del documento
```

```sh
sandra_cli newProject create -t web -l angular -n nombre_proyecto -a autor
```


