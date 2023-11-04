```
   ____                  _
  / __ \__  _____  _____(_)__  _____
 / / / / / / / _ \/ ___/ / _ \/ ___/
/ /_/ / /_/ /  __/ /  / /  __(__  )
\___\_\__,_/\___/_/  /_/\___/____/

```
These Files contain all Queries on the main database.

We are using `Sqlc` to generate the `Go` code. Please refer to [Sqlc Docs](https://docs.sqlc.dev/en/stable/) for more information.

The Files Includes 6 Sections:

1. Selections: `Get?` / `Get?By?`
2. Inserts: `Create?`
3. Joins (insert for ManyToMany): `Add?To?`
4. Updates: `Update?`
5. Deletions: `Delete?`
6. Counts: `Count?`

> Please use PascalCase for function names.
