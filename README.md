## How to Install.

<!-- markdownlint-disable no-inline-html -->

* [`pre-commit`](https://pre-commit.com/#install)

<!-- markdownlint-enable no-inline-html -->

### 2. Install the pre-commit hook globally.

```bash
DIR=~/.git-template
git config --global init.templateDir ${DIR}
pre-commit init-templatedir -t pre-commit ${DIR}
```
### 3. Config for GitHooks.

```bash
git config core.hooksPath githooks
```

### 4. If you want to run Terraform-fmt and Terraform-docs when you commiting 

Go https://github.com/dasmeta/pre-commit-terraform/blob/master/githooks/pre-commit-example

change the " modules/*/ " path with yours it will check all folder under modules/ directory:

```bash
for i in `ls -d modules/*/`;do
```

and add in your repo `githooks` -> `pre-commit`


