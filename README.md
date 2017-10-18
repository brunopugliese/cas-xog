# cas-xog
Execute XOG files reading and writing in a more easy way

This is a new method of creating XOG files. Using a Driver XML file, you can define with objects you would like to read, write and migrate.


### How to use:

1. Download the [lastest stable release](https://github.com/andreluzz/cas-xog/releases/latest) (cas-xog.exe and xogRead.xml);
2. Create a file called [xogEnv.xml](#xog-environment-example), in the same folder of the cas-xog.exe, to define the environments connections configuration;
3. Create a folder called "drivers" with all [xog Driver xml](#xog-driver-example) files defining the objects you want to read and write;
4. Execute the cas-xog.exe and follow the instructions in the screen.

### Description of Driver types: 

| Type | Description |
| ------ | ------ |
| [objects](#type:-objects) | Used to read and write objects attributes, actions and links  |


#### Type: Objects

| Atribute | Description | Required |
| ------ | ------ | ------ |
| Code | Object code | yes | 
| Path | Path to the file to be saved on the file system | yes | 
| Path | Path to the file to be saved on the file system | yes | 

```sh
<?xml version="1.0" encoding="utf-8"?>
<xogdriver version="1.0">
    <file code="test_subobj" path="test_subobj.xml" type="objects" />
</xogdriver>
```

##### Sub Type: Element
Used to read only the selected elements from the environment

| Atribute | Description | Required |
| ------ | ------ | ------ |
| type | Defines what element to read. Available: **attribute**, **action** and **link** | yes | 
| code | Code of the attribute that you want to include | yes | 

```sh
<?xml version="1.0" encoding="utf-8"?>
<xogdriver version="1.0">
    <file code="test_subobj" path="test_subobj.xml" type="objects">
        <element type="attribute" code="attr_auto_number" />
        <element type="attribute" code="novo_atr" />
        <element type="action" code="tst_run_proc_lnk" />
        <element type="link" code="test_subobj.link_test" />
    </file>
</xogdriver>
```


### XOG Environment example:

```sh
<?xml version="1.0" encoding="utf-8"?>
<xogenvs version="1.0">
    <env name="Development">
        <username>username</username>
        <password>12345</password>
        <endpoint>http://development.server.com</endpoint>
    </env>
    <env name="Quality">
        <username>username</username>
        <password>12345</password>
        <endpoint>http://quality.server.com</endpoint>
    </env>
    <env name="Production">
        <username>username</username>
        <password>12345</password>
        <endpoint>http://production.server.com</endpoint>
    </env>
</xogenvs>
```