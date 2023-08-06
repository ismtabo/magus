import caporal from "@caporal/core";
import { dedent } from "ts-dedent";
import dayjs from "dayjs";
import { valid } from "semver";

import createVersionChangelog from "./commands/create-version-changelog.mjs";
import getVersionChangelog from "./commands/get-version-changelog.mjs";

caporal.program
  .option("--changelog <changelog>", "Changelog file path", { default: "CHANGELOG.md", global: true })
  .command(
    "get",
    dedent`Get changelog from version <version>.

    Get changelog sections from version <version> and outputs them.
  
    Example:
    
    $ update-changelog-version get 1.2.3
    
    From:
    
      # Changelog
      
      ## [Unreleased]
      
      ### Added
      
      - Something
      
      ## [1.2.3] - (2021-01-01)
      
      ### Added
      
      - Other thing
      
    Produces:
      
      ## [1.2.3] - (2021-01-01)
      
      ### Added
      
      - Other thing
    `
  )
  .argument("<version>", "Version to get changelog from", { validator: (value) => valid(value) })
  .action(getVersionChangelog)
  .command(
    "create",
    dedent`Update CHANGELOG.md to version <version> with date <date>
  
    Replaces the [Unreleased] section with a new version section and adds a new [Unreleased] section at the top.

    Adds a new version section at the top if it does not exist yet, with name [<version>] - (<date>). Eg: [1.2.3] (2021-01-01)

    Removes empty changes subsections.

    Example:

    $ update-changelog-version 1.2.3 2021-01-01

    From:

      # Changelog

      ## [Unreleased]

      ### Added

      - Something

      ## [1.2.2] - (2020-12-31)

      ### Added

      - Other thing

    Produces:

      # Changelog

      ## [Unreleased]

      ### ...

      ## [1.2.3] - (2021-01-01)

      ### Added

      - Something

      ## [1.2.2] - (2020-12-31)

      ### Added

      - Other thing
    `
  )
  .argument("<version>", "Version to update", { validator: (value) => valid(value) })
  .argument("<date>", "Date to update", {
    default: dayjs(),
    validator: (value) => dayjs(value),
  })
  .action(createVersionChangelog);

caporal.program.run();
