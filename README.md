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

# pre-commit

1. Install pre-commit on your system. You can install it with pip:

```shell
pip install pre-commit
```

2. Create a .pre-commit-config.yaml file in the root of your project and add the hooks you want to run in the pre-commit. Here is an example:

```yaml
repos:
  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v1.3.0
    hooks:
      - id: go-fmt
      - id: go-mod-tidy
      - id: go-vet
```

- Know more options [See more](https://pre-commit.com/hooks.html)

3. Run the following command to install the hooks in your project:

```shell
pre-commit install
```

This will install the hooks in your repository and associate them with the git commit command.

From now on, when you do a git commit, the hooks configured in .pre-commit-config.yaml will be executed automatically before the commit is performed. If any of the hooks fail, the commit will not be performed.

4. Optionally, you can manually execute the pre-commit hooks at any time by running the command:

```shell
pre-commit run --all-files
```

This will execute all hooks configured in .pre-commit-config.yaml in all repository files.
