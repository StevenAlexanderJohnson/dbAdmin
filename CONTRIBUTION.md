# Contribution Guide

## （；´д｀）ゞ Issues

All branches and PRs should be linked to an issue.
If the issue doesn't already exist, create the issue so that it can be linked.

## o(*￣▽￣*)ブ Pull Request

In order for changes to make it to main, you have to make a PR.
Each PR requires at least one reviewer.
This is to help learn from as well as keep up to date on what changes have been made.


## ┗|｀O′|┛ Code Style Guide

All rules within this section is subject to change.
In the case that it is changed, the code will be updated in a PR with changes to this file.

### (☞ﾟヮﾟ)☞ Go ☜(ﾟヮﾟ☜)

All Go files should be in the base directory and within the main package.
That is unless they pertain to an external resource like an API or service.

External service should be put into it's own package, wrapped in a struct, and imported into main.

Some good examples of external resources would be:

- Auth0
- pokeapi.co
- Stripe

These are resources that are not developed by us and which may change in the future.
We put these resources in their own structs so that we may unit test our code.

Odds are that tehse are not going to be used, but if they do we have a plan.

#### File Names

Files should follow ```camelCase```.
If possible make one word file names that accurately describes what the logic does.
If it gets too hard to create a file name in less than two words, or if words in the file name are repeated frequently, it may be time to create a package.

#### Variables

Varaibles should follow ```camelCase``` unless they are to become public, then they should use ```PascalCase```.

#### Functions

Seeing all the logic is within main there is no need for upper-case function names.
Therefore use ```camelCase```.
If the function is within a package other than main and is to be used by main use ```PascalCase```.

#### Structs and Struct Methods

Structs should use ```PascalCase```.
Struct Methods should use ```PascalCase```.

### Constants

Constants should be ```UPPER_SNAKE``` case.

### (╯°□°）╯︵ ┻━┻ JavaScript

### File Names

Svelte files should be capitalized.
In the case that the file name is multiple words use ```PascalCase```.

Folder names should be lowercase.
Most folders that will be used are already created.

### Variables and Functions

Both variables and functions should be ```camelCase```.
I'm going to say if you're exporting either make it ```PascalCase``` so we can know that it doesn't exist within the current file.

### Classes

I don't see a need for classes in this project, but if you were to create one, it should probably be in it's own file and ```PascalCase```.

### Constants

Constants should be ```UPPER_SNAKE``` case, even if they are exported from the file and imported elsewhere.