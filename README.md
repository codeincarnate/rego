# Rego

Rego is a simple utility to rename files in a directory.  You give it a regular expression to match files and a pattern to rename them by and it will rename all of them.

Example Usage:

```
rego '(.*)\.tgz' '${1}.tar.gz'
```

This will rename all files in the current directory that end in .tgz to ones that end in .tar.gz.

For the arguments, it's advisable to use single quotes to stop bash from interpreting their contents.