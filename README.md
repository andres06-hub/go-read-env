# Read Enviroment variables

This is to know the different ways to read environment variables in Golang

## godotenv

---

### Install

Open the terminal in the project root directory.

```shell
go get github.com/joho/godotenv
```

godotenv provides a Load method to load the env files.

```shell
#Load the .env file in the current directory
godotenv.Load()
#or
godotenv.Load(".env")
```

> _Load method can load multiple env files at once. This also supports yaml._

if you want to read more of godotenv.

- [Web Site](https://pkg.go.dev/github.com/joho/godotenv)
- [GitHub](https://github.com/joho/godotenv)

## Viper Package 

---

> Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application and can handle all types of configuration needs and formats. Reading from JSON, TOML, YAML, HCL, envfile and Java properties config files

### Install

Open the terminal in the project root directory.

```shell
go get github.com/spf13/viper
```

if you want to read more of viper.

- [GitHub](https://github.com/spf13/viper)
