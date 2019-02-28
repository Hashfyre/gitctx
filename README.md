# GITCTX

When committing to Git, your email address and name are attached to the commits. If you are like me, you want to have different configurations based on the project you're working on (e.g., for your job and open-source work).

`gitctx` makes this simple.

- `gitctx create` establishes a new context with name and email
- `gitctx use` hooks up the configuration into a local Git repository

This functionality works on a per-project basis, this is how it works:
```shell
cd <myproject>
gitctx use
# Select the right context

# Validate everything is as you expect
git config --get user.name
git config --get user.email
```

This functionality is backed by adding an [`[include] path=<contextpath>`](https://git-scm.com/docs/git-config#_includes) to `.git/config`.
