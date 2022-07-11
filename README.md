## What we have.

We run our hooks on every commit to automatically point out issues in code such as terraform fmt, terraform docs, trailing whitespace, detect-aws-credentials, check-merge-conflict, detect-private-key. 

## If you want to use pre-commit githooks you need to have installed pre-commit see here how to install.

<!-- markdownlint-disable no-inline-html -->

* [`pre-commit`](https://pre-commit.com/#install)  

<!-- markdownlint-enable no-inline-html -->

## You need to have .pre-commit-config.yaml file in the top of your repo

```bash
repos:
- repo: https://github.com/pre-commit/pre-commit-hooks
  rev: v4.2.0
  hooks:
    - id: check-added-large-files
    - id: check-merge-conflict
    - id: check-vcs-permalinks
    - id: end-of-file-fixer
    - id: trailing-whitespace
      args: [--markdown-linebreak-ext=md]
      exclude: CHANGELOG.md
    - id: check-yaml
    - id: check-merge-conflict
    - id: check-executables-have-shebangs
    - id: check-case-conflict
    - id: mixed-line-ending
      args: [--fix=lf]
    - id: detect-aws-credentials
      args: ['--allow-missing-credentials']
    - id: detect-private-key
```

### If you want to install the pre-commit hook globally run this commands.

```bash
DIR=~/.git-template
git config --global init.templateDir ${DIR}
pre-commit init-templatedir -t pre-commit ${DIR}
```

### If you want to install the pre-commit hook manually in one folder run this.

```bash
git config core.hooksPath githooks

OR 

git config core.hooksPath githooks-detect-secrets

OR 

git config core.hooksPath githooks-terraform-docs

OR 

git config core.hooksPath githooks-terraform-fmt
```


### . If you want to run Terraform-docs you need to change folder name with yours.

Go https://github.com/dasmeta/pre-commit-terraform/blob/master/githooks-terraform-docs

change the " modules/*/ " path with yours it will check all folder under modules/ directory:

```bash
for i in `ls -d modules/*/`;do
```