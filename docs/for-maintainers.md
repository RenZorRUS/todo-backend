# For To-Do Backend App Maintainers

Before you start contributing to the To-Do Backend App, please make sure you have the following development tools installed on your local machine, which are used to maintain a clean and maintainable codebase:

- [Make](https://www.gnu.org/software/make/manual/make.html) (a build automation tool)
  - Brief explanation of the circumstances requiring the installation of Make and the procedure for installing the `make` tool

    | Operating System | `GNU` `make` Availability | Installation Instructions if Not Available |
    | :-- | :-- | :-- |
    | **Linux** (Ubuntu, Mint, Arch, etc.) | ✅ Pre-installed | No further action needed |
    | **macOS** | ✅ Comes with an older version (`v3.81`) | To upgrade, use `brew install make`.<br>(available as `gmake` to avoid conflicts) |
    | **Windows** | ❌ Absent by default | Options for installation: Chocolatey (`choco install make`), [GnuWin32](https://gnuwin32.sourceforge.net/packages/make.htm), or [MSYS2](https://www.msys2.org) |

  - To list all available `make` commands, run `make help`

- [Golangci-lint](https://golangci-lint.run) (a fast linters runner for Go)
  > **Note:** The `golangci-lint` configuration is defined in the `.golangci.yaml` file by default.
  - Install linter using `make check` command
    - (Or) install linter according to the official [installation instructions](https://golangci-lint.run/welcome/install/)
  - To lint Golang code, run `make lint`
  - To format Golang code, run `make fmt`
  - To integrate `golangci-lint` with [VSCode](https://code.visualstudio.com), follow these steps:
    1. Open VSCode settings:
        - On macOS press `⌘ + ,` or `Ctrl + ,` on Windows and Linux
    2. Navigate to the settings section (where `User` and `Workspace` are listed), under the `search settings` bar, select `Workspace`, then click on the `Open Settings (JSON)` icon located in the top-right corner next to the `search settings` bar.
       - This will open the `settings.json` file in the `.vscode` folder at the root of the project, where you can add settings that will only be applied to files in the current project.
    3. Add to `setting.json` the following lines:

        ```json
        {
          "go.lintTool":"golangci-lint",
          "go.lintFlags": [
            "--allow-parallel-runners"
          ]
        }
        ```

- [Pre-commit](https://pre-commit.com) (manages and maintains multi-language `pre-commit` hooks)
  - Install `pre-commit` using `make check` command
    - (Or) install `pre-commit` according to the [installation instructions](https://paolozaino.wordpress.com/2023/12/15/software-development-installing-pre-commit-to-check-our-code-repositories-and-improve-consistency-across-teams/)
    - Next, execute `make pre-commit-install` to set up the `pre-commit` hooks in your local `.git` repository

## References

...
