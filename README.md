## Steampipe Plugin Cookiecutter

Quickly create a steampipe plugin repository with a cookiecutter template.

### Installation

```bash
pip install cookiecutter
```

### Usage

```bash
cookiecutter github.com/mr-destructive/plugin-cookiecutter

OR

pipx run cookiecutter github.com/mr-destructive/plugin-cookiecutter
```

Fill in some details like

- Plugin name(in lower case)
- Plugin standard name (in camel case)
- Plugin description (optional)
- Github Username (defaults to turbot)
- Plugin Table name (in lower case)
- Plugin Table standard name (in camel case)


References:

- [Steampipe Plugin Checklist](https://steampipe.io/docs/develop/plugin-release-checklist)
